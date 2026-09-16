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

// Ipv6sharednetworkNetworksModel is the Terraform model for Ipv6sharednetworkNetworks
type Ipv6sharednetworkNetworksModel struct {
	Ref types.String `tfsdk:"ref"`
}

// Ipv6sharednetworkNetworksAttrTypes contains the attribute types for Ipv6sharednetworkNetworksModel
var Ipv6sharednetworkNetworksAttrTypes = map[string]attr.Type{
	"ref": types.StringType,
}

// Ipv6sharednetworkNetworksResourceSchemaAttributes contains the schema attributes for Ipv6sharednetworkNetworksModel
var Ipv6sharednetworkNetworksResourceSchemaAttributes = map[string]schema.Attribute{
	"ref": schema.StringAttribute{
		Optional: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "Reference to the IPv6 Network.",
	},
}

// ExpandIpv6sharednetworkNetworks converts a Terraform Object to SDK type
func ExpandIpv6sharednetworkNetworks(ctx context.Context, o types.Object, diags *diag.Diagnostics) *niosdhcp.Ipv6sharednetworkNetworks {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m Ipv6sharednetworkNetworksModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

// Expand converts the Terraform model to SDK type
func (m *Ipv6sharednetworkNetworksModel) Expand(ctx context.Context, diags *diag.Diagnostics) *niosdhcp.Ipv6sharednetworkNetworks {
	if m == nil {
		return nil
	}
	to := &niosdhcp.Ipv6sharednetworkNetworks{
		Ref: flex.ExpandStringPointerNullAsEmpty(m.Ref),
	}
	return to
}

// FlattenIpv6sharednetworkNetworks converts an SDK type to Terraform Object
func FlattenIpv6sharednetworkNetworks(ctx context.Context, from *niosdhcp.Ipv6sharednetworkNetworks, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(Ipv6sharednetworkNetworksAttrTypes)
	}
	m := &Ipv6sharednetworkNetworksModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, Ipv6sharednetworkNetworksAttrTypes, m)
	diags.Append(d...)
	return t
}

// Flatten populates the Terraform model from SDK type
func (m *Ipv6sharednetworkNetworksModel) Flatten(ctx context.Context, from *niosdhcp.Ipv6sharednetworkNetworks, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.Ref = flex.FlattenStringPointerEmptyAsNull(from.Ref)
}
