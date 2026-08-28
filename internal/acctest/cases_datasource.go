package acctest

import (
	"fmt"
	"os"
	"sort"
	"strings"
	"testing"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclparse"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/zclconf/go-cty/cty"
)

// DataSourceCase is the per-subtest configuration for a data source acceptance test.
// Each case maps to a `case "<name>" { ... }` block in <backend>_datasources.hcl.
type DataSourceCase struct {
	Name             string
	Backend          string
	Skip             bool
	SkipReason       string
	PrerequisitesHCL string
	FilterType       string            // filters | ext_attr_filters | tag_filters
	Filters          map[string]string // filter key -> resource attribute path (e.g. "name" -> "nios.name")
	FilterOrder      []string          // filter keys in deterministic order
	Step             CaseStep
	// PairChecks lists extra backend-prefixed paths (e.g. "nios.comment") for TestCheckResourceAttrPair.
	PairChecks []string
}

// RunDataSourceCases loads all `case` blocks from testdata/<fileRelPath> and runs each as a subtest,
// provisioning the resource and asserting the data source returns matching attribute values.
func RunDataSourceCases(t *testing.T, dsType, resourceType, fileRelPath string, checksByBackend map[string]CheckFuncs) {
	path := GetTestdataPath(fileRelPath)
	cases, err := loadDataSourceCases(path)
	if err != nil {
		cases = nil
	}

	// Also run custom cases from the sibling custom_<file>.
	customCases, mErr := loadCustomDataSourceCases(fileRelPath)
	if mErr != nil {
		t.Fatalf("failed to load custom data source cases for %s: %v", fileRelPath, mErr)
	}
	cases = append(cases, customCases...)

	if len(cases) == 0 {
		t.Skipf("no data source cases at %s: %v", path, err)
		return
	}

	for _, dc := range cases {
		dc := dc
		t.Run(dc.Name, func(t *testing.T) {
			if dc.Skip {
				t.Skipf("Skipped: %s", dc.SkipReason)
				return
			}

			checks, ok := checksByBackend[dc.Backend]
			if !ok {
				t.Skipf("no check functions registered for backend %q", dc.Backend)
				return
			}

			PreCheck(t, dc.Backend)
			dc.materialize()
			runDataSourceCase(t, dsType, resourceType, dc, checks)
		})
	}
}

func runDataSourceCase(t *testing.T, dsType, resourceType string, dc *DataSourceCase, checks CheckFuncs) {
	resourceAddr := resourceType + ".test"
	dsAddr := "data." + dsType + ".test"

	providerConfig := ProviderConfigHCL(dc.Backend)
	resourceHCL := buildCaseHCL(resourceType, "test", dc.Backend, dc.Step)
	dsHCL := buildDataSourceBlock(dsType, resourceType, dc)
	fullConfig := providerConfig + "\n" + dc.PrerequisitesHCL + "\n" + resourceHCL + "\n" + dsHCL

	t.Logf("Case %q HCL:\n%s", dc.Name, fullConfig)

	var checkFuncs []resource.TestCheckFunc
	if checks.Exists != nil {
		checkFuncs = append(checkFuncs, checks.Exists(resourceAddr))
	}
	checkFuncs = append(checkFuncs, resource.TestCheckResourceAttrSet(dsAddr, "results.0.id"))
	checkFuncs = append(checkFuncs, dataSourcePairChecks(dsAddr, resourceAddr, dc)...)

	tc := resource.TestCase{
		ProtoV6ProviderFactories: ProtoV6ProviderFactories,
		Steps: []resource.TestStep{{
			Config: fullConfig,
			Check:  resource.ComposeTestCheckFunc(checkFuncs...),
		}},
	}
	if checks.Destroy != nil {
		tc.CheckDestroy = checks.Destroy(resourceType)
	}

	resource.Test(t, tc)
}

// buildDataSourceBlock renders the data source HCL block with filter values referencing the test resource.
func buildDataSourceBlock(dsType, resourceType string, dc *DataSourceCase) string {
	var sb strings.Builder
	fmt.Fprintf(&sb, "data %q \"test\" {\n", dsType)
	fmt.Fprintf(&sb, "  %s = {\n", dc.FilterType)
	for _, key := range dc.FilterOrder {
		fmt.Fprintf(&sb, "    %q = %s.test.%s\n", key, resourceType, dc.Filters[key])
	}
	sb.WriteString("  }\n")
	fmt.Fprintf(&sb, "  depends_on = [%s.test]\n", resourceType)
	sb.WriteString("}\n")
	return sb.String()
}

// dataSourcePairChecks returns TestCheckResourceAttrPair checks for all scalar fields in the step config.
func dataSourcePairChecks(dsAddr, resourceAddr string, dc *DataSourceCase) []resource.TestCheckFunc {
	var checks []resource.TestCheckFunc

	add := func(prefix string, m map[string]any) {
		for k, v := range m {
			if !isScalarValue(v) {
				continue
			}
			path := k
			if prefix != "" {
				path = prefix + "." + k
			}
			checks = append(checks, resource.TestCheckResourceAttrPair(
				dsAddr, "results.0."+path,
				resourceAddr, path,
			))
		}
	}

	add("", dc.Step.Common)
	switch dc.Backend {
	case "nios":
		add("nios", dc.Step.NIOS)
	case "uddi":
		add("uddi", dc.Step.UDDI)
	}

	// Track paths already covered above to avoid duplicates from PairChecks.
	covered := make(map[string]bool, len(checks))
	for _, fn := range checks {
		_ = fn // not inspectable; re-derive covered paths below
	}
	addCovered := func(prefix string, m map[string]any) {
		for k, v := range m {
			if !isScalarValue(v) {
				continue
			}
			path := k
			if prefix != "" {
				path = prefix + "." + k
			}
			covered[path] = true
		}
	}
	addCovered("", dc.Step.Common)
	switch dc.Backend {
	case "nios":
		addCovered("nios", dc.Step.NIOS)
	case "uddi":
		addCovered("uddi", dc.Step.UDDI)
	}

	for _, path := range dc.PairChecks {
		if covered[path] {
			continue
		}
		checks = append(checks, resource.TestCheckResourceAttrPair(
			dsAddr, "results.0."+path,
			resourceAddr, path,
		))
		covered[path] = true
	}

	return checks
}

// isScalarValue reports whether v is a scalar safe to use with TestCheckResourceAttrPair.
func isScalarValue(v any) bool {
	switch v.(type) {
	case map[string]any, []any, RawExpr:
		return false
	default:
		return true
	}
}

// materialize replaces {{placeholder}} tokens with stable values across the step config and prerequisites.
func (dc *DataSourceCase) materialize() {
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

	if s, ok := replace(dc.PrerequisitesHCL).(string); ok {
		dc.PrerequisitesHCL = s
	}
	for k, v := range dc.Step.Common {
		dc.Step.Common[k] = replace(v)
	}
	for k, v := range dc.Step.NIOS {
		dc.Step.NIOS[k] = replace(v)
	}
	for k, v := range dc.Step.UDDI {
		dc.Step.UDDI[k] = replace(v)
	}
}

// loadCustomDataSourceCases loads cases from the custom_ sibling file; returns nil if absent or empty.
func loadCustomDataSourceCases(fileRelPath string) ([]*DataSourceCase, error) {
	full := GetTestdataPath(customSibling(fileRelPath))
	if _, err := os.Stat(full); err != nil {
		return nil, nil
	}
	cases, err := loadDataSourceCases(full)
	if err != nil {
		if strings.Contains(err.Error(), "no case blocks found") {
			return nil, nil
		}
		return nil, err
	}
	return cases, nil
}

// loadDataSourceCases parses `case "<name>" { ... }` blocks from a data source case file,
// sorted by name for deterministic subtest ordering.
func loadDataSourceCases(path string) ([]*DataSourceCase, error) {
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

	var cases []*DataSourceCase
	for _, block := range content.Blocks {
		dc, err := parseDataSourceCaseBody(block.Body, data)
		if err != nil {
			return nil, fmt.Errorf("case %q: %w", block.Labels[0], err)
		}
		dc.Name = block.Labels[0]
		cases = append(cases, dc)
	}

	if len(cases) == 0 {
		return nil, fmt.Errorf("no case blocks found in %s", path)
	}

	sort.Slice(cases, func(i, j int) bool { return cases[i].Name < cases[j].Name })
	return cases, nil
}

// parseDataSourceCaseBody decodes a `case` block body into a DataSourceCase (Name set by caller).
func parseDataSourceCaseBody(body hcl.Body, src []byte) (*DataSourceCase, error) {
	content, _, diags := body.PartialContent(&hcl.BodySchema{
		Attributes: []hcl.AttributeSchema{
			{Name: "backend"},
			{Name: "skip"},
			{Name: "skip_reason"},
			{Name: "prerequisites_hcl"},
			{Name: "pair_checks"},
		},
		Blocks: []hcl.BlockHeaderSchema{
			{Type: "filter"},
			{Type: "step"},
		},
	})
	if diags.HasErrors() {
		return nil, fmt.Errorf("failed to decode case: %s", diags.Error())
	}

	dc := &DataSourceCase{Filters: make(map[string]string)}
	if attr, ok := content.Attributes["backend"]; ok {
		val, _ := attr.Expr.Value(nil)
		dc.Backend = val.AsString()
	}
	if attr, ok := content.Attributes["skip"]; ok {
		val, _ := attr.Expr.Value(nil)
		dc.Skip = val.True()
	}
	if attr, ok := content.Attributes["skip_reason"]; ok {
		val, _ := attr.Expr.Value(nil)
		dc.SkipReason = val.AsString()
	}
	if attr, ok := content.Attributes["prerequisites_hcl"]; ok {
		val, _ := attr.Expr.Value(nil)
		dc.PrerequisitesHCL = val.AsString()
	}
	if attr, ok := content.Attributes["pair_checks"]; ok {
		val, _ := attr.Expr.Value(nil)
		if val.CanIterateElements() {
			it := val.ElementIterator()
			for it.Next() {
				_, v := it.Element()
				if v.Type() == cty.String {
					dc.PairChecks = append(dc.PairChecks, v.AsString())
				}
			}
		}
	}

	for _, block := range content.Blocks {
		switch block.Type {
		case "filter":
			parseFilterBlock(block.Body, dc)
		case "step":
			st, err := parseCaseStep(block.Body, src)
			if err != nil {
				return nil, err
			}
			dc.Step = st
		}
	}

	return dc, nil
}

func parseFilterBlock(body hcl.Body, dc *DataSourceCase) {
	content, _, _ := body.PartialContent(&hcl.BodySchema{
		Attributes: []hcl.AttributeSchema{
			{Name: "type"},
			{Name: "values"},
		},
	})

	if attr, ok := content.Attributes["type"]; ok {
		val, _ := attr.Expr.Value(nil)
		dc.FilterType = val.AsString()
	}
	if attr, ok := content.Attributes["values"]; ok {
		val, _ := attr.Expr.Value(nil)
		dc.Filters = ctyMapToStringMap(val)
		for k := range dc.Filters {
			dc.FilterOrder = append(dc.FilterOrder, k)
		}
		sort.Strings(dc.FilterOrder)
	}
}
