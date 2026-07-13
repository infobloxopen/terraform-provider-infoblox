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
)

// DataSourceCase is a per-test-case data source acceptance configuration,
// generated from the legacy provider's hand-written data source acceptance
// tests. Each case maps 1:1 to a legacy TestAcc<Object>DataSource_<Case>
// function and is stored as a labelled `case "<name>" { ... }` block in the
// merged testdata/<package>/<object>/<backend>_datasources.tfvars file.
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
}

// RunDataSourceCases loads every `case "<name>" { ... }` block from the merged
// tfvars file at testdata/<fileRelPath> and runs each as an independent
// subtest. Each case provisions the resource described by its step, queries the
// data source using the case's filter, and asserts the data source returns the
// same values as the created resource.
func RunDataSourceCases(t *testing.T, dsType, resourceType, fileRelPath string, checksByBackend map[string]CheckFuncs) {
	path := GetTestdataPath(fileRelPath)
	cases, err := loadDataSourceCases(path)
	if err != nil {
		cases = nil
	}

	// Append user-authored custom cases from the sibling custom_<file> so custom
	// scenarios written by users run automatically alongside the generated ones.
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

// buildDataSourceBlock renders the data source HCL with the case's filter,
// whose values reference attributes of the resource created in the same config.
func buildDataSourceBlock(dsType, resourceType string, dc *DataSourceCase) string {
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("data %q \"test\" {\n", dsType))
	sb.WriteString(fmt.Sprintf("  %s = {\n", dc.FilterType))
	for _, key := range dc.FilterOrder {
		sb.WriteString(fmt.Sprintf("    %q = %s.test.%s\n", key, resourceType, dc.Filters[key]))
	}
	sb.WriteString("  }\n")
	sb.WriteString(fmt.Sprintf("  depends_on = [%s.test]\n", resourceType))
	sb.WriteString("}\n")
	return sb.String()
}

// dataSourcePairChecks asserts that each scalar field configured on the
// resource is returned identically by the data source at results.0.<path>.
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

	return checks
}

// isScalarValue reports whether v is a simple scalar (and thus safe to compare
// with TestCheckResourceAttrPair). Maps, lists and raw expressions are skipped.
func isScalarValue(v any) bool {
	switch v.(type) {
	case map[string]any, []any, RawExpr:
		return false
	default:
		return true
	}
}

// materialize replaces placeholders (e.g. {{random}}) with concrete values,
// keeping each distinct placeholder consistent across the resource config and
// prerequisites so the data source filter resolves to the created resource.
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

// loadDataSourceCases parses a merged tfvars file into its constituent data
// source cases. Each case is a labelled `case "<name>" { ... }` block; the
// label becomes the case (subtest) name. Cases are returned sorted by name for
// a deterministic subtest order.
// loadCustomDataSourceCases loads data source cases from the sibling custom_
// file if present. A missing file, or a file containing only comments (a
// skeleton with no `case` blocks), yields no cases and no error.
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

// parseDataSourceCaseBody decodes a single `case` block body into a
// DataSourceCase (the Name is set by the caller from the block label).
func parseDataSourceCaseBody(body hcl.Body, src []byte) (*DataSourceCase, error) {
	content, _, diags := body.PartialContent(&hcl.BodySchema{
		Attributes: []hcl.AttributeSchema{
			{Name: "backend"},
			{Name: "skip"},
			{Name: "skip_reason"},
			{Name: "prerequisites_hcl"},
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
