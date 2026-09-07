package dhcp

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-nettypes/cidrtypes"
	"github.com/hashicorp/terraform-plugin-framework-nettypes/iptypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/mapvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listdefault"
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

type Ipv6fixedaddressModel struct {
	Id   types.String `tfsdk:"id"`
	NIOS types.Object `tfsdk:"nios"`
	UDDI types.Object `tfsdk:"uddi"`
}

var Ipv6fixedaddressAttrTypes = map[string]attr.Type{
	"id":   types.StringType,
	"nios": types.ObjectType{AttrTypes: NIOSIpv6fixedaddressAttrTypes},
	"uddi": types.ObjectType{AttrTypes: UDDIIpv6fixedaddressAttrTypes},
}

type NIOSIpv6fixedaddressModel struct {
	AddressType              types.String                        `tfsdk:"address_type"`
	AllowTelnet              types.Bool                          `tfsdk:"allow_telnet"`
	CliCredentials           types.List                          `tfsdk:"cli_credentials"`
	Comment                  types.String                        `tfsdk:"comment"`
	DeviceDescription        types.String                        `tfsdk:"device_description"`
	DeviceLocation           types.String                        `tfsdk:"device_location"`
	DeviceType               types.String                        `tfsdk:"device_type"`
	DeviceVendor             types.String                        `tfsdk:"device_vendor"`
	Disable                  types.Bool                          `tfsdk:"disable"`
	DisableDiscovery         types.Bool                          `tfsdk:"disable_discovery"`
	DomainName               internaltypes.CaseInsensitiveString `tfsdk:"domain_name"`
	DomainNameServers        types.List                          `tfsdk:"domain_name_servers"`
	Duid                     types.String                        `tfsdk:"duid"`
	EnableImmediateDiscovery types.Bool                          `tfsdk:"enable_immediate_discovery"`
	ExtAttrs                 types.Map                           `tfsdk:"ext_attrs"`
	ExtAttrsAll              types.Map                           `tfsdk:"ext_attrs_all"`
	Ipv6addr                 iptypes.IPv6Address                 `tfsdk:"ipv6addr"`
	Ipv6prefix               types.String                        `tfsdk:"ipv6prefix"`
	Ipv6prefixBits           types.Int64                         `tfsdk:"ipv6prefix_bits"`
	LogicFilterRules         types.List                          `tfsdk:"logic_filter_rules"`
	MacAddress               internaltypes.MACAddress            `tfsdk:"mac_address"`
	MatchClient              types.String                        `tfsdk:"match_client"`
	Name                     types.String                        `tfsdk:"name"`
	Network                  cidrtypes.IPv6Prefix                `tfsdk:"network"`
	NetworkView              types.String                        `tfsdk:"network_view"`
	Options                  types.List                          `tfsdk:"options"`
	PreferredLifetime        types.Int64                         `tfsdk:"preferred_lifetime"`
	ReservedInterface        types.String                        `tfsdk:"reserved_interface"`
	RestartIfNeeded          types.Bool                          `tfsdk:"restart_if_needed"`
	Snmp3Credential          types.Object                        `tfsdk:"snmp3_credential"`
	SnmpCredential           types.Object                        `tfsdk:"snmp_credential"`
	Template                 types.String                        `tfsdk:"template"`
	ValidLifetime            types.Int64                         `tfsdk:"valid_lifetime"`
	DynamicAllocation        types.Object                        `tfsdk:"dynamic_allocation"`
}

var NIOSIpv6fixedaddressAttrTypes = map[string]attr.Type{
	"address_type":               types.StringType,
	"allow_telnet":               types.BoolType,
	"cli_credentials":            types.ListType{ElemType: types.ObjectType{AttrTypes: Ipv6fixedaddressCliCredentialsAttrTypes}},
	"comment":                    types.StringType,
	"device_description":         types.StringType,
	"device_location":            types.StringType,
	"device_type":                types.StringType,
	"device_vendor":              types.StringType,
	"disable":                    types.BoolType,
	"disable_discovery":          types.BoolType,
	"domain_name":                internaltypes.CaseInsensitiveStringType{},
	"domain_name_servers":        types.ListType{ElemType: types.StringType},
	"duid":                       types.StringType,
	"enable_immediate_discovery": types.BoolType,
	"ext_attrs":                  types.MapType{ElemType: types.StringType},
	"ext_attrs_all":              types.MapType{ElemType: types.StringType},
	"ipv6addr":                   iptypes.IPv6AddressType{},
	"ipv6prefix":                 types.StringType,
	"ipv6prefix_bits":            types.Int64Type,
	"logic_filter_rules":         types.ListType{ElemType: types.ObjectType{AttrTypes: Ipv6fixedaddressLogicFilterRulesAttrTypes}},
	"mac_address":                internaltypes.MACAddressType{},
	"match_client":               types.StringType,
	"name":                       types.StringType,
	"network":                    cidrtypes.IPv6PrefixType{},
	"network_view":               types.StringType,
	"options":                    types.ListType{ElemType: types.ObjectType{AttrTypes: Ipv6fixedaddressOptionsAttrTypes}},
	"preferred_lifetime":         types.Int64Type,
	"reserved_interface":         types.StringType,
	"restart_if_needed":          types.BoolType,
	"snmp3_credential":           types.ObjectType{AttrTypes: Ipv6fixedaddressSnmp3CredentialAttrTypes},
	"snmp_credential":            types.ObjectType{AttrTypes: Ipv6fixedaddressSnmpCredentialAttrTypes},
	"template":                   types.StringType,
	"valid_lifetime":             types.Int64Type,
	"dynamic_allocation":         types.ObjectType{AttrTypes: dynamicallocation.NextAvailableIpAttrTypes},
}

type UDDIIpv6fixedaddressModel struct {
	Address            iptypes.IPv6Address `tfsdk:"address"`
	Comment            types.String        `tfsdk:"comment"`
	DhcpOptions        types.List          `tfsdk:"dhcp_options"`
	DisableDhcp        types.Bool          `tfsdk:"disable_dhcp"`
	Hostname           types.String        `tfsdk:"hostname"`
	InheritanceParent  types.String        `tfsdk:"inheritance_parent"`
	InheritanceSources types.Object        `tfsdk:"inheritance_sources"`
	IpSpace            types.String        `tfsdk:"ip_space"`
	MatchType          types.String        `tfsdk:"match_type"`
	MatchValue         types.String        `tfsdk:"match_value"`
	Name               types.String        `tfsdk:"name"`
	Parent             types.String        `tfsdk:"parent"`
	Tags               types.Map           `tfsdk:"tags"`
	TagsAll            types.Map           `tfsdk:"tags_all"`
}

var UDDIIpv6fixedaddressAttrTypes = map[string]attr.Type{
	"address":             iptypes.IPv6AddressType{},
	"comment":             types.StringType,
	"dhcp_options":        types.ListType{ElemType: types.ObjectType{AttrTypes: OptionItemAttrTypes}},
	"disable_dhcp":        types.BoolType,
	"hostname":            types.StringType,
	"inheritance_parent":  types.StringType,
	"inheritance_sources": types.ObjectType{AttrTypes: FixedAddressInheritanceAttrTypes},
	"ip_space":            types.StringType,
	"match_type":          types.StringType,
	"match_value":         types.StringType,
	"name":                types.StringType,
	"parent":              types.StringType,
	"tags":                types.MapType{ElemType: types.StringType},
	"tags_all":            types.MapType{ElemType: types.StringType},
}

const (
	Ipv6fixedaddressInheritanceType = "full"
	Ipv6fixedaddressReturnFields    = "address_type,allow_telnet,cli_credentials,cloud_info,comment,device_description,device_location,device_type,device_vendor,disable,disable_discovery,discover_now_status,discovered_data,domain_name,domain_name_servers,duid,extattrs,ipv6addr,ipv6prefix,ipv6prefix_bits,logic_filter_rules,mac_address,match_client,ms_ad_user_data,name,network,network_view,options,preferred_lifetime,reserved_interface,snmp3_credential,snmp_credential,use_cli_credentials,use_domain_name,use_domain_name_servers,use_logic_filter_rules,use_options,use_preferred_lifetime,use_snmp3_credential,use_snmp_credential,use_valid_lifetime,valid_lifetime"
)

var Ipv6fixedaddressResourceSchemaAttributes = map[string]schema.Attribute{
	"id": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The reference to the object.",
	},
	"nios": schema.SingleNestedAttribute{
		Optional:            true,
		MarkdownDescription: "NIOS backend-specific fields.",
		Attributes:          Ipv6fixedaddressResourceNiosSchemaAttributes,
	},
	"uddi": schema.SingleNestedAttribute{
		Optional:            true,
		MarkdownDescription: "UDDI backend-specific fields.",
		Attributes:          Ipv6fixedaddressResourceUddiSchemaAttributes,
	},
}

var Ipv6fixedaddressResourceNiosSchemaAttributes = map[string]schema.Attribute{
	"address_type": schema.StringAttribute{
		Default: stringdefault.StaticString("ADDRESS"),
		Validators: []validator.String{
			stringvalidator.OneOf("ADDRESS", "PREFIX", "BOTH"),
		},
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The address type value for this IPv6 fixed address. When the address type is \"ADDRESS\", a value for the 'ipv6addr' member is required. When the address type is \"PREFIX\", values for 'ipv6prefix' and 'ipv6prefix_bits' are required. When the address type is \"BOTH\", values for 'ipv6addr', 'ipv6prefix', and 'ipv6prefix_bits' are all required.",
	},
	"allow_telnet": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "This field controls whether the credential is used for both the Telnet and SSH credentials. If set to False, the credential is used only for SSH.",
	},
	"cli_credentials": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: Ipv6fixedaddressCliCredentialsResourceSchemaAttributes,
		},
		Optional: true,
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
		},
		MarkdownDescription: "The CLI credentials for the IPv6 fixed address.",
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
	"disable": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Determines whether a fixed address is disabled or not. When this is set to False, the IPv6 fixed address is enabled.",
	},
	"disable_discovery": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Determines if the discovery for this IPv6 fixed address is disabled or not. False means that the discovery is enabled.",
	},
	"domain_name": schema.StringAttribute{
		Optional:   true,
		CustomType: internaltypes.CaseInsensitiveStringType{},
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "The domain name for this IPv6 fixed address.",
	},
	"domain_name_servers": schema.ListAttribute{
		ElementType: types.StringType,
		Optional:    true,
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
			listvalidator.ValueStringsAre(customvalidator.IsValidIPv6Address()),
		},
		MarkdownDescription: "The IPv6 addresses of DNS recursive name servers to which the DHCP client can send name resolution requests. The DHCP server includes this information in the DNS Recursive Name Server option in Advertise, Rebind, Information-Request, and Reply messages.",
	},
	"duid": schema.StringAttribute{
		Optional: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
			customvalidator.IsValidDUID(),
		},
		MarkdownDescription: "The DUID value for this IPv6 fixed address.",
	},
	"enable_immediate_discovery": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines if the discovery for the IPv6 fixed address should be immediately enabled.",
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
	"ipv6addr": schema.StringAttribute{
		Optional:   true,
		Computed:   true,
		CustomType: iptypes.IPv6AddressType{},
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "The IPv6 Address of the DHCP IPv6 fixed address.",
	},
	"ipv6prefix": schema.StringAttribute{
		Optional: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "The IPv6 Address prefix of the DHCP IPv6 fixed address.",
	},
	"ipv6prefix_bits": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "Prefix bits of the DHCP IPv6 fixed address.",
	},
	"logic_filter_rules": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: Ipv6fixedaddressLogicFilterRulesResourceSchemaAttributes,
		},
		Optional: true,
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
		},
		MarkdownDescription: "This field contains the logic filters to be applied to this IPv6 fixed address. This list corresponds to the match rules that are written to the DHCPv6 configuration file.",
	},
	"mac_address": schema.StringAttribute{
		Optional:   true,
		CustomType: internaltypes.MACAddressType{},
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "The MAC address for this IPv6 fixed address.",
	},
	"match_client": schema.StringAttribute{
		Default: stringdefault.StaticString("DUID"),
		Validators: []validator.String{
			stringvalidator.OneOf("DUID", "MAC_ADDRESS"),
		},
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The match_client value for this fixed address. Valid values are: \"DUID\": The fixed IP address is leased to the matching DUID. \"MAC_ADDRESS\": The fixed IP address is leased to the matching MAC address.",
	},
	"name": schema.StringAttribute{
		Optional: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "This field contains the name of this IPv6 fixed address.",
	},
	"network": schema.StringAttribute{
		Optional:   true,
		Computed:   true,
		CustomType: cidrtypes.IPv6PrefixType{},
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
			customvalidator.IsValidIPCIDR(),
		},
		MarkdownDescription: "The network to which this IPv6 fixed address belongs, in IPv6 Address/CIDR format.",
	},
	"network_view": schema.StringAttribute{
		Default:  stringdefault.StaticString("default"),
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "The name of the network view in which this IPv6 fixed address resides.",
	},
	"options": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: Ipv6fixedaddressOptionsResourceSchemaAttributes,
		},
		Optional: true,
		Computed: true,
		Default:  listdefault.StaticValue(types.ListValueMust(types.ObjectType{AttrTypes: Ipv6fixedaddressOptionsAttrTypes}, []attr.Value{})),
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
		},
		MarkdownDescription: "An array of DHCP option dhcpoption structs that lists the DHCP options associated with the object.",
	},
	"preferred_lifetime": schema.Int64Attribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The preferred lifetime value for this DHCP IPv6 fixed address object.",
	},
	"reserved_interface": schema.StringAttribute{
		Optional: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "The reference to the reserved interface to which the device belongs.",
	},
	"restart_if_needed": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Restarts the member service. The restart_if_needed flag can trigger a restart on DHCP services only when it is enabled on CP member.",
	},
	"snmp3_credential": schema.SingleNestedAttribute{
		Attributes:          Ipv6fixedaddressSnmp3CredentialResourceSchemaAttributes,
		Optional:            true,
		MarkdownDescription: "",
	},
	"snmp_credential": schema.SingleNestedAttribute{
		Attributes:          Ipv6fixedaddressSnmpCredentialResourceSchemaAttributes,
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
		MarkdownDescription: "If set on creation, the IPv6 fixed address will be created according to the values specified in the named template.",
	},
	"valid_lifetime": schema.Int64Attribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The valid lifetime value for this DHCP IPv6 Fixed Address object.",
	},
	"dynamic_allocation": schema.SingleNestedAttribute{
		Attributes:          dynamicallocation.NextAvailableIpResourceSchemaAttributes,
		Optional:            true,
		MarkdownDescription: "Dynamically allocate the ip using the NIOS next_available_ip function call. Mutually exclusive with the static value field.",
	},
}

var Ipv6fixedaddressResourceUddiSchemaAttributes = map[string]schema.Attribute{
	"address": schema.StringAttribute{
		Required:   true,
		CustomType: iptypes.IPv6AddressType{},
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.RequiresReplaceIfConfigured(),
			stringplanmodifier.UseStateForUnknown(),
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
		Validators: []validator.String{
			stringvalidator.OneOf("mac", "duid"),
		},
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
		Optional:            true,
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
}

// Expand converts the TF model to the infoblox core model
func (m *Ipv6fixedaddressModel) Expand(ctx context.Context, diags *diag.Diagnostics, isCreate bool) *coremodel.Ipv6fixedaddress {
	if m == nil {
		return nil
	}

	obj := &coremodel.Ipv6fixedaddress{}

	// Expand NIOS nested attribute (returns nil if not present)
	niosModel := flex.ExpandNestedObject[NIOSIpv6fixedaddressModel](ctx, m.NIOS, diags)
	if niosModel != nil {
		obj.NIOS = niosModel.Expand(ctx, diags, isCreate)
	}

	// Expand UDDI nested attribute (returns nil if not present)
	uddiModel := flex.ExpandNestedObject[UDDIIpv6fixedaddressModel](ctx, m.UDDI, diags)
	if uddiModel != nil {
		obj.UDDI = uddiModel.Expand(ctx, diags)
	}

	return obj
}

// Expand converts the NIOS TF model to the core model.
func (m *NIOSIpv6fixedaddressModel) Expand(ctx context.Context, diags *diag.Diagnostics, isCreate bool) *coremodel.NIOSIpv6fixedaddressExt {
	ext := &coremodel.NIOSIpv6fixedaddressExt{
		AddressType:              flex.ExpandStringPointerNullAsEmpty(m.AddressType),
		AllowTelnet:              flex.ExpandBoolPointer(m.AllowTelnet),
		CliCredentials:           flex.ExpandFrameworkListNestedBlock(ctx, m.CliCredentials, diags, ExpandIpv6fixedaddressCliCredentials),
		Comment:                  flex.ExpandStringPointerNullAsEmpty(m.Comment),
		DeviceDescription:        flex.ExpandStringPointerNullAsEmpty(m.DeviceDescription),
		DeviceLocation:           flex.ExpandStringPointerNullAsEmpty(m.DeviceLocation),
		DeviceType:               flex.ExpandStringPointerNullAsEmpty(m.DeviceType),
		DeviceVendor:             flex.ExpandStringPointerNullAsEmpty(m.DeviceVendor),
		Disable:                  flex.ExpandBoolPointer(m.Disable),
		DisableDiscovery:         flex.ExpandBoolPointer(m.DisableDiscovery),
		DomainName:               flex.ExpandStringPointer(m.DomainName.StringValue),
		DomainNameServers:        flex.ExpandFrameworkListString(ctx, m.DomainNameServers, diags),
		Duid:                     flex.ExpandStringPointer(m.Duid),
		EnableImmediateDiscovery: flex.ExpandBoolPointer(m.EnableImmediateDiscovery),
		ExtAttrs:                 flex.ExpandMapStringAny(ctx, m.ExtAttrs, diags),
		Ipv6addr:                 flex.ExpandIPv6Address(m.Ipv6addr),
		Ipv6prefix:               flex.ExpandStringPointerNullAsEmpty(m.Ipv6prefix),
		Ipv6prefixBits:           flex.ExpandInt64Pointer(m.Ipv6prefixBits),
		LogicFilterRules:         flex.ExpandFrameworkListNestedBlock(ctx, m.LogicFilterRules, diags, ExpandIpv6fixedaddressLogicFilterRules),
		MacAddress:               flex.ExpandMACAddress(m.MacAddress),
		MatchClient:              flex.ExpandStringPointerNullAsEmpty(m.MatchClient),
		Name:                     flex.ExpandStringPointerNullAsEmpty(m.Name),
		Network:                  flex.ExpandIPv6Prefix(m.Network),
		NetworkView:              flex.ExpandStringPointerNullAsEmpty(m.NetworkView),
		Options:                  flex.ExpandFrameworkListNestedBlock(ctx, m.Options, diags, ExpandIpv6fixedaddressOptions),
		PreferredLifetime:        flex.ExpandInt64Pointer(m.PreferredLifetime),
		ReservedInterface:        flex.ExpandStringPointer(m.ReservedInterface),
		RestartIfNeeded:          flex.ExpandBoolPointer(m.RestartIfNeeded),
		Snmp3Credential:          ExpandIpv6fixedaddressSnmp3Credential(ctx, m.Snmp3Credential, diags),
		SnmpCredential:           ExpandIpv6fixedaddressSnmpCredential(ctx, m.SnmpCredential, diags),
		Template:                 flex.ExpandStringPointer(m.Template),
		ValidLifetime:            flex.ExpandInt64Pointer(m.ValidLifetime),
	}
	if isCreate {
		ext.FuncCall = BuildIpv6fixedaddressFuncCall(ctx, m.DynamicAllocation, diags)
	}
	return ext
}

// ApplyIpv6fixedaddressNIOSUseFlags derives NIOS use flags from the raw config
// value(s) and writes them onto the core model. A flag is true when the user
// set any of its governed value fields in config.
func ApplyIpv6fixedaddressNIOSUseFlags(ctx context.Context, config tfsdk.Config, obj *coremodel.Ipv6fixedaddress, diags *diag.Diagnostics) {
	if obj == nil || obj.NIOS == nil {
		return
	}
	obj.NIOS.UseCliCredentials = flex.DeriveUseFlag(ctx, config, diags, path.Root("nios").AtName("cli_credentials"), path.Root("nios").AtName("snmp3_credential"))
	obj.NIOS.UseDomainName = flex.DeriveUseFlag(ctx, config, diags, path.Root("nios").AtName("domain_name"))
	obj.NIOS.UseDomainNameServers = flex.DeriveUseFlag(ctx, config, diags, path.Root("nios").AtName("domain_name_servers"))
	obj.NIOS.UseLogicFilterRules = flex.DeriveUseFlag(ctx, config, diags, path.Root("nios").AtName("logic_filter_rules"))
	obj.NIOS.UseOptions = flex.DeriveUseFlag(ctx, config, diags, path.Root("nios").AtName("options"))
	obj.NIOS.UsePreferredLifetime = flex.DeriveUseFlag(ctx, config, diags, path.Root("nios").AtName("preferred_lifetime"))
	obj.NIOS.UseSnmp3Credential = flex.DeriveUseFlag(ctx, config, diags, path.Root("nios").AtName("snmp3_credential"))
	obj.NIOS.UseSnmpCredential = flex.DeriveUseFlag(ctx, config, diags, path.Root("nios").AtName("snmp_credential"))
	obj.NIOS.UseValidLifetime = flex.DeriveUseFlag(ctx, config, diags, path.Root("nios").AtName("valid_lifetime"))
}

// Expand converts the UDDI TF model to the core model.
func (m *UDDIIpv6fixedaddressModel) Expand(ctx context.Context, diags *diag.Diagnostics) *coremodel.UDDIIpv6fixedaddressExt {
	return &coremodel.UDDIIpv6fixedaddressExt{
		Address:            flex.ExpandIPv6AddressValue(m.Address),
		Comment:            flex.ExpandStringPointer(m.Comment),
		DhcpOptions:        flex.ExpandFrameworkListNestedBlock(ctx, m.DhcpOptions, diags, ExpandOptionItem),
		DisableDhcp:        flex.ExpandBoolPointer(m.DisableDhcp),
		Hostname:           flex.ExpandStringPointer(m.Hostname),
		InheritanceParent:  flex.ExpandStringPointer(m.InheritanceParent),
		InheritanceSources: ExpandFixedAddressInheritance(ctx, m.InheritanceSources, diags),
		IpSpace:            flex.ExpandStringPointer(m.IpSpace),
		MatchType:          flex.ExpandString(m.MatchType),
		MatchValue:         flex.ExpandString(m.MatchValue),
		Name:               flex.ExpandStringPointer(m.Name),
		Parent:             flex.ExpandStringPointer(m.Parent),
		Tags:               flex.ExpandMapStringAny(ctx, m.Tags, diags),
	}
}

// Flatten populates the TF model from a core response.
func (m *Ipv6fixedaddressModel) Flatten(ctx context.Context, resp *coremodel.Ipv6fixedaddress, diags *diag.Diagnostics) {
	if resp == nil {
		return
	}

	m.Id = flex.FlattenStringPointer(resp.Id)

	// Extract existing NIOS model, flatten API response onto it, convert back
	niosModel := flex.ExpandNestedObject[NIOSIpv6fixedaddressModel](ctx, m.NIOS, diags)
	if niosModel == nil {
		niosModel = &NIOSIpv6fixedaddressModel{}
	}
	plannedNIOS := flex.ExpandNestedObject[NIOSIpv6fixedaddressModel](ctx, m.NIOS, diags)
	niosModel.Flatten(ctx, resp.NIOS, diags)
	if resp.NIOS != nil {
		PostFlattenIpv6fixedaddressNIOS(ctx, plannedNIOS, niosModel, diags)
		m.NIOS = flex.FlattenNestedObject(ctx, niosModel, NIOSIpv6fixedaddressAttrTypes, diags)
	} else {
		m.NIOS = types.ObjectNull(NIOSIpv6fixedaddressAttrTypes)
	}

	// Extract existing UDDI model, flatten API response onto it, convert back
	uddiModel := flex.ExpandNestedObject[UDDIIpv6fixedaddressModel](ctx, m.UDDI, diags)
	if uddiModel == nil {
		uddiModel = &UDDIIpv6fixedaddressModel{}
	}
	uddiModel.Flatten(ctx, resp.UDDI, diags)
	if resp.UDDI != nil {
		m.UDDI = flex.FlattenNestedObject(ctx, uddiModel, UDDIIpv6fixedaddressAttrTypes, diags)
	} else {
		m.UDDI = types.ObjectNull(UDDIIpv6fixedaddressAttrTypes)
	}
}

// Flatten merges API response onto existing NIOS model.
func (m *NIOSIpv6fixedaddressModel) Flatten(ctx context.Context, from *coremodel.NIOSIpv6fixedaddressExt, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	planExtAttrs := m.ExtAttrs
	if planExtAttrs.IsUnknown() {
		planExtAttrs = types.MapNull(types.StringType)
	}
	m.AddressType = flex.FlattenStringPointerEmptyAsNull(from.AddressType)
	m.AllowTelnet = flex.FlattenBoolPointer(from.AllowTelnet)
	m.CliCredentials = flex.FlattenFrameworkListNestedBlock(ctx, from.CliCredentials, Ipv6fixedaddressCliCredentialsAttrTypes, diags, FlattenIpv6fixedaddressCliCredentials)
	m.Comment = flex.FlattenStringPointerEmptyAsNull(from.Comment)
	m.DeviceDescription = flex.FlattenStringPointerEmptyAsNull(from.DeviceDescription)
	m.DeviceLocation = flex.FlattenStringPointerEmptyAsNull(from.DeviceLocation)
	m.DeviceType = flex.FlattenStringPointerEmptyAsNull(from.DeviceType)
	m.DeviceVendor = flex.FlattenStringPointerEmptyAsNull(from.DeviceVendor)
	m.Disable = flex.FlattenBoolPointer(from.Disable)
	m.DisableDiscovery = flex.FlattenBoolPointer(from.DisableDiscovery)
	m.DomainName.StringValue = flex.FlattenStringPointer(from.DomainName)
	m.DomainNameServers = flex.FlattenFrameworkListString(ctx, from.DomainNameServers, diags)
	m.Duid = flex.FlattenStringPointerEmptyAsNull(from.Duid)
	m.EnableImmediateDiscovery = flex.FlattenBoolPointer(from.EnableImmediateDiscovery)
	m.ExtAttrs, m.ExtAttrsAll = flex.FlattenEAs(planExtAttrs, from.ExtAttrs)
	m.Ipv6addr = flex.FlattenIPv6Address(from.Ipv6addr)
	m.Ipv6prefix = flex.FlattenStringPointerEmptyAsNull(from.Ipv6prefix)
	m.Ipv6prefixBits = flex.FlattenInt64Pointer(from.Ipv6prefixBits)
	m.LogicFilterRules = flex.FlattenFrameworkListNestedBlock(ctx, from.LogicFilterRules, Ipv6fixedaddressLogicFilterRulesAttrTypes, diags, FlattenIpv6fixedaddressLogicFilterRules)
	m.MacAddress = flex.FlattenMACAddress(from.MacAddress)
	m.MatchClient = flex.FlattenStringPointerEmptyAsNull(from.MatchClient)
	m.Name = flex.FlattenStringPointerEmptyAsNull(from.Name)
	m.Network = flex.FlattenIPv6Prefix(from.Network)
	m.NetworkView = flex.FlattenStringPointerEmptyAsNull(from.NetworkView)
	m.Options = flex.FlattenFrameworkListNestedBlock(ctx, from.Options, Ipv6fixedaddressOptionsAttrTypes, diags, FlattenIpv6fixedaddressOptions)
	m.PreferredLifetime = flex.FlattenInt64Pointer(from.PreferredLifetime)
	m.ReservedInterface = flex.FlattenStringPointerEmptyAsNull(from.ReservedInterface)
	m.Snmp3Credential = FlattenIpv6fixedaddressSnmp3Credential(ctx, from.Snmp3Credential, diags)
	m.SnmpCredential = FlattenIpv6fixedaddressSnmpCredential(ctx, from.SnmpCredential, diags)
	m.Template = flex.FlattenStringPointerEmptyAsNull(from.Template)
	m.ValidLifetime = flex.FlattenInt64Pointer(from.ValidLifetime)
	if len(m.DynamicAllocation.AttributeTypes(ctx)) == 0 {
		m.DynamicAllocation = types.ObjectNull(dynamicallocation.NextAvailableIpAttrTypes)
	}
}

// Flatten merges API response onto existing UDDI model.
func (m *UDDIIpv6fixedaddressModel) Flatten(ctx context.Context, from *coremodel.UDDIIpv6fixedaddressExt, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.Address = flex.FlattenIPv6AddressValue(from.Address)
	m.Comment = flex.FlattenStringPointer(from.Comment)
	m.DhcpOptions = flex.FlattenFrameworkListNestedBlock(ctx, from.DhcpOptions, OptionItemAttrTypes, diags, FlattenOptionItem)
	m.DisableDhcp = flex.FlattenBoolPointer(from.DisableDhcp)
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
}
