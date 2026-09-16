package dns

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/mapvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/mapdefault"
	objectplanmodifier "github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"

	coremodel "github.com/infobloxopen/terraform-provider-infoblox/internal/core/model/dns"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
	customvalidator "github.com/infobloxopen/terraform-provider-infoblox/internal/validator"
)

type DnsServerModel struct {
	Id   types.String `tfsdk:"id"`
	UDDI types.Object `tfsdk:"uddi"`
}

var DnsServerAttrTypes = map[string]attr.Type{
	"id":   types.StringType,
	"uddi": types.ObjectType{AttrTypes: UDDIDnsServerAttrTypes},
}

type UDDIDnsServerModel struct {
	AddEdnsOptionInOutgoingQuery                types.Bool   `tfsdk:"add_edns_option_in_outgoing_query"`
	AutoSortViews                               types.Bool   `tfsdk:"auto_sort_views"`
	Comment                                     types.String `tfsdk:"comment"`
	CustomRootNs                                types.List   `tfsdk:"custom_root_ns"`
	CustomRootNsEnabled                         types.Bool   `tfsdk:"custom_root_ns_enabled"`
	DnssecEnableValidation                      types.Bool   `tfsdk:"dnssec_enable_validation"`
	DnssecEnabled                               types.Bool   `tfsdk:"dnssec_enabled"`
	DnssecTrustAnchors                          types.List   `tfsdk:"dnssec_trust_anchors"`
	DnssecValidateExpiry                        types.Bool   `tfsdk:"dnssec_validate_expiry"`
	EcsEnabled                                  types.Bool   `tfsdk:"ecs_enabled"`
	EcsForwarding                               types.Bool   `tfsdk:"ecs_forwarding"`
	EcsPrefixV4                                 types.Int64  `tfsdk:"ecs_prefix_v4"`
	EcsPrefixV6                                 types.Int64  `tfsdk:"ecs_prefix_v6"`
	EcsZones                                    types.List   `tfsdk:"ecs_zones"`
	FilterAaaaAcl                               types.List   `tfsdk:"filter_aaaa_acl"`
	FilterAaaaOnV4                              types.String `tfsdk:"filter_aaaa_on_v4"`
	Forwarders                                  types.List   `tfsdk:"forwarders"`
	ForwardersOnly                              types.Bool   `tfsdk:"forwarders_only"`
	GssTsigEnabled                              types.Bool   `tfsdk:"gss_tsig_enabled"`
	InheritanceSources                          types.Object `tfsdk:"inheritance_sources"`
	KerberosKeys                                types.List   `tfsdk:"kerberos_keys"`
	LameTtl                                     types.Int64  `tfsdk:"lame_ttl"`
	LogQueryResponse                            types.Bool   `tfsdk:"log_query_response"`
	MatchRecursiveOnly                          types.Bool   `tfsdk:"match_recursive_only"`
	MaxCacheTtl                                 types.Int64  `tfsdk:"max_cache_ttl"`
	MaxNegativeTtl                              types.Int64  `tfsdk:"max_negative_ttl"`
	MinimalResponses                            types.Bool   `tfsdk:"minimal_responses"`
	Name                                        types.String `tfsdk:"name"`
	Notify                                      types.Bool   `tfsdk:"notify"`
	QueryAcl                                    types.List   `tfsdk:"query_acl"`
	QueryPort                                   types.Int64  `tfsdk:"query_port"`
	RecursionAcl                                types.List   `tfsdk:"recursion_acl"`
	RecursionEnabled                            types.Bool   `tfsdk:"recursion_enabled"`
	RecursiveClients                            types.Int64  `tfsdk:"recursive_clients"`
	ResolverQueryTimeout                        types.Int64  `tfsdk:"resolver_query_timeout"`
	SecondaryAxfrQueryLimit                     types.Int64  `tfsdk:"secondary_axfr_query_limit"`
	SecondarySoaQueryLimit                      types.Int64  `tfsdk:"secondary_soa_query_limit"`
	SortList                                    types.List   `tfsdk:"sort_list"`
	SynthesizeAddressRecordsFromHttps           types.Bool   `tfsdk:"synthesize_address_records_from_https"`
	Tags                                        types.Map    `tfsdk:"tags"`
	TagsAll                                     types.Map    `tfsdk:"tags_all"`
	TransferAcl                                 types.List   `tfsdk:"transfer_acl"`
	UpdateAcl                                   types.List   `tfsdk:"update_acl"`
	UseForwardersForSubzones                    types.Bool   `tfsdk:"use_forwarders_for_subzones"`
	UseRootForwardersForLocalResolutionWithB1td types.Bool   `tfsdk:"use_root_forwarders_for_local_resolution_with_b1td"`
	Views                                       types.List   `tfsdk:"views"`
}

var UDDIDnsServerAttrTypes = map[string]attr.Type{
	"add_edns_option_in_outgoing_query":     types.BoolType,
	"auto_sort_views":                       types.BoolType,
	"comment":                               types.StringType,
	"custom_root_ns":                        types.ListType{ElemType: types.ObjectType{AttrTypes: RootNSAttrTypes}},
	"custom_root_ns_enabled":                types.BoolType,
	"dnssec_enable_validation":              types.BoolType,
	"dnssec_enabled":                        types.BoolType,
	"dnssec_trust_anchors":                  types.ListType{ElemType: types.ObjectType{AttrTypes: TrustAnchorAttrTypes}},
	"dnssec_validate_expiry":                types.BoolType,
	"ecs_enabled":                           types.BoolType,
	"ecs_forwarding":                        types.BoolType,
	"ecs_prefix_v4":                         types.Int64Type,
	"ecs_prefix_v6":                         types.Int64Type,
	"ecs_zones":                             types.ListType{ElemType: types.ObjectType{AttrTypes: ECSZoneAttrTypes}},
	"filter_aaaa_acl":                       types.ListType{ElemType: types.ObjectType{AttrTypes: ACLItemAttrTypes}},
	"filter_aaaa_on_v4":                     types.StringType,
	"forwarders":                            types.ListType{ElemType: types.ObjectType{AttrTypes: ForwarderAttrTypes}},
	"forwarders_only":                       types.BoolType,
	"gss_tsig_enabled":                      types.BoolType,
	"inheritance_sources":                   types.ObjectType{AttrTypes: ServerInheritanceAttrTypes},
	"kerberos_keys":                         types.ListType{ElemType: types.ObjectType{AttrTypes: KerberosKeyAttrTypes}},
	"lame_ttl":                              types.Int64Type,
	"log_query_response":                    types.BoolType,
	"match_recursive_only":                  types.BoolType,
	"max_cache_ttl":                         types.Int64Type,
	"max_negative_ttl":                      types.Int64Type,
	"minimal_responses":                     types.BoolType,
	"name":                                  types.StringType,
	"notify":                                types.BoolType,
	"query_acl":                             types.ListType{ElemType: types.ObjectType{AttrTypes: ACLItemAttrTypes}},
	"query_port":                            types.Int64Type,
	"recursion_acl":                         types.ListType{ElemType: types.ObjectType{AttrTypes: ACLItemAttrTypes}},
	"recursion_enabled":                     types.BoolType,
	"recursive_clients":                     types.Int64Type,
	"resolver_query_timeout":                types.Int64Type,
	"secondary_axfr_query_limit":            types.Int64Type,
	"secondary_soa_query_limit":             types.Int64Type,
	"sort_list":                             types.ListType{ElemType: types.ObjectType{AttrTypes: SortListItemAttrTypes}},
	"synthesize_address_records_from_https": types.BoolType,
	"tags":                                  types.MapType{ElemType: types.StringType},
	"tags_all":                              types.MapType{ElemType: types.StringType},
	"transfer_acl":                          types.ListType{ElemType: types.ObjectType{AttrTypes: ACLItemAttrTypes}},
	"update_acl":                            types.ListType{ElemType: types.ObjectType{AttrTypes: ACLItemAttrTypes}},
	"use_forwarders_for_subzones":           types.BoolType,
	"use_root_forwarders_for_local_resolution_with_b1td": types.BoolType,
	"views": types.ListType{ElemType: types.ObjectType{AttrTypes: DisplayViewAttrTypes}},
}

const (
	DnsServerInheritanceType = "full"
	DnsServerReturnFields    = ""
)

var DnsServerResourceSchemaAttributes = map[string]schema.Attribute{
	"id": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The resource identifier.",
	},
	"uddi": schema.SingleNestedAttribute{
		Optional:            true,
		MarkdownDescription: "UDDI backend-specific fields.",
		Attributes:          DnsServerResourceUddiSchemaAttributes,
	},
}

var DnsServerResourceUddiSchemaAttributes = map[string]schema.Attribute{
	"add_edns_option_in_outgoing_query": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "_add_edns_option_in_outgoing_query_ adds client IP, MAC address and view name into outgoing recursive query. Defaults to _false_.",
	},
	"auto_sort_views": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(true),
		MarkdownDescription: "Optional. Controls manual/automatic views ordering.  Defaults to _true_.",
	},
	"comment": schema.StringAttribute{
		Default:             stringdefault.StaticString(""),
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Optional. Comment for configuration.",
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
	"ecs_enabled": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Optional. _true_ to enable EDNS client subnet for recursive queries. Other _ecs_*_ fields are ignored if this field is not enabled.  Defaults to _false_.",
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
		Attributes: ServerInheritanceResourceSchemaAttributes,
		Optional:   true,
		Computed:   true,
		PlanModifiers: []planmodifier.Object{
			objectplanmodifier.UseStateForUnknown(),
		},
		MarkdownDescription: "Inheritance configuration specifies how and which fields _Server_ object inherits from _Global_ parent.",
	},
	"kerberos_keys": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: KerberosKeyResourceSchemaAttributes,
		},
		Optional: true,
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
		},
		MarkdownDescription: "_kerberos_keys_ contains a list of keys for GSS-TSIG signed dynamic updates.  Defaults to empty.",
	},
	"lame_ttl": schema.Int64Attribute{
		Optional:            true,
		Computed:            true,
		Default:             int64default.StaticInt64(600),
		MarkdownDescription: "Optional. Unused in the current on-prem DNS server implementation.  Unsigned integer, min 0 max 3600 (1h).  Defaults to 600.",
	},
	"log_query_response": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(true),
		MarkdownDescription: "Optional. Control DNS query/response logging functionality.  Defaults to _true_.",
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
	"minimal_responses": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Optional. When enabled, the DNS server will only add records to the authority and additional data sections when they are required.  Defaults to _false_.",
	},
	"name": schema.StringAttribute{
		Required:            true,
		MarkdownDescription: "Name of configuration.",
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
	"query_port": schema.Int64Attribute{
		Optional:            true,
		Computed:            true,
		Default:             int64default.StaticInt64(0),
		MarkdownDescription: "Optional. Source port for outbound DNS queries. When set to 0 the port is unspecified and the implementation may randomize it using any available ports.  Defaults to 0.",
	},
	"recursion_acl": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: ACLItemResourceSchemaAttributes,
		},
		Optional: true,
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
		},
		MarkdownDescription: "Optional. Clients must match this ACL to make recursive queries. If this ACL is empty, then the _query_acl_ field will be used instead.  Defaults to empty.",
	},
	"recursion_enabled": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(true),
		MarkdownDescription: "Optional. _true_ to allow recursive DNS queries.  Defaults to _true_.",
	},
	"recursive_clients": schema.Int64Attribute{
		Optional:            true,
		Computed:            true,
		Default:             int64default.StaticInt64(1000),
		MarkdownDescription: "Optional. Defines the number of simultaneous recursive lookups the server will perform on behalf of its clients.  Defaults to 1000.",
	},
	"resolver_query_timeout": schema.Int64Attribute{
		Optional:            true,
		Computed:            true,
		Default:             int64default.StaticInt64(10),
		MarkdownDescription: "Optional. Seconds before a recursive query times out.  Unsigned integer, min 10 max 30.  Defaults to 10.",
	},
	"secondary_axfr_query_limit": schema.Int64Attribute{
		Optional:            true,
		Computed:            true,
		Default:             int64default.StaticInt64(0),
		MarkdownDescription: "Optional. Maximum concurrent inbound AXFRs. When set to 0 a host-dependent default will be used.  Defaults to 0.",
	},
	"secondary_soa_query_limit": schema.Int64Attribute{
		Optional:            true,
		Computed:            true,
		Default:             int64default.StaticInt64(0),
		MarkdownDescription: "Optional. Maximum concurrent outbound SOA queries. When set to 0 a host-dependent default will be used.  Defaults to 0.",
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
		MarkdownDescription: "_use_root_forwarders_for_local_resolution_with_b1td_ allows DNS recursive queries sent to root forwarders for local resolution when deployed alongside Universal Thread Defense. Defaults to _false_.",
	},
	"views": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: DisplayViewResourceSchemaAttributes,
		},
		Computed: true,
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
		},
		MarkdownDescription: "Optional. Ordered list of _dns/display_view_ objects served by any of _dns/host_ assigned to a particular DNS Config Profile. Automatically determined. Allows re-ordering only.",
	},
}

// Expand converts the TF model to the infoblox core model
func (m *DnsServerModel) Expand(ctx context.Context, diags *diag.Diagnostics, isCreate bool) *coremodel.DnsServer {
	if m == nil {
		return nil
	}

	obj := &coremodel.DnsServer{}

	// Expand UDDI nested attribute (returns nil if not present)
	uddiModel := flex.ExpandNestedObject[UDDIDnsServerModel](ctx, m.UDDI, diags)
	if uddiModel != nil {
		obj.UDDI = uddiModel.Expand(ctx, diags)
	}

	return obj
}

// Expand converts the UDDI TF model to the core model.
func (m *UDDIDnsServerModel) Expand(ctx context.Context, diags *diag.Diagnostics) *coremodel.UDDIDnsServerExt {
	return &coremodel.UDDIDnsServerExt{
		AddEdnsOptionInOutgoingQuery:      flex.ExpandBoolPointer(m.AddEdnsOptionInOutgoingQuery),
		AutoSortViews:                     flex.ExpandBoolPointer(m.AutoSortViews),
		Comment:                           flex.ExpandStringPointer(m.Comment),
		CustomRootNs:                      flex.ExpandFrameworkListNestedBlock(ctx, m.CustomRootNs, diags, ExpandRootNS),
		CustomRootNsEnabled:               flex.ExpandBoolPointer(m.CustomRootNsEnabled),
		DnssecEnableValidation:            flex.ExpandBoolPointer(m.DnssecEnableValidation),
		DnssecEnabled:                     flex.ExpandBoolPointer(m.DnssecEnabled),
		DnssecTrustAnchors:                flex.ExpandFrameworkListNestedBlock(ctx, m.DnssecTrustAnchors, diags, ExpandTrustAnchor),
		DnssecValidateExpiry:              flex.ExpandBoolPointer(m.DnssecValidateExpiry),
		EcsEnabled:                        flex.ExpandBoolPointer(m.EcsEnabled),
		EcsForwarding:                     flex.ExpandBoolPointer(m.EcsForwarding),
		EcsPrefixV4:                       flex.ExpandInt64Pointer(m.EcsPrefixV4),
		EcsPrefixV6:                       flex.ExpandInt64Pointer(m.EcsPrefixV6),
		EcsZones:                          flex.ExpandFrameworkListNestedBlock(ctx, m.EcsZones, diags, ExpandECSZone),
		FilterAaaaAcl:                     flex.ExpandFrameworkListNestedBlock(ctx, m.FilterAaaaAcl, diags, ExpandACLItem),
		FilterAaaaOnV4:                    flex.ExpandStringPointer(m.FilterAaaaOnV4),
		Forwarders:                        flex.ExpandFrameworkListNestedBlock(ctx, m.Forwarders, diags, ExpandForwarder),
		ForwardersOnly:                    flex.ExpandBoolPointer(m.ForwardersOnly),
		GssTsigEnabled:                    flex.ExpandBoolPointer(m.GssTsigEnabled),
		InheritanceSources:                ExpandServerInheritance(ctx, m.InheritanceSources, diags),
		KerberosKeys:                      flex.ExpandFrameworkListNestedBlock(ctx, m.KerberosKeys, diags, ExpandKerberosKey),
		LameTtl:                           flex.ExpandInt64Pointer(m.LameTtl),
		LogQueryResponse:                  flex.ExpandBoolPointer(m.LogQueryResponse),
		MatchRecursiveOnly:                flex.ExpandBoolPointer(m.MatchRecursiveOnly),
		MaxCacheTtl:                       flex.ExpandInt64Pointer(m.MaxCacheTtl),
		MaxNegativeTtl:                    flex.ExpandInt64Pointer(m.MaxNegativeTtl),
		MinimalResponses:                  flex.ExpandBoolPointer(m.MinimalResponses),
		Name:                              flex.ExpandString(m.Name),
		Notify:                            flex.ExpandBoolPointer(m.Notify),
		QueryAcl:                          flex.ExpandFrameworkListNestedBlock(ctx, m.QueryAcl, diags, ExpandACLItem),
		QueryPort:                         flex.ExpandInt64Pointer(m.QueryPort),
		RecursionAcl:                      flex.ExpandFrameworkListNestedBlock(ctx, m.RecursionAcl, diags, ExpandACLItem),
		RecursionEnabled:                  flex.ExpandBoolPointer(m.RecursionEnabled),
		RecursiveClients:                  flex.ExpandInt64Pointer(m.RecursiveClients),
		ResolverQueryTimeout:              flex.ExpandInt64Pointer(m.ResolverQueryTimeout),
		SecondaryAxfrQueryLimit:           flex.ExpandInt64Pointer(m.SecondaryAxfrQueryLimit),
		SecondarySoaQueryLimit:            flex.ExpandInt64Pointer(m.SecondarySoaQueryLimit),
		SortList:                          flex.ExpandFrameworkListNestedBlock(ctx, m.SortList, diags, ExpandSortListItem),
		SynthesizeAddressRecordsFromHttps: flex.ExpandBoolPointer(m.SynthesizeAddressRecordsFromHttps),
		Tags:                              flex.ExpandMapStringAny(ctx, m.Tags, diags),
		TransferAcl:                       flex.ExpandFrameworkListNestedBlock(ctx, m.TransferAcl, diags, ExpandACLItem),
		UpdateAcl:                         flex.ExpandFrameworkListNestedBlock(ctx, m.UpdateAcl, diags, ExpandACLItem),
		UseForwardersForSubzones:          flex.ExpandBoolPointer(m.UseForwardersForSubzones),
		UseRootForwardersForLocalResolutionWithB1td: flex.ExpandBoolPointer(m.UseRootForwardersForLocalResolutionWithB1td),
	}
}

// Flatten populates the TF model from a core response.
func (m *DnsServerModel) Flatten(ctx context.Context, resp *coremodel.DnsServer, diags *diag.Diagnostics) {
	if resp == nil {
		return
	}

	m.Id = flex.FlattenStringPointer(resp.Id)

	// Extract existing UDDI model, flatten API response onto it, convert back
	uddiModel := flex.ExpandNestedObject[UDDIDnsServerModel](ctx, m.UDDI, diags)
	if uddiModel == nil {
		uddiModel = &UDDIDnsServerModel{}
	}
	uddiModel.Flatten(ctx, resp.UDDI, diags)
	if resp.UDDI != nil {
		m.UDDI = flex.FlattenNestedObject(ctx, uddiModel, UDDIDnsServerAttrTypes, diags)
	} else {
		m.UDDI = types.ObjectNull(UDDIDnsServerAttrTypes)
	}
}

// Flatten merges API response onto existing UDDI model.
func (m *UDDIDnsServerModel) Flatten(ctx context.Context, from *coremodel.UDDIDnsServerExt, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.AddEdnsOptionInOutgoingQuery = flex.FlattenBoolPointer(from.AddEdnsOptionInOutgoingQuery)
	m.AutoSortViews = flex.FlattenBoolPointer(from.AutoSortViews)
	m.Comment = flex.FlattenStringPointer(from.Comment)
	m.CustomRootNs = flex.FlattenFrameworkListNestedBlock(ctx, from.CustomRootNs, RootNSAttrTypes, diags, FlattenRootNS)
	m.CustomRootNsEnabled = flex.FlattenBoolPointer(from.CustomRootNsEnabled)
	m.DnssecEnableValidation = flex.FlattenBoolPointer(from.DnssecEnableValidation)
	m.DnssecEnabled = flex.FlattenBoolPointer(from.DnssecEnabled)
	m.DnssecTrustAnchors = flex.FlattenFrameworkListNestedBlock(ctx, from.DnssecTrustAnchors, TrustAnchorAttrTypes, diags, FlattenTrustAnchor)
	m.DnssecValidateExpiry = flex.FlattenBoolPointer(from.DnssecValidateExpiry)
	m.EcsEnabled = flex.FlattenBoolPointer(from.EcsEnabled)
	m.EcsForwarding = flex.FlattenBoolPointer(from.EcsForwarding)
	m.EcsPrefixV4 = flex.FlattenInt64Pointer(from.EcsPrefixV4)
	m.EcsPrefixV6 = flex.FlattenInt64Pointer(from.EcsPrefixV6)
	m.EcsZones = flex.FlattenFrameworkListNestedBlock(ctx, from.EcsZones, ECSZoneAttrTypes, diags, FlattenECSZone)
	m.FilterAaaaAcl = flex.FlattenFrameworkListNestedBlock(ctx, from.FilterAaaaAcl, ACLItemAttrTypes, diags, FlattenACLItem)
	m.FilterAaaaOnV4 = flex.FlattenStringPointer(from.FilterAaaaOnV4)
	m.Forwarders = flex.FlattenFrameworkListNestedBlock(ctx, from.Forwarders, ForwarderAttrTypes, diags, FlattenForwarder)
	m.ForwardersOnly = flex.FlattenBoolPointer(from.ForwardersOnly)
	m.GssTsigEnabled = flex.FlattenBoolPointer(from.GssTsigEnabled)
	m.InheritanceSources = FlattenServerInheritance(ctx, from.InheritanceSources, diags)
	m.KerberosKeys = flex.FlattenFrameworkListNestedBlock(ctx, from.KerberosKeys, KerberosKeyAttrTypes, diags, FlattenKerberosKey)
	m.LameTtl = flex.FlattenInt64Pointer(from.LameTtl)
	m.LogQueryResponse = flex.FlattenBoolPointer(from.LogQueryResponse)
	m.MatchRecursiveOnly = flex.FlattenBoolPointer(from.MatchRecursiveOnly)
	m.MaxCacheTtl = flex.FlattenInt64Pointer(from.MaxCacheTtl)
	m.MaxNegativeTtl = flex.FlattenInt64Pointer(from.MaxNegativeTtl)
	m.MinimalResponses = flex.FlattenBoolPointer(from.MinimalResponses)
	m.Name = flex.FlattenString(from.Name)
	m.Notify = flex.FlattenBoolPointer(from.Notify)
	m.QueryAcl = flex.FlattenFrameworkListNestedBlock(ctx, from.QueryAcl, ACLItemAttrTypes, diags, FlattenACLItem)
	m.QueryPort = flex.FlattenInt64Pointer(from.QueryPort)
	m.RecursionAcl = flex.FlattenFrameworkListNestedBlock(ctx, from.RecursionAcl, ACLItemAttrTypes, diags, FlattenACLItem)
	m.RecursionEnabled = flex.FlattenBoolPointer(from.RecursionEnabled)
	m.RecursiveClients = flex.FlattenInt64Pointer(from.RecursiveClients)
	m.ResolverQueryTimeout = flex.FlattenInt64Pointer(from.ResolverQueryTimeout)
	m.SecondaryAxfrQueryLimit = flex.FlattenInt64Pointer(from.SecondaryAxfrQueryLimit)
	m.SecondarySoaQueryLimit = flex.FlattenInt64Pointer(from.SecondarySoaQueryLimit)
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
	m.Views = flex.FlattenFrameworkListNestedBlock(ctx, from.Views, DisplayViewAttrTypes, diags, FlattenDisplayView)
}
