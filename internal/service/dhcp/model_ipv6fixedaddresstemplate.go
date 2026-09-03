package dhcp

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/mapvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/mapdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"

	coremodel "github.com/infobloxopen/terraform-provider-infoblox/internal/core/model/dhcp"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
	importmod "github.com/infobloxopen/terraform-provider-infoblox/internal/planmodifiers/import"
	customvalidator "github.com/infobloxopen/terraform-provider-infoblox/internal/validator"
)

type Ipv6fixedaddresstemplateModel struct {
	Id   types.String `tfsdk:"id"`
	NIOS types.Object `tfsdk:"nios"`
}

var Ipv6fixedaddresstemplateAttrTypes = map[string]attr.Type{
	"id":   types.StringType,
	"nios": types.ObjectType{AttrTypes: NIOSIpv6fixedaddresstemplateAttrTypes},
}

type NIOSIpv6fixedaddresstemplateModel struct {
	Comment           types.String `tfsdk:"comment"`
	DomainName        types.String `tfsdk:"domain_name"`
	DomainNameServers types.List   `tfsdk:"domain_name_servers"`
	ExtAttrs          types.Map    `tfsdk:"ext_attrs"`
	ExtAttrsAll       types.Map    `tfsdk:"ext_attrs_all"`
	LogicFilterRules  types.List   `tfsdk:"logic_filter_rules"`
	Name              types.String `tfsdk:"name"`
	NumberOfAddresses types.Int64  `tfsdk:"number_of_addresses"`
	Offset            types.Int64  `tfsdk:"offset"`
	Options           types.List   `tfsdk:"options"`
	PreferredLifetime types.Int64  `tfsdk:"preferred_lifetime"`
	ValidLifetime     types.Int64  `tfsdk:"valid_lifetime"`
}

var NIOSIpv6fixedaddresstemplateAttrTypes = map[string]attr.Type{
	"comment":             types.StringType,
	"domain_name":         types.StringType,
	"domain_name_servers": types.ListType{ElemType: types.StringType},
	"ext_attrs":           types.MapType{ElemType: types.StringType},
	"ext_attrs_all":       types.MapType{ElemType: types.StringType},
	"logic_filter_rules":  types.ListType{ElemType: types.ObjectType{AttrTypes: Ipv6fixedaddresstemplateLogicFilterRulesAttrTypes}},
	"name":                types.StringType,
	"number_of_addresses": types.Int64Type,
	"offset":              types.Int64Type,
	"options":             types.ListType{ElemType: types.ObjectType{AttrTypes: Ipv6fixedaddresstemplateOptionsAttrTypes}},
	"preferred_lifetime":  types.Int64Type,
	"valid_lifetime":      types.Int64Type,
}

const (
	Ipv6fixedaddresstemplateReturnFields = "comment,domain_name,domain_name_servers,extattrs,logic_filter_rules,name,number_of_addresses,offset,options,preferred_lifetime,use_domain_name,use_domain_name_servers,use_logic_filter_rules,use_options,use_preferred_lifetime,use_valid_lifetime,valid_lifetime"
)

var Ipv6fixedaddresstemplateResourceSchemaAttributes = map[string]schema.Attribute{
	"id": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The reference to the object.",
	},
	"nios": schema.SingleNestedAttribute{
		Optional:            true,
		MarkdownDescription: "NIOS backend-specific fields.",
		Attributes:          Ipv6fixedaddresstemplateResourceNiosSchemaAttributes,
	},
}

var Ipv6fixedaddresstemplateResourceNiosSchemaAttributes = map[string]schema.Attribute{
	"comment": schema.StringAttribute{
		Optional: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
			stringvalidator.LengthBetween(0, 256),
		},
		MarkdownDescription: "A descriptive comment of an IPv6 fixed address template object.",
	},
	"domain_name": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
			customvalidator.IsValidNIOSDomainName(),
		},
		MarkdownDescription: "Domain name of the IPv6 fixed address template object.",
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
			Attributes: Ipv6fixedaddresstemplateLogicFilterRulesResourceSchemaAttributes,
		},
		Optional: true,
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
		},
		MarkdownDescription: "This field contains the logic filters to be applied to this IPv6 fixed address. This list corresponds to the match rules that are written to the DHCPv6 configuration file.",
	},
	"name": schema.StringAttribute{
		Required: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "Name of an IPv6 fixed address template object.",
	},
	"number_of_addresses": schema.Int64Attribute{
		Optional: true,
		Validators: []validator.Int64{
			int64validator.AlsoRequires(path.MatchRelative().AtParent().AtName("offset")),
		},
		MarkdownDescription: "The number of IPv6 addresses for this fixed address.",
	},
	"offset": schema.Int64Attribute{
		Optional: true,
		Validators: []validator.Int64{
			int64validator.AlsoRequires(path.MatchRelative().AtParent().AtName("number_of_addresses")),
		},
		MarkdownDescription: "The start address offset for this IPv6 fixed address.",
	},
	"options": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: Ipv6fixedaddresstemplateOptionsResourceSchemaAttributes,
		},
		Optional: true,
		Computed: true,
		Default:  listdefault.StaticValue(types.ListValueMust(types.ObjectType{AttrTypes: Ipv6fixedaddresstemplateOptionsAttrTypes}, []attr.Value{})),
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
		},
		MarkdownDescription: "An array of DHCP option dhcpoption structs that lists the DHCP options associated with the object.",
	},
	"preferred_lifetime": schema.Int64Attribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The preferred lifetime value for this DHCP IPv6 fixed address template object.",
	},
	"valid_lifetime": schema.Int64Attribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The valid lifetime value for this DHCP IPv6 fixed address template object.",
	},
}

// Expand converts the TF model to the infoblox core model
func (m *Ipv6fixedaddresstemplateModel) Expand(ctx context.Context, diags *diag.Diagnostics, isCreate bool) *coremodel.Ipv6fixedaddresstemplate {
	if m == nil {
		return nil
	}

	obj := &coremodel.Ipv6fixedaddresstemplate{}

	// Expand NIOS nested attribute (returns nil if not present)
	niosModel := flex.ExpandNestedObject[NIOSIpv6fixedaddresstemplateModel](ctx, m.NIOS, diags)
	if niosModel != nil {
		obj.NIOS = niosModel.Expand(ctx, diags)
	}

	return obj
}

// Expand converts the NIOS TF model to the core model.
func (m *NIOSIpv6fixedaddresstemplateModel) Expand(ctx context.Context, diags *diag.Diagnostics) *coremodel.NIOSIpv6fixedaddresstemplateExt {
	return &coremodel.NIOSIpv6fixedaddresstemplateExt{
		Comment:           flex.ExpandStringPointerNullAsEmpty(m.Comment),
		DomainName:        flex.ExpandStringPointerNullAsEmpty(m.DomainName),
		DomainNameServers: flex.ExpandFrameworkListString(ctx, m.DomainNameServers, diags),
		ExtAttrs:          flex.ExpandMapStringAny(ctx, m.ExtAttrs, diags),
		LogicFilterRules:  flex.ExpandFrameworkListNestedBlock(ctx, m.LogicFilterRules, diags, ExpandIpv6fixedaddresstemplateLogicFilterRules),
		Name:              flex.ExpandStringPointerNullAsEmpty(m.Name),
		NumberOfAddresses: flex.ExpandInt64Pointer(m.NumberOfAddresses),
		Offset:            flex.ExpandInt64Pointer(m.Offset),
		Options:           flex.ExpandFrameworkListNestedBlock(ctx, m.Options, diags, ExpandIpv6fixedaddresstemplateOptions),
		PreferredLifetime: flex.ExpandInt64Pointer(m.PreferredLifetime),
		ValidLifetime:     flex.ExpandInt64Pointer(m.ValidLifetime),
	}
}

// ApplyIpv6fixedaddresstemplateNIOSUseFlags derives NIOS use flags from the raw config
// value(s) and writes them onto the core model. A flag is true when the user
// set any of its governed value fields in config.
func ApplyIpv6fixedaddresstemplateNIOSUseFlags(ctx context.Context, config tfsdk.Config, obj *coremodel.Ipv6fixedaddresstemplate, diags *diag.Diagnostics) {
	if obj == nil || obj.NIOS == nil {
		return
	}
	obj.NIOS.UseDomainName = flex.DeriveUseFlag(ctx, config, diags, path.Root("nios").AtName("domain_name"))
	obj.NIOS.UseDomainNameServers = flex.DeriveUseFlag(ctx, config, diags, path.Root("nios").AtName("domain_name_servers"))
	obj.NIOS.UseLogicFilterRules = flex.DeriveUseFlag(ctx, config, diags, path.Root("nios").AtName("logic_filter_rules"))
	obj.NIOS.UseOptions = flex.DeriveUseFlag(ctx, config, diags, path.Root("nios").AtName("options"))
	obj.NIOS.UsePreferredLifetime = flex.DeriveUseFlag(ctx, config, diags, path.Root("nios").AtName("preferred_lifetime"))
	obj.NIOS.UseValidLifetime = flex.DeriveUseFlag(ctx, config, diags, path.Root("nios").AtName("valid_lifetime"))
}

// Flatten populates the TF model from a core response.
func (m *Ipv6fixedaddresstemplateModel) Flatten(ctx context.Context, resp *coremodel.Ipv6fixedaddresstemplate, diags *diag.Diagnostics) {
	if resp == nil {
		return
	}

	m.Id = flex.FlattenStringPointer(resp.Id)

	// Extract existing NIOS model, flatten API response onto it, convert back
	niosModel := flex.ExpandNestedObject[NIOSIpv6fixedaddresstemplateModel](ctx, m.NIOS, diags)
	if niosModel == nil {
		niosModel = &NIOSIpv6fixedaddresstemplateModel{}
	}
	plannedNIOS := flex.ExpandNestedObject[NIOSIpv6fixedaddresstemplateModel](ctx, m.NIOS, diags)
	niosModel.Flatten(ctx, resp.NIOS, diags)
	if resp.NIOS != nil {
		PostFlattenIpv6fixedaddresstemplateNIOS(ctx, plannedNIOS, niosModel, diags)
		m.NIOS = flex.FlattenNestedObject(ctx, niosModel, NIOSIpv6fixedaddresstemplateAttrTypes, diags)
	} else {
		m.NIOS = types.ObjectNull(NIOSIpv6fixedaddresstemplateAttrTypes)
	}

}

// Flatten merges API response onto existing NIOS model.
func (m *NIOSIpv6fixedaddresstemplateModel) Flatten(ctx context.Context, from *coremodel.NIOSIpv6fixedaddresstemplateExt, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	planExtAttrs := m.ExtAttrs
	if planExtAttrs.IsUnknown() {
		planExtAttrs = types.MapNull(types.StringType)
	}
	m.Comment = flex.FlattenStringPointerEmptyAsNull(from.Comment)
	m.DomainName = flex.FlattenStringPointerEmptyAsNull(from.DomainName)
	m.DomainNameServers = flex.FlattenFrameworkListString(ctx, from.DomainNameServers, diags)
	m.ExtAttrs, m.ExtAttrsAll = flex.FlattenEAs(planExtAttrs, from.ExtAttrs)
	m.LogicFilterRules = flex.FlattenFrameworkListNestedBlock(ctx, from.LogicFilterRules, Ipv6fixedaddresstemplateLogicFilterRulesAttrTypes, diags, FlattenIpv6fixedaddresstemplateLogicFilterRules)
	m.Name = flex.FlattenStringPointerEmptyAsNull(from.Name)
	m.NumberOfAddresses = flex.FlattenInt64Pointer(from.NumberOfAddresses)
	m.Offset = flex.FlattenInt64Pointer(from.Offset)
	m.Options = flex.FlattenFrameworkListNestedBlock(ctx, from.Options, Ipv6fixedaddresstemplateOptionsAttrTypes, diags, FlattenIpv6fixedaddresstemplateOptions)
	m.PreferredLifetime = flex.FlattenInt64Pointer(from.PreferredLifetime)
	m.ValidLifetime = flex.FlattenInt64Pointer(from.ValidLifetime)
}
