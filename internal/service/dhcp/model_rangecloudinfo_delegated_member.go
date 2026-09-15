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

// RangecloudinfoDelegatedMemberModel is the Terraform model for RangecloudinfoDelegatedMember
type RangecloudinfoDelegatedMemberModel struct {
	Ipv4addr types.String `tfsdk:"ipv4addr"`
	Ipv6addr types.String `tfsdk:"ipv6addr"`
	Name     types.String `tfsdk:"name"`
}

// RangecloudinfoDelegatedMemberAttrTypes contains the attribute types for RangecloudinfoDelegatedMemberModel
var RangecloudinfoDelegatedMemberAttrTypes = map[string]attr.Type{
	"ipv4addr": types.StringType,
	"ipv6addr": types.StringType,
	"name":     types.StringType,
}

// RangecloudinfoDelegatedMemberResourceSchemaAttributes contains the schema attributes for RangecloudinfoDelegatedMemberModel
var RangecloudinfoDelegatedMemberResourceSchemaAttributes = map[string]schema.Attribute{
	"ipv4addr": schema.StringAttribute{
		Optional: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "The IPv4 Address of the Grid Member.",
	},
	"ipv6addr": schema.StringAttribute{
		Optional: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "The IPv6 Address of the Grid Member.",
	},
	"name": schema.StringAttribute{
		Optional: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "The Grid member name",
	},
}

// ExpandRangecloudinfoDelegatedMember converts a Terraform Object to SDK type
func ExpandRangecloudinfoDelegatedMember(ctx context.Context, o types.Object, diags *diag.Diagnostics) *niosdhcp.RangecloudinfoDelegatedMember {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m RangecloudinfoDelegatedMemberModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

// Expand converts the Terraform model to SDK type
func (m *RangecloudinfoDelegatedMemberModel) Expand(ctx context.Context, diags *diag.Diagnostics) *niosdhcp.RangecloudinfoDelegatedMember {
	if m == nil {
		return nil
	}
	to := &niosdhcp.RangecloudinfoDelegatedMember{
		Ipv4addr: flex.ExpandStringPointerNullAsEmpty(m.Ipv4addr),
		Ipv6addr: flex.ExpandStringPointerNullAsEmpty(m.Ipv6addr),
		Name:     flex.ExpandStringPointerNullAsEmpty(m.Name),
	}
	return to
}

// FlattenRangecloudinfoDelegatedMember converts an SDK type to Terraform Object
func FlattenRangecloudinfoDelegatedMember(ctx context.Context, from *niosdhcp.RangecloudinfoDelegatedMember, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(RangecloudinfoDelegatedMemberAttrTypes)
	}
	m := &RangecloudinfoDelegatedMemberModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, RangecloudinfoDelegatedMemberAttrTypes, m)
	diags.Append(d...)
	return t
}

// Flatten populates the Terraform model from SDK type
func (m *RangecloudinfoDelegatedMemberModel) Flatten(ctx context.Context, from *niosdhcp.RangecloudinfoDelegatedMember, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.Ipv4addr = flex.FlattenStringPointerEmptyAsNull(from.Ipv4addr)
	m.Ipv6addr = flex.FlattenStringPointerEmptyAsNull(from.Ipv6addr)
	m.Name = flex.FlattenStringPointerEmptyAsNull(from.Name)
}
