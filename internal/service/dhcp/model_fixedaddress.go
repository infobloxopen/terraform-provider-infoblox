package dhcp

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-nettypes/iptypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/mapvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/mapdefault"
	objectplanmodifier "github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	stringplanmodifier "github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"

	coremodel "github.com/infobloxopen/terraform-provider-infoblox/internal/core/model/dhcp"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/dynamicallocation"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
	immutable "github.com/infobloxopen/terraform-provider-infoblox/internal/planmodifiers/immutable"
	importmod "github.com/infobloxopen/terraform-provider-infoblox/internal/planmodifiers/import"
	internaltypes "github.com/infobloxopen/terraform-provider-infoblox/internal/types"
	customvalidator "github.com/infobloxopen/terraform-provider-infoblox/internal/validator"
)

type FixedaddressModel struct {
	Id            types.String `tfsdk:"id"`
	UpdateTrigger types.String `tfsdk:"update_trigger"`
	NIOS          types.Object `tfsdk:"nios"`
	UDDI          types.Object `tfsdk:"uddi"`
}

var FixedaddressAttrTypes = map[string]attr.Type{
	"id":             types.StringType,
	"update_trigger": types.StringType,
	"nios":           types.ObjectType{AttrTypes: NIOSFixedaddressAttrTypes},
	"uddi":           types.ObjectType{AttrTypes: UDDIFixedaddressAttrTypes},
}

type NIOSFixedaddressModel struct {
	AgentCircuitId              types.String             `tfsdk:"agent_circuit_id"`
	AgentRemoteId               types.String             `tfsdk:"agent_remote_id"`
	AllowTelnet                 types.Bool               `tfsdk:"allow_telnet"`
	AlwaysUpdateDns             types.Bool               `tfsdk:"always_update_dns"`
	Bootfile                    types.String             `tfsdk:"bootfile"`
	Bootserver                  types.String             `tfsdk:"bootserver"`
	CliCredentials              types.List               `tfsdk:"cli_credentials"`
	ClientIdentifierPrependZero types.Bool               `tfsdk:"client_identifier_prepend_zero"`
	Comment                     types.String             `tfsdk:"comment"`
	DdnsDomainname              types.String             `tfsdk:"ddns_domainname"`
	DdnsHostname                types.String             `tfsdk:"ddns_hostname"`
	DenyBootp                   types.Bool               `tfsdk:"deny_bootp"`
	DeviceDescription           types.String             `tfsdk:"device_description"`
	DeviceLocation              types.String             `tfsdk:"device_location"`
	DeviceType                  types.String             `tfsdk:"device_type"`
	DeviceVendor                types.String             `tfsdk:"device_vendor"`
	DhcpClientIdentifier        types.String             `tfsdk:"dhcp_client_identifier"`
	Disable                     types.Bool               `tfsdk:"disable"`
	DisableDiscovery            types.Bool               `tfsdk:"disable_discovery"`
	EnableDdns                  types.Bool               `tfsdk:"enable_ddns"`
	EnableImmediateDiscovery    types.Bool               `tfsdk:"enable_immediate_discovery"`
	EnablePxeLeaseTime          types.Bool               `tfsdk:"enable_pxe_lease_time"`
	ExtAttrs                    types.Map                `tfsdk:"ext_attrs"`
	ExtAttrsAll                 types.Map                `tfsdk:"ext_attrs_all"`
	IgnoreDhcpOptionListRequest types.Bool               `tfsdk:"ignore_dhcp_option_list_request"`
	Ipv4addr                    iptypes.IPv4Address      `tfsdk:"ipv4addr"`
	LogicFilterRules            types.List               `tfsdk:"logic_filter_rules"`
	Mac                         internaltypes.MACAddress `tfsdk:"mac"`
	MatchClient                 types.String             `tfsdk:"match_client"`
	MsOptions                   types.List               `tfsdk:"ms_options"`
	MsServer                    types.Object             `tfsdk:"ms_server"`
	Name                        types.String             `tfsdk:"name"`
	Network                     types.String             `tfsdk:"network"`
	NetworkView                 types.String             `tfsdk:"network_view"`
	Nextserver                  types.String             `tfsdk:"nextserver"`
	Options                     types.List               `tfsdk:"options"`
	PxeLeaseTime                types.Int64              `tfsdk:"pxe_lease_time"`
	ReservedInterface           types.String             `tfsdk:"reserved_interface"`
	RestartIfNeeded             types.Bool               `tfsdk:"restart_if_needed"`
	Snmp3Credential             types.Object             `tfsdk:"snmp3_credential"`
	SnmpCredential              types.Object             `tfsdk:"snmp_credential"`
	Template                    types.String             `tfsdk:"template"`
	DynamicAllocation           types.Object             `tfsdk:"dynamic_allocation"`
}

var NIOSFixedaddressAttrTypes = map[string]attr.Type{
	"agent_circuit_id":                types.StringType,
	"agent_remote_id":                 types.StringType,
	"allow_telnet":                    types.BoolType,
	"always_update_dns":               types.BoolType,
	"bootfile":                        types.StringType,
	"bootserver":                      types.StringType,
	"cli_credentials":                 types.ListType{ElemType: types.ObjectType{AttrTypes: FixedaddressCliCredentialsAttrTypes}},
	"client_identifier_prepend_zero":  types.BoolType,
	"comment":                         types.StringType,
	"ddns_domainname":                 types.StringType,
	"ddns_hostname":                   types.StringType,
	"deny_bootp":                      types.BoolType,
	"device_description":              types.StringType,
	"device_location":                 types.StringType,
	"device_type":                     types.StringType,
	"device_vendor":                   types.StringType,
	"dhcp_client_identifier":          types.StringType,
	"disable":                         types.BoolType,
	"disable_discovery":               types.BoolType,
	"enable_ddns":                     types.BoolType,
	"enable_immediate_discovery":      types.BoolType,
	"enable_pxe_lease_time":           types.BoolType,
	"ext_attrs":                       types.MapType{ElemType: types.StringType},
	"ext_attrs_all":                   types.MapType{ElemType: types.StringType},
	"ignore_dhcp_option_list_request": types.BoolType,
	"ipv4addr":                        iptypes.IPv4AddressType{},
	"logic_filter_rules":              types.ListType{ElemType: types.ObjectType{AttrTypes: FixedaddressLogicFilterRulesAttrTypes}},
	"mac":                             internaltypes.MACAddressType{},
	"match_client":                    types.StringType,
	"ms_options":                      types.ListType{ElemType: types.ObjectType{AttrTypes: FixedaddressMsOptionsAttrTypes}},
	"ms_server":                       types.ObjectType{AttrTypes: FixedaddressMsServerAttrTypes},
	"name":                            types.StringType,
	"network":                         types.StringType,
	"network_view":                    types.StringType,
	"nextserver":                      types.StringType,
	"options":                         types.ListType{ElemType: types.ObjectType{AttrTypes: FixedaddressOptionsAttrTypes}},
	"pxe_lease_time":                  types.Int64Type,
	"reserved_interface":              types.StringType,
	"restart_if_needed":               types.BoolType,
	"snmp3_credential":                types.ObjectType{AttrTypes: FixedaddressSnmp3CredentialAttrTypes},
	"snmp_credential":                 types.ObjectType{AttrTypes: FixedaddressSnmpCredentialAttrTypes},
	"template":                        types.StringType,
	"dynamic_allocation":              types.ObjectType{AttrTypes: dynamicallocation.NextAvailableIpAttrTypes},
}

type UDDIFixedaddressModel struct {
	Address                   types.String `tfsdk:"address"`
	Comment                   types.String `tfsdk:"comment"`
	DhcpOptions               types.List   `tfsdk:"dhcp_options"`
	DisableDhcp               types.Bool   `tfsdk:"disable_dhcp"`
	HeaderOptionFilename      types.String `tfsdk:"header_option_filename"`
	HeaderOptionServerAddress types.String `tfsdk:"header_option_server_address"`
	HeaderOptionServerName    types.String `tfsdk:"header_option_server_name"`
	Hostname                  types.String `tfsdk:"hostname"`
	InheritanceParent         types.String `tfsdk:"inheritance_parent"`
	InheritanceSources        types.Object `tfsdk:"inheritance_sources"`
	IpSpace                   types.String `tfsdk:"ip_space"`
	MatchType                 types.String `tfsdk:"match_type"`
	MatchValue                types.String `tfsdk:"match_value"`
	Name                      types.String `tfsdk:"name"`
	Parent                    types.String `tfsdk:"parent"`
	Tags                      types.Map    `tfsdk:"tags"`
	TagsAll                   types.Map    `tfsdk:"tags_all"`
	DynamicAllocation         types.Object `tfsdk:"dynamic_allocation"`
}

var UDDIFixedaddressAttrTypes = map[string]attr.Type{
	"address":                      types.StringType,
	"comment":                      types.StringType,
	"dhcp_options":                 types.ListType{ElemType: types.ObjectType{AttrTypes: OptionItemAttrTypes}},
	"disable_dhcp":                 types.BoolType,
	"header_option_filename":       types.StringType,
	"header_option_server_address": types.StringType,
	"header_option_server_name":    types.StringType,
	"hostname":                     types.StringType,
	"inheritance_parent":           types.StringType,
	"inheritance_sources":          types.ObjectType{AttrTypes: FixedAddressInheritanceAttrTypes},
	"ip_space":                     types.StringType,
	"match_type":                   types.StringType,
	"match_value":                  types.StringType,
	"name":                         types.StringType,
	"parent":                       types.StringType,
	"tags":                         types.MapType{ElemType: types.StringType},
	"tags_all":                     types.MapType{ElemType: types.StringType},
	"dynamic_allocation":           types.ObjectType{AttrTypes: dynamicallocation.NextAvailableAddressAttrTypes},
}

const (
	FixedaddressInheritanceType = "full"
	FixedaddressReturnFields    = "agent_circuit_id,agent_remote_id,allow_telnet,always_update_dns,bootfile,bootserver,cli_credentials,client_identifier_prepend_zero,cloud_info,comment,ddns_domainname,ddns_hostname,deny_bootp,device_description,device_location,device_type,device_vendor,dhcp_client_identifier,disable,disable_discovery,discover_now_status,discovered_data,enable_ddns,enable_pxe_lease_time,extattrs,ignore_dhcp_option_list_request,ipv4addr,is_invalid_mac,logic_filter_rules,mac,match_client,ms_ad_user_data,ms_options,ms_server,name,network,network_view,nextserver,options,pxe_lease_time,reserved_interface,snmp3_credential,snmp_credential,use_bootfile,use_bootserver,use_cli_credentials,use_ddns_domainname,use_deny_bootp,use_enable_ddns,use_ignore_dhcp_option_list_request,use_logic_filter_rules,use_ms_options,use_nextserver,use_options,use_pxe_lease_time,use_snmp3_credential,use_snmp_credential"
)

var FixedaddressResourceSchemaAttributes = map[string]schema.Attribute{
	"id": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The reference to the object.",
	},
	"update_trigger": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "An arbitrary value used to trigger an update. Not sent to the API. Change it when Terraform reports no infrastructure changes.",
	},
	"nios": schema.SingleNestedAttribute{
		Optional:            true,
		MarkdownDescription: "NIOS backend-specific fields.",
		Attributes:          FixedaddressResourceNiosSchemaAttributes,
	},
	"uddi": schema.SingleNestedAttribute{
		Optional:            true,
		MarkdownDescription: "UDDI backend-specific fields.",
		Attributes:          FixedaddressResourceUddiSchemaAttributes,
	},
}

var FixedaddressResourceNiosSchemaAttributes = map[string]schema.Attribute{
	"agent_circuit_id": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "The agent circuit ID for the fixed address.",
	},
	"agent_remote_id": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "The agent remote ID for the fixed address.",
	},
	"allow_telnet": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "This field controls whether the credential is used for both the Telnet and SSH credentials. If set to False, the credential is used only for SSH.",
	},
	"always_update_dns": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "This field controls whether only the DHCP server is allowed to update DNS, regardless of the DHCP client requests.",
	},
	"bootfile": schema.StringAttribute{
		Optional: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "The bootfile name for the fixed address. You can configure the DHCP server to support clients that use the boot file name option in their DHCPREQUEST messages.",
	},
	"bootserver": schema.StringAttribute{
		Optional: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
			customvalidator.IsValidIPv4OrFQDN(),
		},
		MarkdownDescription: "The bootserver address for the fixed address. You can specify the name and/or IP address of the boot server that the host needs to boot. The boot server IPv4 Address or name in FQDN format.",
	},
	"cli_credentials": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: FixedaddressCliCredentialsResourceSchemaAttributes,
		},
		Optional: true,
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
		},
		MarkdownDescription: "The CLI credentials for the fixed address.",
	},
	"client_identifier_prepend_zero": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "This field controls whether there is a prepend for the dhcp-client-identifier of a fixed address.",
	},
	"comment": schema.StringAttribute{
		Optional: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
			customvalidator.ValidateTrimmedString(),
			stringvalidator.LengthBetween(0, 256),
		},
		MarkdownDescription: "Comment for the fixed address; maximum 256 characters.",
	},
	"ddns_domainname": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "The dynamic DNS domain name the appliance uses specifically for DDNS updates for this fixed address.",
	},
	"ddns_hostname": schema.StringAttribute{
		Optional: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "The DDNS host name for this fixed address.",
	},
	"deny_bootp": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "If set to true, BOOTP settings are disabled and BOOTP requests will be denied.",
	},
	"device_description": schema.StringAttribute{
		Optional: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "The description of the device.",
	},
	"device_location": schema.StringAttribute{
		Optional: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "The location of the device.",
	},
	"device_type": schema.StringAttribute{
		Optional: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "The type of the device.",
	},
	"device_vendor": schema.StringAttribute{
		Optional: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "The vendor of the device.",
	},
	"dhcp_client_identifier": schema.StringAttribute{
		Optional: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "The DHCP client ID for the fixed address.",
	},
	"disable": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Determines whether a fixed address is disabled or not. When this is set to False, the fixed address is enabled.",
	},
	"disable_discovery": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Determines if the discovery for this fixed address is disabled or not. False means that the discovery is enabled.",
	},
	"enable_ddns": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "The dynamic DNS updates flag of a DHCP Fixed Address object. If set to True, the DHCP server sends DDNS updates to DNS servers in the same Grid, and to external DNS servers.",
	},
	"enable_immediate_discovery": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines if the discovery for the fixed address should be immediately enabled.",
	},
	"enable_pxe_lease_time": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Set this to True if you want the DHCP server to use a different lease time for PXE clients.",
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
	"ignore_dhcp_option_list_request": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "If this field is set to False, the appliance returns all DHCP options the client is eligible to receive, rather than only the list of options the client has requested.",
	},
	"ipv4addr": schema.StringAttribute{
		Optional:   true,
		Computed:   true,
		CustomType: iptypes.IPv4AddressType{},
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
			stringvalidator.ExactlyOneOf(
				path.MatchRelative().AtParent().AtName("ipv4addr"),
				path.MatchRelative().AtParent().AtName("dynamic_allocation"),
			),
		},
		MarkdownDescription: "The IPv4 Address of the fixed address.",
	},
	"logic_filter_rules": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: FixedaddressLogicFilterRulesResourceSchemaAttributes,
		},
		Optional: true,
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
		},
		MarkdownDescription: "This field contains the logic filters to be applied on the this fixed address. This list corresponds to the match rules that are written to the dhcpd configuration file.",
	},
	"mac": schema.StringAttribute{
		Optional:   true,
		CustomType: internaltypes.MACAddressType{},
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "The MAC address value for this fixed address.",
	},
	"match_client": schema.StringAttribute{
		Default: stringdefault.StaticString("MAC_ADDRESS"),
		Validators: []validator.String{
			stringvalidator.OneOf("MAC_ADDRESS", "CLIENT_ID", "RESERVED", "CIRCUIT_ID", "REMOTE_ID"),
		},
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The match_client value for this fixed address. Valid values are: \"MAC_ADDRESS\": The fixed IP address is leased to the matching MAC address. \"CLIENT_ID\": The fixed IP address is leased to the matching DHCP client identifier. \"RESERVED\": The fixed IP address is reserved for later use with a MAC address that only has zeros. \"CIRCUIT_ID\": The fixed IP address is leased to the DHCP client with a matching circuit ID. Note that the \"agent_circuit_id\" field must be set in this case. \"REMOTE_ID\": The fixed IP address is leased to the DHCP client with a matching remote ID. Note that the \"agent_remote_id\" field must be set in this case.",
	},
	"ms_options": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: FixedaddressMsOptionsResourceSchemaAttributes,
		},
		Optional: true,
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
		},
		MarkdownDescription: "This field contains the Microsoft DHCP options for this fixed address.",
	},
	"ms_server": schema.SingleNestedAttribute{
		Attributes:          FixedaddressMsServerResourceSchemaAttributes,
		Optional:            true,
		MarkdownDescription: "",
	},
	"name": schema.StringAttribute{
		Optional: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "This field contains the name of this fixed address.",
	},
	"network": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
			customvalidator.IsValidIPCIDR(),
		},
		MarkdownDescription: "The network to which this fixed address belongs, in IPv4 Address/CIDR format.",
	},
	"network_view": schema.StringAttribute{
		Default:  stringdefault.StaticString("default"),
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "The name of the network view in which this fixed address resides.",
	},
	"nextserver": schema.StringAttribute{
		Optional: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
			customvalidator.IsValidIPv4OrFQDN(),
		},
		MarkdownDescription: "The name in FQDN and/or IPv4 Address format of the next server that the host needs to boot.",
	},
	"options": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: FixedaddressOptionsResourceSchemaAttributes,
		},
		Optional: true,
		Computed: true,
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
		},
		MarkdownDescription: "An array of DHCP option dhcpoption structs that lists the DHCP options associated with the object.",
	},
	"pxe_lease_time": schema.Int64Attribute{
		Optional: true,
		Computed: true,
		Validators: []validator.Int64{
			int64validator.Between(0, 399999999),
		},
		MarkdownDescription: "The PXE lease time value for a DHCP Fixed Address object. Some hosts use PXE (Preboot Execution Environment) to boot remotely from a server. To better manage your IP resources, set a different lease time for PXE boot requests. You can configure the DHCP server to allocate an IP address with a shorter lease time to hosts that send PXE boot requests, so IP addresses are not leased longer than necessary. A 32-bit unsigned integer that represents the duration, in seconds, for which the update is cached. Zero indicates that the update is not cached.",
	},
	"reserved_interface": schema.StringAttribute{
		Optional: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "The ref to the reserved interface to which the device belongs.",
	},
	"restart_if_needed": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Restarts the member service. The restart_if_needed flag can trigger a restart on DHCP services only when it is enabled on CP member.",
	},
	"snmp3_credential": schema.SingleNestedAttribute{
		Attributes:          FixedaddressSnmp3CredentialResourceSchemaAttributes,
		Optional:            true,
		MarkdownDescription: "",
	},
	"snmp_credential": schema.SingleNestedAttribute{
		Attributes:          FixedaddressSnmpCredentialResourceSchemaAttributes,
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
		MarkdownDescription: "If set on creation, the fixed address will be created according to the values specified in the named template.",
	},
	"dynamic_allocation": schema.SingleNestedAttribute{
		Attributes:          dynamicallocation.NextAvailableIpResourceSchemaAttributes,
		Optional:            true,
		MarkdownDescription: "Dynamically allocate the ip using the NIOS next_available_ip function call. Mutually exclusive with the static value field.",
	},
}

var FixedaddressResourceUddiSchemaAttributes = map[string]schema.Attribute{
	"address": schema.StringAttribute{
		Optional: true,
		Computed: true,
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.RequiresReplaceIfConfigured(),
			stringplanmodifier.UseStateForUnknown(),
		},
		Validators: []validator.String{
			stringvalidator.ExactlyOneOf(
				path.MatchRelative().AtParent().AtName("dynamic_allocation"),
			),
		},
		MarkdownDescription: "The reserved address.",
	},
	"comment": schema.StringAttribute{
		Default:             stringdefault.StaticString(""),
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The description for the fixed address. May contain 0 to 1024 characters. Can include UTF-8.",
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
		MarkdownDescription: "Optional. _true_ to disable object. The fixed address is converted to an exclusion when generating configuration.  Defaults to _false_.",
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
	"hostname": schema.StringAttribute{
		Default:             stringdefault.StaticString(""),
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The DHCP host name associated with this fixed address. It is of FQDN type and it defaults to empty.",
	},
	"inheritance_parent": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The resource identifier.",
	},
	"inheritance_sources": schema.SingleNestedAttribute{
		Attributes: FixedAddressInheritanceResourceSchemaAttributes,
		Optional:   true,
		Computed:   true,
		PlanModifiers: []planmodifier.Object{
			objectplanmodifier.UseStateForUnknown(),
		},
		MarkdownDescription: "The __FixedAddressInheritance__ object specifies how and which fields _FixedAddress_ object inherits from the parent.",
	},
	"ip_space": schema.StringAttribute{
		Optional: true,
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.RequiresReplaceIfConfigured(),
		},
		MarkdownDescription: "The resource identifier.",
	},
	"match_type": schema.StringAttribute{
		Required:            true,
		MarkdownDescription: "Indicates how to match the client:  * _mac_: match the client MAC address for both IPv4 and IPv6,  * _client_text_ or _client_hex_: match the client identifier for IPv4 only,  * _relay_text_ or _relay_hex_: match the circuit ID or remote ID in the DHCP relay agent option (82) for IPv4 only,  * _duid_: match the DHCP unique identifier, currently match only for IPv6 protocol.",
	},
	"match_value": schema.StringAttribute{
		Required:            true,
		MarkdownDescription: "The value to match.",
	},
	"name": schema.StringAttribute{
		Default:             stringdefault.StaticString(""),
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The name of the fixed address. May contain 1 to 256 characters. Can include UTF-8.",
	},
	"parent": schema.StringAttribute{
		Computed:            true,
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
		MarkdownDescription: "The tags for the fixed address in JSON format.",
	},
	"tags_all": schema.MapAttribute{
		Computed:            true,
		ElementType:         types.StringType,
		MarkdownDescription: "All tags including inherited values.",
	},
	"dynamic_allocation": schema.SingleNestedAttribute{
		Attributes:          dynamicallocation.NextAvailableAddressResourceSchemaAttributes,
		Optional:            true,
		MarkdownDescription: "Dynamically allocate the next available address from a parent scope. Mutually exclusive with the static \"address\" field.",
	},
}

// Expand converts the TF model to the infoblox core model
func (m *FixedaddressModel) Expand(ctx context.Context, diags *diag.Diagnostics, isCreate bool) *coremodel.Fixedaddress {
	if m == nil {
		return nil
	}

	obj := &coremodel.Fixedaddress{}

	// Expand NIOS nested attribute (returns nil if not present)
	niosModel := flex.ExpandNestedObject[NIOSFixedaddressModel](ctx, m.NIOS, diags)
	if niosModel != nil {
		obj.NIOS = niosModel.Expand(ctx, diags, isCreate)
	}

	// Expand UDDI nested attribute (returns nil if not present)
	uddiModel := flex.ExpandNestedObject[UDDIFixedaddressModel](ctx, m.UDDI, diags)
	if uddiModel != nil {
		obj.UDDI = uddiModel.Expand(ctx, diags, isCreate)
	}

	return obj
}

// Expand converts the NIOS TF model to the core model.
func (m *NIOSFixedaddressModel) Expand(ctx context.Context, diags *diag.Diagnostics, isCreate bool) *coremodel.NIOSFixedaddressExt {
	ext := &coremodel.NIOSFixedaddressExt{
		AgentCircuitId:              flex.ExpandStringPointer(m.AgentCircuitId),
		AgentRemoteId:               flex.ExpandStringPointer(m.AgentRemoteId),
		AllowTelnet:                 flex.ExpandBoolPointer(m.AllowTelnet),
		AlwaysUpdateDns:             flex.ExpandBoolPointer(m.AlwaysUpdateDns),
		Bootfile:                    flex.ExpandStringPointerNullAsEmpty(m.Bootfile),
		Bootserver:                  flex.ExpandStringPointerNullAsEmpty(m.Bootserver),
		CliCredentials:              flex.ExpandFrameworkListNestedBlock(ctx, m.CliCredentials, diags, ExpandFixedaddressCliCredentials),
		ClientIdentifierPrependZero: flex.ExpandBoolPointer(m.ClientIdentifierPrependZero),
		Comment:                     flex.ExpandStringPointerNullAsEmpty(m.Comment),
		DdnsDomainname:              flex.ExpandStringPointerNullAsEmpty(m.DdnsDomainname),
		DdnsHostname:                flex.ExpandStringPointerNullAsEmpty(m.DdnsHostname),
		DenyBootp:                   flex.ExpandBoolPointer(m.DenyBootp),
		DeviceDescription:           flex.ExpandStringPointerNullAsEmpty(m.DeviceDescription),
		DeviceLocation:              flex.ExpandStringPointerNullAsEmpty(m.DeviceLocation),
		DeviceType:                  flex.ExpandStringPointerNullAsEmpty(m.DeviceType),
		DeviceVendor:                flex.ExpandStringPointerNullAsEmpty(m.DeviceVendor),
		DhcpClientIdentifier:        flex.ExpandStringPointer(m.DhcpClientIdentifier),
		Disable:                     flex.ExpandBoolPointer(m.Disable),
		DisableDiscovery:            flex.ExpandBoolPointer(m.DisableDiscovery),
		EnableDdns:                  flex.ExpandBoolPointer(m.EnableDdns),
		EnableImmediateDiscovery:    flex.ExpandBoolPointer(m.EnableImmediateDiscovery),
		EnablePxeLeaseTime:          flex.ExpandBoolPointer(m.EnablePxeLeaseTime),
		ExtAttrs:                    flex.ExpandMapStringAny(ctx, m.ExtAttrs, diags),
		IgnoreDhcpOptionListRequest: flex.ExpandBoolPointer(m.IgnoreDhcpOptionListRequest),
		Ipv4addr:                    flex.ExpandIPv4Address(m.Ipv4addr),
		LogicFilterRules:            flex.ExpandFrameworkListNestedBlock(ctx, m.LogicFilterRules, diags, ExpandFixedaddressLogicFilterRules),
		Mac:                         flex.ExpandMACAddress(m.Mac),
		MatchClient:                 flex.ExpandStringPointerNullAsEmpty(m.MatchClient),
		MsOptions:                   flex.ExpandFrameworkListNestedBlock(ctx, m.MsOptions, diags, ExpandFixedaddressMsOptions),
		MsServer:                    ExpandFixedaddressMsServer(ctx, m.MsServer, diags),
		Name:                        flex.ExpandStringPointerNullAsEmpty(m.Name),
		Network:                     flex.ExpandStringPointer(m.Network),
		NetworkView:                 flex.ExpandStringPointerNullAsEmpty(m.NetworkView),
		Nextserver:                  flex.ExpandStringPointerNullAsEmpty(m.Nextserver),
		Options:                     flex.ExpandFrameworkListNestedBlock(ctx, m.Options, diags, ExpandFixedaddressOptions),
		PxeLeaseTime:                flex.ExpandInt64Pointer(m.PxeLeaseTime),
		ReservedInterface:           flex.ExpandStringPointer(m.ReservedInterface),
		RestartIfNeeded:             flex.ExpandBoolPointer(m.RestartIfNeeded),
		Snmp3Credential:             ExpandFixedaddressSnmp3Credential(ctx, m.Snmp3Credential, diags),
		SnmpCredential:              ExpandFixedaddressSnmpCredential(ctx, m.SnmpCredential, diags),
	}
	if isCreate {
		ext.Template = flex.ExpandStringPointer(m.Template)
		ext.FuncCall = BuildFixedaddressFuncCall(ctx, m.DynamicAllocation, diags)
	}
	return ext
}

// ApplyFixedaddressNIOSUseFlags derives NIOS use flags from the raw config
// value(s) and writes them onto the core model. A flag is true when the user
// set any of its governed value fields in config.
func ApplyFixedaddressNIOSUseFlags(ctx context.Context, config tfsdk.Config, obj *coremodel.Fixedaddress, diags *diag.Diagnostics) {
	if obj == nil || obj.NIOS == nil {
		return
	}
	obj.NIOS.UseBootfile = flex.DeriveUseFlag(ctx, config, diags, path.Root("nios").AtName("bootfile"))
	obj.NIOS.UseBootserver = flex.DeriveUseFlag(ctx, config, diags, path.Root("nios").AtName("bootserver"))
	obj.NIOS.UseCliCredentials = flex.DeriveUseFlag(ctx, config, diags, path.Root("nios").AtName("cli_credentials"), path.Root("nios").AtName("snmp3_credential"))
	obj.NIOS.UseDdnsDomainname = flex.DeriveUseFlag(ctx, config, diags, path.Root("nios").AtName("ddns_domainname"))
	obj.NIOS.UseDenyBootp = flex.DeriveUseFlag(ctx, config, diags, path.Root("nios").AtName("deny_bootp"))
	obj.NIOS.UseEnableDdns = flex.DeriveUseFlag(ctx, config, diags, path.Root("nios").AtName("enable_ddns"))
	obj.NIOS.UseIgnoreDhcpOptionListRequest = flex.DeriveUseFlag(ctx, config, diags, path.Root("nios").AtName("ignore_dhcp_option_list_request"))
	obj.NIOS.UseLogicFilterRules = flex.DeriveUseFlag(ctx, config, diags, path.Root("nios").AtName("logic_filter_rules"))
	obj.NIOS.UseMsOptions = flex.DeriveUseFlag(ctx, config, diags, path.Root("nios").AtName("ms_options"))
	obj.NIOS.UseNextserver = flex.DeriveUseFlag(ctx, config, diags, path.Root("nios").AtName("nextserver"))
	obj.NIOS.UseOptions = flex.DeriveUseFlag(ctx, config, diags, path.Root("nios").AtName("options"))
	obj.NIOS.UsePxeLeaseTime = flex.DeriveUseFlag(ctx, config, diags, path.Root("nios").AtName("pxe_lease_time"))
	obj.NIOS.UseSnmp3Credential = flex.DeriveUseFlag(ctx, config, diags, path.Root("nios").AtName("snmp3_credential"))
	obj.NIOS.UseSnmpCredential = flex.DeriveUseFlag(ctx, config, diags, path.Root("nios").AtName("snmp_credential"))
}

// Expand converts the UDDI TF model to the core model.
func (m *UDDIFixedaddressModel) Expand(ctx context.Context, diags *diag.Diagnostics, isCreate bool) *coremodel.UDDIFixedaddressExt {
	ext := &coremodel.UDDIFixedaddressExt{
		Address:                   flex.ExpandString(m.Address),
		Comment:                   flex.ExpandStringPointer(m.Comment),
		DhcpOptions:               flex.ExpandFrameworkListNestedBlock(ctx, m.DhcpOptions, diags, ExpandOptionItem),
		DisableDhcp:               flex.ExpandBoolPointer(m.DisableDhcp),
		HeaderOptionFilename:      flex.ExpandStringPointer(m.HeaderOptionFilename),
		HeaderOptionServerAddress: flex.ExpandStringPointer(m.HeaderOptionServerAddress),
		HeaderOptionServerName:    flex.ExpandStringPointer(m.HeaderOptionServerName),
		Hostname:                  flex.ExpandStringPointer(m.Hostname),
		InheritanceParent:         flex.ExpandStringPointer(m.InheritanceParent),
		InheritanceSources:        ExpandFixedAddressInheritance(ctx, m.InheritanceSources, diags),
		IpSpace:                   flex.ExpandStringPointer(m.IpSpace),
		MatchType:                 flex.ExpandString(m.MatchType),
		MatchValue:                flex.ExpandString(m.MatchValue),
		Name:                      flex.ExpandStringPointer(m.Name),
		Parent:                    flex.ExpandStringPointer(m.Parent),
		Tags:                      flex.ExpandMapStringAny(ctx, m.Tags, diags),
	}
	if isCreate {
		if alloc := BuildFixedaddressAllocation(ctx, m.DynamicAllocation, diags); alloc != nil {
			ext.Address = *alloc
		}
	}
	return ext
}

// Flatten populates the TF model from a core response.
func (m *FixedaddressModel) Flatten(ctx context.Context, resp *coremodel.Fixedaddress, diags *diag.Diagnostics) {
	if resp == nil {
		return
	}

	m.Id = flex.FlattenStringPointer(resp.Id)

	// Extract existing NIOS model, flatten API response onto it, convert back
	niosModel := flex.ExpandNestedObject[NIOSFixedaddressModel](ctx, m.NIOS, diags)
	if niosModel == nil {
		niosModel = &NIOSFixedaddressModel{}
	}
	plannedNIOS := flex.ExpandNestedObject[NIOSFixedaddressModel](ctx, m.NIOS, diags)
	niosModel.Flatten(ctx, resp.NIOS, diags)
	if resp.NIOS != nil {
		PostFlattenFixedaddressNIOS(ctx, plannedNIOS, niosModel, diags)
		m.NIOS = flex.FlattenNestedObject(ctx, niosModel, NIOSFixedaddressAttrTypes, diags)
	} else {
		m.NIOS = types.ObjectNull(NIOSFixedaddressAttrTypes)
	}

	// Extract existing UDDI model, flatten API response onto it, convert back
	uddiModel := flex.ExpandNestedObject[UDDIFixedaddressModel](ctx, m.UDDI, diags)
	if uddiModel == nil {
		uddiModel = &UDDIFixedaddressModel{}
	}
	uddiModel.Flatten(ctx, resp.UDDI, diags)
	if resp.UDDI != nil {
		m.UDDI = flex.FlattenNestedObject(ctx, uddiModel, UDDIFixedaddressAttrTypes, diags)
	} else {
		m.UDDI = types.ObjectNull(UDDIFixedaddressAttrTypes)
	}
}

// Flatten merges API response onto existing NIOS model.
func (m *NIOSFixedaddressModel) Flatten(ctx context.Context, from *coremodel.NIOSFixedaddressExt, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	planExtAttrs := m.ExtAttrs
	if planExtAttrs.IsUnknown() {
		planExtAttrs = types.MapNull(types.StringType)
	}
	m.AgentCircuitId = flex.FlattenStringPointerEmptyAsNull(from.AgentCircuitId)
	m.AgentRemoteId = flex.FlattenStringPointerEmptyAsNull(from.AgentRemoteId)
	m.AllowTelnet = flex.FlattenBoolPointer(from.AllowTelnet)
	m.AlwaysUpdateDns = flex.FlattenBoolPointer(from.AlwaysUpdateDns)
	m.Bootfile = flex.FlattenStringPointerEmptyAsNull(from.Bootfile)
	m.Bootserver = flex.FlattenStringPointerEmptyAsNull(from.Bootserver)
	m.CliCredentials = flex.FlattenFrameworkListNestedBlock(ctx, from.CliCredentials, FixedaddressCliCredentialsAttrTypes, diags, FlattenFixedaddressCliCredentials)
	m.ClientIdentifierPrependZero = flex.FlattenBoolPointer(from.ClientIdentifierPrependZero)
	m.Comment = flex.FlattenStringPointerEmptyAsNull(from.Comment)
	m.DdnsDomainname = flex.FlattenStringPointerEmptyAsNull(from.DdnsDomainname)
	m.DdnsHostname = flex.FlattenStringPointerEmptyAsNull(from.DdnsHostname)
	m.DenyBootp = flex.FlattenBoolPointer(from.DenyBootp)
	m.DeviceDescription = flex.FlattenStringPointerEmptyAsNull(from.DeviceDescription)
	m.DeviceLocation = flex.FlattenStringPointerEmptyAsNull(from.DeviceLocation)
	m.DeviceType = flex.FlattenStringPointerEmptyAsNull(from.DeviceType)
	m.DeviceVendor = flex.FlattenStringPointerEmptyAsNull(from.DeviceVendor)
	m.DhcpClientIdentifier = flex.FlattenStringPointerEmptyAsNull(from.DhcpClientIdentifier)
	m.Disable = flex.FlattenBoolPointer(from.Disable)
	m.DisableDiscovery = flex.FlattenBoolPointer(from.DisableDiscovery)
	m.EnableDdns = flex.FlattenBoolPointer(from.EnableDdns)
	m.EnableImmediateDiscovery = flex.FlattenBoolPointer(from.EnableImmediateDiscovery)
	m.EnablePxeLeaseTime = flex.FlattenBoolPointer(from.EnablePxeLeaseTime)
	m.ExtAttrs, m.ExtAttrsAll = flex.FlattenEAs(planExtAttrs, from.ExtAttrs)
	m.IgnoreDhcpOptionListRequest = flex.FlattenBoolPointer(from.IgnoreDhcpOptionListRequest)
	m.Ipv4addr = flex.FlattenIPv4Address(from.Ipv4addr)
	m.LogicFilterRules = flex.FlattenFrameworkListNestedBlock(ctx, from.LogicFilterRules, FixedaddressLogicFilterRulesAttrTypes, diags, FlattenFixedaddressLogicFilterRules)
	m.Mac = flex.FlattenMACAddress(from.Mac)
	m.MatchClient = flex.FlattenStringPointerEmptyAsNull(from.MatchClient)
	m.MsOptions = flex.FlattenFrameworkListNestedBlock(ctx, from.MsOptions, FixedaddressMsOptionsAttrTypes, diags, FlattenFixedaddressMsOptions)
	m.MsServer = FlattenFixedaddressMsServer(ctx, from.MsServer, diags)
	m.Name = flex.FlattenStringPointerEmptyAsNull(from.Name)
	m.Network = flex.FlattenStringPointerEmptyAsNull(from.Network)
	m.NetworkView = flex.FlattenStringPointerEmptyAsNull(from.NetworkView)
	m.Nextserver = flex.FlattenStringPointerEmptyAsNull(from.Nextserver)
	m.Options = flex.FlattenFrameworkListNestedBlock(ctx, from.Options, FixedaddressOptionsAttrTypes, diags, FlattenFixedaddressOptions)
	m.PxeLeaseTime = flex.FlattenInt64Pointer(from.PxeLeaseTime)
	m.ReservedInterface = flex.FlattenStringPointerEmptyAsNull(from.ReservedInterface)
	m.Snmp3Credential = FlattenFixedaddressSnmp3Credential(ctx, from.Snmp3Credential, diags)
	m.SnmpCredential = FlattenFixedaddressSnmpCredential(ctx, from.SnmpCredential, diags)
	if len(m.DynamicAllocation.AttributeTypes(ctx)) == 0 {
		m.DynamicAllocation = types.ObjectNull(dynamicallocation.NextAvailableIpAttrTypes)
	}
}

// Flatten merges API response onto existing UDDI model.
func (m *UDDIFixedaddressModel) Flatten(ctx context.Context, from *coremodel.UDDIFixedaddressExt, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.Address = flex.FlattenString(from.Address)
	m.Comment = flex.FlattenStringPointer(from.Comment)
	m.DhcpOptions = flex.FlattenFrameworkListNestedBlock(ctx, from.DhcpOptions, OptionItemAttrTypes, diags, FlattenOptionItem)
	m.DisableDhcp = flex.FlattenBoolPointer(from.DisableDhcp)
	m.HeaderOptionFilename = flex.FlattenStringPointer(from.HeaderOptionFilename)
	m.HeaderOptionServerAddress = flex.FlattenStringPointer(from.HeaderOptionServerAddress)
	m.HeaderOptionServerName = flex.FlattenStringPointer(from.HeaderOptionServerName)
	m.Hostname = flex.FlattenStringPointer(from.Hostname)
	m.InheritanceParent = flex.FlattenStringPointer(from.InheritanceParent)
	m.InheritanceSources = FlattenFixedAddressInheritance(ctx, from.InheritanceSources, diags)
	m.IpSpace = flex.FlattenStringPointer(from.IpSpace)
	m.MatchType = flex.FlattenString(from.MatchType)
	m.MatchValue = flex.FlattenString(from.MatchValue)
	m.Name = flex.FlattenStringPointer(from.Name)
	m.Parent = flex.FlattenStringPointer(from.Parent)
	tagsAll := flex.FlattenMapStringAny(ctx, from.Tags, diags)
	if m.Tags.IsNull() || m.Tags.IsUnknown() {
		m.Tags = tagsAll
	}
	m.TagsAll = tagsAll
	if len(m.DynamicAllocation.AttributeTypes(ctx)) == 0 {
		m.DynamicAllocation = types.ObjectNull(dynamicallocation.NextAvailableAddressAttrTypes)
	}
}
