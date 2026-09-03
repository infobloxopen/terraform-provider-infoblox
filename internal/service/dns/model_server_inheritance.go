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

// ServerInheritanceModel is the Terraform model for ServerInheritance
type ServerInheritanceModel struct {
	AddEdnsOptionInOutgoingQuery      types.Object `tfsdk:"add_edns_option_in_outgoing_query"`
	CustomRootNsBlock                 types.Object `tfsdk:"custom_root_ns_block"`
	DnssecValidationBlock             types.Object `tfsdk:"dnssec_validation_block"`
	EcsBlock                          types.Object `tfsdk:"ecs_block"`
	FilterAaaaAcl                     types.Object `tfsdk:"filter_aaaa_acl"`
	FilterAaaaOnV4                    types.Object `tfsdk:"filter_aaaa_on_v4"`
	ForwardersBlock                   types.Object `tfsdk:"forwarders_block"`
	GssTsigEnabled                    types.Object `tfsdk:"gss_tsig_enabled"`
	KerberosKeys                      types.Object `tfsdk:"kerberos_keys"`
	LameTtl                           types.Object `tfsdk:"lame_ttl"`
	LogQueryResponse                  types.Object `tfsdk:"log_query_response"`
	MatchRecursiveOnly                types.Object `tfsdk:"match_recursive_only"`
	MaxCacheTtl                       types.Object `tfsdk:"max_cache_ttl"`
	MaxNegativeTtl                    types.Object `tfsdk:"max_negative_ttl"`
	MinimalResponses                  types.Object `tfsdk:"minimal_responses"`
	Notify                            types.Object `tfsdk:"notify"`
	QueryAcl                          types.Object `tfsdk:"query_acl"`
	QueryPort                         types.Object `tfsdk:"query_port"`
	RecursionAcl                      types.Object `tfsdk:"recursion_acl"`
	RecursionEnabled                  types.Object `tfsdk:"recursion_enabled"`
	RecursiveClients                  types.Object `tfsdk:"recursive_clients"`
	ResolverQueryTimeout              types.Object `tfsdk:"resolver_query_timeout"`
	SecondaryAxfrQueryLimit           types.Object `tfsdk:"secondary_axfr_query_limit"`
	SecondarySoaQueryLimit            types.Object `tfsdk:"secondary_soa_query_limit"`
	SortList                          types.Object `tfsdk:"sort_list"`
	SynthesizeAddressRecordsFromHttps types.Object `tfsdk:"synthesize_address_records_from_https"`
	TransferAcl                       types.Object `tfsdk:"transfer_acl"`
	UpdateAcl                         types.Object `tfsdk:"update_acl"`
	UseForwardersForSubzones          types.Object `tfsdk:"use_forwarders_for_subzones"`
}

// ServerInheritanceAttrTypes contains the attribute types for ServerInheritanceModel
var ServerInheritanceAttrTypes = map[string]attr.Type{
	"add_edns_option_in_outgoing_query":     types.ObjectType{AttrTypes: Inheritance2InheritedBoolAttrTypes},
	"custom_root_ns_block":                  types.ObjectType{AttrTypes: InheritedCustomRootNSBlockAttrTypes},
	"dnssec_validation_block":               types.ObjectType{AttrTypes: InheritedDNSSECValidationBlockAttrTypes},
	"ecs_block":                             types.ObjectType{AttrTypes: InheritedECSBlockAttrTypes},
	"filter_aaaa_acl":                       types.ObjectType{AttrTypes: InheritedACLItemsAttrTypes},
	"filter_aaaa_on_v4":                     types.ObjectType{AttrTypes: Inheritance2InheritedStringAttrTypes},
	"forwarders_block":                      types.ObjectType{AttrTypes: InheritedForwardersBlockAttrTypes},
	"gss_tsig_enabled":                      types.ObjectType{AttrTypes: Inheritance2InheritedBoolAttrTypes},
	"kerberos_keys":                         types.ObjectType{AttrTypes: InheritedKerberosKeysAttrTypes},
	"lame_ttl":                              types.ObjectType{AttrTypes: Inheritance2InheritedUInt32DnsconfigAttrTypes},
	"log_query_response":                    types.ObjectType{AttrTypes: Inheritance2InheritedBoolAttrTypes},
	"match_recursive_only":                  types.ObjectType{AttrTypes: Inheritance2InheritedBoolAttrTypes},
	"max_cache_ttl":                         types.ObjectType{AttrTypes: Inheritance2InheritedUInt32DnsconfigAttrTypes},
	"max_negative_ttl":                      types.ObjectType{AttrTypes: Inheritance2InheritedUInt32DnsconfigAttrTypes},
	"minimal_responses":                     types.ObjectType{AttrTypes: Inheritance2InheritedBoolAttrTypes},
	"notify":                                types.ObjectType{AttrTypes: Inheritance2InheritedBoolAttrTypes},
	"query_acl":                             types.ObjectType{AttrTypes: InheritedACLItemsAttrTypes},
	"query_port":                            types.ObjectType{AttrTypes: Inheritance2InheritedUInt32DnsconfigAttrTypes},
	"recursion_acl":                         types.ObjectType{AttrTypes: InheritedACLItemsAttrTypes},
	"recursion_enabled":                     types.ObjectType{AttrTypes: Inheritance2InheritedBoolAttrTypes},
	"recursive_clients":                     types.ObjectType{AttrTypes: Inheritance2InheritedUInt32DnsconfigAttrTypes},
	"resolver_query_timeout":                types.ObjectType{AttrTypes: Inheritance2InheritedUInt32DnsconfigAttrTypes},
	"secondary_axfr_query_limit":            types.ObjectType{AttrTypes: Inheritance2InheritedUInt32DnsconfigAttrTypes},
	"secondary_soa_query_limit":             types.ObjectType{AttrTypes: Inheritance2InheritedUInt32DnsconfigAttrTypes},
	"sort_list":                             types.ObjectType{AttrTypes: InheritedSortListItemsAttrTypes},
	"synthesize_address_records_from_https": types.ObjectType{AttrTypes: Inheritance2InheritedBoolAttrTypes},
	"transfer_acl":                          types.ObjectType{AttrTypes: InheritedACLItemsAttrTypes},
	"update_acl":                            types.ObjectType{AttrTypes: InheritedACLItemsAttrTypes},
	"use_forwarders_for_subzones":           types.ObjectType{AttrTypes: Inheritance2InheritedBoolAttrTypes},
}

// ServerInheritanceResourceSchemaAttributes contains the schema attributes for ServerInheritanceModel
var ServerInheritanceResourceSchemaAttributes = map[string]schema.Attribute{
	"add_edns_option_in_outgoing_query": schema.SingleNestedAttribute{
		Attributes:          Inheritance2InheritedBoolResourceSchemaAttributes,
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Field config for _add_edns_option_in_outgoing_query_ field from _Server_ object.",
	},
	"custom_root_ns_block": schema.SingleNestedAttribute{
		Attributes:          InheritedCustomRootNSBlockResourceSchemaAttributes,
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Optional. Field config for _custom_root_ns_block_ field from _Server_ object.",
	},
	"dnssec_validation_block": schema.SingleNestedAttribute{
		Attributes:          InheritedDNSSECValidationBlockResourceSchemaAttributes,
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Optional. Field config for _dnssec_validation_block_ field from _Server_ object.",
	},
	"ecs_block": schema.SingleNestedAttribute{
		Attributes:          InheritedECSBlockResourceSchemaAttributes,
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Optional. Field config for _ecs_block_ field from _Server_ object.",
	},
	"filter_aaaa_acl": schema.SingleNestedAttribute{
		Attributes:          InheritedACLItemsResourceSchemaAttributes,
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Optional. Field config for _filter_aaaa_acl_ field from _Server_ object.",
	},
	"filter_aaaa_on_v4": schema.SingleNestedAttribute{
		Attributes:          Inheritance2InheritedStringResourceSchemaAttributes,
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Optional. Field config for _filter_aaaa_on_v4_ field from _Server_ object.",
	},
	"forwarders_block": schema.SingleNestedAttribute{
		Attributes:          InheritedForwardersBlockResourceSchemaAttributes,
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Optional. Field config for _forwarders_block_ field from _Server_ object.",
	},
	"gss_tsig_enabled": schema.SingleNestedAttribute{
		Attributes:          Inheritance2InheritedBoolResourceSchemaAttributes,
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Optional. Field config for _gss_tsig_enabled_ field from _Server_ object.",
	},
	"kerberos_keys": schema.SingleNestedAttribute{
		Attributes:          InheritedKerberosKeysResourceSchemaAttributes,
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Optional. Field config for _kerberos_keys_ field from _Server_ object.",
	},
	"lame_ttl": schema.SingleNestedAttribute{
		Attributes:          Inheritance2InheritedUInt32DnsconfigResourceSchemaAttributes,
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Optional. Field config for _lame_ttl_ field from _Server_ object.",
	},
	"log_query_response": schema.SingleNestedAttribute{
		Attributes:          Inheritance2InheritedBoolResourceSchemaAttributes,
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Optional. Field config for _log_queries_response_ field from _Server_ object.",
	},
	"match_recursive_only": schema.SingleNestedAttribute{
		Attributes:          Inheritance2InheritedBoolResourceSchemaAttributes,
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Optional. Field config for _match_recursive_only_ field from _Server_ object.",
	},
	"max_cache_ttl": schema.SingleNestedAttribute{
		Attributes:          Inheritance2InheritedUInt32DnsconfigResourceSchemaAttributes,
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Optional. Field config for _max_cache_ttl_ field from _Server_ object.",
	},
	"max_negative_ttl": schema.SingleNestedAttribute{
		Attributes:          Inheritance2InheritedUInt32DnsconfigResourceSchemaAttributes,
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Optional. Field config for _max_negative_ttl_ field from _Server_ object.",
	},
	"minimal_responses": schema.SingleNestedAttribute{
		Attributes:          Inheritance2InheritedBoolResourceSchemaAttributes,
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Optional. Field config for _minimal_responses_ field from _Server_ object.",
	},
	"notify": schema.SingleNestedAttribute{
		Attributes:          Inheritance2InheritedBoolResourceSchemaAttributes,
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Field config for _notify_ field from _Server_ object.",
	},
	"query_acl": schema.SingleNestedAttribute{
		Attributes:          InheritedACLItemsResourceSchemaAttributes,
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Optional. Field config for _query_acl_ field from _Server_ object.",
	},
	"query_port": schema.SingleNestedAttribute{
		Attributes:          Inheritance2InheritedUInt32DnsconfigResourceSchemaAttributes,
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Optional. Field config for _query_port_ field from _Server_ object.",
	},
	"recursion_acl": schema.SingleNestedAttribute{
		Attributes:          InheritedACLItemsResourceSchemaAttributes,
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Optional. Field config for _recursion_acl_ field from _Server_ object.",
	},
	"recursion_enabled": schema.SingleNestedAttribute{
		Attributes:          Inheritance2InheritedBoolResourceSchemaAttributes,
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Optional. Field config for _recursion_enabled_ field from _Server_ object.",
	},
	"recursive_clients": schema.SingleNestedAttribute{
		Attributes:          Inheritance2InheritedUInt32DnsconfigResourceSchemaAttributes,
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Optional. Field config for _recursive_clients_ field from _Server_ object.",
	},
	"resolver_query_timeout": schema.SingleNestedAttribute{
		Attributes:          Inheritance2InheritedUInt32DnsconfigResourceSchemaAttributes,
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Optional. Field config for _resolver_query_timeout_ field from _Server_ object.",
	},
	"secondary_axfr_query_limit": schema.SingleNestedAttribute{
		Attributes:          Inheritance2InheritedUInt32DnsconfigResourceSchemaAttributes,
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Optional. Field config for _secondary_axfr_query_limit_ field from _Server_ object.",
	},
	"secondary_soa_query_limit": schema.SingleNestedAttribute{
		Attributes:          Inheritance2InheritedUInt32DnsconfigResourceSchemaAttributes,
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Optional. Field config for _secondary_soa_query_limit_ field from _Server_ object.",
	},
	"sort_list": schema.SingleNestedAttribute{
		Attributes:          InheritedSortListItemsResourceSchemaAttributes,
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Optional. Field config for _sort_list_ field from _Server object.",
	},
	"synthesize_address_records_from_https": schema.SingleNestedAttribute{
		Attributes:          Inheritance2InheritedBoolResourceSchemaAttributes,
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Field config for _synthesize_address_records_from_https_ field from _Server_ object.",
	},
	"transfer_acl": schema.SingleNestedAttribute{
		Attributes:          InheritedACLItemsResourceSchemaAttributes,
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Optional. Field config for _transfer_acl_ field from _Server_ object.",
	},
	"update_acl": schema.SingleNestedAttribute{
		Attributes:          InheritedACLItemsResourceSchemaAttributes,
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Optional. Field config for _update_acl_ field from _Server_ object.",
	},
	"use_forwarders_for_subzones": schema.SingleNestedAttribute{
		Attributes:          Inheritance2InheritedBoolResourceSchemaAttributes,
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Optional. Field config for _use_forwarders_for_subzones_ field from _Server_ object.",
	},
}

// ExpandServerInheritance converts a Terraform Object to SDK type
func ExpandServerInheritance(ctx context.Context, o types.Object, diags *diag.Diagnostics) *uddidns.ServerInheritance {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m ServerInheritanceModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

// Expand converts the Terraform model to SDK type
func (m *ServerInheritanceModel) Expand(ctx context.Context, diags *diag.Diagnostics) *uddidns.ServerInheritance {
	if m == nil {
		return nil
	}
	to := &uddidns.ServerInheritance{
		AddEdnsOptionInOutgoingQuery:      ExpandInheritance2InheritedBool(ctx, m.AddEdnsOptionInOutgoingQuery, diags),
		CustomRootNsBlock:                 ExpandInheritedCustomRootNSBlock(ctx, m.CustomRootNsBlock, diags),
		DnssecValidationBlock:             ExpandInheritedDNSSECValidationBlock(ctx, m.DnssecValidationBlock, diags),
		EcsBlock:                          ExpandInheritedECSBlock(ctx, m.EcsBlock, diags),
		FilterAaaaAcl:                     ExpandInheritedACLItems(ctx, m.FilterAaaaAcl, diags),
		FilterAaaaOnV4:                    ExpandInheritance2InheritedString(ctx, m.FilterAaaaOnV4, diags),
		ForwardersBlock:                   ExpandInheritedForwardersBlock(ctx, m.ForwardersBlock, diags),
		GssTsigEnabled:                    ExpandInheritance2InheritedBool(ctx, m.GssTsigEnabled, diags),
		KerberosKeys:                      ExpandInheritedKerberosKeys(ctx, m.KerberosKeys, diags),
		LameTtl:                           ExpandInheritance2InheritedUInt32Dnsconfig(ctx, m.LameTtl, diags),
		LogQueryResponse:                  ExpandInheritance2InheritedBool(ctx, m.LogQueryResponse, diags),
		MatchRecursiveOnly:                ExpandInheritance2InheritedBool(ctx, m.MatchRecursiveOnly, diags),
		MaxCacheTtl:                       ExpandInheritance2InheritedUInt32Dnsconfig(ctx, m.MaxCacheTtl, diags),
		MaxNegativeTtl:                    ExpandInheritance2InheritedUInt32Dnsconfig(ctx, m.MaxNegativeTtl, diags),
		MinimalResponses:                  ExpandInheritance2InheritedBool(ctx, m.MinimalResponses, diags),
		Notify:                            ExpandInheritance2InheritedBool(ctx, m.Notify, diags),
		QueryAcl:                          ExpandInheritedACLItems(ctx, m.QueryAcl, diags),
		QueryPort:                         ExpandInheritance2InheritedUInt32Dnsconfig(ctx, m.QueryPort, diags),
		RecursionAcl:                      ExpandInheritedACLItems(ctx, m.RecursionAcl, diags),
		RecursionEnabled:                  ExpandInheritance2InheritedBool(ctx, m.RecursionEnabled, diags),
		RecursiveClients:                  ExpandInheritance2InheritedUInt32Dnsconfig(ctx, m.RecursiveClients, diags),
		ResolverQueryTimeout:              ExpandInheritance2InheritedUInt32Dnsconfig(ctx, m.ResolverQueryTimeout, diags),
		SecondaryAxfrQueryLimit:           ExpandInheritance2InheritedUInt32Dnsconfig(ctx, m.SecondaryAxfrQueryLimit, diags),
		SecondarySoaQueryLimit:            ExpandInheritance2InheritedUInt32Dnsconfig(ctx, m.SecondarySoaQueryLimit, diags),
		SortList:                          ExpandInheritedSortListItems(ctx, m.SortList, diags),
		SynthesizeAddressRecordsFromHttps: ExpandInheritance2InheritedBool(ctx, m.SynthesizeAddressRecordsFromHttps, diags),
		TransferAcl:                       ExpandInheritedACLItems(ctx, m.TransferAcl, diags),
		UpdateAcl:                         ExpandInheritedACLItems(ctx, m.UpdateAcl, diags),
		UseForwardersForSubzones:          ExpandInheritance2InheritedBool(ctx, m.UseForwardersForSubzones, diags),
	}
	return to
}

// FlattenServerInheritance converts an SDK type to Terraform Object
func FlattenServerInheritance(ctx context.Context, from *uddidns.ServerInheritance, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(ServerInheritanceAttrTypes)
	}
	m := &ServerInheritanceModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, ServerInheritanceAttrTypes, m)
	diags.Append(d...)
	return t
}

// Flatten populates the Terraform model from SDK type
func (m *ServerInheritanceModel) Flatten(ctx context.Context, from *uddidns.ServerInheritance, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.AddEdnsOptionInOutgoingQuery = FlattenInheritance2InheritedBool(ctx, from.AddEdnsOptionInOutgoingQuery, diags)
	m.CustomRootNsBlock = FlattenInheritedCustomRootNSBlock(ctx, from.CustomRootNsBlock, diags)
	m.DnssecValidationBlock = FlattenInheritedDNSSECValidationBlock(ctx, from.DnssecValidationBlock, diags)
	m.EcsBlock = FlattenInheritedECSBlock(ctx, from.EcsBlock, diags)
	m.FilterAaaaAcl = FlattenInheritedACLItems(ctx, from.FilterAaaaAcl, diags)
	m.FilterAaaaOnV4 = FlattenInheritance2InheritedString(ctx, from.FilterAaaaOnV4, diags)
	m.ForwardersBlock = FlattenInheritedForwardersBlock(ctx, from.ForwardersBlock, diags)
	m.GssTsigEnabled = FlattenInheritance2InheritedBool(ctx, from.GssTsigEnabled, diags)
	m.KerberosKeys = FlattenInheritedKerberosKeys(ctx, from.KerberosKeys, diags)
	m.LameTtl = FlattenInheritance2InheritedUInt32Dnsconfig(ctx, from.LameTtl, diags)
	m.LogQueryResponse = FlattenInheritance2InheritedBool(ctx, from.LogQueryResponse, diags)
	m.MatchRecursiveOnly = FlattenInheritance2InheritedBool(ctx, from.MatchRecursiveOnly, diags)
	m.MaxCacheTtl = FlattenInheritance2InheritedUInt32Dnsconfig(ctx, from.MaxCacheTtl, diags)
	m.MaxNegativeTtl = FlattenInheritance2InheritedUInt32Dnsconfig(ctx, from.MaxNegativeTtl, diags)
	m.MinimalResponses = FlattenInheritance2InheritedBool(ctx, from.MinimalResponses, diags)
	m.Notify = FlattenInheritance2InheritedBool(ctx, from.Notify, diags)
	m.QueryAcl = FlattenInheritedACLItems(ctx, from.QueryAcl, diags)
	m.QueryPort = FlattenInheritance2InheritedUInt32Dnsconfig(ctx, from.QueryPort, diags)
	m.RecursionAcl = FlattenInheritedACLItems(ctx, from.RecursionAcl, diags)
	m.RecursionEnabled = FlattenInheritance2InheritedBool(ctx, from.RecursionEnabled, diags)
	m.RecursiveClients = FlattenInheritance2InheritedUInt32Dnsconfig(ctx, from.RecursiveClients, diags)
	m.ResolverQueryTimeout = FlattenInheritance2InheritedUInt32Dnsconfig(ctx, from.ResolverQueryTimeout, diags)
	m.SecondaryAxfrQueryLimit = FlattenInheritance2InheritedUInt32Dnsconfig(ctx, from.SecondaryAxfrQueryLimit, diags)
	m.SecondarySoaQueryLimit = FlattenInheritance2InheritedUInt32Dnsconfig(ctx, from.SecondarySoaQueryLimit, diags)
	m.SortList = FlattenInheritedSortListItems(ctx, from.SortList, diags)
	m.SynthesizeAddressRecordsFromHttps = FlattenInheritance2InheritedBool(ctx, from.SynthesizeAddressRecordsFromHttps, diags)
	m.TransferAcl = FlattenInheritedACLItems(ctx, from.TransferAcl, diags)
	m.UpdateAcl = FlattenInheritedACLItems(ctx, from.UpdateAcl, diags)
	m.UseForwardersForSubzones = FlattenInheritance2InheritedBool(ctx, from.UseForwardersForSubzones, diags)
}
