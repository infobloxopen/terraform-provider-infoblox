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
	stringplanmodifier "github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"

	coremodel "github.com/infobloxopen/terraform-provider-infoblox/internal/core/model/fw"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
	internaltypes "github.com/infobloxopen/terraform-provider-infoblox/internal/types"
	customvalidator "github.com/infobloxopen/terraform-provider-infoblox/internal/validator"
)

type InternalDomainListModel struct {
	Id            types.Int32  `tfsdk:"id"`
	UpdateTrigger types.String `tfsdk:"update_trigger"`
	UDDI          types.Object `tfsdk:"uddi"`
}

var InternalDomainListAttrTypes = map[string]attr.Type{
	"id":             types.Int32Type,
	"update_trigger": types.StringType,
	"uddi":           types.ObjectType{AttrTypes: UDDIInternalDomainListAttrTypes},
}

type UDDIInternalDomainListModel struct {
	Description     types.String                     `tfsdk:"description"`
	InternalDomains internaltypes.UnorderedListValue `tfsdk:"internal_domains"`
	IsDefault       types.Bool                       `tfsdk:"is_default"`
	Name            types.String                     `tfsdk:"name"`
	Tags            types.Map                        `tfsdk:"tags"`
	TagsAll         types.Map                        `tfsdk:"tags_all"`
}

var UDDIInternalDomainListAttrTypes = map[string]attr.Type{
	"description":      types.StringType,
	"internal_domains": internaltypes.UnorderedListOfStringType,
	"is_default":       types.BoolType,
	"name":             types.StringType,
	"tags":             types.MapType{ElemType: types.StringType},
	"tags_all":         types.MapType{ElemType: types.StringType},
}

const (
	InternalDomainListReturnFields = ""
)

var InternalDomainListResourceSchemaAttributes = map[string]schema.Attribute{
	"id": schema.Int32Attribute{
		Computed: true,
		PlanModifiers: []planmodifier.Int32{
			int32planmodifier.UseStateForUnknown(),
		},
		MarkdownDescription: "The Internal Domain object identifier.",
	},
	"update_trigger": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "An arbitrary value used to trigger an update. Not sent to the API. Change it when Terraform reports no infrastructure changes.",
	},
	"uddi": schema.SingleNestedAttribute{
		Optional:            true,
		MarkdownDescription: "UDDI backend-specific fields.",
		Attributes:          InternalDomainListResourceUddiSchemaAttributes,
	},
}

var InternalDomainListResourceUddiSchemaAttributes = map[string]schema.Attribute{
	"description": schema.StringAttribute{
		Default:             stringdefault.StaticString(""),
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The brief description for the internal domain lists .",
	},
	"internal_domains": schema.ListAttribute{
		ElementType: types.StringType,
		Required:    true,
		CustomType:  internaltypes.UnorderedListOfStringType,
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
		},
		MarkdownDescription: "The list of internal domains, should be unique to each other and has to be read-only from the API level.",
	},
	"is_default": schema.BoolAttribute{
		Computed:            true,
		MarkdownDescription: "True if name is 'Default Bypass Domains/CIDRs' otherwise false.",
	},
	"name": schema.StringAttribute{
		Optional: true,
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.RequiresReplaceIfConfigured(),
		},
		MarkdownDescription: "The name of the internal domain lists.",
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
func (m *InternalDomainListModel) Expand(ctx context.Context, diags *diag.Diagnostics, isCreate bool) *coremodel.InternalDomainList {
	if m == nil {
		return nil
	}

	obj := &coremodel.InternalDomainList{}

	// Expand UDDI nested attribute (returns nil if not present)
	uddiModel := flex.ExpandNestedObject[UDDIInternalDomainListModel](ctx, m.UDDI, diags)
	if uddiModel != nil {
		obj.UDDI = uddiModel.Expand(ctx, diags)
	}

	return obj
}

// Expand converts the UDDI TF model to the core model.
func (m *UDDIInternalDomainListModel) Expand(ctx context.Context, diags *diag.Diagnostics) *coremodel.UDDIInternalDomainListExt {
	return &coremodel.UDDIInternalDomainListExt{
		Description:     flex.ExpandStringPointer(m.Description),
		InternalDomains: flex.ExpandFrameworkListString(ctx, m.InternalDomains, diags),
		Name:            flex.ExpandStringPointer(m.Name),
		Tags:            flex.ExpandMapStringAny(ctx, m.Tags, diags),
	}
}

// Flatten populates the TF model from a core response.
func (m *InternalDomainListModel) Flatten(ctx context.Context, resp *coremodel.InternalDomainList, diags *diag.Diagnostics) {
	if resp == nil {
		return
	}

	m.Id = flex.FlattenInt32Pointer(resp.Id)

	// Extract existing UDDI model, flatten API response onto it, convert back
	uddiModel := flex.ExpandNestedObject[UDDIInternalDomainListModel](ctx, m.UDDI, diags)
	if uddiModel == nil {
		uddiModel = &UDDIInternalDomainListModel{}
	}
	uddiModel.Flatten(ctx, resp.UDDI, diags)
	if resp.UDDI != nil {
		m.UDDI = flex.FlattenNestedObject(ctx, uddiModel, UDDIInternalDomainListAttrTypes, diags)
	} else {
		m.UDDI = types.ObjectNull(UDDIInternalDomainListAttrTypes)
	}
}

// Flatten merges API response onto existing UDDI model.
func (m *UDDIInternalDomainListModel) Flatten(ctx context.Context, from *coremodel.UDDIInternalDomainListExt, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.Description = flex.FlattenStringPointer(from.Description)
	m.InternalDomains = flex.FlattenFrameworkUnorderedListString(ctx, from.InternalDomains, diags)
	m.IsDefault = flex.FlattenBoolPointer(from.IsDefault)
	m.Name = flex.FlattenStringPointer(from.Name)
	tagsAll := flex.FlattenMapStringAny(ctx, from.Tags, diags)
	if m.Tags.IsNull() || m.Tags.IsUnknown() {
		m.Tags = tagsAll
	}
	m.TagsAll = tagsAll
}
