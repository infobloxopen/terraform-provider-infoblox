package dhcp

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-nettypes/cidrtypes"
	"github.com/hashicorp/terraform-plugin-framework-nettypes/iptypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/mapvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/mapdefault"
	objectplanmodifier "github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	stringplanmodifier "github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"

	coremodel "github.com/infobloxopen/terraform-provider-infoblox/internal/core/model/dhcp"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
	immutable "github.com/infobloxopen/terraform-provider-infoblox/internal/planmodifiers/immutable"
	importmod "github.com/infobloxopen/terraform-provider-infoblox/internal/planmodifiers/import"
	internaltypes "github.com/infobloxopen/terraform-provider-infoblox/internal/types"
	customvalidator "github.com/infobloxopen/terraform-provider-infoblox/internal/validator"
)

type RangeModel struct {
	Id   types.String `tfsdk:"id"`
	NIOS types.Object `tfsdk:"nios"`
	UDDI types.Object `tfsdk:"uddi"`
}

var RangeAttrTypes = map[string]attr.Type{
	"id":   types.StringType,
	"nios": types.ObjectType{AttrTypes: NIOSRangeAttrTypes},
	"uddi": types.ObjectType{AttrTypes: UDDIRangeAttrTypes},
}

type NIOSRangeModel struct {
	AlwaysUpdateDns                  types.Bool                       `tfsdk:"always_update_dns"`
	Bootfile                         types.String                     `tfsdk:"bootfile"`
	Bootserver                       types.String                     `tfsdk:"bootserver"`
	CloudInfo                        types.Object                     `tfsdk:"cloud_info"`
	Comment                          types.String                     `tfsdk:"comment"`
	DdnsDomainname                   types.String                     `tfsdk:"ddns_domainname"`
	DdnsGenerateHostname             types.Bool                       `tfsdk:"ddns_generate_hostname"`
	DenyAllClients                   types.Bool                       `tfsdk:"deny_all_clients"`
	DenyBootp                        types.Bool                       `tfsdk:"deny_bootp"`
	Disable                          types.Bool                       `tfsdk:"disable"`
	DiscoveryBasicPollSettings       types.Object                     `tfsdk:"discovery_basic_poll_settings"`
	DiscoveryBlackoutSetting         types.Object                     `tfsdk:"discovery_blackout_setting"`
	DiscoveryMember                  types.String                     `tfsdk:"discovery_member"`
	EmailList                        internaltypes.UnorderedListValue `tfsdk:"email_list"`
	EnableDdns                       types.Bool                       `tfsdk:"enable_ddns"`
	EnableDhcpThresholds             types.Bool                       `tfsdk:"enable_dhcp_thresholds"`
	EnableDiscovery                  types.Bool                       `tfsdk:"enable_discovery"`
	EnableEmailWarnings              types.Bool                       `tfsdk:"enable_email_warnings"`
	EnableIfmapPublishing            types.Bool                       `tfsdk:"enable_ifmap_publishing"`
	EnableImmediateDiscovery         types.Bool                       `tfsdk:"enable_immediate_discovery"`
	EnablePxeLeaseTime               types.Bool                       `tfsdk:"enable_pxe_lease_time"`
	EnableSnmpWarnings               types.Bool                       `tfsdk:"enable_snmp_warnings"`
	EndAddr                          iptypes.IPv4Address              `tfsdk:"end_addr"`
	Exclude                          types.List                       `tfsdk:"exclude"`
	ExtAttrs                         types.Map                        `tfsdk:"ext_attrs"`
	ExtAttrsAll                      types.Map                        `tfsdk:"ext_attrs_all"`
	FailoverAssociation              types.String                     `tfsdk:"failover_association"`
	FingerprintFilterRules           types.List                       `tfsdk:"fingerprint_filter_rules"`
	HighWaterMark                    types.Int64                      `tfsdk:"high_water_mark"`
	HighWaterMarkReset               types.Int64                      `tfsdk:"high_water_mark_reset"`
	IgnoreDhcpOptionListRequest      types.Bool                       `tfsdk:"ignore_dhcp_option_list_request"`
	IgnoreId                         types.String                     `tfsdk:"ignore_id"`
	IgnoreMacAddresses               internaltypes.UnorderedListValue `tfsdk:"ignore_mac_addresses"`
	KnownClients                     types.String                     `tfsdk:"known_clients"`
	LeaseScavengeTime                types.Int64                      `tfsdk:"lease_scavenge_time"`
	LogicFilterRules                 types.List                       `tfsdk:"logic_filter_rules"`
	LowWaterMark                     types.Int64                      `tfsdk:"low_water_mark"`
	LowWaterMarkReset                types.Int64                      `tfsdk:"low_water_mark_reset"`
	MacFilterRules                   types.List                       `tfsdk:"mac_filter_rules"`
	Member                           types.Object                     `tfsdk:"member"`
	MsOptions                        types.List                       `tfsdk:"ms_options"`
	MsServer                         types.Object                     `tfsdk:"ms_server"`
	NacFilterRules                   types.List                       `tfsdk:"nac_filter_rules"`
	Name                             types.String                     `tfsdk:"name"`
	Network                          cidrtypes.IPv4Prefix             `tfsdk:"network"`
	NetworkView                      types.String                     `tfsdk:"network_view"`
	Nextserver                       types.String                     `tfsdk:"nextserver"`
	OptionFilterRules                types.List                       `tfsdk:"option_filter_rules"`
	Options                          types.List                       `tfsdk:"options"`
	PortControlBlackoutSetting       types.Object                     `tfsdk:"port_control_blackout_setting"`
	PxeLeaseTime                     types.Int64                      `tfsdk:"pxe_lease_time"`
	RecycleLeases                    types.Bool                       `tfsdk:"recycle_leases"`
	RelayAgentFilterRules            types.List                       `tfsdk:"relay_agent_filter_rules"`
	RestartIfNeeded                  types.Bool                       `tfsdk:"restart_if_needed"`
	SamePortControlDiscoveryBlackout types.Bool                       `tfsdk:"same_port_control_discovery_blackout"`
	ServerAssociationType            types.String                     `tfsdk:"server_association_type"`
	SplitMember                      types.Object                     `tfsdk:"split_member"`
	SplitScopeExclusionPercent       types.Int64                      `tfsdk:"split_scope_exclusion_percent"`
	StartAddr                        iptypes.IPv4Address              `tfsdk:"start_addr"`
	SubscribeSettings                types.Object                     `tfsdk:"subscribe_settings"`
	Template                         types.String                     `tfsdk:"template"`
	UnknownClients                   types.String                     `tfsdk:"unknown_clients"`
	UpdateDnsOnLeaseRenewal          types.Bool                       `tfsdk:"update_dns_on_lease_renewal"`
}

var NIOSRangeAttrTypes = map[string]attr.Type{
	"always_update_dns":                    types.BoolType,
	"bootfile":                             types.StringType,
	"bootserver":                           types.StringType,
	"cloud_info":                           types.ObjectType{AttrTypes: RangeCloudInfoAttrTypes},
	"comment":                              types.StringType,
	"ddns_domainname":                      types.StringType,
	"ddns_generate_hostname":               types.BoolType,
	"deny_all_clients":                     types.BoolType,
	"deny_bootp":                           types.BoolType,
	"disable":                              types.BoolType,
	"discovery_basic_poll_settings":        types.ObjectType{AttrTypes: RangeDiscoveryBasicPollSettingsAttrTypes},
	"discovery_blackout_setting":           types.ObjectType{AttrTypes: RangeDiscoveryBlackoutSettingAttrTypes},
	"discovery_member":                     types.StringType,
	"email_list":                           internaltypes.UnorderedListOfStringType,
	"enable_ddns":                          types.BoolType,
	"enable_dhcp_thresholds":               types.BoolType,
	"enable_discovery":                     types.BoolType,
	"enable_email_warnings":                types.BoolType,
	"enable_ifmap_publishing":              types.BoolType,
	"enable_immediate_discovery":           types.BoolType,
	"enable_pxe_lease_time":                types.BoolType,
	"enable_snmp_warnings":                 types.BoolType,
	"end_addr":                             iptypes.IPv4AddressType{},
	"exclude":                              types.ListType{ElemType: types.ObjectType{AttrTypes: RangeExcludeAttrTypes}},
	"ext_attrs":                            types.MapType{ElemType: types.StringType},
	"ext_attrs_all":                        types.MapType{ElemType: types.StringType},
	"failover_association":                 types.StringType,
	"fingerprint_filter_rules":             types.ListType{ElemType: types.ObjectType{AttrTypes: RangeFingerprintFilterRulesAttrTypes}},
	"high_water_mark":                      types.Int64Type,
	"high_water_mark_reset":                types.Int64Type,
	"ignore_dhcp_option_list_request":      types.BoolType,
	"ignore_id":                            types.StringType,
	"ignore_mac_addresses":                 internaltypes.UnorderedListOfStringType,
	"known_clients":                        types.StringType,
	"lease_scavenge_time":                  types.Int64Type,
	"logic_filter_rules":                   types.ListType{ElemType: types.ObjectType{AttrTypes: RangeLogicFilterRulesAttrTypes}},
	"low_water_mark":                       types.Int64Type,
	"low_water_mark_reset":                 types.Int64Type,
	"mac_filter_rules":                     types.ListType{ElemType: types.ObjectType{AttrTypes: RangeMacFilterRulesAttrTypes}},
	"member":                               types.ObjectType{AttrTypes: RangeMemberAttrTypes},
	"ms_options":                           types.ListType{ElemType: types.ObjectType{AttrTypes: RangeMsOptionsAttrTypes}},
	"ms_server":                            types.ObjectType{AttrTypes: RangeMsServerAttrTypes},
	"nac_filter_rules":                     types.ListType{ElemType: types.ObjectType{AttrTypes: RangeNacFilterRulesAttrTypes}},
	"name":                                 types.StringType,
	"network":                              cidrtypes.IPv4PrefixType{},
	"network_view":                         types.StringType,
	"nextserver":                           types.StringType,
	"option_filter_rules":                  types.ListType{ElemType: types.ObjectType{AttrTypes: RangeOptionFilterRulesAttrTypes}},
	"options":                              types.ListType{ElemType: types.ObjectType{AttrTypes: RangeOptionsAttrTypes}},
	"port_control_blackout_setting":        types.ObjectType{AttrTypes: RangePortControlBlackoutSettingAttrTypes},
	"pxe_lease_time":                       types.Int64Type,
	"recycle_leases":                       types.BoolType,
	"relay_agent_filter_rules":             types.ListType{ElemType: types.ObjectType{AttrTypes: RangeRelayAgentFilterRulesAttrTypes}},
	"restart_if_needed":                    types.BoolType,
	"same_port_control_discovery_blackout": types.BoolType,
	"server_association_type":              types.StringType,
	"split_member":                         types.ObjectType{AttrTypes: RangeSplitMemberAttrTypes},
	"split_scope_exclusion_percent":        types.Int64Type,
	"start_addr":                           iptypes.IPv4AddressType{},
	"subscribe_settings":                   types.ObjectType{AttrTypes: RangeSubscribeSettingsAttrTypes},
	"template":                             types.StringType,
	"unknown_clients":                      types.StringType,
	"update_dns_on_lease_renewal":          types.BoolType,
}

type UDDIRangeModel struct {
	Comment            types.String `tfsdk:"comment"`
	DhcpHost           types.String `tfsdk:"dhcp_host"`
	DhcpOptions        types.List   `tfsdk:"dhcp_options"`
	DisableDhcp        types.Bool   `tfsdk:"disable_dhcp"`
	End                types.String `tfsdk:"end"`
	ExclusionRanges    types.List   `tfsdk:"exclusion_ranges"`
	Filters            types.List   `tfsdk:"filters"`
	InheritanceParent  types.String `tfsdk:"inheritance_parent"`
	InheritanceSources types.Object `tfsdk:"inheritance_sources"`
	Name               types.String `tfsdk:"name"`
	Parent             types.String `tfsdk:"parent"`
	Space              types.String `tfsdk:"space"`
	Start              types.String `tfsdk:"start"`
	Tags               types.Map    `tfsdk:"tags"`
	TagsAll            types.Map    `tfsdk:"tags_all"`
	Threshold          types.Object `tfsdk:"threshold"`
}

var UDDIRangeAttrTypes = map[string]attr.Type{
	"comment":             types.StringType,
	"dhcp_host":           types.StringType,
	"dhcp_options":        types.ListType{ElemType: types.ObjectType{AttrTypes: OptionItemAttrTypes}},
	"disable_dhcp":        types.BoolType,
	"end":                 types.StringType,
	"exclusion_ranges":    types.ListType{ElemType: types.ObjectType{AttrTypes: ExclusionRangeAttrTypes}},
	"filters":             types.ListType{ElemType: types.ObjectType{AttrTypes: AccessFilterAttrTypes}},
	"inheritance_parent":  types.StringType,
	"inheritance_sources": types.ObjectType{AttrTypes: DHCPOptionsInheritanceAttrTypes},
	"name":                types.StringType,
	"parent":              types.StringType,
	"space":               types.StringType,
	"start":               types.StringType,
	"tags":                types.MapType{ElemType: types.StringType},
	"tags_all":            types.MapType{ElemType: types.StringType},
	"threshold":           types.ObjectType{AttrTypes: UtilizationThresholdAttrTypes},
}

const (
	RangeInheritanceType = "full"
	RangeReturnFields    = "always_update_dns,bootfile,bootserver,cloud_info,comment,ddns_domainname,ddns_generate_hostname,deny_all_clients,deny_bootp,dhcp_utilization,dhcp_utilization_status,disable,discover_now_status,discovery_basic_poll_settings,discovery_blackout_setting,discovery_member,dynamic_hosts,email_list,enable_ddns,enable_dhcp_thresholds,enable_discovery,enable_email_warnings,enable_ifmap_publishing,enable_pxe_lease_time,enable_snmp_warnings,end_addr,endpoint_sources,exclude,extattrs,failover_association,fingerprint_filter_rules,high_water_mark,high_water_mark_reset,ignore_dhcp_option_list_request,ignore_id,ignore_mac_addresses,is_split_scope,known_clients,lease_scavenge_time,logic_filter_rules,low_water_mark,low_water_mark_reset,mac_filter_rules,member,ms_ad_user_data,ms_options,ms_server,nac_filter_rules,name,network,network_view,nextserver,option_filter_rules,options,port_control_blackout_setting,pxe_lease_time,recycle_leases,relay_agent_filter_rules,same_port_control_discovery_blackout,server_association_type,start_addr,static_hosts,subscribe_settings,total_hosts,unknown_clients,update_dns_on_lease_renewal,use_blackout_setting,use_bootfile,use_bootserver,use_ddns_domainname,use_ddns_generate_hostname,use_deny_bootp,use_discovery_basic_polling_settings,use_email_list,use_enable_ddns,use_enable_dhcp_thresholds,use_enable_discovery,use_enable_ifmap_publishing,use_ignore_dhcp_option_list_request,use_ignore_id,use_known_clients,use_lease_scavenge_time,use_logic_filter_rules,use_ms_options,use_nextserver,use_options,use_pxe_lease_time,use_recycle_leases,use_subscribe_settings,use_unknown_clients,use_update_dns_on_lease_renewal"
)

var RangeResourceSchemaAttributes = map[string]schema.Attribute{
	"id": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The reference to the object.",
	},
	"nios": schema.SingleNestedAttribute{
		Optional:            true,
		MarkdownDescription: "NIOS backend-specific fields.",
		Attributes:          RangeResourceNiosSchemaAttributes,
	},
	"uddi": schema.SingleNestedAttribute{
		Optional:            true,
		MarkdownDescription: "UDDI backend-specific fields.",
		Attributes:          RangeResourceUddiSchemaAttributes,
	},
}

var RangeResourceNiosSchemaAttributes = map[string]schema.Attribute{
	"always_update_dns": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "This field controls whether only the DHCP server is allowed to update DNS, regardless of the DHCP clients requests.",
	},
	"bootfile": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "The bootfile name for the range. You can configure the DHCP server to support clients that use the boot file name option in their DHCPREQUEST messages.",
	},
	"bootserver": schema.StringAttribute{
		Optional: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
			customvalidator.IsValidIPv4OrFQDN(),
		},
		MarkdownDescription: "The bootserver address for the range. You can specify the name and/or IP address of the boot server that the host needs to boot. The boot server IPv4 Address or name in FQDN format.",
	},
	"cloud_info": schema.SingleNestedAttribute{
		Attributes:          RangeCloudInfoResourceSchemaAttributes,
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "",
	},
	"comment": schema.StringAttribute{
		Optional: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
			customvalidator.ValidateTrimmedString(),
			stringvalidator.LengthBetween(0, 256),
		},
		MarkdownDescription: "Comment for the range; maximum 256 characters.",
	},
	"ddns_domainname": schema.StringAttribute{
		Optional: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "The dynamic DNS domain name the appliance uses specifically for DDNS updates for this range.",
	},
	"ddns_generate_hostname": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "If this field is set to True, the DHCP server generates a hostname and updates DNS with it when the DHCP client request does not contain a hostname.",
	},
	"deny_all_clients": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "If True, send NAK forcing the client to take the new address.",
	},
	"deny_bootp": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "If set to true, BOOTP settings are disabled and BOOTP requests will be denied.",
	},
	"disable": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Determines whether a range is disabled or not. When this is set to False, the range is enabled.",
	},
	"discovery_basic_poll_settings": schema.SingleNestedAttribute{
		Attributes:          RangeDiscoveryBasicPollSettingsResourceSchemaAttributes,
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "",
	},
	"discovery_blackout_setting": schema.SingleNestedAttribute{
		Attributes:          RangeDiscoveryBlackoutSettingResourceSchemaAttributes,
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "",
	},
	"discovery_member": schema.StringAttribute{
		Optional: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "The member that will run discovery for this range.",
	},
	"email_list": schema.ListAttribute{
		ElementType: types.StringType,
		Optional:    true,
		CustomType:  internaltypes.UnorderedListOfStringType,
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
		},
		MarkdownDescription: "The e-mail lists to which the appliance sends DHCP threshold alarm e-mail messages.",
	},
	"enable_ddns": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "The dynamic DNS updates flag of a DHCP range object. If set to True, the DHCP server sends DDNS updates to DNS servers in the same Grid, and to external DNS servers.",
	},
	"enable_dhcp_thresholds": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Determines if DHCP thresholds are enabled for the range.",
	},
	"enable_discovery": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Determines whether a discovery is enabled or not for this range. When this is set to False, the discovery for this range is disabled.",
	},
	"enable_email_warnings": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Determines if DHCP threshold warnings are sent through email.",
	},
	"enable_ifmap_publishing": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Determines if IFMAP publishing is enabled for the range.",
	},
	"enable_immediate_discovery": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines if the discovery for the range should be immediately enabled.",
	},
	"enable_pxe_lease_time": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Set this to True if you want the DHCP server to use a different lease time for PXE clients.",
	},
	"enable_snmp_warnings": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Determines if DHCP threshold warnings are send through SNMP.",
	},
	"end_addr": schema.StringAttribute{
		Required:   true,
		CustomType: iptypes.IPv4AddressType{},
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "The IPv4 Address end address of the range.",
	},
	"exclude": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: RangeExcludeResourceSchemaAttributes,
		},
		Optional: true,
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
		},
		MarkdownDescription: "These are ranges of IP addresses that the appliance does not use to assign to clients. You can use these exclusion addresses as static IP addresses. They contain the start and end addresses of the exclusion range, and optionally, information about this exclusion range.",
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
	"failover_association": schema.StringAttribute{
		Optional: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "The name of the failover association: the server in this failover association will serve the IPv4 range in case the main server is out of service. {range:range} must be set to 'FAILOVER' or 'FAILOVER_MS' if you want the failover association specified here to serve the range.",
	},
	"fingerprint_filter_rules": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: RangeFingerprintFilterRulesResourceSchemaAttributes,
		},
		Optional: true,
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
		},
		MarkdownDescription: "This field contains the fingerprint filters for this DHCP range. The appliance uses matching rules in these filters to select the address range from which it assigns a lease.",
	},
	"high_water_mark": schema.Int64Attribute{
		Optional: true,
		Computed: true,
		Default:  int64default.StaticInt64(95),
		Validators: []validator.Int64{
			int64validator.Between(1, 100),
		},
		MarkdownDescription: "The percentage of DHCP range usage threshold above which range usage is not expected and may warrant your attention. When the high watermark is reached, the Infoblox appliance generates a syslog message and sends a warning (if enabled). A number that specifies the percentage of allocated addresses. The range is from 1 to 100.",
	},
	"high_water_mark_reset": schema.Int64Attribute{
		Optional: true,
		Computed: true,
		Default:  int64default.StaticInt64(85),
		Validators: []validator.Int64{
			int64validator.Between(1, 100),
		},
		MarkdownDescription: "The percentage of DHCP range usage below which the corresponding SNMP trap is reset. A number that specifies the percentage of allocated addresses. The range is from 1 to 100. The high watermark reset value must be lower than the high watermark value.",
	},
	"ignore_dhcp_option_list_request": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "If this field is set to False, the appliance returns all DHCP options the client is eligible to receive, rather than only the list of options the client has requested.",
	},
	"ignore_id": schema.StringAttribute{
		Default: stringdefault.StaticString("NONE"),
		Validators: []validator.String{
			stringvalidator.OneOf("NONE", "CLIENT", "MACADDR"),
		},
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Indicates whether the appliance will ignore DHCP client IDs or MAC addresses. Valid values are \"NONE\", \"CLIENT\", or \"MACADDR\". The default is \"NONE\".",
	},
	"ignore_mac_addresses": schema.ListAttribute{
		ElementType: types.StringType,
		Optional:    true,
		CustomType:  internaltypes.UnorderedListOfStringType,
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
		},
		MarkdownDescription: "A list of MAC addresses the appliance will ignore.",
	},
	"known_clients": schema.StringAttribute{
		Validators: []validator.String{
			stringvalidator.OneOf("Allow", "Deny"),
		},
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Permission for known clients. This can be 'Allow' or 'Deny'. If set to 'Deny' known clients will be denied IP addresses. Known clients include roaming hosts and clients with fixed addresses or DHCP host entries. Unknown clients include clients that are not roaming hosts and clients that do not have fixed addresses or DHCP host entries.",
	},
	"lease_scavenge_time": schema.Int64Attribute{
		Optional: true,
		Computed: true,
		Default:  int64default.StaticInt64(-1),
		Validators: []validator.Int64{
			int64validator.Any(int64validator.OneOf(-1), int64validator.Between(86400, 2147472000)),
		},
		MarkdownDescription: "An integer that specifies the period of time (in seconds) that frees and backs up leases remained in the database before they are automatically deleted. To disable lease scavenging, set the parameter to -1. The minimum positive value must be greater than 86400 seconds (1 day).",
	},
	"logic_filter_rules": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: RangeLogicFilterRulesResourceSchemaAttributes,
		},
		Optional: true,
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
		},
		MarkdownDescription: "This field contains the logic filters to be applied to this range. This list corresponds to the match rules that are written to the dhcpd configuration file.",
	},
	"low_water_mark": schema.Int64Attribute{
		Optional: true,
		Computed: true,
		Default:  int64default.StaticInt64(0),
		Validators: []validator.Int64{
			int64validator.Any(int64validator.Between(0, 100)),
		},
		MarkdownDescription: "The percentage of DHCP range usage below which the Infoblox appliance generates a syslog message and sends a warning (if enabled). A number that specifies the percentage of allocated addresses. The range is from 1 to 100.",
	},
	"low_water_mark_reset": schema.Int64Attribute{
		Optional: true,
		Computed: true,
		Default:  int64default.StaticInt64(10),
		Validators: []validator.Int64{
			int64validator.Any(int64validator.Between(1, 100)),
		},
		MarkdownDescription: "The percentage of DHCP range usage threshold below which range usage is not expected and may warrant your attention. When the low watermark is crossed, the Infoblox appliance generates a syslog message and sends a warning (if enabled). A number that specifies the percentage of allocated addresses. The range is from 1 to 100. The low watermark reset value must be higher than the low watermark value.",
	},
	"mac_filter_rules": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: RangeMacFilterRulesResourceSchemaAttributes,
		},
		Optional: true,
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
		},
		MarkdownDescription: "This field contains the MAC filters to be applied to this range. The appliance uses the matching rules of these filters to select the address range from which it assigns a lease.",
	},
	"member": schema.SingleNestedAttribute{
		Attributes:          RangeMemberResourceSchemaAttributes,
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "",
	},
	"ms_options": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: RangeMsOptionsResourceSchemaAttributes,
		},
		Optional: true,
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
		},
		MarkdownDescription: "This field contains the Microsoft DHCP options for this range.",
	},
	"ms_server": schema.SingleNestedAttribute{
		Attributes:          RangeMsServerResourceSchemaAttributes,
		Optional:            true,
		MarkdownDescription: "",
	},
	"nac_filter_rules": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: RangeNacFilterRulesResourceSchemaAttributes,
		},
		Optional: true,
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
		},
		MarkdownDescription: "This field contains the NAC filters to be applied to this range. The appliance uses the matching rules of these filters to select the address range from which it assigns a lease.",
	},
	"name": schema.StringAttribute{
		Optional: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "This field contains the name of the Microsoft scope.",
	},
	"network": schema.StringAttribute{
		Optional:   true,
		Computed:   true,
		CustomType: cidrtypes.IPv4PrefixType{},
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "The network to which this range belongs, in IPv4 Address/CIDR format.",
	},
	"network_view": schema.StringAttribute{
		Default:  stringdefault.StaticString("default"),
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "The name of the network view in which this range resides.",
	},
	"nextserver": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
			customvalidator.IsValidIPv4OrFQDN(),
		},
		MarkdownDescription: "The name in FQDN and/or IPv4 Address of the next server that the host needs to boot.",
	},
	"option_filter_rules": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: RangeOptionFilterRulesResourceSchemaAttributes,
		},
		Optional: true,
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
		},
		MarkdownDescription: "This field contains the Option filters to be applied to this range. The appliance uses the matching rules of these filters to select the address range from which it assigns a lease.",
	},
	"options": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: RangeOptionsResourceSchemaAttributes,
		},
		Optional: true,
		Computed: true,
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
		},
		MarkdownDescription: "An array of DHCP option dhcpoption structs that lists the DHCP options associated with the object.",
	},
	"port_control_blackout_setting": schema.SingleNestedAttribute{
		Attributes:          RangePortControlBlackoutSettingResourceSchemaAttributes,
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "",
	},
	"pxe_lease_time": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "The PXE lease time value of a DHCP Range object. Some hosts use PXE (Preboot Execution Environment) to boot remotely from a server. To better manage your IP resources, set a different lease time for PXE boot requests. You can configure the DHCP server to allocate an IP address with a shorter lease time to hosts that send PXE boot requests, so IP addresses are not leased longer than necessary. A 32-bit unsigned integer that represents the duration, in seconds, for which the update is cached. Zero indicates that the update is not cached.",
	},
	"recycle_leases": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(true),
		MarkdownDescription: "If the field is set to True, the leases are kept in the Recycle Bin until one week after expiration. Otherwise, the leases are permanently deleted.",
	},
	"relay_agent_filter_rules": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: RangeRelayAgentFilterRulesResourceSchemaAttributes,
		},
		Optional: true,
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
		},
		MarkdownDescription: "This field contains the Relay Agent filters to be applied to this range. The appliance uses the matching rules of these filters to select the address range from which it assigns a lease.",
	},
	"restart_if_needed": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Restarts the member service.",
	},
	"same_port_control_discovery_blackout": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "If the field is set to True, the discovery blackout setting will be used for port control blackout setting.",
	},
	"server_association_type": schema.StringAttribute{
		Default: stringdefault.StaticString("NONE"),
		Validators: []validator.String{
			stringvalidator.OneOf("NONE", "MEMBER", "FAILOVER", "MS_SERVER", "MS_FAILOVER"),
		},
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The type of server that is going to serve the range.",
	},
	"split_member": schema.SingleNestedAttribute{
		Attributes: RangeSplitMemberResourceSchemaAttributes,
		Optional:   true,
		PlanModifiers: []planmodifier.Object{
			immutable.ImmutableObject(),
		},
		MarkdownDescription: "",
	},
	"split_scope_exclusion_percent": schema.Int64Attribute{
		Optional: true,
		PlanModifiers: []planmodifier.Int64{
			immutable.ImmutableInt64(),
		},
		MarkdownDescription: "This field controls the percentage used when creating a split scope. Valid values are numbers between 1 and 99. If the value is 40, it means that the top 40% of the exclusion will be created on the DHCP range assigned to {next_available_ip:next_available_ip} and the lower 60% of the range will be assigned to DHCP range assigned to {next_available_ip:next_available_ip}",
	},
	"start_addr": schema.StringAttribute{
		Required:   true,
		CustomType: iptypes.IPv4AddressType{},
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "The IPv4 Address starting address of the range.",
	},
	"subscribe_settings": schema.SingleNestedAttribute{
		Attributes:          RangeSubscribeSettingsResourceSchemaAttributes,
		Optional:            true,
		MarkdownDescription: "",
	},
	"template": schema.StringAttribute{
		Optional: true,
		PlanModifiers: []planmodifier.String{
			immutable.ImmutableString(),
		},
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "If set on creation, the range will be created according to the values specified in the named template.",
	},
	"unknown_clients": schema.StringAttribute{
		Validators: []validator.String{
			stringvalidator.OneOf("Allow", "Deny"),
		},
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Permission for unknown clients. This can be 'Allow' or 'Deny'. If set to 'Deny', unknown clients will be denied IP addresses. Known clients include roaming hosts and clients with fixed addresses or DHCP host entries. Unknown clients include clients that are not roaming hosts and clients that do not have fixed addresses or DHCP host entries.",
	},
	"update_dns_on_lease_renewal": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "This field controls whether the DHCP server updates DNS when a DHCP lease is renewed.",
	},
}

var RangeResourceUddiSchemaAttributes = map[string]schema.Attribute{
	"comment": schema.StringAttribute{
		Optional: true,
		Validators: []validator.String{
			stringvalidator.LengthBetween(0, 1024),
		},
		MarkdownDescription: "The description for the range. May contain 0 to 1024 characters. Can include UTF-8.",
	},
	"dhcp_host": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The resource identifier.",
	},
	"dhcp_options": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: OptionItemResourceSchemaAttributes,
		},
		Optional: true,
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
		},
		MarkdownDescription: "The list of DHCP options. May be either a specific option or a group of options.",
	},
	"disable_dhcp": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Optional. _true_ to disable object. A disabled object is effectively non-existent when generating configuration.  Defaults to _false_.",
	},
	"end": schema.StringAttribute{
		Required:            true,
		MarkdownDescription: "The end IP address of the range.",
	},
	"exclusion_ranges": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: ExclusionRangeResourceSchemaAttributes,
		},
		Optional: true,
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
		},
		MarkdownDescription: "The list of all exclusion ranges in the scope of the range.",
	},
	"filters": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: AccessFilterResourceSchemaAttributes,
		},
		Optional: true,
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
		},
		MarkdownDescription: "The list of all allow/deny filters of the range.",
	},
	"inheritance_parent": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The resource identifier.",
	},
	"inheritance_sources": schema.SingleNestedAttribute{
		Attributes: DHCPOptionsInheritanceResourceSchemaAttributes,
		Optional:   true,
		PlanModifiers: []planmodifier.Object{
			objectplanmodifier.UseStateForUnknown(),
		},
		MarkdownDescription: "The inheritance configuration that specifies how the _dhcp_options_ field is inherited from the parent object.",
	},
	"name": schema.StringAttribute{
		Optional: true,
		Validators: []validator.String{
			stringvalidator.LengthBetween(1, 256),
		},
		MarkdownDescription: "The name of the range. May contain 1 to 256 characters. Can include UTF-8.",
	},
	"parent": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The resource identifier.",
	},
	"space": schema.StringAttribute{
		Required: true,
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.RequiresReplaceIfConfigured(),
		},
		MarkdownDescription: "The resource identifier.",
	},
	"start": schema.StringAttribute{
		Required:            true,
		MarkdownDescription: "The start IP address of the range.",
	},
	"tags": schema.MapAttribute{
		Optional:    true,
		Computed:    true,
		ElementType: types.StringType,
		Default:     mapdefault.StaticValue(types.MapNull(types.StringType)),
		Validators: []validator.Map{
			mapvalidator.SizeAtLeast(1),
		},
		MarkdownDescription: "The tags for the range in JSON format.",
	},
	"tags_all": schema.MapAttribute{
		Computed:            true,
		ElementType:         types.StringType,
		MarkdownDescription: "All tags including inherited values.",
	},
	"threshold": schema.SingleNestedAttribute{
		Attributes:          UtilizationThresholdResourceSchemaAttributes,
		Optional:            true,
		MarkdownDescription: "A __UtilizationThreshold__ object represents IP address utilization threshold settings.",
	},
}

// Expand converts the TF model to the infoblox core model
func (m *RangeModel) Expand(ctx context.Context, diags *diag.Diagnostics, isCreate bool) *coremodel.Range {
	if m == nil {
		return nil
	}

	obj := &coremodel.Range{}

	// Expand NIOS nested attribute (returns nil if not present)
	niosModel := flex.ExpandNestedObject[NIOSRangeModel](ctx, m.NIOS, diags)
	if niosModel != nil {
		obj.NIOS = niosModel.Expand(ctx, diags, isCreate)
	}

	// Expand UDDI nested attribute (returns nil if not present)
	uddiModel := flex.ExpandNestedObject[UDDIRangeModel](ctx, m.UDDI, diags)
	if uddiModel != nil {
		obj.UDDI = uddiModel.Expand(ctx, diags)
	}

	return obj
}

// Expand converts the NIOS TF model to the core model.
func (m *NIOSRangeModel) Expand(ctx context.Context, diags *diag.Diagnostics, isCreate bool) *coremodel.NIOSRangeExt {
	ext := &coremodel.NIOSRangeExt{
		AlwaysUpdateDns:                  flex.ExpandBoolPointer(m.AlwaysUpdateDns),
		Bootfile:                         flex.ExpandStringPointer(m.Bootfile),
		Bootserver:                       flex.ExpandStringPointerNullAsEmpty(m.Bootserver),
		CloudInfo:                        ExpandRangeCloudInfo(ctx, m.CloudInfo, diags),
		Comment:                          flex.ExpandStringPointerNullAsEmpty(m.Comment),
		DdnsDomainname:                   flex.ExpandStringPointerNullAsEmpty(m.DdnsDomainname),
		DdnsGenerateHostname:             flex.ExpandBoolPointer(m.DdnsGenerateHostname),
		DenyAllClients:                   flex.ExpandBoolPointer(m.DenyAllClients),
		DenyBootp:                        flex.ExpandBoolPointer(m.DenyBootp),
		Disable:                          flex.ExpandBoolPointer(m.Disable),
		DiscoveryBasicPollSettings:       ExpandRangeDiscoveryBasicPollSettings(ctx, m.DiscoveryBasicPollSettings, diags),
		DiscoveryBlackoutSetting:         ExpandRangeDiscoveryBlackoutSetting(ctx, m.DiscoveryBlackoutSetting, diags),
		DiscoveryMember:                  flex.ExpandStringPointer(m.DiscoveryMember),
		EmailList:                        flex.ExpandFrameworkListString(ctx, m.EmailList, diags),
		EnableDdns:                       flex.ExpandBoolPointer(m.EnableDdns),
		EnableDhcpThresholds:             flex.ExpandBoolPointer(m.EnableDhcpThresholds),
		EnableDiscovery:                  flex.ExpandBoolPointer(m.EnableDiscovery),
		EnableEmailWarnings:              flex.ExpandBoolPointer(m.EnableEmailWarnings),
		EnableIfmapPublishing:            flex.ExpandBoolPointer(m.EnableIfmapPublishing),
		EnableImmediateDiscovery:         flex.ExpandBoolPointer(m.EnableImmediateDiscovery),
		EnablePxeLeaseTime:               flex.ExpandBoolPointer(m.EnablePxeLeaseTime),
		EnableSnmpWarnings:               flex.ExpandBoolPointer(m.EnableSnmpWarnings),
		EndAddr:                          flex.ExpandIPv4Address(m.EndAddr),
		Exclude:                          flex.ExpandFrameworkListNestedBlock(ctx, m.Exclude, diags, ExpandRangeExclude),
		ExtAttrs:                         flex.ExpandMapStringAny(ctx, m.ExtAttrs, diags),
		FailoverAssociation:              flex.ExpandStringPointerNullAsEmpty(m.FailoverAssociation),
		FingerprintFilterRules:           flex.ExpandFrameworkListNestedBlock(ctx, m.FingerprintFilterRules, diags, ExpandRangeFingerprintFilterRules),
		HighWaterMark:                    flex.ExpandInt64Pointer(m.HighWaterMark),
		HighWaterMarkReset:               flex.ExpandInt64Pointer(m.HighWaterMarkReset),
		IgnoreDhcpOptionListRequest:      flex.ExpandBoolPointer(m.IgnoreDhcpOptionListRequest),
		IgnoreId:                         flex.ExpandStringPointerNullAsEmpty(m.IgnoreId),
		IgnoreMacAddresses:               flex.ExpandFrameworkListString(ctx, m.IgnoreMacAddresses, diags),
		KnownClients:                     flex.ExpandStringPointer(m.KnownClients),
		LeaseScavengeTime:                flex.ExpandInt64Pointer(m.LeaseScavengeTime),
		LogicFilterRules:                 flex.ExpandFrameworkListNestedBlock(ctx, m.LogicFilterRules, diags, ExpandRangeLogicFilterRules),
		LowWaterMark:                     flex.ExpandInt64Pointer(m.LowWaterMark),
		LowWaterMarkReset:                flex.ExpandInt64Pointer(m.LowWaterMarkReset),
		MacFilterRules:                   flex.ExpandFrameworkListNestedBlock(ctx, m.MacFilterRules, diags, ExpandRangeMacFilterRules),
		Member:                           ExpandRangeMember(ctx, m.Member, diags),
		MsOptions:                        flex.ExpandFrameworkListNestedBlock(ctx, m.MsOptions, diags, ExpandRangeMsOptions),
		MsServer:                         ExpandRangeMsServer(ctx, m.MsServer, diags),
		NacFilterRules:                   flex.ExpandFrameworkListNestedBlock(ctx, m.NacFilterRules, diags, ExpandRangeNacFilterRules),
		Name:                             flex.ExpandStringPointerNullAsEmpty(m.Name),
		Network:                          flex.ExpandIPv4Prefix(m.Network),
		NetworkView:                      flex.ExpandStringPointerNullAsEmpty(m.NetworkView),
		Nextserver:                       flex.ExpandStringPointer(m.Nextserver),
		OptionFilterRules:                flex.ExpandFrameworkListNestedBlock(ctx, m.OptionFilterRules, diags, ExpandRangeOptionFilterRules),
		Options:                          flex.ExpandFrameworkListNestedBlock(ctx, m.Options, diags, ExpandRangeOptions),
		PortControlBlackoutSetting:       ExpandRangePortControlBlackoutSetting(ctx, m.PortControlBlackoutSetting, diags),
		PxeLeaseTime:                     flex.ExpandInt64Pointer(m.PxeLeaseTime),
		RecycleLeases:                    flex.ExpandBoolPointer(m.RecycleLeases),
		RelayAgentFilterRules:            flex.ExpandFrameworkListNestedBlock(ctx, m.RelayAgentFilterRules, diags, ExpandRangeRelayAgentFilterRules),
		RestartIfNeeded:                  flex.ExpandBoolPointer(m.RestartIfNeeded),
		SamePortControlDiscoveryBlackout: flex.ExpandBoolPointer(m.SamePortControlDiscoveryBlackout),
		ServerAssociationType:            flex.ExpandStringPointerNullAsEmpty(m.ServerAssociationType),
		StartAddr:                        flex.ExpandIPv4Address(m.StartAddr),
		SubscribeSettings:                ExpandRangeSubscribeSettings(ctx, m.SubscribeSettings, diags),
		UnknownClients:                   flex.ExpandStringPointer(m.UnknownClients),
		UpdateDnsOnLeaseRenewal:          flex.ExpandBoolPointer(m.UpdateDnsOnLeaseRenewal),
	}
	if isCreate {
		ext.SplitMember = ExpandRangeSplitMember(ctx, m.SplitMember, diags)
		ext.SplitScopeExclusionPercent = flex.ExpandInt64Pointer(m.SplitScopeExclusionPercent)
		ext.Template = flex.ExpandStringPointer(m.Template)
	}
	return ext
}

// ApplyRangeNIOSUseFlags derives NIOS use flags from the raw config
// value(s) and writes them onto the core model. A flag is true when the user
// set any of its governed value fields in config.
func ApplyRangeNIOSUseFlags(ctx context.Context, config tfsdk.Config, obj *coremodel.Range, diags *diag.Diagnostics) {
	if obj == nil || obj.NIOS == nil {
		return
	}
	obj.NIOS.UseBlackoutSetting = flex.DeriveUseFlag(ctx, config, diags, path.Root("nios").AtName("discovery_blackout_setting"), path.Root("nios").AtName("port_control_blackout_setting"), path.Root("nios").AtName("same_port_control_discovery_blackout"))
	obj.NIOS.UseBootfile = flex.DeriveUseFlag(ctx, config, diags, path.Root("nios").AtName("bootfile"))
	obj.NIOS.UseBootserver = flex.DeriveUseFlag(ctx, config, diags, path.Root("nios").AtName("bootserver"))
	obj.NIOS.UseDdnsDomainname = flex.DeriveUseFlag(ctx, config, diags, path.Root("nios").AtName("ddns_domainname"))
	obj.NIOS.UseDdnsGenerateHostname = flex.DeriveUseFlag(ctx, config, diags, path.Root("nios").AtName("ddns_generate_hostname"))
	obj.NIOS.UseDenyBootp = flex.DeriveUseFlag(ctx, config, diags, path.Root("nios").AtName("deny_bootp"))
	obj.NIOS.UseDiscoveryBasicPollingSettings = flex.DeriveUseFlag(ctx, config, diags, path.Root("nios").AtName("discovery_basic_poll_settings"))
	obj.NIOS.UseEmailList = flex.DeriveUseFlag(ctx, config, diags, path.Root("nios").AtName("email_list"))
	obj.NIOS.UseEnableDdns = flex.DeriveUseFlag(ctx, config, diags, path.Root("nios").AtName("enable_ddns"))
	obj.NIOS.UseEnableDhcpThresholds = flex.DeriveUseFlag(ctx, config, diags, path.Root("nios").AtName("enable_dhcp_thresholds"))
	obj.NIOS.UseEnableDiscovery = flex.DeriveUseFlag(ctx, config, diags, path.Root("nios").AtName("discovery_member"), path.Root("nios").AtName("enable_discovery"))
	obj.NIOS.UseEnableIfmapPublishing = flex.DeriveUseFlag(ctx, config, diags, path.Root("nios").AtName("enable_ifmap_publishing"))
	obj.NIOS.UseIgnoreDhcpOptionListRequest = flex.DeriveUseFlag(ctx, config, diags, path.Root("nios").AtName("ignore_dhcp_option_list_request"))
	obj.NIOS.UseIgnoreId = flex.DeriveUseFlag(ctx, config, diags, path.Root("nios").AtName("ignore_id"))
	obj.NIOS.UseKnownClients = flex.DeriveUseFlag(ctx, config, diags, path.Root("nios").AtName("known_clients"))
	obj.NIOS.UseLeaseScavengeTime = flex.DeriveUseFlag(ctx, config, diags, path.Root("nios").AtName("lease_scavenge_time"))
	obj.NIOS.UseLogicFilterRules = flex.DeriveUseFlag(ctx, config, diags, path.Root("nios").AtName("logic_filter_rules"))
	obj.NIOS.UseMsOptions = flex.DeriveUseFlag(ctx, config, diags, path.Root("nios").AtName("ms_options"))
	obj.NIOS.UseNextserver = flex.DeriveUseFlag(ctx, config, diags, path.Root("nios").AtName("nextserver"))
	obj.NIOS.UseOptions = flex.DeriveUseFlag(ctx, config, diags, path.Root("nios").AtName("options"))
	obj.NIOS.UsePxeLeaseTime = flex.DeriveUseFlag(ctx, config, diags, path.Root("nios").AtName("pxe_lease_time"))
	obj.NIOS.UseRecycleLeases = flex.DeriveUseFlag(ctx, config, diags, path.Root("nios").AtName("recycle_leases"))
	obj.NIOS.UseSubscribeSettings = flex.DeriveUseFlag(ctx, config, diags, path.Root("nios").AtName("subscribe_settings"))
	obj.NIOS.UseUnknownClients = flex.DeriveUseFlag(ctx, config, diags, path.Root("nios").AtName("unknown_clients"))
	obj.NIOS.UseUpdateDnsOnLeaseRenewal = flex.DeriveUseFlag(ctx, config, diags, path.Root("nios").AtName("update_dns_on_lease_renewal"))
}

// Expand converts the UDDI TF model to the core model.
func (m *UDDIRangeModel) Expand(ctx context.Context, diags *diag.Diagnostics) *coremodel.UDDIRangeExt {
	return &coremodel.UDDIRangeExt{
		Comment:            flex.ExpandStringPointer(m.Comment),
		DhcpHost:           flex.ExpandStringPointer(m.DhcpHost),
		DhcpOptions:        flex.ExpandFrameworkListNestedBlock(ctx, m.DhcpOptions, diags, ExpandOptionItem),
		DisableDhcp:        flex.ExpandBoolPointer(m.DisableDhcp),
		End:                flex.ExpandString(m.End),
		ExclusionRanges:    flex.ExpandFrameworkListNestedBlock(ctx, m.ExclusionRanges, diags, ExpandExclusionRange),
		Filters:            flex.ExpandFrameworkListNestedBlock(ctx, m.Filters, diags, ExpandAccessFilter),
		InheritanceParent:  flex.ExpandStringPointer(m.InheritanceParent),
		InheritanceSources: ExpandDHCPOptionsInheritance(ctx, m.InheritanceSources, diags),
		Name:               flex.ExpandStringPointer(m.Name),
		Parent:             flex.ExpandStringPointer(m.Parent),
		Space:              flex.ExpandStringPointer(m.Space),
		Start:              flex.ExpandString(m.Start),
		Tags:               flex.ExpandMapStringAny(ctx, m.Tags, diags),
		Threshold:          ExpandUtilizationThreshold(ctx, m.Threshold, diags),
	}
}

// Flatten populates the TF model from a core response.
func (m *RangeModel) Flatten(ctx context.Context, resp *coremodel.Range, diags *diag.Diagnostics) {
	if resp == nil {
		return
	}

	m.Id = flex.FlattenStringPointer(resp.Id)

	// Extract existing NIOS model, flatten API response onto it, convert back
	niosModel := flex.ExpandNestedObject[NIOSRangeModel](ctx, m.NIOS, diags)
	if niosModel == nil {
		niosModel = &NIOSRangeModel{}
	}
	plannedNIOS := flex.ExpandNestedObject[NIOSRangeModel](ctx, m.NIOS, diags)
	niosModel.Flatten(ctx, resp.NIOS, diags)
	if resp.NIOS != nil {
		PostFlattenRangeNIOS(ctx, plannedNIOS, niosModel, diags)
		m.NIOS = flex.FlattenNestedObject(ctx, niosModel, NIOSRangeAttrTypes, diags)
	} else {
		m.NIOS = types.ObjectNull(NIOSRangeAttrTypes)
	}

	// Extract existing UDDI model, flatten API response onto it, convert back
	uddiModel := flex.ExpandNestedObject[UDDIRangeModel](ctx, m.UDDI, diags)
	if uddiModel == nil {
		uddiModel = &UDDIRangeModel{}
	}
	uddiModel.Flatten(ctx, resp.UDDI, diags)
	if resp.UDDI != nil {
		m.UDDI = flex.FlattenNestedObject(ctx, uddiModel, UDDIRangeAttrTypes, diags)
	} else {
		m.UDDI = types.ObjectNull(UDDIRangeAttrTypes)
	}
}

// Flatten merges API response onto existing NIOS model.
func (m *NIOSRangeModel) Flatten(ctx context.Context, from *coremodel.NIOSRangeExt, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	planExtAttrs := m.ExtAttrs
	if planExtAttrs.IsUnknown() {
		planExtAttrs = types.MapNull(types.StringType)
	}
	m.AlwaysUpdateDns = flex.FlattenBoolPointer(from.AlwaysUpdateDns)
	m.Bootfile = flex.FlattenStringPointerEmptyAsNull(from.Bootfile)
	m.Bootserver = flex.FlattenStringPointerEmptyAsNull(from.Bootserver)
	m.CloudInfo = FlattenRangeCloudInfo(ctx, from.CloudInfo, diags)
	m.Comment = flex.FlattenStringPointerEmptyAsNull(from.Comment)
	m.DdnsDomainname = flex.FlattenStringPointerEmptyAsNull(from.DdnsDomainname)
	m.DdnsGenerateHostname = flex.FlattenBoolPointer(from.DdnsGenerateHostname)
	m.DenyAllClients = flex.FlattenBoolPointer(from.DenyAllClients)
	m.DenyBootp = flex.FlattenBoolPointer(from.DenyBootp)
	m.Disable = flex.FlattenBoolPointer(from.Disable)
	m.DiscoveryBasicPollSettings = FlattenRangeDiscoveryBasicPollSettings(ctx, from.DiscoveryBasicPollSettings, diags)
	m.DiscoveryBlackoutSetting = FlattenRangeDiscoveryBlackoutSetting(ctx, from.DiscoveryBlackoutSetting, diags)
	m.DiscoveryMember = flex.FlattenStringPointerEmptyAsNull(from.DiscoveryMember)
	m.EmailList = flex.FlattenFrameworkUnorderedListString(ctx, from.EmailList, diags)
	m.EnableDdns = flex.FlattenBoolPointer(from.EnableDdns)
	m.EnableDhcpThresholds = flex.FlattenBoolPointer(from.EnableDhcpThresholds)
	m.EnableDiscovery = flex.FlattenBoolPointer(from.EnableDiscovery)
	m.EnableEmailWarnings = flex.FlattenBoolPointer(from.EnableEmailWarnings)
	m.EnableIfmapPublishing = flex.FlattenBoolPointer(from.EnableIfmapPublishing)
	m.EnableImmediateDiscovery = flex.FlattenBoolPointer(from.EnableImmediateDiscovery)
	m.EnablePxeLeaseTime = flex.FlattenBoolPointer(from.EnablePxeLeaseTime)
	m.EnableSnmpWarnings = flex.FlattenBoolPointer(from.EnableSnmpWarnings)
	m.EndAddr = flex.FlattenIPv4Address(from.EndAddr)
	m.Exclude = flex.FlattenFrameworkListNestedBlock(ctx, from.Exclude, RangeExcludeAttrTypes, diags, FlattenRangeExclude)
	m.ExtAttrs, m.ExtAttrsAll = flex.FlattenEAs(planExtAttrs, from.ExtAttrs)
	m.FailoverAssociation = flex.FlattenStringPointerEmptyAsNull(from.FailoverAssociation)
	m.FingerprintFilterRules = flex.FlattenFrameworkListNestedBlock(ctx, from.FingerprintFilterRules, RangeFingerprintFilterRulesAttrTypes, diags, FlattenRangeFingerprintFilterRules)
	m.HighWaterMark = flex.FlattenInt64Pointer(from.HighWaterMark)
	m.HighWaterMarkReset = flex.FlattenInt64Pointer(from.HighWaterMarkReset)
	m.IgnoreDhcpOptionListRequest = flex.FlattenBoolPointer(from.IgnoreDhcpOptionListRequest)
	m.IgnoreId = flex.FlattenStringPointerEmptyAsNull(from.IgnoreId)
	m.IgnoreMacAddresses = flex.FlattenFrameworkUnorderedListString(ctx, from.IgnoreMacAddresses, diags)
	m.KnownClients = flex.FlattenStringPointerEmptyAsNull(from.KnownClients)
	m.LeaseScavengeTime = flex.FlattenInt64Pointer(from.LeaseScavengeTime)
	m.LogicFilterRules = flex.FlattenFrameworkListNestedBlock(ctx, from.LogicFilterRules, RangeLogicFilterRulesAttrTypes, diags, FlattenRangeLogicFilterRules)
	m.LowWaterMark = flex.FlattenInt64Pointer(from.LowWaterMark)
	m.LowWaterMarkReset = flex.FlattenInt64Pointer(from.LowWaterMarkReset)
	m.MacFilterRules = flex.FlattenFrameworkListNestedBlock(ctx, from.MacFilterRules, RangeMacFilterRulesAttrTypes, diags, FlattenRangeMacFilterRules)
	m.Member = FlattenRangeMember(ctx, from.Member, diags)
	m.MsOptions = flex.FlattenFrameworkListNestedBlock(ctx, from.MsOptions, RangeMsOptionsAttrTypes, diags, FlattenRangeMsOptions)
	m.MsServer = FlattenRangeMsServer(ctx, from.MsServer, diags)
	m.NacFilterRules = flex.FlattenFrameworkListNestedBlock(ctx, from.NacFilterRules, RangeNacFilterRulesAttrTypes, diags, FlattenRangeNacFilterRules)
	m.Name = flex.FlattenStringPointerEmptyAsNull(from.Name)
	m.Network = flex.FlattenIPv4Prefix(from.Network)
	m.NetworkView = flex.FlattenStringPointerEmptyAsNull(from.NetworkView)
	m.Nextserver = flex.FlattenStringPointerEmptyAsNull(from.Nextserver)
	m.OptionFilterRules = flex.FlattenFrameworkListNestedBlock(ctx, from.OptionFilterRules, RangeOptionFilterRulesAttrTypes, diags, FlattenRangeOptionFilterRules)
	m.Options = flex.FlattenFrameworkListNestedBlock(ctx, from.Options, RangeOptionsAttrTypes, diags, FlattenRangeOptions)
	m.PortControlBlackoutSetting = FlattenRangePortControlBlackoutSetting(ctx, from.PortControlBlackoutSetting, diags)
	m.PxeLeaseTime = flex.FlattenInt64Pointer(from.PxeLeaseTime)
	m.RecycleLeases = flex.FlattenBoolPointer(from.RecycleLeases)
	m.RelayAgentFilterRules = flex.FlattenFrameworkListNestedBlock(ctx, from.RelayAgentFilterRules, RangeRelayAgentFilterRulesAttrTypes, diags, FlattenRangeRelayAgentFilterRules)
	m.RestartIfNeeded = flex.FlattenBoolPointer(from.RestartIfNeeded)
	m.SamePortControlDiscoveryBlackout = flex.FlattenBoolPointer(from.SamePortControlDiscoveryBlackout)
	m.ServerAssociationType = flex.FlattenStringPointerEmptyAsNull(from.ServerAssociationType)
	m.SplitMember = FlattenRangeSplitMember(ctx, from.SplitMember, diags)
	m.SplitScopeExclusionPercent = flex.FlattenInt64Pointer(from.SplitScopeExclusionPercent)
	m.StartAddr = flex.FlattenIPv4Address(from.StartAddr)
	m.SubscribeSettings = FlattenRangeSubscribeSettings(ctx, from.SubscribeSettings, diags)
	m.Template = flex.FlattenStringPointerEmptyAsNull(from.Template)
	m.UnknownClients = flex.FlattenStringPointerEmptyAsNull(from.UnknownClients)
	m.UpdateDnsOnLeaseRenewal = flex.FlattenBoolPointer(from.UpdateDnsOnLeaseRenewal)
}

// Flatten merges API response onto existing UDDI model.
func (m *UDDIRangeModel) Flatten(ctx context.Context, from *coremodel.UDDIRangeExt, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.Comment = flex.FlattenStringPointer(from.Comment)
	m.DhcpHost = flex.FlattenStringPointer(from.DhcpHost)
	m.DhcpOptions = flex.FlattenFrameworkListNestedBlock(ctx, from.DhcpOptions, OptionItemAttrTypes, diags, FlattenOptionItem)
	m.DisableDhcp = flex.FlattenBoolPointer(from.DisableDhcp)
	m.End = flex.FlattenString(from.End)
	m.ExclusionRanges = flex.FlattenFrameworkListNestedBlock(ctx, from.ExclusionRanges, ExclusionRangeAttrTypes, diags, FlattenExclusionRange)
	m.Filters = flex.FlattenFrameworkListNestedBlock(ctx, from.Filters, AccessFilterAttrTypes, diags, FlattenAccessFilter)
	m.InheritanceParent = flex.FlattenStringPointer(from.InheritanceParent)
	m.InheritanceSources = FlattenDHCPOptionsInheritance(ctx, from.InheritanceSources, diags)
	m.Name = flex.FlattenStringPointer(from.Name)
	m.Parent = flex.FlattenStringPointer(from.Parent)
	m.Space = flex.FlattenStringPointer(from.Space)
	m.Start = flex.FlattenString(from.Start)
	tagsAll := flex.FlattenMapStringAny(ctx, from.Tags, diags)
	if m.Tags.IsNull() || m.Tags.IsUnknown() {
		m.Tags = tagsAll
	}
	m.TagsAll = tagsAll
	m.Threshold = FlattenUtilizationThreshold(ctx, from.Threshold, diags)
}
