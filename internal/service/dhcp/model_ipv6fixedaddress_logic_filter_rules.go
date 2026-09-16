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

// Ipv6fixedaddressLogicFilterRulesModel is the Terraform model for Ipv6fixedaddressLogicFilterRules
type Ipv6fixedaddressLogicFilterRulesModel struct {
	Filter types.String `tfsdk:"filter"`
	Type   types.String `tfsdk:"type"`
}

// Ipv6fixedaddressLogicFilterRulesAttrTypes contains the attribute types for Ipv6fixedaddressLogicFilterRulesModel
var Ipv6fixedaddressLogicFilterRulesAttrTypes = map[string]attr.Type{
	"filter": types.StringType,
	"type":   types.StringType,
}

// Ipv6fixedaddressLogicFilterRulesResourceSchemaAttributes contains the schema attributes for Ipv6fixedaddressLogicFilterRulesModel
var Ipv6fixedaddressLogicFilterRulesResourceSchemaAttributes = map[string]schema.Attribute{
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

// ExpandIpv6fixedaddressLogicFilterRules converts a Terraform Object to SDK type
func ExpandIpv6fixedaddressLogicFilterRules(ctx context.Context, o types.Object, diags *diag.Diagnostics) *niosdhcp.Ipv6fixedaddressLogicFilterRules {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m Ipv6fixedaddressLogicFilterRulesModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

// Expand converts the Terraform model to SDK type
func (m *Ipv6fixedaddressLogicFilterRulesModel) Expand(ctx context.Context, diags *diag.Diagnostics) *niosdhcp.Ipv6fixedaddressLogicFilterRules {
	if m == nil {
		return nil
	}
	to := &niosdhcp.Ipv6fixedaddressLogicFilterRules{
		Filter: flex.ExpandStringPointerNullAsEmpty(m.Filter),
		Type:   flex.ExpandStringPointerNullAsEmpty(m.Type),
	}
	return to
}

// FlattenIpv6fixedaddressLogicFilterRules converts an SDK type to Terraform Object
func FlattenIpv6fixedaddressLogicFilterRules(ctx context.Context, from *niosdhcp.Ipv6fixedaddressLogicFilterRules, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(Ipv6fixedaddressLogicFilterRulesAttrTypes)
	}
	m := &Ipv6fixedaddressLogicFilterRulesModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, Ipv6fixedaddressLogicFilterRulesAttrTypes, m)
	diags.Append(d...)
	return t
}

// Flatten populates the Terraform model from SDK type
func (m *Ipv6fixedaddressLogicFilterRulesModel) Flatten(ctx context.Context, from *niosdhcp.Ipv6fixedaddressLogicFilterRules, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.Filter = flex.FlattenStringPointerEmptyAsNull(from.Filter)
	m.Type = flex.FlattenStringPointerEmptyAsNull(from.Type)
}
