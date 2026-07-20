package acctest

import (
	"fmt"
	"os"
	"regexp"
	"sort"
	"strings"
	"testing"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclparse"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/core"
)

// ExtractNIOSRef strips the leading object-type segment from a full NIOS _ref
// (e.g. "record:a/ZG5z..." -> "ZG5z..."). Acceptance check helpers that call
// the raw NIOS SDK must use this, since the SDK re-applies the type prefix.
func ExtractNIOSRef(ref string) string {
	return core.ExtractNIOSRef(ref)
}

// CaseStep is a single Terraform step within a generated resource case.
type CaseStep struct {
	Common map[string]any
	NIOS   map[string]any
	UDDI   map[string]any
	Checks map[string]string
}

// ResourceCase is a per-test-case acceptance configuration, generated from the
// legacy provider's hand-written acceptance tests. Each case maps 1:1 to a
// legacy TestAcc<Object>Resource_<Case> function and is stored as a labelled
// `case "<name>" { ... }` block in the merged
// testdata/<package>/<object>/<backend>_resources.tfvars file.
type ResourceCase struct {
	Name               string
	Backend            string
	Skip               bool
	SkipReason         string
	Disappears         bool
	ExpectNonEmptyPlan bool
	PrerequisitesHCL   string
	Steps              []CaseStep
}

var placeholderPattern = regexp.MustCompile(`\{\{[a-z0-9_]+\}\}`)

// RunResourceCases loads every `case "<name>" { ... }` block from the merged
// tfvars file at testdata/<fileRelPath> and runs each as an independent
// subtest, replaying its steps verbatim. checksByBackend supplies the
// verification functions for each backend a case may target (e.g. "nios",
// "uddi").
func RunResourceCases(t *testing.T, resourceType, fileRelPath string, checksByBackend map[string]CheckFuncs) {
	path := GetTestdataPath(fileRelPath)
	cases, err := loadResourceCases(path)
	if err != nil {
		cases = nil
	}

	// Append user-authored custom cases from the sibling custom_<file> so custom
	// scenarios written by users run automatically alongside the generated ones.
	customCases, mErr := loadCustomResourceCases(fileRelPath)
	if mErr != nil {
		t.Fatalf("failed to load custom resource cases for %s: %v", fileRelPath, mErr)
	}
	cases = append(cases, customCases...)

	if len(cases) == 0 {
		t.Skipf("no resource cases at %s: %v", path, err)
		return
	}

	for _, rc := range cases {
		rc := rc
		t.Run(rc.Name, func(t *testing.T) {
			if rc.Skip {
				t.Skipf("Skipped: %s", rc.SkipReason)
				return
			}

			checks, ok := checksByBackend[rc.Backend]
			if !ok {
				t.Skipf("no check functions registered for backend %q", rc.Backend)
				return
			}

			PreCheck(t, rc.Backend)
			rc.materialize()
			runResourceCase(t, resourceType, rc, checks)
		})
	}
}

func runResourceCase(t *testing.T, resourceType string, rc *ResourceCase, checks CheckFuncs) {
	resourceAddr := resourceType + ".test"
	providerConfig := ProviderConfigHCL(rc.Backend)

	var steps []resource.TestStep
	for i, st := range rc.Steps {
		config := buildCaseHCL(resourceType, "test", rc.Backend, st)
		fullConfig := providerConfig + "\n" + rc.PrerequisitesHCL + "\n" + config

		t.Logf("Case %q step %d HCL:\n%s", rc.Name, i+1, fullConfig)

		var checkFuncs []resource.TestCheckFunc
		if checks.Exists != nil {
			checkFuncs = append(checkFuncs, checks.Exists(resourceAddr))
		}
		if rc.Disappears && checks.Disappears != nil {
			checkFuncs = append(checkFuncs, checks.Disappears(resourceAddr))
		}
		for attr, expected := range st.Checks {
			checkFuncs = append(checkFuncs, resource.TestCheckResourceAttr(resourceAddr, attr, expected))
		}

		step := resource.TestStep{
			Config: fullConfig,
			Check:  resource.ComposeTestCheckFunc(checkFuncs...),
		}
		if rc.ExpectNonEmptyPlan {
			step.ExpectNonEmptyPlan = true
		}
		steps = append(steps, step)
	}

	tc := resource.TestCase{
		ProtoV6ProviderFactories: ProtoV6ProviderFactories,
		Steps:                    steps,
	}
	if checks.Destroy != nil {
		tc.CheckDestroy = checks.Destroy(resourceType)
	}

	resource.ParallelTest(t, tc)
}

// buildCaseHCL renders a single step into resource HCL for the infoblox provider.
func buildCaseHCL(resourceType, label, backend string, st CaseStep) string {
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("resource %q %q {\n", resourceType, label))

	for k, v := range st.Common {
		sb.WriteString(fmt.Sprintf("  %s = %s\n", k, formatHCLValue(v)))
	}

	writeSection := func(name string, m map[string]any) {
		if len(m) == 0 {
			sb.WriteString(fmt.Sprintf("  %s = {}\n", name))
			return
		}
		sb.WriteString(fmt.Sprintf("  %s = {\n", name))
		for k, v := range m {
			sb.WriteString(fmt.Sprintf("    %s = %s\n", k, formatHCLValue(v)))
		}
		sb.WriteString("  }\n")
	}

	switch backend {
	case "nios":
		writeSection("nios", st.NIOS)
	case "uddi":
		writeSection("uddi", st.UDDI)
	}

	sb.WriteString("}\n")
	return sb.String()
}

// materialize replaces placeholders (e.g. {{random}}) with concrete values,
// keeping each distinct placeholder consistent across all steps and checks of
// the case so that, for example, a name used in config matches its check.
func (rc *ResourceCase) materialize() {
	cache := make(map[string]string)

	value := func(ph string) string {
		if v, ok := cache[ph]; ok {
			return v
		}
		v := ResolvePlaceholder(ph)
		cache[ph] = v
		return v
	}

	replace := func(v any) any {
		s, ok := v.(string)
		if !ok {
			return v
		}
		for _, ph := range placeholderPattern.FindAllString(s, -1) {
			s = strings.ReplaceAll(s, ph, value(ph))
		}
		return s
	}

	// replaceDeep recurses into nested objects/lists (e.g. members, options,
	// ext_attrs) so placeholders inside nested attribute values are materialized
	// too, keeping each distinct placeholder consistent across the whole case.
	var replaceDeep func(v any) any
	replaceDeep = func(v any) any {
		switch val := v.(type) {
		case map[string]any:
			for k, sub := range val {
				val[k] = replaceDeep(sub)
			}
			return val
		case []any:
			for i, sub := range val {
				val[i] = replaceDeep(sub)
			}
			return val
		default:
			return replace(v)
		}
	}

	replaceMap := func(m map[string]any) {
		for k, v := range m {
			m[k] = replaceDeep(v)
		}
	}

	if s, ok := replace(rc.PrerequisitesHCL).(string); ok {
		rc.PrerequisitesHCL = s
	}

	for i := range rc.Steps {
		replaceMap(rc.Steps[i].Common)
		replaceMap(rc.Steps[i].NIOS)
		replaceMap(rc.Steps[i].UDDI)
		for k, v := range rc.Steps[i].Checks {
			if s, ok := replace(v).(string); ok {
				rc.Steps[i].Checks[k] = s
			}
		}
	}
}

// customSibling returns the "custom_"-prefixed sibling of a testdata-relative
// case file (e.g. "dns/record_a/nios_resources.tfvars" ->
// "dns/record_a/custom_nios_resources.tfvars"). Custom files are hand-authored
// by users to add custom scenarios and are never overwritten by codegen.
func customSibling(fileRelPath string) string {
	i := strings.LastIndex(fileRelPath, "/")
	if i < 0 {
		return "custom_" + fileRelPath
	}
	return fileRelPath[:i+1] + "custom_" + fileRelPath[i+1:]
}

// loadCustomResourceCases loads resource cases from the sibling custom_ file if
// present. A missing file, or a file containing only comments (a skeleton with
// no `case` blocks), yields no cases and no error.
func loadCustomResourceCases(fileRelPath string) ([]*ResourceCase, error) {
	full := GetTestdataPath(customSibling(fileRelPath))
	if _, err := os.Stat(full); err != nil {
		return nil, nil
	}
	cases, err := loadResourceCases(full)
	if err != nil {
		if strings.Contains(err.Error(), "no case blocks found") {
			return nil, nil
		}
		return nil, err
	}
	return cases, nil
}

// loadResourceCases parses a merged tfvars file into its constituent resource
// cases. Each case is a labelled `case "<name>" { ... }` block; the label
// becomes the case (subtest) name. Cases are returned sorted by name for a
// deterministic subtest order.
func loadResourceCases(path string) ([]*ResourceCase, error) {
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

	var cases []*ResourceCase
	for _, block := range content.Blocks {
		rc, err := parseResourceCaseBody(block.Body, data)
		if err != nil {
			return nil, fmt.Errorf("case %q: %w", block.Labels[0], err)
		}
		rc.Name = block.Labels[0]
		cases = append(cases, rc)
	}

	if len(cases) == 0 {
		return nil, fmt.Errorf("no case blocks found in %s", path)
	}

	sort.Slice(cases, func(i, j int) bool { return cases[i].Name < cases[j].Name })
	return cases, nil
}

// parseResourceCaseBody decodes a single `case` block body into a ResourceCase
// (the Name is set by the caller from the block label).
func parseResourceCaseBody(body hcl.Body, src []byte) (*ResourceCase, error) {
	content, _, diags := body.PartialContent(&hcl.BodySchema{
		Attributes: []hcl.AttributeSchema{
			{Name: "backend"},
			{Name: "skip"},
			{Name: "skip_reason"},
			{Name: "disappears"},
			{Name: "expect_non_empty_plan"},
			{Name: "prerequisites_hcl"},
		},
		Blocks: []hcl.BlockHeaderSchema{
			{Type: "step"},
		},
	})
	if diags.HasErrors() {
		return nil, fmt.Errorf("failed to decode case: %s", diags.Error())
	}

	rc := &ResourceCase{}
	if attr, ok := content.Attributes["backend"]; ok {
		val, _ := attr.Expr.Value(nil)
		rc.Backend = val.AsString()
	}
	if attr, ok := content.Attributes["skip"]; ok {
		val, _ := attr.Expr.Value(nil)
		rc.Skip = val.True()
	}
	if attr, ok := content.Attributes["skip_reason"]; ok {
		val, _ := attr.Expr.Value(nil)
		rc.SkipReason = val.AsString()
	}
	if attr, ok := content.Attributes["disappears"]; ok {
		val, _ := attr.Expr.Value(nil)
		rc.Disappears = val.True()
	}
	if attr, ok := content.Attributes["expect_non_empty_plan"]; ok {
		val, _ := attr.Expr.Value(nil)
		rc.ExpectNonEmptyPlan = val.True()
	}
	if attr, ok := content.Attributes["prerequisites_hcl"]; ok {
		val, _ := attr.Expr.Value(nil)
		rc.PrerequisitesHCL = val.AsString()
	}

	for _, block := range content.Blocks {
		if block.Type != "step" {
			continue
		}
		st, err := parseCaseStep(block.Body, src)
		if err != nil {
			return nil, err
		}
		rc.Steps = append(rc.Steps, st)
	}

	return rc, nil
}

func parseCaseStep(body hcl.Body, src []byte) (CaseStep, error) {
	st := CaseStep{
		Common: make(map[string]any),
		NIOS:   make(map[string]any),
		UDDI:   make(map[string]any),
		Checks: make(map[string]string),
	}

	content, _, diags := body.PartialContent(&hcl.BodySchema{
		Attributes: []hcl.AttributeSchema{
			{Name: "check"},
		},
		Blocks: []hcl.BlockHeaderSchema{
			{Type: "common"},
			{Type: "nios"},
			{Type: "uddi"},
		},
	})
	if diags.HasErrors() {
		return st, fmt.Errorf("failed to parse step: %s", diags.Error())
	}

	if attr, ok := content.Attributes["check"]; ok {
		val, _ := attr.Expr.Value(nil)
		st.Checks = ctyMapToStringMap(val)
	}

	for _, block := range content.Blocks {
		switch block.Type {
		case "common":
			st.Common = parseCaseBlock(block.Body, src)
		case "nios":
			st.NIOS = parseCaseBlock(block.Body, src)
		case "uddi":
			st.UDDI = parseCaseBlock(block.Body, src)
		}
	}

	return st, nil
}

// parseCaseBlock parses a config block (common/nios/uddi) into a value map.
// Literal values are reduced to Go values; non-literal expressions (references
// to other resources, e.g. ${infoblox_zone_auth.prereq.nios.fqdn}, or
// depends_on) are preserved verbatim as RawExpr so Terraform resolves them at
// apply time.
func parseCaseBlock(body hcl.Body, src []byte) map[string]any {
	result := make(map[string]any)
	attrs, _ := body.JustAttributes()
	for name, attr := range attrs {
		val, diags := attr.Expr.Value(nil)
		if diags.HasErrors() || !val.IsWhollyKnown() {
			rng := attr.Expr.Range()
			if src != nil && rng.End.Byte <= len(src) && rng.Start.Byte >= 0 {
				result[name] = RawExpr(strings.TrimSpace(string(src[rng.Start.Byte:rng.End.Byte])))
			}
			continue
		}
		result[name] = ctyToGo(val)
	}
	return result
}
