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

// Ipv6rangeOptionFilterRulesModel is the Terraform model for Ipv6rangeOptionFilterRules
type Ipv6rangeOptionFilterRulesModel struct {
	Filter     types.String `tfsdk:"filter"`
	Permission types.String `tfsdk:"permission"`
}

// Ipv6rangeOptionFilterRulesAttrTypes contains the attribute types for Ipv6rangeOptionFilterRulesModel
var Ipv6rangeOptionFilterRulesAttrTypes = map[string]attr.Type{
	"filter":     types.StringType,
	"permission": types.StringType,
}

// Ipv6rangeOptionFilterRulesResourceSchemaAttributes contains the schema attributes for Ipv6rangeOptionFilterRulesModel
var Ipv6rangeOptionFilterRulesResourceSchemaAttributes = map[string]schema.Attribute{
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

// ExpandIpv6rangeOptionFilterRules converts a Terraform Object to SDK type
func ExpandIpv6rangeOptionFilterRules(ctx context.Context, o types.Object, diags *diag.Diagnostics) *niosdhcp.Ipv6rangeOptionFilterRules {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m Ipv6rangeOptionFilterRulesModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

// Expand converts the Terraform model to SDK type
func (m *Ipv6rangeOptionFilterRulesModel) Expand(ctx context.Context, diags *diag.Diagnostics) *niosdhcp.Ipv6rangeOptionFilterRules {
	if m == nil {
		return nil
	}
	to := &niosdhcp.Ipv6rangeOptionFilterRules{
		Filter:     flex.ExpandStringPointerNullAsEmpty(m.Filter),
		Permission: flex.ExpandStringPointerNullAsEmpty(m.Permission),
	}
	return to
}

// FlattenIpv6rangeOptionFilterRules converts an SDK type to Terraform Object
func FlattenIpv6rangeOptionFilterRules(ctx context.Context, from *niosdhcp.Ipv6rangeOptionFilterRules, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(Ipv6rangeOptionFilterRulesAttrTypes)
	}
	m := &Ipv6rangeOptionFilterRulesModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, Ipv6rangeOptionFilterRulesAttrTypes, m)
	diags.Append(d...)
	return t
}

// Flatten populates the Terraform model from SDK type
func (m *Ipv6rangeOptionFilterRulesModel) Flatten(ctx context.Context, from *niosdhcp.Ipv6rangeOptionFilterRules, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.Filter = flex.FlattenStringPointerEmptyAsNull(from.Filter)
	m.Permission = flex.FlattenStringPointerEmptyAsNull(from.Permission)
}
