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

// Ipv6networkVlansModel is the Terraform model for Ipv6networkVlans
type Ipv6networkVlansModel struct {
	Vlan types.String `tfsdk:"vlan"`
}

// Ipv6networkVlansAttrTypes contains the attribute types for Ipv6networkVlansModel
var Ipv6networkVlansAttrTypes = map[string]attr.Type{
	"vlan": types.StringType,
}

// Ipv6networkVlansResourceSchemaAttributes contains the schema attributes for Ipv6networkVlansModel
var Ipv6networkVlansResourceSchemaAttributes = map[string]schema.Attribute{
	"vlan": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "Reference to the underlying StaticVlan object vlan.",
	},
}

// ExpandIpv6networkVlans converts a Terraform Object to SDK type
func ExpandIpv6networkVlans(ctx context.Context, o types.Object, diags *diag.Diagnostics) *niosipam.Ipv6networkVlans {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m Ipv6networkVlansModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

// Expand converts the Terraform model to SDK type
func (m *Ipv6networkVlansModel) Expand(ctx context.Context, diags *diag.Diagnostics) *niosipam.Ipv6networkVlans {
	if m == nil {
		return nil
	}
	to := &niosipam.Ipv6networkVlans{
		Vlan: flex.ExpandStringPointerNullAsEmpty(m.Vlan),
	}
	return to
}

// FlattenIpv6networkVlans converts an SDK type to Terraform Object
func FlattenIpv6networkVlans(ctx context.Context, from *niosipam.Ipv6networkVlans, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(Ipv6networkVlansAttrTypes)
	}
	m := &Ipv6networkVlansModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, Ipv6networkVlansAttrTypes, m)
	diags.Append(d...)
	return t
}

// Flatten populates the Terraform model from SDK type
func (m *Ipv6networkVlansModel) Flatten(ctx context.Context, from *niosipam.Ipv6networkVlans, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.Vlan = flex.FlattenStringPointerEmptyAsNull(from.Vlan)
}
