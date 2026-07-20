package dns

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	uddidns "github.com/infobloxopen/universal-ddi-go-client/dnsconfig"
)

// ViewInheritanceModel is the Terraform model for ViewInheritance
type ViewInheritanceModel struct {
	AddEdnsOptionInOutgoingQuery      types.Object `tfsdk:"add_edns_option_in_outgoing_query"`
	CustomRootNsBlock                 types.Object `tfsdk:"custom_root_ns_block"`
	DnssecValidationBlock             types.Object `tfsdk:"dnssec_validation_block"`
	DtcConfig                         types.Object `tfsdk:"dtc_config"`
	EcsBlock                          types.Object `tfsdk:"ecs_block"`
	EdnsUdpSize                       types.Object `tfsdk:"edns_udp_size"`
	FilterAaaaAcl                     types.Object `tfsdk:"filter_aaaa_acl"`
	FilterAaaaOnV4                    types.Object `tfsdk:"filter_aaaa_on_v4"`
	ForwardersBlock                   types.Object `tfsdk:"forwarders_block"`
	GssTsigEnabled                    types.Object `tfsdk:"gss_tsig_enabled"`
	LameTtl                           types.Object `tfsdk:"lame_ttl"`
	MatchRecursiveOnly                types.Object `tfsdk:"match_recursive_only"`
	MaxCacheTtl                       types.Object `tfsdk:"max_cache_ttl"`
	MaxNegativeTtl                    types.Object `tfsdk:"max_negative_ttl"`
	MaxUdpSize                        types.Object `tfsdk:"max_udp_size"`
	MinimalResponses                  types.Object `tfsdk:"minimal_responses"`
	Notify                            types.Object `tfsdk:"notify"`
	QueryAcl                          types.Object `tfsdk:"query_acl"`
	RecursionAcl                      types.Object `tfsdk:"recursion_acl"`
	RecursionEnabled                  types.Object `tfsdk:"recursion_enabled"`
	SortList                          types.Object `tfsdk:"sort_list"`
	SynthesizeAddressRecordsFromHttps types.Object `tfsdk:"synthesize_address_records_from_https"`
	TransferAcl                       types.Object `tfsdk:"transfer_acl"`
	UpdateAcl                         types.Object `tfsdk:"update_acl"`
	UseForwardersForSubzones          types.Object `tfsdk:"use_forwarders_for_subzones"`
	ZoneAuthority                     types.Object `tfsdk:"zone_authority"`
}

// ViewInheritanceAttrTypes contains the attribute types for ViewInheritanceModel
var ViewInheritanceAttrTypes = map[string]attr.Type{
	"add_edns_option_in_outgoing_query":     types.ObjectType{AttrTypes: Inheritance2InheritedBoolAttrTypes},
	"custom_root_ns_block":                  types.ObjectType{AttrTypes: InheritedCustomRootNSBlockAttrTypes},
	"dnssec_validation_block":               types.ObjectType{AttrTypes: InheritedDNSSECValidationBlockAttrTypes},
	"dtc_config":                            types.ObjectType{AttrTypes: InheritedDtcConfigAttrTypes},
	"ecs_block":                             types.ObjectType{AttrTypes: InheritedECSBlockAttrTypes},
	"edns_udp_size":                         types.ObjectType{AttrTypes: Inheritance2InheritedUInt32DnsconfigAttrTypes},
	"filter_aaaa_acl":                       types.ObjectType{AttrTypes: InheritedACLItemsAttrTypes},
	"filter_aaaa_on_v4":                     types.ObjectType{AttrTypes: Inheritance2InheritedStringAttrTypes},
	"forwarders_block":                      types.ObjectType{AttrTypes: InheritedForwardersBlockAttrTypes},
	"gss_tsig_enabled":                      types.ObjectType{AttrTypes: Inheritance2InheritedBoolAttrTypes},
	"lame_ttl":                              types.ObjectType{AttrTypes: Inheritance2InheritedUInt32DnsconfigAttrTypes},
	"match_recursive_only":                  types.ObjectType{AttrTypes: Inheritance2InheritedBoolAttrTypes},
	"max_cache_ttl":                         types.ObjectType{AttrTypes: Inheritance2InheritedUInt32DnsconfigAttrTypes},
	"max_negative_ttl":                      types.ObjectType{AttrTypes: Inheritance2InheritedUInt32DnsconfigAttrTypes},
	"max_udp_size":                          types.ObjectType{AttrTypes: Inheritance2InheritedUInt32DnsconfigAttrTypes},
	"minimal_responses":                     types.ObjectType{AttrTypes: Inheritance2InheritedBoolAttrTypes},
	"notify":                                types.ObjectType{AttrTypes: Inheritance2InheritedBoolAttrTypes},
	"query_acl":                             types.ObjectType{AttrTypes: InheritedACLItemsAttrTypes},
	"recursion_acl":                         types.ObjectType{AttrTypes: InheritedACLItemsAttrTypes},
	"recursion_enabled":                     types.ObjectType{AttrTypes: Inheritance2InheritedBoolAttrTypes},
	"sort_list":                             types.ObjectType{AttrTypes: InheritedSortListItemsAttrTypes},
	"synthesize_address_records_from_https": types.ObjectType{AttrTypes: Inheritance2InheritedBoolAttrTypes},
	"transfer_acl":                          types.ObjectType{AttrTypes: InheritedACLItemsAttrTypes},
	"update_acl":                            types.ObjectType{AttrTypes: InheritedACLItemsAttrTypes},
	"use_forwarders_for_subzones":           types.ObjectType{AttrTypes: Inheritance2InheritedBoolAttrTypes},
	"zone_authority":                        types.ObjectType{AttrTypes: InheritedZoneAuthorityAttrTypes},
}

// ViewInheritanceResourceSchemaAttributes contains the schema attributes for ViewInheritanceModel
var ViewInheritanceResourceSchemaAttributes = map[string]schema.Attribute{
	"add_edns_option_in_outgoing_query": schema.SingleNestedAttribute{
		Attributes:          Inheritance2InheritedBoolResourceSchemaAttributes,
		Optional:            true,
		MarkdownDescription: "Field config for _add_edns_option_in_outgoing_query_ field from _View_ object.",
	},
	"custom_root_ns_block": schema.SingleNestedAttribute{
		Attributes:          InheritedCustomRootNSBlockResourceSchemaAttributes,
		Optional:            true,
		MarkdownDescription: "Optional. Field config for _custom_root_ns_block_ field from _View_ object.",
	},
	"dnssec_validation_block": schema.SingleNestedAttribute{
		Attributes:          InheritedDNSSECValidationBlockResourceSchemaAttributes,
		Optional:            true,
		MarkdownDescription: "Optional. Field config for _dnssec_validation_block_ field from _View_ object.",
	},
	"dtc_config": schema.SingleNestedAttribute{
		Attributes:          InheritedDtcConfigResourceSchemaAttributes,
		Optional:            true,
		MarkdownDescription: "Optional. Field config for _dtc_config_ field from _View_ object.",
	},
	"ecs_block": schema.SingleNestedAttribute{
		Attributes:          InheritedECSBlockResourceSchemaAttributes,
		Optional:            true,
		MarkdownDescription: "Optional. Field config for _ecs_block_ field from _View_ object.",
	},
	"edns_udp_size": schema.SingleNestedAttribute{
		Attributes:          Inheritance2InheritedUInt32DnsconfigResourceSchemaAttributes,
		Optional:            true,
		MarkdownDescription: "Optional. Field config for _edns_udp_size_ field from [View] object.",
	},
	"filter_aaaa_acl": schema.SingleNestedAttribute{
		Attributes:          InheritedACLItemsResourceSchemaAttributes,
		Optional:            true,
		MarkdownDescription: "Optional. Field config for _filter_aaaa_acl_ field from _View_ object.",
	},
	"filter_aaaa_on_v4": schema.SingleNestedAttribute{
		Attributes:          Inheritance2InheritedStringResourceSchemaAttributes,
		Optional:            true,
		MarkdownDescription: "Optional. Field config for _filter_aaaa_on_v4_ field from _View_ object.",
	},
	"forwarders_block": schema.SingleNestedAttribute{
		Attributes:          InheritedForwardersBlockResourceSchemaAttributes,
		Optional:            true,
		MarkdownDescription: "Optional. Field config for _forwarders_block_ field from _View_ object.",
	},
	"gss_tsig_enabled": schema.SingleNestedAttribute{
		Attributes:          Inheritance2InheritedBoolResourceSchemaAttributes,
		Optional:            true,
		MarkdownDescription: "Optional. Field config for _gss_tsig_enabled_ field from _View_ object.",
	},
	"lame_ttl": schema.SingleNestedAttribute{
		Attributes:          Inheritance2InheritedUInt32DnsconfigResourceSchemaAttributes,
		Optional:            true,
		MarkdownDescription: "Optional. Field config for _lame_ttl_ field from _View_ object.",
	},
	"match_recursive_only": schema.SingleNestedAttribute{
		Attributes:          Inheritance2InheritedBoolResourceSchemaAttributes,
		Optional:            true,
		MarkdownDescription: "Optional. Field config for _match_recursive_only_ field from _View_ object.",
	},
	"max_cache_ttl": schema.SingleNestedAttribute{
		Attributes:          Inheritance2InheritedUInt32DnsconfigResourceSchemaAttributes,
		Optional:            true,
		MarkdownDescription: "Optional. Field config for _max_cache_ttl_ field from _View_ object.",
	},
	"max_negative_ttl": schema.SingleNestedAttribute{
		Attributes:          Inheritance2InheritedUInt32DnsconfigResourceSchemaAttributes,
		Optional:            true,
		MarkdownDescription: "Optional. Field config for _max_negative_ttl_ field from _View_ object.",
	},
	"max_udp_size": schema.SingleNestedAttribute{
		Attributes:          Inheritance2InheritedUInt32DnsconfigResourceSchemaAttributes,
		Optional:            true,
		MarkdownDescription: "Optional. Field config for _max_udp_size_ field from [View] object.",
	},
	"minimal_responses": schema.SingleNestedAttribute{
		Attributes:          Inheritance2InheritedBoolResourceSchemaAttributes,
		Optional:            true,
		MarkdownDescription: "Optional. Field config for _minimal_responses_ field from _View_ object.",
	},
	"notify": schema.SingleNestedAttribute{
		Attributes:          Inheritance2InheritedBoolResourceSchemaAttributes,
		Optional:            true,
		MarkdownDescription: "Field config for _notify_ field from _View_ object.",
	},
	"query_acl": schema.SingleNestedAttribute{
		Attributes:          InheritedACLItemsResourceSchemaAttributes,
		Optional:            true,
		MarkdownDescription: "Optional. Field config for _query_acl_ field from _View_ object.",
	},
	"recursion_acl": schema.SingleNestedAttribute{
		Attributes:          InheritedACLItemsResourceSchemaAttributes,
		Optional:            true,
		MarkdownDescription: "Optional. Field config for _recursion_acl_ field from _View_ object.",
	},
	"recursion_enabled": schema.SingleNestedAttribute{
		Attributes:          Inheritance2InheritedBoolResourceSchemaAttributes,
		Optional:            true,
		MarkdownDescription: "Optional. Field config for _recursion_enabled_ field from _View_ object.",
	},
	"sort_list": schema.SingleNestedAttribute{
		Attributes:          InheritedSortListItemsResourceSchemaAttributes,
		Optional:            true,
		MarkdownDescription: "Optional. Field config for _sort_list_ field from _View_ object.",
	},
	"synthesize_address_records_from_https": schema.SingleNestedAttribute{
		Attributes:          Inheritance2InheritedBoolResourceSchemaAttributes,
		Optional:            true,
		MarkdownDescription: "Field config for _synthesize_address_records_from_https_ field from _View_ object.",
	},
	"transfer_acl": schema.SingleNestedAttribute{
		Attributes:          InheritedACLItemsResourceSchemaAttributes,
		Optional:            true,
		MarkdownDescription: "Optional. Field config for _transfer_acl_ field from _View_ object.",
	},
	"update_acl": schema.SingleNestedAttribute{
		Attributes:          InheritedACLItemsResourceSchemaAttributes,
		Optional:            true,
		MarkdownDescription: "Optional. Field config for _update_acl_ field from _View_ object.",
	},
	"use_forwarders_for_subzones": schema.SingleNestedAttribute{
		Attributes:          Inheritance2InheritedBoolResourceSchemaAttributes,
		Optional:            true,
		MarkdownDescription: "Optional. Field config for _use_forwarders_for_subzones_ field from _View_ object.",
	},
	"zone_authority": schema.SingleNestedAttribute{
		Attributes:          InheritedZoneAuthorityResourceSchemaAttributes,
		Optional:            true,
		MarkdownDescription: "Optional. Field config for _zone_authority_ field from _View_ object.",
	},
}

// ExpandViewInheritance converts a Terraform Object to SDK type
func ExpandViewInheritance(ctx context.Context, o types.Object, diags *diag.Diagnostics) *uddidns.ViewInheritance {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m ViewInheritanceModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

// Expand converts the Terraform model to SDK type
func (m *ViewInheritanceModel) Expand(ctx context.Context, diags *diag.Diagnostics) *uddidns.ViewInheritance {
	if m == nil {
		return nil
	}
	to := &uddidns.ViewInheritance{
		AddEdnsOptionInOutgoingQuery:      ExpandInheritance2InheritedBool(ctx, m.AddEdnsOptionInOutgoingQuery, diags),
		CustomRootNsBlock:                 ExpandInheritedCustomRootNSBlock(ctx, m.CustomRootNsBlock, diags),
		DnssecValidationBlock:             ExpandInheritedDNSSECValidationBlock(ctx, m.DnssecValidationBlock, diags),
		DtcConfig:                         ExpandInheritedDtcConfig(ctx, m.DtcConfig, diags),
		EcsBlock:                          ExpandInheritedECSBlock(ctx, m.EcsBlock, diags),
		EdnsUdpSize:                       ExpandInheritance2InheritedUInt32Dnsconfig(ctx, m.EdnsUdpSize, diags),
		FilterAaaaAcl:                     ExpandInheritedACLItems(ctx, m.FilterAaaaAcl, diags),
		FilterAaaaOnV4:                    ExpandInheritance2InheritedString(ctx, m.FilterAaaaOnV4, diags),
		ForwardersBlock:                   ExpandInheritedForwardersBlock(ctx, m.ForwardersBlock, diags),
		GssTsigEnabled:                    ExpandInheritance2InheritedBool(ctx, m.GssTsigEnabled, diags),
		LameTtl:                           ExpandInheritance2InheritedUInt32Dnsconfig(ctx, m.LameTtl, diags),
		MatchRecursiveOnly:                ExpandInheritance2InheritedBool(ctx, m.MatchRecursiveOnly, diags),
		MaxCacheTtl:                       ExpandInheritance2InheritedUInt32Dnsconfig(ctx, m.MaxCacheTtl, diags),
		MaxNegativeTtl:                    ExpandInheritance2InheritedUInt32Dnsconfig(ctx, m.MaxNegativeTtl, diags),
		MaxUdpSize:                        ExpandInheritance2InheritedUInt32Dnsconfig(ctx, m.MaxUdpSize, diags),
		MinimalResponses:                  ExpandInheritance2InheritedBool(ctx, m.MinimalResponses, diags),
		Notify:                            ExpandInheritance2InheritedBool(ctx, m.Notify, diags),
		QueryAcl:                          ExpandInheritedACLItems(ctx, m.QueryAcl, diags),
		RecursionAcl:                      ExpandInheritedACLItems(ctx, m.RecursionAcl, diags),
		RecursionEnabled:                  ExpandInheritance2InheritedBool(ctx, m.RecursionEnabled, diags),
		SortList:                          ExpandInheritedSortListItems(ctx, m.SortList, diags),
		SynthesizeAddressRecordsFromHttps: ExpandInheritance2InheritedBool(ctx, m.SynthesizeAddressRecordsFromHttps, diags),
		TransferAcl:                       ExpandInheritedACLItems(ctx, m.TransferAcl, diags),
		UpdateAcl:                         ExpandInheritedACLItems(ctx, m.UpdateAcl, diags),
		UseForwardersForSubzones:          ExpandInheritance2InheritedBool(ctx, m.UseForwardersForSubzones, diags),
		ZoneAuthority:                     ExpandInheritedZoneAuthority(ctx, m.ZoneAuthority, diags),
	}
	return to
}

// FlattenViewInheritance converts an SDK type to Terraform Object
func FlattenViewInheritance(ctx context.Context, from *uddidns.ViewInheritance, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(ViewInheritanceAttrTypes)
	}
	m := &ViewInheritanceModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, ViewInheritanceAttrTypes, m)
	diags.Append(d...)
	return t
}

// Flatten populates the Terraform model from SDK type
func (m *ViewInheritanceModel) Flatten(ctx context.Context, from *uddidns.ViewInheritance, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.AddEdnsOptionInOutgoingQuery = FlattenInheritance2InheritedBool(ctx, from.AddEdnsOptionInOutgoingQuery, diags)
	m.CustomRootNsBlock = FlattenInheritedCustomRootNSBlock(ctx, from.CustomRootNsBlock, diags)
	m.DnssecValidationBlock = FlattenInheritedDNSSECValidationBlock(ctx, from.DnssecValidationBlock, diags)
	m.DtcConfig = FlattenInheritedDtcConfig(ctx, from.DtcConfig, diags)
	m.EcsBlock = FlattenInheritedECSBlock(ctx, from.EcsBlock, diags)
	m.EdnsUdpSize = FlattenInheritance2InheritedUInt32Dnsconfig(ctx, from.EdnsUdpSize, diags)
	m.FilterAaaaAcl = FlattenInheritedACLItems(ctx, from.FilterAaaaAcl, diags)
	m.FilterAaaaOnV4 = FlattenInheritance2InheritedString(ctx, from.FilterAaaaOnV4, diags)
	m.ForwardersBlock = FlattenInheritedForwardersBlock(ctx, from.ForwardersBlock, diags)
	m.GssTsigEnabled = FlattenInheritance2InheritedBool(ctx, from.GssTsigEnabled, diags)
	m.LameTtl = FlattenInheritance2InheritedUInt32Dnsconfig(ctx, from.LameTtl, diags)
	m.MatchRecursiveOnly = FlattenInheritance2InheritedBool(ctx, from.MatchRecursiveOnly, diags)
	m.MaxCacheTtl = FlattenInheritance2InheritedUInt32Dnsconfig(ctx, from.MaxCacheTtl, diags)
	m.MaxNegativeTtl = FlattenInheritance2InheritedUInt32Dnsconfig(ctx, from.MaxNegativeTtl, diags)
	m.MaxUdpSize = FlattenInheritance2InheritedUInt32Dnsconfig(ctx, from.MaxUdpSize, diags)
	m.MinimalResponses = FlattenInheritance2InheritedBool(ctx, from.MinimalResponses, diags)
	m.Notify = FlattenInheritance2InheritedBool(ctx, from.Notify, diags)
	m.QueryAcl = FlattenInheritedACLItems(ctx, from.QueryAcl, diags)
	m.RecursionAcl = FlattenInheritedACLItems(ctx, from.RecursionAcl, diags)
	m.RecursionEnabled = FlattenInheritance2InheritedBool(ctx, from.RecursionEnabled, diags)
	m.SortList = FlattenInheritedSortListItems(ctx, from.SortList, diags)
	m.SynthesizeAddressRecordsFromHttps = FlattenInheritance2InheritedBool(ctx, from.SynthesizeAddressRecordsFromHttps, diags)
	m.TransferAcl = FlattenInheritedACLItems(ctx, from.TransferAcl, diags)
	m.UpdateAcl = FlattenInheritedACLItems(ctx, from.UpdateAcl, diags)
	m.UseForwardersForSubzones = FlattenInheritance2InheritedBool(ctx, from.UseForwardersForSubzones, diags)
	m.ZoneAuthority = FlattenInheritedZoneAuthority(ctx, from.ZoneAuthority, diags)
}
