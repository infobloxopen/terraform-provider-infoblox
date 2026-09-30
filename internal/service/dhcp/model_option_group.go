package dhcp

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/mapvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/mapdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"

	coremodel "github.com/infobloxopen/terraform-provider-infoblox/internal/core/model/dhcp"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
	immutable "github.com/infobloxopen/terraform-provider-infoblox/internal/planmodifiers/immutable"
	customvalidator "github.com/infobloxopen/terraform-provider-infoblox/internal/validator"
)

type OptionGroupModel struct {
	Id   types.String `tfsdk:"id"`
	UDDI types.Object `tfsdk:"uddi"`
}

var OptionGroupAttrTypes = map[string]attr.Type{
	"id":   types.StringType,
	"uddi": types.ObjectType{AttrTypes: UDDIOptionGroupAttrTypes},
}

type UDDIOptionGroupModel struct {
	Comment     types.String `tfsdk:"comment"`
	DhcpOptions types.List   `tfsdk:"dhcp_options"`
	Name        types.String `tfsdk:"name"`
	Protocol    types.String `tfsdk:"protocol"`
	Tags        types.Map    `tfsdk:"tags"`
	TagsAll     types.Map    `tfsdk:"tags_all"`
}

var UDDIOptionGroupAttrTypes = map[string]attr.Type{
	"comment":      types.StringType,
	"dhcp_options": types.ListType{ElemType: types.ObjectType{AttrTypes: OptionItemAttrTypes}},
	"name":         types.StringType,
	"protocol":     types.StringType,
	"tags":         types.MapType{ElemType: types.StringType},
	"tags_all":     types.MapType{ElemType: types.StringType},
}

const (
	OptionGroupReturnFields = ""
)

var OptionGroupResourceSchemaAttributes = map[string]schema.Attribute{
	"id": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The resource identifier.",
	},
	"uddi": schema.SingleNestedAttribute{
		Optional:            true,
		MarkdownDescription: "UDDI backend-specific fields.",
		Attributes:          OptionGroupResourceUddiSchemaAttributes,
	},
}

var OptionGroupResourceUddiSchemaAttributes = map[string]schema.Attribute{
	"comment": schema.StringAttribute{
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			stringvalidator.LengthBetween(0, 1024),
		},
		MarkdownDescription: "The description for the option group. May contain 0 to 1024 characters. Can include UTF-8.",
	},
	"dhcp_options": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: OptionItemResourceSchemaAttributes,
		},
		Optional: true,
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
		},
		MarkdownDescription: "The list of DHCP options for the option group. May be either a specific option or a group of options.",
	},
	"name": schema.StringAttribute{
		Required: true,
		Validators: []validator.String{
			stringvalidator.LengthBetween(1, 256),
		},
		MarkdownDescription: "The name of the option group. Must contain 1 to 256 characters. Can include UTF-8.",
	},
	"protocol": schema.StringAttribute{
		Validators: []validator.String{
			stringvalidator.OneOf("ip4", "ip6"),
		},
		Required: true,
		PlanModifiers: []planmodifier.String{
			immutable.ImmutableString(),
		},
		MarkdownDescription: "The type of protocol (_ip4_ or _ip6_).",
	},
	"tags": schema.MapAttribute{
		Optional:    true,
		Computed:    true,
		ElementType: types.StringType,
		Default:     mapdefault.StaticValue(types.MapNull(types.StringType)),
		Validators: []validator.Map{
			mapvalidator.SizeAtLeast(1),
		},
		MarkdownDescription: "The tags for the option group in JSON format.",
	},
	"tags_all": schema.MapAttribute{
		Computed:            true,
		ElementType:         types.StringType,
		MarkdownDescription: "All tags including inherited values.",
	},
}

// Expand converts the TF model to the infoblox core model
func (m *OptionGroupModel) Expand(ctx context.Context, diags *diag.Diagnostics, isCreate bool) *coremodel.OptionGroup {
	if m == nil {
		return nil
	}

	obj := &coremodel.OptionGroup{}

	// Expand UDDI nested attribute (returns nil if not present)
	uddiModel := flex.ExpandNestedObject[UDDIOptionGroupModel](ctx, m.UDDI, diags)
	if uddiModel != nil {
		obj.UDDI = uddiModel.Expand(ctx, diags, isCreate)
	}

	return obj
}

// Expand converts the UDDI TF model to the core model.
func (m *UDDIOptionGroupModel) Expand(ctx context.Context, diags *diag.Diagnostics, isCreate bool) *coremodel.UDDIOptionGroupExt {
	ext := &coremodel.UDDIOptionGroupExt{
		Comment:     flex.ExpandStringPointer(m.Comment),
		DhcpOptions: flex.ExpandFrameworkListNestedBlock(ctx, m.DhcpOptions, diags, ExpandOptionItem),
		Name:        flex.ExpandString(m.Name),
		Tags:        flex.ExpandMapStringAny(ctx, m.Tags, diags),
	}
	if isCreate {
		ext.Protocol = flex.ExpandStringPointer(m.Protocol)
	}
	return ext
}

// Flatten populates the TF model from a core response.
func (m *OptionGroupModel) Flatten(ctx context.Context, resp *coremodel.OptionGroup, diags *diag.Diagnostics) {
	if resp == nil {
		return
	}

	m.Id = flex.FlattenStringPointer(resp.Id)

	// Extract existing UDDI model, flatten API response onto it, convert back
	uddiModel := flex.ExpandNestedObject[UDDIOptionGroupModel](ctx, m.UDDI, diags)
	if uddiModel == nil {
		uddiModel = &UDDIOptionGroupModel{}
	}
	uddiModel.Flatten(ctx, resp.UDDI, diags)
	if resp.UDDI != nil {
		m.UDDI = flex.FlattenNestedObject(ctx, uddiModel, UDDIOptionGroupAttrTypes, diags)
	} else {
		m.UDDI = types.ObjectNull(UDDIOptionGroupAttrTypes)
	}
}

// Flatten merges API response onto existing UDDI model.
func (m *UDDIOptionGroupModel) Flatten(ctx context.Context, from *coremodel.UDDIOptionGroupExt, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.Comment = flex.FlattenStringPointer(from.Comment)
	m.DhcpOptions = flex.FlattenFrameworkListNestedBlock(ctx, from.DhcpOptions, OptionItemAttrTypes, diags, FlattenOptionItem)
	m.Name = flex.FlattenString(from.Name)
	m.Protocol = flex.FlattenStringPointer(from.Protocol)
	tagsAll := flex.FlattenMapStringAny(ctx, from.Tags, diags)
	if m.Tags.IsNull() || m.Tags.IsUnknown() {
		m.Tags = tagsAll
	}
	m.TagsAll = tagsAll
}
