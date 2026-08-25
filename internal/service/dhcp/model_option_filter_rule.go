package dhcp

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
	uddidhcp "github.com/infobloxopen/universal-ddi-go-client/ipam"
)

// OptionFilterRuleModel is the Terraform model for OptionFilterRule
type OptionFilterRuleModel struct {
	Compare         types.String `tfsdk:"compare"`
	OptionCode      types.String `tfsdk:"option_code"`
	OptionValue     types.String `tfsdk:"option_value"`
	SubstringOffset types.Int64  `tfsdk:"substring_offset"`
}

// OptionFilterRuleAttrTypes contains the attribute types for OptionFilterRuleModel
var OptionFilterRuleAttrTypes = map[string]attr.Type{
	"compare":          types.StringType,
	"option_code":      types.StringType,
	"option_value":     types.StringType,
	"substring_offset": types.Int64Type,
}

// OptionFilterRuleResourceSchemaAttributes contains the schema attributes for OptionFilterRuleModel
var OptionFilterRuleResourceSchemaAttributes = map[string]schema.Attribute{
	"compare": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "Indicates how to compare the _option_value_ to the client option.  Success by comparison:  * _equals_: value and client option are the same,  * _not_equals_: value and client option are not the same,  * _exists_: client option exists,  * _not_exists_: client option does not exist,  * _text_substring_: value is the specified substring of the option,  * _not_text_substring_: value is not the specified substring of the option.  * _hex_substring_: value is the specified hexadecimal substring of the option,  * _not_hex_substring_: value is not the specified hexadecimal substring of the option.",
	},
	"option_code": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The resource identifier.",
	},
	"option_value": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The value to match against.",
	},
	"substring_offset": schema.Int64Attribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The offset where the substring match starts. This is used only if comparing the _option_value_ using any of the substring modes.",
	},
}

// ExpandOptionFilterRule converts a Terraform Object to SDK type
func ExpandOptionFilterRule(ctx context.Context, o types.Object, diags *diag.Diagnostics) *uddidhcp.OptionFilterRule {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m OptionFilterRuleModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

// Expand converts the Terraform model to SDK type
func (m *OptionFilterRuleModel) Expand(ctx context.Context, diags *diag.Diagnostics) *uddidhcp.OptionFilterRule {
	if m == nil {
		return nil
	}
	to := &uddidhcp.OptionFilterRule{
		Compare:         flex.ExpandString(m.Compare),
		OptionCode:      flex.ExpandString(m.OptionCode),
		OptionValue:     flex.ExpandStringPointer(m.OptionValue),
		SubstringOffset: flex.ExpandInt64Pointer(m.SubstringOffset),
	}
	return to
}

// FlattenOptionFilterRule converts an SDK type to Terraform Object
func FlattenOptionFilterRule(ctx context.Context, from *uddidhcp.OptionFilterRule, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(OptionFilterRuleAttrTypes)
	}
	m := &OptionFilterRuleModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, OptionFilterRuleAttrTypes, m)
	diags.Append(d...)
	return t
}

// Flatten populates the Terraform model from SDK type
func (m *OptionFilterRuleModel) Flatten(ctx context.Context, from *uddidhcp.OptionFilterRule, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.Compare = flex.FlattenString(from.Compare)
	m.OptionCode = flex.FlattenString(from.OptionCode)
	m.OptionValue = flex.FlattenStringPointer(from.OptionValue)
	m.SubstringOffset = flex.FlattenInt64Pointer(from.SubstringOffset)
}
