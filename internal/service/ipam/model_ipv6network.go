package ipam

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-nettypes/cidrtypes"
	"github.com/hashicorp/terraform-plugin-framework-nettypes/iptypes"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/boolvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/mapvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/mapdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectdefault"
	objectplanmodifier "github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	stringplanmodifier "github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"

	coremodel "github.com/infobloxopen/terraform-provider-infoblox/internal/core/model/ipam"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/dynamicallocation"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
	immutable "github.com/infobloxopen/terraform-provider-infoblox/internal/planmodifiers/immutable"
	importmod "github.com/infobloxopen/terraform-provider-infoblox/internal/planmodifiers/import"
	internaltypes "github.com/infobloxopen/terraform-provider-infoblox/internal/types"
	customvalidator "github.com/infobloxopen/terraform-provider-infoblox/internal/validator"
)

type Ipv6networkModel struct {
	Id   types.String `tfsdk:"id"`
	NIOS types.Object `tfsdk:"nios"`
	UDDI types.Object `tfsdk:"uddi"`
}

var Ipv6networkAttrTypes = map[string]attr.Type{
	"id":   types.StringType,
	"nios": types.ObjectType{AttrTypes: NIOSIpv6networkAttrTypes},
	"uddi": types.ObjectType{AttrTypes: UDDIIpv6networkAttrTypes},
}

type NIOSIpv6networkModel struct {
	AutoCreateReversezone            types.Bool           `tfsdk:"auto_create_reversezone"`
	CloudInfo                        types.Object         `tfsdk:"cloud_info"`
	Comment                          types.String         `tfsdk:"comment"`
	DdnsDomainname                   types.String         `tfsdk:"ddns_domainname"`
	DdnsEnableOptionFqdn             types.Bool           `tfsdk:"ddns_enable_option_fqdn"`
	DdnsGenerateHostname             types.Bool           `tfsdk:"ddns_generate_hostname"`
	DdnsServerAlwaysUpdates          types.Bool           `tfsdk:"ddns_server_always_updates"`
	DdnsTtl                          types.Int64          `tfsdk:"ddns_ttl"`
	DeleteReason                     types.String         `tfsdk:"delete_reason"`
	Disable                          types.Bool           `tfsdk:"disable"`
	DiscoveredBridgeDomain           types.String         `tfsdk:"discovered_bridge_domain"`
	DiscoveredTenant                 types.String         `tfsdk:"discovered_tenant"`
	DiscoveryBasicPollSettings       types.Object         `tfsdk:"discovery_basic_poll_settings"`
	DiscoveryBlackoutSetting         types.Object         `tfsdk:"discovery_blackout_setting"`
	DiscoveryMember                  types.String         `tfsdk:"discovery_member"`
	DomainName                       types.String         `tfsdk:"domain_name"`
	DomainNameServers                types.List           `tfsdk:"domain_name_servers"`
	EnableDdns                       types.Bool           `tfsdk:"enable_ddns"`
	EnableDiscovery                  types.Bool           `tfsdk:"enable_discovery"`
	EnableIfmapPublishing            types.Bool           `tfsdk:"enable_ifmap_publishing"`
	EnableImmediateDiscovery         types.Bool           `tfsdk:"enable_immediate_discovery"`
	ExtAttrs                         types.Map            `tfsdk:"ext_attrs"`
	ExtAttrsAll                      types.Map            `tfsdk:"ext_attrs_all"`
	FederatedRealms                  types.List           `tfsdk:"federated_realms"`
	LogicFilterRules                 types.List           `tfsdk:"logic_filter_rules"`
	Members                          types.List           `tfsdk:"members"`
	MgmPrivate                       types.Bool           `tfsdk:"mgm_private"`
	Network                          cidrtypes.IPv6Prefix `tfsdk:"network"`
	NetworkView                      types.String         `tfsdk:"network_view"`
	Options                          types.List           `tfsdk:"options"`
	PortControlBlackoutSetting       types.Object         `tfsdk:"port_control_blackout_setting"`
	PreferredLifetime                types.Int64          `tfsdk:"preferred_lifetime"`
	RecycleLeases                    types.Bool           `tfsdk:"recycle_leases"`
	RestartIfNeeded                  types.Bool           `tfsdk:"restart_if_needed"`
	RirOrganization                  types.String         `tfsdk:"rir_organization"`
	RirRegistrationAction            types.String         `tfsdk:"rir_registration_action"`
	RirRegistrationStatus            types.String         `tfsdk:"rir_registration_status"`
	SamePortControlDiscoveryBlackout types.Bool           `tfsdk:"same_port_control_discovery_blackout"`
	SendRirRequest                   types.Bool           `tfsdk:"send_rir_request"`
	SubscribeSettings                types.Object         `tfsdk:"subscribe_settings"`
	Unmanaged                        types.Bool           `tfsdk:"unmanaged"`
	UpdateDnsOnLeaseRenewal          types.Bool           `tfsdk:"update_dns_on_lease_renewal"`
	ValidLifetime                    types.Int64          `tfsdk:"valid_lifetime"`
	Vlans                            types.List           `tfsdk:"vlans"`
	ZoneAssociations                 types.List           `tfsdk:"zone_associations"`
	DynamicAllocation                types.Object         `tfsdk:"dynamic_allocation"`
}

var NIOSIpv6networkAttrTypes = map[string]attr.Type{
	"auto_create_reversezone":              types.BoolType,
	"cloud_info":                           types.ObjectType{AttrTypes: Ipv6networkCloudInfoAttrTypes},
	"comment":                              types.StringType,
	"ddns_domainname":                      types.StringType,
	"ddns_enable_option_fqdn":              types.BoolType,
	"ddns_generate_hostname":               types.BoolType,
	"ddns_server_always_updates":           types.BoolType,
	"ddns_ttl":                             types.Int64Type,
	"delete_reason":                        types.StringType,
	"disable":                              types.BoolType,
	"discovered_bridge_domain":             types.StringType,
	"discovered_tenant":                    types.StringType,
	"discovery_basic_poll_settings":        types.ObjectType{AttrTypes: Ipv6networkDiscoveryBasicPollSettingsAttrTypes},
	"discovery_blackout_setting":           types.ObjectType{AttrTypes: Ipv6networkDiscoveryBlackoutSettingAttrTypes},
	"discovery_member":                     types.StringType,
	"domain_name":                          types.StringType,
	"domain_name_servers":                  types.ListType{ElemType: types.StringType},
	"enable_ddns":                          types.BoolType,
	"enable_discovery":                     types.BoolType,
	"enable_ifmap_publishing":              types.BoolType,
	"enable_immediate_discovery":           types.BoolType,
	"ext_attrs":                            types.MapType{ElemType: types.StringType},
	"ext_attrs_all":                        types.MapType{ElemType: types.StringType},
	"federated_realms":                     types.ListType{ElemType: types.ObjectType{AttrTypes: Ipv6networkFederatedRealmsAttrTypes}},
	"logic_filter_rules":                   types.ListType{ElemType: types.ObjectType{AttrTypes: Ipv6networkLogicFilterRulesAttrTypes}},
	"members":                              types.ListType{ElemType: types.ObjectType{AttrTypes: Ipv6networkMembersAttrTypes}},
	"mgm_private":                          types.BoolType,
	"network":                              cidrtypes.IPv6PrefixType{},
	"network_view":                         types.StringType,
	"options":                              types.ListType{ElemType: types.ObjectType{AttrTypes: Ipv6networkOptionsAttrTypes}},
	"port_control_blackout_setting":        types.ObjectType{AttrTypes: Ipv6networkPortControlBlackoutSettingAttrTypes},
	"preferred_lifetime":                   types.Int64Type,
	"recycle_leases":                       types.BoolType,
	"restart_if_needed":                    types.BoolType,
	"rir_organization":                     types.StringType,
	"rir_registration_action":              types.StringType,
	"rir_registration_status":              types.StringType,
	"same_port_control_discovery_blackout": types.BoolType,
	"send_rir_request":                     types.BoolType,
	"subscribe_settings":                   types.ObjectType{AttrTypes: Ipv6networkSubscribeSettingsAttrTypes},
	"unmanaged":                            types.BoolType,
	"update_dns_on_lease_renewal":          types.BoolType,
	"valid_lifetime":                       types.Int64Type,
	"vlans":                                types.ListType{ElemType: types.ObjectType{AttrTypes: Ipv6networkVlansAttrTypes}},
	"zone_associations":                    types.ListType{ElemType: types.ObjectType{AttrTypes: Ipv6networkZoneAssociationsAttrTypes}},
	"dynamic_allocation":                   types.ObjectType{AttrTypes: dynamicallocation.NextAvailableNetworkAttrTypes},
}

type UDDIIpv6networkModel struct {
	Address                    iptypes.IPv6Address              `tfsdk:"address"`
	AsmConfig                  types.Object                     `tfsdk:"asm_config"`
	Cidr                       types.Int64                      `tfsdk:"cidr"`
	Comment                    types.String                     `tfsdk:"comment"`
	ConfigProfiles             types.List                       `tfsdk:"config_profiles"`
	DdnsClientUpdate           types.String                     `tfsdk:"ddns_client_update"`
	DdnsConflictResolutionMode types.String                     `tfsdk:"ddns_conflict_resolution_mode"`
	DdnsDomain                 types.String                     `tfsdk:"ddns_domain"`
	DdnsGenerateName           types.Bool                       `tfsdk:"ddns_generate_name"`
	DdnsGeneratedPrefix        types.String                     `tfsdk:"ddns_generated_prefix"`
	DdnsSendUpdates            types.Bool                       `tfsdk:"ddns_send_updates"`
	DdnsTtlPercent             types.Float64                    `tfsdk:"ddns_ttl_percent"`
	DdnsUpdateOnRenew          types.Bool                       `tfsdk:"ddns_update_on_renew"`
	DdnsUseConflictResolution  types.Bool                       `tfsdk:"ddns_use_conflict_resolution"`
	DhcpConfig                 types.Object                     `tfsdk:"dhcp_config"`
	DhcpHost                   types.String                     `tfsdk:"dhcp_host"`
	DhcpOptions                internaltypes.UnorderedListValue `tfsdk:"dhcp_options"`
	DisableDhcp                types.Bool                       `tfsdk:"disable_dhcp"`
	ExternalKeys               types.Map                        `tfsdk:"external_keys"`
	FederatedRealms            internaltypes.UnorderedListValue `tfsdk:"federated_realms"`
	HostnameRewriteChar        types.String                     `tfsdk:"hostname_rewrite_char"`
	HostnameRewriteEnabled     types.Bool                       `tfsdk:"hostname_rewrite_enabled"`
	HostnameRewriteRegex       types.String                     `tfsdk:"hostname_rewrite_regex"`
	InheritanceSources         types.Object                     `tfsdk:"inheritance_sources"`
	Name                       types.String                     `tfsdk:"name"`
	RebindTime                 types.Int64                      `tfsdk:"rebind_time"`
	RenewTime                  types.Int64                      `tfsdk:"renew_time"`
	Space                      types.String                     `tfsdk:"space"`
	Tags                       types.Map                        `tfsdk:"tags"`
	TagsAll                    types.Map                        `tfsdk:"tags_all"`
}

var UDDIIpv6networkAttrTypes = map[string]attr.Type{
	"address":                       iptypes.IPv6AddressType{},
	"asm_config":                    types.ObjectType{AttrTypes: ASMConfigAttrTypes},
	"cidr":                          types.Int64Type,
	"comment":                       types.StringType,
	"config_profiles":               types.ListType{ElemType: types.StringType},
	"ddns_client_update":            types.StringType,
	"ddns_conflict_resolution_mode": types.StringType,
	"ddns_domain":                   types.StringType,
	"ddns_generate_name":            types.BoolType,
	"ddns_generated_prefix":         types.StringType,
	"ddns_send_updates":             types.BoolType,
	"ddns_ttl_percent":              types.Float64Type,
	"ddns_update_on_renew":          types.BoolType,
	"ddns_use_conflict_resolution":  types.BoolType,
	"dhcp_config":                   types.ObjectType{AttrTypes: Ipv6networkDhcpConfigAttrTypes},
	"dhcp_host":                     types.StringType,
	"dhcp_options":                  internaltypes.UnorderedList{ListType: types.ListType{ElemType: types.ObjectType{AttrTypes: OptionItemAttrTypes}}},
	"disable_dhcp":                  types.BoolType,
	"external_keys":                 types.MapType{ElemType: types.StringType},
	"federated_realms":              internaltypes.UnorderedListOfStringType,
	"hostname_rewrite_char":         types.StringType,
	"hostname_rewrite_enabled":      types.BoolType,
	"hostname_rewrite_regex":        types.StringType,
	"inheritance_sources":           types.ObjectType{AttrTypes: DHCPInheritanceAttrTypes},
	"name":                          types.StringType,
	"rebind_time":                   types.Int64Type,
	"renew_time":                    types.Int64Type,
	"space":                         types.StringType,
	"tags":                          types.MapType{ElemType: types.StringType},
	"tags_all":                      types.MapType{ElemType: types.StringType},
}

const (
	Ipv6networkInheritanceType = "full"
	Ipv6networkReturnFields    = "cloud_info,comment,ddns_domainname,ddns_enable_option_fqdn,ddns_generate_hostname,ddns_server_always_updates,ddns_ttl,disable,discover_now_status,discovered_bgp_as,discovered_bridge_domain,discovered_tenant,discovered_vlan_id,discovered_vlan_name,discovered_vrf_description,discovered_vrf_name,discovered_vrf_rd,discovery_basic_poll_settings,discovery_blackout_setting,discovery_engine_type,discovery_member,domain_name,domain_name_servers,enable_ddns,enable_discovery,enable_ifmap_publishing,endpoint_sources,extattrs,federated_realms,last_rir_registration_update_sent,last_rir_registration_update_status,logic_filter_rules,members,mgm_private,mgm_private_overridable,ms_ad_user_data,network,network_container,network_view,options,port_control_blackout_setting,preferred_lifetime,recycle_leases,rir,rir_organization,rir_registration_status,same_port_control_discovery_blackout,subscribe_settings,unmanaged,unmanaged_count,update_dns_on_lease_renewal,use_blackout_setting,use_ddns_domainname,use_ddns_enable_option_fqdn,use_ddns_generate_hostname,use_ddns_ttl,use_discovery_basic_polling_settings,use_domain_name,use_domain_name_servers,use_enable_ddns,use_enable_discovery,use_enable_ifmap_publishing,use_logic_filter_rules,use_mgm_private,use_options,use_preferred_lifetime,use_recycle_leases,use_subscribe_settings,use_update_dns_on_lease_renewal,use_valid_lifetime,use_zone_associations,valid_lifetime,vlans,zone_associations"
)

var Ipv6networkResourceSchemaAttributes = map[string]schema.Attribute{
	"id": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The reference to the object.",
	},
	"nios": schema.SingleNestedAttribute{
		Optional:            true,
		MarkdownDescription: "NIOS backend-specific fields.",
		Attributes:          Ipv6networkResourceNiosSchemaAttributes,
	},
	"uddi": schema.SingleNestedAttribute{
		Optional:            true,
		MarkdownDescription: "UDDI backend-specific fields.",
		Attributes:          Ipv6networkResourceUddiSchemaAttributes,
	},
}

var Ipv6networkResourceNiosSchemaAttributes = map[string]schema.Attribute{
	"auto_create_reversezone": schema.BoolAttribute{
		Optional: true,
		Computed: true,
		Default:  booldefault.StaticBool(false),
		PlanModifiers: []planmodifier.Bool{
			immutable.ImmutableBool(),
		},
		MarkdownDescription: "This flag controls whether reverse zones are automatically created when the network is added.",
	},
	"cloud_info": schema.SingleNestedAttribute{
		Attributes:          Ipv6networkCloudInfoResourceSchemaAttributes,
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Structure containing all cloud API related information for this object.",
	},
	"comment": schema.StringAttribute{
		Optional: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "Comment for the network; maximum 256 characters.",
	},
	"ddns_domainname": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "The dynamic DNS domain name the appliance uses specifically for DDNS updates for this network.",
	},
	"ddns_enable_option_fqdn": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Use this method to set or retrieve the ddns_enable_option_fqdn flag of a DHCP IPv6 Network object. This method controls whether the FQDN option sent by the client is to be used, or if the server can automatically generate the FQDN. This setting overrides the upper-level settings.",
	},
	"ddns_generate_hostname": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "If this field is set to True, the DHCP server generates a hostname and updates DNS with it when the DHCP client request does not contain a hostname.",
	},
	"ddns_server_always_updates": schema.BoolAttribute{
		Optional: true,
		Computed: true,
		Default:  booldefault.StaticBool(true),
		Validators: []validator.Bool{
			boolvalidator.AlsoRequires(path.MatchRelative().AtParent().AtName("ddns_enable_option_fqdn")),
		},
		MarkdownDescription: "This field controls whether only the DHCP server is allowed to update DNS, regardless of the DHCP clients requests. Note that changes for this field take effect only if ddns_enable_option_fqdn is True.",
	},
	"ddns_ttl": schema.Int64Attribute{
		Optional:            true,
		Computed:            true,
		Default:             int64default.StaticInt64(0),
		MarkdownDescription: "The DNS update Time to Live (TTL) value of a DHCP network object. The TTL is a 32-bit unsigned integer that represents the duration, in seconds, for which the update is cached. Zero indicates that the update is not cached.",
	},
	"delete_reason": schema.StringAttribute{
		Optional: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "The reason for deleting the RIR registration request.",
	},
	"disable": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Determines whether a network is disabled or not. When this is set to False, the network is enabled.",
	},
	"discovered_bridge_domain": schema.StringAttribute{
		Optional: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "Discovered bridge domain.",
	},
	"discovered_tenant": schema.StringAttribute{
		Optional: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "Discovered tenant.",
	},
	"discovery_basic_poll_settings": schema.SingleNestedAttribute{
		Attributes:          Ipv6networkDiscoveryBasicPollSettingsResourceSchemaAttributes,
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The discovery basic poll settings for this network.",
	},
	"discovery_blackout_setting": schema.SingleNestedAttribute{
		Attributes:          Ipv6networkDiscoveryBlackoutSettingResourceSchemaAttributes,
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The discovery blackout setting for this network.",
	},
	"discovery_member": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "The member that will run discovery for this network.",
	},
	"domain_name": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "Use this method to set or retrieve the domain_name value of a DHCP IPv6 Network object.",
	},
	"domain_name_servers": schema.ListAttribute{
		ElementType: types.StringType,
		Optional:    true,
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
			listvalidator.ValueStringsAre(customvalidator.IsValidIPv6Address()),
		},
		MarkdownDescription: "Use this method to set or retrieve the dynamic DNS updates flag of a DHCP IPv6 Network object. The DHCP server can send DDNS updates to DNS servers in the same Grid and to external DNS servers. This setting overrides the member level settings.",
	},
	"enable_ddns": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "The dynamic DNS updates flag of a DHCP IPv6 network object. If set to True, the DHCP server sends DDNS updates to DNS servers in the same Grid, and to external DNS servers.",
	},
	"enable_discovery": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Determines whether a discovery is enabled or not for this network. When this is set to False, the network discovery is disabled.",
	},
	"enable_ifmap_publishing": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Determines if IFMAP publishing is enabled for the network.",
	},
	"enable_immediate_discovery": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines if the discovery for the network should be immediately enabled.",
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
	"federated_realms": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: Ipv6networkFederatedRealmsResourceSchemaAttributes,
		},
		Optional: true,
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
		},
		MarkdownDescription: "This field contains the federated realms associated to this network",
	},
	"logic_filter_rules": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: Ipv6networkLogicFilterRulesResourceSchemaAttributes,
		},
		Optional: true,
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
		},
		MarkdownDescription: "This field contains the logic filters to be applied on this IPv6 network. This list corresponds to the match rules that are written to the DHCPv6 configuration file.",
	},
	"members": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: Ipv6networkMembersResourceSchemaAttributes,
		},
		Optional: true,
		Computed: true,
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
		},
		MarkdownDescription: "A list of members servers that serve DHCP for the network. All members in the array must be of the same type. The struct type must be indicated in each element, by setting the \"_struct\" member to the struct type.",
	},
	"mgm_private": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "This field controls whether this object is synchronized with the Multi-Grid Master. If this field is set to True, objects are not synchronized.",
	},
	"network": schema.StringAttribute{
		Optional:   true,
		Computed:   true,
		CustomType: cidrtypes.IPv6PrefixType{},
		PlanModifiers: []planmodifier.String{
			immutable.ImmutableString(),
		},
		Validators: []validator.String{
			stringvalidator.ExactlyOneOf(
				path.MatchRelative().AtParent().AtName("dynamic_allocation"),
			),
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "The network address in IPv6 Address/CIDR format. For regular expression searches, only the IPv6 Address portion is supported. Searches for the CIDR portion is always an exact match. For example, both network containers 16::0/28 and 26::0/24 are matched by expression '.6' and only 26::0/24 is matched by '.6/24'.",
	},
	"network_view": schema.StringAttribute{
		Default:  stringdefault.StaticString("default"),
		Optional: true,
		Computed: true,
		PlanModifiers: []planmodifier.String{
			immutable.ImmutableString(),
		},
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "The name of the network view in which this network resides.",
	},
	"options": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: Ipv6networkOptionsResourceSchemaAttributes,
		},
		Optional: true,
		Computed: true,
		Default:  listdefault.StaticValue(types.ListValueMust(types.ObjectType{AttrTypes: Ipv6networkOptionsAttrTypes}, []attr.Value{})),
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
		},
		MarkdownDescription: "An array of DHCP option dhcpoption structs that lists the DHCP options associated with the object.",
	},
	"port_control_blackout_setting": schema.SingleNestedAttribute{
		Attributes:          Ipv6networkPortControlBlackoutSettingResourceSchemaAttributes,
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The port control blackout setting for this network.",
	},
	"preferred_lifetime": schema.Int64Attribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Use this method to set or retrieve the preferred lifetime value of a DHCP IPv6 Network object.",
	},
	"recycle_leases": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(true),
		MarkdownDescription: "If the field is set to True, the leases are kept in the Recycle Bin until one week after expiration. Otherwise, the leases are permanently deleted.",
	},
	"restart_if_needed": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Restarts the member service.",
	},
	"rir_organization": schema.StringAttribute{
		Optional: true,
		PlanModifiers: []planmodifier.String{
			immutable.ImmutableString(),
		},
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "The RIR organization associated with the IPv6 network.",
	},
	"rir_registration_action": schema.StringAttribute{
		Validators: []validator.String{
			stringvalidator.OneOf("NONE", "CREATE", "MODIFY", "DELETE"),
		},
		Optional:            true,
		MarkdownDescription: "The RIR registration action.",
	},
	"rir_registration_status": schema.StringAttribute{
		Default: stringdefault.StaticString("NOT_REGISTERED"),
		Validators: []validator.String{
			stringvalidator.OneOf("REGISTERED", "NOT_REGISTERED"),
		},
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The registration status of the IPv6 network in RIR.",
	},
	"same_port_control_discovery_blackout": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "If the field is set to True, the discovery blackout setting will be used for port control blackout setting.",
	},
	"send_rir_request": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines whether to send the RIR registration request.",
	},
	"subscribe_settings": schema.SingleNestedAttribute{
		Attributes:          Ipv6networkSubscribeSettingsResourceSchemaAttributes,
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The DHCP IPv6 Network Cisco ISE subscribe settings.",
	},
	"unmanaged": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Determines whether the DHCP IPv6 Network is unmanaged or not.",
	},
	"update_dns_on_lease_renewal": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "This field controls whether the DHCP server updates DNS when a DHCP lease is renewed.",
	},
	"valid_lifetime": schema.Int64Attribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Use this method to set or retrieve the valid lifetime value of a DHCP IPv6 Network object.",
	},
	"vlans": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: Ipv6networkVlansResourceSchemaAttributes,
		},
		Optional: true,
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
		},
		MarkdownDescription: "List of VLANs assigned to Network.",
	},
	"zone_associations": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: Ipv6networkZoneAssociationsResourceSchemaAttributes,
		},
		Optional: true,
		Computed: true,
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
		},
		MarkdownDescription: "The list of zones associated with this network.",
	},
	"dynamic_allocation": schema.SingleNestedAttribute{
		Attributes:          dynamicallocation.NextAvailableNetworkResourceSchemaAttributes,
		Optional:            true,
		MarkdownDescription: "Dynamically allocate the network using the NIOS next_available_network function call. Mutually exclusive with the static value field.",
	},
}

var Ipv6networkResourceUddiSchemaAttributes = map[string]schema.Attribute{
	"address": schema.StringAttribute{
		Optional:   true,
		Computed:   true,
		CustomType: iptypes.IPv6AddressType{},
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.RequiresReplaceIfConfigured(),
			stringplanmodifier.UseStateForUnknown(),
		},
		MarkdownDescription: "The address of the subnet in the form “a.b.c.d/n” where the “/n” may be omitted. In this case, the CIDR value must be defined in the _cidr_ field. When reading, the _address_ field is always in the form “a.b.c.d”.",
	},
	"asm_config": schema.SingleNestedAttribute{
		Attributes: ASMConfigResourceSchemaAttributes,
		Optional:   true,
		Computed:   true,
		Default: objectdefault.StaticValue(types.ObjectValueMust(ASMConfigAttrTypes, map[string]attr.Value{
			"asm_threshold":       types.Int64Value(90),
			"enable":              types.BoolValue(true),
			"enable_notification": types.BoolValue(true),
			"forecast_period":     types.Int64Value(14),
			"growth_factor":       types.Int64Value(20),
			"growth_type":         types.StringValue("percent"),
			"history":             types.Int64Value(30),
			"min_total":           types.Int64Value(10),
			"min_unused":          types.Int64Value(10),
			"reenable_date":       timetypes.NewRFC3339ValueMust("1970-01-01T00:00:00Z"),
		})),
		MarkdownDescription: "The __ASMConfig__ object represents Automated Scope Management configuration.",
	},
	"cidr": schema.Int64Attribute{
		Required:            true,
		MarkdownDescription: "The CIDR of the subnet. This is required if _address_ does not include CIDR.",
	},
	"comment": schema.StringAttribute{
		Default:             stringdefault.StaticString(""),
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The description for the subnet. May contain 0 to 1024 characters. Can include UTF-8.",
	},
	"config_profiles": schema.ListAttribute{
		ElementType: types.StringType,
		Optional:    true,
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
		},
		MarkdownDescription: "The resource identifier.",
	},
	"ddns_client_update": schema.StringAttribute{
		Default:             stringdefault.StaticString("client"),
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Controls who does the DDNS updates.  Valid values are: * _client_: DHCP server updates DNS if requested by client. * _server_: DHCP server always updates DNS, overriding an update request from the client, unless the client requests no updates. * _ignore_: DHCP server always updates DNS, even if the client says not to. * _over_client_update_: Same as _server_. DHCP server always updates DNS, overriding an update request from the client, unless the client requests no updates. * _over_no_update_: DHCP server updates DNS even if the client requests that no updates be done. If the client requests to do the update, DHCP server allows it.  Defaults to _client_.",
	},
	"ddns_conflict_resolution_mode": schema.StringAttribute{
		Default:             stringdefault.StaticString("check_with_dhcid"),
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The mode used for resolving conflicts while performing DDNS updates.  Valid values are: * _check_with_dhcid_: It includes adding a DHCID record and checking that record via conflict detection as per RFC 4703. * _no_check_with_dhcid_: This will ignore conflict detection but add a DHCID record when creating/updating an entry. * _check_exists_with_dhcid_: This will check if there is an existing DHCID record but does not verify the value of the record matches the update. This will also update the DHCID record for the entry. * _no_check_without_dhcid_: This ignores conflict detection and will not add a DHCID record when creating/updating a DDNS entry.  Defaults to _check_with_dhcid_.",
	},
	"ddns_domain": schema.StringAttribute{
		Default:             stringdefault.StaticString(""),
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The domain suffix for DDNS updates. FQDN, may be empty.  Defaults to empty.",
	},
	"ddns_generate_name": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Indicates if DDNS needs to generate a hostname when not supplied by the client.  Defaults to _false_.",
	},
	"ddns_generated_prefix": schema.StringAttribute{
		Default:             stringdefault.StaticString("myhost"),
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The prefix used in the generation of an FQDN.  When generating a name, DHCP server will construct the name in the format: [ddns-generated-prefix]-[address-text].[ddns-qualifying-suffix]. where address-text is simply the lease IP address converted to a hyphenated string.  Defaults to \"myhost\".",
	},
	"ddns_send_updates": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(true),
		MarkdownDescription: "Determines if DDNS updates are enabled at the subnet level. Defaults to _true_.",
	},
	"ddns_ttl_percent": schema.Float64Attribute{
		Optional:            true,
		MarkdownDescription: "DDNS TTL value - to be calculated as a simple percentage of the lease's lifetime, using the parameter's value as the percentage. It is specified as a percentage (e.g. 25, 75). Defaults to unspecified.",
	},
	"ddns_update_on_renew": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Instructs the DHCP server to always update the DNS information when a lease is renewed even if its DNS information has not changed.  Defaults to _false_.",
	},
	"ddns_use_conflict_resolution": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(true),
		MarkdownDescription: "When true, DHCP server will apply conflict resolution, as described in RFC 4703, when attempting to fulfill the update request.  When false, DHCP server will simply attempt to update the DNS entries per the request, regardless of whether or not they conflict with existing entries owned by other DHCP4 clients.  Defaults to _true_.",
	},
	"dhcp_config": schema.SingleNestedAttribute{
		Attributes: Ipv6networkDhcpConfigResourceSchemaAttributes,
		Optional:   true,
		Computed:   true,
		Default: objectdefault.StaticValue(types.ObjectValueMust(Ipv6networkDhcpConfigAttrTypes, map[string]attr.Value{
			"allow_unknown_v6":        types.BoolValue(true),
			"authoritative_dhcp":      types.BoolValue(false),
			"filters_large_selection": types.ListNull(types.StringType),
			"filters_v6":              types.ListNull(types.StringType),
			"ignore_client_uid":       types.BoolValue(false),
			"ignore_list":             types.ListNull(types.ObjectType{AttrTypes: IgnoreItemAttrTypes}),
			"lease_time_v6":           types.Int64Value(3600),
		})),
		MarkdownDescription: "A DHCP Config object (_dhcp/dhcp_config_) represents a shared DHCP configuration that controls how leases are issued.",
	},
	"dhcp_host": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The resource identifier for the DHCP Host associated with this subnet. Omit or set to `null` to inherit from the parent address block (if applicable). Set to empty string (`\"\"`) to explicitly unset the DHCP host. Provide a resource ID to assign a specific DHCP host.",
	},
	"dhcp_options": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: OptionItemResourceSchemaAttributes,
		},
		CustomType: internaltypes.UnorderedList{ListType: types.ListType{ElemType: types.ObjectType{AttrTypes: OptionItemAttrTypes}}},
		Optional:   true,
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
		},
		MarkdownDescription: "The DHCP options of the subnet. This can either be a specific option or a group of options.",
	},
	"disable_dhcp": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Optional. _true_ to disable object. A disabled object is effectively non-existent when generating configuration.  Defaults to _false_.",
	},
	"external_keys": schema.MapAttribute{
		ElementType:         types.StringType,
		Optional:            true,
		MarkdownDescription: "The external keys (source key) for this subnet in JSON format.",
	},
	"federated_realms": schema.ListAttribute{
		ElementType: types.StringType,
		Optional:    true,
		CustomType:  internaltypes.UnorderedListOfStringType,
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
		},
		MarkdownDescription: "Reserved for future use.",
	},
	"hostname_rewrite_char": schema.StringAttribute{
		Default:             stringdefault.StaticString("-"),
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The character to replace non-matching characters with, when hostname rewrite is enabled.  Any single ASCII character or no character if the invalid characters should be removed without replacement.  Defaults to \"-\".",
	},
	"hostname_rewrite_enabled": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Indicates if client supplied hostnames will be rewritten prior to DDNS update by replacing every character that does not match _hostname_rewrite_regex_ by _hostname_rewrite_char_.  Defaults to _false_.",
	},
	"hostname_rewrite_regex": schema.StringAttribute{
		Default:             stringdefault.StaticString("[^a-zA-Z0-9_.]"),
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The regex bracket expression to match valid characters.  Must begin with \"[\" and end with \"]\" and be a compilable POSIX regex.  Defaults to \"[^a-zA-Z0-9_.]\".",
	},
	"inheritance_sources": schema.SingleNestedAttribute{
		Attributes: DHCPInheritanceResourceSchemaAttributes,
		Optional:   true,
		Computed:   true,
		PlanModifiers: []planmodifier.Object{
			objectplanmodifier.UseStateForUnknown(),
		},
		MarkdownDescription: "The __DHCPInheritance__ object specifies how the _dhcp_config_, _dhcp_options_ and _asm_config_ configuration fields are inherited from the parent object.",
	},
	"name": schema.StringAttribute{
		Default:             stringdefault.StaticString(""),
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The name of the subnet. May contain 1 to 256 characters. Can include UTF-8.",
	},
	"rebind_time": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "The lease rebind time (T2) in seconds.",
	},
	"renew_time": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "The lease renew time (T1) in seconds.",
	},
	"space": schema.StringAttribute{
		Required: true,
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.RequiresReplaceIfConfigured(),
		},
		MarkdownDescription: "The resource identifier.",
	},
	"tags": schema.MapAttribute{
		Optional:    true,
		Computed:    true,
		ElementType: types.StringType,
		Default:     mapdefault.StaticValue(types.MapNull(types.StringType)),
		Validators: []validator.Map{
			mapvalidator.SizeAtLeast(1),
		},
		MarkdownDescription: "The tags for the subnet in JSON format.",
	},
	"tags_all": schema.MapAttribute{
		Computed:            true,
		ElementType:         types.StringType,
		MarkdownDescription: "All tags including inherited values.",
	},
}

// Expand converts the TF model to the infoblox core model
func (m *Ipv6networkModel) Expand(ctx context.Context, diags *diag.Diagnostics, isCreate bool) *coremodel.Ipv6network {
	if m == nil {
		return nil
	}

	obj := &coremodel.Ipv6network{}

	// Expand NIOS nested attribute (returns nil if not present)
	niosModel := flex.ExpandNestedObject[NIOSIpv6networkModel](ctx, m.NIOS, diags)
	if niosModel != nil {
		obj.NIOS = niosModel.Expand(ctx, diags, isCreate)
	}

	// Expand UDDI nested attribute (returns nil if not present)
	uddiModel := flex.ExpandNestedObject[UDDIIpv6networkModel](ctx, m.UDDI, diags)
	if uddiModel != nil {
		obj.UDDI = uddiModel.Expand(ctx, diags, isCreate)
	}

	return obj
}

// Expand converts the NIOS TF model to the core model.
func (m *NIOSIpv6networkModel) Expand(ctx context.Context, diags *diag.Diagnostics, isCreate bool) *coremodel.NIOSIpv6networkExt {
	ext := &coremodel.NIOSIpv6networkExt{
		CloudInfo:                        ExpandIpv6networkCloudInfo(ctx, m.CloudInfo, diags),
		Comment:                          flex.ExpandStringPointerNullAsEmpty(m.Comment),
		DdnsDomainname:                   flex.ExpandStringPointerNullAsEmpty(m.DdnsDomainname),
		DdnsEnableOptionFqdn:             flex.ExpandBoolPointer(m.DdnsEnableOptionFqdn),
		DdnsGenerateHostname:             flex.ExpandBoolPointer(m.DdnsGenerateHostname),
		DdnsServerAlwaysUpdates:          flex.ExpandBoolPointer(m.DdnsServerAlwaysUpdates),
		DdnsTtl:                          flex.ExpandInt64Pointer(m.DdnsTtl),
		DeleteReason:                     flex.ExpandStringPointerNullAsEmpty(m.DeleteReason),
		Disable:                          flex.ExpandBoolPointer(m.Disable),
		DiscoveredBridgeDomain:           flex.ExpandStringPointerNullAsEmpty(m.DiscoveredBridgeDomain),
		DiscoveredTenant:                 flex.ExpandStringPointerNullAsEmpty(m.DiscoveredTenant),
		DiscoveryBasicPollSettings:       ExpandIpv6networkDiscoveryBasicPollSettings(ctx, m.DiscoveryBasicPollSettings, diags),
		DiscoveryBlackoutSetting:         ExpandIpv6networkDiscoveryBlackoutSetting(ctx, m.DiscoveryBlackoutSetting, diags),
		DiscoveryMember:                  flex.ExpandStringPointer(m.DiscoveryMember),
		DomainName:                       flex.ExpandStringPointerNullAsEmpty(m.DomainName),
		DomainNameServers:                flex.ExpandFrameworkListString(ctx, m.DomainNameServers, diags),
		EnableDdns:                       flex.ExpandBoolPointer(m.EnableDdns),
		EnableDiscovery:                  flex.ExpandBoolPointer(m.EnableDiscovery),
		EnableIfmapPublishing:            flex.ExpandBoolPointer(m.EnableIfmapPublishing),
		EnableImmediateDiscovery:         flex.ExpandBoolPointer(m.EnableImmediateDiscovery),
		ExtAttrs:                         flex.ExpandMapStringAny(ctx, m.ExtAttrs, diags),
		FederatedRealms:                  flex.ExpandFrameworkListNestedBlock(ctx, m.FederatedRealms, diags, ExpandIpv6networkFederatedRealms),
		LogicFilterRules:                 flex.ExpandFrameworkListNestedBlock(ctx, m.LogicFilterRules, diags, ExpandIpv6networkLogicFilterRules),
		Members:                          flex.ExpandFrameworkListNestedBlock(ctx, m.Members, diags, ExpandIpv6networkMembers),
		MgmPrivate:                       flex.ExpandBoolPointer(m.MgmPrivate),
		Options:                          flex.ExpandFrameworkListNestedBlock(ctx, m.Options, diags, ExpandIpv6networkOptions),
		PortControlBlackoutSetting:       ExpandIpv6networkPortControlBlackoutSetting(ctx, m.PortControlBlackoutSetting, diags),
		PreferredLifetime:                flex.ExpandInt64Pointer(m.PreferredLifetime),
		RecycleLeases:                    flex.ExpandBoolPointer(m.RecycleLeases),
		RestartIfNeeded:                  flex.ExpandBoolPointer(m.RestartIfNeeded),
		RirOrganization:                  flex.ExpandStringPointer(m.RirOrganization),
		RirRegistrationAction:            flex.ExpandStringPointer(m.RirRegistrationAction),
		RirRegistrationStatus:            flex.ExpandStringPointerNullAsEmpty(m.RirRegistrationStatus),
		SamePortControlDiscoveryBlackout: flex.ExpandBoolPointer(m.SamePortControlDiscoveryBlackout),
		SendRirRequest:                   flex.ExpandBoolPointer(m.SendRirRequest),
		SubscribeSettings:                ExpandIpv6networkSubscribeSettings(ctx, m.SubscribeSettings, diags),
		Unmanaged:                        flex.ExpandBoolPointer(m.Unmanaged),
		UpdateDnsOnLeaseRenewal:          flex.ExpandBoolPointer(m.UpdateDnsOnLeaseRenewal),
		ValidLifetime:                    flex.ExpandInt64Pointer(m.ValidLifetime),
		Vlans:                            flex.ExpandFrameworkListNestedBlock(ctx, m.Vlans, diags, ExpandIpv6networkVlans),
		ZoneAssociations:                 flex.ExpandFrameworkListNestedBlock(ctx, m.ZoneAssociations, diags, ExpandIpv6networkZoneAssociations),
	}
	if isCreate {
		ext.AutoCreateReversezone = flex.ExpandBoolPointer(m.AutoCreateReversezone)
		ext.Network = flex.ExpandIPv6Prefix(m.Network)
		ext.NetworkView = flex.ExpandStringPointerNullAsEmpty(m.NetworkView)
		ext.FuncCall = BuildIpv6networkFuncCall(ctx, m.DynamicAllocation, diags)
	}
	return ext
}

// ApplyIpv6networkNIOSUseFlags derives NIOS use flags from the raw config
// value(s) and writes them onto the core model. A flag is true when the user
// set any of its governed value fields in config.
func ApplyIpv6networkNIOSUseFlags(ctx context.Context, config tfsdk.Config, obj *coremodel.Ipv6network, diags *diag.Diagnostics) {
	if obj == nil || obj.NIOS == nil {
		return
	}
	obj.NIOS.UseBlackoutSetting = flex.DeriveUseFlag(ctx, config, diags, path.Root("nios").AtName("discovery_blackout_setting"), path.Root("nios").AtName("port_control_blackout_setting"), path.Root("nios").AtName("same_port_control_discovery_blackout"))
	obj.NIOS.UseDdnsDomainname = flex.DeriveUseFlag(ctx, config, diags, path.Root("nios").AtName("ddns_domainname"))
	obj.NIOS.UseDdnsEnableOptionFqdn = flex.DeriveUseFlag(ctx, config, diags, path.Root("nios").AtName("ddns_enable_option_fqdn"))
	obj.NIOS.UseDdnsGenerateHostname = flex.DeriveUseFlag(ctx, config, diags, path.Root("nios").AtName("ddns_generate_hostname"))
	obj.NIOS.UseDdnsTtl = flex.DeriveUseFlag(ctx, config, diags, path.Root("nios").AtName("ddns_ttl"))
	obj.NIOS.UseDiscoveryBasicPollingSettings = flex.DeriveUseFlag(ctx, config, diags, path.Root("nios").AtName("discovery_basic_poll_settings"))
	obj.NIOS.UseDomainName = flex.DeriveUseFlag(ctx, config, diags, path.Root("nios").AtName("domain_name"))
	obj.NIOS.UseDomainNameServers = flex.DeriveUseFlag(ctx, config, diags, path.Root("nios").AtName("domain_name_servers"))
	obj.NIOS.UseEnableDdns = flex.DeriveUseFlag(ctx, config, diags, path.Root("nios").AtName("enable_ddns"))
	obj.NIOS.UseEnableDiscovery = flex.DeriveUseFlag(ctx, config, diags, path.Root("nios").AtName("discovery_member"), path.Root("nios").AtName("enable_discovery"))
	obj.NIOS.UseEnableIfmapPublishing = flex.DeriveUseFlag(ctx, config, diags, path.Root("nios").AtName("enable_ifmap_publishing"))
	obj.NIOS.UseLogicFilterRules = flex.DeriveUseFlag(ctx, config, diags, path.Root("nios").AtName("logic_filter_rules"))
	obj.NIOS.UseMgmPrivate = flex.DeriveUseFlag(ctx, config, diags, path.Root("nios").AtName("mgm_private"))
	obj.NIOS.UseOptions = flex.DeriveUseFlag(ctx, config, diags, path.Root("nios").AtName("options"))
	obj.NIOS.UsePreferredLifetime = flex.DeriveUseFlag(ctx, config, diags, path.Root("nios").AtName("preferred_lifetime"))
	obj.NIOS.UseRecycleLeases = flex.DeriveUseFlag(ctx, config, diags, path.Root("nios").AtName("recycle_leases"))
	obj.NIOS.UseSubscribeSettings = flex.DeriveUseFlag(ctx, config, diags, path.Root("nios").AtName("subscribe_settings"))
	obj.NIOS.UseUpdateDnsOnLeaseRenewal = flex.DeriveUseFlag(ctx, config, diags, path.Root("nios").AtName("update_dns_on_lease_renewal"))
	obj.NIOS.UseValidLifetime = flex.DeriveUseFlag(ctx, config, diags, path.Root("nios").AtName("valid_lifetime"))
	obj.NIOS.UseZoneAssociations = flex.DeriveUseFlag(ctx, config, diags, path.Root("nios").AtName("zone_associations"))
}

// Expand converts the UDDI TF model to the core model.
func (m *UDDIIpv6networkModel) Expand(ctx context.Context, diags *diag.Diagnostics, isCreate bool) *coremodel.UDDIIpv6networkExt {
	ext := &coremodel.UDDIIpv6networkExt{
		AsmConfig:                  ExpandASMConfig(ctx, m.AsmConfig, diags),
		Cidr:                       flex.ExpandInt64Pointer(m.Cidr),
		Comment:                    flex.ExpandStringPointer(m.Comment),
		ConfigProfiles:             flex.ExpandFrameworkListString(ctx, m.ConfigProfiles, diags),
		DdnsClientUpdate:           flex.ExpandStringPointer(m.DdnsClientUpdate),
		DdnsConflictResolutionMode: flex.ExpandStringPointer(m.DdnsConflictResolutionMode),
		DdnsDomain:                 flex.ExpandStringPointer(m.DdnsDomain),
		DdnsGenerateName:           flex.ExpandBoolPointer(m.DdnsGenerateName),
		DdnsGeneratedPrefix:        flex.ExpandStringPointer(m.DdnsGeneratedPrefix),
		DdnsSendUpdates:            flex.ExpandBoolPointer(m.DdnsSendUpdates),
		DdnsTtlPercent:             flex.ExpandFloat32Pointer(m.DdnsTtlPercent),
		DdnsUpdateOnRenew:          flex.ExpandBoolPointer(m.DdnsUpdateOnRenew),
		DdnsUseConflictResolution:  flex.ExpandBoolPointer(m.DdnsUseConflictResolution),
		DhcpConfig:                 ExpandIpv6networkDhcpConfig(ctx, m.DhcpConfig, diags),
		DhcpHost:                   flex.ExpandStringPointer(m.DhcpHost),
		DhcpOptions:                flex.ExpandFrameworkListNestedBlock(ctx, m.DhcpOptions, diags, ExpandOptionItem),
		DisableDhcp:                flex.ExpandBoolPointer(m.DisableDhcp),
		ExternalKeys:               flex.ExpandMapStringAny(ctx, m.ExternalKeys, diags),
		FederatedRealms:            flex.ExpandFrameworkListString(ctx, m.FederatedRealms, diags),
		HostnameRewriteChar:        flex.ExpandStringPointer(m.HostnameRewriteChar),
		HostnameRewriteEnabled:     flex.ExpandBoolPointer(m.HostnameRewriteEnabled),
		HostnameRewriteRegex:       flex.ExpandStringPointer(m.HostnameRewriteRegex),
		InheritanceSources:         ExpandDHCPInheritance(ctx, m.InheritanceSources, diags),
		Name:                       flex.ExpandStringPointer(m.Name),
		RebindTime:                 flex.ExpandInt64Pointer(m.RebindTime),
		RenewTime:                  flex.ExpandInt64Pointer(m.RenewTime),
		Tags:                       flex.ExpandMapStringAny(ctx, m.Tags, diags),
	}
	if isCreate {
		ext.Address = flex.ExpandIPv6Address(m.Address)
		ext.Space = flex.ExpandStringPointer(m.Space)
	}
	return ext
}

// Flatten populates the TF model from a core response.
func (m *Ipv6networkModel) Flatten(ctx context.Context, resp *coremodel.Ipv6network, diags *diag.Diagnostics) {
	if resp == nil {
		return
	}

	m.Id = flex.FlattenStringPointer(resp.Id)

	// Extract existing NIOS model, flatten API response onto it, convert back
	niosModel := flex.ExpandNestedObject[NIOSIpv6networkModel](ctx, m.NIOS, diags)
	if niosModel == nil {
		niosModel = &NIOSIpv6networkModel{}
	}
	plannedNIOS := flex.ExpandNestedObject[NIOSIpv6networkModel](ctx, m.NIOS, diags)
	niosModel.Flatten(ctx, resp.NIOS, diags)
	if resp.NIOS != nil {
		PostFlattenIpv6networkNIOS(ctx, plannedNIOS, niosModel, diags)
		m.NIOS = flex.FlattenNestedObject(ctx, niosModel, NIOSIpv6networkAttrTypes, diags)
	} else {
		m.NIOS = types.ObjectNull(NIOSIpv6networkAttrTypes)
	}

	// Extract existing UDDI model, flatten API response onto it, convert back
	uddiModel := flex.ExpandNestedObject[UDDIIpv6networkModel](ctx, m.UDDI, diags)
	if uddiModel == nil {
		uddiModel = &UDDIIpv6networkModel{}
	}
	uddiModel.Flatten(ctx, resp.UDDI, diags)
	if resp.UDDI != nil {
		m.UDDI = flex.FlattenNestedObject(ctx, uddiModel, UDDIIpv6networkAttrTypes, diags)
	} else {
		m.UDDI = types.ObjectNull(UDDIIpv6networkAttrTypes)
	}
}

// Flatten merges API response onto existing NIOS model.
func (m *NIOSIpv6networkModel) Flatten(ctx context.Context, from *coremodel.NIOSIpv6networkExt, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	planExtAttrs := m.ExtAttrs
	if planExtAttrs.IsUnknown() {
		planExtAttrs = types.MapNull(types.StringType)
	}
	m.CloudInfo = FlattenIpv6networkCloudInfo(ctx, from.CloudInfo, diags)
	m.Comment = flex.FlattenStringPointerEmptyAsNull(from.Comment)
	m.DdnsDomainname = flex.FlattenStringPointerEmptyAsNull(from.DdnsDomainname)
	m.DdnsEnableOptionFqdn = flex.FlattenBoolPointer(from.DdnsEnableOptionFqdn)
	m.DdnsGenerateHostname = flex.FlattenBoolPointer(from.DdnsGenerateHostname)
	m.DdnsServerAlwaysUpdates = flex.FlattenBoolPointer(from.DdnsServerAlwaysUpdates)
	m.DdnsTtl = flex.FlattenInt64Pointer(from.DdnsTtl)
	m.Disable = flex.FlattenBoolPointer(from.Disable)
	m.DiscoveredBridgeDomain = flex.FlattenStringPointerEmptyAsNull(from.DiscoveredBridgeDomain)
	m.DiscoveredTenant = flex.FlattenStringPointerEmptyAsNull(from.DiscoveredTenant)
	m.DiscoveryBasicPollSettings = FlattenIpv6networkDiscoveryBasicPollSettings(ctx, from.DiscoveryBasicPollSettings, diags)
	m.DiscoveryBlackoutSetting = FlattenIpv6networkDiscoveryBlackoutSetting(ctx, from.DiscoveryBlackoutSetting, diags)
	m.DiscoveryMember = flex.FlattenStringPointerEmptyAsNull(from.DiscoveryMember)
	m.DomainName = flex.FlattenStringPointerEmptyAsNull(from.DomainName)
	m.DomainNameServers = flex.FlattenFrameworkListString(ctx, from.DomainNameServers, diags)
	m.EnableDdns = flex.FlattenBoolPointer(from.EnableDdns)
	m.EnableDiscovery = flex.FlattenBoolPointer(from.EnableDiscovery)
	m.EnableIfmapPublishing = flex.FlattenBoolPointer(from.EnableIfmapPublishing)
	m.ExtAttrs, m.ExtAttrsAll = flex.FlattenEAs(planExtAttrs, from.ExtAttrs)
	m.FederatedRealms = flex.FlattenFrameworkListNestedBlock(ctx, from.FederatedRealms, Ipv6networkFederatedRealmsAttrTypes, diags, FlattenIpv6networkFederatedRealms)
	m.LogicFilterRules = flex.FlattenFrameworkListNestedBlock(ctx, from.LogicFilterRules, Ipv6networkLogicFilterRulesAttrTypes, diags, FlattenIpv6networkLogicFilterRules)
	m.Members = flex.FlattenFrameworkListNestedBlock(ctx, from.Members, Ipv6networkMembersAttrTypes, diags, FlattenIpv6networkMembers)
	m.MgmPrivate = flex.FlattenBoolPointer(from.MgmPrivate)
	m.Network = flex.FlattenIPv6Prefix(from.Network)
	m.NetworkView = flex.FlattenStringPointerEmptyAsNull(from.NetworkView)
	m.Options = flex.FlattenFrameworkListNestedBlock(ctx, from.Options, Ipv6networkOptionsAttrTypes, diags, FlattenIpv6networkOptions)
	m.PortControlBlackoutSetting = FlattenIpv6networkPortControlBlackoutSetting(ctx, from.PortControlBlackoutSetting, diags)
	m.PreferredLifetime = flex.FlattenInt64Pointer(from.PreferredLifetime)
	m.RecycleLeases = flex.FlattenBoolPointer(from.RecycleLeases)
	m.RirOrganization = flex.FlattenStringPointerEmptyAsNull(from.RirOrganization)
	m.RirRegistrationStatus = flex.FlattenStringPointerEmptyAsNull(from.RirRegistrationStatus)
	m.SamePortControlDiscoveryBlackout = flex.FlattenBoolPointer(from.SamePortControlDiscoveryBlackout)
	m.SubscribeSettings = FlattenIpv6networkSubscribeSettings(ctx, from.SubscribeSettings, diags)
	m.Unmanaged = flex.FlattenBoolPointer(from.Unmanaged)
	m.UpdateDnsOnLeaseRenewal = flex.FlattenBoolPointer(from.UpdateDnsOnLeaseRenewal)
	m.ValidLifetime = flex.FlattenInt64Pointer(from.ValidLifetime)
	m.Vlans = flex.FlattenFrameworkListNestedBlock(ctx, from.Vlans, Ipv6networkVlansAttrTypes, diags, FlattenIpv6networkVlans)
	m.ZoneAssociations = flex.FlattenFrameworkListNestedBlock(ctx, from.ZoneAssociations, Ipv6networkZoneAssociationsAttrTypes, diags, FlattenIpv6networkZoneAssociations)
	if len(m.DynamicAllocation.AttributeTypes(ctx)) == 0 {
		m.DynamicAllocation = types.ObjectNull(dynamicallocation.NextAvailableNetworkAttrTypes)
	}
}

// Flatten merges API response onto existing UDDI model.
func (m *UDDIIpv6networkModel) Flatten(ctx context.Context, from *coremodel.UDDIIpv6networkExt, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.Address = flex.FlattenIPv6Address(from.Address)
	m.AsmConfig = FlattenASMConfig(ctx, from.AsmConfig, diags)
	m.Cidr = flex.FlattenInt64Pointer(from.Cidr)
	m.Comment = flex.FlattenStringPointer(from.Comment)
	m.ConfigProfiles = flex.FlattenFrameworkListString(ctx, from.ConfigProfiles, diags)
	m.DdnsClientUpdate = flex.FlattenStringPointer(from.DdnsClientUpdate)
	m.DdnsConflictResolutionMode = flex.FlattenStringPointer(from.DdnsConflictResolutionMode)
	m.DdnsDomain = flex.FlattenStringPointer(from.DdnsDomain)
	m.DdnsGenerateName = flex.FlattenBoolPointer(from.DdnsGenerateName)
	m.DdnsGeneratedPrefix = flex.FlattenStringPointer(from.DdnsGeneratedPrefix)
	m.DdnsSendUpdates = flex.FlattenBoolPointer(from.DdnsSendUpdates)
	m.DdnsTtlPercent = flex.FlattenFloat32PointerZeroAsNull(from.DdnsTtlPercent)
	m.DdnsUpdateOnRenew = flex.FlattenBoolPointer(from.DdnsUpdateOnRenew)
	m.DdnsUseConflictResolution = flex.FlattenBoolPointer(from.DdnsUseConflictResolution)
	m.DhcpConfig = FlattenIpv6networkDhcpConfig(ctx, from.DhcpConfig, diags)
	m.DhcpHost = flex.FlattenStringPointer(from.DhcpHost)
	m.DhcpOptions = flex.FlattenFrameworkUnorderedListNestedBlock(ctx, from.DhcpOptions, OptionItemAttrTypes, diags, FlattenOptionItem)
	m.DisableDhcp = flex.FlattenBoolPointer(from.DisableDhcp)
	m.ExternalKeys = flex.FlattenMapStringAny(ctx, from.ExternalKeys, diags)
	m.FederatedRealms = flex.FlattenFrameworkUnorderedListString(ctx, from.FederatedRealms, diags)
	m.HostnameRewriteChar = flex.FlattenStringPointer(from.HostnameRewriteChar)
	m.HostnameRewriteEnabled = flex.FlattenBoolPointer(from.HostnameRewriteEnabled)
	m.HostnameRewriteRegex = flex.FlattenStringPointer(from.HostnameRewriteRegex)
	m.InheritanceSources = FlattenDHCPInheritance(ctx, from.InheritanceSources, diags)
	m.Name = flex.FlattenStringPointer(from.Name)
	m.RebindTime = flex.FlattenInt64PointerZeroAsNull(from.RebindTime)
	m.RenewTime = flex.FlattenInt64PointerZeroAsNull(from.RenewTime)
	m.Space = flex.FlattenStringPointer(from.Space)
	tagsAll := flex.FlattenMapStringAny(ctx, from.Tags, diags)
	if m.Tags.IsNull() || m.Tags.IsUnknown() {
		m.Tags = tagsAll
	}
	m.TagsAll = tagsAll
}
