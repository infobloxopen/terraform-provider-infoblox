package dns

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/hashicorp/terraform-plugin-framework-nettypes/iptypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listdefault"
	niosdns "github.com/infobloxopen/infoblox-nios-go-client/dns"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/dynamicallocation"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
	customvalidator "github.com/infobloxopen/terraform-provider-infoblox/internal/validator"
)

// RecordHostIpv6addrModel is the Terraform model for RecordHostIpv6addr
type RecordHostIpv6addrModel struct {
	AddressType          types.String        `tfsdk:"address_type"`
	ConfigureForDhcp     types.Bool          `tfsdk:"configure_for_dhcp"`
	DomainName           types.String        `tfsdk:"domain_name"`
	DomainNameServers    types.List          `tfsdk:"domain_name_servers"`
	Duid                 types.String        `tfsdk:"duid"`
	Ipv6addr             iptypes.IPv6Address `tfsdk:"ipv6addr"`
	Ipv6prefix           types.String        `tfsdk:"ipv6prefix"`
	Ipv6prefixBits       types.Int64         `tfsdk:"ipv6prefix_bits"`
	LogicFilterRules     types.List          `tfsdk:"logic_filter_rules"`
	Mac                  types.String        `tfsdk:"mac"`
	MatchClient          types.String        `tfsdk:"match_client"`
	Options              types.List          `tfsdk:"options"`
	PreferredLifetime    types.Int64         `tfsdk:"preferred_lifetime"`
	ReservedInterface    types.String        `tfsdk:"reserved_interface"`
	UseDomainName        types.Bool          `tfsdk:"use_domain_name"`
	UseDomainNameServers types.Bool          `tfsdk:"use_domain_name_servers"`
	UseForEaInheritance  types.Bool          `tfsdk:"use_for_ea_inheritance"`
	UseLogicFilterRules  types.Bool          `tfsdk:"use_logic_filter_rules"`
	UseOptions           types.Bool          `tfsdk:"use_options"`
	UsePreferredLifetime types.Bool          `tfsdk:"use_preferred_lifetime"`
	UseValidLifetime     types.Bool          `tfsdk:"use_valid_lifetime"`
	ValidLifetime        types.Int64         `tfsdk:"valid_lifetime"`
	DynamicAllocation    types.Object        `tfsdk:"dynamic_allocation"`
}

// RecordHostIpv6addrAttrTypes contains the attribute types for RecordHostIpv6addrModel
var RecordHostIpv6addrAttrTypes = map[string]attr.Type{
	"address_type":            types.StringType,
	"configure_for_dhcp":      types.BoolType,
	"domain_name":             types.StringType,
	"domain_name_servers":     types.ListType{ElemType: types.StringType},
	"duid":                    types.StringType,
	"ipv6addr":                iptypes.IPv6AddressType{},
	"ipv6prefix":              types.StringType,
	"ipv6prefix_bits":         types.Int64Type,
	"logic_filter_rules":      types.ListType{ElemType: types.ObjectType{AttrTypes: RecordHostIpv6addrLogicFilterRulesAttrTypes}},
	"mac":                     types.StringType,
	"match_client":            types.StringType,
	"options":                 types.ListType{ElemType: types.ObjectType{AttrTypes: RecordHostIpv6addrOptionsAttrTypes}},
	"preferred_lifetime":      types.Int64Type,
	"reserved_interface":      types.StringType,
	"use_domain_name":         types.BoolType,
	"use_domain_name_servers": types.BoolType,
	"use_for_ea_inheritance":  types.BoolType,
	"use_logic_filter_rules":  types.BoolType,
	"use_options":             types.BoolType,
	"use_preferred_lifetime":  types.BoolType,
	"use_valid_lifetime":      types.BoolType,
	"valid_lifetime":          types.Int64Type,
	"dynamic_allocation":      types.ObjectType{AttrTypes: dynamicallocation.NextAvailableIpAttrTypes},
}

// RecordHostIpv6addrResourceSchemaAttributes contains the schema attributes for RecordHostIpv6addrModel
var RecordHostIpv6addrResourceSchemaAttributes = map[string]schema.Attribute{
	"address_type": schema.StringAttribute{
		Validators: []validator.String{
			stringvalidator.OneOf("ADDRESS", "PREFIX", "BOTH"),
		},
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Type of the DHCP IPv6 Host Address object.",
	},
	"configure_for_dhcp": schema.BoolAttribute{
		Computed:            true,
		MarkdownDescription: "Set this to True to enable the DHCP configuration for this IPv6 host address.",
	},
	"domain_name": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "Use this method to set or retrieve the domain_name value of the DHCP IPv6 Host Address object.",
	},
	"domain_name_servers": schema.ListAttribute{
		ElementType: types.StringType,
		Optional:    true,
		Computed:    true,
		Default:     listdefault.StaticValue(types.ListNull(types.StringType)),
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
			listvalidator.AlsoRequires(path.MatchRelative().AtParent().AtName("use_domain_name_servers")),
			listvalidator.ValueStringsAre(customvalidator.IsValidIPv6Address()),
		},
		MarkdownDescription: "The IPv6 addresses of DNS recursive name servers to which the DHCP client can send name resolution requests. The DHCP server includes this information in the DNS Recursive Name Server option in Advertise, Rebind, Information-Request, and Reply messages.",
	},
	"duid": schema.StringAttribute{
		Computed: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "DHCPv6 Unique Identifier (DUID) of the address object.",
	},
	"ipv6addr": schema.StringAttribute{
		Optional:   true,
		Computed:   true,
		CustomType: iptypes.IPv6AddressType{},
		Validators: []validator.String{
			stringvalidator.ExactlyOneOf(
				path.MatchRelative().AtParent().AtName("dynamic_allocation"),
			),
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "The IPv6 Address of the record.",
	},
	"ipv6prefix": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "The IPv6 Address prefix of the DHCP IPv6 Host Address object.",
	},
	"ipv6prefix_bits": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "Prefix bits of the DHCP IPv6 Host Address object.",
	},
	"logic_filter_rules": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: RecordHostIpv6addrLogicFilterRulesResourceSchemaAttributes,
		},
		Optional: true,
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
		},
		MarkdownDescription: "This field contains the logic filters to be applied on the this host address. This list corresponds to the match rules that are written to the dhcpd configuration file.",
	},
	"mac": schema.StringAttribute{
		Computed: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "The MAC address for this host address.",
	},
	"match_client": schema.StringAttribute{
		Validators: []validator.String{
			stringvalidator.OneOf("DUID", "MAC_ADDRESS"),
		},
		Computed:            true,
		MarkdownDescription: "The match_client value for this fixed address. Valid values are: \"DUID\": The host IP address is leased to the matching DUID. \"MAC_ADDRESS\": The host IP address is leased to the matching MAC address.",
	},
	"options": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: RecordHostIpv6addrOptionsResourceSchemaAttributes,
		},
		Optional: true,
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
		},
		MarkdownDescription: "An array of DHCP option dhcpoption structs that lists the DHCP options associated with the object.",
	},
	"preferred_lifetime": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "Use this method to set or retrieve the preferred lifetime value of the DHCP IPv6 Host Address object.",
	},
	"reserved_interface": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "The reference to the reserved interface to which the device belongs.",
	},
	"use_domain_name": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Use flag for: domain_name",
	},
	"use_domain_name_servers": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Use flag for: domain_name_servers",
	},
	"use_for_ea_inheritance": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Set this to True when using this host address for EA inheritance.",
	},
	"use_logic_filter_rules": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Use flag for: logic_filter_rules",
	},
	"use_options": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Use flag for: options",
	},
	"use_preferred_lifetime": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Use flag for: preferred_lifetime",
	},
	"use_valid_lifetime": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "Use flag for: valid_lifetime",
	},
	"valid_lifetime": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "Use this method to set or retrieve the valid lifetime value of the DHCP IPv6 Host Address object.",
	},
	"dynamic_allocation": schema.SingleNestedAttribute{
		Attributes:          dynamicallocation.NextAvailableIpResourceSchemaAttributes,
		Optional:            true,
		MarkdownDescription: "Dynamically allocate the ip using the NIOS next_available_ip function call. Mutually exclusive with the static value field.",
	},
}

// ExpandRecordHostIpv6addr converts a Terraform Object to SDK type
func ExpandRecordHostIpv6addr(ctx context.Context, o types.Object, diags *diag.Diagnostics) *niosdns.RecordHostIpv6addr {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m RecordHostIpv6addrModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

// Expand converts the Terraform model to SDK type
func (m *RecordHostIpv6addrModel) Expand(ctx context.Context, diags *diag.Diagnostics) *niosdns.RecordHostIpv6addr {
	if m == nil {
		return nil
	}
	to := &niosdns.RecordHostIpv6addr{
		AddressType:          flex.ExpandStringPointer(m.AddressType),
		ConfigureForDhcp:     flex.ExpandBoolPointer(m.ConfigureForDhcp),
		DomainName:           flex.ExpandStringPointerNullAsEmpty(m.DomainName),
		DomainNameServers:    flex.ExpandFrameworkListString(ctx, m.DomainNameServers, diags),
		Duid:                 flex.ExpandStringPointer(m.Duid),
		Ipv6addr:             ExpandRecordHostIpv6addrIpv6addr(m.Ipv6addr),
		Ipv6prefix:           flex.ExpandStringPointer(m.Ipv6prefix),
		Ipv6prefixBits:       flex.ExpandInt64Pointer(m.Ipv6prefixBits),
		LogicFilterRules:     flex.ExpandFrameworkListNestedBlock(ctx, m.LogicFilterRules, diags, ExpandRecordHostIpv6addrLogicFilterRules),
		Mac:                  flex.ExpandStringPointer(m.Mac),
		MatchClient:          flex.ExpandStringPointer(m.MatchClient),
		Options:              flex.ExpandFrameworkListNestedBlock(ctx, m.Options, diags, ExpandRecordHostIpv6addrOptions),
		PreferredLifetime:    flex.ExpandInt64Pointer(m.PreferredLifetime),
		ReservedInterface:    flex.ExpandStringPointer(m.ReservedInterface),
		UseDomainName:        flex.ExpandBoolPointer(m.UseDomainName),
		UseDomainNameServers: flex.ExpandBoolPointer(m.UseDomainNameServers),
		UseForEaInheritance:  flex.ExpandBoolPointer(m.UseForEaInheritance),
		UseLogicFilterRules:  flex.ExpandBoolPointer(m.UseLogicFilterRules),
		UseOptions:           flex.ExpandBoolPointer(m.UseOptions),
		UsePreferredLifetime: flex.ExpandBoolPointer(m.UsePreferredLifetime),
		UseValidLifetime:     flex.ExpandBoolPointer(m.UseValidLifetime),
		ValidLifetime:        flex.ExpandInt64Pointer(m.ValidLifetime),
		FuncCall:             BuildRecordHostIpv6addrFuncCall(ctx, m.DynamicAllocation, diags),
	}
	return to
}

// FlattenRecordHostIpv6addr converts an SDK type to Terraform Object
func FlattenRecordHostIpv6addr(ctx context.Context, from *niosdns.RecordHostIpv6addr, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(RecordHostIpv6addrAttrTypes)
	}
	m := &RecordHostIpv6addrModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, RecordHostIpv6addrAttrTypes, m)
	diags.Append(d...)
	return t
}

// Flatten populates the Terraform model from SDK type
func (m *RecordHostIpv6addrModel) Flatten(ctx context.Context, from *niosdns.RecordHostIpv6addr, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.AddressType = flex.FlattenStringPointerEmptyAsNull(from.AddressType)
	m.ConfigureForDhcp = flex.FlattenBoolPointer(from.ConfigureForDhcp)
	m.DomainName = flex.FlattenStringPointerEmptyAsNull(from.DomainName)
	m.DomainNameServers = flex.FlattenFrameworkListString(ctx, from.DomainNameServers, diags)
	m.Duid = flex.FlattenStringPointerEmptyAsNull(from.Duid)
	m.Ipv6addr = FlattenRecordHostIpv6addrIpv6addr(from.Ipv6addr)
	m.Ipv6prefix = flex.FlattenStringPointerEmptyAsNull(from.Ipv6prefix)
	m.Ipv6prefixBits = flex.FlattenInt64Pointer(from.Ipv6prefixBits)
	m.LogicFilterRules = flex.FlattenFrameworkListNestedBlock(ctx, from.LogicFilterRules, RecordHostIpv6addrLogicFilterRulesAttrTypes, diags, FlattenRecordHostIpv6addrLogicFilterRules)
	m.Mac = flex.FlattenStringPointerEmptyAsNull(from.Mac)
	m.MatchClient = flex.FlattenStringPointerEmptyAsNull(from.MatchClient)
	m.Options = flex.FlattenFrameworkListNestedBlock(ctx, from.Options, RecordHostIpv6addrOptionsAttrTypes, diags, FlattenRecordHostIpv6addrOptions)
	m.PreferredLifetime = flex.FlattenInt64Pointer(from.PreferredLifetime)
	m.ReservedInterface = flex.FlattenStringPointerEmptyAsNull(from.ReservedInterface)
	m.UseDomainName = flex.FlattenBoolPointer(from.UseDomainName)
	m.UseDomainNameServers = flex.FlattenBoolPointer(from.UseDomainNameServers)
	m.UseForEaInheritance = flex.FlattenBoolPointer(from.UseForEaInheritance)
	m.UseLogicFilterRules = flex.FlattenBoolPointer(from.UseLogicFilterRules)
	m.UseOptions = flex.FlattenBoolPointer(from.UseOptions)
	m.UsePreferredLifetime = flex.FlattenBoolPointer(from.UsePreferredLifetime)
	m.UseValidLifetime = flex.FlattenBoolPointer(from.UseValidLifetime)
	m.ValidLifetime = flex.FlattenInt64Pointer(from.ValidLifetime)
	if len(m.DynamicAllocation.AttributeTypes(ctx)) == 0 {
		m.DynamicAllocation = types.ObjectNull(dynamicallocation.NextAvailableIpAttrTypes)
	}
}

func ExpandRecordHostIpv6addrIpv6addr(v iptypes.IPv6Address) *niosdns.RecordHostIpv6addrIpv6addr {
	return &niosdns.RecordHostIpv6addrIpv6addr{String: flex.ExpandIPv6Address(v)}
}

func FlattenRecordHostIpv6addrIpv6addr(v *niosdns.RecordHostIpv6addrIpv6addr) iptypes.IPv6Address {
	var value *string
	if v != nil {
		value = v.String
	}
	return flex.FlattenIPv6Address(value)
}
