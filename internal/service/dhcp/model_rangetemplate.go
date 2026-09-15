package dhcp

import (
	"context"

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
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"

	coremodel "github.com/infobloxopen/terraform-provider-infoblox/internal/core/model/dhcp"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
	importmod "github.com/infobloxopen/terraform-provider-infoblox/internal/planmodifiers/import"
	internaltypes "github.com/infobloxopen/terraform-provider-infoblox/internal/types"
	customvalidator "github.com/infobloxopen/terraform-provider-infoblox/internal/validator"
)

type RangetemplateModel struct {
	Id   types.String `tfsdk:"id"`
	NIOS types.Object `tfsdk:"nios"`
}

var RangetemplateAttrTypes = map[string]attr.Type{
	"id":   types.StringType,
	"nios": types.ObjectType{AttrTypes: NIOSRangetemplateAttrTypes},
}

type NIOSRangetemplateModel struct {
	Bootfile                    types.String                     `tfsdk:"bootfile"`
	Bootserver                  types.String                     `tfsdk:"bootserver"`
	CloudApiCompatible          types.Bool                       `tfsdk:"cloud_api_compatible"`
	Comment                     types.String                     `tfsdk:"comment"`
	DdnsDomainname              types.String                     `tfsdk:"ddns_domainname"`
	DdnsGenerateHostname        types.Bool                       `tfsdk:"ddns_generate_hostname"`
	DelegatedMember             types.Object                     `tfsdk:"delegated_member"`
	DenyAllClients              types.Bool                       `tfsdk:"deny_all_clients"`
	DenyBootp                   types.Bool                       `tfsdk:"deny_bootp"`
	EmailList                   internaltypes.UnorderedListValue `tfsdk:"email_list"`
	EnableDdns                  types.Bool                       `tfsdk:"enable_ddns"`
	EnableDhcpThresholds        types.Bool                       `tfsdk:"enable_dhcp_thresholds"`
	EnableEmailWarnings         types.Bool                       `tfsdk:"enable_email_warnings"`
	EnablePxeLeaseTime          types.Bool                       `tfsdk:"enable_pxe_lease_time"`
	EnableSnmpWarnings          types.Bool                       `tfsdk:"enable_snmp_warnings"`
	Exclude                     types.List                       `tfsdk:"exclude"`
	ExtAttrs                    types.Map                        `tfsdk:"ext_attrs"`
	ExtAttrsAll                 types.Map                        `tfsdk:"ext_attrs_all"`
	FailoverAssociation         types.String                     `tfsdk:"failover_association"`
	FingerprintFilterRules      types.List                       `tfsdk:"fingerprint_filter_rules"`
	HighWaterMark               types.Int64                      `tfsdk:"high_water_mark"`
	HighWaterMarkReset          types.Int64                      `tfsdk:"high_water_mark_reset"`
	IgnoreDhcpOptionListRequest types.Bool                       `tfsdk:"ignore_dhcp_option_list_request"`
	KnownClients                types.String                     `tfsdk:"known_clients"`
	LeaseScavengeTime           types.Int64                      `tfsdk:"lease_scavenge_time"`
	LogicFilterRules            types.List                       `tfsdk:"logic_filter_rules"`
	LowWaterMark                types.Int64                      `tfsdk:"low_water_mark"`
	LowWaterMarkReset           types.Int64                      `tfsdk:"low_water_mark_reset"`
	MacFilterRules              types.List                       `tfsdk:"mac_filter_rules"`
	Member                      types.Object                     `tfsdk:"member"`
	MsOptions                   types.List                       `tfsdk:"ms_options"`
	MsServer                    types.Object                     `tfsdk:"ms_server"`
	NacFilterRules              types.List                       `tfsdk:"nac_filter_rules"`
	Name                        types.String                     `tfsdk:"name"`
	Nextserver                  types.String                     `tfsdk:"nextserver"`
	NumberOfAddresses           types.Int64                      `tfsdk:"number_of_addresses"`
	Offset                      types.Int64                      `tfsdk:"offset"`
	OptionFilterRules           types.List                       `tfsdk:"option_filter_rules"`
	Options                     types.List                       `tfsdk:"options"`
	PxeLeaseTime                types.Int64                      `tfsdk:"pxe_lease_time"`
	RecycleLeases               types.Bool                       `tfsdk:"recycle_leases"`
	RelayAgentFilterRules       types.List                       `tfsdk:"relay_agent_filter_rules"`
	ServerAssociationType       types.String                     `tfsdk:"server_association_type"`
	UnknownClients              types.String                     `tfsdk:"unknown_clients"`
	UpdateDnsOnLeaseRenewal     types.Bool                       `tfsdk:"update_dns_on_lease_renewal"`
}

var NIOSRangetemplateAttrTypes = map[string]attr.Type{
	"bootfile":                        types.StringType,
	"bootserver":                      types.StringType,
	"cloud_api_compatible":            types.BoolType,
	"comment":                         types.StringType,
	"ddns_domainname":                 types.StringType,
	"ddns_generate_hostname":          types.BoolType,
	"delegated_member":                types.ObjectType{AttrTypes: RangetemplateDelegatedMemberAttrTypes},
	"deny_all_clients":                types.BoolType,
	"deny_bootp":                      types.BoolType,
	"email_list":                      internaltypes.UnorderedListOfStringType,
	"enable_ddns":                     types.BoolType,
	"enable_dhcp_thresholds":          types.BoolType,
	"enable_email_warnings":           types.BoolType,
	"enable_pxe_lease_time":           types.BoolType,
	"enable_snmp_warnings":            types.BoolType,
	"exclude":                         types.ListType{ElemType: types.ObjectType{AttrTypes: RangetemplateExcludeAttrTypes}},
	"ext_attrs":                       types.MapType{ElemType: types.StringType},
	"ext_attrs_all":                   types.MapType{ElemType: types.StringType},
	"failover_association":            types.StringType,
	"fingerprint_filter_rules":        types.ListType{ElemType: types.ObjectType{AttrTypes: RangetemplateFingerprintFilterRulesAttrTypes}},
	"high_water_mark":                 types.Int64Type,
	"high_water_mark_reset":           types.Int64Type,
	"ignore_dhcp_option_list_request": types.BoolType,
	"known_clients":                   types.StringType,
	"lease_scavenge_time":             types.Int64Type,
	"logic_filter_rules":              types.ListType{ElemType: types.ObjectType{AttrTypes: RangetemplateLogicFilterRulesAttrTypes}},
	"low_water_mark":                  types.Int64Type,
	"low_water_mark_reset":            types.Int64Type,
	"mac_filter_rules":                types.ListType{ElemType: types.ObjectType{AttrTypes: RangetemplateMacFilterRulesAttrTypes}},
	"member":                          types.ObjectType{AttrTypes: RangetemplateMemberAttrTypes},
	"ms_options":                      types.ListType{ElemType: types.ObjectType{AttrTypes: RangetemplateMsOptionsAttrTypes}},
	"ms_server":                       types.ObjectType{AttrTypes: RangetemplateMsServerAttrTypes},
	"nac_filter_rules":                types.ListType{ElemType: types.ObjectType{AttrTypes: RangetemplateNacFilterRulesAttrTypes}},
	"name":                            types.StringType,
	"nextserver":                      types.StringType,
	"number_of_addresses":             types.Int64Type,
	"offset":                          types.Int64Type,
	"option_filter_rules":             types.ListType{ElemType: types.ObjectType{AttrTypes: RangetemplateOptionFilterRulesAttrTypes}},
	"options":                         types.ListType{ElemType: types.ObjectType{AttrTypes: RangetemplateOptionsAttrTypes}},
	"pxe_lease_time":                  types.Int64Type,
	"recycle_leases":                  types.BoolType,
	"relay_agent_filter_rules":        types.ListType{ElemType: types.ObjectType{AttrTypes: RangetemplateRelayAgentFilterRulesAttrTypes}},
	"server_association_type":         types.StringType,
	"unknown_clients":                 types.StringType,
	"update_dns_on_lease_renewal":     types.BoolType,
}

const (
	RangetemplateReturnFields = "bootfile,bootserver,cloud_api_compatible,comment,ddns_domainname,ddns_generate_hostname,delegated_member,deny_all_clients,deny_bootp,email_list,enable_ddns,enable_dhcp_thresholds,enable_email_warnings,enable_pxe_lease_time,enable_snmp_warnings,exclude,extattrs,failover_association,fingerprint_filter_rules,high_water_mark,high_water_mark_reset,ignore_dhcp_option_list_request,known_clients,lease_scavenge_time,logic_filter_rules,low_water_mark,low_water_mark_reset,mac_filter_rules,member,ms_options,ms_server,nac_filter_rules,name,nextserver,number_of_addresses,offset,option_filter_rules,options,pxe_lease_time,recycle_leases,relay_agent_filter_rules,server_association_type,unknown_clients,update_dns_on_lease_renewal,use_bootfile,use_bootserver,use_ddns_domainname,use_ddns_generate_hostname,use_deny_bootp,use_email_list,use_enable_ddns,use_enable_dhcp_thresholds,use_ignore_dhcp_option_list_request,use_known_clients,use_lease_scavenge_time,use_logic_filter_rules,use_ms_options,use_nextserver,use_options,use_pxe_lease_time,use_recycle_leases,use_unknown_clients,use_update_dns_on_lease_renewal"
)

var RangetemplateResourceSchemaAttributes = map[string]schema.Attribute{
	"id": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The reference to the object.",
	},
	"nios": schema.SingleNestedAttribute{
		Optional:            true,
		MarkdownDescription: "NIOS backend-specific fields.",
		Attributes:          RangetemplateResourceNiosSchemaAttributes,
	},
}

var RangetemplateResourceNiosSchemaAttributes = map[string]schema.Attribute{
	"bootfile": schema.StringAttribute{
		Optional: true,
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
	"cloud_api_compatible": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(true),
		MarkdownDescription: "This flag controls whether this template can be used to create network objects in a cloud-computing deployment.",
	},
	"comment": schema.StringAttribute{
		Optional: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
			customvalidator.ValidateTrimmedString(),
			stringvalidator.LengthBetween(0, 256),
		},
		MarkdownDescription: "A descriptive comment of a range template object.",
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
	"delegated_member": schema.SingleNestedAttribute{
		Attributes:          RangetemplateDelegatedMemberResourceSchemaAttributes,
		Optional:            true,
		MarkdownDescription: "",
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
		MarkdownDescription: "Determines if BOOTP settings are disabled and BOOTP requests will be denied.",
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
		MarkdownDescription: "Determines if the DHCP server sends DDNS updates to DNS servers in the same Grid, and to external DNS servers.",
	},
	"enable_dhcp_thresholds": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Determines if DHCP thresholds are enabled for the range.",
	},
	"enable_email_warnings": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Determines if DHCP threshold warnings are sent through email.",
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
		MarkdownDescription: "Determines if DHCP threshold warnings are sent through SNMP.",
	},
	"exclude": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: RangetemplateExcludeResourceSchemaAttributes,
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
			stringvalidator.AlsoRequires(path.MatchRelative().AtParent().AtName("server_association_type")),
		},
		MarkdownDescription: "The name of the failover association: the server in this failover association will serve the IPv4 range in case the main server is out of service. {rangetemplate:rangetemplate} must be set to 'FAILOVER' or 'FAILOVER_MS' if you want the failover association specified here to serve the range.",
	},
	"fingerprint_filter_rules": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: RangetemplateFingerprintFilterRulesResourceSchemaAttributes,
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
	"known_clients": schema.StringAttribute{
		Validators: []validator.String{
			stringvalidator.OneOf("Allow", "Deny"),
		},
		Optional:            true,
		MarkdownDescription: "Permission for known clients. If set to 'Deny' known clients will be denied IP addresses. Known clients include roaming hosts and clients with fixed addresses or DHCP host entries. Unknown clients include clients that are not roaming hosts and clients that do not have fixed addresses or DHCP host entries.",
	},
	"lease_scavenge_time": schema.Int64Attribute{
		Optional:            true,
		Computed:            true,
		Default:             int64default.StaticInt64(-1),
		MarkdownDescription: "An integer that specifies the period of time (in seconds) that frees and backs up leases remained in the database before they are automatically deleted. To disable lease scavenging, set the parameter to -1. The minimum positive value must be greater than 86400 seconds (1 day).",
	},
	"logic_filter_rules": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: RangetemplateLogicFilterRulesResourceSchemaAttributes,
		},
		Optional: true,
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
		},
		MarkdownDescription: "This field contains the logic filters to be applied on this range. This list corresponds to the match rules that are written to the dhcpd configuration file.",
	},
	"low_water_mark": schema.Int64Attribute{
		Optional: true,
		Computed: true,
		Default:  int64default.StaticInt64(0),
		Validators: []validator.Int64{
			int64validator.Between(1, 100),
		},
		MarkdownDescription: "The percentage of DHCP range usage below which the Infoblox appliance generates a syslog message and sends a warning (if enabled). A number that specifies the percentage of allocated addresses. The range is from 1 to 100.",
	},
	"low_water_mark_reset": schema.Int64Attribute{
		Optional: true,
		Computed: true,
		Default:  int64default.StaticInt64(10),
		Validators: []validator.Int64{
			int64validator.Between(1, 100),
		},
		MarkdownDescription: "The percentage of DHCP range usage threshold below which range usage is not expected and may warrant your attention. When the low watermark is crossed, the Infoblox appliance generates a syslog message and sends a warning (if enabled). A number that specifies the percentage of allocated addresses. The range is from 1 to 100. The low watermark reset value must be higher than the low watermark value.",
	},
	"mac_filter_rules": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: RangetemplateMacFilterRulesResourceSchemaAttributes,
		},
		Optional: true,
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
		},
		MarkdownDescription: "This field contains the MAC filters to be applied to this range. The appliance uses the matching rules of these filters to select the address range from which it assigns a lease.",
	},
	"member": schema.SingleNestedAttribute{
		Attributes:          RangetemplateMemberResourceSchemaAttributes,
		Optional:            true,
		MarkdownDescription: "",
	},
	"ms_options": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: RangetemplateMsOptionsResourceSchemaAttributes,
		},
		Optional: true,
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
		},
		MarkdownDescription: "The Microsoft DHCP options for this range.",
	},
	"ms_server": schema.SingleNestedAttribute{
		Attributes:          RangetemplateMsServerResourceSchemaAttributes,
		Optional:            true,
		MarkdownDescription: "",
	},
	"nac_filter_rules": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: RangetemplateNacFilterRulesResourceSchemaAttributes,
		},
		Optional: true,
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
		},
		MarkdownDescription: "This field contains the NAC filters to be applied to this range. The appliance uses the matching rules of these filters to select the address range from which it assigns a lease.",
	},
	"name": schema.StringAttribute{
		Required: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "The name of a range template object.",
	},
	"nextserver": schema.StringAttribute{
		Optional: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
			customvalidator.IsValidIPv4OrFQDN(),
		},
		MarkdownDescription: "The name in FQDN and/or IPv4 Address format of the next server that the host needs to boot.",
	},
	"number_of_addresses": schema.Int64Attribute{
		Required:            true,
		MarkdownDescription: "The number of addresses for this range.",
	},
	"offset": schema.Int64Attribute{
		Required:            true,
		MarkdownDescription: "The start address offset for this range.",
	},
	"option_filter_rules": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: RangetemplateOptionFilterRulesResourceSchemaAttributes,
		},
		Optional: true,
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
		},
		MarkdownDescription: "This field contains the Option filters to be applied to this range. The appliance uses the matching rules of these filters to select the address range from which it assigns a lease.",
	},
	"options": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: RangetemplateOptionsResourceSchemaAttributes,
		},
		Optional: true,
		Computed: true,
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
		},
		MarkdownDescription: "An array of DHCP option dhcpoption structs that lists the DHCP options associated with the object.",
	},
	"pxe_lease_time": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "The PXE lease time value for a range object. Some hosts use PXE (Preboot Execution Environment) to boot remotely from a server. To better manage your IP resources, set a different lease time for PXE boot requests. You can configure the DHCP server to allocate an IP address with a shorter lease time to hosts that send PXE boot requests, so IP addresses are not leased longer than necessary. A 32-bit unsigned integer that represents the duration, in seconds, for which the update is cached. Zero indicates that the update is not cached.",
	},
	"recycle_leases": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(true),
		MarkdownDescription: "If the field is set to True, the leases are kept in the Recycle Bin until one week after expiration. Otherwise, the leases are permanently deleted.",
	},
	"relay_agent_filter_rules": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: RangetemplateRelayAgentFilterRulesResourceSchemaAttributes,
		},
		Optional: true,
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
		},
		MarkdownDescription: "This field contains the Relay Agent filters to be applied to this range. The appliance uses the matching rules of these filters to select the address range from which it assigns a lease.",
	},
	"server_association_type": schema.StringAttribute{
		Default: stringdefault.StaticString("NONE"),
		Validators: []validator.String{
			stringvalidator.OneOf("FAILOVER", "MEMBER", "MS_FAILOVER", "MS_SERVER", "NONE"),
		},
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The type of server that is going to serve the range.",
	},
	"unknown_clients": schema.StringAttribute{
		Validators: []validator.String{
			stringvalidator.OneOf("Allow", "Deny"),
		},
		Optional:            true,
		MarkdownDescription: "Permission for unknown clients. If set to 'Deny' unknown clients will be denied IP addresses. Known clients include roaming hosts and clients with fixed addresses or DHCP host entries. Unknown clients include clients that are not roaming hosts and clients that do not have fixed addresses or DHCP host entries.",
	},
	"update_dns_on_lease_renewal": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "This field controls whether the DHCP server updates DNS when a DHCP lease is renewed.",
	},
}

// Expand converts the TF model to the infoblox core model
func (m *RangetemplateModel) Expand(ctx context.Context, diags *diag.Diagnostics, isCreate bool) *coremodel.Rangetemplate {
	if m == nil {
		return nil
	}

	obj := &coremodel.Rangetemplate{}

	// Expand NIOS nested attribute (returns nil if not present)
	niosModel := flex.ExpandNestedObject[NIOSRangetemplateModel](ctx, m.NIOS, diags)
	if niosModel != nil {
		obj.NIOS = niosModel.Expand(ctx, diags)
	}

	return obj
}

// Expand converts the NIOS TF model to the core model.
func (m *NIOSRangetemplateModel) Expand(ctx context.Context, diags *diag.Diagnostics) *coremodel.NIOSRangetemplateExt {
	return &coremodel.NIOSRangetemplateExt{
		Bootfile:                    flex.ExpandStringPointerNullAsEmpty(m.Bootfile),
		Bootserver:                  flex.ExpandStringPointerNullAsEmpty(m.Bootserver),
		CloudApiCompatible:          flex.ExpandBoolPointer(m.CloudApiCompatible),
		Comment:                     flex.ExpandStringPointerNullAsEmpty(m.Comment),
		DdnsDomainname:              flex.ExpandStringPointerNullAsEmpty(m.DdnsDomainname),
		DdnsGenerateHostname:        flex.ExpandBoolPointer(m.DdnsGenerateHostname),
		DelegatedMember:             ExpandRangetemplateDelegatedMember(ctx, m.DelegatedMember, diags),
		DenyAllClients:              flex.ExpandBoolPointer(m.DenyAllClients),
		DenyBootp:                   flex.ExpandBoolPointer(m.DenyBootp),
		EmailList:                   flex.ExpandFrameworkListString(ctx, m.EmailList, diags),
		EnableDdns:                  flex.ExpandBoolPointer(m.EnableDdns),
		EnableDhcpThresholds:        flex.ExpandBoolPointer(m.EnableDhcpThresholds),
		EnableEmailWarnings:         flex.ExpandBoolPointer(m.EnableEmailWarnings),
		EnablePxeLeaseTime:          flex.ExpandBoolPointer(m.EnablePxeLeaseTime),
		EnableSnmpWarnings:          flex.ExpandBoolPointer(m.EnableSnmpWarnings),
		Exclude:                     flex.ExpandFrameworkListNestedBlock(ctx, m.Exclude, diags, ExpandRangetemplateExclude),
		ExtAttrs:                    flex.ExpandMapStringAny(ctx, m.ExtAttrs, diags),
		FailoverAssociation:         flex.ExpandStringPointerNullAsEmpty(m.FailoverAssociation),
		FingerprintFilterRules:      flex.ExpandFrameworkListNestedBlock(ctx, m.FingerprintFilterRules, diags, ExpandRangetemplateFingerprintFilterRules),
		HighWaterMark:               flex.ExpandInt64Pointer(m.HighWaterMark),
		HighWaterMarkReset:          flex.ExpandInt64Pointer(m.HighWaterMarkReset),
		IgnoreDhcpOptionListRequest: flex.ExpandBoolPointer(m.IgnoreDhcpOptionListRequest),
		KnownClients:                flex.ExpandStringPointer(m.KnownClients),
		LeaseScavengeTime:           flex.ExpandInt64Pointer(m.LeaseScavengeTime),
		LogicFilterRules:            flex.ExpandFrameworkListNestedBlock(ctx, m.LogicFilterRules, diags, ExpandRangetemplateLogicFilterRules),
		LowWaterMark:                flex.ExpandInt64Pointer(m.LowWaterMark),
		LowWaterMarkReset:           flex.ExpandInt64Pointer(m.LowWaterMarkReset),
		MacFilterRules:              flex.ExpandFrameworkListNestedBlock(ctx, m.MacFilterRules, diags, ExpandRangetemplateMacFilterRules),
		Member:                      ExpandRangetemplateMember(ctx, m.Member, diags),
		MsOptions:                   flex.ExpandFrameworkListNestedBlock(ctx, m.MsOptions, diags, ExpandRangetemplateMsOptions),
		MsServer:                    ExpandRangetemplateMsServer(ctx, m.MsServer, diags),
		NacFilterRules:              flex.ExpandFrameworkListNestedBlock(ctx, m.NacFilterRules, diags, ExpandRangetemplateNacFilterRules),
		Name:                        flex.ExpandStringPointerNullAsEmpty(m.Name),
		Nextserver:                  flex.ExpandStringPointerNullAsEmpty(m.Nextserver),
		NumberOfAddresses:           flex.ExpandInt64Pointer(m.NumberOfAddresses),
		Offset:                      flex.ExpandInt64Pointer(m.Offset),
		OptionFilterRules:           flex.ExpandFrameworkListNestedBlock(ctx, m.OptionFilterRules, diags, ExpandRangetemplateOptionFilterRules),
		Options:                     flex.ExpandFrameworkListNestedBlock(ctx, m.Options, diags, ExpandRangetemplateOptions),
		PxeLeaseTime:                flex.ExpandInt64Pointer(m.PxeLeaseTime),
		RecycleLeases:               flex.ExpandBoolPointer(m.RecycleLeases),
		RelayAgentFilterRules:       flex.ExpandFrameworkListNestedBlock(ctx, m.RelayAgentFilterRules, diags, ExpandRangetemplateRelayAgentFilterRules),
		ServerAssociationType:       flex.ExpandStringPointerNullAsEmpty(m.ServerAssociationType),
		UnknownClients:              flex.ExpandStringPointer(m.UnknownClients),
		UpdateDnsOnLeaseRenewal:     flex.ExpandBoolPointer(m.UpdateDnsOnLeaseRenewal),
	}
}

// ApplyRangetemplateNIOSUseFlags derives NIOS use flags from the raw config
// value(s) and writes them onto the core model. A flag is true when the user
// set any of its governed value fields in config.
func ApplyRangetemplateNIOSUseFlags(ctx context.Context, config tfsdk.Config, obj *coremodel.Rangetemplate, diags *diag.Diagnostics) {
	if obj == nil || obj.NIOS == nil {
		return
	}
	obj.NIOS.UseBootfile = flex.DeriveUseFlag(ctx, config, diags, path.Root("nios").AtName("bootfile"))
	obj.NIOS.UseBootserver = flex.DeriveUseFlag(ctx, config, diags, path.Root("nios").AtName("bootserver"))
	obj.NIOS.UseDdnsDomainname = flex.DeriveUseFlag(ctx, config, diags, path.Root("nios").AtName("ddns_domainname"))
	obj.NIOS.UseDdnsGenerateHostname = flex.DeriveUseFlag(ctx, config, diags, path.Root("nios").AtName("ddns_generate_hostname"))
	obj.NIOS.UseDenyBootp = flex.DeriveUseFlag(ctx, config, diags, path.Root("nios").AtName("deny_bootp"))
	obj.NIOS.UseEmailList = flex.DeriveUseFlag(ctx, config, diags, path.Root("nios").AtName("email_list"))
	obj.NIOS.UseEnableDdns = flex.DeriveUseFlag(ctx, config, diags, path.Root("nios").AtName("enable_ddns"))
	obj.NIOS.UseEnableDhcpThresholds = flex.DeriveUseFlag(ctx, config, diags, path.Root("nios").AtName("enable_dhcp_thresholds"))
	obj.NIOS.UseIgnoreDhcpOptionListRequest = flex.DeriveUseFlag(ctx, config, diags, path.Root("nios").AtName("ignore_dhcp_option_list_request"))
	obj.NIOS.UseKnownClients = flex.DeriveUseFlag(ctx, config, diags, path.Root("nios").AtName("known_clients"))
	obj.NIOS.UseLeaseScavengeTime = flex.DeriveUseFlag(ctx, config, diags, path.Root("nios").AtName("lease_scavenge_time"))
	obj.NIOS.UseLogicFilterRules = flex.DeriveUseFlag(ctx, config, diags, path.Root("nios").AtName("logic_filter_rules"))
	obj.NIOS.UseMsOptions = flex.DeriveUseFlag(ctx, config, diags, path.Root("nios").AtName("ms_options"))
	obj.NIOS.UseNextserver = flex.DeriveUseFlag(ctx, config, diags, path.Root("nios").AtName("nextserver"))
	obj.NIOS.UseOptions = flex.DeriveUseFlag(ctx, config, diags, path.Root("nios").AtName("options"))
	obj.NIOS.UsePxeLeaseTime = flex.DeriveUseFlag(ctx, config, diags, path.Root("nios").AtName("pxe_lease_time"))
	obj.NIOS.UseRecycleLeases = flex.DeriveUseFlag(ctx, config, diags, path.Root("nios").AtName("recycle_leases"))
	obj.NIOS.UseUnknownClients = flex.DeriveUseFlag(ctx, config, diags, path.Root("nios").AtName("unknown_clients"))
	obj.NIOS.UseUpdateDnsOnLeaseRenewal = flex.DeriveUseFlag(ctx, config, diags, path.Root("nios").AtName("update_dns_on_lease_renewal"))
}

// Flatten populates the TF model from a core response.
func (m *RangetemplateModel) Flatten(ctx context.Context, resp *coremodel.Rangetemplate, diags *diag.Diagnostics) {
	if resp == nil {
		return
	}

	m.Id = flex.FlattenStringPointer(resp.Id)

	// Extract existing NIOS model, flatten API response onto it, convert back
	niosModel := flex.ExpandNestedObject[NIOSRangetemplateModel](ctx, m.NIOS, diags)
	if niosModel == nil {
		niosModel = &NIOSRangetemplateModel{}
	}
	plannedNIOS := flex.ExpandNestedObject[NIOSRangetemplateModel](ctx, m.NIOS, diags)
	niosModel.Flatten(ctx, resp.NIOS, diags)
	if resp.NIOS != nil {
		PostFlattenRangetemplateNIOS(ctx, plannedNIOS, niosModel, diags)
		m.NIOS = flex.FlattenNestedObject(ctx, niosModel, NIOSRangetemplateAttrTypes, diags)
	} else {
		m.NIOS = types.ObjectNull(NIOSRangetemplateAttrTypes)
	}

}

// Flatten merges API response onto existing NIOS model.
func (m *NIOSRangetemplateModel) Flatten(ctx context.Context, from *coremodel.NIOSRangetemplateExt, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	planExtAttrs := m.ExtAttrs
	if planExtAttrs.IsUnknown() {
		planExtAttrs = types.MapNull(types.StringType)
	}
	m.Bootfile = flex.FlattenStringPointerEmptyAsNull(from.Bootfile)
	m.Bootserver = flex.FlattenStringPointerEmptyAsNull(from.Bootserver)
	m.CloudApiCompatible = flex.FlattenBoolPointer(from.CloudApiCompatible)
	m.Comment = flex.FlattenStringPointerEmptyAsNull(from.Comment)
	m.DdnsDomainname = flex.FlattenStringPointerEmptyAsNull(from.DdnsDomainname)
	m.DdnsGenerateHostname = flex.FlattenBoolPointer(from.DdnsGenerateHostname)
	m.DelegatedMember = FlattenRangetemplateDelegatedMember(ctx, from.DelegatedMember, diags)
	m.DenyAllClients = flex.FlattenBoolPointer(from.DenyAllClients)
	m.DenyBootp = flex.FlattenBoolPointer(from.DenyBootp)
	m.EmailList = flex.FlattenFrameworkUnorderedListString(ctx, from.EmailList, diags)
	m.EnableDdns = flex.FlattenBoolPointer(from.EnableDdns)
	m.EnableDhcpThresholds = flex.FlattenBoolPointer(from.EnableDhcpThresholds)
	m.EnableEmailWarnings = flex.FlattenBoolPointer(from.EnableEmailWarnings)
	m.EnablePxeLeaseTime = flex.FlattenBoolPointer(from.EnablePxeLeaseTime)
	m.EnableSnmpWarnings = flex.FlattenBoolPointer(from.EnableSnmpWarnings)
	m.Exclude = flex.FlattenFrameworkListNestedBlock(ctx, from.Exclude, RangetemplateExcludeAttrTypes, diags, FlattenRangetemplateExclude)
	m.ExtAttrs, m.ExtAttrsAll = flex.FlattenEAs(planExtAttrs, from.ExtAttrs)
	m.FailoverAssociation = flex.FlattenStringPointerEmptyAsNull(from.FailoverAssociation)
	m.FingerprintFilterRules = flex.FlattenFrameworkListNestedBlock(ctx, from.FingerprintFilterRules, RangetemplateFingerprintFilterRulesAttrTypes, diags, FlattenRangetemplateFingerprintFilterRules)
	m.HighWaterMark = flex.FlattenInt64Pointer(from.HighWaterMark)
	m.HighWaterMarkReset = flex.FlattenInt64Pointer(from.HighWaterMarkReset)
	m.IgnoreDhcpOptionListRequest = flex.FlattenBoolPointer(from.IgnoreDhcpOptionListRequest)
	m.KnownClients = flex.FlattenStringPointerEmptyAsNull(from.KnownClients)
	m.LeaseScavengeTime = flex.FlattenInt64Pointer(from.LeaseScavengeTime)
	m.LogicFilterRules = flex.FlattenFrameworkListNestedBlock(ctx, from.LogicFilterRules, RangetemplateLogicFilterRulesAttrTypes, diags, FlattenRangetemplateLogicFilterRules)
	m.LowWaterMark = flex.FlattenInt64Pointer(from.LowWaterMark)
	m.LowWaterMarkReset = flex.FlattenInt64Pointer(from.LowWaterMarkReset)
	m.MacFilterRules = flex.FlattenFrameworkListNestedBlock(ctx, from.MacFilterRules, RangetemplateMacFilterRulesAttrTypes, diags, FlattenRangetemplateMacFilterRules)
	m.Member = FlattenRangetemplateMember(ctx, from.Member, diags)
	m.MsOptions = flex.FlattenFrameworkListNestedBlock(ctx, from.MsOptions, RangetemplateMsOptionsAttrTypes, diags, FlattenRangetemplateMsOptions)
	m.MsServer = FlattenRangetemplateMsServer(ctx, from.MsServer, diags)
	m.NacFilterRules = flex.FlattenFrameworkListNestedBlock(ctx, from.NacFilterRules, RangetemplateNacFilterRulesAttrTypes, diags, FlattenRangetemplateNacFilterRules)
	m.Name = flex.FlattenStringPointerEmptyAsNull(from.Name)
	m.Nextserver = flex.FlattenStringPointerEmptyAsNull(from.Nextserver)
	m.NumberOfAddresses = flex.FlattenInt64Pointer(from.NumberOfAddresses)
	m.Offset = flex.FlattenInt64Pointer(from.Offset)
	m.OptionFilterRules = flex.FlattenFrameworkListNestedBlock(ctx, from.OptionFilterRules, RangetemplateOptionFilterRulesAttrTypes, diags, FlattenRangetemplateOptionFilterRules)
	m.Options = flex.FlattenFrameworkListNestedBlock(ctx, from.Options, RangetemplateOptionsAttrTypes, diags, FlattenRangetemplateOptions)
	m.PxeLeaseTime = flex.FlattenInt64Pointer(from.PxeLeaseTime)
	m.RecycleLeases = flex.FlattenBoolPointer(from.RecycleLeases)
	m.RelayAgentFilterRules = flex.FlattenFrameworkListNestedBlock(ctx, from.RelayAgentFilterRules, RangetemplateRelayAgentFilterRulesAttrTypes, diags, FlattenRangetemplateRelayAgentFilterRules)
	m.ServerAssociationType = flex.FlattenStringPointerEmptyAsNull(from.ServerAssociationType)
	m.UnknownClients = flex.FlattenStringPointerEmptyAsNull(from.UnknownClients)
	m.UpdateDnsOnLeaseRenewal = flex.FlattenBoolPointer(from.UpdateDnsOnLeaseRenewal)
}
