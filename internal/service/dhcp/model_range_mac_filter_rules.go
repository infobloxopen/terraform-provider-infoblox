package dhcp

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
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

// RangeMacFilterRulesModel is the Terraform model for RangeMacFilterRules
type RangeMacFilterRulesModel struct {
	Filter     types.String `tfsdk:"filter"`
	Permission types.String `tfsdk:"permission"`
}

// RangeMacFilterRulesAttrTypes contains the attribute types for RangeMacFilterRulesModel
var RangeMacFilterRulesAttrTypes = map[string]attr.Type{
	"filter":     types.StringType,
	"permission": types.StringType,
}

// RangeMacFilterRulesResourceSchemaAttributes contains the schema attributes for RangeMacFilterRulesModel
var RangeMacFilterRulesResourceSchemaAttributes = map[string]schema.Attribute{
	"filter": schema.StringAttribute{
		Required: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "The name of the DHCP filter.",
	},
	"permission": schema.StringAttribute{
		Validators: []validator.String{
			stringvalidator.OneOf("Allow", "Deny"),
		},
		Required:            true,
		MarkdownDescription: "The permission to be applied.",
	},
}

// ExpandRangeMacFilterRules converts a Terraform Object to SDK type
func ExpandRangeMacFilterRules(ctx context.Context, o types.Object, diags *diag.Diagnostics) *niosdhcp.RangeMacFilterRules {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m RangeMacFilterRulesModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

// Expand converts the Terraform model to SDK type
func (m *RangeMacFilterRulesModel) Expand(ctx context.Context, diags *diag.Diagnostics) *niosdhcp.RangeMacFilterRules {
	if m == nil {
		return nil
	}
	to := &niosdhcp.RangeMacFilterRules{
		Filter:     flex.ExpandStringPointerNullAsEmpty(m.Filter),
		Permission: flex.ExpandStringPointerNullAsEmpty(m.Permission),
	}
	return to
}

// FlattenRangeMacFilterRules converts an SDK type to Terraform Object
func FlattenRangeMacFilterRules(ctx context.Context, from *niosdhcp.RangeMacFilterRules, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(RangeMacFilterRulesAttrTypes)
	}
	m := &RangeMacFilterRulesModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, RangeMacFilterRulesAttrTypes, m)
	diags.Append(d...)
	return t
}

// Flatten populates the Terraform model from SDK type
func (m *RangeMacFilterRulesModel) Flatten(ctx context.Context, from *niosdhcp.RangeMacFilterRules, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.Filter = flex.FlattenStringPointerEmptyAsNull(from.Filter)
	m.Permission = flex.FlattenStringPointerEmptyAsNull(from.Permission)
}
