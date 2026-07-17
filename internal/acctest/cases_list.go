package acctest

import (
	"fmt"
	"os"
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

// ListCase is a per-test-case list/query acceptance configuration, generated
// from the legacy provider's hand-written list acceptance tests. Each case maps
// 1:1 to a legacy TestAcc<Object>List_<Case> function and is stored as a
// labelled `case "<name>" { ... }` block in the merged
// testdata/<package>/<object>/<backend>_lists.tfvars file.
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
}

// RunListCases loads every `case "<name>" { ... }` block from the merged
// tfvars file at testdata/<fileRelPath> and runs each as an independent
// two-step subtest: step 1 creates the resource, step 2 runs a list query
// and asserts at least one result is returned.
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

	t.Logf("Case %q create HCL:\n%s", lc.Name, providerConfig+"\n"+resourceHCL)
	t.Logf("Case %q list HCL:\n%s", lc.Name, providerConfig+"\n"+listHCL)

	resourceAddr := resourceType + ".test"

	queryChecks := buildQueryChecks(resourceAddr, lc)

	tc := resource.TestCase{
		ProtoV6ProviderFactories: ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: providerConfig + "\n" + resourceHCL,
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
// Filtered cases assert an exact count of 1 and validate each filter attribute
// on the returned resource object. List-all cases (empty FilterType) only
// assert that at least one result is present.
func buildQueryChecks(resourceAddr string, lc *ListCase) []querycheck.QueryResultCheck {
	if lc.FilterType == "" {
		return []querycheck.QueryResultCheck{
			querycheck.ExpectLengthAtLeast(resourceAddr, 1),
		}
	}

	checks := []querycheck.QueryResultCheck{
		querycheck.ExpectLength(resourceAddr, 1),
	}

	// ExpectResourceKnownValues only works for simple filters (not EA/tag
	// filters where the attribute path in the query result is structured
	// differently from regular fields).
	if lc.FilterType == "filters" {
		var knownChecks []querycheck.KnownValueCheck
		for _, key := range lc.FilterOrder {
			refPath := lc.Filters[key]
			val := resolveStepValue(refPath, lc.Step)
			if val == "" {
				continue
			}
			knownChecks = append(knownChecks, querycheck.KnownValueCheck{
				Path:       refPathToTFJSONPath(refPath),
				KnownValue: knownvalue.StringExact(val),
			})
		}
		if len(knownChecks) > 0 {
			checks = append(checks, querycheck.ExpectResourceKnownValues(
				resourceAddr, nil, knownChecks,
			))
		}
	}

	return checks
}

// buildListBlock renders the `list "<resourceType>" "test" { ... }` HCL for a
// list query step. An empty FilterType means list-all (limit only, no config
// block); otherwise a config block with filters/extattrfilters is emitted.
// Filter values are resolved from the materialized step config (not resource
// references) because the Query step config contains only the list block.
func buildListBlock(resourceType string, lc *ListCase) string {
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("list %q \"test\" {\n", resourceType))
	sb.WriteString("  provider = infoblox\n")

	if lc.FilterType == "" {
		sb.WriteString("  limit = 5\n")
	} else {
		sb.WriteString("  include_resource = true\n")
		sb.WriteString("  config {\n")

		sb.WriteString(fmt.Sprintf("    %s = {\n", lc.FilterType))
		for _, key := range lc.FilterOrder {
			refPath := lc.Filters[key]
			val := resolveStepValue(refPath, lc.Step)
			if val != "" {
				sb.WriteString(fmt.Sprintf("      %s = %q\n", key, val))
			}
		}
		sb.WriteString("    }\n")
		sb.WriteString("  }\n")
	}

	sb.WriteString("}\n")
	return sb.String()
}

// resolveStepValue looks up the materialized value for a resource attribute
// reference path (e.g. "nios.network", "nios.ext_attrs.Site") in the step.
func resolveStepValue(refPath string, step CaseStep) string {
	parts := strings.SplitN(refPath, ".", 2)
	if len(parts) < 2 {
		return ""
	}
	section, rest := parts[0], parts[1]

	var m map[string]any
	switch section {
	case "nios":
		m = step.NIOS
	case "uddi":
		m = step.UDDI
	default:
		m = step.Common
	}

	// Handle simple path like "network" and nested like "ext_attrs.Site"
	attrParts := strings.SplitN(rest, ".", 2)
	raw, ok := m[attrParts[0]]
	if !ok {
		return ""
	}
	if len(attrParts) == 1 {
		if s, ok := raw.(string); ok {
			return s
		}
		return ""
	}
	// Nested map (e.g. ext_attrs.Site)
	nested, ok := raw.(map[string]any)
	if !ok {
		return ""
	}
	if s, ok := nested[attrParts[1]].(string); ok {
		return s
	}
	return ""
}

// refPathToTFJSONPath converts a resource attribute ref path like "nios.network"
// or "nios.ext_attrs.Site" into a tfjsonpath.Path for use in KnownValueChecks.
func refPathToTFJSONPath(refPath string) tfjsonpath.Path {
	parts := strings.Split(refPath, ".")
	p := tfjsonpath.New(parts[0])
	for _, part := range parts[1:] {
		p = p.AtMapKey(part)
	}
	return p
}

// materialize replaces {{placeholder}} tokens with concrete runtime values,
// keeping each distinct placeholder consistent across the step config.
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

	var replace func(v any) any
	replace = func(v any) any {
		switch t := v.(type) {
		case string:
			s := t
			for _, ph := range placeholderPattern.FindAllString(s, -1) {
				s = strings.ReplaceAll(s, ph, value(ph))
			}
			return s
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

	for _, block := range content.Blocks {
		if block.Type != "step" {
			continue
		}
		// Peek at query/display attributes and the optional filter block inside
		// this step. PartialContent only consumes what matches — remaining
		// content is still available for parseCaseStep on the same body.
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
