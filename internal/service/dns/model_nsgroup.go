package dns

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/mapvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/mapdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"

	coremodel "github.com/infobloxopen/terraform-provider-infoblox/internal/core/model/dns"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
	immutable "github.com/infobloxopen/terraform-provider-infoblox/internal/planmodifiers/immutable"
	importmod "github.com/infobloxopen/terraform-provider-infoblox/internal/planmodifiers/import"
	customvalidator "github.com/infobloxopen/terraform-provider-infoblox/internal/validator"
)

type NsgroupModel struct {
	Id   types.String `tfsdk:"id"`
	NIOS types.Object `tfsdk:"nios"`
}

var NsgroupAttrTypes = map[string]attr.Type{
	"id":   types.StringType,
	"nios": types.ObjectType{AttrTypes: NIOSNsgroupAttrTypes},
}

type NIOSNsgroupModel struct {
	Comment             types.String `tfsdk:"comment"`
	ExtAttrs            types.Map    `tfsdk:"ext_attrs"`
	ExtAttrsAll         types.Map    `tfsdk:"ext_attrs_all"`
	ExternalPrimaries   types.List   `tfsdk:"external_primaries"`
	ExternalSecondaries types.List   `tfsdk:"external_secondaries"`
	GridPrimary         types.List   `tfsdk:"grid_primary"`
	GridSecondaries     types.List   `tfsdk:"grid_secondaries"`
	IsGridDefault       types.Bool   `tfsdk:"is_grid_default"`
	IsMultimaster       types.Bool   `tfsdk:"is_multimaster"`
	Name                types.String `tfsdk:"name"`
	UseExternalPrimary  types.Bool   `tfsdk:"use_external_primary"`
}

var NIOSNsgroupAttrTypes = map[string]attr.Type{
	"comment":              types.StringType,
	"ext_attrs":            types.MapType{ElemType: types.StringType},
	"ext_attrs_all":        types.MapType{ElemType: types.StringType},
	"external_primaries":   types.ListType{ElemType: types.ObjectType{AttrTypes: NsgroupExternalPrimariesAttrTypes}},
	"external_secondaries": types.ListType{ElemType: types.ObjectType{AttrTypes: NsgroupExternalSecondariesAttrTypes}},
	"grid_primary":         types.ListType{ElemType: types.ObjectType{AttrTypes: NsgroupGridPrimaryAttrTypes}},
	"grid_secondaries":     types.ListType{ElemType: types.ObjectType{AttrTypes: NsgroupGridSecondariesAttrTypes}},
	"is_grid_default":      types.BoolType,
	"is_multimaster":       types.BoolType,
	"name":                 types.StringType,
	"use_external_primary": types.BoolType,
}

const (
	NsgroupReturnFields = "comment,extattrs,external_primaries,external_secondaries,grid_primary,grid_secondaries,is_grid_default,is_multimaster,name,use_external_primary"
)

var NsgroupResourceSchemaAttributes = map[string]schema.Attribute{
	"id": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The reference to the object.",
	},
	"nios": schema.SingleNestedAttribute{
		Optional:            true,
		MarkdownDescription: "NIOS backend-specific fields.",
		Attributes:          NsgroupResourceNiosSchemaAttributes,
	},
}

var NsgroupResourceNiosSchemaAttributes = map[string]schema.Attribute{
	"comment": schema.StringAttribute{
		Default:  stringdefault.StaticString(""),
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
			customvalidator.ValidateTrimmedString(),
			stringvalidator.LengthBetween(0, 256),
		},
		MarkdownDescription: "Comment for the name server group; maximum 256 characters.",
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
	"external_primaries": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: NsgroupExternalPrimariesResourceSchemaAttributes,
		},
		Optional: true,
		Computed: true,
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
			listvalidator.AlsoRequires(path.MatchRelative().AtParent().AtName("use_external_primary")),
		},
		MarkdownDescription: "The list of external primary servers.",
	},
	"external_secondaries": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: NsgroupExternalSecondariesResourceSchemaAttributes,
		},
		Optional: true,
		Computed: true,
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
		},
		MarkdownDescription: "The list of external secondary servers.",
	},
	"grid_primary": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: NsgroupGridPrimaryResourceSchemaAttributes,
		},
		Optional: true,
		Computed: true,
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
			listvalidator.ExactlyOneOf(
				path.MatchRelative().AtParent().AtName("grid_primary"),
				path.MatchRelative().AtParent().AtName("external_primaries"),
			),
		},
		MarkdownDescription: "The grid primary servers for this group.",
	},
	"grid_secondaries": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: NsgroupGridSecondariesResourceSchemaAttributes,
		},
		Optional: true,
		Computed: true,
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
			listvalidator.AlsoRequires(path.MatchRelative().AtParent().AtName("use_external_primary")),
		},
		MarkdownDescription: "The list with Grid members that are secondary servers for this group.",
	},
	"is_grid_default": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Determines if this name server group is the Grid default.",
	},
	"is_multimaster": schema.BoolAttribute{
		Optional: true,
		Computed: true,
		PlanModifiers: []planmodifier.Bool{
			immutable.ImmutableBool(),
		},
		MarkdownDescription: "Determines if the \"multiple DNS primaries\" feature is enabled for the group.",
	},
	"name": schema.StringAttribute{
		Required: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "The name of this name server group.",
	},
	"use_external_primary": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "This flag controls whether the group is using an external primary. Note that modification of this field requires passing values for \"grid_secondaries\" and \"external_primaries\".",
	},
}

// Expand converts the TF model to the infoblox core model
func (m *NsgroupModel) Expand(ctx context.Context, diags *diag.Diagnostics, isCreate bool) *coremodel.Nsgroup {
	if m == nil {
		return nil
	}

	obj := &coremodel.Nsgroup{}

	// Expand NIOS nested attribute (returns nil if not present)
	niosModel := flex.ExpandNestedObject[NIOSNsgroupModel](ctx, m.NIOS, diags)
	if niosModel != nil {
		obj.NIOS = niosModel.Expand(ctx, diags, isCreate)
	}

	return obj
}

// Expand converts the NIOS TF model to the core model.
func (m *NIOSNsgroupModel) Expand(ctx context.Context, diags *diag.Diagnostics, isCreate bool) *coremodel.NIOSNsgroupExt {
	ext := &coremodel.NIOSNsgroupExt{
		Comment:             flex.ExpandStringPointerNullAsEmpty(m.Comment),
		ExtAttrs:            flex.ExpandMapStringAny(ctx, m.ExtAttrs, diags),
		ExternalPrimaries:   flex.ExpandFrameworkListNestedBlock(ctx, m.ExternalPrimaries, diags, ExpandNsgroupExternalPrimaries),
		ExternalSecondaries: flex.ExpandFrameworkListNestedBlock(ctx, m.ExternalSecondaries, diags, ExpandNsgroupExternalSecondaries),
		GridPrimary:         flex.ExpandFrameworkListNestedBlock(ctx, m.GridPrimary, diags, ExpandNsgroupGridPrimary),
		GridSecondaries:     flex.ExpandFrameworkListNestedBlock(ctx, m.GridSecondaries, diags, ExpandNsgroupGridSecondaries),
		IsGridDefault:       flex.ExpandBoolPointer(m.IsGridDefault),
		Name:                flex.ExpandStringPointerNullAsEmpty(m.Name),
		UseExternalPrimary:  flex.ExpandBoolPointer(m.UseExternalPrimary),
	}
	if isCreate {
		ext.IsMultimaster = flex.ExpandBoolPointer(m.IsMultimaster)
	}
	return ext
}

// Flatten populates the TF model from a core response.
func (m *NsgroupModel) Flatten(ctx context.Context, resp *coremodel.Nsgroup, diags *diag.Diagnostics) {
	if resp == nil {
		return
	}

	m.Id = flex.FlattenStringPointer(resp.Id)

	// Extract existing NIOS model, flatten API response onto it, convert back
	niosModel := flex.ExpandNestedObject[NIOSNsgroupModel](ctx, m.NIOS, diags)
	if niosModel == nil {
		niosModel = &NIOSNsgroupModel{}
	}
	plannedNIOS := flex.ExpandNestedObject[NIOSNsgroupModel](ctx, m.NIOS, diags)
	niosModel.Flatten(ctx, resp.NIOS, diags)
	if resp.NIOS != nil {
		PostFlattenNsgroupNIOS(ctx, plannedNIOS, niosModel, diags)
		m.NIOS = flex.FlattenNestedObject(ctx, niosModel, NIOSNsgroupAttrTypes, diags)
	} else {
		m.NIOS = types.ObjectNull(NIOSNsgroupAttrTypes)
	}

}

// Flatten merges API response onto existing NIOS model.
func (m *NIOSNsgroupModel) Flatten(ctx context.Context, from *coremodel.NIOSNsgroupExt, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	planExtAttrs := m.ExtAttrs
	if planExtAttrs.IsUnknown() {
		planExtAttrs = types.MapNull(types.StringType)
	}
	m.Comment = flex.FlattenStringPointerEmptyAsNull(from.Comment)
	m.ExtAttrs, m.ExtAttrsAll = flex.FlattenEAs(planExtAttrs, from.ExtAttrs)
	m.ExternalPrimaries = flex.FlattenFrameworkListNestedBlock(ctx, from.ExternalPrimaries, NsgroupExternalPrimariesAttrTypes, diags, FlattenNsgroupExternalPrimaries)
	m.ExternalSecondaries = flex.FlattenFrameworkListNestedBlock(ctx, from.ExternalSecondaries, NsgroupExternalSecondariesAttrTypes, diags, FlattenNsgroupExternalSecondaries)
	m.GridPrimary = flex.FlattenFrameworkListNestedBlock(ctx, from.GridPrimary, NsgroupGridPrimaryAttrTypes, diags, FlattenNsgroupGridPrimary)
	m.GridSecondaries = flex.FlattenFrameworkListNestedBlock(ctx, from.GridSecondaries, NsgroupGridSecondariesAttrTypes, diags, FlattenNsgroupGridSecondaries)
	m.IsGridDefault = flex.FlattenBoolPointer(from.IsGridDefault)
	m.IsMultimaster = flex.FlattenBoolPointer(from.IsMultimaster)
	m.Name = flex.FlattenStringPointerEmptyAsNull(from.Name)
	m.UseExternalPrimary = flex.FlattenBoolPointer(from.UseExternalPrimary)
}
