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

// RangeSplitMemberModel is the Terraform model for RangeSplitMember
type RangeSplitMemberModel struct {
	Ipv4addr types.String `tfsdk:"ipv4addr"`
}

// RangeSplitMemberAttrTypes contains the attribute types for RangeSplitMemberModel
var RangeSplitMemberAttrTypes = map[string]attr.Type{
	"ipv4addr": types.StringType,
}

// RangeSplitMemberResourceSchemaAttributes contains the schema attributes for RangeSplitMemberModel
var RangeSplitMemberResourceSchemaAttributes = map[string]schema.Attribute{
	"ipv4addr": schema.StringAttribute{
		Required: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
			customvalidator.IsValidIPv4OrFQDN(),
		},
		MarkdownDescription: "The IPv4 Address or FQDN of the Microsoft server.",
	},
}

// ExpandRangeSplitMember converts a Terraform Object to SDK type
func ExpandRangeSplitMember(ctx context.Context, o types.Object, diags *diag.Diagnostics) *niosdhcp.RangeSplitMember {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m RangeSplitMemberModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

// Expand converts the Terraform model to SDK type
func (m *RangeSplitMemberModel) Expand(ctx context.Context, diags *diag.Diagnostics) *niosdhcp.RangeSplitMember {
	if m == nil {
		return nil
	}
	to := &niosdhcp.RangeSplitMember{
		Ipv4addr: flex.ExpandStringPointer(m.Ipv4addr),
	}
	return to
}

// FlattenRangeSplitMember converts an SDK type to Terraform Object
func FlattenRangeSplitMember(ctx context.Context, from *niosdhcp.RangeSplitMember, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(RangeSplitMemberAttrTypes)
	}
	m := &RangeSplitMemberModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, RangeSplitMemberAttrTypes, m)
	diags.Append(d...)
	return t
}

// Flatten populates the Terraform model from SDK type
func (m *RangeSplitMemberModel) Flatten(ctx context.Context, from *niosdhcp.RangeSplitMember, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.Ipv4addr = flex.FlattenStringPointerEmptyAsNull(from.Ipv4addr)
}
