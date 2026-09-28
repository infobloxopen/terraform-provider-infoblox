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

// RangetemplateMemberModel is the Terraform model for RangetemplateMember
type RangetemplateMemberModel struct {
	Ipv4addr types.String `tfsdk:"ipv4addr"`
	Ipv6addr types.String `tfsdk:"ipv6addr"`
	Name     types.String `tfsdk:"name"`
}

// RangetemplateMemberAttrTypes contains the attribute types for RangetemplateMemberModel
var RangetemplateMemberAttrTypes = map[string]attr.Type{
	"ipv4addr": types.StringType,
	"ipv6addr": types.StringType,
	"name":     types.StringType,
}

// RangetemplateMemberResourceSchemaAttributes contains the schema attributes for RangetemplateMemberModel
var RangetemplateMemberResourceSchemaAttributes = map[string]schema.Attribute{
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

// ExpandRangetemplateMember converts a Terraform Object to SDK type
func ExpandRangetemplateMember(ctx context.Context, o types.Object, diags *diag.Diagnostics) *niosdhcp.RangetemplateMember {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m RangetemplateMemberModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

// Expand converts the Terraform model to SDK type
func (m *RangetemplateMemberModel) Expand(ctx context.Context, diags *diag.Diagnostics) *niosdhcp.RangetemplateMember {
	if m == nil {
		return nil
	}
	to := &niosdhcp.RangetemplateMember{
		Ipv4addr: flex.ExpandStringPointer(m.Ipv4addr),
		Ipv6addr: flex.ExpandStringPointer(m.Ipv6addr),
		Name:     flex.ExpandStringPointer(m.Name),
	}
	return to
}

// FlattenRangetemplateMember converts an SDK type to Terraform Object
func FlattenRangetemplateMember(ctx context.Context, from *niosdhcp.RangetemplateMember, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(RangetemplateMemberAttrTypes)
	}
	m := &RangetemplateMemberModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, RangetemplateMemberAttrTypes, m)
	diags.Append(d...)
	return t
}

// Flatten populates the Terraform model from SDK type
func (m *RangetemplateMemberModel) Flatten(ctx context.Context, from *niosdhcp.RangetemplateMember, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.Ipv4addr = flex.FlattenStringPointerEmptyAsNull(from.Ipv4addr)
	m.Ipv6addr = flex.FlattenStringPointerEmptyAsNull(from.Ipv6addr)
	m.Name = flex.FlattenStringPointerEmptyAsNull(from.Name)
}
