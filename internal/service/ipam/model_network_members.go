package ipam

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
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

// NetworkMembersModel is the Terraform model for NetworkMembers
type NetworkMembersModel struct {
	Struct   types.String `tfsdk:"struct"`
	Ipv4addr types.String `tfsdk:"ipv4addr"`
	Ipv6addr types.String `tfsdk:"ipv6addr"`
	Name     types.String `tfsdk:"name"`
}

// NetworkMembersAttrTypes contains the attribute types for NetworkMembersModel
var NetworkMembersAttrTypes = map[string]attr.Type{
	"struct":   types.StringType,
	"ipv4addr": types.StringType,
	"ipv6addr": types.StringType,
	"name":     types.StringType,
}

// NetworkMembersResourceSchemaAttributes contains the schema attributes for NetworkMembersModel
var NetworkMembersResourceSchemaAttributes = map[string]schema.Attribute{
	"struct": schema.StringAttribute{
		Validators: []validator.String{
			stringvalidator.OneOf("dhcpmember", "msdhcpserver"),
		},
		Required:            true,
		MarkdownDescription: "The struct type of the object. The value must be one of 'dhcpmember' and 'msdhcpserver'.",
	},
	"ipv4addr": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
			customvalidator.IsValidIPv4OrFQDN(),
		},
		MarkdownDescription: "The IPv4 Address or FQDN of the Microsoft server.",
	},
	"ipv6addr": schema.StringAttribute{
		Optional: true,
		Computed: true,
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

// ExpandNetworkMembers converts a Terraform Object to SDK type
func ExpandNetworkMembers(ctx context.Context, o types.Object, diags *diag.Diagnostics) *niosipam.NetworkMembers {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m NetworkMembersModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

// Expand converts the Terraform model to SDK type
func (m *NetworkMembersModel) Expand(ctx context.Context, diags *diag.Diagnostics) *niosipam.NetworkMembers {
	if m == nil {
		return nil
	}
	to := &niosipam.NetworkMembers{
		Struct:   flex.ExpandStringPointerNullAsEmpty(m.Struct),
		Ipv4addr: flex.ExpandStringPointer(m.Ipv4addr),
		Ipv6addr: flex.ExpandStringPointer(m.Ipv6addr),
		Name:     flex.ExpandStringPointer(m.Name),
	}
	return to
}

// FlattenNetworkMembers converts an SDK type to Terraform Object
func FlattenNetworkMembers(ctx context.Context, from *niosipam.NetworkMembers, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(NetworkMembersAttrTypes)
	}
	m := &NetworkMembersModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, NetworkMembersAttrTypes, m)
	diags.Append(d...)
	return t
}

// Flatten populates the Terraform model from SDK type
func (m *NetworkMembersModel) Flatten(ctx context.Context, from *niosipam.NetworkMembers, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.Struct = flex.FlattenStringPointerEmptyAsNull(from.Struct)
	m.Ipv4addr = flex.FlattenStringPointerEmptyAsNull(from.Ipv4addr)
	m.Ipv6addr = flex.FlattenStringPointerEmptyAsNull(from.Ipv6addr)
	m.Name = flex.FlattenStringPointerEmptyAsNull(from.Name)
}
