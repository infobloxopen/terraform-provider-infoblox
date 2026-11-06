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

// FixedaddresstemplateLogicFilterRulesModel is the Terraform model for FixedaddresstemplateLogicFilterRules
type FixedaddresstemplateLogicFilterRulesModel struct {
	Filter types.String `tfsdk:"filter"`
	Type   types.String `tfsdk:"type"`
}

// FixedaddresstemplateLogicFilterRulesAttrTypes contains the attribute types for FixedaddresstemplateLogicFilterRulesModel
var FixedaddresstemplateLogicFilterRulesAttrTypes = map[string]attr.Type{
	"filter": types.StringType,
	"type":   types.StringType,
}

// FixedaddresstemplateLogicFilterRulesResourceSchemaAttributes contains the schema attributes for FixedaddresstemplateLogicFilterRulesModel
var FixedaddresstemplateLogicFilterRulesResourceSchemaAttributes = map[string]schema.Attribute{
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

// ExpandFixedaddresstemplateLogicFilterRules converts a Terraform Object to SDK type
func ExpandFixedaddresstemplateLogicFilterRules(ctx context.Context, o types.Object, diags *diag.Diagnostics) *niosdhcp.FixedaddresstemplateLogicFilterRules {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m FixedaddresstemplateLogicFilterRulesModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

// Expand converts the Terraform model to SDK type
func (m *FixedaddresstemplateLogicFilterRulesModel) Expand(ctx context.Context, diags *diag.Diagnostics) *niosdhcp.FixedaddresstemplateLogicFilterRules {
	if m == nil {
		return nil
	}
	to := &niosdhcp.FixedaddresstemplateLogicFilterRules{
		Filter: flex.ExpandStringPointerNullAsEmpty(m.Filter),
		Type:   flex.ExpandStringPointerNullAsEmpty(m.Type),
	}
	return to
}

// FlattenFixedaddresstemplateLogicFilterRules converts an SDK type to Terraform Object
func FlattenFixedaddresstemplateLogicFilterRules(ctx context.Context, from *niosdhcp.FixedaddresstemplateLogicFilterRules, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(FixedaddresstemplateLogicFilterRulesAttrTypes)
	}
	m := &FixedaddresstemplateLogicFilterRulesModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, FixedaddresstemplateLogicFilterRulesAttrTypes, m)
	diags.Append(d...)
	return t
}

// Flatten populates the Terraform model from SDK type
func (m *FixedaddresstemplateLogicFilterRulesModel) Flatten(ctx context.Context, from *niosdhcp.FixedaddresstemplateLogicFilterRules, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.Filter = flex.FlattenStringPointerEmptyAsNull(from.Filter)
	m.Type = flex.FlattenStringPointerEmptyAsNull(from.Type)
}
