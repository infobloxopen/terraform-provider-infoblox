package acctest

import (
	"fmt"
	"os"
	"regexp"
	"sort"
	"strings"
	"testing"

	goversion "github.com/hashicorp/go-version"
	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclparse"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/querycheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
	"github.com/hashicorp/terraform-plugin-testing/tfversion"
)

// ListCase is the per-subtest configuration for a list query acceptance test.
// Each case maps to a `case "<name>" { ... }` block in <backend>_lists.hcl.
type ListCase struct {
	Name         string
	Backend      string
	Skip         bool
	SkipReason   string
	MinTFVersion string            // minimum Terraform version string (e.g. "1.14.0")
	FilterType   string            // filters | ext_attr_filters | tag_filters; empty = list-all
	Filters      map[string]string // filter key -> resource attribute path (e.g. "network" -> "nios.network")
	FilterOrder  []string          // filter keys in deterministic order
	Step         CaseStep          // resource-create step
	// PrerequisitesHCL is prepended to both the create and query steps so the resource stays alive.
	PrerequisitesHCL string
}

// RunListCases loads all `case` blocks from testdata/<fileRelPath> and runs each as a two-step subtest:
// step 1 creates the resource, step 2 issues a list query and asserts at least one result is returned.
func RunListCases(t *testing.T, resourceType, fileRelPath string, checksByBackend map[string]CheckFuncs) {
	t.Helper()
	path := GetTestdataPath(fileRelPath)
	cases, err := loadListCases(path)
	if err != nil {
		cases = nil
	}

	if len(cases) == 0 {
		t.Skipf("no list cases at %s: %v", path, err)
		return
	}

	for _, lc := range cases {
		lc := lc
		t.Run(lc.Name, func(t *testing.T) {
			if lc.Skip {
				t.Skipf("Skipped: %s", lc.SkipReason)
				return
			}

			checks, ok := checksByBackend[lc.Backend]
			if !ok {
				t.Skipf("no check functions registered for backend %q", lc.Backend)
				return
			}

			PreCheck(t, lc.Backend)
			lc.materialize()
			runListCase(t, resourceType, lc, checks)
		})
	}
}

func runListCase(t *testing.T, resourceType string, lc *ListCase, checks CheckFuncs) {
	t.Helper()

	providerConfig := ProviderConfigHCL(lc.Backend)
	resourceHCL := buildCaseHCL(resourceType, "test", lc.Backend, lc.Step)
	listHCL := buildListBlock(resourceType, lc)

	createConfig := providerConfig + "\n" + lc.PrerequisitesHCL + "\n" + resourceHCL

	t.Logf("Case %q create HCL:\n%s", lc.Name, createConfig)
	t.Logf("Case %q list HCL:\n%s", lc.Name, listHCL)

	resourceAddr := resourceType + ".test"

	queryChecks := buildQueryChecks(resourceAddr, lc)

	tc := resource.TestCase{
		ProtoV6ProviderFactories: ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: createConfig,
				Check:  resource.ComposeTestCheckFunc(checks.Exists(resourceAddr)),
			},
			{
				Query:             true,
				Config:            listHCL,
				QueryResultChecks: queryChecks,
			},
		},
	}

	if lc.MinTFVersion != "" {
		if v, err := goversion.NewVersion(lc.MinTFVersion); err == nil {
			tc.TerraformVersionChecks = []tfversion.TerraformVersionCheck{
				tfversion.SkipBelow(v),
			}
		}
	}

	resource.Test(t, tc)
}

// buildQueryChecks assembles QueryResultChecks for the list query step.
// Filtered cases assert an exact count of 1 and verify resource attribute values.
// List-all (basic) cases assert at least one result — known-value checks are skipped
// because multiple pre-existing objects may be returned and include_resource is not set.
func buildQueryChecks(resourceAddr string, lc *ListCase) []querycheck.QueryResultCheck {
	var checks []querycheck.QueryResultCheck

	if lc.FilterType == "" {
		checks = append(checks, querycheck.ExpectLengthAtLeast(resourceAddr, 1))
	} else {
		checks = append(checks, querycheck.ExpectLength(resourceAddr, 1))
		// Filtered cases use include_resource=true and return exactly 1 result, so
		// known-value checks are safe to add.
		knownChecks := buildStepKnownValueChecks(lc)
		if len(knownChecks) > 0 {
			checks = append(checks, querycheck.ExpectResourceKnownValues(
				resourceAddr, nil, knownChecks,
			))
		}
	}

	return checks
}

// buildStepKnownValueChecks returns KnownValueChecks for all resolvable scalar fields in the create step.
// Fields that are resource references (RawExpr) are skipped; only literal values are verified.
func buildStepKnownValueChecks(lc *ListCase) []querycheck.KnownValueCheck {
	var knownChecks []querycheck.KnownValueCheck
	seen := make(map[string]bool)

	addField := func(refPath string) {
		if seen[refPath] {
			return
		}
		val := resolveStepValue(refPath, lc)
		if val == "" {
			return
		}
		seen[refPath] = true
		knownChecks = append(knownChecks, querycheck.KnownValueCheck{
			Path:       refPathToTFJSONPath(refPath),
			KnownValue: knownvalue.StringExact(val),
		})
	}

	addSection := func(prefix string, m map[string]any) {
		for k, v := range m {
			if !isScalarValue(v) {
				continue
			}
			path := k
			if prefix != "" {
				path = prefix + "." + k
			}
			addField(path)
		}
	}

	addSection("", lc.Step.Common)
	switch lc.Backend {
	case "nios":
		addSection("nios", lc.Step.NIOS)
	case "uddi":
		addSection("uddi", lc.Step.UDDI)
	}

	return knownChecks
}

// buildListBlock renders the `list "<resourceType>" "test" { ... }` HCL for the query step.
// An empty FilterType produces a list-all block; otherwise a filtered config block is emitted.
func buildListBlock(resourceType string, lc *ListCase) string {
	var sb strings.Builder
	fmt.Fprintf(&sb, "list %q \"test\" {\n", resourceType)
	sb.WriteString("  provider = infoblox\n")

	if lc.FilterType == "" {
		sb.WriteString("  limit = 5\n")
	} else {
		sb.WriteString("  include_resource = true\n")
		sb.WriteString("  config {\n")

		fmt.Fprintf(&sb, "    %s = {\n", lc.FilterType)
		for _, key := range lc.FilterOrder {
			refPath := lc.Filters[key]
			val := resolveStepValue(refPath, lc)
			if val != "" {
				fmt.Fprintf(&sb, "      %s = %q\n", key, val)
			}
		}
		sb.WriteString("    }\n")
		sb.WriteString("  }\n")
	}

	sb.WriteString("}\n")
	return sb.String()
}

// resolveStepValue returns the materialized value for a step attribute ref path (e.g. "nios.network").
func resolveStepValue(refPath string, lc *ListCase) string {
	parts := strings.SplitN(refPath, ".", 2)
	if len(parts) < 2 {
		return ""
	}
	section, rest := parts[0], parts[1]

	var m map[string]any
	switch section {
	case "nios":
		m = lc.Step.NIOS
	case "uddi":
		m = lc.Step.UDDI
	default:
		m = lc.Step.Common
	}

	// Handle simple path like "network" and nested like "ext_attrs.Site"
	attrParts := strings.SplitN(rest, ".", 2)
	raw, ok := m[attrParts[0]]
	if !ok {
		return ""
	}
	if len(attrParts) == 1 {
		return stepScalarValue(raw, lc.PrerequisitesHCL)
	}
	// Nested map (e.g. ext_attrs.Site)
	nested, ok := raw.(map[string]any)
	if !ok {
		return ""
	}
	return stepScalarValue(nested[attrParts[1]], lc.PrerequisitesHCL)
}

// stepScalarValue renders a step value as a filter string.
// String literals pass through; RawExpr is resolved against prereqHCL since the query step holds only the list block.
// Non-string scalars yield "".
func stepScalarValue(raw any, prereqHCL string) string {
	switch v := raw.(type) {
	case string:
		return v
	case RawExpr:
		return resolveRawExpr(string(v), prereqHCL)
	default:
		return ""
	}
}

// prereqRefPattern matches interpolated ("${...}") or bare attribute references in HCL expressions.
var prereqRefPattern = regexp.MustCompile(`\$\{\s*([a-zA-Z_][\w-]*(?:\.[\w-]+)+)\s*\}|([a-zA-Z_][\w-]*(?:\.[\w-]+)+)`)

// resolveRawExpr reduces a raw HCL expression to a literal by substituting prerequisite attribute refs.
// Returns "" if any referenced attribute is not a literal in the prerequisites.
func resolveRawExpr(raw, prereqHCL string) string {
	expr := strings.TrimSpace(raw)
	if len(expr) >= 2 && strings.HasPrefix(expr, `"`) && strings.HasSuffix(expr, `"`) {
		expr = expr[1 : len(expr)-1]
	}

	matches := prereqRefPattern.FindAllStringSubmatch(expr, -1)
	if len(matches) == 0 {
		return ""
	}

	literals := prereqAttrLiterals(prereqHCL)
	for _, m := range matches {
		ref := m[1]
		if ref == "" {
			ref = m[2]
		}
		val, ok := literals[ref]
		if !ok {
			return ""
		}
		expr = strings.ReplaceAll(expr, m[0], val)
	}
	return expr
}

// prereqAttrLiterals extracts literal attribute values from materialized prerequisites HCL,
// keyed by Terraform reference path (e.g. "infoblox_zone_auth.test.nios.fqdn").
func prereqAttrLiterals(prereqHCL string) map[string]string {
	literals := make(map[string]string)
	if strings.TrimSpace(prereqHCL) == "" {
		return literals
	}

	file, diags := hclparse.NewParser().ParseHCL([]byte(prereqHCL), "prerequisites.hcl")
	if diags.HasErrors() {
		return literals
	}
	content, _, _ := file.Body.PartialContent(&hcl.BodySchema{
		Blocks: []hcl.BlockHeaderSchema{
			{Type: "resource", LabelNames: []string{"type", "name"}},
		},
	})

	for _, block := range content.Blocks {
		prefix := block.Labels[0] + "." + block.Labels[1]
		attrs, _ := block.Body.JustAttributes()
		for name, attr := range attrs {
			val, valDiags := attr.Expr.Value(nil)
			if valDiags.HasErrors() || !val.IsWhollyKnown() {
				continue
			}
			flattenLiteral(literals, prefix+"."+name, ctyToGo(val))
		}
	}
	return literals
}

// flattenLiteral recursively records scalar values at their full reference path, skipping lists and nulls.
func flattenLiteral(out map[string]string, path string, v any) {
	switch t := v.(type) {
	case nil, []any:
		return
	case map[string]any:
		for k, vv := range t {
			flattenLiteral(out, path+"."+k, vv)
		}
	default:
		out[path] = fmt.Sprintf("%v", t)
	}
}

// refPathToTFJSONPath converts a ref path like "nios.ext_attrs.Site" into a tfjsonpath.Path.
func refPathToTFJSONPath(refPath string) tfjsonpath.Path {
	parts := strings.Split(refPath, ".")
	p := tfjsonpath.New(parts[0])
	for _, part := range parts[1:] {
		p = p.AtMapKey(part)
	}
	return p
}

// materialize replaces {{placeholder}} tokens with stable values across the step config and prerequisites.
func (lc *ListCase) materialize() {
	cache := make(map[string]string)

	value := func(ph string) string {
		if v, ok := cache[ph]; ok {
			return v
		}
		v := ResolvePlaceholder(ph)
		cache[ph] = v
		return v
	}

	substitute := func(s string) string {
		for _, ph := range placeholderPattern.FindAllString(s, -1) {
			s = strings.ReplaceAll(s, ph, value(ph))
		}
		return s
	}

	var replace func(v any) any
	replace = func(v any) any {
		switch t := v.(type) {
		case RawExpr:
			// RawExpr values may still contain placeholders; substitute and keep raw.
			return RawExpr(substitute(string(t)))
		case string:
			return substitute(t)
		case map[string]any:
			for k, vv := range t {
				t[k] = replace(vv)
			}
			return t
		case []any:
			for i, vv := range t {
				t[i] = replace(vv)
			}
			return t
		default:
			return v
		}
	}

	lc.PrerequisitesHCL = substitute(lc.PrerequisitesHCL)

	for k, v := range lc.Step.Common {
		lc.Step.Common[k] = replace(v)
	}
	for k, v := range lc.Step.NIOS {
		lc.Step.NIOS[k] = replace(v)
	}
	for k, v := range lc.Step.UDDI {
		lc.Step.UDDI[k] = replace(v)
	}
}

func loadListCases(path string) ([]*ListCase, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read case file %s: %w", path, err)
	}

	parser := hclparse.NewParser()
	file, diags := parser.ParseHCL(data, path)
	if diags.HasErrors() {
		return nil, fmt.Errorf("failed to parse HCL: %s", diags.Error())
	}

	content, _, diags := file.Body.PartialContent(&hcl.BodySchema{
		Blocks: []hcl.BlockHeaderSchema{
			{Type: "case", LabelNames: []string{"name"}},
		},
	})
	if diags.HasErrors() {
		return nil, fmt.Errorf("failed to decode cases: %s", diags.Error())
	}

	var cases []*ListCase
	for _, block := range content.Blocks {
		lc, err := parseListCaseBody(block.Body, data)
		if err != nil {
			return nil, fmt.Errorf("case %q: %w", block.Labels[0], err)
		}
		lc.Name = block.Labels[0]
		cases = append(cases, lc)
	}

	if len(cases) == 0 {
		return nil, fmt.Errorf("no case blocks found in %s", path)
	}

	sort.Slice(cases, func(i, j int) bool { return cases[i].Name < cases[j].Name })
	return cases, nil
}

func parseListCaseBody(body hcl.Body, src []byte) (*ListCase, error) {
	content, _, diags := body.PartialContent(&hcl.BodySchema{
		Attributes: []hcl.AttributeSchema{
			{Name: "backend"},
			{Name: "skip"},
			{Name: "skip_reason"},
			{Name: "min_tf_version"},
			{Name: "prerequisites_hcl"},
		},
		Blocks: []hcl.BlockHeaderSchema{
			{Type: "step"},
		},
	})
	if diags.HasErrors() {
		return nil, fmt.Errorf("failed to decode case: %s", diags.Error())
	}

	lc := &ListCase{Filters: make(map[string]string)}
	if attr, ok := content.Attributes["backend"]; ok {
		val, _ := attr.Expr.Value(nil)
		lc.Backend = val.AsString()
	}
	if attr, ok := content.Attributes["skip"]; ok {
		val, _ := attr.Expr.Value(nil)
		lc.Skip = val.True()
	}
	if attr, ok := content.Attributes["skip_reason"]; ok {
		val, _ := attr.Expr.Value(nil)
		lc.SkipReason = val.AsString()
	}
	if attr, ok := content.Attributes["min_tf_version"]; ok {
		val, _ := attr.Expr.Value(nil)
		lc.MinTFVersion = val.AsString()
	}
	if attr, ok := content.Attributes["prerequisites_hcl"]; ok {
		val, _ := attr.Expr.Value(nil)
		lc.PrerequisitesHCL = val.AsString()
	}

	for _, block := range content.Blocks {
		if block.Type != "step" {
			continue
		}
		// Peek at query/filter metadata; PartialContent leaves remaining content available for parseCaseStep.
		stepMeta, _, _ := block.Body.PartialContent(&hcl.BodySchema{
			Attributes: []hcl.AttributeSchema{
				{Name: "query"},
				{Name: "provider"},
				{Name: "include_resource"},
				{Name: "limit"},
			},
			Blocks: []hcl.BlockHeaderSchema{
				{Type: "filter"},
			},
		})

		isQuery := false
		if attr, ok := stepMeta.Attributes["query"]; ok {
			val, _ := attr.Expr.Value(nil)
			isQuery = val.True()
		}

		if isQuery {
			for _, fb := range stepMeta.Blocks {
				if fb.Type == "filter" {
					parseListFilterBlock(fb.Body, lc)
				}
			}
		} else {
			st, err := parseCaseStep(block.Body, src)
			if err != nil {
				return nil, err
			}
			lc.Step = st
		}
	}

	return lc, nil
}

func parseListFilterBlock(body hcl.Body, lc *ListCase) {
	content, _, _ := body.PartialContent(&hcl.BodySchema{
		Attributes: []hcl.AttributeSchema{
			{Name: "type"},
			{Name: "values"},
		},
	})

	if attr, ok := content.Attributes["type"]; ok {
		val, _ := attr.Expr.Value(nil)
		lc.FilterType = val.AsString()
	}
	if attr, ok := content.Attributes["values"]; ok {
		val, _ := attr.Expr.Value(nil)
		lc.Filters = ctyMapToStringMap(val)
		for k := range lc.Filters {
			lc.FilterOrder = append(lc.FilterOrder, k)
		}
		sort.Strings(lc.FilterOrder)
	}
}
