package dtc

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/mapvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/mapdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"

	coremodel "github.com/infobloxopen/terraform-provider-infoblox/internal/core/model/dtc"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
	importmod "github.com/infobloxopen/terraform-provider-infoblox/internal/planmodifiers/import"
	customvalidator "github.com/infobloxopen/terraform-provider-infoblox/internal/validator"
)

type DtcTopologyModel struct {
	Id   types.String `tfsdk:"id"`
	NIOS types.Object `tfsdk:"nios"`
	UDDI types.Object `tfsdk:"uddi"`
}

var DtcTopologyAttrTypes = map[string]attr.Type{
	"id":   types.StringType,
	"nios": types.ObjectType{AttrTypes: NIOSDtcTopologyAttrTypes},
	"uddi": types.ObjectType{AttrTypes: UDDIDtcTopologyAttrTypes},
}

type NIOSDtcTopologyModel struct {
	Comment     types.String `tfsdk:"comment"`
	ExtAttrs    types.Map    `tfsdk:"ext_attrs"`
	ExtAttrsAll types.Map    `tfsdk:"ext_attrs_all"`
	Name        types.String `tfsdk:"name"`
	Rules       types.List   `tfsdk:"rules"`
}

var NIOSDtcTopologyAttrTypes = map[string]attr.Type{
	"comment":       types.StringType,
	"ext_attrs":     types.MapType{ElemType: types.StringType},
	"ext_attrs_all": types.MapType{ElemType: types.StringType},
	"name":          types.StringType,
	"rules":         types.ListType{ElemType: types.ObjectType{AttrTypes: TopologyRulesInnerAttrTypes}},
}

type UDDIDtcTopologyModel struct {
	Comment  types.String `tfsdk:"comment"`
	Disabled types.Bool   `tfsdk:"disabled"`
	Name     types.String `tfsdk:"name"`
	Sources  types.List   `tfsdk:"sources"`
	Tags     types.Map    `tfsdk:"tags"`
	TagsAll  types.Map    `tfsdk:"tags_all"`
}

var UDDIDtcTopologyAttrTypes = map[string]attr.Type{
	"comment":  types.StringType,
	"disabled": types.BoolType,
	"name":     types.StringType,
	"sources":  types.ListType{ElemType: types.ObjectType{AttrTypes: TopologySourceAttrTypes}},
	"tags":     types.MapType{ElemType: types.StringType},
	"tags_all": types.MapType{ElemType: types.StringType},
}

const (
	DtcTopologyReturnFields = "comment,extattrs,name,rules"
)

var DtcTopologyResourceSchemaAttributes = map[string]schema.Attribute{
	"id": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The reference to the object.",
	},
	"nios": schema.SingleNestedAttribute{
		Optional:            true,
		MarkdownDescription: "NIOS backend-specific fields.",
		Attributes:          DtcTopologyResourceNiosSchemaAttributes,
	},
	"uddi": schema.SingleNestedAttribute{
		Optional:            true,
		MarkdownDescription: "UDDI backend-specific fields.",
		Attributes:          DtcTopologyResourceUddiSchemaAttributes,
	},
}

var DtcTopologyResourceNiosSchemaAttributes = map[string]schema.Attribute{
	"comment": schema.StringAttribute{
		Optional: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
			stringvalidator.LengthBetween(0, 256),
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "The comment for the DTC TOPOLOGY monitor object; maximum 256 characters.",
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
	"name": schema.StringAttribute{
		Required: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
			stringvalidator.LengthBetween(0, 256),
		},
		MarkdownDescription: "Display name of the DTC Topology.",
	},
	"rules": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: TopologyRulesInnerResourceSchemaAttributes,
		},
		Optional: true,
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
		},
		MarkdownDescription: "Topology rules.",
	},
}

var DtcTopologyResourceUddiSchemaAttributes = map[string]schema.Attribute{
	"comment": schema.StringAttribute{
		Default:             stringdefault.StaticString(""),
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Optional. Comment for __Topology__.",
	},
	"disabled": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Optional. Flag which enables/disables __Topology__.  Defaults to _false_.",
	},
	"name": schema.StringAttribute{
		Required:            true,
		MarkdownDescription: "Display name of __Topology__.",
	},
	"sources": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: TopologySourceResourceSchemaAttributes,
		},
		Optional: true,
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
		},
		MarkdownDescription: "Required. List of __TopologySource__ objects with unique names.",
	},
	"tags": schema.MapAttribute{
		Optional:    true,
		Computed:    true,
		ElementType: types.StringType,
		Default:     mapdefault.StaticValue(types.MapNull(types.StringType)),
		Validators: []validator.Map{
			mapvalidator.SizeAtLeast(1),
		},
		MarkdownDescription: "Optional. The tags for __Topology__ in JSON format.",
	},
	"tags_all": schema.MapAttribute{
		Computed:            true,
		ElementType:         types.StringType,
		MarkdownDescription: "All tags including inherited values.",
	},
}

// Expand converts the TF model to the infoblox core model
func (m *DtcTopologyModel) Expand(ctx context.Context, diags *diag.Diagnostics, isCreate bool) *coremodel.DtcTopology {
	if m == nil {
		return nil
	}

	obj := &coremodel.DtcTopology{}

	// Expand NIOS nested attribute (returns nil if not present)
	niosModel := flex.ExpandNestedObject[NIOSDtcTopologyModel](ctx, m.NIOS, diags)
	if niosModel != nil {
		obj.NIOS = niosModel.Expand(ctx, diags)
	}

	// Expand UDDI nested attribute (returns nil if not present)
	uddiModel := flex.ExpandNestedObject[UDDIDtcTopologyModel](ctx, m.UDDI, diags)
	if uddiModel != nil {
		obj.UDDI = uddiModel.Expand(ctx, diags)
	}

	return obj
}

// Expand converts the NIOS TF model to the core model.
func (m *NIOSDtcTopologyModel) Expand(ctx context.Context, diags *diag.Diagnostics) *coremodel.NIOSDtcTopologyExt {
	return &coremodel.NIOSDtcTopologyExt{
		Comment:  flex.ExpandStringPointerNullAsEmpty(m.Comment),
		ExtAttrs: flex.ExpandMapStringAny(ctx, m.ExtAttrs, diags),
		Name:     flex.ExpandStringPointerNullAsEmpty(m.Name),
		Rules:    flex.ExpandFrameworkListNestedBlock(ctx, m.Rules, diags, ExpandTopologyRulesInner),
	}
}

// Expand converts the UDDI TF model to the core model.
func (m *UDDIDtcTopologyModel) Expand(ctx context.Context, diags *diag.Diagnostics) *coremodel.UDDIDtcTopologyExt {
	return &coremodel.UDDIDtcTopologyExt{
		Comment:  flex.ExpandStringPointer(m.Comment),
		Disabled: flex.ExpandBoolPointer(m.Disabled),
		Name:     flex.ExpandString(m.Name),
		Sources:  flex.ExpandFrameworkListNestedBlock(ctx, m.Sources, diags, ExpandTopologySource),
		Tags:     flex.ExpandMapStringAny(ctx, m.Tags, diags),
	}
}

// Flatten populates the TF model from a core response.
func (m *DtcTopologyModel) Flatten(ctx context.Context, resp *coremodel.DtcTopology, diags *diag.Diagnostics) {
	if resp == nil {
		return
	}

	m.Id = flex.FlattenStringPointer(resp.Id)

	// Extract existing NIOS model, flatten API response onto it, convert back
	niosModel := flex.ExpandNestedObject[NIOSDtcTopologyModel](ctx, m.NIOS, diags)
	if niosModel == nil {
		niosModel = &NIOSDtcTopologyModel{}
	}
	niosModel.Flatten(ctx, resp.NIOS, diags)
	if resp.NIOS != nil {
		m.NIOS = flex.FlattenNestedObject(ctx, niosModel, NIOSDtcTopologyAttrTypes, diags)
	} else {
		m.NIOS = types.ObjectNull(NIOSDtcTopologyAttrTypes)
	}

	// Extract existing UDDI model, flatten API response onto it, convert back
	uddiModel := flex.ExpandNestedObject[UDDIDtcTopologyModel](ctx, m.UDDI, diags)
	if uddiModel == nil {
		uddiModel = &UDDIDtcTopologyModel{}
	}
	uddiModel.Flatten(ctx, resp.UDDI, diags)
	if resp.UDDI != nil {
		m.UDDI = flex.FlattenNestedObject(ctx, uddiModel, UDDIDtcTopologyAttrTypes, diags)
	} else {
		m.UDDI = types.ObjectNull(UDDIDtcTopologyAttrTypes)
	}
}

// Flatten merges API response onto existing NIOS model.
func (m *NIOSDtcTopologyModel) Flatten(ctx context.Context, from *coremodel.NIOSDtcTopologyExt, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	planExtAttrs := m.ExtAttrs
	if planExtAttrs.IsUnknown() {
		planExtAttrs = types.MapNull(types.StringType)
	}
	m.Comment = flex.FlattenStringPointerEmptyAsNull(from.Comment)
	m.ExtAttrs, m.ExtAttrsAll = flex.FlattenEAs(planExtAttrs, from.ExtAttrs)
	m.Name = flex.FlattenStringPointerEmptyAsNull(from.Name)
	m.Rules = flex.FlattenFrameworkListNestedBlock(ctx, from.Rules, TopologyRulesInnerAttrTypes, diags, FlattenTopologyRulesInner)
}

// Flatten merges API response onto existing UDDI model.
func (m *UDDIDtcTopologyModel) Flatten(ctx context.Context, from *coremodel.UDDIDtcTopologyExt, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.Comment = flex.FlattenStringPointer(from.Comment)
	m.Disabled = flex.FlattenBoolPointer(from.Disabled)
	m.Name = flex.FlattenString(from.Name)
	m.Sources = flex.FlattenFrameworkListNestedBlock(ctx, from.Sources, TopologySourceAttrTypes, diags, FlattenTopologySource)
	tagsAll := flex.FlattenMapStringAny(ctx, from.Tags, diags)
	if m.Tags.IsNull() || m.Tags.IsUnknown() {
		m.Tags = tagsAll
	}
	m.TagsAll = tagsAll
}
