package fw

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/mapvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/mapdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"

	coremodel "github.com/infobloxopen/terraform-provider-infoblox/internal/core/model/fw"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
	customvalidator "github.com/infobloxopen/terraform-provider-infoblox/internal/validator"
)

type ApplicationFilterModel struct {
	Id   types.Int32  `tfsdk:"id"`
	UDDI types.Object `tfsdk:"uddi"`
}

var ApplicationFilterAttrTypes = map[string]attr.Type{
	"id":   types.Int32Type,
	"uddi": types.ObjectType{AttrTypes: UDDIApplicationFilterAttrTypes},
}

type UDDIApplicationFilterModel struct {
	Criteria    types.List   `tfsdk:"criteria"`
	Description types.String `tfsdk:"description"`
	Name        types.String `tfsdk:"name"`
	Readonly    types.Bool   `tfsdk:"readonly"`
	Tags        types.Map    `tfsdk:"tags"`
	TagsAll     types.Map    `tfsdk:"tags_all"`
}

var UDDIApplicationFilterAttrTypes = map[string]attr.Type{
	"criteria":    types.ListType{ElemType: types.ObjectType{AttrTypes: ApplicationCriterionAttrTypes}},
	"description": types.StringType,
	"name":        types.StringType,
	"readonly":    types.BoolType,
	"tags":        types.MapType{ElemType: types.StringType},
	"tags_all":    types.MapType{ElemType: types.StringType},
}

const (
	ApplicationFilterReturnFields = ""
)

var ApplicationFilterResourceSchemaAttributes = map[string]schema.Attribute{
	"id": schema.Int32Attribute{
		Computed:            true,
		MarkdownDescription: "The Application Filter object identifier.",
	},
	"uddi": schema.SingleNestedAttribute{
		Optional:            true,
		MarkdownDescription: "UDDI backend-specific fields.",
		Attributes:          ApplicationFilterResourceUddiSchemaAttributes,
	},
}

var ApplicationFilterResourceUddiSchemaAttributes = map[string]schema.Attribute{
	"criteria": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: ApplicationCriterionResourceSchemaAttributes,
		},
		Required: true,
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
		},
		MarkdownDescription: "The array of key-value pairs specifying criteria for the search.",
	},
	"description": schema.StringAttribute{
		Default:             stringdefault.StaticString(""),
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The brief description for the application filter.",
	},
	"name": schema.StringAttribute{
		Required:            true,
		MarkdownDescription: "The name of the application filter.",
	},
	"readonly": schema.BoolAttribute{
		Computed:            true,
		MarkdownDescription: "True if it is a predefined application filter",
	},
	"tags": schema.MapAttribute{
		Optional:    true,
		Computed:    true,
		ElementType: types.StringType,
		Default:     mapdefault.StaticValue(types.MapNull(types.StringType)),
		Validators: []validator.Map{
			mapvalidator.SizeAtLeast(1),
		},
		MarkdownDescription: "Enables tag support for resource where tags attribute contains user-defined key value pairs",
	},
	"tags_all": schema.MapAttribute{
		Computed:            true,
		ElementType:         types.StringType,
		MarkdownDescription: "All tags including inherited values.",
	},
}

// Expand converts the TF model to the infoblox core model
func (m *ApplicationFilterModel) Expand(ctx context.Context, diags *diag.Diagnostics, isCreate bool) *coremodel.ApplicationFilter {
	if m == nil {
		return nil
	}

	obj := &coremodel.ApplicationFilter{}

	// Expand UDDI nested attribute (returns nil if not present)
	uddiModel := flex.ExpandNestedObject[UDDIApplicationFilterModel](ctx, m.UDDI, diags)
	if uddiModel != nil {
		obj.UDDI = uddiModel.Expand(ctx, diags)
	}

	return obj
}

// Expand converts the UDDI TF model to the core model.
func (m *UDDIApplicationFilterModel) Expand(ctx context.Context, diags *diag.Diagnostics) *coremodel.UDDIApplicationFilterExt {
	return &coremodel.UDDIApplicationFilterExt{
		Criteria:    flex.ExpandFrameworkListNestedBlock(ctx, m.Criteria, diags, ExpandApplicationCriterion),
		Description: flex.ExpandStringPointer(m.Description),
		Name:        flex.ExpandStringPointer(m.Name),
		Tags:        flex.ExpandMapStringAny(ctx, m.Tags, diags),
	}
}

// Flatten populates the TF model from a core response.
func (m *ApplicationFilterModel) Flatten(ctx context.Context, resp *coremodel.ApplicationFilter, diags *diag.Diagnostics) {
	if resp == nil {
		return
	}

	m.Id = flex.FlattenInt32Pointer(resp.Id)

	// Extract existing UDDI model, flatten API response onto it, convert back
	uddiModel := flex.ExpandNestedObject[UDDIApplicationFilterModel](ctx, m.UDDI, diags)
	if uddiModel == nil {
		uddiModel = &UDDIApplicationFilterModel{}
	}
	uddiModel.Flatten(ctx, resp.UDDI, diags)
	if resp.UDDI != nil {
		m.UDDI = flex.FlattenNestedObject(ctx, uddiModel, UDDIApplicationFilterAttrTypes, diags)
	} else {
		m.UDDI = types.ObjectNull(UDDIApplicationFilterAttrTypes)
	}
}

// Flatten merges API response onto existing UDDI model.
func (m *UDDIApplicationFilterModel) Flatten(ctx context.Context, from *coremodel.UDDIApplicationFilterExt, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.Criteria = flex.FlattenFrameworkListNestedBlock(ctx, from.Criteria, ApplicationCriterionAttrTypes, diags, FlattenApplicationCriterion)
	m.Description = flex.FlattenStringPointer(from.Description)
	m.Name = flex.FlattenStringPointer(from.Name)
	m.Readonly = flex.FlattenBoolPointer(from.Readonly)
	tagsAll := flex.FlattenMapStringAny(ctx, from.Tags, diags)
	if m.Tags.IsNull() || m.Tags.IsUnknown() {
		m.Tags = tagsAll
	}
	m.TagsAll = tagsAll
}
