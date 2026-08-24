package ipam

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-nettypes/cidrtypes"
	"github.com/hashicorp/terraform-plugin-framework-nettypes/iptypes"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/boolvalidator"
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

type NetworkcontainerModel struct {
	Id   types.String `tfsdk:"id"`
	NIOS types.Object `tfsdk:"nios"`
	UDDI types.Object `tfsdk:"uddi"`
}

var NetworkcontainerAttrTypes = map[string]attr.Type{
	"id":   types.StringType,
	"nios": types.ObjectType{AttrTypes: NIOSNetworkcontainerAttrTypes},
	"uddi": types.ObjectType{AttrTypes: UDDINetworkcontainerAttrTypes},
}

type NIOSNetworkcontainerModel struct {
	Authority                        types.Bool           `tfsdk:"authority"`
	AutoCreateReversezone            types.Bool           `tfsdk:"auto_create_reversezone"`
	Bootfile                         types.String         `tfsdk:"bootfile"`
	Bootserver                       types.String         `tfsdk:"bootserver"`
	CloudInfo                        types.Object         `tfsdk:"cloud_info"`
	Comment                          types.String         `tfsdk:"comment"`
	DdnsDomainname                   types.String         `tfsdk:"ddns_domainname"`
	DdnsGenerateHostname             types.Bool           `tfsdk:"ddns_generate_hostname"`
	DdnsServerAlwaysUpdates          types.Bool           `tfsdk:"ddns_server_always_updates"`
	DdnsTtl                          types.Int64          `tfsdk:"ddns_ttl"`
	DdnsUpdateFixedAddresses         types.Bool           `tfsdk:"ddns_update_fixed_addresses"`
	DdnsUseOption81                  types.Bool           `tfsdk:"ddns_use_option81"`
	DeleteReason                     types.String         `tfsdk:"delete_reason"`
	DenyBootp                        types.Bool           `tfsdk:"deny_bootp"`
	DiscoveryBasicPollSettings       types.Object         `tfsdk:"discovery_basic_poll_settings"`
	DiscoveryBlackoutSetting         types.Object         `tfsdk:"discovery_blackout_setting"`
	DiscoveryMember                  types.String         `tfsdk:"discovery_member"`
	EmailList                        types.List           `tfsdk:"email_list"`
	EnableDdns                       types.Bool           `tfsdk:"enable_ddns"`
	EnableDhcpThresholds             types.Bool           `tfsdk:"enable_dhcp_thresholds"`
	EnableDiscovery                  types.Bool           `tfsdk:"enable_discovery"`
	EnableEmailWarnings              types.Bool           `tfsdk:"enable_email_warnings"`
	EnableImmediateDiscovery         types.Bool           `tfsdk:"enable_immediate_discovery"`
	EnablePxeLeaseTime               types.Bool           `tfsdk:"enable_pxe_lease_time"`
	EnableSnmpWarnings               types.Bool           `tfsdk:"enable_snmp_warnings"`
	ExtAttrs                         types.Map            `tfsdk:"ext_attrs"`
	ExtAttrsAll                      types.Map            `tfsdk:"ext_attrs_all"`
	FederatedRealms                  types.List           `tfsdk:"federated_realms"`
	HighWaterMark                    types.Int64          `tfsdk:"high_water_mark"`
	HighWaterMarkReset               types.Int64          `tfsdk:"high_water_mark_reset"`
	IgnoreDhcpOptionListRequest      types.Bool           `tfsdk:"ignore_dhcp_option_list_request"`
	IgnoreId                         types.String         `tfsdk:"ignore_id"`
	IgnoreMacAddresses               types.List           `tfsdk:"ignore_mac_addresses"`
	IpamEmailAddresses               types.List           `tfsdk:"ipam_email_addresses"`
	IpamThresholdSettings            types.Object         `tfsdk:"ipam_threshold_settings"`
	IpamTrapSettings                 types.Object         `tfsdk:"ipam_trap_settings"`
	LeaseScavengeTime                types.Int64          `tfsdk:"lease_scavenge_time"`
	LogicFilterRules                 types.List           `tfsdk:"logic_filter_rules"`
	LowWaterMark                     types.Int64          `tfsdk:"low_water_mark"`
	LowWaterMarkReset                types.Int64          `tfsdk:"low_water_mark_reset"`
	MgmPrivate                       types.Bool           `tfsdk:"mgm_private"`
	Network                          cidrtypes.IPv4Prefix `tfsdk:"network"`
	NetworkView                      types.String         `tfsdk:"network_view"`
	Nextserver                       types.String         `tfsdk:"nextserver"`
	Options                          types.List           `tfsdk:"options"`
	PortControlBlackoutSetting       types.Object         `tfsdk:"port_control_blackout_setting"`
	PxeLeaseTime                     types.Int64          `tfsdk:"pxe_lease_time"`
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
	ZoneAssociations                 types.List           `tfsdk:"zone_associations"`
	DynamicAllocation                types.Object         `tfsdk:"dynamic_allocation"`
}

var NIOSNetworkcontainerAttrTypes = map[string]attr.Type{
	"authority":                            types.BoolType,
	"auto_create_reversezone":              types.BoolType,
	"bootfile":                             types.StringType,
	"bootserver":                           types.StringType,
	"cloud_info":                           types.ObjectType{AttrTypes: NetworkcontainerCloudInfoAttrTypes},
	"comment":                              types.StringType,
	"ddns_domainname":                      types.StringType,
	"ddns_generate_hostname":               types.BoolType,
	"ddns_server_always_updates":           types.BoolType,
	"ddns_ttl":                             types.Int64Type,
	"ddns_update_fixed_addresses":          types.BoolType,
	"ddns_use_option81":                    types.BoolType,
	"delete_reason":                        types.StringType,
	"deny_bootp":                           types.BoolType,
	"discovery_basic_poll_settings":        types.ObjectType{AttrTypes: NetworkcontainerDiscoveryBasicPollSettingsAttrTypes},
	"discovery_blackout_setting":           types.ObjectType{AttrTypes: NetworkcontainerDiscoveryBlackoutSettingAttrTypes},
	"discovery_member":                     types.StringType,
	"email_list":                           types.ListType{ElemType: types.StringType},
	"enable_ddns":                          types.BoolType,
	"enable_dhcp_thresholds":               types.BoolType,
	"enable_discovery":                     types.BoolType,
	"enable_email_warnings":                types.BoolType,
	"enable_immediate_discovery":           types.BoolType,
	"enable_pxe_lease_time":                types.BoolType,
	"enable_snmp_warnings":                 types.BoolType,
	"ext_attrs":                            types.MapType{ElemType: types.StringType},
	"ext_attrs_all":                        types.MapType{ElemType: types.StringType},
	"federated_realms":                     types.ListType{ElemType: types.ObjectType{AttrTypes: NetworkcontainerFederatedRealmsAttrTypes}},
	"high_water_mark":                      types.Int64Type,
	"high_water_mark_reset":                types.Int64Type,
	"ignore_dhcp_option_list_request":      types.BoolType,
	"ignore_id":                            types.StringType,
	"ignore_mac_addresses":                 types.ListType{ElemType: types.StringType},
	"ipam_email_addresses":                 types.ListType{ElemType: types.StringType},
	"ipam_threshold_settings":              types.ObjectType{AttrTypes: NetworkcontainerIpamThresholdSettingsAttrTypes},
	"ipam_trap_settings":                   types.ObjectType{AttrTypes: NetworkcontainerIpamTrapSettingsAttrTypes},
	"lease_scavenge_time":                  types.Int64Type,
	"logic_filter_rules":                   types.ListType{ElemType: types.ObjectType{AttrTypes: NetworkcontainerLogicFilterRulesAttrTypes}},
	"low_water_mark":                       types.Int64Type,
	"low_water_mark_reset":                 types.Int64Type,
	"mgm_private":                          types.BoolType,
	"network":                              cidrtypes.IPv4PrefixType{},
	"network_view":                         types.StringType,
	"nextserver":                           types.StringType,
	"options":                              types.ListType{ElemType: types.ObjectType{AttrTypes: NetworkcontainerOptionsAttrTypes}},
	"port_control_blackout_setting":        types.ObjectType{AttrTypes: NetworkcontainerPortControlBlackoutSettingAttrTypes},
	"pxe_lease_time":                       types.Int64Type,
	"recycle_leases":                       types.BoolType,
	"restart_if_needed":                    types.BoolType,
	"rir_organization":                     types.StringType,
	"rir_registration_action":              types.StringType,
	"rir_registration_status":              types.StringType,
	"same_port_control_discovery_blackout": types.BoolType,
	"send_rir_request":                     types.BoolType,
	"subscribe_settings":                   types.ObjectType{AttrTypes: NetworkcontainerSubscribeSettingsAttrTypes},
	"unmanaged":                            types.BoolType,
	"update_dns_on_lease_renewal":          types.BoolType,
	"zone_associations":                    types.ListType{ElemType: types.ObjectType{AttrTypes: NetworkcontainerZoneAssociationsAttrTypes}},
	"dynamic_allocation":                   types.ObjectType{AttrTypes: dynamicallocation.NextAvailableNetworkAttrTypes},
}

type UDDINetworkcontainerModel struct {
	Address                    iptypes.IPv4Address              `tfsdk:"address"`
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
	HeaderOptionFilename       types.String                     `tfsdk:"header_option_filename"`
	HeaderOptionServerAddress  types.String                     `tfsdk:"header_option_server_address"`
	HeaderOptionServerName     types.String                     `tfsdk:"header_option_server_name"`
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
	DynamicAllocation          types.Object                     `tfsdk:"dynamic_allocation"`
}

var UDDINetworkcontainerAttrTypes = map[string]attr.Type{
	"address":                       iptypes.IPv4AddressType{},
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
	"dhcp_config":                   types.ObjectType{AttrTypes: NetworkcontainerDHCPConfigAttrTypes},
	"dhcp_options":                  types.ListType{ElemType: types.ObjectType{AttrTypes: OptionItemAttrTypes}},
	"external_keys":                 types.MapType{ElemType: types.StringType},
	"federated_realms":              internaltypes.UnorderedListOfStringType,
	"header_option_filename":        types.StringType,
	"header_option_server_address":  types.StringType,
	"header_option_server_name":     types.StringType,
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
	"dynamic_allocation":            types.ObjectType{AttrTypes: dynamicallocation.NextAvailableAddressBlockAttrTypes},
}

const (
	NetworkcontainerInheritanceType = "full"
	NetworkcontainerReturnFields    = "authority,bootfile,bootserver,cloud_info,comment,ddns_domainname,ddns_generate_hostname,ddns_server_always_updates,ddns_ttl,ddns_update_fixed_addresses,ddns_use_option81,deny_bootp,discover_now_status,discovery_basic_poll_settings,discovery_blackout_setting,discovery_engine_type,discovery_member,email_list,enable_ddns,enable_dhcp_thresholds,enable_discovery,enable_email_warnings,enable_pxe_lease_time,enable_snmp_warnings,endpoint_sources,extattrs,federated_realms,high_water_mark,high_water_mark_reset,ignore_dhcp_option_list_request,ignore_id,ignore_mac_addresses,ipam_email_addresses,ipam_threshold_settings,ipam_trap_settings,last_rir_registration_update_sent,last_rir_registration_update_status,lease_scavenge_time,logic_filter_rules,low_water_mark,low_water_mark_reset,mgm_private,mgm_private_overridable,ms_ad_user_data,network,network_container,network_view,nextserver,options,port_control_blackout_setting,pxe_lease_time,recycle_leases,rir,rir_organization,rir_registration_status,same_port_control_discovery_blackout,subscribe_settings,unmanaged,update_dns_on_lease_renewal,use_authority,use_blackout_setting,use_bootfile,use_bootserver,use_ddns_domainname,use_ddns_generate_hostname,use_ddns_ttl,use_ddns_update_fixed_addresses,use_ddns_use_option81,use_deny_bootp,use_discovery_basic_polling_settings,use_email_list,use_enable_ddns,use_enable_dhcp_thresholds,use_enable_discovery,use_ignore_dhcp_option_list_request,use_ignore_id,use_ipam_email_addresses,use_ipam_threshold_settings,use_ipam_trap_settings,use_lease_scavenge_time,use_logic_filter_rules,use_mgm_private,use_nextserver,use_options,use_pxe_lease_time,use_recycle_leases,use_subscribe_settings,use_update_dns_on_lease_renewal,use_zone_associations,utilization,zone_associations"
)

var NetworkcontainerResourceSchemaAttributes = map[string]schema.Attribute{
	"id": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The reference to the object.",
	},
	"nios": schema.SingleNestedAttribute{
		Optional:            true,
		MarkdownDescription: "NIOS backend-specific fields.",
		Attributes:          NetworkcontainerResourceNiosSchemaAttributes,
	},
	"uddi": schema.SingleNestedAttribute{
		Optional:            true,
		MarkdownDescription: "UDDI backend-specific fields.",
		Attributes:          NetworkcontainerResourceUddiSchemaAttributes,
	},
}

var NetworkcontainerResourceNiosSchemaAttributes = map[string]schema.Attribute{
	"authority": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Authority for the DHCP network container.",
	},
	"auto_create_reversezone": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "This flag controls whether reverse zones are automatically created when the network is added.",
	},
	"bootfile": schema.StringAttribute{
		Optional: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
			customvalidator.IsValidIPv4OrFQDN(),
		},
		MarkdownDescription: "The boot server IPv4 Address or name in FQDN format for the network container. You can specify the name and/or IP address of the boot server that the host needs to boot.",
	},
	"bootserver": schema.StringAttribute{
		Optional: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
			customvalidator.IsValidIPv4OrFQDN(),
		},
		MarkdownDescription: "The bootserver address for the network container. You can specify the name and/or IP address of the boot server that the host needs to boot. The boot server IPv4 Address or name in FQDN format.",
	},
	"cloud_info": schema.SingleNestedAttribute{
		Attributes:          NetworkcontainerCloudInfoResourceSchemaAttributes,
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Structure containing all cloud API related information for this object.",
	},
	"comment": schema.StringAttribute{
		Optional: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "Comment for the network container; maximum 256 characters.",
	},
	"ddns_domainname": schema.StringAttribute{
		Optional: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "The dynamic DNS domain name the appliance uses specifically for DDNS updates for this network container.",
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
			boolvalidator.AlsoRequires(path.MatchRelative().AtParent().AtName("ddns_use_option81")),
		},
		MarkdownDescription: "This field controls whether the DHCP server is allowed to update DNS, regardless of the DHCP client requests. Note that changes for this field take effect only if ddns_use_option81 is True.",
	},
	"ddns_ttl": schema.Int64Attribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The DNS update Time to Live (TTL) value of a DHCP network container object. The TTL is a 32-bit unsigned integer that represents the duration, in seconds, for which the update is cached. Zero indicates that the update is not cached.",
	},
	"ddns_update_fixed_addresses": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "By default, the DHCP server does not update DNS when it allocates a fixed address to a client. You can configure the DHCP server to update the A and PTR records of a client with a fixed address. When this feature is enabled and the DHCP server adds A and PTR records for a fixed address, the DHCP server never discards the records.",
	},
	"ddns_use_option81": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "The support for DHCP Option 81 at the network container level.",
	},
	"delete_reason": schema.StringAttribute{
		Optional: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "The reason for deleting the RIR registration request.",
	},
	"deny_bootp": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "If set to True, BOOTP settings are disabled and BOOTP requests will be denied.",
	},
	"discovery_basic_poll_settings": schema.SingleNestedAttribute{
		Attributes:          NetworkcontainerDiscoveryBasicPollSettingsResourceSchemaAttributes,
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "",
	},
	"discovery_blackout_setting": schema.SingleNestedAttribute{
		Attributes:          NetworkcontainerDiscoveryBlackoutSettingResourceSchemaAttributes,
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
	"email_list": schema.ListAttribute{
		ElementType: types.StringType,
		Optional:    true,
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
		},
		MarkdownDescription: "The e-mail lists to which the appliance sends DHCP threshold alarm e-mail messages.",
	},
	"enable_ddns": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "The dynamic DNS updates flag of a DHCP network container object. If set to True, the DHCP server sends DDNS updates to DNS servers in the same Grid, and to external DNS servers.",
	},
	"enable_dhcp_thresholds": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Determines if DHCP thresholds are enabled for the network container.",
	},
	"enable_discovery": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Determines whether a discovery is enabled or not for this network container. When this is set to False, the network container discovery is disabled.",
	},
	"enable_email_warnings": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Determines if DHCP threshold warnings are sent through email.",
	},
	"enable_immediate_discovery": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines if the discovery for the network container should be immediately enabled.",
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
			Attributes: NetworkcontainerFederatedRealmsResourceSchemaAttributes,
		},
		Optional: true,
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
		},
		MarkdownDescription: "This field contains the federated realms associated to this network container.",
	},
	"high_water_mark": schema.Int64Attribute{
		Optional:            true,
		Computed:            true,
		Default:             int64default.StaticInt64(95),
		MarkdownDescription: "The percentage of DHCP network container usage threshold above which network container usage is not expected and may warrant your attention. When the high watermark is reached, the Infoblox appliance generates a syslog message and sends a warning (if enabled). A number that specifies the percentage of allocated addresses. The range is from 1 to 100.",
	},
	"high_water_mark_reset": schema.Int64Attribute{
		Optional:            true,
		Computed:            true,
		Default:             int64default.StaticInt64(85),
		MarkdownDescription: "The percentage of DHCP network container usage below which the corresponding SNMP trap is reset. A number that specifies the percentage of allocated addresses. The range is from 1 to 100. The high watermark reset value must be lower than the high watermark value.",
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
		MarkdownDescription: "Indicates whether the appliance will ignore DHCP client IDs or MAC addresses.",
	},
	"ignore_mac_addresses": schema.ListAttribute{
		ElementType: types.StringType,
		Optional:    true,
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
		},
		MarkdownDescription: "A list of MAC addresses the appliance will ignore.",
	},
	"ipam_email_addresses": schema.ListAttribute{
		ElementType: types.StringType,
		Optional:    true,
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
		},
		MarkdownDescription: "The e-mail lists to which the appliance sends IPAM threshold alarm e-mail messages.",
	},
	"ipam_threshold_settings": schema.SingleNestedAttribute{
		Attributes:          NetworkcontainerIpamThresholdSettingsResourceSchemaAttributes,
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "",
	},
	"ipam_trap_settings": schema.SingleNestedAttribute{
		Attributes:          NetworkcontainerIpamTrapSettingsResourceSchemaAttributes,
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "",
	},
	"lease_scavenge_time": schema.Int64Attribute{
		Optional: true,
		Computed: true,
		Validators: []validator.Int64{
			int64validator.Any(int64validator.OneOf(-1), int64validator.Between(86400, 2147472000)),
		},
		MarkdownDescription: "An integer that specifies the period of time (in seconds) that frees and backs up leases remained in the database before they are automatically deleted. To disable lease scavenging, set the parameter to -1. The minimum positive value must be greater than 86400 seconds (1 day).",
	},
	"logic_filter_rules": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: NetworkcontainerLogicFilterRulesResourceSchemaAttributes,
		},
		Optional: true,
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
		},
		MarkdownDescription: "This field contains the logic filters to be applied on the this network container. This list corresponds to the match rules that are written to the dhcpd configuration file.",
	},
	"low_water_mark": schema.Int64Attribute{
		Optional: true,
		Computed: true,
		Default:  int64default.StaticInt64(0),
		Validators: []validator.Int64{
			int64validator.Any(int64validator.Between(0, 100)),
		},
		MarkdownDescription: "The percentage of DHCP network container usage below which the Infoblox appliance generates a syslog message and sends a warning (if enabled). A number that specifies the percentage of allocated addresses. The range is from 1 to 100.",
	},
	"low_water_mark_reset": schema.Int64Attribute{
		Optional: true,
		Computed: true,
		Default:  int64default.StaticInt64(10),
		Validators: []validator.Int64{
			int64validator.Any(int64validator.Between(1, 100)),
		},
		MarkdownDescription: "The percentage of DHCP network container usage threshold below which network container usage is not expected and may warrant your attention. When the low watermark is crossed, the Infoblox appliance generates a syslog message and sends a warning (if enabled). A number that specifies the percentage of allocated addresses. The range is from 1 to 100. The low watermark reset value must be higher than the low watermark value.",
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
		CustomType: cidrtypes.IPv4PrefixType{},
		PlanModifiers: []planmodifier.String{
			immutable.ImmutableString(),
		},
		Validators: []validator.String{
			stringvalidator.ExactlyOneOf(
				path.MatchRelative().AtParent().AtName("dynamic_allocation"),
			),
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "The network address in IPv4 Address/CIDR format. For regular expression searches, only the IPv4 Address portion is supported. Searches for the CIDR portion is always an exact match. For example, both network containers 10.0.0.0/8 and 20.1.0.0/16 are matched by expression '.0' and only 20.1.0.0/16 is matched by '.0/16'.",
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
	"nextserver": schema.StringAttribute{
		Optional: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
			customvalidator.IsValidIPv4OrFQDN(),
		},
		MarkdownDescription: "The name in FQDN and/or IPv4 Address of the next server that the host needs to boot.",
	},
	"options": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: NetworkcontainerOptionsResourceSchemaAttributes,
		},
		Optional: true,
		Computed: true,
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
		},
		MarkdownDescription: "An array of DHCP option dhcpoption structs that lists the DHCP options associated with the object.",
	},
	"port_control_blackout_setting": schema.SingleNestedAttribute{
		Attributes:          NetworkcontainerPortControlBlackoutSettingResourceSchemaAttributes,
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "",
	},
	"pxe_lease_time": schema.Int64Attribute{
		Optional: true,
		Validators: []validator.Int64{
			int64validator.Any(int64validator.Between(0, 4294967295)),
		},
		MarkdownDescription: "The PXE lease time value of a DHCP Network container object. Some hosts use PXE (Preboot Execution Environment) to boot remotely from a server. To better manage your IP resources, set a different lease time for PXE boot requests. You can configure the DHCP server to allocate an IP address with a shorter lease time to hosts that send PXE boot requests, so IP addresses are not leased longer than necessary. A 32-bit unsigned integer that represents the duration, in seconds, for which the update is cached. Zero indicates that the update is not cached.",
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
		MarkdownDescription: "The RIR organization assoicated with the network container.",
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
		MarkdownDescription: "The registration status of the network container in RIR.",
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
		Attributes:          NetworkcontainerSubscribeSettingsResourceSchemaAttributes,
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
	"zone_associations": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: NetworkcontainerZoneAssociationsResourceSchemaAttributes,
		},
		Optional: true,
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

var NetworkcontainerResourceUddiSchemaAttributes = map[string]schema.Attribute{
	"address": schema.StringAttribute{
		Optional:   true,
		Computed:   true,
		CustomType: iptypes.IPv4AddressType{},
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.RequiresReplaceIfConfigured(),
		},
		Validators: []validator.String{
			stringvalidator.ExactlyOneOf(
				path.MatchRelative().AtParent().AtName("dynamic_allocation"),
			),
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
		Attributes: NetworkcontainerDHCPConfigResourceSchemaAttributes,
		Optional:   true,
		Computed:   true,
		Default: objectdefault.StaticValue(types.ObjectValueMust(NetworkcontainerDHCPConfigAttrTypes, map[string]attr.Value{
			"allow_unknown":           types.BoolValue(true),
			"authoritative_dhcp":      types.BoolValue(false),
			"filters":                 types.ListNull(types.StringType),
			"filters_large_selection": types.ListNull(types.StringType),
			"ignore_client_uid":       types.BoolValue(false),
			"ignore_list":             internaltypes.NewUnorderedListValueNull(types.ObjectType{AttrTypes: IgnoreItemAttrTypes}),
			"lease_time":              types.Int64Value(3600),
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
	"header_option_filename": schema.StringAttribute{
		Default:             stringdefault.StaticString(""),
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The configuration for header option filename field.",
	},
	"header_option_server_address": schema.StringAttribute{
		Default:             stringdefault.StaticString(""),
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The configuration for header option server address field.",
	},
	"header_option_server_name": schema.StringAttribute{
		Default:             stringdefault.StaticString(""),
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The configuration for header option server name field.",
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
	"dynamic_allocation": schema.SingleNestedAttribute{
		Attributes:          dynamicallocation.NextAvailableAddressBlockResourceSchemaAttributes,
		Optional:            true,
		MarkdownDescription: "Dynamically allocate the next available address block from a parent scope. Mutually exclusive with the static \"address\" field.",
	},
}

// Expand converts the TF model to the infoblox core model
func (m *NetworkcontainerModel) Expand(ctx context.Context, diags *diag.Diagnostics, isCreate bool) *coremodel.Networkcontainer {
	if m == nil {
		return nil
	}

	obj := &coremodel.Networkcontainer{}

	// Expand NIOS nested attribute (returns nil if not present)
	niosModel := flex.ExpandNestedObject[NIOSNetworkcontainerModel](ctx, m.NIOS, diags)
	if niosModel != nil {
		obj.NIOS = niosModel.Expand(ctx, diags, isCreate)
	}

	// Expand UDDI nested attribute (returns nil if not present)
	uddiModel := flex.ExpandNestedObject[UDDINetworkcontainerModel](ctx, m.UDDI, diags)
	if uddiModel != nil {
		obj.UDDI = uddiModel.Expand(ctx, diags, isCreate)
	}

	return obj
}

// Expand converts the NIOS TF model to the core model.
func (m *NIOSNetworkcontainerModel) Expand(ctx context.Context, diags *diag.Diagnostics, isCreate bool) *coremodel.NIOSNetworkcontainerExt {
	ext := &coremodel.NIOSNetworkcontainerExt{
		Authority:                        flex.ExpandBoolPointer(m.Authority),
		Bootfile:                         flex.ExpandStringPointerNullAsEmpty(m.Bootfile),
		Bootserver:                       flex.ExpandStringPointerNullAsEmpty(m.Bootserver),
		CloudInfo:                        ExpandNetworkcontainerCloudInfo(ctx, m.CloudInfo, diags),
		Comment:                          flex.ExpandStringPointerNullAsEmpty(m.Comment),
		DdnsDomainname:                   flex.ExpandStringPointerNullAsEmpty(m.DdnsDomainname),
		DdnsGenerateHostname:             flex.ExpandBoolPointer(m.DdnsGenerateHostname),
		DdnsServerAlwaysUpdates:          flex.ExpandBoolPointer(m.DdnsServerAlwaysUpdates),
		DdnsTtl:                          flex.ExpandInt64Pointer(m.DdnsTtl),
		DdnsUpdateFixedAddresses:         flex.ExpandBoolPointer(m.DdnsUpdateFixedAddresses),
		DdnsUseOption81:                  flex.ExpandBoolPointer(m.DdnsUseOption81),
		DeleteReason:                     flex.ExpandStringPointerNullAsEmpty(m.DeleteReason),
		DenyBootp:                        flex.ExpandBoolPointer(m.DenyBootp),
		DiscoveryBasicPollSettings:       ExpandNetworkcontainerDiscoveryBasicPollSettings(ctx, m.DiscoveryBasicPollSettings, diags),
		DiscoveryBlackoutSetting:         ExpandNetworkcontainerDiscoveryBlackoutSetting(ctx, m.DiscoveryBlackoutSetting, diags),
		DiscoveryMember:                  flex.ExpandStringPointer(m.DiscoveryMember),
		EmailList:                        flex.ExpandFrameworkListString(ctx, m.EmailList, diags),
		EnableDdns:                       flex.ExpandBoolPointer(m.EnableDdns),
		EnableDhcpThresholds:             flex.ExpandBoolPointer(m.EnableDhcpThresholds),
		EnableDiscovery:                  flex.ExpandBoolPointer(m.EnableDiscovery),
		EnableEmailWarnings:              flex.ExpandBoolPointer(m.EnableEmailWarnings),
		EnableImmediateDiscovery:         flex.ExpandBoolPointer(m.EnableImmediateDiscovery),
		EnablePxeLeaseTime:               flex.ExpandBoolPointer(m.EnablePxeLeaseTime),
		EnableSnmpWarnings:               flex.ExpandBoolPointer(m.EnableSnmpWarnings),
		ExtAttrs:                         flex.ExpandMapStringAny(ctx, m.ExtAttrs, diags),
		FederatedRealms:                  flex.ExpandFrameworkListNestedBlock(ctx, m.FederatedRealms, diags, ExpandNetworkcontainerFederatedRealms),
		HighWaterMark:                    flex.ExpandInt64Pointer(m.HighWaterMark),
		HighWaterMarkReset:               flex.ExpandInt64Pointer(m.HighWaterMarkReset),
		IgnoreDhcpOptionListRequest:      flex.ExpandBoolPointer(m.IgnoreDhcpOptionListRequest),
		IgnoreId:                         flex.ExpandStringPointerNullAsEmpty(m.IgnoreId),
		IgnoreMacAddresses:               flex.ExpandFrameworkListString(ctx, m.IgnoreMacAddresses, diags),
		IpamEmailAddresses:               flex.ExpandFrameworkListString(ctx, m.IpamEmailAddresses, diags),
		IpamThresholdSettings:            ExpandNetworkcontainerIpamThresholdSettings(ctx, m.IpamThresholdSettings, diags),
		IpamTrapSettings:                 ExpandNetworkcontainerIpamTrapSettings(ctx, m.IpamTrapSettings, diags),
		LeaseScavengeTime:                flex.ExpandInt64Pointer(m.LeaseScavengeTime),
		LogicFilterRules:                 flex.ExpandFrameworkListNestedBlock(ctx, m.LogicFilterRules, diags, ExpandNetworkcontainerLogicFilterRules),
		LowWaterMark:                     flex.ExpandInt64Pointer(m.LowWaterMark),
		LowWaterMarkReset:                flex.ExpandInt64Pointer(m.LowWaterMarkReset),
		MgmPrivate:                       flex.ExpandBoolPointer(m.MgmPrivate),
		Nextserver:                       flex.ExpandStringPointerNullAsEmpty(m.Nextserver),
		Options:                          flex.ExpandFrameworkListNestedBlock(ctx, m.Options, diags, ExpandNetworkcontainerOptions),
		PortControlBlackoutSetting:       ExpandNetworkcontainerPortControlBlackoutSetting(ctx, m.PortControlBlackoutSetting, diags),
		PxeLeaseTime:                     flex.ExpandInt64Pointer(m.PxeLeaseTime),
		RecycleLeases:                    flex.ExpandBoolPointer(m.RecycleLeases),
		RestartIfNeeded:                  flex.ExpandBoolPointer(m.RestartIfNeeded),
		RirOrganization:                  flex.ExpandStringPointer(m.RirOrganization),
		RirRegistrationAction:            flex.ExpandStringPointer(m.RirRegistrationAction),
		RirRegistrationStatus:            flex.ExpandStringPointerNullAsEmpty(m.RirRegistrationStatus),
		SamePortControlDiscoveryBlackout: flex.ExpandBoolPointer(m.SamePortControlDiscoveryBlackout),
		SendRirRequest:                   flex.ExpandBoolPointer(m.SendRirRequest),
		SubscribeSettings:                ExpandNetworkcontainerSubscribeSettings(ctx, m.SubscribeSettings, diags),
		Unmanaged:                        flex.ExpandBoolPointer(m.Unmanaged),
		UpdateDnsOnLeaseRenewal:          flex.ExpandBoolPointer(m.UpdateDnsOnLeaseRenewal),
		ZoneAssociations:                 flex.ExpandFrameworkListNestedBlock(ctx, m.ZoneAssociations, diags, ExpandNetworkcontainerZoneAssociations),
	}
	if isCreate {
		ext.Network = flex.ExpandIPv4Prefix(m.Network)
		ext.NetworkView = flex.ExpandStringPointerNullAsEmpty(m.NetworkView)
		ext.FuncCall = BuildNetworkcontainerFuncCall(ctx, m.DynamicAllocation, diags)
	}
	return ext
}

// ApplyNetworkcontainerNIOSUseFlags derives NIOS use flags from the raw config
// value(s) and writes them onto the core model. A flag is true when the user
// set any of its governed value fields in config.
func ApplyNetworkcontainerNIOSUseFlags(ctx context.Context, config tfsdk.Config, obj *coremodel.Networkcontainer, diags *diag.Diagnostics) {
	if obj == nil || obj.NIOS == nil {
		return
	}
	obj.NIOS.UseAuthority = flex.DeriveUseFlag(ctx, config, diags, path.Root("nios").AtName("authority"))
	obj.NIOS.UseBlackoutSetting = flex.DeriveUseFlag(ctx, config, diags, path.Root("nios").AtName("discovery_blackout_setting"), path.Root("nios").AtName("port_control_blackout_setting"), path.Root("nios").AtName("same_port_control_discovery_blackout"))
	obj.NIOS.UseBootfile = flex.DeriveUseFlag(ctx, config, diags, path.Root("nios").AtName("bootfile"))
	obj.NIOS.UseBootserver = flex.DeriveUseFlag(ctx, config, diags, path.Root("nios").AtName("bootserver"))
	obj.NIOS.UseDdnsDomainname = flex.DeriveUseFlag(ctx, config, diags, path.Root("nios").AtName("ddns_domainname"))
	obj.NIOS.UseDdnsGenerateHostname = flex.DeriveUseFlag(ctx, config, diags, path.Root("nios").AtName("ddns_generate_hostname"))
	obj.NIOS.UseDdnsTtl = flex.DeriveUseFlag(ctx, config, diags, path.Root("nios").AtName("ddns_ttl"))
	obj.NIOS.UseDdnsUpdateFixedAddresses = flex.DeriveUseFlag(ctx, config, diags, path.Root("nios").AtName("ddns_update_fixed_addresses"))
	obj.NIOS.UseDdnsUseOption81 = flex.DeriveUseFlag(ctx, config, diags, path.Root("nios").AtName("ddns_use_option81"))
	obj.NIOS.UseDenyBootp = flex.DeriveUseFlag(ctx, config, diags, path.Root("nios").AtName("deny_bootp"))
	obj.NIOS.UseDiscoveryBasicPollingSettings = flex.DeriveUseFlag(ctx, config, diags, path.Root("nios").AtName("discovery_basic_poll_settings"))
	obj.NIOS.UseEmailList = flex.DeriveUseFlag(ctx, config, diags, path.Root("nios").AtName("email_list"))
	obj.NIOS.UseEnableDdns = flex.DeriveUseFlag(ctx, config, diags, path.Root("nios").AtName("enable_ddns"))
	obj.NIOS.UseEnableDhcpThresholds = flex.DeriveUseFlag(ctx, config, diags, path.Root("nios").AtName("enable_dhcp_thresholds"))
	obj.NIOS.UseEnableDiscovery = flex.DeriveUseFlag(ctx, config, diags, path.Root("nios").AtName("discovery_member"), path.Root("nios").AtName("enable_discovery"))
	obj.NIOS.UseIgnoreDhcpOptionListRequest = flex.DeriveUseFlag(ctx, config, diags, path.Root("nios").AtName("ignore_dhcp_option_list_request"))
	obj.NIOS.UseIgnoreId = flex.DeriveUseFlag(ctx, config, diags, path.Root("nios").AtName("ignore_id"))
	obj.NIOS.UseIpamEmailAddresses = flex.DeriveUseFlag(ctx, config, diags, path.Root("nios").AtName("ipam_email_addresses"))
	obj.NIOS.UseIpamThresholdSettings = flex.DeriveUseFlag(ctx, config, diags, path.Root("nios").AtName("ipam_threshold_settings"))
	obj.NIOS.UseIpamTrapSettings = flex.DeriveUseFlag(ctx, config, diags, path.Root("nios").AtName("ipam_trap_settings"))
	obj.NIOS.UseLeaseScavengeTime = flex.DeriveUseFlag(ctx, config, diags, path.Root("nios").AtName("lease_scavenge_time"))
	obj.NIOS.UseLogicFilterRules = flex.DeriveUseFlag(ctx, config, diags, path.Root("nios").AtName("logic_filter_rules"))
	obj.NIOS.UseMgmPrivate = flex.DeriveUseFlag(ctx, config, diags, path.Root("nios").AtName("mgm_private"))
	obj.NIOS.UseNextserver = flex.DeriveUseFlag(ctx, config, diags, path.Root("nios").AtName("nextserver"))
	obj.NIOS.UseOptions = flex.DeriveUseFlag(ctx, config, diags, path.Root("nios").AtName("options"))
	obj.NIOS.UsePxeLeaseTime = flex.DeriveUseFlag(ctx, config, diags, path.Root("nios").AtName("pxe_lease_time"))
	obj.NIOS.UseRecycleLeases = flex.DeriveUseFlag(ctx, config, diags, path.Root("nios").AtName("recycle_leases"))
	obj.NIOS.UseSubscribeSettings = flex.DeriveUseFlag(ctx, config, diags, path.Root("nios").AtName("subscribe_settings"))
	obj.NIOS.UseUpdateDnsOnLeaseRenewal = flex.DeriveUseFlag(ctx, config, diags, path.Root("nios").AtName("update_dns_on_lease_renewal"))
	obj.NIOS.UseZoneAssociations = flex.DeriveUseFlag(ctx, config, diags, path.Root("nios").AtName("zone_associations"))
}

// Expand converts the UDDI TF model to the core model.
func (m *UDDINetworkcontainerModel) Expand(ctx context.Context, diags *diag.Diagnostics, isCreate bool) *coremodel.UDDINetworkcontainerExt {
	ext := &coremodel.UDDINetworkcontainerExt{
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
		DhcpConfig:                 ExpandNetworkcontainerDHCPConfig(ctx, m.DhcpConfig, diags),
		DhcpOptions:                flex.ExpandFrameworkListNestedBlock(ctx, m.DhcpOptions, diags, ExpandOptionItem),
		ExternalKeys:               flex.ExpandMapStringAny(ctx, m.ExternalKeys, diags),
		FederatedRealms:            flex.ExpandFrameworkListString(ctx, m.FederatedRealms, diags),
		HeaderOptionFilename:       flex.ExpandStringPointer(m.HeaderOptionFilename),
		HeaderOptionServerAddress:  flex.ExpandStringPointer(m.HeaderOptionServerAddress),
		HeaderOptionServerName:     flex.ExpandStringPointer(m.HeaderOptionServerName),
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
		ext.Address = flex.ExpandIPv4Address(m.Address)
		ext.Space = flex.ExpandStringPointer(m.Space)
		if alloc := BuildNetworkcontainerAllocation(ctx, m.DynamicAllocation, diags); alloc != nil {
			ext.Address = alloc
		}
	}
	return ext
}

// Flatten populates the TF model from a core response.
func (m *NetworkcontainerModel) Flatten(ctx context.Context, resp *coremodel.Networkcontainer, diags *diag.Diagnostics) {
	if resp == nil {
		return
	}

	m.Id = flex.FlattenStringPointer(resp.Id)

	// Extract existing NIOS model, flatten API response onto it, convert back
	niosModel := flex.ExpandNestedObject[NIOSNetworkcontainerModel](ctx, m.NIOS, diags)
	if niosModel == nil {
		niosModel = &NIOSNetworkcontainerModel{}
	}
	plannedNIOS := flex.ExpandNestedObject[NIOSNetworkcontainerModel](ctx, m.NIOS, diags)
	niosModel.Flatten(ctx, resp.NIOS, diags)
	if resp.NIOS != nil {
		PostFlattenNetworkcontainerNIOS(ctx, plannedNIOS, niosModel, diags)
		m.NIOS = flex.FlattenNestedObject(ctx, niosModel, NIOSNetworkcontainerAttrTypes, diags)
	} else {
		m.NIOS = types.ObjectNull(NIOSNetworkcontainerAttrTypes)
	}

	// Extract existing UDDI model, flatten API response onto it, convert back
	uddiModel := flex.ExpandNestedObject[UDDINetworkcontainerModel](ctx, m.UDDI, diags)
	if uddiModel == nil {
		uddiModel = &UDDINetworkcontainerModel{}
	}
	uddiModel.Flatten(ctx, resp.UDDI, diags)
	if resp.UDDI != nil {
		m.UDDI = flex.FlattenNestedObject(ctx, uddiModel, UDDINetworkcontainerAttrTypes, diags)
	} else {
		m.UDDI = types.ObjectNull(UDDINetworkcontainerAttrTypes)
	}
}

// Flatten merges API response onto existing NIOS model.
func (m *NIOSNetworkcontainerModel) Flatten(ctx context.Context, from *coremodel.NIOSNetworkcontainerExt, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	planExtAttrs := m.ExtAttrs
	if planExtAttrs.IsUnknown() {
		planExtAttrs = types.MapNull(types.StringType)
	}
	m.Authority = flex.FlattenBoolPointer(from.Authority)
	m.Bootfile = flex.FlattenStringPointerEmptyAsNull(from.Bootfile)
	m.Bootserver = flex.FlattenStringPointerEmptyAsNull(from.Bootserver)
	m.CloudInfo = FlattenNetworkcontainerCloudInfo(ctx, from.CloudInfo, diags)
	m.Comment = flex.FlattenStringPointerEmptyAsNull(from.Comment)
	m.DdnsDomainname = flex.FlattenStringPointerEmptyAsNull(from.DdnsDomainname)
	m.DdnsGenerateHostname = flex.FlattenBoolPointer(from.DdnsGenerateHostname)
	m.DdnsServerAlwaysUpdates = flex.FlattenBoolPointer(from.DdnsServerAlwaysUpdates)
	m.DdnsTtl = flex.FlattenInt64Pointer(from.DdnsTtl)
	m.DdnsUpdateFixedAddresses = flex.FlattenBoolPointer(from.DdnsUpdateFixedAddresses)
	m.DdnsUseOption81 = flex.FlattenBoolPointer(from.DdnsUseOption81)
	m.DenyBootp = flex.FlattenBoolPointer(from.DenyBootp)
	m.DiscoveryBasicPollSettings = FlattenNetworkcontainerDiscoveryBasicPollSettings(ctx, from.DiscoveryBasicPollSettings, diags)
	m.DiscoveryBlackoutSetting = FlattenNetworkcontainerDiscoveryBlackoutSetting(ctx, from.DiscoveryBlackoutSetting, diags)
	m.DiscoveryMember = flex.FlattenStringPointerEmptyAsNull(from.DiscoveryMember)
	m.EmailList = flex.FlattenFrameworkListString(ctx, from.EmailList, diags)
	m.EnableDdns = flex.FlattenBoolPointer(from.EnableDdns)
	m.EnableDhcpThresholds = flex.FlattenBoolPointer(from.EnableDhcpThresholds)
	m.EnableDiscovery = flex.FlattenBoolPointer(from.EnableDiscovery)
	m.EnableEmailWarnings = flex.FlattenBoolPointer(from.EnableEmailWarnings)
	m.EnablePxeLeaseTime = flex.FlattenBoolPointer(from.EnablePxeLeaseTime)
	m.EnableSnmpWarnings = flex.FlattenBoolPointer(from.EnableSnmpWarnings)
	m.ExtAttrs, m.ExtAttrsAll = flex.FlattenEAs(planExtAttrs, from.ExtAttrs)
	m.FederatedRealms = flex.FlattenFrameworkListNestedBlock(ctx, from.FederatedRealms, NetworkcontainerFederatedRealmsAttrTypes, diags, FlattenNetworkcontainerFederatedRealms)
	m.HighWaterMark = flex.FlattenInt64Pointer(from.HighWaterMark)
	m.HighWaterMarkReset = flex.FlattenInt64Pointer(from.HighWaterMarkReset)
	m.IgnoreDhcpOptionListRequest = flex.FlattenBoolPointer(from.IgnoreDhcpOptionListRequest)
	m.IgnoreId = flex.FlattenStringPointerEmptyAsNull(from.IgnoreId)
	m.IgnoreMacAddresses = flex.FlattenFrameworkListString(ctx, from.IgnoreMacAddresses, diags)
	m.IpamEmailAddresses = flex.FlattenFrameworkListString(ctx, from.IpamEmailAddresses, diags)
	m.IpamThresholdSettings = FlattenNetworkcontainerIpamThresholdSettings(ctx, from.IpamThresholdSettings, diags)
	m.IpamTrapSettings = FlattenNetworkcontainerIpamTrapSettings(ctx, from.IpamTrapSettings, diags)
	m.LeaseScavengeTime = flex.FlattenInt64Pointer(from.LeaseScavengeTime)
	m.LogicFilterRules = flex.FlattenFrameworkListNestedBlock(ctx, from.LogicFilterRules, NetworkcontainerLogicFilterRulesAttrTypes, diags, FlattenNetworkcontainerLogicFilterRules)
	m.LowWaterMark = flex.FlattenInt64Pointer(from.LowWaterMark)
	m.LowWaterMarkReset = flex.FlattenInt64Pointer(from.LowWaterMarkReset)
	m.MgmPrivate = flex.FlattenBoolPointer(from.MgmPrivate)
	m.Network = flex.FlattenIPv4Prefix(from.Network)
	m.NetworkView = flex.FlattenStringPointerEmptyAsNull(from.NetworkView)
	m.Nextserver = flex.FlattenStringPointerEmptyAsNull(from.Nextserver)
	m.Options = flex.FlattenFrameworkListNestedBlock(ctx, from.Options, NetworkcontainerOptionsAttrTypes, diags, FlattenNetworkcontainerOptions)
	m.PortControlBlackoutSetting = FlattenNetworkcontainerPortControlBlackoutSetting(ctx, from.PortControlBlackoutSetting, diags)
	m.PxeLeaseTime = flex.FlattenInt64Pointer(from.PxeLeaseTime)
	m.RecycleLeases = flex.FlattenBoolPointer(from.RecycleLeases)
	m.RirOrganization = flex.FlattenStringPointerEmptyAsNull(from.RirOrganization)
	m.RirRegistrationAction = flex.FlattenStringPointerEmptyAsNull(from.RirRegistrationAction)
	m.RirRegistrationStatus = flex.FlattenStringPointerEmptyAsNull(from.RirRegistrationStatus)
	m.SamePortControlDiscoveryBlackout = flex.FlattenBoolPointer(from.SamePortControlDiscoveryBlackout)
	m.SubscribeSettings = FlattenNetworkcontainerSubscribeSettings(ctx, from.SubscribeSettings, diags)
	m.Unmanaged = flex.FlattenBoolPointer(from.Unmanaged)
	m.UpdateDnsOnLeaseRenewal = flex.FlattenBoolPointer(from.UpdateDnsOnLeaseRenewal)
	m.ZoneAssociations = flex.FlattenFrameworkListNestedBlock(ctx, from.ZoneAssociations, NetworkcontainerZoneAssociationsAttrTypes, diags, FlattenNetworkcontainerZoneAssociations)
	if len(m.DynamicAllocation.AttributeTypes(ctx)) == 0 {
		m.DynamicAllocation = types.ObjectNull(dynamicallocation.NextAvailableNetworkAttrTypes)
	}
}

// Flatten merges API response onto existing UDDI model.
func (m *UDDINetworkcontainerModel) Flatten(ctx context.Context, from *coremodel.UDDINetworkcontainerExt, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.Address = flex.FlattenIPv4Address(from.Address)
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
	m.DhcpConfig = FlattenNetworkcontainerDHCPConfig(ctx, from.DhcpConfig, diags)
	m.DhcpOptions = flex.FlattenFrameworkListNestedBlock(ctx, from.DhcpOptions, OptionItemAttrTypes, diags, FlattenOptionItem)
	m.ExternalKeys = flex.FlattenMapStringAny(ctx, from.ExternalKeys, diags)
	m.FederatedRealms = flex.FlattenFrameworkUnorderedListString(ctx, from.FederatedRealms, diags)
	m.HeaderOptionFilename = flex.FlattenStringPointer(from.HeaderOptionFilename)
	m.HeaderOptionServerAddress = flex.FlattenStringPointer(from.HeaderOptionServerAddress)
	m.HeaderOptionServerName = flex.FlattenStringPointer(from.HeaderOptionServerName)
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
	if len(m.DynamicAllocation.AttributeTypes(ctx)) == 0 {
		m.DynamicAllocation = types.ObjectNull(dynamicallocation.NextAvailableAddressBlockAttrTypes)
	}
}
