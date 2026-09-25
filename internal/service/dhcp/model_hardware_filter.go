package dhcp

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/mapvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/mapdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	stringplanmodifier "github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"

	coremodel "github.com/infobloxopen/terraform-provider-infoblox/internal/core/model/dhcp"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
	customvalidator "github.com/infobloxopen/terraform-provider-infoblox/internal/validator"
)

type HardwareFilterModel struct {
	Id   types.String `tfsdk:"id"`
	UDDI types.Object `tfsdk:"uddi"`
}

var HardwareFilterAttrTypes = map[string]attr.Type{
	"id":   types.StringType,
	"uddi": types.ObjectType{AttrTypes: UDDIHardwareFilterAttrTypes},
}

type UDDIHardwareFilterModel struct {
	Addresses                       types.List        `tfsdk:"addresses"`
	Comment                         types.String      `tfsdk:"comment"`
	CreatedAt                       timetypes.RFC3339 `tfsdk:"created_at"`
	DhcpOptions                     types.List        `tfsdk:"dhcp_options"`
	HeaderOptionFilename            types.String      `tfsdk:"header_option_filename"`
	HeaderOptionServerAddress       types.String      `tfsdk:"header_option_server_address"`
	HeaderOptionServerName          types.String      `tfsdk:"header_option_server_name"`
	LeaseTime                       types.Int64       `tfsdk:"lease_time"`
	Name                            types.String      `tfsdk:"name"`
	Role                            types.String      `tfsdk:"role"`
	Tags                            types.Map         `tfsdk:"tags"`
	TagsAll                         types.Map         `tfsdk:"tags_all"`
	UpdatedAt                       timetypes.RFC3339 `tfsdk:"updated_at"`
	VendorSpecificOptionOptionSpace types.String      `tfsdk:"vendor_specific_option_option_space"`
}

var UDDIHardwareFilterAttrTypes = map[string]attr.Type{
	"addresses":                           types.ListType{ElemType: types.StringType},
	"comment":                             types.StringType,
	"created_at":                          timetypes.RFC3339Type{},
	"dhcp_options":                        types.ListType{ElemType: types.ObjectType{AttrTypes: OptionItemAttrTypes}},
	"header_option_filename":              types.StringType,
	"header_option_server_address":        types.StringType,
	"header_option_server_name":           types.StringType,
	"lease_time":                          types.Int64Type,
	"name":                                types.StringType,
	"role":                                types.StringType,
	"tags":                                types.MapType{ElemType: types.StringType},
	"tags_all":                            types.MapType{ElemType: types.StringType},
	"updated_at":                          timetypes.RFC3339Type{},
	"vendor_specific_option_option_space": types.StringType,
}

const (
	HardwareFilterReturnFields = ""
)

var HardwareFilterResourceSchemaAttributes = map[string]schema.Attribute{
	"id": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The resource identifier.",
	},
	"uddi": schema.SingleNestedAttribute{
		Optional:            true,
		MarkdownDescription: "UDDI backend-specific fields.",
		Attributes:          HardwareFilterResourceUddiSchemaAttributes,
	},
}

var HardwareFilterResourceUddiSchemaAttributes = map[string]schema.Attribute{
	"addresses": schema.ListAttribute{
		ElementType: types.StringType,
		Optional:    true,
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
		},
		MarkdownDescription: "The list of addresses to match for the hardware filter.",
	},
	"comment": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			stringvalidator.LengthBetween(0, 1024),
		},
		MarkdownDescription: "The description for the hardware filter. May contain 0 to 1024 characters. Can include UTF-8.",
	},
	"created_at": schema.StringAttribute{
		Computed:            true,
		CustomType:          timetypes.RFC3339Type{},
		MarkdownDescription: "Time when the object has been created.",
	},
	"dhcp_options": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: OptionItemResourceSchemaAttributes,
		},
		Optional: true,
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
		},
		MarkdownDescription: "The list of DHCP options for the hardware filter. May be either a specific option or a group of options.",
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
		Required: true,
		Validators: []validator.String{
			stringvalidator.LengthBetween(1, 256),
		},
		MarkdownDescription: "The name of the hardware filter. Must contain 1 to 256 characters. Can include UTF-8.",
	},
	"role": schema.StringAttribute{
		Validators: []validator.String{
			stringvalidator.OneOf("values", "selection"),
		},
		Optional: true,
		Computed: true,
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.RequiresReplaceIfConfigured(),
		},
		MarkdownDescription: "The role of DHCP filter (_values_ or _selection_).  Defaults to _values_.",
	},
	"tags": schema.MapAttribute{
		Optional:    true,
		Computed:    true,
		ElementType: types.StringType,
		Default:     mapdefault.StaticValue(types.MapNull(types.StringType)),
		Validators: []validator.Map{
			mapvalidator.SizeAtLeast(1),
		},
		MarkdownDescription: "The tags for the hardware filter in JSON format.",
	},
	"tags_all": schema.MapAttribute{
		Computed:            true,
		ElementType:         types.StringType,
		MarkdownDescription: "All tags including inherited values.",
	},
	"updated_at": schema.StringAttribute{
		Computed:            true,
		CustomType:          timetypes.RFC3339Type{},
		MarkdownDescription: "Time when the object has been updated. Equals to _created_at_ if not updated after creation.",
	},
	"vendor_specific_option_option_space": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The resource identifier.",
	},
}

// Expand converts the TF model to the infoblox core model
func (m *HardwareFilterModel) Expand(ctx context.Context, diags *diag.Diagnostics, isCreate bool) *coremodel.HardwareFilter {
	if m == nil {
		return nil
	}

	obj := &coremodel.HardwareFilter{}

	// Expand UDDI nested attribute (returns nil if not present)
	uddiModel := flex.ExpandNestedObject[UDDIHardwareFilterModel](ctx, m.UDDI, diags)
	if uddiModel != nil {
		obj.UDDI = uddiModel.Expand(ctx, diags)
	}

	return obj
}

// Expand converts the UDDI TF model to the core model.
func (m *UDDIHardwareFilterModel) Expand(ctx context.Context, diags *diag.Diagnostics) *coremodel.UDDIHardwareFilterExt {
	return &coremodel.UDDIHardwareFilterExt{
		Addresses:                       flex.ExpandFrameworkListString(ctx, m.Addresses, diags),
		Comment:                         flex.ExpandStringPointer(m.Comment),
		DhcpOptions:                     flex.ExpandFrameworkListNestedBlock(ctx, m.DhcpOptions, diags, ExpandOptionItem),
		HeaderOptionFilename:            flex.ExpandStringPointer(m.HeaderOptionFilename),
		HeaderOptionServerAddress:       flex.ExpandStringPointer(m.HeaderOptionServerAddress),
		HeaderOptionServerName:          flex.ExpandStringPointer(m.HeaderOptionServerName),
		LeaseTime:                       flex.ExpandInt64Pointer(m.LeaseTime),
		Name:                            flex.ExpandString(m.Name),
		Role:                            flex.ExpandStringPointer(m.Role),
		Tags:                            flex.ExpandMapStringAny(ctx, m.Tags, diags),
		VendorSpecificOptionOptionSpace: flex.ExpandStringPointer(m.VendorSpecificOptionOptionSpace),
	}
}

// Flatten populates the TF model from a core response.
func (m *HardwareFilterModel) Flatten(ctx context.Context, resp *coremodel.HardwareFilter, diags *diag.Diagnostics) {
	if resp == nil {
		return
	}

	m.Id = flex.FlattenStringPointer(resp.Id)

	// Extract existing UDDI model, flatten API response onto it, convert back
	uddiModel := flex.ExpandNestedObject[UDDIHardwareFilterModel](ctx, m.UDDI, diags)
	if uddiModel == nil {
		uddiModel = &UDDIHardwareFilterModel{}
	}
	uddiModel.Flatten(ctx, resp.UDDI, diags)
	if resp.UDDI != nil {
		m.UDDI = flex.FlattenNestedObject(ctx, uddiModel, UDDIHardwareFilterAttrTypes, diags)
	} else {
		m.UDDI = types.ObjectNull(UDDIHardwareFilterAttrTypes)
	}
}

// Flatten merges API response onto existing UDDI model.
func (m *UDDIHardwareFilterModel) Flatten(ctx context.Context, from *coremodel.UDDIHardwareFilterExt, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.Addresses = flex.FlattenFrameworkListString(ctx, from.Addresses, diags)
	m.Comment = flex.FlattenStringPointer(from.Comment)
	m.CreatedAt = flex.FlattenRFC3339(from.CreatedAt)
	m.DhcpOptions = flex.FlattenFrameworkListNestedBlock(ctx, from.DhcpOptions, OptionItemAttrTypes, diags, FlattenOptionItem)
	m.HeaderOptionFilename = flex.FlattenStringPointer(from.HeaderOptionFilename)
	m.HeaderOptionServerAddress = flex.FlattenStringPointer(from.HeaderOptionServerAddress)
	m.HeaderOptionServerName = flex.FlattenStringPointer(from.HeaderOptionServerName)
	m.LeaseTime = flex.FlattenInt64Pointer(from.LeaseTime)
	m.Name = flex.FlattenString(from.Name)
	m.Role = flex.FlattenStringPointer(from.Role)
	tagsAll := flex.FlattenMapStringAny(ctx, from.Tags, diags)
	if m.Tags.IsNull() || m.Tags.IsUnknown() {
		m.Tags = tagsAll
	}
	m.TagsAll = tagsAll
	m.UpdatedAt = flex.FlattenRFC3339(from.UpdatedAt)
	m.VendorSpecificOptionOptionSpace = flex.FlattenStringPointer(from.VendorSpecificOptionOptionSpace)
}
