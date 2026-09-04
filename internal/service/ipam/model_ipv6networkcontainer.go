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

type Ipv6networkcontainerModel struct {
	Id   types.String `tfsdk:"id"`
	NIOS types.Object `tfsdk:"nios"`
	UDDI types.Object `tfsdk:"uddi"`
}

var Ipv6networkcontainerAttrTypes = map[string]attr.Type{
	"id":   types.StringType,
	"nios": types.ObjectType{AttrTypes: NIOSIpv6networkcontainerAttrTypes},
	"uddi": types.ObjectType{AttrTypes: UDDIIpv6networkcontainerAttrTypes},
}

type NIOSIpv6networkcontainerModel struct {
	AutoCreateReversezone            types.Bool           `tfsdk:"auto_create_reversezone"`
	CloudInfo                        types.Object         `tfsdk:"cloud_info"`
	Comment                          types.String         `tfsdk:"comment"`
	DdnsDomainname                   types.String         `tfsdk:"ddns_domainname"`
	DdnsEnableOptionFqdn             types.Bool           `tfsdk:"ddns_enable_option_fqdn"`
	DdnsGenerateHostname             types.Bool           `tfsdk:"ddns_generate_hostname"`
	DdnsServerAlwaysUpdates          types.Bool           `tfsdk:"ddns_server_always_updates"`
	DdnsTtl                          types.Int64          `tfsdk:"ddns_ttl"`
	DeleteReason                     types.String         `tfsdk:"delete_reason"`
	DiscoveryBasicPollSettings       types.Object         `tfsdk:"discovery_basic_poll_settings"`
	DiscoveryBlackoutSetting         types.Object         `tfsdk:"discovery_blackout_setting"`
	DiscoveryMember                  types.String         `tfsdk:"discovery_member"`
	DomainNameServers                types.List           `tfsdk:"domain_name_servers"`
	EnableDdns                       types.Bool           `tfsdk:"enable_ddns"`
	EnableDiscovery                  types.Bool           `tfsdk:"enable_discovery"`
	EnableImmediateDiscovery         types.Bool           `tfsdk:"enable_immediate_discovery"`
	ExtAttrs                         types.Map            `tfsdk:"ext_attrs"`
	ExtAttrsAll                      types.Map            `tfsdk:"ext_attrs_all"`
	FederatedRealms                  types.List           `tfsdk:"federated_realms"`
	LogicFilterRules                 types.List           `tfsdk:"logic_filter_rules"`
	MgmPrivate                       types.Bool           `tfsdk:"mgm_private"`
	Network                          cidrtypes.IPv6Prefix `tfsdk:"network"`
	NetworkView                      types.String         `tfsdk:"network_view"`
	Options                          types.List           `tfsdk:"options"`
	PortControlBlackoutSetting       types.Object         `tfsdk:"port_control_blackout_setting"`
	PreferredLifetime                types.Int64          `tfsdk:"preferred_lifetime"`
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
	ZoneAssociations                 types.List           `tfsdk:"zone_associations"`
	DynamicAllocation                types.Object         `tfsdk:"dynamic_allocation"`
}

var NIOSIpv6networkcontainerAttrTypes = map[string]attr.Type{
	"auto_create_reversezone":              types.BoolType,
	"cloud_info":                           types.ObjectType{AttrTypes: Ipv6networkcontainerCloudInfoAttrTypes},
	"comment":                              types.StringType,
	"ddns_domainname":                      types.StringType,
	"ddns_enable_option_fqdn":              types.BoolType,
	"ddns_generate_hostname":               types.BoolType,
	"ddns_server_always_updates":           types.BoolType,
	"ddns_ttl":                             types.Int64Type,
	"delete_reason":                        types.StringType,
	"discovery_basic_poll_settings":        types.ObjectType{AttrTypes: Ipv6networkcontainerDiscoveryBasicPollSettingsAttrTypes},
	"discovery_blackout_setting":           types.ObjectType{AttrTypes: Ipv6networkcontainerDiscoveryBlackoutSettingAttrTypes},
	"discovery_member":                     types.StringType,
	"domain_name_servers":                  types.ListType{ElemType: types.StringType},
	"enable_ddns":                          types.BoolType,
	"enable_discovery":                     types.BoolType,
	"enable_immediate_discovery":           types.BoolType,
	"ext_attrs":                            types.MapType{ElemType: types.StringType},
	"ext_attrs_all":                        types.MapType{ElemType: types.StringType},
	"federated_realms":                     types.ListType{ElemType: types.ObjectType{AttrTypes: Ipv6networkcontainerFederatedRealmsAttrTypes}},
	"logic_filter_rules":                   types.ListType{ElemType: types.ObjectType{AttrTypes: Ipv6networkcontainerLogicFilterRulesAttrTypes}},
	"mgm_private":                          types.BoolType,
	"network":                              cidrtypes.IPv6PrefixType{},
	"network_view":                         types.StringType,
	"options":                              types.ListType{ElemType: types.ObjectType{AttrTypes: Ipv6networkcontainerOptionsAttrTypes}},
	"port_control_blackout_setting":        types.ObjectType{AttrTypes: Ipv6networkcontainerPortControlBlackoutSettingAttrTypes},
	"preferred_lifetime":                   types.Int64Type,
	"restart_if_needed":                    types.BoolType,
	"rir_organization":                     types.StringType,
	"rir_registration_action":              types.StringType,
	"rir_registration_status":              types.StringType,
	"same_port_control_discovery_blackout": types.BoolType,
	"send_rir_request":                     types.BoolType,
	"subscribe_settings":                   types.ObjectType{AttrTypes: Ipv6networkcontainerSubscribeSettingsAttrTypes},
	"unmanaged":                            types.BoolType,
	"update_dns_on_lease_renewal":          types.BoolType,
	"valid_lifetime":                       types.Int64Type,
	"zone_associations":                    types.ListType{ElemType: types.ObjectType{AttrTypes: Ipv6networkcontainerZoneAssociationsAttrTypes}},
	"dynamic_allocation":                   types.ObjectType{AttrTypes: dynamicallocation.NextAvailableNetworkAttrTypes},
}

type UDDIIpv6networkcontainerModel struct {
	Address                    iptypes.IPv6Address              `tfsdk:"address"`
	AsmConfig                  types.Object                     `tfsdk:"asm_config"`
	Cidr                       types.Int64                      `tfsdk:"cidr"`
	Comment                    types.String                     `tfsdk:"comment"`
	CompartmentId              types.String                     `tfsdk:"compartment_id"`
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
	DhcpOptions                types.List                       `tfsdk:"dhcp_options"`
	ExternalKeys               types.Map                        `tfsdk:"external_keys"`
	FederatedRealms            internaltypes.UnorderedListValue `tfsdk:"federated_realms"`
	HostnameRewriteChar        types.String                     `tfsdk:"hostname_rewrite_char"`
	HostnameRewriteEnabled     types.Bool                       `tfsdk:"hostname_rewrite_enabled"`
	HostnameRewriteRegex       types.String                     `tfsdk:"hostname_rewrite_regex"`
	InheritanceParent          types.String                     `tfsdk:"inheritance_parent"`
	InheritanceSources         types.Object                     `tfsdk:"inheritance_sources"`
	Name                       types.String                     `tfsdk:"name"`
	Parent                     types.String                     `tfsdk:"parent"`
	Space                      types.String                     `tfsdk:"space"`
	Tags                       types.Map                        `tfsdk:"tags"`
	TagsAll                    types.Map                        `tfsdk:"tags_all"`
	Threshold                  types.Object                     `tfsdk:"threshold"`
}

var UDDIIpv6networkcontainerAttrTypes = map[string]attr.Type{
	"address":                       iptypes.IPv6AddressType{},
	"asm_config":                    types.ObjectType{AttrTypes: ASMConfigAttrTypes},
	"cidr":                          types.Int64Type,
	"comment":                       types.StringType,
	"compartment_id":                types.StringType,
	"ddns_client_update":            types.StringType,
	"ddns_conflict_resolution_mode": types.StringType,
	"ddns_domain":                   types.StringType,
	"ddns_generate_name":            types.BoolType,
	"ddns_generated_prefix":         types.StringType,
	"ddns_send_updates":             types.BoolType,
	"ddns_ttl_percent":              types.Float64Type,
	"ddns_update_on_renew":          types.BoolType,
	"ddns_use_conflict_resolution":  types.BoolType,
	"dhcp_config":                   types.ObjectType{AttrTypes: Ipv6NetworkcontainerDHCPConfigAttrTypes},
	"dhcp_options":                  types.ListType{ElemType: types.ObjectType{AttrTypes: OptionItemAttrTypes}},
	"external_keys":                 types.MapType{ElemType: types.StringType},
	"federated_realms":              internaltypes.UnorderedListOfStringType,
	"hostname_rewrite_char":         types.StringType,
	"hostname_rewrite_enabled":      types.BoolType,
	"hostname_rewrite_regex":        types.StringType,
	"inheritance_parent":            types.StringType,
	"inheritance_sources":           types.ObjectType{AttrTypes: DHCPInheritanceAttrTypes},
	"name":                          types.StringType,
	"parent":                        types.StringType,
	"space":                         types.StringType,
	"tags":                          types.MapType{ElemType: types.StringType},
	"tags_all":                      types.MapType{ElemType: types.StringType},
	"threshold":                     types.ObjectType{AttrTypes: UtilizationThresholdAttrTypes},
}

const (
	Ipv6networkcontainerInheritanceType = "full"
	Ipv6networkcontainerReturnFields    = "cloud_info,comment,ddns_domainname,ddns_enable_option_fqdn,ddns_generate_hostname,ddns_server_always_updates,ddns_ttl,discover_now_status,discovery_basic_poll_settings,discovery_blackout_setting,discovery_engine_type,discovery_member,domain_name_servers,enable_ddns,enable_discovery,endpoint_sources,extattrs,federated_realms,last_rir_registration_update_sent,last_rir_registration_update_status,logic_filter_rules,mgm_private,mgm_private_overridable,ms_ad_user_data,network,network_container,network_view,options,port_control_blackout_setting,preferred_lifetime,rir,rir_organization,rir_registration_status,same_port_control_discovery_blackout,subscribe_settings,unmanaged,update_dns_on_lease_renewal,use_blackout_setting,use_ddns_domainname,use_ddns_enable_option_fqdn,use_ddns_generate_hostname,use_ddns_ttl,use_discovery_basic_polling_settings,use_domain_name_servers,use_enable_ddns,use_enable_discovery,use_logic_filter_rules,use_mgm_private,use_options,use_preferred_lifetime,use_subscribe_settings,use_update_dns_on_lease_renewal,use_valid_lifetime,use_zone_associations,utilization,valid_lifetime,zone_associations"
)

var Ipv6networkcontainerResourceSchemaAttributes = map[string]schema.Attribute{
	"id": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The reference to the object.",
	},
	"nios": schema.SingleNestedAttribute{
		Optional:            true,
		MarkdownDescription: "NIOS backend-specific fields.",
		Attributes:          Ipv6networkcontainerResourceNiosSchemaAttributes,
	},
	"uddi": schema.SingleNestedAttribute{
		Optional:            true,
		MarkdownDescription: "UDDI backend-specific fields.",
		Attributes:          Ipv6networkcontainerResourceUddiSchemaAttributes,
	},
}

var Ipv6networkcontainerResourceNiosSchemaAttributes = map[string]schema.Attribute{
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
		Attributes:          Ipv6networkcontainerCloudInfoResourceSchemaAttributes,
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
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "The dynamic DNS domain name the appliance uses specifically for DDNS updates for this network container.",
	},
	"ddns_enable_option_fqdn": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Use this method to set or retrieve the ddns_enable_option_fqdn flag of a DHCP IPv6 Network Container object. This method controls whether the FQDN option sent by the client is to be used, or if the server can automatically generate the FQDN. This setting overrides the upper-level settings.",
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
		MarkdownDescription: "This field controls whether the DHCP server is allowed to update DNS, regardless of the DHCP client requests. Note that changes for this field take effect only if ddns_enable_option_fqdn is True.",
	},
	"ddns_ttl": schema.Int64Attribute{
		Optional:            true,
		Computed:            true,
		Default:             int64default.StaticInt64(0),
		MarkdownDescription: "The DNS update Time to Live (TTL) value of a DHCP network container object. The TTL is a 32-bit unsigned integer that represents the duration, in seconds, for which the update is cached. Zero indicates that the update is not cached.",
	},
	"delete_reason": schema.StringAttribute{
		Optional: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "The reason for deleting the RIR registration request.",
	},
	"discovery_basic_poll_settings": schema.SingleNestedAttribute{
		Attributes:          Ipv6networkcontainerDiscoveryBasicPollSettingsResourceSchemaAttributes,
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "",
	},
	"discovery_blackout_setting": schema.SingleNestedAttribute{
		Attributes:          Ipv6networkcontainerDiscoveryBlackoutSettingResourceSchemaAttributes,
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "",
	},
	"discovery_member": schema.StringAttribute{
		Optional: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "The member that will run discovery for this network container.",
	},
	"domain_name_servers": schema.ListAttribute{
		ElementType: types.StringType,
		Optional:    true,
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
			listvalidator.ValueStringsAre(customvalidator.IsValidIPv6Address()),
		},
		MarkdownDescription: "Use this method to set or retrieve the dynamic DNS updates flag of a DHCP IPv6 Network Container object. The DHCP server can send DDNS updates to DNS servers in the same Grid and to external DNS servers. This setting overrides the member level settings.",
	},
	"enable_ddns": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "The dynamic DNS updates flag of a DHCP IPv6 network container object. If set to True, the DHCP server sends DDNS updates to DNS servers in the same Grid, and to external DNS servers.",
	},
	"enable_discovery": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Determines whether a discovery is enabled or not for this network container. When this is set to False, the network container discovery is disabled.",
	},
	"enable_immediate_discovery": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines if the discovery for the network container should be immediately enabled.",
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
			Attributes: Ipv6networkcontainerFederatedRealmsResourceSchemaAttributes,
		},
		Optional: true,
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
		},
		MarkdownDescription: "This field contains the federated realms associated to this network container.",
	},
	"logic_filter_rules": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: Ipv6networkcontainerLogicFilterRulesResourceSchemaAttributes,
		},
		Optional: true,
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
		},
		MarkdownDescription: "This field contains the logic filters to be applied on the this network container. This list corresponds to the match rules that are written to the dhcpd configuration file.",
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
			Attributes: Ipv6networkcontainerOptionsResourceSchemaAttributes,
		},
		Optional: true,
		Computed: true,
		Default:  listdefault.StaticValue(types.ListValueMust(types.ObjectType{AttrTypes: Ipv6networkcontainerOptionsAttrTypes}, []attr.Value{})),
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
		},
		MarkdownDescription: "An array of DHCP option dhcpoption structs that lists the DHCP options associated with the object.",
	},
	"port_control_blackout_setting": schema.SingleNestedAttribute{
		Attributes:          Ipv6networkcontainerPortControlBlackoutSettingResourceSchemaAttributes,
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "",
	},
	"preferred_lifetime": schema.Int64Attribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Use this method to set or retrieve the preferred lifetime value of a DHCP IPv6 Network Container object.",
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
		MarkdownDescription: "The RIR organization associated with the IPv6 network container.",
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
		MarkdownDescription: "The registration status of the IPv6 network container in RIR.",
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
		Attributes:          Ipv6networkcontainerSubscribeSettingsResourceSchemaAttributes,
		Optional:            true,
		MarkdownDescription: "",
	},
	"unmanaged": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Determines whether the network container is unmanaged or not.",
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
		MarkdownDescription: "Use this method to set or retrieve the valid lifetime value of a DHCP IPv6 Network Container object.",
	},
	"zone_associations": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: Ipv6networkcontainerZoneAssociationsResourceSchemaAttributes,
		},
		Optional: true,
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
		},
		MarkdownDescription: "The list of zones associated with this network container.",
	},
	"dynamic_allocation": schema.SingleNestedAttribute{
		Attributes:          dynamicallocation.NextAvailableNetworkResourceSchemaAttributes,
		Optional:            true,
		MarkdownDescription: "Dynamically allocate the network using the NIOS next_available_network function call. Mutually exclusive with the static value field.",
	},
}

var Ipv6networkcontainerResourceUddiSchemaAttributes = map[string]schema.Attribute{
	"address": schema.StringAttribute{
		Optional:   true,
		CustomType: iptypes.IPv6AddressType{},
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.RequiresReplaceIfConfigured(),
			stringplanmodifier.UseStateForUnknown(),
		},
		MarkdownDescription: "The address field in form “a.b.c.d/n” where the “/n” may be omitted. In this case, the CIDR value must be defined in the _cidr_ field. When reading, the _address_ field is always in the form “a.b.c.d”.",
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
		MarkdownDescription: "The CIDR of the address block. This is required, if _address_ does not specify it in its input.",
	},
	"comment": schema.StringAttribute{
		Default:             stringdefault.StaticString(""),
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The description for the address block. May contain 0 to 1024 characters. Can include UTF-8.",
	},
	"compartment_id": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The compartment associated with the object. If no compartment is associated with the object, the value defaults to empty.",
	},
	"ddns_client_update": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Controls who does the DDNS updates.  Valid values are: * _client_: DHCP server updates DNS if requested by client. * _server_: DHCP server always updates DNS, overriding an update request from the client, unless the client requests no updates. * _ignore_: DHCP server always updates DNS, even if the client says not to. * _over_client_update_: Same as _server_. DHCP server always updates DNS, overriding an update request from the client, unless the client requests no updates. * _over_no_update_: DHCP server updates DNS even if the client requests that no updates be done. If the client requests to do the update, DHCP server allows it.  Defaults to _client_.",
	},
	"ddns_conflict_resolution_mode": schema.StringAttribute{
		Validators: []validator.String{
			stringvalidator.OneOf("check_with_dhcid", "no_check_with_dhcid", "check_exists_with_dhcid", "no_check_without_dhcid"),
		},
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
		MarkdownDescription: "Determines if DDNS updates are enabled at the address block level. Defaults to _true_.",
	},
	"ddns_ttl_percent": schema.Float64Attribute{
		Optional:            true,
		Computed:            true,
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
		Attributes: Ipv6NetworkcontainerDHCPConfigResourceSchemaAttributes,
		Optional:   true,
		Computed:   true,
		Default: objectdefault.StaticValue(types.ObjectValueMust(Ipv6NetworkcontainerDHCPConfigAttrTypes, map[string]attr.Value{
			"allow_unknown_v6":        types.BoolValue(true),
			"authoritative_dhcp":      types.BoolValue(false),
			"filters_v6":              types.ListNull(types.StringType),
			"filters_large_selection": types.ListNull(types.StringType),
			"ignore_client_uid":       types.BoolValue(false),
			"ignore_list":             internaltypes.NewUnorderedListValueNull(types.ObjectType{AttrTypes: IgnoreItemAttrTypes}),
			"lease_time_v6":           types.Int64Value(3600),
		})),
		MarkdownDescription: "A DHCP Config object (_dhcp/dhcp_config_) represents a shared DHCP configuration that controls how leases are issued.",
	},
	"dhcp_options": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: OptionItemResourceSchemaAttributes,
		},
		Optional: true,
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
		},
		MarkdownDescription: "The list of DHCP options for the address block. May be either a specific option or a group of options.",
	},
	"external_keys": schema.MapAttribute{
		ElementType:         types.StringType,
		Optional:            true,
		MarkdownDescription: "The external keys (source key) for this address block in JSON format.",
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
	"inheritance_parent": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The resource identifier.",
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
		MarkdownDescription: "The name of the address block. May contain 1 to 256 characters. Can include UTF-8.",
	},
	"parent": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The resource identifier.",
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
		MarkdownDescription: "The tags for the address block in JSON format.",
	},
	"tags_all": schema.MapAttribute{
		Computed:            true,
		ElementType:         types.StringType,
		MarkdownDescription: "All tags including inherited values.",
	},
	"threshold": schema.SingleNestedAttribute{
		Attributes:          UtilizationThresholdResourceSchemaAttributes,
		Computed:            true,
		MarkdownDescription: "A __UtilizationThreshold__ object represents IP address utilization threshold settings.",
	},
}

// Expand converts the TF model to the infoblox core model
func (m *Ipv6networkcontainerModel) Expand(ctx context.Context, diags *diag.Diagnostics, isCreate bool) *coremodel.Ipv6networkcontainer {
	if m == nil {
		return nil
	}

	obj := &coremodel.Ipv6networkcontainer{}

	// Expand NIOS nested attribute (returns nil if not present)
	niosModel := flex.ExpandNestedObject[NIOSIpv6networkcontainerModel](ctx, m.NIOS, diags)
	if niosModel != nil {
		obj.NIOS = niosModel.Expand(ctx, diags, isCreate)
	}

	// Expand UDDI nested attribute (returns nil if not present)
	uddiModel := flex.ExpandNestedObject[UDDIIpv6networkcontainerModel](ctx, m.UDDI, diags)
	if uddiModel != nil {
		obj.UDDI = uddiModel.Expand(ctx, diags, isCreate)
	}

	return obj
}

// Expand converts the NIOS TF model to the core model.
func (m *NIOSIpv6networkcontainerModel) Expand(ctx context.Context, diags *diag.Diagnostics, isCreate bool) *coremodel.NIOSIpv6networkcontainerExt {
	ext := &coremodel.NIOSIpv6networkcontainerExt{
		CloudInfo:                        ExpandIpv6networkcontainerCloudInfo(ctx, m.CloudInfo, diags),
		Comment:                          flex.ExpandStringPointerNullAsEmpty(m.Comment),
		DdnsDomainname:                   flex.ExpandStringPointerNullAsEmpty(m.DdnsDomainname),
		DdnsEnableOptionFqdn:             flex.ExpandBoolPointer(m.DdnsEnableOptionFqdn),
		DdnsGenerateHostname:             flex.ExpandBoolPointer(m.DdnsGenerateHostname),
		DdnsServerAlwaysUpdates:          flex.ExpandBoolPointer(m.DdnsServerAlwaysUpdates),
		DdnsTtl:                          flex.ExpandInt64Pointer(m.DdnsTtl),
		DeleteReason:                     flex.ExpandStringPointerNullAsEmpty(m.DeleteReason),
		DiscoveryBasicPollSettings:       ExpandIpv6networkcontainerDiscoveryBasicPollSettings(ctx, m.DiscoveryBasicPollSettings, diags),
		DiscoveryBlackoutSetting:         ExpandIpv6networkcontainerDiscoveryBlackoutSetting(ctx, m.DiscoveryBlackoutSetting, diags),
		DiscoveryMember:                  flex.ExpandStringPointer(m.DiscoveryMember),
		DomainNameServers:                flex.ExpandFrameworkListString(ctx, m.DomainNameServers, diags),
		EnableDdns:                       flex.ExpandBoolPointer(m.EnableDdns),
		EnableDiscovery:                  flex.ExpandBoolPointer(m.EnableDiscovery),
		EnableImmediateDiscovery:         flex.ExpandBoolPointer(m.EnableImmediateDiscovery),
		ExtAttrs:                         flex.ExpandMapStringAny(ctx, m.ExtAttrs, diags),
		FederatedRealms:                  flex.ExpandFrameworkListNestedBlock(ctx, m.FederatedRealms, diags, ExpandIpv6networkcontainerFederatedRealms),
		LogicFilterRules:                 flex.ExpandFrameworkListNestedBlock(ctx, m.LogicFilterRules, diags, ExpandIpv6networkcontainerLogicFilterRules),
		MgmPrivate:                       flex.ExpandBoolPointer(m.MgmPrivate),
		Options:                          flex.ExpandFrameworkListNestedBlock(ctx, m.Options, diags, ExpandIpv6networkcontainerOptions),
		PortControlBlackoutSetting:       ExpandIpv6networkcontainerPortControlBlackoutSetting(ctx, m.PortControlBlackoutSetting, diags),
		PreferredLifetime:                flex.ExpandInt64Pointer(m.PreferredLifetime),
		RestartIfNeeded:                  flex.ExpandBoolPointer(m.RestartIfNeeded),
		RirOrganization:                  flex.ExpandStringPointer(m.RirOrganization),
		RirRegistrationAction:            flex.ExpandStringPointer(m.RirRegistrationAction),
		RirRegistrationStatus:            flex.ExpandStringPointerNullAsEmpty(m.RirRegistrationStatus),
		SamePortControlDiscoveryBlackout: flex.ExpandBoolPointer(m.SamePortControlDiscoveryBlackout),
		SendRirRequest:                   flex.ExpandBoolPointer(m.SendRirRequest),
		SubscribeSettings:                ExpandIpv6networkcontainerSubscribeSettings(ctx, m.SubscribeSettings, diags),
		Unmanaged:                        flex.ExpandBoolPointer(m.Unmanaged),
		UpdateDnsOnLeaseRenewal:          flex.ExpandBoolPointer(m.UpdateDnsOnLeaseRenewal),
		ValidLifetime:                    flex.ExpandInt64Pointer(m.ValidLifetime),
		ZoneAssociations:                 flex.ExpandFrameworkListNestedBlock(ctx, m.ZoneAssociations, diags, ExpandIpv6networkcontainerZoneAssociations),
	}
	if isCreate {
		ext.AutoCreateReversezone = flex.ExpandBoolPointer(m.AutoCreateReversezone)
		ext.Network = flex.ExpandIPv6Prefix(m.Network)
		ext.NetworkView = flex.ExpandStringPointerNullAsEmpty(m.NetworkView)
		ext.FuncCall = BuildIpv6networkcontainerFuncCall(ctx, m.DynamicAllocation, diags)
	}
	return ext
}

// ApplyIpv6networkcontainerNIOSUseFlags derives NIOS use flags from the raw config
// value(s) and writes them onto the core model. A flag is true when the user
// set any of its governed value fields in config.
func ApplyIpv6networkcontainerNIOSUseFlags(ctx context.Context, config tfsdk.Config, obj *coremodel.Ipv6networkcontainer, diags *diag.Diagnostics) {
	if obj == nil || obj.NIOS == nil {
		return
	}
	obj.NIOS.UseBlackoutSetting = flex.DeriveUseFlag(ctx, config, diags, path.Root("nios").AtName("discovery_blackout_setting"), path.Root("nios").AtName("port_control_blackout_setting"), path.Root("nios").AtName("same_port_control_discovery_blackout"))
	obj.NIOS.UseDdnsDomainname = flex.DeriveUseFlag(ctx, config, diags, path.Root("nios").AtName("ddns_domainname"))
	obj.NIOS.UseDdnsEnableOptionFqdn = flex.DeriveUseFlag(ctx, config, diags, path.Root("nios").AtName("ddns_enable_option_fqdn"))
	obj.NIOS.UseDdnsGenerateHostname = flex.DeriveUseFlag(ctx, config, diags, path.Root("nios").AtName("ddns_generate_hostname"))
	obj.NIOS.UseDdnsTtl = flex.DeriveUseFlag(ctx, config, diags, path.Root("nios").AtName("ddns_ttl"))
	obj.NIOS.UseDiscoveryBasicPollingSettings = flex.DeriveUseFlag(ctx, config, diags, path.Root("nios").AtName("discovery_basic_poll_settings"))
	obj.NIOS.UseDomainNameServers = flex.DeriveUseFlag(ctx, config, diags, path.Root("nios").AtName("domain_name_servers"))
	obj.NIOS.UseEnableDdns = flex.DeriveUseFlag(ctx, config, diags, path.Root("nios").AtName("enable_ddns"))
	obj.NIOS.UseEnableDiscovery = flex.DeriveUseFlag(ctx, config, diags, path.Root("nios").AtName("discovery_member"), path.Root("nios").AtName("enable_discovery"))
	obj.NIOS.UseLogicFilterRules = flex.DeriveUseFlag(ctx, config, diags, path.Root("nios").AtName("logic_filter_rules"))
	obj.NIOS.UseMgmPrivate = flex.DeriveUseFlag(ctx, config, diags, path.Root("nios").AtName("mgm_private"))
	obj.NIOS.UseOptions = flex.DeriveUseFlag(ctx, config, diags, path.Root("nios").AtName("options"))
	obj.NIOS.UsePreferredLifetime = flex.DeriveUseFlag(ctx, config, diags, path.Root("nios").AtName("preferred_lifetime"))
	obj.NIOS.UseSubscribeSettings = flex.DeriveUseFlag(ctx, config, diags, path.Root("nios").AtName("subscribe_settings"))
	obj.NIOS.UseUpdateDnsOnLeaseRenewal = flex.DeriveUseFlag(ctx, config, diags, path.Root("nios").AtName("update_dns_on_lease_renewal"))
	obj.NIOS.UseValidLifetime = flex.DeriveUseFlag(ctx, config, diags, path.Root("nios").AtName("valid_lifetime"))
	obj.NIOS.UseZoneAssociations = flex.DeriveUseFlag(ctx, config, diags, path.Root("nios").AtName("zone_associations"))
}

// Expand converts the UDDI TF model to the core model.
func (m *UDDIIpv6networkcontainerModel) Expand(ctx context.Context, diags *diag.Diagnostics, isCreate bool) *coremodel.UDDIIpv6networkcontainerExt {
	ext := &coremodel.UDDIIpv6networkcontainerExt{
		AsmConfig:                  ExpandASMConfig(ctx, m.AsmConfig, diags),
		Cidr:                       flex.ExpandInt64Pointer(m.Cidr),
		Comment:                    flex.ExpandStringPointer(m.Comment),
		CompartmentId:              flex.ExpandStringPointer(m.CompartmentId),
		DdnsClientUpdate:           flex.ExpandStringPointer(m.DdnsClientUpdate),
		DdnsConflictResolutionMode: flex.ExpandStringPointer(m.DdnsConflictResolutionMode),
		DdnsDomain:                 flex.ExpandStringPointer(m.DdnsDomain),
		DdnsGenerateName:           flex.ExpandBoolPointer(m.DdnsGenerateName),
		DdnsGeneratedPrefix:        flex.ExpandStringPointer(m.DdnsGeneratedPrefix),
		DdnsSendUpdates:            flex.ExpandBoolPointer(m.DdnsSendUpdates),
		DdnsTtlPercent:             flex.ExpandFloat32Pointer(m.DdnsTtlPercent),
		DdnsUpdateOnRenew:          flex.ExpandBoolPointer(m.DdnsUpdateOnRenew),
		DdnsUseConflictResolution:  flex.ExpandBoolPointer(m.DdnsUseConflictResolution),
		DhcpConfig:                 ExpandIpv6NetworkcontainerDHCPConfig(ctx, m.DhcpConfig, diags),
		DhcpOptions:                flex.ExpandFrameworkListNestedBlock(ctx, m.DhcpOptions, diags, ExpandOptionItem),
		ExternalKeys:               flex.ExpandMapStringAny(ctx, m.ExternalKeys, diags),
		FederatedRealms:            flex.ExpandFrameworkListString(ctx, m.FederatedRealms, diags),
		HostnameRewriteChar:        flex.ExpandStringPointer(m.HostnameRewriteChar),
		HostnameRewriteEnabled:     flex.ExpandBoolPointer(m.HostnameRewriteEnabled),
		HostnameRewriteRegex:       flex.ExpandStringPointer(m.HostnameRewriteRegex),
		InheritanceParent:          flex.ExpandStringPointer(m.InheritanceParent),
		InheritanceSources:         ExpandDHCPInheritance(ctx, m.InheritanceSources, diags),
		Name:                       flex.ExpandStringPointer(m.Name),
		Parent:                     flex.ExpandStringPointer(m.Parent),
		Tags:                       flex.ExpandMapStringAny(ctx, m.Tags, diags),
		Threshold:                  ExpandUtilizationThreshold(ctx, m.Threshold, diags),
	}
	if isCreate {
		ext.Address = flex.ExpandIPv6Address(m.Address)
		ext.Space = flex.ExpandStringPointer(m.Space)
	}
	return ext
}

// Flatten populates the TF model from a core response.
func (m *Ipv6networkcontainerModel) Flatten(ctx context.Context, resp *coremodel.Ipv6networkcontainer, diags *diag.Diagnostics) {
	if resp == nil {
		return
	}

	m.Id = flex.FlattenStringPointer(resp.Id)

	// Extract existing NIOS model, flatten API response onto it, convert back
	niosModel := flex.ExpandNestedObject[NIOSIpv6networkcontainerModel](ctx, m.NIOS, diags)
	if niosModel == nil {
		niosModel = &NIOSIpv6networkcontainerModel{}
	}
	plannedNIOS := flex.ExpandNestedObject[NIOSIpv6networkcontainerModel](ctx, m.NIOS, diags)
	niosModel.Flatten(ctx, resp.NIOS, diags)
	if resp.NIOS != nil {
		PostFlattenIpv6networkcontainerNIOS(ctx, plannedNIOS, niosModel, diags)
		m.NIOS = flex.FlattenNestedObject(ctx, niosModel, NIOSIpv6networkcontainerAttrTypes, diags)
	} else {
		m.NIOS = types.ObjectNull(NIOSIpv6networkcontainerAttrTypes)
	}

	// Extract existing UDDI model, flatten API response onto it, convert back
	uddiModel := flex.ExpandNestedObject[UDDIIpv6networkcontainerModel](ctx, m.UDDI, diags)
	if uddiModel == nil {
		uddiModel = &UDDIIpv6networkcontainerModel{}
	}
	uddiModel.Flatten(ctx, resp.UDDI, diags)
	if resp.UDDI != nil {
		m.UDDI = flex.FlattenNestedObject(ctx, uddiModel, UDDIIpv6networkcontainerAttrTypes, diags)
	} else {
		m.UDDI = types.ObjectNull(UDDIIpv6networkcontainerAttrTypes)
	}
}

// Flatten merges API response onto existing NIOS model.
func (m *NIOSIpv6networkcontainerModel) Flatten(ctx context.Context, from *coremodel.NIOSIpv6networkcontainerExt, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	planExtAttrs := m.ExtAttrs
	if planExtAttrs.IsUnknown() {
		planExtAttrs = types.MapNull(types.StringType)
	}
	m.CloudInfo = FlattenIpv6networkcontainerCloudInfo(ctx, from.CloudInfo, diags)
	m.Comment = flex.FlattenStringPointerEmptyAsNull(from.Comment)
	m.DdnsDomainname = flex.FlattenStringPointerEmptyAsNull(from.DdnsDomainname)
	m.DdnsEnableOptionFqdn = flex.FlattenBoolPointer(from.DdnsEnableOptionFqdn)
	m.DdnsGenerateHostname = flex.FlattenBoolPointer(from.DdnsGenerateHostname)
	m.DdnsServerAlwaysUpdates = flex.FlattenBoolPointer(from.DdnsServerAlwaysUpdates)
	m.DdnsTtl = flex.FlattenInt64Pointer(from.DdnsTtl)
	m.DiscoveryBasicPollSettings = FlattenIpv6networkcontainerDiscoveryBasicPollSettings(ctx, from.DiscoveryBasicPollSettings, diags)
	m.DiscoveryBlackoutSetting = FlattenIpv6networkcontainerDiscoveryBlackoutSetting(ctx, from.DiscoveryBlackoutSetting, diags)
	m.DiscoveryMember = flex.FlattenStringPointerEmptyAsNull(from.DiscoveryMember)
	m.DomainNameServers = flex.FlattenFrameworkListString(ctx, from.DomainNameServers, diags)
	m.EnableDdns = flex.FlattenBoolPointer(from.EnableDdns)
	m.EnableDiscovery = flex.FlattenBoolPointer(from.EnableDiscovery)
	m.ExtAttrs, m.ExtAttrsAll = flex.FlattenEAs(planExtAttrs, from.ExtAttrs)
	m.FederatedRealms = flex.FlattenFrameworkListNestedBlock(ctx, from.FederatedRealms, Ipv6networkcontainerFederatedRealmsAttrTypes, diags, FlattenIpv6networkcontainerFederatedRealms)
	m.LogicFilterRules = flex.FlattenFrameworkListNestedBlock(ctx, from.LogicFilterRules, Ipv6networkcontainerLogicFilterRulesAttrTypes, diags, FlattenIpv6networkcontainerLogicFilterRules)
	m.MgmPrivate = flex.FlattenBoolPointer(from.MgmPrivate)
	m.Network = flex.FlattenIPv6Prefix(from.Network)
	m.NetworkView = flex.FlattenStringPointerEmptyAsNull(from.NetworkView)
	m.Options = flex.FlattenFrameworkListNestedBlock(ctx, from.Options, Ipv6networkcontainerOptionsAttrTypes, diags, FlattenIpv6networkcontainerOptions)
	m.PortControlBlackoutSetting = FlattenIpv6networkcontainerPortControlBlackoutSetting(ctx, from.PortControlBlackoutSetting, diags)
	m.PreferredLifetime = flex.FlattenInt64Pointer(from.PreferredLifetime)
	m.RirOrganization = flex.FlattenStringPointerEmptyAsNull(from.RirOrganization)
	m.RirRegistrationStatus = flex.FlattenStringPointerEmptyAsNull(from.RirRegistrationStatus)
	m.SamePortControlDiscoveryBlackout = flex.FlattenBoolPointer(from.SamePortControlDiscoveryBlackout)
	m.SubscribeSettings = FlattenIpv6networkcontainerSubscribeSettings(ctx, from.SubscribeSettings, diags)
	m.Unmanaged = flex.FlattenBoolPointer(from.Unmanaged)
	m.UpdateDnsOnLeaseRenewal = flex.FlattenBoolPointer(from.UpdateDnsOnLeaseRenewal)
	m.ValidLifetime = flex.FlattenInt64Pointer(from.ValidLifetime)
	m.ZoneAssociations = flex.FlattenFrameworkListNestedBlock(ctx, from.ZoneAssociations, Ipv6networkcontainerZoneAssociationsAttrTypes, diags, FlattenIpv6networkcontainerZoneAssociations)
	if len(m.DynamicAllocation.AttributeTypes(ctx)) == 0 {
		m.DynamicAllocation = types.ObjectNull(dynamicallocation.NextAvailableNetworkAttrTypes)
	}
}

// Flatten merges API response onto existing UDDI model.
func (m *UDDIIpv6networkcontainerModel) Flatten(ctx context.Context, from *coremodel.UDDIIpv6networkcontainerExt, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.Address = flex.FlattenIPv6Address(from.Address)
	m.AsmConfig = FlattenASMConfig(ctx, from.AsmConfig, diags)
	m.Cidr = flex.FlattenInt64Pointer(from.Cidr)
	m.Comment = flex.FlattenStringPointer(from.Comment)
	m.CompartmentId = flex.FlattenStringPointer(from.CompartmentId)
	m.DdnsClientUpdate = flex.FlattenStringPointer(from.DdnsClientUpdate)
	m.DdnsConflictResolutionMode = flex.FlattenStringPointer(from.DdnsConflictResolutionMode)
	m.DdnsDomain = flex.FlattenStringPointer(from.DdnsDomain)
	m.DdnsGenerateName = flex.FlattenBoolPointer(from.DdnsGenerateName)
	m.DdnsGeneratedPrefix = flex.FlattenStringPointer(from.DdnsGeneratedPrefix)
	m.DdnsSendUpdates = flex.FlattenBoolPointer(from.DdnsSendUpdates)
	m.DdnsTtlPercent = flex.FlattenFloat32Pointer(from.DdnsTtlPercent)
	m.DdnsUpdateOnRenew = flex.FlattenBoolPointer(from.DdnsUpdateOnRenew)
	m.DdnsUseConflictResolution = flex.FlattenBoolPointer(from.DdnsUseConflictResolution)
	m.DhcpConfig = FlattenIpv6NetworkcontainerDHCPConfig(ctx, from.DhcpConfig, diags)
	m.DhcpOptions = flex.FlattenFrameworkListNestedBlock(ctx, from.DhcpOptions, OptionItemAttrTypes, diags, FlattenOptionItem)
	m.ExternalKeys = flex.FlattenMapStringAny(ctx, from.ExternalKeys, diags)
	m.FederatedRealms = flex.FlattenFrameworkUnorderedListString(ctx, from.FederatedRealms, diags)
	m.HostnameRewriteChar = flex.FlattenStringPointer(from.HostnameRewriteChar)
	m.HostnameRewriteEnabled = flex.FlattenBoolPointer(from.HostnameRewriteEnabled)
	m.HostnameRewriteRegex = flex.FlattenStringPointer(from.HostnameRewriteRegex)
	m.InheritanceParent = flex.FlattenStringPointer(from.InheritanceParent)
	m.InheritanceSources = FlattenDHCPInheritance(ctx, from.InheritanceSources, diags)
	m.Name = flex.FlattenStringPointer(from.Name)
	m.Parent = flex.FlattenStringPointer(from.Parent)
	m.Space = flex.FlattenStringPointer(from.Space)
	tagsAll := flex.FlattenMapStringAny(ctx, from.Tags, diags)
	if m.Tags.IsNull() || m.Tags.IsUnknown() {
		m.Tags = tagsAll
	}
	m.TagsAll = tagsAll
	m.Threshold = FlattenUtilizationThreshold(ctx, from.Threshold, diags)
}
