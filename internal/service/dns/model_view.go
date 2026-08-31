package dns

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/mapvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/mapdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectdefault"
	objectplanmodifier "github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"

	coremodel "github.com/infobloxopen/terraform-provider-infoblox/internal/core/model/dns"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
	planmod "github.com/infobloxopen/terraform-provider-infoblox/internal/planmodifiers"
	importmod "github.com/infobloxopen/terraform-provider-infoblox/internal/planmodifiers/import"
	internaltypes "github.com/infobloxopen/terraform-provider-infoblox/internal/types"
	customvalidator "github.com/infobloxopen/terraform-provider-infoblox/internal/validator"
)

type ViewModel struct {
	Id   types.String `tfsdk:"id"`
	NIOS types.Object `tfsdk:"nios"`
	UDDI types.Object `tfsdk:"uddi"`
}

var ViewAttrTypes = map[string]attr.Type{
	"id":   types.StringType,
	"nios": types.ObjectType{AttrTypes: NIOSViewAttrTypes},
	"uddi": types.ObjectType{AttrTypes: UDDIViewAttrTypes},
}

type NIOSViewModel struct {
	BlacklistAction                  types.String                     `tfsdk:"blacklist_action"`
	BlacklistLogQuery                types.Bool                       `tfsdk:"blacklist_log_query"`
	BlacklistRedirectAddresses       types.List                       `tfsdk:"blacklist_redirect_addresses"`
	BlacklistRedirectTtl             types.Int64                      `tfsdk:"blacklist_redirect_ttl"`
	BlacklistRulesets                types.List                       `tfsdk:"blacklist_rulesets"`
	Comment                          types.String                     `tfsdk:"comment"`
	CustomRootNameServers            types.List                       `tfsdk:"custom_root_name_servers"`
	DdnsForceCreationTimestampUpdate types.Bool                       `tfsdk:"ddns_force_creation_timestamp_update"`
	DdnsPrincipalGroup               types.String                     `tfsdk:"ddns_principal_group"`
	DdnsPrincipalTracking            types.Bool                       `tfsdk:"ddns_principal_tracking"`
	DdnsRestrictPatterns             types.Bool                       `tfsdk:"ddns_restrict_patterns"`
	DdnsRestrictPatternsList         types.List                       `tfsdk:"ddns_restrict_patterns_list"`
	DdnsRestrictProtected            types.Bool                       `tfsdk:"ddns_restrict_protected"`
	DdnsRestrictSecure               types.Bool                       `tfsdk:"ddns_restrict_secure"`
	DdnsRestrictStatic               types.Bool                       `tfsdk:"ddns_restrict_static"`
	Disable                          types.Bool                       `tfsdk:"disable"`
	Dns64Enabled                     types.Bool                       `tfsdk:"dns64_enabled"`
	Dns64Groups                      internaltypes.UnorderedListValue `tfsdk:"dns64_groups"`
	DnssecEnabled                    types.Bool                       `tfsdk:"dnssec_enabled"`
	DnssecExpiredSignaturesEnabled   types.Bool                       `tfsdk:"dnssec_expired_signatures_enabled"`
	DnssecNegativeTrustAnchors       types.List                       `tfsdk:"dnssec_negative_trust_anchors"`
	DnssecTrustedKeys                types.List                       `tfsdk:"dnssec_trusted_keys"`
	DnssecValidationEnabled          types.Bool                       `tfsdk:"dnssec_validation_enabled"`
	EdnsUdpSize                      types.Int64                      `tfsdk:"edns_udp_size"`
	EnableBlacklist                  types.Bool                       `tfsdk:"enable_blacklist"`
	EnableFixedRrsetOrderFqdns       types.Bool                       `tfsdk:"enable_fixed_rrset_order_fqdns"`
	EnableMatchRecursiveOnly         types.Bool                       `tfsdk:"enable_match_recursive_only"`
	ExtAttrs                         types.Map                        `tfsdk:"ext_attrs"`
	ExtAttrsAll                      types.Map                        `tfsdk:"ext_attrs_all"`
	FilterAaaa                       types.String                     `tfsdk:"filter_aaaa"`
	FilterAaaaList                   types.List                       `tfsdk:"filter_aaaa_list"`
	FixedRrsetOrderFqdns             types.List                       `tfsdk:"fixed_rrset_order_fqdns"`
	ForwardOnly                      types.Bool                       `tfsdk:"forward_only"`
	Forwarders                       types.List                       `tfsdk:"forwarders"`
	LastQueriedAcl                   types.List                       `tfsdk:"last_queried_acl"`
	MatchClients                     types.List                       `tfsdk:"match_clients"`
	MatchDestinations                types.List                       `tfsdk:"match_destinations"`
	MaxCacheTtl                      types.Int64                      `tfsdk:"max_cache_ttl"`
	MaxNcacheTtl                     types.Int64                      `tfsdk:"max_ncache_ttl"`
	MaxUdpSize                       types.Int64                      `tfsdk:"max_udp_size"`
	Name                             types.String                     `tfsdk:"name"`
	NetworkView                      types.String                     `tfsdk:"network_view"`
	NotifyDelay                      types.Int64                      `tfsdk:"notify_delay"`
	NxdomainLogQuery                 types.Bool                       `tfsdk:"nxdomain_log_query"`
	NxdomainRedirect                 types.Bool                       `tfsdk:"nxdomain_redirect"`
	NxdomainRedirectAddresses        types.List                       `tfsdk:"nxdomain_redirect_addresses"`
	NxdomainRedirectAddressesV6      types.List                       `tfsdk:"nxdomain_redirect_addresses_v6"`
	NxdomainRedirectTtl              types.Int64                      `tfsdk:"nxdomain_redirect_ttl"`
	NxdomainRulesets                 types.List                       `tfsdk:"nxdomain_rulesets"`
	Recursion                        types.Bool                       `tfsdk:"recursion"`
	ResponseRateLimiting             types.Object                     `tfsdk:"response_rate_limiting"`
	RootNameServerType               types.String                     `tfsdk:"root_name_server_type"`
	RpzDropIpRuleEnabled             types.Bool                       `tfsdk:"rpz_drop_ip_rule_enabled"`
	RpzDropIpRuleMinPrefixLengthIpv4 types.Int64                      `tfsdk:"rpz_drop_ip_rule_min_prefix_length_ipv4"`
	RpzDropIpRuleMinPrefixLengthIpv6 types.Int64                      `tfsdk:"rpz_drop_ip_rule_min_prefix_length_ipv6"`
	RpzQnameWaitRecurse              types.Bool                       `tfsdk:"rpz_qname_wait_recurse"`
	ScavengingSettings               types.Object                     `tfsdk:"scavenging_settings"`
	Sortlist                         types.List                       `tfsdk:"sortlist"`
}

var NIOSViewAttrTypes = map[string]attr.Type{
	"blacklist_action":                        types.StringType,
	"blacklist_log_query":                     types.BoolType,
	"blacklist_redirect_addresses":            types.ListType{ElemType: types.StringType},
	"blacklist_redirect_ttl":                  types.Int64Type,
	"blacklist_rulesets":                      types.ListType{ElemType: types.StringType},
	"comment":                                 types.StringType,
	"custom_root_name_servers":                types.ListType{ElemType: types.ObjectType{AttrTypes: ViewCustomRootNameServersAttrTypes}},
	"ddns_force_creation_timestamp_update":    types.BoolType,
	"ddns_principal_group":                    types.StringType,
	"ddns_principal_tracking":                 types.BoolType,
	"ddns_restrict_patterns":                  types.BoolType,
	"ddns_restrict_patterns_list":             types.ListType{ElemType: types.StringType},
	"ddns_restrict_protected":                 types.BoolType,
	"ddns_restrict_secure":                    types.BoolType,
	"ddns_restrict_static":                    types.BoolType,
	"disable":                                 types.BoolType,
	"dns64_enabled":                           types.BoolType,
	"dns64_groups":                            internaltypes.UnorderedListOfStringType,
	"dnssec_enabled":                          types.BoolType,
	"dnssec_expired_signatures_enabled":       types.BoolType,
	"dnssec_negative_trust_anchors":           types.ListType{ElemType: types.StringType},
	"dnssec_trusted_keys":                     types.ListType{ElemType: types.ObjectType{AttrTypes: ViewDnssecTrustedKeysAttrTypes}},
	"dnssec_validation_enabled":               types.BoolType,
	"edns_udp_size":                           types.Int64Type,
	"enable_blacklist":                        types.BoolType,
	"enable_fixed_rrset_order_fqdns":          types.BoolType,
	"enable_match_recursive_only":             types.BoolType,
	"ext_attrs":                               types.MapType{ElemType: types.StringType},
	"ext_attrs_all":                           types.MapType{ElemType: types.StringType},
	"filter_aaaa":                             types.StringType,
	"filter_aaaa_list":                        types.ListType{ElemType: types.ObjectType{AttrTypes: ViewFilterAaaaListAttrTypes}},
	"fixed_rrset_order_fqdns":                 types.ListType{ElemType: types.ObjectType{AttrTypes: ViewFixedRrsetOrderFqdnsAttrTypes}},
	"forward_only":                            types.BoolType,
	"forwarders":                              types.ListType{ElemType: types.StringType},
	"last_queried_acl":                        types.ListType{ElemType: types.ObjectType{AttrTypes: ViewLastQueriedAclAttrTypes}},
	"match_clients":                           types.ListType{ElemType: types.ObjectType{AttrTypes: ViewMatchClientsAttrTypes}},
	"match_destinations":                      types.ListType{ElemType: types.ObjectType{AttrTypes: ViewMatchDestinationsAttrTypes}},
	"max_cache_ttl":                           types.Int64Type,
	"max_ncache_ttl":                          types.Int64Type,
	"max_udp_size":                            types.Int64Type,
	"name":                                    types.StringType,
	"network_view":                            types.StringType,
	"notify_delay":                            types.Int64Type,
	"nxdomain_log_query":                      types.BoolType,
	"nxdomain_redirect":                       types.BoolType,
	"nxdomain_redirect_addresses":             types.ListType{ElemType: types.StringType},
	"nxdomain_redirect_addresses_v6":          types.ListType{ElemType: types.StringType},
	"nxdomain_redirect_ttl":                   types.Int64Type,
	"nxdomain_rulesets":                       types.ListType{ElemType: types.StringType},
	"recursion":                               types.BoolType,
	"response_rate_limiting":                  types.ObjectType{AttrTypes: ViewResponseRateLimitingAttrTypes},
	"root_name_server_type":                   types.StringType,
	"rpz_drop_ip_rule_enabled":                types.BoolType,
	"rpz_drop_ip_rule_min_prefix_length_ipv4": types.Int64Type,
	"rpz_drop_ip_rule_min_prefix_length_ipv6": types.Int64Type,
	"rpz_qname_wait_recurse":                  types.BoolType,
	"scavenging_settings":                     types.ObjectType{AttrTypes: ViewScavengingSettingsAttrTypes},
	"sortlist":                                types.ListType{ElemType: types.ObjectType{AttrTypes: ViewSortlistAttrTypes}},
}

type UDDIViewModel struct {
	AddEdnsOptionInOutgoingQuery                types.Bool   `tfsdk:"add_edns_option_in_outgoing_query"`
	Comment                                     types.String `tfsdk:"comment"`
	CompartmentId                               types.String `tfsdk:"compartment_id"`
	CustomRootNs                                types.List   `tfsdk:"custom_root_ns"`
	CustomRootNsEnabled                         types.Bool   `tfsdk:"custom_root_ns_enabled"`
	Disabled                                    types.Bool   `tfsdk:"disabled"`
	DnssecEnableValidation                      types.Bool   `tfsdk:"dnssec_enable_validation"`
	DnssecEnabled                               types.Bool   `tfsdk:"dnssec_enabled"`
	DnssecTrustAnchors                          types.List   `tfsdk:"dnssec_trust_anchors"`
	DnssecValidateExpiry                        types.Bool   `tfsdk:"dnssec_validate_expiry"`
	DtcConfig                                   types.Object `tfsdk:"dtc_config"`
	EcsEnabled                                  types.Bool   `tfsdk:"ecs_enabled"`
	EcsForwarding                               types.Bool   `tfsdk:"ecs_forwarding"`
	EcsPrefixV4                                 types.Int64  `tfsdk:"ecs_prefix_v4"`
	EcsPrefixV6                                 types.Int64  `tfsdk:"ecs_prefix_v6"`
	EcsZones                                    types.List   `tfsdk:"ecs_zones"`
	EdnsUdpSize                                 types.Int64  `tfsdk:"edns_udp_size"`
	FilterAaaaAcl                               types.List   `tfsdk:"filter_aaaa_acl"`
	FilterAaaaOnV4                              types.String `tfsdk:"filter_aaaa_on_v4"`
	Forwarders                                  types.List   `tfsdk:"forwarders"`
	ForwardersOnly                              types.Bool   `tfsdk:"forwarders_only"`
	GssTsigEnabled                              types.Bool   `tfsdk:"gss_tsig_enabled"`
	InheritanceSources                          types.Object `tfsdk:"inheritance_sources"`
	IpSpaces                                    types.List   `tfsdk:"ip_spaces"`
	LameTtl                                     types.Int64  `tfsdk:"lame_ttl"`
	MatchClientsAcl                             types.List   `tfsdk:"match_clients_acl"`
	MatchDestinationsAcl                        types.List   `tfsdk:"match_destinations_acl"`
	MatchRecursiveOnly                          types.Bool   `tfsdk:"match_recursive_only"`
	MaxCacheTtl                                 types.Int64  `tfsdk:"max_cache_ttl"`
	MaxNegativeTtl                              types.Int64  `tfsdk:"max_negative_ttl"`
	MaxUdpSize                                  types.Int64  `tfsdk:"max_udp_size"`
	MinimalResponses                            types.Bool   `tfsdk:"minimal_responses"`
	Name                                        types.String `tfsdk:"name"`
	Notify                                      types.Bool   `tfsdk:"notify"`
	QueryAcl                                    types.List   `tfsdk:"query_acl"`
	RecursionAcl                                types.List   `tfsdk:"recursion_acl"`
	RecursionEnabled                            types.Bool   `tfsdk:"recursion_enabled"`
	SortList                                    types.List   `tfsdk:"sort_list"`
	SynthesizeAddressRecordsFromHttps           types.Bool   `tfsdk:"synthesize_address_records_from_https"`
	Tags                                        types.Map    `tfsdk:"tags"`
	TagsAll                                     types.Map    `tfsdk:"tags_all"`
	TransferAcl                                 types.List   `tfsdk:"transfer_acl"`
	UpdateAcl                                   types.List   `tfsdk:"update_acl"`
	UseForwardersForSubzones                    types.Bool   `tfsdk:"use_forwarders_for_subzones"`
	UseRootForwardersForLocalResolutionWithB1td types.Bool   `tfsdk:"use_root_forwarders_for_local_resolution_with_b1td"`
	ZoneAuthority                               types.Object `tfsdk:"zone_authority"`
}

var UDDIViewAttrTypes = map[string]attr.Type{
	"add_edns_option_in_outgoing_query":     types.BoolType,
	"comment":                               types.StringType,
	"compartment_id":                        types.StringType,
	"custom_root_ns":                        types.ListType{ElemType: types.ObjectType{AttrTypes: RootNSAttrTypes}},
	"custom_root_ns_enabled":                types.BoolType,
	"disabled":                              types.BoolType,
	"dnssec_enable_validation":              types.BoolType,
	"dnssec_enabled":                        types.BoolType,
	"dnssec_trust_anchors":                  types.ListType{ElemType: types.ObjectType{AttrTypes: TrustAnchorAttrTypes}},
	"dnssec_validate_expiry":                types.BoolType,
	"dtc_config":                            types.ObjectType{AttrTypes: DTCConfigAttrTypes},
	"ecs_enabled":                           types.BoolType,
	"ecs_forwarding":                        types.BoolType,
	"ecs_prefix_v4":                         types.Int64Type,
	"ecs_prefix_v6":                         types.Int64Type,
	"ecs_zones":                             types.ListType{ElemType: types.ObjectType{AttrTypes: ECSZoneAttrTypes}},
	"edns_udp_size":                         types.Int64Type,
	"filter_aaaa_acl":                       types.ListType{ElemType: types.ObjectType{AttrTypes: ACLItemAttrTypes}},
	"filter_aaaa_on_v4":                     types.StringType,
	"forwarders":                            types.ListType{ElemType: types.ObjectType{AttrTypes: ForwarderAttrTypes}},
	"forwarders_only":                       types.BoolType,
	"gss_tsig_enabled":                      types.BoolType,
	"inheritance_sources":                   types.ObjectType{AttrTypes: ViewInheritanceAttrTypes},
	"ip_spaces":                             types.ListType{ElemType: types.StringType},
	"lame_ttl":                              types.Int64Type,
	"match_clients_acl":                     types.ListType{ElemType: types.ObjectType{AttrTypes: ACLItemAttrTypes}},
	"match_destinations_acl":                types.ListType{ElemType: types.ObjectType{AttrTypes: ACLItemAttrTypes}},
	"match_recursive_only":                  types.BoolType,
	"max_cache_ttl":                         types.Int64Type,
	"max_negative_ttl":                      types.Int64Type,
	"max_udp_size":                          types.Int64Type,
	"minimal_responses":                     types.BoolType,
	"name":                                  types.StringType,
	"notify":                                types.BoolType,
	"query_acl":                             types.ListType{ElemType: types.ObjectType{AttrTypes: ACLItemAttrTypes}},
	"recursion_acl":                         types.ListType{ElemType: types.ObjectType{AttrTypes: ACLItemAttrTypes}},
	"recursion_enabled":                     types.BoolType,
	"sort_list":                             types.ListType{ElemType: types.ObjectType{AttrTypes: SortListItemAttrTypes}},
	"synthesize_address_records_from_https": types.BoolType,
	"tags":                                  types.MapType{ElemType: types.StringType},
	"tags_all":                              types.MapType{ElemType: types.StringType},
	"transfer_acl":                          types.ListType{ElemType: types.ObjectType{AttrTypes: ACLItemAttrTypes}},
	"update_acl":                            types.ListType{ElemType: types.ObjectType{AttrTypes: ACLItemAttrTypes}},
	"use_forwarders_for_subzones":           types.BoolType,
	"use_root_forwarders_for_local_resolution_with_b1td": types.BoolType,
	"zone_authority": types.ObjectType{AttrTypes: ZoneAuthorityAttrTypes},
}

const (
	ViewInheritanceType = "full"
	ViewReturnFields    = "blacklist_action,blacklist_log_query,blacklist_redirect_addresses,blacklist_redirect_ttl,blacklist_rulesets,cloud_info,comment,custom_root_name_servers,ddns_force_creation_timestamp_update,ddns_principal_group,ddns_principal_tracking,ddns_restrict_patterns,ddns_restrict_patterns_list,ddns_restrict_protected,ddns_restrict_secure,ddns_restrict_static,disable,dns64_enabled,dns64_groups,dnssec_enabled,dnssec_expired_signatures_enabled,dnssec_negative_trust_anchors,dnssec_trusted_keys,dnssec_validation_enabled,edns_udp_size,enable_blacklist,enable_fixed_rrset_order_fqdns,enable_match_recursive_only,extattrs,filter_aaaa,filter_aaaa_list,fixed_rrset_order_fqdns,forward_only,forwarders,is_default,last_queried_acl,match_clients,match_destinations,max_cache_ttl,max_ncache_ttl,max_udp_size,name,network_view,notify_delay,nxdomain_log_query,nxdomain_redirect,nxdomain_redirect_addresses,nxdomain_redirect_addresses_v6,nxdomain_redirect_ttl,nxdomain_rulesets,recursion,response_rate_limiting,root_name_server_type,rpz_drop_ip_rule_enabled,rpz_drop_ip_rule_min_prefix_length_ipv4,rpz_drop_ip_rule_min_prefix_length_ipv6,rpz_qname_wait_recurse,scavenging_settings,sortlist,use_blacklist,use_ddns_force_creation_timestamp_update,use_ddns_patterns_restriction,use_ddns_principal_security,use_ddns_restrict_protected,use_ddns_restrict_static,use_dns64,use_dnssec,use_edns_udp_size,use_filter_aaaa,use_fixed_rrset_order_fqdns,use_forwarders,use_max_cache_ttl,use_max_ncache_ttl,use_max_udp_size,use_nxdomain_redirect,use_recursion,use_response_rate_limiting,use_root_name_server,use_rpz_drop_ip_rule,use_rpz_qname_wait_recurse,use_scavenging_settings,use_sortlist"
)

var ViewResourceSchemaAttributes = map[string]schema.Attribute{
	"id": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The reference to the object.",
	},
	"nios": schema.SingleNestedAttribute{
		Optional:            true,
		MarkdownDescription: "NIOS backend-specific fields.",
		Attributes:          ViewResourceNiosSchemaAttributes,
	},
	"uddi": schema.SingleNestedAttribute{
		Optional:            true,
		MarkdownDescription: "UDDI backend-specific fields.",
		Attributes:          ViewResourceUddiSchemaAttributes,
	},
}

var ViewResourceNiosSchemaAttributes = map[string]schema.Attribute{
	"blacklist_action": schema.StringAttribute{
		Default: stringdefault.StaticString("REDIRECT"),
		Validators: []validator.String{
			stringvalidator.OneOf("REDIRECT", "REFUSE"),
		},
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The action to perform when a domain name matches the pattern defined in a rule that is specified by the blacklist_ruleset method. Valid values are \"REDIRECT\" or \"REFUSE\". The default value is \"REFUSE\".",
	},
	"blacklist_log_query": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "The flag that indicates whether blacklist redirection queries are logged. Specify \"true\" to enable logging, or \"false\" to disable it. The default value is \"false\".",
	},
	"blacklist_redirect_addresses": schema.ListAttribute{
		ElementType: types.StringType,
		Optional:    true,
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
		},
		MarkdownDescription: "The array of IP addresses the appliance includes in the response it sends in place of a blacklisted IP address.",
	},
	"blacklist_redirect_ttl": schema.Int64Attribute{
		Optional:            true,
		Computed:            true,
		Default:             int64default.StaticInt64(60),
		MarkdownDescription: "The Time To Live (TTL) value of the synthetic DNS responses resulted from blacklist redirection. The TTL value is a 32-bit unsigned integer that represents the TTL in seconds.",
	},
	"blacklist_rulesets": schema.ListAttribute{
		ElementType: types.StringType,
		Optional:    true,
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
		},
		MarkdownDescription: "The name of the Ruleset object assigned at the Grid level for blacklist redirection.",
	},
	"comment": schema.StringAttribute{
		Optional: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "Comment for the DNS view; maximum 64 characters.",
	},
	"custom_root_name_servers": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: ViewCustomRootNameServersResourceSchemaAttributes,
		},
		Optional: true,
		Computed: true,
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
		},
		MarkdownDescription: "The list of customized root name servers. You can either select and use Internet root name servers or specify custom root name servers by providing a host name and IP address to which the Infoblox appliance can send queries. Include the specified parameter to set the attribute value. Omit the parameter to retrieve the attribute value.",
	},
	"ddns_force_creation_timestamp_update": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Defines whether creation timestamp of RR should be updated ' when DDNS update happens even if there is no change to ' the RR.",
	},
	"ddns_principal_group": schema.StringAttribute{
		Optional: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "The DDNS Principal cluster group name.",
	},
	"ddns_principal_tracking": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "The flag that indicates whether the DDNS principal track is enabled or disabled.",
	},
	"ddns_restrict_patterns": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "The flag that indicates whether an option to restrict DDNS update request based on FQDN patterns is enabled or disabled.",
	},
	"ddns_restrict_patterns_list": schema.ListAttribute{
		ElementType: types.StringType,
		Optional:    true,
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
		},
		MarkdownDescription: "The unordered list of restriction patterns for an option of to restrict DDNS updates based on FQDN patterns.",
	},
	"ddns_restrict_protected": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "The flag that indicates whether an option to restrict DDNS update request to protected resource records is enabled or disabled.",
	},
	"ddns_restrict_secure": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "The flag that indicates whether DDNS update request for principal other than target resource record's principal is restricted.",
	},
	"ddns_restrict_static": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "The flag that indicates whether an option to restrict DDNS update request to resource records which are marked as 'STATIC' is enabled or disabled.",
	},
	"disable": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Determines if the DNS view is disabled or not. When this is set to False, the DNS view is enabled.",
	},
	"dns64_enabled": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Determines if the DNS64 s enabled or not.",
	},
	"dns64_groups": schema.ListAttribute{
		ElementType: types.StringType,
		Optional:    true,
		CustomType:  internaltypes.UnorderedListOfStringType,
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
		},
		MarkdownDescription: "The list of DNS64 synthesis groups associated with this DNS view.",
	},
	"dnssec_enabled": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Determines if the DNS security extension is enabled or not.",
	},
	"dnssec_expired_signatures_enabled": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Determines if the DNS security extension accepts expired signatures or not.",
	},
	"dnssec_negative_trust_anchors": schema.ListAttribute{
		ElementType: types.StringType,
		Optional:    true,
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
		},
		MarkdownDescription: "A list of zones for which the server does not perform DNSSEC validation.",
	},
	"dnssec_trusted_keys": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: ViewDnssecTrustedKeysResourceSchemaAttributes,
		},
		Optional: true,
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
		},
		MarkdownDescription: "The list of trusted keys for the DNS security extension.",
	},
	"dnssec_validation_enabled": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(true),
		MarkdownDescription: "Determines if the DNS security validation is enabled or not.",
	},
	"edns_udp_size": schema.Int64Attribute{
		Optional:            true,
		Computed:            true,
		Default:             int64default.StaticInt64(1220),
		MarkdownDescription: "Advertises the EDNS0 buffer size to the upstream server. The value should be between 512 and 4096 bytes. The recommended value is between 512 and 1220 bytes.",
	},
	"enable_blacklist": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Determines if the blacklist in a DNS view is enabled or not.",
	},
	"enable_fixed_rrset_order_fqdns": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Determines if the fixed RRset order FQDN is enabled or not.",
	},
	"enable_match_recursive_only": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Determines if the 'match-recursive-only' option in a DNS view is enabled or not.",
	},
	"ext_attrs": schema.MapAttribute{
		Optional:    true,
		Computed:    true,
		ElementType: types.StringType,
		Default:     mapdefault.StaticValue(types.MapNull(types.StringType)),
		Validators: []validator.Map{
			mapvalidator.SizeAtLeast(1),
		},
		MarkdownDescription: "Extensible attributes associated with the object. For valid values for extensible attributes, see {extattrs:values}.",
	},
	"ext_attrs_all": schema.MapAttribute{
		Computed:            true,
		ElementType:         types.StringType,
		MarkdownDescription: "All ext_attrs including Terraform Internal ID and inherited attributes.",
		PlanModifiers: []planmodifier.Map{
			importmod.AssociateInternalId(),
		},
	},
	"filter_aaaa": schema.StringAttribute{
		Default: stringdefault.StaticString("NO"),
		Validators: []validator.String{
			stringvalidator.OneOf("YES", "NO", "BREAK_DNSSEC"),
		},
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The type of AAAA filtering for this DNS view object.",
	},
	"filter_aaaa_list": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: ViewFilterAaaaListResourceSchemaAttributes,
		},
		Optional: true,
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
		},
		MarkdownDescription: "Applies AAAA filtering to a named ACL, or to a list of IPv4/IPv6 addresses and networks from which queries are received. This field does not allow TSIG keys.",
	},
	"fixed_rrset_order_fqdns": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: ViewFixedRrsetOrderFqdnsResourceSchemaAttributes,
		},
		Optional: true,
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
		},
		MarkdownDescription: "The fixed RRset order FQDN. If this field does not contain an empty value, the appliance will automatically set the enable_fixed_rrset_order_fqdns field to 'true', unless the same request sets the enable field to 'false'.",
	},
	"forward_only": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Determines if this DNS view sends queries to forwarders only or not. When the value is True, queries are sent to forwarders only, and not to other internal or Internet root servers.",
	},
	"forwarders": schema.ListAttribute{
		ElementType: types.StringType,
		Optional:    true,
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
		},
		MarkdownDescription: "The list of forwarders for the DNS view. A forwarder is a name server to which other name servers first send their off-site queries. The forwarder builds up a cache of information, avoiding the need for other name servers to send queries off-site.",
	},
	"last_queried_acl": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: ViewLastQueriedAclResourceSchemaAttributes,
		},
		Optional: true,
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
		},
		MarkdownDescription: "Determines last queried ACL for the specified IPv4 or IPv6 addresses and networks in scavenging settings.",
	},
	"match_clients": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: ViewMatchClientsResourceSchemaAttributes,
		},
		Optional: true,
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
		},
		MarkdownDescription: "A list of forwarders for the match clients. This list specifies a named ACL, or a list of IPv4/IPv6 addresses, networks, TSIG keys of clients that are allowed or denied access to the DNS view.",
	},
	"match_destinations": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: ViewMatchDestinationsResourceSchemaAttributes,
		},
		Optional: true,
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
		},
		MarkdownDescription: "A list of forwarders for the match destinations. This list specifies a name ACL, or a list of IPv4/IPv6 addresses, networks, TSIG keys of clients that are allowed or denied access to the DNS view.",
	},
	"max_cache_ttl": schema.Int64Attribute{
		Optional:            true,
		Computed:            true,
		Default:             int64default.StaticInt64(604800),
		MarkdownDescription: "The maximum number of seconds to cache ordinary (positive) answers.",
	},
	"max_ncache_ttl": schema.Int64Attribute{
		Optional:            true,
		Computed:            true,
		Default:             int64default.StaticInt64(10800),
		MarkdownDescription: "The maximum number of seconds to cache negative (NXDOMAIN) answers.",
	},
	"max_udp_size": schema.Int64Attribute{
		Optional:            true,
		Computed:            true,
		Default:             int64default.StaticInt64(1220),
		MarkdownDescription: "The value is used by authoritative DNS servers to never send DNS responses larger than the configured value. The value should be between 512 and 4096 bytes. The recommended value is between 512 and 1220 bytes.",
	},
	"name": schema.StringAttribute{
		Required: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "Name of the DNS view.",
	},
	"network_view": schema.StringAttribute{
		Default:  stringdefault.StaticString("default"),
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "The name of the network view object associated with this DNS view.",
	},
	"notify_delay": schema.Int64Attribute{
		Optional:            true,
		Computed:            true,
		Default:             int64default.StaticInt64(5),
		MarkdownDescription: "The number of seconds of delay the notify messages are sent to secondaries.",
	},
	"nxdomain_log_query": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "The flag that indicates whether NXDOMAIN redirection queries are logged. Specify \"true\" to enable logging, or \"false\" to disable it. The default value is \"false\".",
	},
	"nxdomain_redirect": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Determines if NXDOMAIN redirection in a DNS view is enabled or not.",
	},
	"nxdomain_redirect_addresses": schema.ListAttribute{
		ElementType: types.StringType,
		Optional:    true,
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
		},
		MarkdownDescription: "The array with IPv4 addresses the appliance includes in the response it sends in place of an NXDOMAIN response.",
	},
	"nxdomain_redirect_addresses_v6": schema.ListAttribute{
		ElementType: types.StringType,
		Optional:    true,
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
		},
		MarkdownDescription: "The array with IPv6 addresses the appliance includes in the response it sends in place of an NXDOMAIN response.",
	},
	"nxdomain_redirect_ttl": schema.Int64Attribute{
		Optional:            true,
		Computed:            true,
		Default:             int64default.StaticInt64(60),
		MarkdownDescription: "The Time To Live (TTL) value of the synthetic DNS responses resulted from NXDOMAIN redirection. The TTL value is a 32-bit unsigned integer that represents the TTL in seconds.",
	},
	"nxdomain_rulesets": schema.ListAttribute{
		ElementType: types.StringType,
		Optional:    true,
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
		},
		MarkdownDescription: "The names of the Ruleset objects assigned at the grid level for NXDOMAIN redirection.",
	},
	"recursion": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Determines if recursion is enabled or not.",
	},
	"response_rate_limiting": schema.SingleNestedAttribute{
		Attributes:          ViewResponseRateLimitingResourceSchemaAttributes,
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The response rate limiting settings for the DNS view. This feature is used to limit the number of responses sent to a client in a given time period.",
	},
	"root_name_server_type": schema.StringAttribute{
		Default: stringdefault.StaticString("INTERNET"),
		Validators: []validator.String{
			stringvalidator.OneOf("CUSTOM", "INTERNET"),
		},
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Determines the type of root name servers.",
	},
	"rpz_drop_ip_rule_enabled": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Enables the appliance to ignore RPZ-IP triggers with prefix lengths less than the specified minimum prefix length.",
	},
	"rpz_drop_ip_rule_min_prefix_length_ipv4": schema.Int64Attribute{
		Optional:            true,
		Computed:            true,
		Default:             int64default.StaticInt64(29),
		MarkdownDescription: "The minimum prefix length for IPv4 RPZ-IP triggers. The appliance ignores RPZ-IP triggers with prefix lengths less than the specified minimum IPv4 prefix length.",
	},
	"rpz_drop_ip_rule_min_prefix_length_ipv6": schema.Int64Attribute{
		Optional:            true,
		Computed:            true,
		Default:             int64default.StaticInt64(112),
		MarkdownDescription: "The minimum prefix length for IPv6 RPZ-IP triggers. The appliance ignores RPZ-IP triggers with prefix lengths less than the specified minimum IPv6 prefix length.",
	},
	"rpz_qname_wait_recurse": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "The flag that indicates whether recursive RPZ lookups are enabled.",
	},
	"scavenging_settings": schema.SingleNestedAttribute{
		Attributes:          ViewScavengingSettingsResourceSchemaAttributes,
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Scavenging settings for the DNS view",
	},
	"sortlist": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: ViewSortlistResourceSchemaAttributes,
		},
		Optional: true,
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
		},
		MarkdownDescription: "A sort list that determines the order of IP addresses in responses sent to DNS queries.",
	},
}

var ViewResourceUddiSchemaAttributes = map[string]schema.Attribute{
	"add_edns_option_in_outgoing_query": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "_add_edns_option_in_outgoing_query_ adds client IP, MAC address and view name into outgoing recursive query. Defaults to _false_.",
	},
	"comment": schema.StringAttribute{
		Default:             stringdefault.StaticString(""),
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Optional. Comment for view.",
	},
	"compartment_id": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The access view associated with the object. If no access view is associated with the object, the value defaults to empty.",
	},
	"custom_root_ns": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: RootNSResourceSchemaAttributes,
		},
		Optional: true,
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
		},
		MarkdownDescription: "Optional. List of custom root nameservers. The order does not matter.  Error if empty while _custom_root_ns_enabled_ is _true_. Error if there are duplicate items in the list.  Defaults to empty.",
	},
	"custom_root_ns_enabled": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Optional. _true_ to use custom root nameservers instead of the default ones.  The _custom_root_ns_ is validated when enabled.  Defaults to _false_.",
	},
	"disabled": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Optional. _true_ to disable object. A disabled object is effectively non-existent when generating configuration.",
	},
	"dnssec_enable_validation": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(true),
		MarkdownDescription: "Optional. _true_ to perform DNSSEC validation. Ignored if _dnssec_enabled_ is _false_.  Defaults to _true_.",
	},
	"dnssec_enabled": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(true),
		MarkdownDescription: "Optional. Master toggle for all DNSSEC processing. Other _dnssec_*_ configuration is unused if this is disabled.  Defaults to _true_.",
	},
	"dnssec_trust_anchors": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: TrustAnchorResourceSchemaAttributes,
		},
		Optional: true,
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
		},
		MarkdownDescription: "Optional. DNSSEC trust anchors.  Error if there are list items with duplicate (_zone_, _sep_, _algorithm_) combinations.  Defaults to empty.",
	},
	"dnssec_validate_expiry": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(true),
		MarkdownDescription: "Optional. _true_ to reject expired DNSSEC keys. Ignored if either _dnssec_enabled_ or _dnssec_enable_validation_ is _false_.  Defaults to _true_.",
	},
	"dtc_config": schema.SingleNestedAttribute{
		Attributes:          DTCConfigResourceSchemaAttributes,
		Optional:            true,
		Computed:            true,
		Default:             objectdefault.StaticValue(types.ObjectValueMust(DTCConfigAttrTypes, map[string]attr.Value{"default_ttl": types.Int64Value(300)})),
		MarkdownDescription: "Construct for fields: _default_ttl_.",
	},
	"ecs_enabled": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Optional. _true_ to enable EDNS client subnet for recursive queries. Other _ecs_*_ fields are ignored if this field is not enabled.  Defaults to _false-.",
	},
	"ecs_forwarding": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Optional. _true_ to enable ECS options in outbound queries. This functionality has additional overhead so it is disabled by default.  Defaults to _false_.",
	},
	"ecs_prefix_v4": schema.Int64Attribute{
		Optional:            true,
		Computed:            true,
		Default:             int64default.StaticInt64(24),
		MarkdownDescription: "Optional. Maximum scope length for v4 ECS.  Unsigned integer, min 1 max 24  Defaults to 24.",
	},
	"ecs_prefix_v6": schema.Int64Attribute{
		Optional:            true,
		Computed:            true,
		Default:             int64default.StaticInt64(56),
		MarkdownDescription: "Optional. Maximum scope length for v6 ECS.  Unsigned integer, min 1 max 56  Defaults to 56.",
	},
	"ecs_zones": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: ECSZoneResourceSchemaAttributes,
		},
		Optional: true,
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
		},
		MarkdownDescription: "Optional. List of zones where ECS queries may be sent.  Error if empty while _ecs_enabled_ is _true_. Error if there are duplicate FQDNs in the list.  Defaults to empty.",
	},
	"edns_udp_size": schema.Int64Attribute{
		Optional:            true,
		Computed:            true,
		Default:             int64default.StaticInt64(1232),
		MarkdownDescription: "Optional. _edns_udp_size_ represents the edns UDP size. The size a querying DNS server advertises to the DNS server it’s sending a query to.  Defaults to 1232 bytes.",
	},
	"filter_aaaa_acl": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: ACLItemResourceSchemaAttributes,
		},
		Optional: true,
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
		},
		MarkdownDescription: "Optional. Specifies a list of client addresses for which AAAA filtering is to be applied.  Defaults to _empty_.",
	},
	"filter_aaaa_on_v4": schema.StringAttribute{
		Default:             stringdefault.StaticString("no"),
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "_filter_aaaa_on_v4_ allows named to omit some IPv6 addresses when responding to IPv4 clients.  Allowed values: * _yes_, * _no_, * _break_dnssec_.  Defaults to _no_",
	},
	"forwarders": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: ForwarderResourceSchemaAttributes(false),
		},
		Optional: true,
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
		},
		MarkdownDescription: "Optional. List of forwarders.  Error if empty while _forwarders_only_ or _use_root_forwarders_for_local_resolution_with_b1td_ is _true_. Error if there are items in the list with duplicate addresses.  Defaults to empty.",
	},
	"forwarders_only": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Optional. _true_ to only forward.  Defaults to _false_.",
	},
	"gss_tsig_enabled": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "_gss_tsig_enabled_ enables/disables GSS-TSIG signed dynamic updates.  Defaults to _false_.",
	},
	"inheritance_sources": schema.SingleNestedAttribute{
		Attributes: ViewInheritanceResourceSchemaAttributes,
		Optional:   true,
		Computed:   true,
		PlanModifiers: []planmodifier.Object{
			objectplanmodifier.UseStateForUnknown(),
		},
		MarkdownDescription: "Inheritance configuration specifies how and which fields _View_ object inherits from [ _Global_, _Server_ ] parent.",
	},
	"ip_spaces": schema.ListAttribute{
		ElementType: types.StringType,
		Optional:    true,
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
			listvalidator.SizeAtMost(1),
		},
		MarkdownDescription: "The resource identifier.",
	},
	"lame_ttl": schema.Int64Attribute{
		Optional:            true,
		Computed:            true,
		Default:             int64default.StaticInt64(600),
		MarkdownDescription: "Optional. Unused in the current on-prem DNS server implementation.  Unsigned integer, min 0 max 3600 (1h).  Defaults to 600.",
	},
	"match_clients_acl": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: ACLItemResourceSchemaAttributes,
		},
		Optional: true,
		Computed: true,
		PlanModifiers: []planmodifier.List{
			planmod.UseDefaultAclForNull(),
		},
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
		},
		MarkdownDescription: "Optional. Specifies which clients have access to the view.  Defaults to empty.",
	},
	"match_destinations_acl": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: ACLItemResourceSchemaAttributes,
		},
		Optional: true,
		Computed: true,
		PlanModifiers: []planmodifier.List{
			planmod.UseDefaultAclForNull(),
		},
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
		},
		MarkdownDescription: "Optional. Specifies which destination addresses have access to the view.  Defaults to empty.",
	},
	"match_recursive_only": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Optional. If _true_ only recursive queries from matching clients access the view.  Defaults to _false_.",
	},
	"max_cache_ttl": schema.Int64Attribute{
		Optional:            true,
		Computed:            true,
		Default:             int64default.StaticInt64(604800),
		MarkdownDescription: "Optional. Seconds to cache positive responses.  Unsigned integer, min 1 max 604800 (7d).  Defaults to 604800 (7d).",
	},
	"max_negative_ttl": schema.Int64Attribute{
		Optional:            true,
		Computed:            true,
		Default:             int64default.StaticInt64(10800),
		MarkdownDescription: "Optional. Seconds to cache negative responses.  Unsigned integer, min 1 max 604800 (7d).  Defaults to 10800 (3h).",
	},
	"max_udp_size": schema.Int64Attribute{
		Optional:            true,
		Computed:            true,
		Default:             int64default.StaticInt64(1232),
		MarkdownDescription: "Optional. _max_udp_size_ represents maximum UDP payload size. The maximum number of bytes a responding DNS server will send to a UDP datagram.  Defaults to 1232 bytes.",
	},
	"minimal_responses": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Optional. When enabled, the DNS server will only add records to the authority and additional data sections when they are required.  Defaults to _false_.",
	},
	"name": schema.StringAttribute{
		Required:            true,
		MarkdownDescription: "Name of view.",
	},
	"notify": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "_notify_ all external secondary DNS servers.  Defaults to _false_.",
	},
	"query_acl": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: ACLItemResourceSchemaAttributes,
		},
		Optional: true,
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
		},
		MarkdownDescription: "Optional. Clients must match this ACL to make authoritative queries. Also used for recursive queries if that ACL is unset.  Defaults to empty.",
	},
	"recursion_acl": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: ACLItemResourceSchemaAttributes,
		},
		Optional: true,
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
		},
		MarkdownDescription: "Optional. Clients must match this ACL to make recursive queries. If this ACL is empty, then the _query_acl_ will be used instead.  Defaults to empty.",
	},
	"recursion_enabled": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(true),
		MarkdownDescription: "Optional. _true_ to allow recursive DNS queries.  Defaults to _true_.",
	},
	"sort_list": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: SortListItemResourceSchemaAttributes,
		},
		Optional: true,
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
		},
		MarkdownDescription: "Optional. Specifies a sorted network list for A/AAAA records in DNS query response.  Defaults to _empty_.",
	},
	"synthesize_address_records_from_https": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "_synthesize_address_records_from_https_ enables/disables creation of A/AAAA records from HTTPS RR Defaults to _false_.",
	},
	"tags": schema.MapAttribute{
		Optional:    true,
		Computed:    true,
		ElementType: types.StringType,
		Default:     mapdefault.StaticValue(types.MapNull(types.StringType)),
		Validators: []validator.Map{
			mapvalidator.SizeAtLeast(1),
		},
		MarkdownDescription: "Tagging specifics.",
	},
	"tags_all": schema.MapAttribute{
		Computed:            true,
		ElementType:         types.StringType,
		MarkdownDescription: "All tags including inherited values.",
	},
	"transfer_acl": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: ACLItemResourceSchemaAttributes,
		},
		Optional: true,
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
		},
		MarkdownDescription: "Optional. Clients must match this ACL to receive zone transfers.  Defaults to empty.",
	},
	"update_acl": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: ACLItemResourceSchemaAttributes,
		},
		Optional: true,
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
		},
		MarkdownDescription: "Optional. Specifies which hosts are allowed to issue Dynamic DNS updates for authoritative zones of _primary_type_ _cloud_.  Defaults to empty.",
	},
	"use_forwarders_for_subzones": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(true),
		MarkdownDescription: "Optional. Use default forwarders to resolve queries for subzones.  Defaults to _true_.",
	},
	"use_root_forwarders_for_local_resolution_with_b1td": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "_use_root_forwarders_for_local_resolution_with_b1td_ allows DNS recursive queries sent to root forwarders for local resolution when deployed alongside BloxOne Thread Defense. Defaults to _false_.",
	},
	"zone_authority": schema.SingleNestedAttribute{
		Attributes: ZoneAuthorityResourceSchemaAttributes,
		Optional:   true,
		Computed:   true,
		Default: objectdefault.StaticValue(types.ObjectValueMust(ZoneAuthorityAttrTypes, map[string]attr.Value{
			"default_ttl":       types.Int64Value(28800),
			"expire":            types.Int64Value(2.4192e+06),
			"mname":             types.StringValue("ns.b1ddi"),
			"negative_ttl":      types.Int64Value(900),
			"protocol_mname":    types.StringValue("ns.b1ddi"),
			"protocol_rname":    types.StringValue("hostmaster"),
			"refresh":           types.Int64Value(10800),
			"retry":             types.Int64Value(3600),
			"rname":             types.StringValue("hostmaster"),
			"use_default_mname": types.BoolValue(true),
		})),
		PlanModifiers: []planmodifier.Object{
			objectplanmodifier.UseStateForUnknown(),
		},
		MarkdownDescription: "Construct for fields: _refresh_, _retry_, _expire_, _default_ttl_, _negative_ttl_, _rname_, _protocol_rname_, _mname_, _protocol_mname_, _use_default_mname_.",
	},
}

// Expand converts the TF model to the infoblox core model
func (m *ViewModel) Expand(ctx context.Context, diags *diag.Diagnostics, isCreate bool) *coremodel.View {
	if m == nil {
		return nil
	}

	obj := &coremodel.View{}

	// Expand NIOS nested attribute (returns nil if not present)
	niosModel := flex.ExpandNestedObject[NIOSViewModel](ctx, m.NIOS, diags)
	if niosModel != nil {
		obj.NIOS = niosModel.Expand(ctx, diags)
	}

	// Expand UDDI nested attribute (returns nil if not present)
	uddiModel := flex.ExpandNestedObject[UDDIViewModel](ctx, m.UDDI, diags)
	if uddiModel != nil {
		obj.UDDI = uddiModel.Expand(ctx, diags)
	}

	return obj
}

// Expand converts the NIOS TF model to the core model.
func (m *NIOSViewModel) Expand(ctx context.Context, diags *diag.Diagnostics) *coremodel.NIOSViewExt {
	return &coremodel.NIOSViewExt{
		BlacklistAction:                  flex.ExpandStringPointerNullAsEmpty(m.BlacklistAction),
		BlacklistLogQuery:                flex.ExpandBoolPointer(m.BlacklistLogQuery),
		BlacklistRedirectAddresses:       flex.ExpandFrameworkListString(ctx, m.BlacklistRedirectAddresses, diags),
		BlacklistRedirectTtl:             flex.ExpandInt64Pointer(m.BlacklistRedirectTtl),
		BlacklistRulesets:                flex.ExpandFrameworkListString(ctx, m.BlacklistRulesets, diags),
		Comment:                          flex.ExpandStringPointerNullAsEmpty(m.Comment),
		CustomRootNameServers:            flex.ExpandFrameworkListNestedBlock(ctx, m.CustomRootNameServers, diags, ExpandViewCustomRootNameServers),
		DdnsForceCreationTimestampUpdate: flex.ExpandBoolPointer(m.DdnsForceCreationTimestampUpdate),
		DdnsPrincipalGroup:               flex.ExpandStringPointer(m.DdnsPrincipalGroup),
		DdnsPrincipalTracking:            flex.ExpandBoolPointer(m.DdnsPrincipalTracking),
		DdnsRestrictPatterns:             flex.ExpandBoolPointer(m.DdnsRestrictPatterns),
		DdnsRestrictPatternsList:         flex.ExpandFrameworkListString(ctx, m.DdnsRestrictPatternsList, diags),
		DdnsRestrictProtected:            flex.ExpandBoolPointer(m.DdnsRestrictProtected),
		DdnsRestrictSecure:               flex.ExpandBoolPointer(m.DdnsRestrictSecure),
		DdnsRestrictStatic:               flex.ExpandBoolPointer(m.DdnsRestrictStatic),
		Disable:                          flex.ExpandBoolPointer(m.Disable),
		Dns64Enabled:                     flex.ExpandBoolPointer(m.Dns64Enabled),
		Dns64Groups:                      flex.ExpandFrameworkListString(ctx, m.Dns64Groups, diags),
		DnssecEnabled:                    flex.ExpandBoolPointer(m.DnssecEnabled),
		DnssecExpiredSignaturesEnabled:   flex.ExpandBoolPointer(m.DnssecExpiredSignaturesEnabled),
		DnssecNegativeTrustAnchors:       flex.ExpandFrameworkListString(ctx, m.DnssecNegativeTrustAnchors, diags),
		DnssecTrustedKeys:                flex.ExpandFrameworkListNestedBlock(ctx, m.DnssecTrustedKeys, diags, ExpandViewDnssecTrustedKeys),
		DnssecValidationEnabled:          flex.ExpandBoolPointer(m.DnssecValidationEnabled),
		EdnsUdpSize:                      flex.ExpandInt64Pointer(m.EdnsUdpSize),
		EnableBlacklist:                  flex.ExpandBoolPointer(m.EnableBlacklist),
		EnableFixedRrsetOrderFqdns:       flex.ExpandBoolPointer(m.EnableFixedRrsetOrderFqdns),
		EnableMatchRecursiveOnly:         flex.ExpandBoolPointer(m.EnableMatchRecursiveOnly),
		ExtAttrs:                         flex.ExpandMapStringAny(ctx, m.ExtAttrs, diags),
		FilterAaaa:                       flex.ExpandStringPointerNullAsEmpty(m.FilterAaaa),
		FilterAaaaList:                   flex.ExpandFrameworkListNestedBlock(ctx, m.FilterAaaaList, diags, ExpandViewFilterAaaaList),
		FixedRrsetOrderFqdns:             flex.ExpandFrameworkListNestedBlock(ctx, m.FixedRrsetOrderFqdns, diags, ExpandViewFixedRrsetOrderFqdns),
		ForwardOnly:                      flex.ExpandBoolPointer(m.ForwardOnly),
		Forwarders:                       flex.ExpandFrameworkListString(ctx, m.Forwarders, diags),
		LastQueriedAcl:                   flex.ExpandFrameworkListNestedBlock(ctx, m.LastQueriedAcl, diags, ExpandViewLastQueriedAcl),
		MatchClients:                     flex.ExpandFrameworkListNestedBlock(ctx, m.MatchClients, diags, ExpandViewMatchClients),
		MatchDestinations:                flex.ExpandFrameworkListNestedBlock(ctx, m.MatchDestinations, diags, ExpandViewMatchDestinations),
		MaxCacheTtl:                      flex.ExpandInt64Pointer(m.MaxCacheTtl),
		MaxNcacheTtl:                     flex.ExpandInt64Pointer(m.MaxNcacheTtl),
		MaxUdpSize:                       flex.ExpandInt64Pointer(m.MaxUdpSize),
		Name:                             flex.ExpandStringPointerNullAsEmpty(m.Name),
		NetworkView:                      flex.ExpandStringPointerNullAsEmpty(m.NetworkView),
		NotifyDelay:                      flex.ExpandInt64Pointer(m.NotifyDelay),
		NxdomainLogQuery:                 flex.ExpandBoolPointer(m.NxdomainLogQuery),
		NxdomainRedirect:                 flex.ExpandBoolPointer(m.NxdomainRedirect),
		NxdomainRedirectAddresses:        flex.ExpandFrameworkListString(ctx, m.NxdomainRedirectAddresses, diags),
		NxdomainRedirectAddressesV6:      flex.ExpandFrameworkListString(ctx, m.NxdomainRedirectAddressesV6, diags),
		NxdomainRedirectTtl:              flex.ExpandInt64Pointer(m.NxdomainRedirectTtl),
		NxdomainRulesets:                 flex.ExpandFrameworkListString(ctx, m.NxdomainRulesets, diags),
		Recursion:                        flex.ExpandBoolPointer(m.Recursion),
		ResponseRateLimiting:             ExpandViewResponseRateLimiting(ctx, m.ResponseRateLimiting, diags),
		RootNameServerType:               flex.ExpandStringPointerNullAsEmpty(m.RootNameServerType),
		RpzDropIpRuleEnabled:             flex.ExpandBoolPointer(m.RpzDropIpRuleEnabled),
		RpzDropIpRuleMinPrefixLengthIpv4: flex.ExpandInt64Pointer(m.RpzDropIpRuleMinPrefixLengthIpv4),
		RpzDropIpRuleMinPrefixLengthIpv6: flex.ExpandInt64Pointer(m.RpzDropIpRuleMinPrefixLengthIpv6),
		RpzQnameWaitRecurse:              flex.ExpandBoolPointer(m.RpzQnameWaitRecurse),
		ScavengingSettings:               ExpandViewScavengingSettings(ctx, m.ScavengingSettings, diags),
		Sortlist:                         flex.ExpandFrameworkListNestedBlock(ctx, m.Sortlist, diags, ExpandViewSortlist),
	}
}

// ApplyViewNIOSUseFlags derives NIOS use flags from the raw config
// value(s) and writes them onto the core model. A flag is true when the user
// set any of its governed value fields in config.
func ApplyViewNIOSUseFlags(ctx context.Context, config tfsdk.Config, obj *coremodel.View, diags *diag.Diagnostics) {
	if obj == nil || obj.NIOS == nil {
		return
	}
	obj.NIOS.UseBlacklist = flex.DeriveUseFlag(ctx, config, diags, path.Root("nios").AtName("blacklist_action"), path.Root("nios").AtName("blacklist_log_query"), path.Root("nios").AtName("blacklist_redirect_addresses"), path.Root("nios").AtName("blacklist_redirect_ttl"), path.Root("nios").AtName("blacklist_rulesets"), path.Root("nios").AtName("enable_blacklist"))
	obj.NIOS.UseDdnsForceCreationTimestampUpdate = flex.DeriveUseFlag(ctx, config, diags, path.Root("nios").AtName("ddns_force_creation_timestamp_update"))
	obj.NIOS.UseDdnsPatternsRestriction = flex.DeriveUseFlag(ctx, config, diags, path.Root("nios").AtName("ddns_restrict_patterns_list"), path.Root("nios").AtName("ddns_restrict_patterns"))
	obj.NIOS.UseDdnsPrincipalSecurity = flex.DeriveUseFlag(ctx, config, diags, path.Root("nios").AtName("ddns_restrict_secure"), path.Root("nios").AtName("ddns_principal_tracking"), path.Root("nios").AtName("ddns_principal_group"))
	obj.NIOS.UseDdnsRestrictProtected = flex.DeriveUseFlag(ctx, config, diags, path.Root("nios").AtName("ddns_restrict_protected"))
	obj.NIOS.UseDdnsRestrictStatic = flex.DeriveUseFlag(ctx, config, diags, path.Root("nios").AtName("ddns_restrict_static"))
	obj.NIOS.UseDns64 = flex.DeriveUseFlag(ctx, config, diags, path.Root("nios").AtName("dns64_enabled"), path.Root("nios").AtName("dns64_groups"))
	obj.NIOS.UseDnssec = flex.DeriveUseFlag(ctx, config, diags, path.Root("nios").AtName("dnssec_enabled"), path.Root("nios").AtName("dnssec_expired_signatures_enabled"), path.Root("nios").AtName("dnssec_validation_enabled"), path.Root("nios").AtName("dnssec_trusted_keys"))
	obj.NIOS.UseEdnsUdpSize = flex.DeriveUseFlag(ctx, config, diags, path.Root("nios").AtName("edns_udp_size"))
	obj.NIOS.UseFilterAaaa = flex.DeriveUseFlag(ctx, config, diags, path.Root("nios").AtName("filter_aaaa"), path.Root("nios").AtName("filter_aaaa_list"))
	obj.NIOS.UseFixedRrsetOrderFqdns = flex.DeriveUseFlag(ctx, config, diags, path.Root("nios").AtName("fixed_rrset_order_fqdns"), path.Root("nios").AtName("enable_fixed_rrset_order_fqdns"))
	obj.NIOS.UseForwarders = flex.DeriveUseFlag(ctx, config, diags, path.Root("nios").AtName("forwarders"), path.Root("nios").AtName("forward_only"))
	obj.NIOS.UseMaxCacheTtl = flex.DeriveUseFlag(ctx, config, diags, path.Root("nios").AtName("max_cache_ttl"))
	obj.NIOS.UseMaxNcacheTtl = flex.DeriveUseFlag(ctx, config, diags, path.Root("nios").AtName("max_ncache_ttl"))
	obj.NIOS.UseMaxUdpSize = flex.DeriveUseFlag(ctx, config, diags, path.Root("nios").AtName("max_udp_size"))
	obj.NIOS.UseNxdomainRedirect = flex.DeriveUseFlag(ctx, config, diags, path.Root("nios").AtName("nxdomain_redirect"), path.Root("nios").AtName("nxdomain_redirect_addresses"), path.Root("nios").AtName("nxdomain_redirect_addresses_v6"), path.Root("nios").AtName("nxdomain_redirect_ttl"), path.Root("nios").AtName("nxdomain_log_query"), path.Root("nios").AtName("nxdomain_rulesets"))
	obj.NIOS.UseRecursion = flex.DeriveUseFlag(ctx, config, diags, path.Root("nios").AtName("recursion"))
	obj.NIOS.UseResponseRateLimiting = flex.DeriveUseFlag(ctx, config, diags, path.Root("nios").AtName("response_rate_limiting"))
	obj.NIOS.UseRootNameServer = flex.DeriveUseFlag(ctx, config, diags, path.Root("nios").AtName("custom_root_name_servers"), path.Root("nios").AtName("root_name_server_type"))
	obj.NIOS.UseRpzDropIpRule = flex.DeriveUseFlag(ctx, config, diags, path.Root("nios").AtName("rpz_drop_ip_rule_enabled"), path.Root("nios").AtName("rpz_drop_ip_rule_min_prefix_length_ipv4"), path.Root("nios").AtName("rpz_drop_ip_rule_min_prefix_length_ipv6"))
	obj.NIOS.UseRpzQnameWaitRecurse = flex.DeriveUseFlag(ctx, config, diags, path.Root("nios").AtName("rpz_qname_wait_recurse"))
	obj.NIOS.UseScavengingSettings = flex.DeriveUseFlag(ctx, config, diags, path.Root("nios").AtName("scavenging_settings"), path.Root("nios").AtName("last_queried_acl"))
	obj.NIOS.UseSortlist = flex.DeriveUseFlag(ctx, config, diags, path.Root("nios").AtName("sortlist"))
}

// Expand converts the UDDI TF model to the core model.
func (m *UDDIViewModel) Expand(ctx context.Context, diags *diag.Diagnostics) *coremodel.UDDIViewExt {
	return &coremodel.UDDIViewExt{
		AddEdnsOptionInOutgoingQuery:      flex.ExpandBoolPointer(m.AddEdnsOptionInOutgoingQuery),
		Comment:                           flex.ExpandStringPointer(m.Comment),
		CompartmentId:                     flex.ExpandStringPointer(m.CompartmentId),
		CustomRootNs:                      flex.ExpandFrameworkListNestedBlock(ctx, m.CustomRootNs, diags, ExpandRootNS),
		CustomRootNsEnabled:               flex.ExpandBoolPointer(m.CustomRootNsEnabled),
		Disabled:                          flex.ExpandBoolPointer(m.Disabled),
		DnssecEnableValidation:            flex.ExpandBoolPointer(m.DnssecEnableValidation),
		DnssecEnabled:                     flex.ExpandBoolPointer(m.DnssecEnabled),
		DnssecTrustAnchors:                flex.ExpandFrameworkListNestedBlock(ctx, m.DnssecTrustAnchors, diags, ExpandTrustAnchor),
		DnssecValidateExpiry:              flex.ExpandBoolPointer(m.DnssecValidateExpiry),
		DtcConfig:                         ExpandDTCConfig(ctx, m.DtcConfig, diags),
		EcsEnabled:                        flex.ExpandBoolPointer(m.EcsEnabled),
		EcsForwarding:                     flex.ExpandBoolPointer(m.EcsForwarding),
		EcsPrefixV4:                       flex.ExpandInt64Pointer(m.EcsPrefixV4),
		EcsPrefixV6:                       flex.ExpandInt64Pointer(m.EcsPrefixV6),
		EcsZones:                          flex.ExpandFrameworkListNestedBlock(ctx, m.EcsZones, diags, ExpandECSZone),
		EdnsUdpSize:                       flex.ExpandInt64Pointer(m.EdnsUdpSize),
		FilterAaaaAcl:                     flex.ExpandFrameworkListNestedBlock(ctx, m.FilterAaaaAcl, diags, ExpandACLItem),
		FilterAaaaOnV4:                    flex.ExpandStringPointer(m.FilterAaaaOnV4),
		Forwarders:                        flex.ExpandFrameworkListNestedBlock(ctx, m.Forwarders, diags, ExpandForwarder),
		ForwardersOnly:                    flex.ExpandBoolPointer(m.ForwardersOnly),
		GssTsigEnabled:                    flex.ExpandBoolPointer(m.GssTsigEnabled),
		InheritanceSources:                ExpandViewInheritance(ctx, m.InheritanceSources, diags),
		IpSpaces:                          flex.ExpandFrameworkListString(ctx, m.IpSpaces, diags),
		LameTtl:                           flex.ExpandInt64Pointer(m.LameTtl),
		MatchClientsAcl:                   flex.ExpandFrameworkListNestedBlock(ctx, m.MatchClientsAcl, diags, ExpandACLItem),
		MatchDestinationsAcl:              flex.ExpandFrameworkListNestedBlock(ctx, m.MatchDestinationsAcl, diags, ExpandACLItem),
		MatchRecursiveOnly:                flex.ExpandBoolPointer(m.MatchRecursiveOnly),
		MaxCacheTtl:                       flex.ExpandInt64Pointer(m.MaxCacheTtl),
		MaxNegativeTtl:                    flex.ExpandInt64Pointer(m.MaxNegativeTtl),
		MaxUdpSize:                        flex.ExpandInt64Pointer(m.MaxUdpSize),
		MinimalResponses:                  flex.ExpandBoolPointer(m.MinimalResponses),
		Name:                              flex.ExpandString(m.Name),
		Notify:                            flex.ExpandBoolPointer(m.Notify),
		QueryAcl:                          flex.ExpandFrameworkListNestedBlock(ctx, m.QueryAcl, diags, ExpandACLItem),
		RecursionAcl:                      flex.ExpandFrameworkListNestedBlock(ctx, m.RecursionAcl, diags, ExpandACLItem),
		RecursionEnabled:                  flex.ExpandBoolPointer(m.RecursionEnabled),
		SortList:                          flex.ExpandFrameworkListNestedBlock(ctx, m.SortList, diags, ExpandSortListItem),
		SynthesizeAddressRecordsFromHttps: flex.ExpandBoolPointer(m.SynthesizeAddressRecordsFromHttps),
		Tags:                              flex.ExpandMapStringAny(ctx, m.Tags, diags),
		TransferAcl:                       flex.ExpandFrameworkListNestedBlock(ctx, m.TransferAcl, diags, ExpandACLItem),
		UpdateAcl:                         flex.ExpandFrameworkListNestedBlock(ctx, m.UpdateAcl, diags, ExpandACLItem),
		UseForwardersForSubzones:          flex.ExpandBoolPointer(m.UseForwardersForSubzones),
		UseRootForwardersForLocalResolutionWithB1td: flex.ExpandBoolPointer(m.UseRootForwardersForLocalResolutionWithB1td),
		ZoneAuthority: ExpandZoneAuthority(ctx, m.ZoneAuthority, diags),
	}
}

// Flatten populates the TF model from a core response.
func (m *ViewModel) Flatten(ctx context.Context, resp *coremodel.View, diags *diag.Diagnostics) {
	if resp == nil {
		return
	}

	m.Id = flex.FlattenStringPointer(resp.Id)

	// Extract existing NIOS model, flatten API response onto it, convert back
	niosModel := flex.ExpandNestedObject[NIOSViewModel](ctx, m.NIOS, diags)
	if niosModel == nil {
		niosModel = &NIOSViewModel{}
	}
	niosModel.Flatten(ctx, resp.NIOS, diags)
	if resp.NIOS != nil {
		m.NIOS = flex.FlattenNestedObject(ctx, niosModel, NIOSViewAttrTypes, diags)
	} else {
		m.NIOS = types.ObjectNull(NIOSViewAttrTypes)
	}

	// Extract existing UDDI model, flatten API response onto it, convert back
	uddiModel := flex.ExpandNestedObject[UDDIViewModel](ctx, m.UDDI, diags)
	if uddiModel == nil {
		uddiModel = &UDDIViewModel{}
	}
	uddiModel.Flatten(ctx, resp.UDDI, diags)
	if resp.UDDI != nil {
		m.UDDI = flex.FlattenNestedObject(ctx, uddiModel, UDDIViewAttrTypes, diags)
	} else {
		m.UDDI = types.ObjectNull(UDDIViewAttrTypes)
	}
}

// Flatten merges API response onto existing NIOS model.
func (m *NIOSViewModel) Flatten(ctx context.Context, from *coremodel.NIOSViewExt, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	planExtAttrs := m.ExtAttrs
	if planExtAttrs.IsUnknown() {
		planExtAttrs = types.MapNull(types.StringType)
	}
	m.BlacklistAction = flex.FlattenStringPointerEmptyAsNull(from.BlacklistAction)
	m.BlacklistLogQuery = flex.FlattenBoolPointer(from.BlacklistLogQuery)
	m.BlacklistRedirectAddresses = flex.FlattenFrameworkListString(ctx, from.BlacklistRedirectAddresses, diags)
	m.BlacklistRedirectTtl = flex.FlattenInt64Pointer(from.BlacklistRedirectTtl)
	m.BlacklistRulesets = flex.FlattenFrameworkListString(ctx, from.BlacklistRulesets, diags)
	m.Comment = flex.FlattenStringPointerEmptyAsNull(from.Comment)
	m.CustomRootNameServers = flex.FlattenFrameworkListNestedBlock(ctx, from.CustomRootNameServers, ViewCustomRootNameServersAttrTypes, diags, FlattenViewCustomRootNameServers)
	m.DdnsForceCreationTimestampUpdate = flex.FlattenBoolPointer(from.DdnsForceCreationTimestampUpdate)
	m.DdnsPrincipalGroup = flex.FlattenStringPointerEmptyAsNull(from.DdnsPrincipalGroup)
	m.DdnsPrincipalTracking = flex.FlattenBoolPointer(from.DdnsPrincipalTracking)
	m.DdnsRestrictPatterns = flex.FlattenBoolPointer(from.DdnsRestrictPatterns)
	m.DdnsRestrictPatternsList = flex.FlattenFrameworkListString(ctx, from.DdnsRestrictPatternsList, diags)
	m.DdnsRestrictProtected = flex.FlattenBoolPointer(from.DdnsRestrictProtected)
	m.DdnsRestrictSecure = flex.FlattenBoolPointer(from.DdnsRestrictSecure)
	m.DdnsRestrictStatic = flex.FlattenBoolPointer(from.DdnsRestrictStatic)
	m.Disable = flex.FlattenBoolPointer(from.Disable)
	m.Dns64Enabled = flex.FlattenBoolPointer(from.Dns64Enabled)
	m.Dns64Groups = flex.FlattenFrameworkUnorderedListString(ctx, from.Dns64Groups, diags)
	m.DnssecEnabled = flex.FlattenBoolPointer(from.DnssecEnabled)
	m.DnssecExpiredSignaturesEnabled = flex.FlattenBoolPointer(from.DnssecExpiredSignaturesEnabled)
	m.DnssecNegativeTrustAnchors = flex.FlattenFrameworkListString(ctx, from.DnssecNegativeTrustAnchors, diags)
	m.DnssecTrustedKeys = flex.FlattenFrameworkListNestedBlock(ctx, from.DnssecTrustedKeys, ViewDnssecTrustedKeysAttrTypes, diags, FlattenViewDnssecTrustedKeys)
	m.DnssecValidationEnabled = flex.FlattenBoolPointer(from.DnssecValidationEnabled)
	m.EdnsUdpSize = flex.FlattenInt64Pointer(from.EdnsUdpSize)
	m.EnableBlacklist = flex.FlattenBoolPointer(from.EnableBlacklist)
	m.EnableFixedRrsetOrderFqdns = flex.FlattenBoolPointer(from.EnableFixedRrsetOrderFqdns)
	m.EnableMatchRecursiveOnly = flex.FlattenBoolPointer(from.EnableMatchRecursiveOnly)
	m.ExtAttrs, m.ExtAttrsAll = flex.FlattenEAs(planExtAttrs, from.ExtAttrs)
	m.FilterAaaa = flex.FlattenStringPointerEmptyAsNull(from.FilterAaaa)
	m.FilterAaaaList = flex.FlattenFrameworkListNestedBlock(ctx, from.FilterAaaaList, ViewFilterAaaaListAttrTypes, diags, FlattenViewFilterAaaaList)
	m.FixedRrsetOrderFqdns = flex.FlattenFrameworkListNestedBlock(ctx, from.FixedRrsetOrderFqdns, ViewFixedRrsetOrderFqdnsAttrTypes, diags, FlattenViewFixedRrsetOrderFqdns)
	m.ForwardOnly = flex.FlattenBoolPointer(from.ForwardOnly)
	m.Forwarders = flex.FlattenFrameworkListString(ctx, from.Forwarders, diags)
	m.LastQueriedAcl = flex.FlattenFrameworkListNestedBlock(ctx, from.LastQueriedAcl, ViewLastQueriedAclAttrTypes, diags, FlattenViewLastQueriedAcl)
	m.MatchClients = flex.FlattenFrameworkListNestedBlock(ctx, from.MatchClients, ViewMatchClientsAttrTypes, diags, FlattenViewMatchClients)
	m.MatchDestinations = flex.FlattenFrameworkListNestedBlock(ctx, from.MatchDestinations, ViewMatchDestinationsAttrTypes, diags, FlattenViewMatchDestinations)
	m.MaxCacheTtl = flex.FlattenInt64Pointer(from.MaxCacheTtl)
	m.MaxNcacheTtl = flex.FlattenInt64Pointer(from.MaxNcacheTtl)
	m.MaxUdpSize = flex.FlattenInt64Pointer(from.MaxUdpSize)
	m.Name = flex.FlattenStringPointerEmptyAsNull(from.Name)
	m.NetworkView = flex.FlattenStringPointerEmptyAsNull(from.NetworkView)
	m.NotifyDelay = flex.FlattenInt64Pointer(from.NotifyDelay)
	m.NxdomainLogQuery = flex.FlattenBoolPointer(from.NxdomainLogQuery)
	m.NxdomainRedirect = flex.FlattenBoolPointer(from.NxdomainRedirect)
	m.NxdomainRedirectAddresses = flex.FlattenFrameworkListString(ctx, from.NxdomainRedirectAddresses, diags)
	m.NxdomainRedirectAddressesV6 = flex.FlattenFrameworkListString(ctx, from.NxdomainRedirectAddressesV6, diags)
	m.NxdomainRedirectTtl = flex.FlattenInt64Pointer(from.NxdomainRedirectTtl)
	m.NxdomainRulesets = flex.FlattenFrameworkListString(ctx, from.NxdomainRulesets, diags)
	m.Recursion = flex.FlattenBoolPointer(from.Recursion)
	m.ResponseRateLimiting = FlattenViewResponseRateLimiting(ctx, from.ResponseRateLimiting, diags)
	m.RootNameServerType = flex.FlattenStringPointerEmptyAsNull(from.RootNameServerType)
	m.RpzDropIpRuleEnabled = flex.FlattenBoolPointer(from.RpzDropIpRuleEnabled)
	m.RpzDropIpRuleMinPrefixLengthIpv4 = flex.FlattenInt64Pointer(from.RpzDropIpRuleMinPrefixLengthIpv4)
	m.RpzDropIpRuleMinPrefixLengthIpv6 = flex.FlattenInt64Pointer(from.RpzDropIpRuleMinPrefixLengthIpv6)
	m.RpzQnameWaitRecurse = flex.FlattenBoolPointer(from.RpzQnameWaitRecurse)
	m.ScavengingSettings = FlattenViewScavengingSettings(ctx, from.ScavengingSettings, diags)
	m.Sortlist = flex.FlattenFrameworkListNestedBlock(ctx, from.Sortlist, ViewSortlistAttrTypes, diags, FlattenViewSortlist)
}

// Flatten merges API response onto existing UDDI model.
func (m *UDDIViewModel) Flatten(ctx context.Context, from *coremodel.UDDIViewExt, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.AddEdnsOptionInOutgoingQuery = flex.FlattenBoolPointer(from.AddEdnsOptionInOutgoingQuery)
	m.Comment = flex.FlattenStringPointer(from.Comment)
	m.CompartmentId = flex.FlattenStringPointer(from.CompartmentId)
	m.CustomRootNs = flex.FlattenFrameworkListNestedBlock(ctx, from.CustomRootNs, RootNSAttrTypes, diags, FlattenRootNS)
	m.CustomRootNsEnabled = flex.FlattenBoolPointer(from.CustomRootNsEnabled)
	m.Disabled = flex.FlattenBoolPointer(from.Disabled)
	m.DnssecEnableValidation = flex.FlattenBoolPointer(from.DnssecEnableValidation)
	m.DnssecEnabled = flex.FlattenBoolPointer(from.DnssecEnabled)
	m.DnssecTrustAnchors = flex.FlattenFrameworkListNestedBlock(ctx, from.DnssecTrustAnchors, TrustAnchorAttrTypes, diags, FlattenTrustAnchor)
	m.DnssecValidateExpiry = flex.FlattenBoolPointer(from.DnssecValidateExpiry)
	m.DtcConfig = FlattenDTCConfig(ctx, from.DtcConfig, diags)
	m.EcsEnabled = flex.FlattenBoolPointer(from.EcsEnabled)
	m.EcsForwarding = flex.FlattenBoolPointer(from.EcsForwarding)
	m.EcsPrefixV4 = flex.FlattenInt64Pointer(from.EcsPrefixV4)
	m.EcsPrefixV6 = flex.FlattenInt64Pointer(from.EcsPrefixV6)
	m.EcsZones = flex.FlattenFrameworkListNestedBlock(ctx, from.EcsZones, ECSZoneAttrTypes, diags, FlattenECSZone)
	m.EdnsUdpSize = flex.FlattenInt64Pointer(from.EdnsUdpSize)
	m.FilterAaaaAcl = flex.FlattenFrameworkListNestedBlock(ctx, from.FilterAaaaAcl, ACLItemAttrTypes, diags, FlattenACLItem)
	m.FilterAaaaOnV4 = flex.FlattenStringPointer(from.FilterAaaaOnV4)
	m.Forwarders = flex.FlattenFrameworkListNestedBlock(ctx, from.Forwarders, ForwarderAttrTypes, diags, FlattenForwarder)
	m.ForwardersOnly = flex.FlattenBoolPointer(from.ForwardersOnly)
	m.GssTsigEnabled = flex.FlattenBoolPointer(from.GssTsigEnabled)
	m.InheritanceSources = FlattenViewInheritance(ctx, from.InheritanceSources, diags)
	m.IpSpaces = flex.FlattenFrameworkListString(ctx, from.IpSpaces, diags)
	m.LameTtl = flex.FlattenInt64Pointer(from.LameTtl)
	m.MatchClientsAcl = flex.FlattenFrameworkListNestedBlock(ctx, from.MatchClientsAcl, ACLItemAttrTypes, diags, FlattenACLItem)
	m.MatchDestinationsAcl = flex.FlattenFrameworkListNestedBlock(ctx, from.MatchDestinationsAcl, ACLItemAttrTypes, diags, FlattenACLItem)
	m.MatchRecursiveOnly = flex.FlattenBoolPointer(from.MatchRecursiveOnly)
	m.MaxCacheTtl = flex.FlattenInt64Pointer(from.MaxCacheTtl)
	m.MaxNegativeTtl = flex.FlattenInt64Pointer(from.MaxNegativeTtl)
	m.MaxUdpSize = flex.FlattenInt64Pointer(from.MaxUdpSize)
	m.MinimalResponses = flex.FlattenBoolPointer(from.MinimalResponses)
	m.Name = flex.FlattenString(from.Name)
	m.Notify = flex.FlattenBoolPointer(from.Notify)
	m.QueryAcl = flex.FlattenFrameworkListNestedBlock(ctx, from.QueryAcl, ACLItemAttrTypes, diags, FlattenACLItem)
	m.RecursionAcl = flex.FlattenFrameworkListNestedBlock(ctx, from.RecursionAcl, ACLItemAttrTypes, diags, FlattenACLItem)
	m.RecursionEnabled = flex.FlattenBoolPointer(from.RecursionEnabled)
	m.SortList = flex.FlattenFrameworkListNestedBlock(ctx, from.SortList, SortListItemAttrTypes, diags, FlattenSortListItem)
	m.SynthesizeAddressRecordsFromHttps = flex.FlattenBoolPointer(from.SynthesizeAddressRecordsFromHttps)
	tagsAll := flex.FlattenMapStringAny(ctx, from.Tags, diags)
	if m.Tags.IsNull() || m.Tags.IsUnknown() {
		m.Tags = tagsAll
	}
	m.TagsAll = tagsAll
	m.TransferAcl = flex.FlattenFrameworkListNestedBlock(ctx, from.TransferAcl, ACLItemAttrTypes, diags, FlattenACLItem)
	m.UpdateAcl = flex.FlattenFrameworkListNestedBlock(ctx, from.UpdateAcl, ACLItemAttrTypes, diags, FlattenACLItem)
	m.UseForwardersForSubzones = flex.FlattenBoolPointer(from.UseForwardersForSubzones)
	m.UseRootForwardersForLocalResolutionWithB1td = flex.FlattenBoolPointer(from.UseRootForwardersForLocalResolutionWithB1td)
	m.ZoneAuthority = FlattenZoneAuthority(ctx, from.ZoneAuthority, diags)
}
