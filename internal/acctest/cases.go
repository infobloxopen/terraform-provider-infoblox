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
	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/core"
)

// ExtractNIOSRef strips the object-type prefix from a NIOS _ref (e.g. "record:a/ZG5z..." -> "ZG5z..."),
// since the raw NIOS SDK re-applies the prefix and would otherwise double it.
func ExtractNIOSRef(ref string) string {
	return core.ExtractNIOSRef(ref)
}

// CaseStep is a single Terraform step within a generated resource case.
type CaseStep struct {
	Common    map[string]any
	NIOS      map[string]any
	UDDI      map[string]any
	Checks    map[string]string
	DependsOn []string
	// PrerequisitesHCL overrides ResourceCase.PrerequisitesHCL for steps that need different prereqs.
	PrerequisitesHCL string
	PairChecks       map[string]string
}

// ResourceCase is the per-subtest configuration for a Terraform resource acceptance test.
// Each case maps to a `case "<name>" { ... }` block in <backend>_resources.hcl.
type ResourceCase struct {
	Name               string
	Backend            string
	Skip               bool
	SkipReason         string
	Disappears         bool
	ExpectNonEmptyPlan bool
	Parallel           bool
	PrerequisitesHCL   string
	Steps              []CaseStep
}

var placeholderPattern = regexp.MustCompile(`\{\{[a-z0-9_]+\}\}`)

// RunResourceCases loads all `case` blocks from testdata/<fileRelPath> and runs each as an independent
// subtest, replaying its steps. checksByBackend provides verification functions per backend.
func RunResourceCases(t *testing.T, resourceType, fileRelPath string, checksByBackend map[string]CheckFuncs) {
	path := GetTestdataPath(fileRelPath)
	cases, err := loadResourceCases(path)
	if err != nil {
		cases = nil
	}

	// Also run custom cases from the sibling custom_<file>.
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
		prereq := rc.PrerequisitesHCL
		if st.PrerequisitesHCL != "" {
			prereq = st.PrerequisitesHCL
		}
		fullConfig := providerConfig + "\n" + prereq + "\n" + config

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
		for attr, ref := range st.PairChecks {
			// "infoblox_x.test1.uddi.name" -> address "infoblox_x.test1", attribute "uddi.name".
			parts := strings.SplitN(ref, ".", 3)
			if len(parts) < 3 {
				t.Fatalf("case %q step %d: check_pair %q: %q must be <type>.<name>.<attr>", rc.Name, i+1, attr, ref)
			}
			checkFuncs = append(checkFuncs, resource.TestCheckResourceAttrPair(
				resourceAddr, attr, parts[0]+"."+parts[1], parts[2]))
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

	if rc.Parallel {
		resource.ParallelTest(t, tc)
	} else {
		resource.Test(t, tc)
	}
}

// buildCaseHCL renders a single step into resource HCL for the infoblox provider.
func buildCaseHCL(resourceType, label, backend string, st CaseStep) string {
	var sb strings.Builder
	fmt.Fprintf(&sb, "resource %q %q {\n", resourceType, label)

	for k, v := range st.Common {
		fmt.Fprintf(&sb, "  %s = %s\n", k, formatHCLValue(v))
	}

	writeSection := func(name string, m map[string]any) {
		if len(m) == 0 {
			fmt.Fprintf(&sb, "  %s = {}\n", name)
			return
		}
		fmt.Fprintf(&sb, "  %s = {\n", name)
		for k, v := range m {
			fmt.Fprintf(&sb, "    %s = %s\n", k, formatHCLValue(v))
		}
		sb.WriteString("  }\n")
	}

	switch backend {
	case "nios":
		writeSection("nios", st.NIOS)
	case "uddi":
		writeSection("uddi", st.UDDI)
	}

	if len(st.DependsOn) > 0 {
		fmt.Fprintf(&sb, "  depends_on = [%s]\n", strings.Join(st.DependsOn, ", "))
	}

	sb.WriteString("}\n")
	return sb.String()
}

// materialize replaces {{placeholder}} tokens with stable concrete values across all steps and checks.
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

	substitute := func(s string) string {
		for _, ph := range placeholderPattern.FindAllString(s, -1) {
			s = strings.ReplaceAll(s, ph, value(ph))
		}
		return s
	}

	replace := func(v any) any {
		switch t := v.(type) {
		case RawExpr:
			// RawExpr values may still contain placeholders; substitute and keep raw.
			return RawExpr(substitute(string(t)))
		case string:
			return substitute(t)
		default:
			return v
		}
	}

	// replaceDeep recurses into nested objects/lists so placeholders inside nested attributes are also materialized.
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
		if s, ok := replace(rc.Steps[i].PrerequisitesHCL).(string); ok {
			rc.Steps[i].PrerequisitesHCL = s
		}
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

// customSibling returns the "custom_"-prefixed sibling path for user-authored cases (never overwritten).
// e.g. "dns/record_a/nios_resources.hcl" -> "dns/record_a/custom_nios_resources.hcl".
func customSibling(fileRelPath string) string {
	i := strings.LastIndex(fileRelPath, "/")
	if i < 0 {
		return "custom_" + fileRelPath
	}
	return fileRelPath[:i+1] + "custom_" + fileRelPath[i+1:]
}

// loadCustomResourceCases loads cases from the custom_ sibling file; returns nil if absent or empty.
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

// loadResourceCases parses `case "<name>" { ... }` blocks from a case file,
// sorted by name for deterministic subtest ordering.
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

// parseResourceCaseBody decodes a `case` block body into a ResourceCase (Name set by caller).
func parseResourceCaseBody(body hcl.Body, src []byte) (*ResourceCase, error) {
	content, _, diags := body.PartialContent(&hcl.BodySchema{
		Attributes: []hcl.AttributeSchema{
			{Name: "backend"},
			{Name: "skip"},
			{Name: "skip_reason"},
			{Name: "disappears"},
			{Name: "expect_non_empty_plan"},
			{Name: "parallel"},
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
	if attr, ok := content.Attributes["parallel"]; ok {
		val, _ := attr.Expr.Value(nil)
		rc.Parallel = val.True()
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
			{Name: "check_pair"},
			{Name: "depends_on"},
			{Name: "prerequisites_hcl"},
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

	// check_pair values are kept as raw source text, since a reference to another resource has no
	// static value at load time and would otherwise be dropped by ctyMapToStringMap.
	if attr, ok := content.Attributes["check_pair"]; ok {
		obj, ok := attr.Expr.(*hclsyntax.ObjectConsExpr)
		if !ok {
			return st, fmt.Errorf("check_pair must be an object literal")
		}
		st.PairChecks = make(map[string]string, len(obj.Items))
		for _, item := range obj.Items {
			key, _ := item.KeyExpr.Value(nil)
			rng := item.ValueExpr.Range()
			st.PairChecks[key.AsString()] = strings.TrimSpace(string(src[rng.Start.Byte:rng.End.Byte]))
		}
	}

	if attr, ok := content.Attributes["prerequisites_hcl"]; ok {
		val, _ := attr.Expr.Value(nil)
		st.PrerequisitesHCL = val.AsString()
	}

	if attr, ok := content.Attributes["depends_on"]; ok {
		rng := attr.Expr.Range()
		if src != nil && rng.Start.Byte >= 0 && rng.End.Byte <= len(src) {
			raw := strings.TrimSpace(string(src[rng.Start.Byte:rng.End.Byte]))
			raw = strings.TrimPrefix(strings.TrimSuffix(raw, "]"), "[")
			for _, ref := range strings.Split(raw, ",") {
				if ref = strings.TrimSpace(ref); ref != "" {
					st.DependsOn = append(st.DependsOn, ref)
				}
			}
		}
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

// parseCaseBlock parses a common/nios/uddi block into a value map; non-literal expressions are kept as RawExpr.
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
