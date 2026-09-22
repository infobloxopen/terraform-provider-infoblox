package dhcp

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/mapvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/mapdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	stringplanmodifier "github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"

	coremodel "github.com/infobloxopen/terraform-provider-infoblox/internal/core/model/dhcp"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
	importmod "github.com/infobloxopen/terraform-provider-infoblox/internal/planmodifiers/import"
	customvalidator "github.com/infobloxopen/terraform-provider-infoblox/internal/validator"
)

type Ipv6filteroptionModel struct {
	Id   types.String `tfsdk:"id"`
	NIOS types.Object `tfsdk:"nios"`
	UDDI types.Object `tfsdk:"uddi"`
}

var Ipv6filteroptionAttrTypes = map[string]attr.Type{
	"id":   types.StringType,
	"nios": types.ObjectType{AttrTypes: NIOSIpv6filteroptionAttrTypes},
	"uddi": types.ObjectType{AttrTypes: UDDIIpv6filteroptionAttrTypes},
}

type NIOSIpv6filteroptionModel struct {
	ApplyAsClass types.Bool   `tfsdk:"apply_as_class"`
	Comment      types.String `tfsdk:"comment"`
	Expression   types.String `tfsdk:"expression"`
	ExtAttrs     types.Map    `tfsdk:"ext_attrs"`
	ExtAttrsAll  types.Map    `tfsdk:"ext_attrs_all"`
	LeaseTime    types.Int64  `tfsdk:"lease_time"`
	Name         types.String `tfsdk:"name"`
	OptionList   types.List   `tfsdk:"option_list"`
	OptionSpace  types.String `tfsdk:"option_space"`
}

var NIOSIpv6filteroptionAttrTypes = map[string]attr.Type{
	"apply_as_class": types.BoolType,
	"comment":        types.StringType,
	"expression":     types.StringType,
	"ext_attrs":      types.MapType{ElemType: types.StringType},
	"ext_attrs_all":  types.MapType{ElemType: types.StringType},
	"lease_time":     types.Int64Type,
	"name":           types.StringType,
	"option_list":    types.ListType{ElemType: types.ObjectType{AttrTypes: Ipv6filteroptionOptionListAttrTypes}},
	"option_space":   types.StringType,
}

type UDDIIpv6filteroptionModel struct {
	Comment                         types.String `tfsdk:"comment"`
	DhcpOptions                     types.List   `tfsdk:"dhcp_options"`
	LeaseTime                       types.Int64  `tfsdk:"lease_time"`
	Name                            types.String `tfsdk:"name"`
	Protocol                        types.String `tfsdk:"protocol"`
	Role                            types.String `tfsdk:"role"`
	Rules                           types.Object `tfsdk:"rules"`
	Tags                            types.Map    `tfsdk:"tags"`
	TagsAll                         types.Map    `tfsdk:"tags_all"`
	VendorSpecificOptionOptionSpace types.String `tfsdk:"vendor_specific_option_option_space"`
}

var UDDIIpv6filteroptionAttrTypes = map[string]attr.Type{
	"comment":                             types.StringType,
	"dhcp_options":                        types.ListType{ElemType: types.ObjectType{AttrTypes: OptionItemAttrTypes}},
	"lease_time":                          types.Int64Type,
	"name":                                types.StringType,
	"protocol":                            types.StringType,
	"role":                                types.StringType,
	"rules":                               types.ObjectType{AttrTypes: OptionFilterRuleListAttrTypes},
	"tags":                                types.MapType{ElemType: types.StringType},
	"tags_all":                            types.MapType{ElemType: types.StringType},
	"vendor_specific_option_option_space": types.StringType,
}

const (
	Ipv6filteroptionReturnFields = "apply_as_class,comment,expression,extattrs,lease_time,name,option_list,option_space"
)

var Ipv6filteroptionResourceSchemaAttributes = map[string]schema.Attribute{
	"id": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The reference to the object.",
	},
	"nios": schema.SingleNestedAttribute{
		Optional:            true,
		MarkdownDescription: "NIOS backend-specific fields.",
		Attributes:          Ipv6filteroptionResourceNiosSchemaAttributes,
	},
	"uddi": schema.SingleNestedAttribute{
		Optional:            true,
		MarkdownDescription: "UDDI backend-specific fields.",
		Attributes:          Ipv6filteroptionResourceUddiSchemaAttributes,
	},
}

var Ipv6filteroptionResourceNiosSchemaAttributes = map[string]schema.Attribute{
	"apply_as_class": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(true),
		MarkdownDescription: "Determines if apply as class is enabled or not. If this flag is set to \"true\" the filter is treated as global DHCP class, e.g it is written to DHCPv6 configuration file even if it is not present in any DHCP range.",
	},
	"comment": schema.StringAttribute{
		Optional: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
			stringvalidator.LengthBetween(0, 256),
		},
		MarkdownDescription: "The descriptive comment of a DHCP IPv6 filter option object.",
	},
	"expression": schema.StringAttribute{
		Optional: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "The conditional expression of a DHCP IPv6 filter option object.",
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
	"lease_time": schema.Int64Attribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Determines the lease time of a DHCP IPv6 filter option object.",
	},
	"name": schema.StringAttribute{
		Required: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "The name of a DHCP IPv6 option filter object.",
	},
	"option_list": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: Ipv6filteroptionOptionListResourceSchemaAttributes,
		},
		Optional: true,
		Computed: true,
		Default:  listdefault.StaticValue(types.ListNull(types.ObjectType{AttrTypes: Ipv6filteroptionOptionListAttrTypes})),
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
		},
		MarkdownDescription: "An array of DHCP option dhcpoption structs that lists the DHCP options associated with the object.",
	},
	"option_space": schema.StringAttribute{
		Default:  stringdefault.StaticString("DHCPv6"),
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "The option space of a DHCP IPv6 filter option object.",
	},
}

var Ipv6filteroptionResourceUddiSchemaAttributes = map[string]schema.Attribute{
	"comment": schema.StringAttribute{
		Default:             stringdefault.StaticString(""),
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The description for the option filter. May contain 0 to 1024 characters. Can include UTF-8.",
	},
	"dhcp_options": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: OptionItemResourceSchemaAttributes,
		},
		Optional: true,
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
		},
		MarkdownDescription: "The list of DHCP options for the option filter. May be either a specific option or a group of options.",
	},
	"lease_time": schema.Int64Attribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The lease lifetime duration in seconds.",
	},
	"name": schema.StringAttribute{
		Required:            true,
		MarkdownDescription: "The name of the option filter. Must contain 1 to 256 characters. Can include UTF-8.",
	},
	"protocol": schema.StringAttribute{
		Default:             stringdefault.StaticString("ip6"),
		Computed:            true,
		MarkdownDescription: "The type of protocol of option filter (_ip4_ or _ip6_).",
	},
	"role": schema.StringAttribute{
		Optional: true,
		Computed: true,
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.RequiresReplaceIfConfigured(),
		},
		MarkdownDescription: "The role of DHCP filter (_values_ or _selection_).  Defaults to _values_.",
	},
	"rules": schema.SingleNestedAttribute{
		Attributes:          OptionFilterRuleListResourceSchemaAttributes,
		Required:            true,
		MarkdownDescription: "An __OptionFilterRuleList__ object (_dhcp/option_filter_rule_list_) represents a collection of DHCP option filter rules that supports matching all or any rules.",
	},
	"tags": schema.MapAttribute{
		Optional:    true,
		Computed:    true,
		ElementType: types.StringType,
		Default:     mapdefault.StaticValue(types.MapNull(types.StringType)),
		Validators: []validator.Map{
			mapvalidator.SizeAtLeast(1),
		},
		MarkdownDescription: "The tags for the option filter in JSON format.",
	},
	"tags_all": schema.MapAttribute{
		Computed:            true,
		ElementType:         types.StringType,
		MarkdownDescription: "All tags including inherited values.",
	},
	"vendor_specific_option_option_space": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The resource identifier.",
	},
}

// Expand converts the TF model to the infoblox core model
func (m *Ipv6filteroptionModel) Expand(ctx context.Context, diags *diag.Diagnostics, isCreate bool) *coremodel.Ipv6filteroption {
	if m == nil {
		return nil
	}

	obj := &coremodel.Ipv6filteroption{}

	// Expand NIOS nested attribute (returns nil if not present)
	niosModel := flex.ExpandNestedObject[NIOSIpv6filteroptionModel](ctx, m.NIOS, diags)
	if niosModel != nil {
		obj.NIOS = niosModel.Expand(ctx, diags)
	}

	// Expand UDDI nested attribute (returns nil if not present)
	uddiModel := flex.ExpandNestedObject[UDDIIpv6filteroptionModel](ctx, m.UDDI, diags)
	if uddiModel != nil {
		obj.UDDI = uddiModel.Expand(ctx, diags, isCreate)
	}

	return obj
}

// Expand converts the NIOS TF model to the core model.
func (m *NIOSIpv6filteroptionModel) Expand(ctx context.Context, diags *diag.Diagnostics) *coremodel.NIOSIpv6filteroptionExt {
	return &coremodel.NIOSIpv6filteroptionExt{
		ApplyAsClass: flex.ExpandBoolPointer(m.ApplyAsClass),
		Comment:      flex.ExpandStringPointerNullAsEmpty(m.Comment),
		Expression:   flex.ExpandStringPointerNullAsEmpty(m.Expression),
		ExtAttrs:     flex.ExpandMapStringAny(ctx, m.ExtAttrs, diags),
		LeaseTime:    flex.ExpandInt64Pointer(m.LeaseTime),
		Name:         flex.ExpandStringPointerNullAsEmpty(m.Name),
		OptionList:   flex.ExpandFrameworkListNestedBlock(ctx, m.OptionList, diags, ExpandIpv6filteroptionOptionList),
		OptionSpace:  flex.ExpandStringPointerNullAsEmpty(m.OptionSpace),
	}
}

// Expand converts the UDDI TF model to the core model.
func (m *UDDIIpv6filteroptionModel) Expand(ctx context.Context, diags *diag.Diagnostics, isCreate bool) *coremodel.UDDIIpv6filteroptionExt {
	ext := &coremodel.UDDIIpv6filteroptionExt{
		Comment:                         flex.ExpandStringPointer(m.Comment),
		DhcpOptions:                     flex.ExpandFrameworkListNestedBlock(ctx, m.DhcpOptions, diags, ExpandOptionItem),
		LeaseTime:                       flex.ExpandInt64Pointer(m.LeaseTime),
		Name:                            flex.ExpandString(m.Name),
		Role:                            flex.ExpandStringPointer(m.Role),
		Rules:                           ExpandOptionFilterRuleList(ctx, m.Rules, diags),
		Tags:                            flex.ExpandMapStringAny(ctx, m.Tags, diags),
		VendorSpecificOptionOptionSpace: flex.ExpandStringPointer(m.VendorSpecificOptionOptionSpace),
	}
	if isCreate {
		ext.Protocol = flex.ExpandStringPointer(m.Protocol)
	}
	return ext
}

// Flatten populates the TF model from a core response.
func (m *Ipv6filteroptionModel) Flatten(ctx context.Context, resp *coremodel.Ipv6filteroption, diags *diag.Diagnostics) {
	if resp == nil {
		return
	}

	m.Id = flex.FlattenStringPointer(resp.Id)

	// Extract existing NIOS model, flatten API response onto it, convert back
	niosModel := flex.ExpandNestedObject[NIOSIpv6filteroptionModel](ctx, m.NIOS, diags)
	if niosModel == nil {
		niosModel = &NIOSIpv6filteroptionModel{}
	}
	plannedNIOS := flex.ExpandNestedObject[NIOSIpv6filteroptionModel](ctx, m.NIOS, diags)
	niosModel.Flatten(ctx, resp.NIOS, diags)
	if resp.NIOS != nil {
		PostFlattenIpv6filteroptionNIOS(ctx, plannedNIOS, niosModel, diags)
		m.NIOS = flex.FlattenNestedObject(ctx, niosModel, NIOSIpv6filteroptionAttrTypes, diags)
	} else {
		m.NIOS = types.ObjectNull(NIOSIpv6filteroptionAttrTypes)
	}

	// Extract existing UDDI model, flatten API response onto it, convert back
	uddiModel := flex.ExpandNestedObject[UDDIIpv6filteroptionModel](ctx, m.UDDI, diags)
	if uddiModel == nil {
		uddiModel = &UDDIIpv6filteroptionModel{}
	}
	uddiModel.Flatten(ctx, resp.UDDI, diags)
	if resp.UDDI != nil {
		m.UDDI = flex.FlattenNestedObject(ctx, uddiModel, UDDIIpv6filteroptionAttrTypes, diags)
	} else {
		m.UDDI = types.ObjectNull(UDDIIpv6filteroptionAttrTypes)
	}
}

// Flatten merges API response onto existing NIOS model.
func (m *NIOSIpv6filteroptionModel) Flatten(ctx context.Context, from *coremodel.NIOSIpv6filteroptionExt, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	planExtAttrs := m.ExtAttrs
	if planExtAttrs.IsUnknown() {
		planExtAttrs = types.MapNull(types.StringType)
	}
	m.ApplyAsClass = flex.FlattenBoolPointer(from.ApplyAsClass)
	m.Comment = flex.FlattenStringPointerEmptyAsNull(from.Comment)
	m.Expression = flex.FlattenStringPointerEmptyAsNull(from.Expression)
	m.ExtAttrs, m.ExtAttrsAll = flex.FlattenEAs(planExtAttrs, from.ExtAttrs)
	m.LeaseTime = flex.FlattenInt64Pointer(from.LeaseTime)
	m.Name = flex.FlattenStringPointerEmptyAsNull(from.Name)
	m.OptionList = flex.FlattenFrameworkListNestedBlock(ctx, from.OptionList, Ipv6filteroptionOptionListAttrTypes, diags, FlattenIpv6filteroptionOptionList)
	m.OptionSpace = flex.FlattenStringPointerEmptyAsNull(from.OptionSpace)
}

// Flatten merges API response onto existing UDDI model.
func (m *UDDIIpv6filteroptionModel) Flatten(ctx context.Context, from *coremodel.UDDIIpv6filteroptionExt, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.Comment = flex.FlattenStringPointer(from.Comment)
	m.DhcpOptions = flex.FlattenFrameworkListNestedBlock(ctx, from.DhcpOptions, OptionItemAttrTypes, diags, FlattenOptionItem)
	m.LeaseTime = flex.FlattenInt64Pointer(from.LeaseTime)
	m.Name = flex.FlattenString(from.Name)
	m.Protocol = flex.FlattenStringPointer(from.Protocol)
	m.Role = flex.FlattenStringPointer(from.Role)
	m.Rules = FlattenOptionFilterRuleList(ctx, from.Rules, diags)
	tagsAll := flex.FlattenMapStringAny(ctx, from.Tags, diags)
	if m.Tags.IsNull() || m.Tags.IsUnknown() {
		m.Tags = tagsAll
	}
	m.TagsAll = tagsAll
	m.VendorSpecificOptionOptionSpace = flex.FlattenStringPointer(from.VendorSpecificOptionOptionSpace)
}
