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

// Ipv6rangecloudinfoDelegatedMemberModel is the Terraform model for Ipv6rangecloudinfoDelegatedMember
type Ipv6rangecloudinfoDelegatedMemberModel struct {
	Ipv4addr types.String `tfsdk:"ipv4addr"`
	Ipv6addr types.String `tfsdk:"ipv6addr"`
	Name     types.String `tfsdk:"name"`
}

// Ipv6rangecloudinfoDelegatedMemberAttrTypes contains the attribute types for Ipv6rangecloudinfoDelegatedMemberModel
var Ipv6rangecloudinfoDelegatedMemberAttrTypes = map[string]attr.Type{
	"ipv4addr": types.StringType,
	"ipv6addr": types.StringType,
	"name":     types.StringType,
}

// Ipv6rangecloudinfoDelegatedMemberResourceSchemaAttributes contains the schema attributes for Ipv6rangecloudinfoDelegatedMemberModel
var Ipv6rangecloudinfoDelegatedMemberResourceSchemaAttributes = map[string]schema.Attribute{
	"ipv4addr": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "The IPv4 Address of the Grid Member.",
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

// ExpandIpv6rangecloudinfoDelegatedMember converts a Terraform Object to SDK type
func ExpandIpv6rangecloudinfoDelegatedMember(ctx context.Context, o types.Object, diags *diag.Diagnostics) *niosdhcp.Ipv6rangecloudinfoDelegatedMember {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m Ipv6rangecloudinfoDelegatedMemberModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

// Expand converts the Terraform model to SDK type
func (m *Ipv6rangecloudinfoDelegatedMemberModel) Expand(ctx context.Context, diags *diag.Diagnostics) *niosdhcp.Ipv6rangecloudinfoDelegatedMember {
	if m == nil {
		return nil
	}
	to := &niosdhcp.Ipv6rangecloudinfoDelegatedMember{
		Ipv4addr: flex.ExpandStringPointerNullAsEmpty(m.Ipv4addr),
		Ipv6addr: flex.ExpandStringPointerNullAsEmpty(m.Ipv6addr),
		Name:     flex.ExpandStringPointerNullAsEmpty(m.Name),
	}
	return to
}

// FlattenIpv6rangecloudinfoDelegatedMember converts an SDK type to Terraform Object
func FlattenIpv6rangecloudinfoDelegatedMember(ctx context.Context, from *niosdhcp.Ipv6rangecloudinfoDelegatedMember, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(Ipv6rangecloudinfoDelegatedMemberAttrTypes)
	}
	m := &Ipv6rangecloudinfoDelegatedMemberModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, Ipv6rangecloudinfoDelegatedMemberAttrTypes, m)
	diags.Append(d...)
	return t
}

// Flatten populates the Terraform model from SDK type
func (m *Ipv6rangecloudinfoDelegatedMemberModel) Flatten(ctx context.Context, from *niosdhcp.Ipv6rangecloudinfoDelegatedMember, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.Ipv4addr = flex.FlattenStringPointerEmptyAsNull(from.Ipv4addr)
	m.Ipv6addr = flex.FlattenStringPointerEmptyAsNull(from.Ipv6addr)
	m.Name = flex.FlattenStringPointerEmptyAsNull(from.Name)
}
