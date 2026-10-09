package ipam

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	niosipam "github.com/infobloxopen/infoblox-nios-go-client/ipam"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
	customvalidator "github.com/infobloxopen/terraform-provider-infoblox/internal/validator"
)

// Ipv6networkcontainerLogicFilterRulesModel is the Terraform model for Ipv6networkcontainerLogicFilterRules
type Ipv6networkcontainerLogicFilterRulesModel struct {
	Filter types.String `tfsdk:"filter"`
	Type   types.String `tfsdk:"type"`
}

// Ipv6networkcontainerLogicFilterRulesAttrTypes contains the attribute types for Ipv6networkcontainerLogicFilterRulesModel
var Ipv6networkcontainerLogicFilterRulesAttrTypes = map[string]attr.Type{
	"filter": types.StringType,
	"type":   types.StringType,
}

// Ipv6networkcontainerLogicFilterRulesResourceSchemaAttributes contains the schema attributes for Ipv6networkcontainerLogicFilterRulesModel
var Ipv6networkcontainerLogicFilterRulesResourceSchemaAttributes = map[string]schema.Attribute{
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

// ExpandIpv6networkcontainerLogicFilterRules converts a Terraform Object to SDK type
func ExpandIpv6networkcontainerLogicFilterRules(ctx context.Context, o types.Object, diags *diag.Diagnostics) *niosipam.Ipv6networkcontainerLogicFilterRules {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m Ipv6networkcontainerLogicFilterRulesModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

// Expand converts the Terraform model to SDK type
func (m *Ipv6networkcontainerLogicFilterRulesModel) Expand(ctx context.Context, diags *diag.Diagnostics) *niosipam.Ipv6networkcontainerLogicFilterRules {
	if m == nil {
		return nil
	}
	to := &niosipam.Ipv6networkcontainerLogicFilterRules{
		Filter: flex.ExpandStringPointerNullAsEmpty(m.Filter),
		Type:   flex.ExpandStringPointerNullAsEmpty(m.Type),
	}
	return to
}

// FlattenIpv6networkcontainerLogicFilterRules converts an SDK type to Terraform Object
func FlattenIpv6networkcontainerLogicFilterRules(ctx context.Context, from *niosipam.Ipv6networkcontainerLogicFilterRules, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(Ipv6networkcontainerLogicFilterRulesAttrTypes)
	}
	m := &Ipv6networkcontainerLogicFilterRulesModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, Ipv6networkcontainerLogicFilterRulesAttrTypes, m)
	diags.Append(d...)
	return t
}

// Flatten populates the Terraform model from SDK type
func (m *Ipv6networkcontainerLogicFilterRulesModel) Flatten(ctx context.Context, from *niosipam.Ipv6networkcontainerLogicFilterRules, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.Filter = flex.FlattenStringPointerEmptyAsNull(from.Filter)
	m.Type = flex.FlattenStringPointerEmptyAsNull(from.Type)
}
