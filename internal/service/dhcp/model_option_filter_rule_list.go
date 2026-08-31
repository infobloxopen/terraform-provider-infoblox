package dhcp

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
	customvalidator "github.com/infobloxopen/terraform-provider-infoblox/internal/validator"
	uddidhcp "github.com/infobloxopen/universal-ddi-go-client/ipam"
)

// OptionFilterRuleListModel is the Terraform model for OptionFilterRuleList
type OptionFilterRuleListModel struct {
	Match types.String `tfsdk:"match"`
	Rules types.List   `tfsdk:"rules"`
}

// OptionFilterRuleListAttrTypes contains the attribute types for OptionFilterRuleListModel
var OptionFilterRuleListAttrTypes = map[string]attr.Type{
	"match": types.StringType,
	"rules": types.ListType{ElemType: types.ObjectType{AttrTypes: OptionFilterRuleAttrTypes}},
}

// OptionFilterRuleListResourceSchemaAttributes contains the schema attributes for OptionFilterRuleListModel
var OptionFilterRuleListResourceSchemaAttributes = map[string]schema.Attribute{
	"match": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "Indicates if this list should match if any or all rules match (_any_ or _all_).",
	},
	"rules": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: OptionFilterRuleResourceSchemaAttributes,
		},
		Optional: true,
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
		},
		MarkdownDescription: "The list of child rules.",
	},
}

// ExpandOptionFilterRuleList converts a Terraform Object to SDK type
func ExpandOptionFilterRuleList(ctx context.Context, o types.Object, diags *diag.Diagnostics) *uddidhcp.OptionFilterRuleList {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m OptionFilterRuleListModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

// Expand converts the Terraform model to SDK type
func (m *OptionFilterRuleListModel) Expand(ctx context.Context, diags *diag.Diagnostics) *uddidhcp.OptionFilterRuleList {
	if m == nil {
		return nil
	}
	to := &uddidhcp.OptionFilterRuleList{
		Match: flex.ExpandStringPointer(m.Match),
		Rules: flex.ExpandFrameworkListNestedBlock(ctx, m.Rules, diags, ExpandOptionFilterRule),
	}
	return to
}

// FlattenOptionFilterRuleList converts an SDK type to Terraform Object
func FlattenOptionFilterRuleList(ctx context.Context, from *uddidhcp.OptionFilterRuleList, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(OptionFilterRuleListAttrTypes)
	}
	m := &OptionFilterRuleListModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, OptionFilterRuleListAttrTypes, m)
	diags.Append(d...)
	return t
}

// Flatten populates the Terraform model from SDK type
func (m *OptionFilterRuleListModel) Flatten(ctx context.Context, from *uddidhcp.OptionFilterRuleList, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.Match = flex.FlattenStringPointer(from.Match)
	m.Rules = flex.FlattenFrameworkListNestedBlock(ctx, from.Rules, OptionFilterRuleAttrTypes, diags, FlattenOptionFilterRule)
}
