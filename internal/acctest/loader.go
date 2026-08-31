package acctest

import (
	"fmt"
	"os"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclparse"
	"github.com/zclconf/go-cty/cty"
)

// LoadCaseFile reads and parses the HCL case file at relativePath, substituting all {{placeholder}} tokens.
func LoadCaseFile(relativePath string) (*CaseConfig, error) {
	path := GetTestdataPath(relativePath)
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read case file %s: %w", path, err)
	}

	content := ReplacePlaceholders(string(data))

	parser := hclparse.NewParser()
	file, diags := parser.ParseHCL([]byte(content), path)
	if diags.HasErrors() {
		return nil, fmt.Errorf("failed to parse HCL: %s", diags.Error())
	}

	tv := &CaseConfig{
		Common: make(map[string]any),
		NIOS:   make(map[string]any),
		UDDI:   make(map[string]any),
	}

	body := file.Body
	contentBody, _, diags := body.PartialContent(&hcl.BodySchema{
		Attributes: []hcl.AttributeSchema{
			{Name: "backend"},
			{Name: "parallel"},
			{Name: "prerequisites_hcl"},
			{Name: "ds_filter_field"},
		},
		Blocks: []hcl.BlockHeaderSchema{
			{Type: "common"},
			{Type: "nios"},
			{Type: "uddi"},
		},
	})
	if diags.HasErrors() {
		return nil, fmt.Errorf("failed to decode case file: %s", diags.Error())
	}

	if attr, exists := contentBody.Attributes["backend"]; exists {
		val, _ := attr.Expr.Value(nil)
		tv.Backend = val.AsString()
	}
	if attr, exists := contentBody.Attributes["parallel"]; exists {
		val, _ := attr.Expr.Value(nil)
		tv.Parallel = val.True()
	}
	if attr, exists := contentBody.Attributes["prerequisites_hcl"]; exists {
		val, _ := attr.Expr.Value(nil)
		tv.PrerequisitesHCL = val.AsString()
	}
	if attr, exists := contentBody.Attributes["ds_filter_field"]; exists {
		val, _ := attr.Expr.Value(nil)
		tv.DSFilterField = val.AsString()
	}

	for _, block := range contentBody.Blocks {
		switch block.Type {
		case "common":
			tv.Common = parseBlockToMap(block.Body)
		case "nios":
			tv.NIOS = parseBlockToMap(block.Body)
		case "uddi":
			tv.UDDI = parseBlockToMap(block.Body)
		}
	}

	return tv, nil
}

func parseBlockToMap(body hcl.Body) map[string]any {
	result := make(map[string]any)
	attrs, _ := body.JustAttributes()
	for name, attr := range attrs {
		val, _ := attr.Expr.Value(nil)
		result[name] = ctyToGo(val)
	}
	return result
}

func ctyToGo(val cty.Value) any {
	if val.IsNull() {
		return nil
	}
	if !val.IsKnown() {
		// Unknown/reference values cannot be reduced to a literal; callers handle source text upstream.
		return nil
	}
	switch val.Type() {
	case cty.String:
		return val.AsString()
	case cty.Number:
		f, _ := val.AsBigFloat().Float64()
		if f == float64(int(f)) {
			return int(f)
		}
		return f
	case cty.Bool:
		return val.True()
	default:
		if val.Type().IsMapType() || val.Type().IsObjectType() {
			result := make(map[string]any)
			for it := val.ElementIterator(); it.Next(); {
				k, v := it.Element()
				result[k.AsString()] = ctyToGo(v)
			}
			return result
		}
		if val.Type().IsListType() || val.Type().IsTupleType() {
			var result []any
			for it := val.ElementIterator(); it.Next(); {
				_, v := it.Element()
				result = append(result, ctyToGo(v))
			}
			return result
		}
		return val.GoString()
	}
}

func ctyMapToStringMap(val cty.Value) map[string]string {
	result := make(map[string]string)
	if val.IsNull() || !val.CanIterateElements() {
		return result
	}
	for it := val.ElementIterator(); it.Next(); {
		k, v := it.Element()
		if !k.IsKnown() || !v.IsWhollyKnown() {
			// Skip entries referencing another resource; they have no static expected value.
			continue
		}
		result[k.AsString()] = fmt.Sprintf("%v", ctyToGo(v))
	}
	return result
}
