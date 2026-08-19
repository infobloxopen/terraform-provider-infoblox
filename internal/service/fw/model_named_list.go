package fw

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/mapvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/mapdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	stringplanmodifier "github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"

	coremodel "github.com/infobloxopen/terraform-provider-infoblox/internal/core/model/fw"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
)

type NamedListModel struct {
	Id   types.Int32  `tfsdk:"id"`
	UDDI types.Object `tfsdk:"uddi"`
}

var NamedListAttrTypes = map[string]attr.Type{
	"id":   types.Int32Type,
	"uddi": types.ObjectType{AttrTypes: UDDINamedListAttrTypes},
}

type UDDINamedListModel struct {
	ConfidenceLevel types.String `tfsdk:"confidence_level"`
	Description     types.String `tfsdk:"description"`
	Items           types.List   `tfsdk:"items"`
	ItemsDescribed  types.List   `tfsdk:"items_described"`
	Name            types.String `tfsdk:"name"`
	Policies        types.List   `tfsdk:"policies"`
	Tags            types.Map    `tfsdk:"tags"`
	TagsAll         types.Map    `tfsdk:"tags_all"`
	ThreatLevel     types.String `tfsdk:"threat_level"`
	Type            types.String `tfsdk:"type"`
}

var UDDINamedListAttrTypes = map[string]attr.Type{
	"confidence_level": types.StringType,
	"description":      types.StringType,
	"items":            types.ListType{ElemType: types.StringType},
	"items_described":  types.ListType{ElemType: types.ObjectType{AttrTypes: ItemStructsAttrTypes}},
	"name":             types.StringType,
	"policies":         types.ListType{ElemType: types.StringType},
	"tags":             types.MapType{ElemType: types.StringType},
	"tags_all":         types.MapType{ElemType: types.StringType},
	"threat_level":     types.StringType,
	"type":             types.StringType,
}

const (
	NamedListType         = "NamedList"
	NamedListReturnFields = ""
)

var NamedListResourceSchemaAttributes = map[string]schema.Attribute{
	"id": schema.Int32Attribute{
		Computed:            true,
		MarkdownDescription: "The Named List object identifier.",
	},
	"uddi": schema.SingleNestedAttribute{
		Optional:            true,
		MarkdownDescription: "UDDI backend-specific fields.",
		Attributes:          NamedListResourceUddiSchemaAttributes,
	},
}

var NamedListResourceUddiSchemaAttributes = map[string]schema.Attribute{
	"confidence_level": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The confidence level for a custom list. The possible values are [\"LOW\", \"MEDIUM\", \"HIGH\"]",
	},
	"description": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The brief description for the named list.",
	},
	"items": schema.ListAttribute{
		ElementType:         types.StringType,
		Optional:            true,
		MarkdownDescription: "The list of the FQDN or IPv4/IPv6 CIDRs to define whitelists and blacklists for additional protection.",
	},
	"items_described": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: ItemStructsResourceSchemaAttributes,
		},
		Optional:            true,
		MarkdownDescription: "The List of ItemStructs structure which contains the item and its description",
	},
	"name": schema.StringAttribute{
		Required:            true,
		MarkdownDescription: "The name of the named list.",
	},
	"policies": schema.ListAttribute{
		ElementType:         types.StringType,
		Optional:            true,
		MarkdownDescription: "The list of the security policy names with which the named list is associated.",
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
	"threat_level": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "The threat level for a custom list. The possible values are [\"INFO\", \"LOW\", \"MEDIUM\", \"HIGH\"]",
	},
	"type": schema.StringAttribute{
		Validators: []validator.String{
			stringvalidator.OneOf("custom_list", "threat_insight", "fast_flux", "dga", "dnsm", "threat_insight_nde", "default_allow", "default_block"),
		},
		Optional: true,
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.RequiresReplaceIfConfigured(),
		},
		MarkdownDescription: "The type of the named list, that can be \"custom_list\", \"threat_insight\", \"fast_flux\", \"dga\", \"dnsm\", \"threat_insight_nde\", \"default_allow\", \"default_block\" or \"threat_insight_nde\".",
	},
}

// Expand converts the TF model to the infoblox core model
func (m *NamedListModel) Expand(ctx context.Context, diags *diag.Diagnostics, isCreate bool) *coremodel.NamedList {
	if m == nil {
		return nil
	}

	obj := &coremodel.NamedList{}

	// Expand UDDI nested attribute (returns nil if not present)
	uddiModel := flex.ExpandNestedObject[UDDINamedListModel](ctx, m.UDDI, diags)
	if uddiModel != nil {
		obj.UDDI = uddiModel.Expand(ctx, diags)
	}

	return obj
}

// Expand converts the UDDI TF model to the core model.
func (m *UDDINamedListModel) Expand(ctx context.Context, diags *diag.Diagnostics) *coremodel.UDDINamedListExt {
	return &coremodel.UDDINamedListExt{
		ConfidenceLevel: flex.ExpandStringPointer(m.ConfidenceLevel),
		Description:     flex.ExpandStringPointer(m.Description),
		Items:           flex.ExpandFrameworkListString(ctx, m.Items, diags),
		ItemsDescribed:  flex.ExpandFrameworkListNestedBlock(ctx, m.ItemsDescribed, diags, ExpandItemStructs),
		Name:            flex.ExpandStringPointer(m.Name),
		Policies:        flex.ExpandFrameworkListString(ctx, m.Policies, diags),
		Tags:            flex.ExpandMapStringAny(ctx, m.Tags, diags),
		ThreatLevel:     flex.ExpandStringPointer(m.ThreatLevel),
		Type:            flex.ExpandStringPointer(m.Type),
	}
}

// Flatten populates the TF model from a core response.
func (m *NamedListModel) Flatten(ctx context.Context, resp *coremodel.NamedList, diags *diag.Diagnostics) {
	if resp == nil {
		return
	}

	m.Id = flex.FlattenInt32Pointer(resp.Id)

	// Extract existing UDDI model, flatten API response onto it, convert back
	uddiModel := flex.ExpandNestedObject[UDDINamedListModel](ctx, m.UDDI, diags)
	if uddiModel == nil {
		uddiModel = &UDDINamedListModel{}
	}
	uddiModel.Flatten(ctx, resp.UDDI, diags)
	if resp.UDDI != nil {
		m.UDDI = flex.FlattenNestedObject(ctx, uddiModel, UDDINamedListAttrTypes, diags)
	} else {
		m.UDDI = types.ObjectNull(UDDINamedListAttrTypes)
	}
}

// Flatten merges API response onto existing UDDI model.
func (m *UDDINamedListModel) Flatten(ctx context.Context, from *coremodel.UDDINamedListExt, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.ConfidenceLevel = flex.FlattenStringPointer(from.ConfidenceLevel)
	m.Description = flex.FlattenStringPointer(from.Description)
	m.Items = flex.FlattenFrameworkListString(ctx, from.Items, diags)
	m.ItemsDescribed = flex.FlattenFrameworkListNestedBlock(ctx, from.ItemsDescribed, ItemStructsAttrTypes, diags, FlattenItemStructs)
	m.Name = flex.FlattenStringPointer(from.Name)
	m.Policies = flex.FlattenFrameworkListString(ctx, from.Policies, diags)
	tagsAll := flex.FlattenMapStringAny(ctx, from.Tags, diags)
	if m.Tags.IsNull() || m.Tags.IsUnknown() {
		m.Tags = tagsAll
	}
	m.TagsAll = tagsAll
	m.ThreatLevel = flex.FlattenStringPointer(from.ThreatLevel)
	m.Type = flex.FlattenStringPointer(from.Type)
}
