package dhcp

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	niosdhcp "github.com/infobloxopen/infoblox-nios-go-client/dhcp"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
	customvalidator "github.com/infobloxopen/terraform-provider-infoblox/internal/validator"
)

// FixedaddressLogicFilterRulesModel is the Terraform model for FixedaddressLogicFilterRules
type FixedaddressLogicFilterRulesModel struct {
	Filter types.String `tfsdk:"filter"`
	Type   types.String `tfsdk:"type"`
}

// FixedaddressLogicFilterRulesAttrTypes contains the attribute types for FixedaddressLogicFilterRulesModel
var FixedaddressLogicFilterRulesAttrTypes = map[string]attr.Type{
	"filter": types.StringType,
	"type":   types.StringType,
}

// FixedaddressLogicFilterRulesResourceSchemaAttributes contains the schema attributes for FixedaddressLogicFilterRulesModel
var FixedaddressLogicFilterRulesResourceSchemaAttributes = map[string]schema.Attribute{
	"filter": schema.StringAttribute{
		Required: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "The filter name.",
	},
	"type": schema.StringAttribute{
		Required: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "The filter type. Valid values are: * MAC * NAC * Option",
	},
}

// ExpandFixedaddressLogicFilterRules converts a Terraform Object to SDK type
func ExpandFixedaddressLogicFilterRules(ctx context.Context, o types.Object, diags *diag.Diagnostics) *niosdhcp.FixedaddressLogicFilterRules {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m FixedaddressLogicFilterRulesModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

// Expand converts the Terraform model to SDK type
func (m *FixedaddressLogicFilterRulesModel) Expand(ctx context.Context, diags *diag.Diagnostics) *niosdhcp.FixedaddressLogicFilterRules {
	if m == nil {
		return nil
	}
	to := &niosdhcp.FixedaddressLogicFilterRules{
		Filter: flex.ExpandStringPointerNullAsEmpty(m.Filter),
		Type:   flex.ExpandStringPointerNullAsEmpty(m.Type),
	}
	return to
}

// FlattenFixedaddressLogicFilterRules converts an SDK type to Terraform Object
func FlattenFixedaddressLogicFilterRules(ctx context.Context, from *niosdhcp.FixedaddressLogicFilterRules, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(FixedaddressLogicFilterRulesAttrTypes)
	}
	m := &FixedaddressLogicFilterRulesModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, FixedaddressLogicFilterRulesAttrTypes, m)
	diags.Append(d...)
	return t
}

// Flatten populates the Terraform model from SDK type
func (m *FixedaddressLogicFilterRulesModel) Flatten(ctx context.Context, from *niosdhcp.FixedaddressLogicFilterRules, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.Filter = flex.FlattenStringPointerEmptyAsNull(from.Filter)
	m.Type = flex.FlattenStringPointerEmptyAsNull(from.Type)
}
