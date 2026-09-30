package fw

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/mapvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	int32planmodifier "github.com/hashicorp/terraform-plugin-framework/resource/schema/int32planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/mapdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"

	coremodel "github.com/infobloxopen/terraform-provider-infoblox/internal/core/model/fw"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
	customvalidator "github.com/infobloxopen/terraform-provider-infoblox/internal/validator"
)

type CategoryFilterModel struct {
	Id            types.Int32  `tfsdk:"id"`
	UpdateTrigger types.String `tfsdk:"update_trigger"`
	UDDI          types.Object `tfsdk:"uddi"`
}

var CategoryFilterAttrTypes = map[string]attr.Type{
	"id":             types.Int32Type,
	"update_trigger": types.StringType,
	"uddi":           types.ObjectType{AttrTypes: UDDICategoryFilterAttrTypes},
}

type UDDICategoryFilterModel struct {
	Categories  types.List   `tfsdk:"categories"`
	Description types.String `tfsdk:"description"`
	Name        types.String `tfsdk:"name"`
	Tags        types.Map    `tfsdk:"tags"`
	TagsAll     types.Map    `tfsdk:"tags_all"`
}

var UDDICategoryFilterAttrTypes = map[string]attr.Type{
	"categories":  types.ListType{ElemType: types.StringType},
	"description": types.StringType,
	"name":        types.StringType,
	"tags":        types.MapType{ElemType: types.StringType},
	"tags_all":    types.MapType{ElemType: types.StringType},
}

const (
	CategoryFilterReturnFields = ""
)

var CategoryFilterResourceSchemaAttributes = map[string]schema.Attribute{
	"id": schema.Int32Attribute{
		Computed: true,
		PlanModifiers: []planmodifier.Int32{
			int32planmodifier.UseStateForUnknown(),
		},
		MarkdownDescription: "The Category Filter object identifier.",
	},
	"update_trigger": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "An arbitrary value used to trigger an update. Not sent to the API. Change it when Terraform reports no infrastructure changes.",
	},
	"uddi": schema.SingleNestedAttribute{
		Optional:            true,
		MarkdownDescription: "UDDI backend-specific fields.",
		Attributes:          CategoryFilterResourceUddiSchemaAttributes,
	},
}

var CategoryFilterResourceUddiSchemaAttributes = map[string]schema.Attribute{
	"categories": schema.ListAttribute{
		ElementType: types.StringType,
		Required:    true,
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
		},
		MarkdownDescription: "The list of content category names that falls into this category filter.",
	},
	"description": schema.StringAttribute{
		Default:             stringdefault.StaticString(""),
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The brief description for the category filter.",
	},
	"name": schema.StringAttribute{
		Required:            true,
		MarkdownDescription: "The name of the category filter.",
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
func (m *CategoryFilterModel) Expand(ctx context.Context, diags *diag.Diagnostics, isCreate bool) *coremodel.CategoryFilter {
	if m == nil {
		return nil
	}

	obj := &coremodel.CategoryFilter{}

	// Expand UDDI nested attribute (returns nil if not present)
	uddiModel := flex.ExpandNestedObject[UDDICategoryFilterModel](ctx, m.UDDI, diags)
	if uddiModel != nil {
		obj.UDDI = uddiModel.Expand(ctx, diags)
	}

	return obj
}

// Expand converts the UDDI TF model to the core model.
func (m *UDDICategoryFilterModel) Expand(ctx context.Context, diags *diag.Diagnostics) *coremodel.UDDICategoryFilterExt {
	return &coremodel.UDDICategoryFilterExt{
		Categories:  flex.ExpandFrameworkListString(ctx, m.Categories, diags),
		Description: flex.ExpandStringPointer(m.Description),
		Name:        flex.ExpandStringPointer(m.Name),
		Tags:        flex.ExpandMapStringAny(ctx, m.Tags, diags),
	}
}

// Flatten populates the TF model from a core response.
func (m *CategoryFilterModel) Flatten(ctx context.Context, resp *coremodel.CategoryFilter, diags *diag.Diagnostics) {
	if resp == nil {
		return
	}

	m.Id = flex.FlattenInt32Pointer(resp.Id)

	// Extract existing UDDI model, flatten API response onto it, convert back
	uddiModel := flex.ExpandNestedObject[UDDICategoryFilterModel](ctx, m.UDDI, diags)
	if uddiModel == nil {
		uddiModel = &UDDICategoryFilterModel{}
	}
	uddiModel.Flatten(ctx, resp.UDDI, diags)
	if resp.UDDI != nil {
		m.UDDI = flex.FlattenNestedObject(ctx, uddiModel, UDDICategoryFilterAttrTypes, diags)
	} else {
		m.UDDI = types.ObjectNull(UDDICategoryFilterAttrTypes)
	}
}

// Flatten merges API response onto existing UDDI model.
func (m *UDDICategoryFilterModel) Flatten(ctx context.Context, from *coremodel.UDDICategoryFilterExt, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.Categories = flex.FlattenFrameworkListString(ctx, from.Categories, diags)
	m.Description = flex.FlattenStringPointer(from.Description)
	m.Name = flex.FlattenStringPointer(from.Name)
	tagsAll := flex.FlattenMapStringAny(ctx, from.Tags, diags)
	if m.Tags.IsNull() || m.Tags.IsUnknown() {
		m.Tags = tagsAll
	}
	m.TagsAll = tagsAll
}
