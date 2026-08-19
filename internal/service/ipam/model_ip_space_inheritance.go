package ipam

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	uddiipam "github.com/infobloxopen/universal-ddi-go-client/ipam"
)

// IPSpaceInheritanceModel is the Terraform model for IPSpaceInheritance
type IPSpaceInheritanceModel struct {
	AsmConfig                       types.Object `tfsdk:"asm_config"`
	DdnsClientUpdate                types.Object `tfsdk:"ddns_client_update"`
	DdnsConflictResolutionMode      types.Object `tfsdk:"ddns_conflict_resolution_mode"`
	DdnsEnabled                     types.Object `tfsdk:"ddns_enabled"`
	DdnsHostnameBlock               types.Object `tfsdk:"ddns_hostname_block"`
	DdnsTtlPercent                  types.Object `tfsdk:"ddns_ttl_percent"`
	DdnsUpdateBlock                 types.Object `tfsdk:"ddns_update_block"`
	DdnsUpdateOnRenew               types.Object `tfsdk:"ddns_update_on_renew"`
	DdnsUseConflictResolution       types.Object `tfsdk:"ddns_use_conflict_resolution"`
	DhcpConfig                      types.Object `tfsdk:"dhcp_config"`
	DhcpOptions                     types.Object `tfsdk:"dhcp_options"`
	DhcpOptionsV6                   types.Object `tfsdk:"dhcp_options_v6"`
	HeaderOptionFilename            types.Object `tfsdk:"header_option_filename"`
	HeaderOptionServerAddress       types.Object `tfsdk:"header_option_server_address"`
	HeaderOptionServerName          types.Object `tfsdk:"header_option_server_name"`
	HostnameRewriteBlock            types.Object `tfsdk:"hostname_rewrite_block"`
	VendorSpecificOptionOptionSpace types.Object `tfsdk:"vendor_specific_option_option_space"`
}

// IPSpaceInheritanceAttrTypes contains the attribute types for IPSpaceInheritanceModel
var IPSpaceInheritanceAttrTypes = map[string]attr.Type{
	"asm_config":                          types.ObjectType{AttrTypes: InheritedASMConfigAttrTypes},
	"ddns_client_update":                  types.ObjectType{AttrTypes: InheritanceInheritedStringAttrTypes},
	"ddns_conflict_resolution_mode":       types.ObjectType{AttrTypes: InheritanceInheritedStringAttrTypes},
	"ddns_enabled":                        types.ObjectType{AttrTypes: InheritanceInheritedBoolAttrTypes},
	"ddns_hostname_block":                 types.ObjectType{AttrTypes: InheritedDDNSHostnameBlockAttrTypes},
	"ddns_ttl_percent":                    types.ObjectType{AttrTypes: InheritanceInheritedFloatAttrTypes},
	"ddns_update_block":                   types.ObjectType{AttrTypes: InheritedDDNSUpdateBlockAttrTypes},
	"ddns_update_on_renew":                types.ObjectType{AttrTypes: InheritanceInheritedBoolAttrTypes},
	"ddns_use_conflict_resolution":        types.ObjectType{AttrTypes: InheritanceInheritedBoolAttrTypes},
	"dhcp_config":                         types.ObjectType{AttrTypes: InheritedDHCPConfigAttrTypes},
	"dhcp_options":                        types.ObjectType{AttrTypes: InheritedDHCPOptionListAttrTypes},
	"dhcp_options_v6":                     types.ObjectType{AttrTypes: InheritedDHCPOptionListAttrTypes},
	"header_option_filename":              types.ObjectType{AttrTypes: InheritanceInheritedStringAttrTypes},
	"header_option_server_address":        types.ObjectType{AttrTypes: InheritanceInheritedStringAttrTypes},
	"header_option_server_name":           types.ObjectType{AttrTypes: InheritanceInheritedStringAttrTypes},
	"hostname_rewrite_block":              types.ObjectType{AttrTypes: InheritedHostnameRewriteBlockAttrTypes},
	"vendor_specific_option_option_space": types.ObjectType{AttrTypes: InheritanceInheritedIdentifierAttrTypes},
}

// IPSpaceInheritanceResourceSchemaAttributes contains the schema attributes for IPSpaceInheritanceModel
var IPSpaceInheritanceResourceSchemaAttributes = map[string]schema.Attribute{
	"asm_config": schema.SingleNestedAttribute{
		Attributes:          InheritedASMConfigResourceSchemaAttributes,
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The inheritance configuration for _asm_config_ field.",
	},
	"ddns_client_update": schema.SingleNestedAttribute{
		Attributes:          InheritanceInheritedStringResourceSchemaAttributes,
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The inheritance configuration for _ddns_client_update_ field from _IPSpace_ object.",
	},
	"ddns_conflict_resolution_mode": schema.SingleNestedAttribute{
		Attributes:          InheritanceInheritedStringResourceSchemaAttributes,
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The inheritance configuration for _ddns_conflict_resolution_mode_ field from _IPSpace_ object.",
	},
	"ddns_enabled": schema.SingleNestedAttribute{
		Attributes:          InheritanceInheritedBoolResourceSchemaAttributes,
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The inheritance configuration for _ddns_enabled_ field. Only action allowed is 'inherit'.",
	},
	"ddns_hostname_block": schema.SingleNestedAttribute{
		Attributes:          InheritedDDNSHostnameBlockResourceSchemaAttributes,
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The inheritance configuration for _ddns_generate_name_ and _ddns_generated_prefix_ fields from _IPSpace_ object.",
	},
	"ddns_ttl_percent": schema.SingleNestedAttribute{
		Attributes:          InheritanceInheritedFloatResourceSchemaAttributes,
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The inheritance configuration for _ddns_ttl_percent_ field from _IPSpace_ object.",
	},
	"ddns_update_block": schema.SingleNestedAttribute{
		Attributes:          InheritedDDNSUpdateBlockResourceSchemaAttributes,
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The inheritance configuration for _ddns_send_updates_ and _ddns_domain_ fields from _IPSpace_ object.",
	},
	"ddns_update_on_renew": schema.SingleNestedAttribute{
		Attributes:          InheritanceInheritedBoolResourceSchemaAttributes,
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The inheritance configuration for _ddns_update_on_renew_ field from _IPSpace_ object.",
	},
	"ddns_use_conflict_resolution": schema.SingleNestedAttribute{
		Attributes:          InheritanceInheritedBoolResourceSchemaAttributes,
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The inheritance configuration for _ddns_use_conflict_resolution_ field from _IPSpace_ object.",
	},
	"dhcp_config": schema.SingleNestedAttribute{
		Attributes:          InheritedDHCPConfigResourceSchemaAttributes,
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The inheritance configuration for _dhcp_config_ field.",
	},
	"dhcp_options": schema.SingleNestedAttribute{
		Attributes:          InheritedDHCPOptionListResourceSchemaAttributes,
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The inheritance configuration for _dhcp_options_ field.",
	},
	"dhcp_options_v6": schema.SingleNestedAttribute{
		Attributes:          InheritedDHCPOptionListResourceSchemaAttributes,
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The inheritance configuration for _dhcp_options_v6_ field.",
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
	"hostname_rewrite_block": schema.SingleNestedAttribute{
		Attributes:          InheritedHostnameRewriteBlockResourceSchemaAttributes,
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The inheritance configuration for _hostname_rewrite_enabled_, _hostname_rewrite_regex_, and _hostname_rewrite_char_ fields from _IPSpace_ object.",
	},
	"vendor_specific_option_option_space": schema.SingleNestedAttribute{
		Attributes:          InheritanceInheritedIdentifierResourceSchemaAttributes,
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The inheritance configuration for _vendor_specific_option_option_space_ field.",
	},
}

// ExpandIPSpaceInheritance converts a Terraform Object to SDK type
func ExpandIPSpaceInheritance(ctx context.Context, o types.Object, diags *diag.Diagnostics) *uddiipam.IPSpaceInheritance {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m IPSpaceInheritanceModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

// Expand converts the Terraform model to SDK type
func (m *IPSpaceInheritanceModel) Expand(ctx context.Context, diags *diag.Diagnostics) *uddiipam.IPSpaceInheritance {
	if m == nil {
		return nil
	}
	to := &uddiipam.IPSpaceInheritance{
		AsmConfig:                       ExpandInheritedASMConfig(ctx, m.AsmConfig, diags),
		DdnsClientUpdate:                ExpandInheritanceInheritedString(ctx, m.DdnsClientUpdate, diags),
		DdnsConflictResolutionMode:      ExpandInheritanceInheritedString(ctx, m.DdnsConflictResolutionMode, diags),
		DdnsEnabled:                     ExpandInheritanceInheritedBool(ctx, m.DdnsEnabled, diags),
		DdnsHostnameBlock:               ExpandInheritedDDNSHostnameBlock(ctx, m.DdnsHostnameBlock, diags),
		DdnsTtlPercent:                  ExpandInheritanceInheritedFloat(ctx, m.DdnsTtlPercent, diags),
		DdnsUpdateBlock:                 ExpandInheritedDDNSUpdateBlock(ctx, m.DdnsUpdateBlock, diags),
		DdnsUpdateOnRenew:               ExpandInheritanceInheritedBool(ctx, m.DdnsUpdateOnRenew, diags),
		DdnsUseConflictResolution:       ExpandInheritanceInheritedBool(ctx, m.DdnsUseConflictResolution, diags),
		DhcpConfig:                      ExpandInheritedDHCPConfig(ctx, m.DhcpConfig, diags),
		DhcpOptions:                     ExpandInheritedDHCPOptionList(ctx, m.DhcpOptions, diags),
		DhcpOptionsV6:                   ExpandInheritedDHCPOptionList(ctx, m.DhcpOptionsV6, diags),
		HeaderOptionFilename:            ExpandInheritanceInheritedString(ctx, m.HeaderOptionFilename, diags),
		HeaderOptionServerAddress:       ExpandInheritanceInheritedString(ctx, m.HeaderOptionServerAddress, diags),
		HeaderOptionServerName:          ExpandInheritanceInheritedString(ctx, m.HeaderOptionServerName, diags),
		HostnameRewriteBlock:            ExpandInheritedHostnameRewriteBlock(ctx, m.HostnameRewriteBlock, diags),
		VendorSpecificOptionOptionSpace: ExpandInheritanceInheritedIdentifier(ctx, m.VendorSpecificOptionOptionSpace, diags),
	}
	return to
}

// FlattenIPSpaceInheritance converts an SDK type to Terraform Object
func FlattenIPSpaceInheritance(ctx context.Context, from *uddiipam.IPSpaceInheritance, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(IPSpaceInheritanceAttrTypes)
	}
	m := &IPSpaceInheritanceModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, IPSpaceInheritanceAttrTypes, m)
	diags.Append(d...)
	return t
}

// Flatten populates the Terraform model from SDK type
func (m *IPSpaceInheritanceModel) Flatten(ctx context.Context, from *uddiipam.IPSpaceInheritance, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.AsmConfig = FlattenInheritedASMConfig(ctx, from.AsmConfig, diags)
	m.DdnsClientUpdate = FlattenInheritanceInheritedString(ctx, from.DdnsClientUpdate, diags)
	m.DdnsConflictResolutionMode = FlattenInheritanceInheritedString(ctx, from.DdnsConflictResolutionMode, diags)
	m.DdnsEnabled = FlattenInheritanceInheritedBool(ctx, from.DdnsEnabled, diags)
	m.DdnsHostnameBlock = FlattenInheritedDDNSHostnameBlock(ctx, from.DdnsHostnameBlock, diags)
	m.DdnsTtlPercent = FlattenInheritanceInheritedFloat(ctx, from.DdnsTtlPercent, diags)
	m.DdnsUpdateBlock = FlattenInheritedDDNSUpdateBlock(ctx, from.DdnsUpdateBlock, diags)
	m.DdnsUpdateOnRenew = FlattenInheritanceInheritedBool(ctx, from.DdnsUpdateOnRenew, diags)
	m.DdnsUseConflictResolution = FlattenInheritanceInheritedBool(ctx, from.DdnsUseConflictResolution, diags)
	m.DhcpConfig = FlattenInheritedDHCPConfig(ctx, from.DhcpConfig, diags)
	m.DhcpOptions = FlattenInheritedDHCPOptionList(ctx, from.DhcpOptions, diags)
	m.DhcpOptionsV6 = FlattenInheritedDHCPOptionList(ctx, from.DhcpOptionsV6, diags)
	m.HeaderOptionFilename = FlattenInheritanceInheritedString(ctx, from.HeaderOptionFilename, diags)
	m.HeaderOptionServerAddress = FlattenInheritanceInheritedString(ctx, from.HeaderOptionServerAddress, diags)
	m.HeaderOptionServerName = FlattenInheritanceInheritedString(ctx, from.HeaderOptionServerName, diags)
	m.HostnameRewriteBlock = FlattenInheritedHostnameRewriteBlock(ctx, from.HostnameRewriteBlock, diags)
	m.VendorSpecificOptionOptionSpace = FlattenInheritanceInheritedIdentifier(ctx, from.VendorSpecificOptionOptionSpace, diags)
}
