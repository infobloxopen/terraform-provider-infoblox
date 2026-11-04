package dhcp

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-nettypes/cidrtypes"
	"github.com/hashicorp/terraform-plugin-framework-nettypes/iptypes"
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
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
	immutable "github.com/infobloxopen/terraform-provider-infoblox/internal/planmodifiers/immutable"
	importmod "github.com/infobloxopen/terraform-provider-infoblox/internal/planmodifiers/import"
	customvalidator "github.com/infobloxopen/terraform-provider-infoblox/internal/validator"
)

type Ipv6rangeModel struct {
	Id            types.String `tfsdk:"id"`
	UpdateTrigger types.String `tfsdk:"update_trigger"`
	NIOS          types.Object `tfsdk:"nios"`
	UDDI          types.Object `tfsdk:"uddi"`
}

var Ipv6rangeAttrTypes = map[string]attr.Type{
	"id":             types.StringType,
	"update_trigger": types.StringType,
	"nios":           types.ObjectType{AttrTypes: NIOSIpv6rangeAttrTypes},
	"uddi":           types.ObjectType{AttrTypes: UDDIIpv6rangeAttrTypes},
}

type NIOSIpv6rangeModel struct {
	AddressType                      types.String         `tfsdk:"address_type"`
	CloudInfo                        types.Object         `tfsdk:"cloud_info"`
	Comment                          types.String         `tfsdk:"comment"`
	Disable                          types.Bool           `tfsdk:"disable"`
	DiscoveryBasicPollSettings       types.Object         `tfsdk:"discovery_basic_poll_settings"`
	DiscoveryBlackoutSetting         types.Object         `tfsdk:"discovery_blackout_setting"`
	DiscoveryMember                  types.String         `tfsdk:"discovery_member"`
	EnableDiscovery                  types.Bool           `tfsdk:"enable_discovery"`
	EnableImmediateDiscovery         types.Bool           `tfsdk:"enable_immediate_discovery"`
	EndAddr                          iptypes.IPv6Address  `tfsdk:"end_addr"`
	Exclude                          types.List           `tfsdk:"exclude"`
	ExtAttrs                         types.Map            `tfsdk:"ext_attrs"`
	ExtAttrsAll                      types.Map            `tfsdk:"ext_attrs_all"`
	Ipv6EndPrefix                    iptypes.IPv6Address  `tfsdk:"ipv6_end_prefix"`
	Ipv6PrefixBits                   types.Int64          `tfsdk:"ipv6_prefix_bits"`
	Ipv6StartPrefix                  iptypes.IPv6Address  `tfsdk:"ipv6_start_prefix"`
	LogicFilterRules                 types.List           `tfsdk:"logic_filter_rules"`
	Member                           types.Object         `tfsdk:"member"`
	Name                             types.String         `tfsdk:"name"`
	Network                          cidrtypes.IPv6Prefix `tfsdk:"network"`
	NetworkView                      types.String         `tfsdk:"network_view"`
	OptionFilterRules                types.List           `tfsdk:"option_filter_rules"`
	PortControlBlackoutSetting       types.Object         `tfsdk:"port_control_blackout_setting"`
	RecycleLeases                    types.Bool           `tfsdk:"recycle_leases"`
	RestartIfNeeded                  types.Bool           `tfsdk:"restart_if_needed"`
	SamePortControlDiscoveryBlackout types.Bool           `tfsdk:"same_port_control_discovery_blackout"`
	ServerAssociationType            types.String         `tfsdk:"server_association_type"`
	StartAddr                        iptypes.IPv6Address  `tfsdk:"start_addr"`
	SubscribeSettings                types.Object         `tfsdk:"subscribe_settings"`
	Template                         types.String         `tfsdk:"template"`
}

var NIOSIpv6rangeAttrTypes = map[string]attr.Type{
	"address_type":                         types.StringType,
	"cloud_info":                           types.ObjectType{AttrTypes: Ipv6rangeCloudInfoAttrTypes},
	"comment":                              types.StringType,
	"disable":                              types.BoolType,
	"discovery_basic_poll_settings":        types.ObjectType{AttrTypes: Ipv6rangeDiscoveryBasicPollSettingsAttrTypes},
	"discovery_blackout_setting":           types.ObjectType{AttrTypes: Ipv6rangeDiscoveryBlackoutSettingAttrTypes},
	"discovery_member":                     types.StringType,
	"enable_discovery":                     types.BoolType,
	"enable_immediate_discovery":           types.BoolType,
	"end_addr":                             iptypes.IPv6AddressType{},
	"exclude":                              types.ListType{ElemType: types.ObjectType{AttrTypes: Ipv6rangeExcludeAttrTypes}},
	"ext_attrs":                            types.MapType{ElemType: types.StringType},
	"ext_attrs_all":                        types.MapType{ElemType: types.StringType},
	"ipv6_end_prefix":                      iptypes.IPv6AddressType{},
	"ipv6_prefix_bits":                     types.Int64Type,
	"ipv6_start_prefix":                    iptypes.IPv6AddressType{},
	"logic_filter_rules":                   types.ListType{ElemType: types.ObjectType{AttrTypes: Ipv6rangeLogicFilterRulesAttrTypes}},
	"member":                               types.ObjectType{AttrTypes: Ipv6rangeMemberAttrTypes},
	"name":                                 types.StringType,
	"network":                              cidrtypes.IPv6PrefixType{},
	"network_view":                         types.StringType,
	"option_filter_rules":                  types.ListType{ElemType: types.ObjectType{AttrTypes: Ipv6rangeOptionFilterRulesAttrTypes}},
	"port_control_blackout_setting":        types.ObjectType{AttrTypes: Ipv6rangePortControlBlackoutSettingAttrTypes},
	"recycle_leases":                       types.BoolType,
	"restart_if_needed":                    types.BoolType,
	"same_port_control_discovery_blackout": types.BoolType,
	"server_association_type":              types.StringType,
	"start_addr":                           iptypes.IPv6AddressType{},
	"subscribe_settings":                   types.ObjectType{AttrTypes: Ipv6rangeSubscribeSettingsAttrTypes},
	"template":                             types.StringType,
}

type UDDIIpv6rangeModel struct {
	Comment            types.String        `tfsdk:"comment"`
	DhcpHost           types.String        `tfsdk:"dhcp_host"`
	DhcpOptions        types.List          `tfsdk:"dhcp_options"`
	DisableDhcp        types.Bool          `tfsdk:"disable_dhcp"`
	End                iptypes.IPv6Address `tfsdk:"end"`
	ExclusionRanges    types.List          `tfsdk:"exclusion_ranges"`
	Filters            types.List          `tfsdk:"filters"`
	InheritanceSources types.Object        `tfsdk:"inheritance_sources"`
	Name               types.String        `tfsdk:"name"`
	Space              types.String        `tfsdk:"space"`
	Start              iptypes.IPv6Address `tfsdk:"start"`
	Tags               types.Map           `tfsdk:"tags"`
	TagsAll            types.Map           `tfsdk:"tags_all"`
	Threshold          types.Object        `tfsdk:"threshold"`
}

var UDDIIpv6rangeAttrTypes = map[string]attr.Type{
	"comment":             types.StringType,
	"dhcp_host":           types.StringType,
	"dhcp_options":        types.ListType{ElemType: types.ObjectType{AttrTypes: OptionItemAttrTypes}},
	"disable_dhcp":        types.BoolType,
	"end":                 iptypes.IPv6AddressType{},
	"exclusion_ranges":    types.ListType{ElemType: types.ObjectType{AttrTypes: ExclusionRangeAttrTypes}},
	"filters":             types.ListType{ElemType: types.ObjectType{AttrTypes: AccessFilterAttrTypes}},
	"inheritance_sources": types.ObjectType{AttrTypes: DHCPOptionsInheritanceAttrTypes},
	"name":                types.StringType,
	"space":               types.StringType,
	"start":               iptypes.IPv6AddressType{},
	"tags":                types.MapType{ElemType: types.StringType},
	"tags_all":            types.MapType{ElemType: types.StringType},
	"threshold":           types.ObjectType{AttrTypes: UtilizationThresholdAttrTypes},
}

const (
	Ipv6rangeInheritanceType = "full"
	Ipv6rangeReturnFields    = "address_type,cloud_info,comment,disable,discover_now_status,discovery_basic_poll_settings,discovery_blackout_setting,discovery_member,enable_discovery,end_addr,endpoint_sources,exclude,extattrs,ipv6_end_prefix,ipv6_prefix_bits,ipv6_start_prefix,logic_filter_rules,member,name,network,network_view,option_filter_rules,port_control_blackout_setting,recycle_leases,same_port_control_discovery_blackout,server_association_type,start_addr,subscribe_settings,use_blackout_setting,use_discovery_basic_polling_settings,use_enable_discovery,use_logic_filter_rules,use_recycle_leases,use_subscribe_settings"
)

var Ipv6rangeResourceSchemaAttributes = map[string]schema.Attribute{
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
		Attributes:          Ipv6rangeResourceNiosSchemaAttributes,
	},
	"uddi": schema.SingleNestedAttribute{
		Optional:            true,
		MarkdownDescription: "UDDI backend-specific fields.",
		Attributes:          Ipv6rangeResourceUddiSchemaAttributes,
	},
}

var Ipv6rangeResourceNiosSchemaAttributes = map[string]schema.Attribute{
	"address_type": schema.StringAttribute{
		Default:  stringdefault.StaticString("ADDRESS"),
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "Type of a DHCP IPv6 Range object. Valid values are \"ADDRESS\", \"PREFIX\", or \"BOTH\". When the address type is \"ADDRESS\", values for the 'start_addr' and 'end_addr' members are required. When the address type is \"PREFIX\", values for 'ipv6_start_prefix', 'ipv6_end_prefix', and 'ipv6_prefix_bits' are required. When the address type is \"BOTH\", values for 'start_addr', 'end_addr', 'ipv6_start_prefix', 'ipv6_end_prefix', and 'ipv6_prefix_bits' are all required.",
	},
	"cloud_info": schema.SingleNestedAttribute{
		Attributes:          Ipv6rangeCloudInfoResourceSchemaAttributes,
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Structure containing all cloud API related information for this object.",
	},
	"comment": schema.StringAttribute{
		Optional: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
			stringvalidator.LengthBetween(0, 256),
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "Comment for the range; maximum 256 characters.",
	},
	"disable": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Determines whether a range is disabled or not. When this is set to False, the range is enabled.",
	},
	"discovery_basic_poll_settings": schema.SingleNestedAttribute{
		Attributes:          Ipv6rangeDiscoveryBasicPollSettingsResourceSchemaAttributes,
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "",
	},
	"discovery_blackout_setting": schema.SingleNestedAttribute{
		Attributes:          Ipv6rangeDiscoveryBlackoutSettingResourceSchemaAttributes,
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "",
	},
	"discovery_member": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "The member that will run discovery for this range.",
	},
	"enable_discovery": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Determines whether a discovery is enabled or not for this range. When this is set to False, the discovery for this range is disabled.",
	},
	"enable_immediate_discovery": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Determines if the discovery for the range should be immediately enabled.",
	},
	"end_addr": schema.StringAttribute{
		Optional:   true,
		CustomType: iptypes.IPv6AddressType{},
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "The IPv6 Address end address of the DHCP IPv6 range.",
	},
	"exclude": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: Ipv6rangeExcludeResourceSchemaAttributes,
		},
		Optional: true,
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
		},
		MarkdownDescription: "These are ranges of IP addresses that the appliance does not use to assign to clients. You can use these exclusion addresses as static IP addresses. They contain the start and end addresses of the exclusion range, and optionally,information about this exclusion range.",
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
	"ipv6_end_prefix": schema.StringAttribute{
		Optional:   true,
		CustomType: iptypes.IPv6AddressType{},
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "The IPv6 Address end prefix of the DHCP IPv6 range.",
	},
	"ipv6_prefix_bits": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "Prefix bits of the DHCP IPv6 range.",
	},
	"ipv6_start_prefix": schema.StringAttribute{
		Optional:   true,
		CustomType: iptypes.IPv6AddressType{},
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "The IPv6 Address starting prefix of the DHCP IPv6 range.",
	},
	"logic_filter_rules": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: Ipv6rangeLogicFilterRulesResourceSchemaAttributes,
		},
		Optional: true,
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
		},
		MarkdownDescription: "This field contains the logic filters to be applied to this IPv6 range. This list corresponds to the match rules that are written to the DHCPv6 configuration file.",
	},
	"member": schema.SingleNestedAttribute{
		Attributes:          Ipv6rangeMemberResourceSchemaAttributes,
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "",
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
		Required:   true,
		CustomType: cidrtypes.IPv6PrefixType{},
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "The network this range belongs to, in IPv6 Address/CIDR format.",
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
		MarkdownDescription: "The name of the network view in which this range resides.",
	},
	"option_filter_rules": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: Ipv6rangeOptionFilterRulesResourceSchemaAttributes,
		},
		Optional: true,
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
		},
		MarkdownDescription: "This field contains the Option filters to be applied to this IPv6 range. The appliance uses the matching rules of these filters to select the address range from which it assigns a lease.",
	},
	"port_control_blackout_setting": schema.SingleNestedAttribute{
		Attributes:          Ipv6rangePortControlBlackoutSettingResourceSchemaAttributes,
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "",
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
	"same_port_control_discovery_blackout": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "If the field is set to True, the discovery blackout setting will be used for port control blackout setting.",
	},
	"server_association_type": schema.StringAttribute{
		Default:  stringdefault.StaticString("NONE"),
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "The type of server that is going to serve the range. Valid values are: * MEMBER * NONE",
	},
	"start_addr": schema.StringAttribute{
		Optional:   true,
		CustomType: iptypes.IPv6AddressType{},
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "The IPv6 Address starting address of the DHCP IPv6 range.",
	},
	"subscribe_settings": schema.SingleNestedAttribute{
		Attributes:          Ipv6rangeSubscribeSettingsResourceSchemaAttributes,
		Optional:            true,
		Computed:            true,
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
}

var Ipv6rangeResourceUddiSchemaAttributes = map[string]schema.Attribute{
	"comment": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			stringvalidator.LengthBetween(0, 1024),
		},
		MarkdownDescription: "The description for the range. May contain 0 to 1024 characters. Can include UTF-8.",
	},
	"dhcp_host": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The resource identifier.",
	},
	"dhcp_options": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: OptionItemResourceSchemaAttributes,
		},
		Optional: true,
		Computed: true,
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
		CustomType:          iptypes.IPv6AddressType{},
		MarkdownDescription: "The end IP address of the range.",
	},
	"exclusion_ranges": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: ExclusionRangeResourceSchemaAttributes,
		},
		Optional: true,
		Computed: true,
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
		Computed: true,
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
		},
		MarkdownDescription: "The list of all allow/deny filters of the range.",
	},
	"inheritance_sources": schema.SingleNestedAttribute{
		Attributes: DHCPOptionsInheritanceResourceSchemaAttributes,
		Optional:   true,
		Computed:   true,
		PlanModifiers: []planmodifier.Object{
			objectplanmodifier.UseStateForUnknown(),
		},
		MarkdownDescription: "The inheritance configuration that specifies how the _dhcp_options_ field is inherited from the parent object.",
	},
	"name": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			stringvalidator.LengthBetween(1, 256),
		},
		MarkdownDescription: "The name of the range. May contain 1 to 256 characters. Can include UTF-8.",
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
		CustomType:          iptypes.IPv6AddressType{},
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
		Computed:            true,
		MarkdownDescription: "A __UtilizationThreshold__ object represents IP address utilization threshold settings.",
	},
}

// Expand converts the TF model to the infoblox core model
func (m *Ipv6rangeModel) Expand(ctx context.Context, diags *diag.Diagnostics, isCreate bool) *coremodel.Ipv6range {
	if m == nil {
		return nil
	}

	obj := &coremodel.Ipv6range{}

	// Expand NIOS nested attribute (returns nil if not present)
	niosModel := flex.ExpandNestedObject[NIOSIpv6rangeModel](ctx, m.NIOS, diags)
	if niosModel != nil {
		obj.NIOS = niosModel.Expand(ctx, diags, isCreate)
	}

	// Expand UDDI nested attribute (returns nil if not present)
	uddiModel := flex.ExpandNestedObject[UDDIIpv6rangeModel](ctx, m.UDDI, diags)
	if uddiModel != nil {
		obj.UDDI = uddiModel.Expand(ctx, diags, isCreate)
	}

	return obj
}

// Expand converts the NIOS TF model to the core model.
func (m *NIOSIpv6rangeModel) Expand(ctx context.Context, diags *diag.Diagnostics, isCreate bool) *coremodel.NIOSIpv6rangeExt {
	ext := &coremodel.NIOSIpv6rangeExt{
		AddressType:                      flex.ExpandStringPointerNullAsEmpty(m.AddressType),
		CloudInfo:                        ExpandIpv6rangeCloudInfo(ctx, m.CloudInfo, diags),
		Comment:                          flex.ExpandStringPointerNullAsEmpty(m.Comment),
		Disable:                          flex.ExpandBoolPointer(m.Disable),
		DiscoveryBasicPollSettings:       ExpandIpv6rangeDiscoveryBasicPollSettings(ctx, m.DiscoveryBasicPollSettings, diags),
		DiscoveryBlackoutSetting:         ExpandIpv6rangeDiscoveryBlackoutSetting(ctx, m.DiscoveryBlackoutSetting, diags),
		DiscoveryMember:                  flex.ExpandStringPointer(m.DiscoveryMember),
		EnableDiscovery:                  flex.ExpandBoolPointer(m.EnableDiscovery),
		EnableImmediateDiscovery:         flex.ExpandBoolPointer(m.EnableImmediateDiscovery),
		EndAddr:                          flex.ExpandIPv6Address(m.EndAddr),
		Exclude:                          flex.ExpandFrameworkListNestedBlock(ctx, m.Exclude, diags, ExpandIpv6rangeExclude),
		ExtAttrs:                         flex.ExpandMapStringAny(ctx, m.ExtAttrs, diags),
		Ipv6EndPrefix:                    flex.ExpandIPv6Address(m.Ipv6EndPrefix),
		Ipv6PrefixBits:                   flex.ExpandInt64Pointer(m.Ipv6PrefixBits),
		Ipv6StartPrefix:                  flex.ExpandIPv6Address(m.Ipv6StartPrefix),
		LogicFilterRules:                 flex.ExpandFrameworkListNestedBlock(ctx, m.LogicFilterRules, diags, ExpandIpv6rangeLogicFilterRules),
		Member:                           ExpandIpv6rangeMember(ctx, m.Member, diags),
		Name:                             flex.ExpandStringPointerNullAsEmpty(m.Name),
		Network:                          flex.ExpandIPv6Prefix(m.Network),
		NetworkView:                      flex.ExpandStringPointerNullAsEmpty(m.NetworkView),
		OptionFilterRules:                flex.ExpandFrameworkListNestedBlock(ctx, m.OptionFilterRules, diags, ExpandIpv6rangeOptionFilterRules),
		PortControlBlackoutSetting:       ExpandIpv6rangePortControlBlackoutSetting(ctx, m.PortControlBlackoutSetting, diags),
		RecycleLeases:                    flex.ExpandBoolPointer(m.RecycleLeases),
		RestartIfNeeded:                  flex.ExpandBoolPointer(m.RestartIfNeeded),
		SamePortControlDiscoveryBlackout: flex.ExpandBoolPointer(m.SamePortControlDiscoveryBlackout),
		ServerAssociationType:            flex.ExpandStringPointerNullAsEmpty(m.ServerAssociationType),
		StartAddr:                        flex.ExpandIPv6Address(m.StartAddr),
		SubscribeSettings:                ExpandIpv6rangeSubscribeSettings(ctx, m.SubscribeSettings, diags),
	}
	if isCreate {
		ext.Template = flex.ExpandStringPointer(m.Template)
	}
	return ext
}

// ApplyIpv6rangeNIOSUseFlags derives NIOS use flags from the raw config
// value(s) and writes them onto the core model. A flag is true when the user
// set any of its governed value fields in config.
func ApplyIpv6rangeNIOSUseFlags(ctx context.Context, config tfsdk.Config, obj *coremodel.Ipv6range, diags *diag.Diagnostics) {
	if obj == nil || obj.NIOS == nil {
		return
	}
	obj.NIOS.UseBlackoutSetting = flex.DeriveUseFlag(ctx, config, diags, path.Root("nios").AtName("discovery_blackout_setting"), path.Root("nios").AtName("port_control_blackout_setting"), path.Root("nios").AtName("same_port_control_discovery_blackout"))
	obj.NIOS.UseDiscoveryBasicPollingSettings = flex.DeriveUseFlag(ctx, config, diags, path.Root("nios").AtName("discovery_basic_poll_settings"))
	obj.NIOS.UseEnableDiscovery = flex.DeriveUseFlag(ctx, config, diags, path.Root("nios").AtName("discovery_member"), path.Root("nios").AtName("enable_discovery"))
	obj.NIOS.UseLogicFilterRules = flex.DeriveUseFlag(ctx, config, diags, path.Root("nios").AtName("logic_filter_rules"))
	obj.NIOS.UseRecycleLeases = flex.DeriveUseFlag(ctx, config, diags, path.Root("nios").AtName("recycle_leases"))
	obj.NIOS.UseSubscribeSettings = flex.DeriveUseFlag(ctx, config, diags, path.Root("nios").AtName("subscribe_settings"))
}

// Expand converts the UDDI TF model to the core model.
func (m *UDDIIpv6rangeModel) Expand(ctx context.Context, diags *diag.Diagnostics, isCreate bool) *coremodel.UDDIIpv6rangeExt {
	ext := &coremodel.UDDIIpv6rangeExt{
		Comment:            flex.ExpandStringPointer(m.Comment),
		DhcpHost:           flex.ExpandStringPointer(m.DhcpHost),
		DhcpOptions:        flex.ExpandFrameworkListNestedBlock(ctx, m.DhcpOptions, diags, ExpandOptionItem),
		DisableDhcp:        flex.ExpandBoolPointer(m.DisableDhcp),
		End:                flex.ExpandIPv6AddressValue(m.End),
		ExclusionRanges:    flex.ExpandFrameworkListNestedBlock(ctx, m.ExclusionRanges, diags, ExpandExclusionRange),
		Filters:            flex.ExpandFrameworkListNestedBlock(ctx, m.Filters, diags, ExpandAccessFilter),
		InheritanceSources: ExpandDHCPOptionsInheritance(ctx, m.InheritanceSources, diags),
		Name:               flex.ExpandStringPointer(m.Name),
		Start:              flex.ExpandIPv6AddressValue(m.Start),
		Tags:               flex.ExpandMapStringAny(ctx, m.Tags, diags),
		Threshold:          ExpandUtilizationThreshold(ctx, m.Threshold, diags),
	}
	if isCreate {
		ext.Space = flex.ExpandStringPointer(m.Space)
	}
	return ext
}

// Flatten populates the TF model from a core response.
func (m *Ipv6rangeModel) Flatten(ctx context.Context, resp *coremodel.Ipv6range, diags *diag.Diagnostics) {
	if resp == nil {
		return
	}

	m.Id = flex.FlattenStringPointer(resp.Id)

	// Extract existing NIOS model, flatten API response onto it, convert back
	niosModel := flex.ExpandNestedObject[NIOSIpv6rangeModel](ctx, m.NIOS, diags)
	if niosModel == nil {
		niosModel = &NIOSIpv6rangeModel{}
	}
	niosModel.Flatten(ctx, resp.NIOS, diags)
	if resp.NIOS != nil {
		m.NIOS = flex.FlattenNestedObject(ctx, niosModel, NIOSIpv6rangeAttrTypes, diags)
	} else {
		m.NIOS = types.ObjectNull(NIOSIpv6rangeAttrTypes)
	}

	// Extract existing UDDI model, flatten API response onto it, convert back
	uddiModel := flex.ExpandNestedObject[UDDIIpv6rangeModel](ctx, m.UDDI, diags)
	if uddiModel == nil {
		uddiModel = &UDDIIpv6rangeModel{}
	}
	uddiModel.Flatten(ctx, resp.UDDI, diags)
	if resp.UDDI != nil {
		m.UDDI = flex.FlattenNestedObject(ctx, uddiModel, UDDIIpv6rangeAttrTypes, diags)
	} else {
		m.UDDI = types.ObjectNull(UDDIIpv6rangeAttrTypes)
	}
}

// Flatten merges API response onto existing NIOS model.
func (m *NIOSIpv6rangeModel) Flatten(ctx context.Context, from *coremodel.NIOSIpv6rangeExt, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	planExtAttrs := m.ExtAttrs
	if planExtAttrs.IsUnknown() {
		planExtAttrs = types.MapNull(types.StringType)
	}
	m.AddressType = flex.FlattenStringPointerEmptyAsNull(from.AddressType)
	m.CloudInfo = FlattenIpv6rangeCloudInfo(ctx, from.CloudInfo, diags)
	m.Comment = flex.FlattenStringPointerEmptyAsNull(from.Comment)
	m.Disable = flex.FlattenBoolPointer(from.Disable)
	m.DiscoveryBasicPollSettings = FlattenIpv6rangeDiscoveryBasicPollSettings(ctx, from.DiscoveryBasicPollSettings, diags)
	m.DiscoveryBlackoutSetting = FlattenIpv6rangeDiscoveryBlackoutSetting(ctx, from.DiscoveryBlackoutSetting, diags)
	m.DiscoveryMember = flex.FlattenStringPointerEmptyAsNull(from.DiscoveryMember)
	m.EnableDiscovery = flex.FlattenBoolPointer(from.EnableDiscovery)
	m.EndAddr = flex.FlattenIPv6Address(from.EndAddr)
	m.Exclude = flex.FlattenFrameworkListNestedBlock(ctx, from.Exclude, Ipv6rangeExcludeAttrTypes, diags, FlattenIpv6rangeExclude)
	m.ExtAttrs, m.ExtAttrsAll = flex.FlattenEAs(planExtAttrs, from.ExtAttrs)
	m.Ipv6EndPrefix = flex.FlattenIPv6Address(from.Ipv6EndPrefix)
	m.Ipv6PrefixBits = flex.FlattenInt64Pointer(from.Ipv6PrefixBits)
	m.Ipv6StartPrefix = flex.FlattenIPv6Address(from.Ipv6StartPrefix)
	m.LogicFilterRules = flex.FlattenFrameworkListNestedBlock(ctx, from.LogicFilterRules, Ipv6rangeLogicFilterRulesAttrTypes, diags, FlattenIpv6rangeLogicFilterRules)
	m.Member = FlattenIpv6rangeMember(ctx, from.Member, diags)
	m.Name = flex.FlattenStringPointerEmptyAsNull(from.Name)
	m.Network = flex.FlattenIPv6Prefix(from.Network)
	m.NetworkView = flex.FlattenStringPointerEmptyAsNull(from.NetworkView)
	m.OptionFilterRules = flex.FlattenFrameworkListNestedBlock(ctx, from.OptionFilterRules, Ipv6rangeOptionFilterRulesAttrTypes, diags, FlattenIpv6rangeOptionFilterRules)
	m.PortControlBlackoutSetting = FlattenIpv6rangePortControlBlackoutSetting(ctx, from.PortControlBlackoutSetting, diags)
	m.RecycleLeases = flex.FlattenBoolPointer(from.RecycleLeases)
	m.SamePortControlDiscoveryBlackout = flex.FlattenBoolPointer(from.SamePortControlDiscoveryBlackout)
	m.ServerAssociationType = flex.FlattenStringPointerEmptyAsNull(from.ServerAssociationType)
	m.StartAddr = flex.FlattenIPv6Address(from.StartAddr)
	m.SubscribeSettings = FlattenIpv6rangeSubscribeSettings(ctx, from.SubscribeSettings, diags)
}

// Flatten merges API response onto existing UDDI model.
func (m *UDDIIpv6rangeModel) Flatten(ctx context.Context, from *coremodel.UDDIIpv6rangeExt, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.Comment = flex.FlattenStringPointer(from.Comment)
	m.DhcpHost = flex.FlattenStringPointer(from.DhcpHost)
	m.DhcpOptions = flex.FlattenFrameworkListNestedBlock(ctx, from.DhcpOptions, OptionItemAttrTypes, diags, FlattenOptionItem)
	m.DisableDhcp = flex.FlattenBoolPointer(from.DisableDhcp)
	m.End = flex.FlattenIPv6AddressValue(from.End)
	m.ExclusionRanges = flex.FlattenFrameworkListNestedBlock(ctx, from.ExclusionRanges, ExclusionRangeAttrTypes, diags, FlattenExclusionRange)
	m.Filters = flex.FlattenFrameworkListNestedBlock(ctx, from.Filters, AccessFilterAttrTypes, diags, FlattenAccessFilter)
	m.InheritanceSources = FlattenDHCPOptionsInheritance(ctx, from.InheritanceSources, diags)
	m.Name = flex.FlattenStringPointer(from.Name)
	m.Space = flex.FlattenStringPointer(from.Space)
	m.Start = flex.FlattenIPv6AddressValue(from.Start)
	tagsAll := flex.FlattenMapStringAny(ctx, from.Tags, diags)
	if m.Tags.IsNull() || m.Tags.IsUnknown() {
		m.Tags = tagsAll
	}
	m.TagsAll = tagsAll
	m.Threshold = FlattenUtilizationThreshold(ctx, from.Threshold, diags)
}
