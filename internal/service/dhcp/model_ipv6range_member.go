package dhcp

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/hashicorp/terraform-plugin-framework-nettypes/iptypes"
	niosdhcp "github.com/infobloxopen/infoblox-nios-go-client/dhcp"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
	customvalidator "github.com/infobloxopen/terraform-provider-infoblox/internal/validator"
)

// Ipv6rangeMemberModel is the Terraform model for Ipv6rangeMember
type Ipv6rangeMemberModel struct {
	Ipv4addr iptypes.IPv4Address `tfsdk:"ipv4addr"`
	Ipv6addr iptypes.IPv6Address `tfsdk:"ipv6addr"`
	Name     types.String        `tfsdk:"name"`
}

// Ipv6rangeMemberAttrTypes contains the attribute types for Ipv6rangeMemberModel
var Ipv6rangeMemberAttrTypes = map[string]attr.Type{
	"ipv4addr": iptypes.IPv4AddressType{},
	"ipv6addr": iptypes.IPv6AddressType{},
	"name":     types.StringType,
}

// Ipv6rangeMemberResourceSchemaAttributes contains the schema attributes for Ipv6rangeMemberModel
var Ipv6rangeMemberResourceSchemaAttributes = map[string]schema.Attribute{
	"ipv4addr": schema.StringAttribute{
		Optional:   true,
		Computed:   true,
		CustomType: iptypes.IPv4AddressType{},
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "The IPv4 Address of the Grid Member.",
	},
	"ipv6addr": schema.StringAttribute{
		Optional:   true,
		Computed:   true,
		CustomType: iptypes.IPv6AddressType{},
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "The IPv6 Address of the Grid Member.",
	},
	"name": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "The Grid member name",
	},
}

// ExpandIpv6rangeMember converts a Terraform Object to SDK type
func ExpandIpv6rangeMember(ctx context.Context, o types.Object, diags *diag.Diagnostics) *niosdhcp.Ipv6rangeMember {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m Ipv6rangeMemberModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

// Expand converts the Terraform model to SDK type
func (m *Ipv6rangeMemberModel) Expand(ctx context.Context, diags *diag.Diagnostics) *niosdhcp.Ipv6rangeMember {
	if m == nil {
		return nil
	}
	to := &niosdhcp.Ipv6rangeMember{
		Ipv4addr: flex.ExpandIPv4Address(m.Ipv4addr),
		Ipv6addr: flex.ExpandIPv6Address(m.Ipv6addr),
		Name:     flex.ExpandStringPointer(m.Name),
	}
	return to
}

// FlattenIpv6rangeMember converts an SDK type to Terraform Object
func FlattenIpv6rangeMember(ctx context.Context, from *niosdhcp.Ipv6rangeMember, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(Ipv6rangeMemberAttrTypes)
	}
	m := &Ipv6rangeMemberModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, Ipv6rangeMemberAttrTypes, m)
	diags.Append(d...)
	return t
}

// Flatten populates the Terraform model from SDK type
func (m *Ipv6rangeMemberModel) Flatten(ctx context.Context, from *niosdhcp.Ipv6rangeMember, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.Ipv4addr = flex.FlattenIPv4Address(from.Ipv4addr)
	m.Ipv6addr = flex.FlattenIPv6Address(from.Ipv6addr)
	m.Name = flex.FlattenStringPointerEmptyAsNull(from.Name)
}
