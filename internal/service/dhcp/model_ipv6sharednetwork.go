package dhcp

import (
	"context"

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
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/mapdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
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

type Ipv6sharednetworkModel struct {
	Id   types.String `tfsdk:"id"`
	NIOS types.Object `tfsdk:"nios"`
}

var Ipv6sharednetworkAttrTypes = map[string]attr.Type{
	"id":   types.StringType,
	"nios": types.ObjectType{AttrTypes: NIOSIpv6sharednetworkAttrTypes},
}

type NIOSIpv6sharednetworkModel struct {
	Comment                 types.String                        `tfsdk:"comment"`
	DdnsDomainname          internaltypes.CaseInsensitiveString `tfsdk:"ddns_domainname"`
	DdnsGenerateHostname    types.Bool                          `tfsdk:"ddns_generate_hostname"`
	DdnsServerAlwaysUpdates types.Bool                          `tfsdk:"ddns_server_always_updates"`
	DdnsTtl                 types.Int64                         `tfsdk:"ddns_ttl"`
	DdnsUseOption81         types.Bool                          `tfsdk:"ddns_use_option81"`
	Disable                 types.Bool                          `tfsdk:"disable"`
	DomainName              internaltypes.CaseInsensitiveString `tfsdk:"domain_name"`
	DomainNameServers       types.List                          `tfsdk:"domain_name_servers"`
	EnableDdns              types.Bool                          `tfsdk:"enable_ddns"`
	ExtAttrs                types.Map                           `tfsdk:"ext_attrs"`
	ExtAttrsAll             types.Map                           `tfsdk:"ext_attrs_all"`
	LogicFilterRules        types.List                          `tfsdk:"logic_filter_rules"`
	Name                    types.String                        `tfsdk:"name"`
	NetworkView             types.String                        `tfsdk:"network_view"`
	Networks                internaltypes.UnorderedListValue    `tfsdk:"networks"`
	Options                 types.List                          `tfsdk:"options"`
	PreferredLifetime       types.Int64                         `tfsdk:"preferred_lifetime"`
	UpdateDnsOnLeaseRenewal types.Bool                          `tfsdk:"update_dns_on_lease_renewal"`
	ValidLifetime           types.Int64                         `tfsdk:"valid_lifetime"`
}

var NIOSIpv6sharednetworkAttrTypes = map[string]attr.Type{
	"comment":                     types.StringType,
	"ddns_domainname":             internaltypes.CaseInsensitiveStringType{},
	"ddns_generate_hostname":      types.BoolType,
	"ddns_server_always_updates":  types.BoolType,
	"ddns_ttl":                    types.Int64Type,
	"ddns_use_option81":           types.BoolType,
	"disable":                     types.BoolType,
	"domain_name":                 internaltypes.CaseInsensitiveStringType{},
	"domain_name_servers":         types.ListType{ElemType: types.StringType},
	"enable_ddns":                 types.BoolType,
	"ext_attrs":                   types.MapType{ElemType: types.StringType},
	"ext_attrs_all":               types.MapType{ElemType: types.StringType},
	"logic_filter_rules":          types.ListType{ElemType: types.ObjectType{AttrTypes: Ipv6sharednetworkLogicFilterRulesAttrTypes}},
	"name":                        types.StringType,
	"network_view":                types.StringType,
	"networks":                    internaltypes.UnorderedListOfStringType,
	"options":                     types.ListType{ElemType: types.ObjectType{AttrTypes: Ipv6sharednetworkOptionsAttrTypes}},
	"preferred_lifetime":          types.Int64Type,
	"update_dns_on_lease_renewal": types.BoolType,
	"valid_lifetime":              types.Int64Type,
}

const (
	Ipv6sharednetworkReturnFields = "comment,ddns_domainname,ddns_generate_hostname,ddns_server_always_updates,ddns_ttl,ddns_use_option81,disable,domain_name,domain_name_servers,enable_ddns,extattrs,logic_filter_rules,name,network_view,networks,options,preferred_lifetime,update_dns_on_lease_renewal,use_ddns_domainname,use_ddns_generate_hostname,use_ddns_ttl,use_ddns_use_option81,use_domain_name,use_domain_name_servers,use_enable_ddns,use_logic_filter_rules,use_options,use_preferred_lifetime,use_update_dns_on_lease_renewal,use_valid_lifetime,valid_lifetime"
)

var Ipv6sharednetworkResourceSchemaAttributes = map[string]schema.Attribute{
	"id": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The reference to the object.",
	},
	"nios": schema.SingleNestedAttribute{
		Optional:            true,
		MarkdownDescription: "NIOS backend-specific fields.",
		Attributes:          Ipv6sharednetworkResourceNiosSchemaAttributes,
	},
}

var Ipv6sharednetworkResourceNiosSchemaAttributes = map[string]schema.Attribute{
	"comment": schema.StringAttribute{
		Optional: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
			stringvalidator.LengthBetween(0, 256),
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "Comment for the IPv6 shared network, maximum 256 characters.",
	},
	"ddns_domainname": schema.StringAttribute{
		Optional:   true,
		CustomType: internaltypes.CaseInsensitiveStringType{},
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
			customvalidator.IsValidNIOSDomainName(),
		},
		MarkdownDescription: "The dynamic DNS domain name the appliance uses specifically for DDNS updates for this network.",
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
		MarkdownDescription: "This field controls whether only the DHCP server is allowed to update DNS, regardless of the DHCP clients requests. Note that changes for this field take effect only if ddns_use_option81 is True.",
	},
	"ddns_ttl": schema.Int64Attribute{
		Optional:            true,
		Computed:            true,
		Default:             int64default.StaticInt64(0),
		MarkdownDescription: "The DNS update Time to Live (TTL) value of an IPv6 shared network object. The TTL is a 32-bit unsigned integer that represents the duration, in seconds, for which the update is cached. Zero indicates that the update is not cached.",
	},
	"ddns_use_option81": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "The support for DHCP Option 81 at the IPv6 shared network level.",
	},
	"disable": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Determines whether an IPv6 shared network is disabled or not. When this is set to False, the IPv6 shared network is enabled.",
	},
	"domain_name": schema.StringAttribute{
		Optional:   true,
		CustomType: internaltypes.CaseInsensitiveStringType{},
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
			customvalidator.IsValidNIOSDomainName(),
		},
		MarkdownDescription: "Use this method to set or retrieve the domain_name value of a DHCP IPv6 Shared Network object.",
	},
	"domain_name_servers": schema.ListAttribute{
		ElementType: types.StringType,
		Optional:    true,
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
			listvalidator.ValueStringsAre(customvalidator.IsValidIPv6Address()),
		},
		MarkdownDescription: "Use this method to set or retrieve the dynamic DNS updates flag of a DHCP IPv6 Shared Network object. The DHCP server can send DDNS updates to DNS servers in the same Grid and to external DNS servers. This setting overrides the member level settings.",
	},
	"enable_ddns": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "The dynamic DNS updates flag of an IPv6 shared network object. If set to True, the DHCP server sends DDNS updates to DNS servers in the same Grid, and to external DNS servers.",
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
	"logic_filter_rules": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: Ipv6sharednetworkLogicFilterRulesResourceSchemaAttributes,
		},
		Optional: true,
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
		},
		MarkdownDescription: "This field contains the logic filters to be applied on the this IPv6 shared network. This list corresponds to the match rules that are written to the DHCPv6 configuration file.",
	},
	"name": schema.StringAttribute{
		Required: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "The name of the IPv6 Shared Network.",
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
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "The name of the network view in which this IPv6 shared network resides.",
	},
	"networks": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: Ipv6sharednetworkNetworksResourceSchemaAttributes,
		},
		CustomType: internaltypes.UnorderedListOfStringType,
		Required:   true,
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
		},
		MarkdownDescription: "A list of IPv6 networks belonging to the shared network Each individual list item must be specified as an object containing a '_ref' parameter to a network reference, for example:: [{ \"_ref\": \"ipv6network/ZG5zdHdvcmskMTAuAvMTYvMA\", }] if the reference of the wanted network is not known, it is possible to specify search parameters for the network instead in the following way:: [{ \"_ref\": { 'network': 'aabb::/64', } }] note that in this case the search must match exactly one network for the assignment to be successful.",
	},
	"options": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: Ipv6sharednetworkOptionsResourceSchemaAttributes,
		},
		Optional: true,
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
		},
		MarkdownDescription: "An array of DHCP option dhcpoption structs that lists the DHCP options associated with the object.",
	},
	"preferred_lifetime": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "Use this method to set or retrieve the preferred lifetime value of a DHCP IPv6 Shared Network object.",
	},
	"update_dns_on_lease_renewal": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "This field controls whether the DHCP server updates DNS when a DHCP lease is renewed.",
	},
	"valid_lifetime": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "Use this method to set or retrieve the valid lifetime value of a DHCP IPv6 Shared Network object.",
	},
}

// Expand converts the TF model to the infoblox core model
func (m *Ipv6sharednetworkModel) Expand(ctx context.Context, diags *diag.Diagnostics, isCreate bool) *coremodel.Ipv6sharednetwork {
	if m == nil {
		return nil
	}

	obj := &coremodel.Ipv6sharednetwork{}

	// Expand NIOS nested attribute (returns nil if not present)
	niosModel := flex.ExpandNestedObject[NIOSIpv6sharednetworkModel](ctx, m.NIOS, diags)
	if niosModel != nil {
		obj.NIOS = niosModel.Expand(ctx, diags)
	}

	return obj
}

// Expand converts the NIOS TF model to the core model.
func (m *NIOSIpv6sharednetworkModel) Expand(ctx context.Context, diags *diag.Diagnostics) *coremodel.NIOSIpv6sharednetworkExt {
	return &coremodel.NIOSIpv6sharednetworkExt{
		Comment:                 flex.ExpandStringPointerNullAsEmpty(m.Comment),
		DdnsDomainname:          flex.ExpandStringPointer(m.DdnsDomainname.StringValue),
		DdnsGenerateHostname:    flex.ExpandBoolPointer(m.DdnsGenerateHostname),
		DdnsServerAlwaysUpdates: flex.ExpandBoolPointer(m.DdnsServerAlwaysUpdates),
		DdnsTtl:                 flex.ExpandInt64Pointer(m.DdnsTtl),
		DdnsUseOption81:         flex.ExpandBoolPointer(m.DdnsUseOption81),
		Disable:                 flex.ExpandBoolPointer(m.Disable),
		DomainName:              flex.ExpandStringPointer(m.DomainName.StringValue),
		DomainNameServers:       flex.ExpandFrameworkListString(ctx, m.DomainNameServers, diags),
		EnableDdns:              flex.ExpandBoolPointer(m.EnableDdns),
		ExtAttrs:                flex.ExpandMapStringAny(ctx, m.ExtAttrs, diags),
		LogicFilterRules:        flex.ExpandFrameworkListNestedBlock(ctx, m.LogicFilterRules, diags, ExpandIpv6sharednetworkLogicFilterRules),
		Name:                    flex.ExpandStringPointerNullAsEmpty(m.Name),
		NetworkView:             flex.ExpandStringPointerNullAsEmpty(m.NetworkView),
		Networks:                flex.ExpandFrameworkListNestedBlock(ctx, m.Networks, diags, ExpandIpv6sharednetworkNetworks),
		Options:                 flex.ExpandFrameworkListNestedBlock(ctx, m.Options, diags, ExpandIpv6sharednetworkOptions),
		PreferredLifetime:       flex.ExpandInt64Pointer(m.PreferredLifetime),
		UpdateDnsOnLeaseRenewal: flex.ExpandBoolPointer(m.UpdateDnsOnLeaseRenewal),
		ValidLifetime:           flex.ExpandInt64Pointer(m.ValidLifetime),
	}
}

// ApplyIpv6sharednetworkNIOSUseFlags derives NIOS use flags from the raw config
// value(s) and writes them onto the core model. A flag is true when the user
// set any of its governed value fields in config.
func ApplyIpv6sharednetworkNIOSUseFlags(ctx context.Context, config tfsdk.Config, obj *coremodel.Ipv6sharednetwork, diags *diag.Diagnostics) {
	if obj == nil || obj.NIOS == nil {
		return
	}
	obj.NIOS.UseDdnsDomainname = flex.DeriveUseFlag(ctx, config, diags, path.Root("nios").AtName("ddns_domainname"))
	obj.NIOS.UseDdnsGenerateHostname = flex.DeriveUseFlag(ctx, config, diags, path.Root("nios").AtName("ddns_generate_hostname"))
	obj.NIOS.UseDdnsTtl = flex.DeriveUseFlag(ctx, config, diags, path.Root("nios").AtName("ddns_ttl"))
	obj.NIOS.UseDdnsUseOption81 = flex.DeriveUseFlag(ctx, config, diags, path.Root("nios").AtName("ddns_use_option81"))
	obj.NIOS.UseDomainName = flex.DeriveUseFlag(ctx, config, diags, path.Root("nios").AtName("domain_name"))
	obj.NIOS.UseDomainNameServers = flex.DeriveUseFlag(ctx, config, diags, path.Root("nios").AtName("domain_name_servers"))
	obj.NIOS.UseEnableDdns = flex.DeriveUseFlag(ctx, config, diags, path.Root("nios").AtName("enable_ddns"))
	obj.NIOS.UseLogicFilterRules = flex.DeriveUseFlag(ctx, config, diags, path.Root("nios").AtName("logic_filter_rules"))
	obj.NIOS.UseOptions = flex.DeriveUseFlag(ctx, config, diags, path.Root("nios").AtName("options"))
	obj.NIOS.UsePreferredLifetime = flex.DeriveUseFlag(ctx, config, diags, path.Root("nios").AtName("preferred_lifetime"))
	obj.NIOS.UseUpdateDnsOnLeaseRenewal = flex.DeriveUseFlag(ctx, config, diags, path.Root("nios").AtName("update_dns_on_lease_renewal"))
	obj.NIOS.UseValidLifetime = flex.DeriveUseFlag(ctx, config, diags, path.Root("nios").AtName("valid_lifetime"))
}

// Flatten populates the TF model from a core response.
func (m *Ipv6sharednetworkModel) Flatten(ctx context.Context, resp *coremodel.Ipv6sharednetwork, diags *diag.Diagnostics) {
	if resp == nil {
		return
	}

	m.Id = flex.FlattenStringPointer(resp.Id)

	// Extract existing NIOS model, flatten API response onto it, convert back
	niosModel := flex.ExpandNestedObject[NIOSIpv6sharednetworkModel](ctx, m.NIOS, diags)
	if niosModel == nil {
		niosModel = &NIOSIpv6sharednetworkModel{}
	}
	niosModel.Flatten(ctx, resp.NIOS, diags)
	if resp.NIOS != nil {
		m.NIOS = flex.FlattenNestedObject(ctx, niosModel, NIOSIpv6sharednetworkAttrTypes, diags)
	} else {
		m.NIOS = types.ObjectNull(NIOSIpv6sharednetworkAttrTypes)
	}

}

// Flatten merges API response onto existing NIOS model.
func (m *NIOSIpv6sharednetworkModel) Flatten(ctx context.Context, from *coremodel.NIOSIpv6sharednetworkExt, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	planExtAttrs := m.ExtAttrs
	if planExtAttrs.IsUnknown() {
		planExtAttrs = types.MapNull(types.StringType)
	}
	m.Comment = flex.FlattenStringPointerEmptyAsNull(from.Comment)
	m.DdnsDomainname.StringValue = flex.FlattenStringPointer(from.DdnsDomainname)
	m.DdnsGenerateHostname = flex.FlattenBoolPointer(from.DdnsGenerateHostname)
	m.DdnsServerAlwaysUpdates = flex.FlattenBoolPointer(from.DdnsServerAlwaysUpdates)
	m.DdnsTtl = flex.FlattenInt64Pointer(from.DdnsTtl)
	m.DdnsUseOption81 = flex.FlattenBoolPointer(from.DdnsUseOption81)
	m.Disable = flex.FlattenBoolPointer(from.Disable)
	m.DomainName.StringValue = flex.FlattenStringPointer(from.DomainName)
	m.DomainNameServers = flex.FlattenFrameworkListString(ctx, from.DomainNameServers, diags)
	m.EnableDdns = flex.FlattenBoolPointer(from.EnableDdns)
	m.ExtAttrs, m.ExtAttrsAll = flex.FlattenEAs(planExtAttrs, from.ExtAttrs)
	m.LogicFilterRules = flex.FlattenFrameworkListNestedBlock(ctx, from.LogicFilterRules, Ipv6sharednetworkLogicFilterRulesAttrTypes, diags, FlattenIpv6sharednetworkLogicFilterRules)
	m.Name = flex.FlattenStringPointerEmptyAsNull(from.Name)
	m.NetworkView = flex.FlattenStringPointerEmptyAsNull(from.NetworkView)
	m.Networks = flex.FlattenFrameworkListNestedBlock(ctx, from.Networks, Ipv6sharednetworkNetworksAttrTypes, diags, FlattenIpv6sharednetworkNetworks)
	m.Options = flex.FlattenFrameworkListNestedBlock(ctx, from.Options, Ipv6sharednetworkOptionsAttrTypes, diags, FlattenIpv6sharednetworkOptions)
	m.PreferredLifetime = flex.FlattenInt64Pointer(from.PreferredLifetime)
	m.UpdateDnsOnLeaseRenewal = flex.FlattenBoolPointer(from.UpdateDnsOnLeaseRenewal)
	m.ValidLifetime = flex.FlattenInt64Pointer(from.ValidLifetime)
}
