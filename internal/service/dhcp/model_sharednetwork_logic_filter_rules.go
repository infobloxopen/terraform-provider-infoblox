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

// SharednetworkLogicFilterRulesModel is the Terraform model for SharednetworkLogicFilterRules
type SharednetworkLogicFilterRulesModel struct {
	Filter types.String `tfsdk:"filter"`
	Type   types.String `tfsdk:"type"`
}

// SharednetworkLogicFilterRulesAttrTypes contains the attribute types for SharednetworkLogicFilterRulesModel
var SharednetworkLogicFilterRulesAttrTypes = map[string]attr.Type{
	"filter": types.StringType,
	"type":   types.StringType,
}

// SharednetworkLogicFilterRulesResourceSchemaAttributes contains the schema attributes for SharednetworkLogicFilterRulesModel
var SharednetworkLogicFilterRulesResourceSchemaAttributes = map[string]schema.Attribute{
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

// ExpandSharednetworkLogicFilterRules converts a Terraform Object to SDK type
func ExpandSharednetworkLogicFilterRules(ctx context.Context, o types.Object, diags *diag.Diagnostics) *niosdhcp.SharednetworkLogicFilterRules {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m SharednetworkLogicFilterRulesModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

// Expand converts the Terraform model to SDK type
func (m *SharednetworkLogicFilterRulesModel) Expand(ctx context.Context, diags *diag.Diagnostics) *niosdhcp.SharednetworkLogicFilterRules {
	if m == nil {
		return nil
	}
	to := &niosdhcp.SharednetworkLogicFilterRules{
		Filter: flex.ExpandStringPointerNullAsEmpty(m.Filter),
		Type:   flex.ExpandStringPointerNullAsEmpty(m.Type),
	}
	return to
}

// FlattenSharednetworkLogicFilterRules converts an SDK type to Terraform Object
func FlattenSharednetworkLogicFilterRules(ctx context.Context, from *niosdhcp.SharednetworkLogicFilterRules, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(SharednetworkLogicFilterRulesAttrTypes)
	}
	m := &SharednetworkLogicFilterRulesModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, SharednetworkLogicFilterRulesAttrTypes, m)
	diags.Append(d...)
	return t
}

// Flatten populates the Terraform model from SDK type
func (m *SharednetworkLogicFilterRulesModel) Flatten(ctx context.Context, from *niosdhcp.SharednetworkLogicFilterRules, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.Filter = flex.FlattenStringPointerEmptyAsNull(from.Filter)
	m.Type = flex.FlattenStringPointerEmptyAsNull(from.Type)
}
