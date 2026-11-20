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
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	niosdhcp "github.com/infobloxopen/infoblox-nios-go-client/dhcp"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
	customvalidator "github.com/infobloxopen/terraform-provider-infoblox/internal/validator"
)

// Ipv6rangeExcludeModel is the Terraform model for Ipv6rangeExclude
type Ipv6rangeExcludeModel struct {
	StartAddress iptypes.IPv6Address `tfsdk:"start_address"`
	EndAddress   iptypes.IPv6Address `tfsdk:"end_address"`
	Comment      types.String        `tfsdk:"comment"`
}

// Ipv6rangeExcludeAttrTypes contains the attribute types for Ipv6rangeExcludeModel
var Ipv6rangeExcludeAttrTypes = map[string]attr.Type{
	"start_address": iptypes.IPv6AddressType{},
	"end_address":   iptypes.IPv6AddressType{},
	"comment":       types.StringType,
}

// Ipv6rangeExcludeResourceSchemaAttributes contains the schema attributes for Ipv6rangeExcludeModel
var Ipv6rangeExcludeResourceSchemaAttributes = map[string]schema.Attribute{
	"start_address": schema.StringAttribute{
		Required:   true,
		CustomType: iptypes.IPv6AddressType{},
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "The IPv4 Address starting address of the exclusion range.",
	},
	"end_address": schema.StringAttribute{
		Required:   true,
		CustomType: iptypes.IPv6AddressType{},
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "The IPv4 Address ending address of the exclusion range.",
	},
	"comment": schema.StringAttribute{
		Optional: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
			stringvalidator.LengthBetween(0, 256),
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "Comment for the exclusion range; maximum 256 characters.",
	},
}

// ExpandIpv6rangeExclude converts a Terraform Object to SDK type
func ExpandIpv6rangeExclude(ctx context.Context, o types.Object, diags *diag.Diagnostics) *niosdhcp.Ipv6rangeExclude {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m Ipv6rangeExcludeModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

// Expand converts the Terraform model to SDK type
func (m *Ipv6rangeExcludeModel) Expand(ctx context.Context, diags *diag.Diagnostics) *niosdhcp.Ipv6rangeExclude {
	if m == nil {
		return nil
	}
	to := &niosdhcp.Ipv6rangeExclude{
		StartAddress: flex.ExpandIPv6Address(m.StartAddress),
		EndAddress:   flex.ExpandIPv6Address(m.EndAddress),
		Comment:      flex.ExpandStringPointerNullAsEmpty(m.Comment),
	}
	return to
}

// FlattenIpv6rangeExclude converts an SDK type to Terraform Object
func FlattenIpv6rangeExclude(ctx context.Context, from *niosdhcp.Ipv6rangeExclude, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(Ipv6rangeExcludeAttrTypes)
	}
	m := &Ipv6rangeExcludeModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, Ipv6rangeExcludeAttrTypes, m)
	diags.Append(d...)
	return t
}

// Flatten populates the Terraform model from SDK type
func (m *Ipv6rangeExcludeModel) Flatten(ctx context.Context, from *niosdhcp.Ipv6rangeExclude, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.StartAddress = flex.FlattenIPv6Address(from.StartAddress)
	m.EndAddress = flex.FlattenIPv6Address(from.EndAddress)
	m.Comment = flex.FlattenStringPointerEmptyAsNull(from.Comment)
}
