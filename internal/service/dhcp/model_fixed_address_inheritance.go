package dhcp

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	uddidhcp "github.com/infobloxopen/universal-ddi-go-client/ipam"
)

// FixedAddressInheritanceModel is the Terraform model for FixedAddressInheritance
type FixedAddressInheritanceModel struct {
	DhcpOptions               types.Object `tfsdk:"dhcp_options"`
	HeaderOptionFilename      types.Object `tfsdk:"header_option_filename"`
	HeaderOptionServerAddress types.Object `tfsdk:"header_option_server_address"`
	HeaderOptionServerName    types.Object `tfsdk:"header_option_server_name"`
}

// FixedAddressInheritanceAttrTypes contains the attribute types for FixedAddressInheritanceModel
var FixedAddressInheritanceAttrTypes = map[string]attr.Type{
	"dhcp_options":                 types.ObjectType{AttrTypes: InheritedDHCPOptionListAttrTypes},
	"header_option_filename":       types.ObjectType{AttrTypes: InheritanceInheritedStringAttrTypes},
	"header_option_server_address": types.ObjectType{AttrTypes: InheritanceInheritedStringAttrTypes},
	"header_option_server_name":    types.ObjectType{AttrTypes: InheritanceInheritedStringAttrTypes},
}

// FixedAddressInheritanceResourceSchemaAttributes contains the schema attributes for FixedAddressInheritanceModel
var FixedAddressInheritanceResourceSchemaAttributes = map[string]schema.Attribute{
	"dhcp_options": schema.SingleNestedAttribute{
		Attributes:          InheritedDHCPOptionListResourceSchemaAttributes,
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The inheritance configuration for _dhcp_options_ field.",
	},
	"header_option_filename": schema.SingleNestedAttribute{
		Attributes:          InheritanceInheritedStringResourceSchemaAttributes,
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The inheritance configuration for _header_option_filename_ field.",
	},
	"header_option_server_address": schema.SingleNestedAttribute{
		Attributes:          InheritanceInheritedStringResourceSchemaAttributes,
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The inheritance configuration for _header_option_server_address_ field.",
	},
	"header_option_server_name": schema.SingleNestedAttribute{
		Attributes:          InheritanceInheritedStringResourceSchemaAttributes,
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The inheritance configuration for _header_option_server_name_ field.",
	},
}

// ExpandFixedAddressInheritance converts a Terraform Object to SDK type
func ExpandFixedAddressInheritance(ctx context.Context, o types.Object, diags *diag.Diagnostics) *uddidhcp.FixedAddressInheritance {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m FixedAddressInheritanceModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

// Expand converts the Terraform model to SDK type
func (m *FixedAddressInheritanceModel) Expand(ctx context.Context, diags *diag.Diagnostics) *uddidhcp.FixedAddressInheritance {
	if m == nil {
		return nil
	}
	to := &uddidhcp.FixedAddressInheritance{
		DhcpOptions:               ExpandInheritedDHCPOptionList(ctx, m.DhcpOptions, diags),
		HeaderOptionFilename:      ExpandInheritanceInheritedString(ctx, m.HeaderOptionFilename, diags),
		HeaderOptionServerAddress: ExpandInheritanceInheritedString(ctx, m.HeaderOptionServerAddress, diags),
		HeaderOptionServerName:    ExpandInheritanceInheritedString(ctx, m.HeaderOptionServerName, diags),
	}
	return to
}

// FlattenFixedAddressInheritance converts an SDK type to Terraform Object
func FlattenFixedAddressInheritance(ctx context.Context, from *uddidhcp.FixedAddressInheritance, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(FixedAddressInheritanceAttrTypes)
	}
	m := &FixedAddressInheritanceModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, FixedAddressInheritanceAttrTypes, m)
	diags.Append(d...)
	return t
}

// Flatten populates the Terraform model from SDK type
func (m *FixedAddressInheritanceModel) Flatten(ctx context.Context, from *uddidhcp.FixedAddressInheritance, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.DhcpOptions = FlattenInheritedDHCPOptionList(ctx, from.DhcpOptions, diags)
	m.HeaderOptionFilename = FlattenInheritanceInheritedString(ctx, from.HeaderOptionFilename, diags)
	m.HeaderOptionServerAddress = FlattenInheritanceInheritedString(ctx, from.HeaderOptionServerAddress, diags)
	m.HeaderOptionServerName = FlattenInheritanceInheritedString(ctx, from.HeaderOptionServerName, diags)
}
