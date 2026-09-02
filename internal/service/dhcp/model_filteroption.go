package dhcp

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
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

type FilteroptionModel struct {
	Id   types.String `tfsdk:"id"`
	NIOS types.Object `tfsdk:"nios"`
	UDDI types.Object `tfsdk:"uddi"`
}

var FilteroptionAttrTypes = map[string]attr.Type{
	"id":   types.StringType,
	"nios": types.ObjectType{AttrTypes: NIOSFilteroptionAttrTypes},
	"uddi": types.ObjectType{AttrTypes: UDDIFilteroptionAttrTypes},
}

type NIOSFilteroptionModel struct {
	ApplyAsClass types.Bool   `tfsdk:"apply_as_class"`
	Bootfile     types.String `tfsdk:"bootfile"`
	Bootserver   types.String `tfsdk:"bootserver"`
	Comment      types.String `tfsdk:"comment"`
	Expression   types.String `tfsdk:"expression"`
	ExtAttrs     types.Map    `tfsdk:"ext_attrs"`
	ExtAttrsAll  types.Map    `tfsdk:"ext_attrs_all"`
	LeaseTime    types.Int64  `tfsdk:"lease_time"`
	Name         types.String `tfsdk:"name"`
	NextServer   types.String `tfsdk:"next_server"`
	OptionList   types.List   `tfsdk:"option_list"`
	OptionSpace  types.String `tfsdk:"option_space"`
	PxeLeaseTime types.Int64  `tfsdk:"pxe_lease_time"`
}

var NIOSFilteroptionAttrTypes = map[string]attr.Type{
	"apply_as_class": types.BoolType,
	"bootfile":       types.StringType,
	"bootserver":     types.StringType,
	"comment":        types.StringType,
	"expression":     types.StringType,
	"ext_attrs":      types.MapType{ElemType: types.StringType},
	"ext_attrs_all":  types.MapType{ElemType: types.StringType},
	"lease_time":     types.Int64Type,
	"name":           types.StringType,
	"next_server":    types.StringType,
	"option_list":    types.ListType{ElemType: types.ObjectType{AttrTypes: FilteroptionOptionListAttrTypes}},
	"option_space":   types.StringType,
	"pxe_lease_time": types.Int64Type,
}

type UDDIFilteroptionModel struct {
	Comment                         types.String `tfsdk:"comment"`
	DhcpOptions                     types.List   `tfsdk:"dhcp_options"`
	HeaderOptionFilename            types.String `tfsdk:"header_option_filename"`
	HeaderOptionServerAddress       types.String `tfsdk:"header_option_server_address"`
	HeaderOptionServerName          types.String `tfsdk:"header_option_server_name"`
	LeaseTime                       types.Int64  `tfsdk:"lease_time"`
	Name                            types.String `tfsdk:"name"`
	Protocol                        types.String `tfsdk:"protocol"`
	Role                            types.String `tfsdk:"role"`
	Rules                           types.Object `tfsdk:"rules"`
	Tags                            types.Map    `tfsdk:"tags"`
	TagsAll                         types.Map    `tfsdk:"tags_all"`
	VendorSpecificOptionOptionSpace types.String `tfsdk:"vendor_specific_option_option_space"`
}

var UDDIFilteroptionAttrTypes = map[string]attr.Type{
	"comment":                             types.StringType,
	"dhcp_options":                        types.ListType{ElemType: types.ObjectType{AttrTypes: OptionItemAttrTypes}},
	"header_option_filename":              types.StringType,
	"header_option_server_address":        types.StringType,
	"header_option_server_name":           types.StringType,
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
	FilteroptionReturnFields = "apply_as_class,bootfile,bootserver,comment,expression,extattrs,lease_time,name,next_server,option_list,option_space,pxe_lease_time"
)

var FilteroptionResourceSchemaAttributes = map[string]schema.Attribute{
	"id": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The reference to the object.",
	},
	"nios": schema.SingleNestedAttribute{
		Optional:            true,
		MarkdownDescription: "NIOS backend-specific fields.",
		Attributes:          FilteroptionResourceNiosSchemaAttributes,
	},
	"uddi": schema.SingleNestedAttribute{
		Optional:            true,
		MarkdownDescription: "UDDI backend-specific fields.",
		Attributes:          FilteroptionResourceUddiSchemaAttributes,
	},
}

var FilteroptionResourceNiosSchemaAttributes = map[string]schema.Attribute{
	"apply_as_class": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(true),
		MarkdownDescription: "Determines if apply as class is enabled or not. If this flag is set to \"true\" the filter is treated as global DHCP class, e.g it is written to dhcpd config file even if it is not present in any DHCP range.",
	},
	"bootfile": schema.StringAttribute{
		Optional: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "A name of boot file of a DHCP filter option object.",
	},
	"bootserver": schema.StringAttribute{
		Optional: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "Determines the boot server of a DHCP filter option object. You can specify the name and/or IP address of the boot server that host needs to boot.",
	},
	"comment": schema.StringAttribute{
		Optional: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
			stringvalidator.LengthBetween(0, 256),
		},
		MarkdownDescription: "The descriptive comment of a DHCP filter option object.",
	},
	"expression": schema.StringAttribute{
		Optional: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "The conditional expression of a DHCP filter option object.",
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
		MarkdownDescription: "Determines the lease time of a DHCP filter option object.",
	},
	"name": schema.StringAttribute{
		Required: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "The name of a DHCP option filter object.",
	},
	"next_server": schema.StringAttribute{
		Optional: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "Determines the next server of a DHCP filter option object. You can specify the name and/or IP address of the next server that the host needs to boot.",
	},
	"option_list": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: FilteroptionOptionListResourceSchemaAttributes,
		},
		Optional: true,
		Computed: true,
		Default:  listdefault.StaticValue(types.ListValueMust(types.ObjectType{AttrTypes: FilteroptionOptionListAttrTypes}, []attr.Value{})),
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
		},
		MarkdownDescription: "An array of DHCP option dhcpoption structs that lists the DHCP options associated with the object.",
	},
	"option_space": schema.StringAttribute{
		Default:  stringdefault.StaticString("DHCP"),
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "The option space of a DHCP filter option object.",
	},
	"pxe_lease_time": schema.Int64Attribute{
		Optional: true,
		Computed: true,
		Validators: []validator.Int64{
			int64validator.Between(0, 2147483647),
		},
		MarkdownDescription: "Determines the PXE (Preboot Execution Environment) lease time of a DHCP filter option object. To specify the duration of time it takes a host to connect to a boot server, such as a TFTP server, and download the file it needs to boot.",
	},
}

var FilteroptionResourceUddiSchemaAttributes = map[string]schema.Attribute{
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
		Default:             stringdefault.StaticString("ip4"),
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
func (m *FilteroptionModel) Expand(ctx context.Context, diags *diag.Diagnostics, isCreate bool) *coremodel.Filteroption {
	if m == nil {
		return nil
	}

	obj := &coremodel.Filteroption{}

	// Expand NIOS nested attribute (returns nil if not present)
	niosModel := flex.ExpandNestedObject[NIOSFilteroptionModel](ctx, m.NIOS, diags)
	if niosModel != nil {
		obj.NIOS = niosModel.Expand(ctx, diags)
	}

	// Expand UDDI nested attribute (returns nil if not present)
	uddiModel := flex.ExpandNestedObject[UDDIFilteroptionModel](ctx, m.UDDI, diags)
	if uddiModel != nil {
		obj.UDDI = uddiModel.Expand(ctx, diags, isCreate)
	}

	return obj
}

// Expand converts the NIOS TF model to the core model.
func (m *NIOSFilteroptionModel) Expand(ctx context.Context, diags *diag.Diagnostics) *coremodel.NIOSFilteroptionExt {
	return &coremodel.NIOSFilteroptionExt{
		ApplyAsClass: flex.ExpandBoolPointer(m.ApplyAsClass),
		Bootfile:     flex.ExpandStringPointerNullAsEmpty(m.Bootfile),
		Bootserver:   flex.ExpandStringPointerNullAsEmpty(m.Bootserver),
		Comment:      flex.ExpandStringPointerNullAsEmpty(m.Comment),
		Expression:   flex.ExpandStringPointerNullAsEmpty(m.Expression),
		ExtAttrs:     flex.ExpandMapStringAny(ctx, m.ExtAttrs, diags),
		LeaseTime:    flex.ExpandInt64Pointer(m.LeaseTime),
		Name:         flex.ExpandStringPointerNullAsEmpty(m.Name),
		NextServer:   flex.ExpandStringPointerNullAsEmpty(m.NextServer),
		OptionList:   flex.ExpandFrameworkListNestedBlock(ctx, m.OptionList, diags, ExpandFilteroptionOptionList),
		OptionSpace:  flex.ExpandStringPointerNullAsEmpty(m.OptionSpace),
		PxeLeaseTime: flex.ExpandInt64Pointer(m.PxeLeaseTime),
	}
}

// Expand converts the UDDI TF model to the core model.
func (m *UDDIFilteroptionModel) Expand(ctx context.Context, diags *diag.Diagnostics, isCreate bool) *coremodel.UDDIFilteroptionExt {
	ext := &coremodel.UDDIFilteroptionExt{
		Comment:                         flex.ExpandStringPointer(m.Comment),
		DhcpOptions:                     flex.ExpandFrameworkListNestedBlock(ctx, m.DhcpOptions, diags, ExpandOptionItem),
		HeaderOptionFilename:            flex.ExpandStringPointer(m.HeaderOptionFilename),
		HeaderOptionServerAddress:       flex.ExpandStringPointer(m.HeaderOptionServerAddress),
		HeaderOptionServerName:          flex.ExpandStringPointer(m.HeaderOptionServerName),
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
func (m *FilteroptionModel) Flatten(ctx context.Context, resp *coremodel.Filteroption, diags *diag.Diagnostics) {
	if resp == nil {
		return
	}

	m.Id = flex.FlattenStringPointer(resp.Id)

	// Extract existing NIOS model, flatten API response onto it, convert back
	niosModel := flex.ExpandNestedObject[NIOSFilteroptionModel](ctx, m.NIOS, diags)
	if niosModel == nil {
		niosModel = &NIOSFilteroptionModel{}
	}
	plannedNIOS := flex.ExpandNestedObject[NIOSFilteroptionModel](ctx, m.NIOS, diags)
	niosModel.Flatten(ctx, resp.NIOS, diags)
	if resp.NIOS != nil {
		PostFlattenFilteroptionNIOS(ctx, plannedNIOS, niosModel, diags)
		m.NIOS = flex.FlattenNestedObject(ctx, niosModel, NIOSFilteroptionAttrTypes, diags)
	} else {
		m.NIOS = types.ObjectNull(NIOSFilteroptionAttrTypes)
	}

	// Extract existing UDDI model, flatten API response onto it, convert back
	uddiModel := flex.ExpandNestedObject[UDDIFilteroptionModel](ctx, m.UDDI, diags)
	if uddiModel == nil {
		uddiModel = &UDDIFilteroptionModel{}
	}
	uddiModel.Flatten(ctx, resp.UDDI, diags)
	if resp.UDDI != nil {
		m.UDDI = flex.FlattenNestedObject(ctx, uddiModel, UDDIFilteroptionAttrTypes, diags)
	} else {
		m.UDDI = types.ObjectNull(UDDIFilteroptionAttrTypes)
	}
}

// Flatten merges API response onto existing NIOS model.
func (m *NIOSFilteroptionModel) Flatten(ctx context.Context, from *coremodel.NIOSFilteroptionExt, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	planExtAttrs := m.ExtAttrs
	if planExtAttrs.IsUnknown() {
		planExtAttrs = types.MapNull(types.StringType)
	}
	m.ApplyAsClass = flex.FlattenBoolPointer(from.ApplyAsClass)
	m.Bootfile = flex.FlattenStringPointerEmptyAsNull(from.Bootfile)
	m.Bootserver = flex.FlattenStringPointerEmptyAsNull(from.Bootserver)
	m.Comment = flex.FlattenStringPointerEmptyAsNull(from.Comment)
	m.Expression = flex.FlattenStringPointerEmptyAsNull(from.Expression)
	m.ExtAttrs, m.ExtAttrsAll = flex.FlattenEAs(planExtAttrs, from.ExtAttrs)
	m.LeaseTime = flex.FlattenInt64Pointer(from.LeaseTime)
	m.Name = flex.FlattenStringPointerEmptyAsNull(from.Name)
	m.NextServer = flex.FlattenStringPointerEmptyAsNull(from.NextServer)
	m.OptionList = flex.FlattenFrameworkListNestedBlock(ctx, from.OptionList, FilteroptionOptionListAttrTypes, diags, FlattenFilteroptionOptionList)
	m.OptionSpace = flex.FlattenStringPointerEmptyAsNull(from.OptionSpace)
	m.PxeLeaseTime = flex.FlattenInt64Pointer(from.PxeLeaseTime)
}

// Flatten merges API response onto existing UDDI model.
func (m *UDDIFilteroptionModel) Flatten(ctx context.Context, from *coremodel.UDDIFilteroptionExt, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.Comment = flex.FlattenStringPointer(from.Comment)
	m.DhcpOptions = flex.FlattenFrameworkListNestedBlock(ctx, from.DhcpOptions, OptionItemAttrTypes, diags, FlattenOptionItem)
	m.HeaderOptionFilename = flex.FlattenStringPointer(from.HeaderOptionFilename)
	m.HeaderOptionServerAddress = flex.FlattenStringPointer(from.HeaderOptionServerAddress)
	m.HeaderOptionServerName = flex.FlattenStringPointer(from.HeaderOptionServerName)
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
