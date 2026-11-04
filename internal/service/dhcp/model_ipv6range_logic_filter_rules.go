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

// Ipv6rangeLogicFilterRulesModel is the Terraform model for Ipv6rangeLogicFilterRules
type Ipv6rangeLogicFilterRulesModel struct {
	Filter types.String `tfsdk:"filter"`
	Type   types.String `tfsdk:"type"`
}

// Ipv6rangeLogicFilterRulesAttrTypes contains the attribute types for Ipv6rangeLogicFilterRulesModel
var Ipv6rangeLogicFilterRulesAttrTypes = map[string]attr.Type{
	"filter": types.StringType,
	"type":   types.StringType,
}

// Ipv6rangeLogicFilterRulesResourceSchemaAttributes contains the schema attributes for Ipv6rangeLogicFilterRulesModel
var Ipv6rangeLogicFilterRulesResourceSchemaAttributes = map[string]schema.Attribute{
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

// ExpandIpv6rangeLogicFilterRules converts a Terraform Object to SDK type
func ExpandIpv6rangeLogicFilterRules(ctx context.Context, o types.Object, diags *diag.Diagnostics) *niosdhcp.Ipv6rangeLogicFilterRules {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m Ipv6rangeLogicFilterRulesModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

// Expand converts the Terraform model to SDK type
func (m *Ipv6rangeLogicFilterRulesModel) Expand(ctx context.Context, diags *diag.Diagnostics) *niosdhcp.Ipv6rangeLogicFilterRules {
	if m == nil {
		return nil
	}
	to := &niosdhcp.Ipv6rangeLogicFilterRules{
		Filter: flex.ExpandStringPointerNullAsEmpty(m.Filter),
		Type:   flex.ExpandStringPointerNullAsEmpty(m.Type),
	}
	return to
}

// FlattenIpv6rangeLogicFilterRules converts an SDK type to Terraform Object
func FlattenIpv6rangeLogicFilterRules(ctx context.Context, from *niosdhcp.Ipv6rangeLogicFilterRules, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(Ipv6rangeLogicFilterRulesAttrTypes)
	}
	m := &Ipv6rangeLogicFilterRulesModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, Ipv6rangeLogicFilterRulesAttrTypes, m)
	diags.Append(d...)
	return t
}

// Flatten populates the Terraform model from SDK type
func (m *Ipv6rangeLogicFilterRulesModel) Flatten(ctx context.Context, from *niosdhcp.Ipv6rangeLogicFilterRules, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.Filter = flex.FlattenStringPointerEmptyAsNull(from.Filter)
	m.Type = flex.FlattenStringPointerEmptyAsNull(from.Type)
}
