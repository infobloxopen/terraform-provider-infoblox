package dns

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/mapvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/mapdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"

	coremodel "github.com/infobloxopen/terraform-provider-infoblox/internal/core/model/dns"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
	importmod "github.com/infobloxopen/terraform-provider-infoblox/internal/planmodifiers/import"
	customvalidator "github.com/infobloxopen/terraform-provider-infoblox/internal/validator"
)

type NsgroupDelegationModel struct {
	Id   types.String `tfsdk:"id"`
	NIOS types.Object `tfsdk:"nios"`
}

var NsgroupDelegationAttrTypes = map[string]attr.Type{
	"id":   types.StringType,
	"nios": types.ObjectType{AttrTypes: NIOSNsgroupDelegationAttrTypes},
}

type NIOSNsgroupDelegationModel struct {
	Comment     types.String `tfsdk:"comment"`
	DelegateTo  types.List   `tfsdk:"delegate_to"`
	ExtAttrs    types.Map    `tfsdk:"ext_attrs"`
	ExtAttrsAll types.Map    `tfsdk:"ext_attrs_all"`
	Name        types.String `tfsdk:"name"`
}

var NIOSNsgroupDelegationAttrTypes = map[string]attr.Type{
	"comment":       types.StringType,
	"delegate_to":   types.ListType{ElemType: types.ObjectType{AttrTypes: NsgroupDelegationDelegateToAttrTypes}},
	"ext_attrs":     types.MapType{ElemType: types.StringType},
	"ext_attrs_all": types.MapType{ElemType: types.StringType},
	"name":          types.StringType,
}

const (
	NsgroupDelegationReturnFields = "comment,delegate_to,extattrs,name"
)

var NsgroupDelegationResourceSchemaAttributes = map[string]schema.Attribute{
	"id": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The reference to the object.",
	},
	"nios": schema.SingleNestedAttribute{
		Optional:            true,
		MarkdownDescription: "NIOS backend-specific fields.",
		Attributes:          NsgroupDelegationResourceNiosSchemaAttributes,
	},
}

var NsgroupDelegationResourceNiosSchemaAttributes = map[string]schema.Attribute{
	"comment": schema.StringAttribute{
		Default:  stringdefault.StaticString(""),
		Optional: true,
		Computed: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
			customvalidator.ValidateTrimmedString(),
			stringvalidator.LengthBetween(0, 256),
		},
		MarkdownDescription: "The comment for the delegated NS group.",
	},
	"delegate_to": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: NsgroupDelegationDelegateToResourceSchemaAttributes,
		},
		Required: true,
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
		},
		MarkdownDescription: "The list of delegated servers for the delegated NS group.",
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
		},
		MarkdownDescription: "The name of the delegated NS group.",
	},
}

// Expand converts the TF model to the infoblox core model
func (m *NsgroupDelegationModel) Expand(ctx context.Context, diags *diag.Diagnostics, isCreate bool) *coremodel.NsgroupDelegation {
	if m == nil {
		return nil
	}

	obj := &coremodel.NsgroupDelegation{}

	// Expand NIOS nested attribute (returns nil if not present)
	niosModel := flex.ExpandNestedObject[NIOSNsgroupDelegationModel](ctx, m.NIOS, diags)
	if niosModel != nil {
		obj.NIOS = niosModel.Expand(ctx, diags)
	}

	return obj
}

// Expand converts the NIOS TF model to the core model.
func (m *NIOSNsgroupDelegationModel) Expand(ctx context.Context, diags *diag.Diagnostics) *coremodel.NIOSNsgroupDelegationExt {
	return &coremodel.NIOSNsgroupDelegationExt{
		Comment:    flex.ExpandStringPointerNullAsEmpty(m.Comment),
		DelegateTo: flex.ExpandFrameworkListNestedBlock(ctx, m.DelegateTo, diags, ExpandNsgroupDelegationDelegateTo),
		ExtAttrs:   flex.ExpandMapStringAny(ctx, m.ExtAttrs, diags),
		Name:       flex.ExpandStringPointerNullAsEmpty(m.Name),
	}
}

// Flatten populates the TF model from a core response.
func (m *NsgroupDelegationModel) Flatten(ctx context.Context, resp *coremodel.NsgroupDelegation, diags *diag.Diagnostics) {
	if resp == nil {
		return
	}

	m.Id = flex.FlattenStringPointer(resp.Id)

	// Extract existing NIOS model, flatten API response onto it, convert back
	niosModel := flex.ExpandNestedObject[NIOSNsgroupDelegationModel](ctx, m.NIOS, diags)
	if niosModel == nil {
		niosModel = &NIOSNsgroupDelegationModel{}
	}
	plannedNIOS := flex.ExpandNestedObject[NIOSNsgroupDelegationModel](ctx, m.NIOS, diags)
	niosModel.Flatten(ctx, resp.NIOS, diags)
	if resp.NIOS != nil {
		PostFlattenNsgroupDelegationNIOS(ctx, plannedNIOS, niosModel, diags)
		m.NIOS = flex.FlattenNestedObject(ctx, niosModel, NIOSNsgroupDelegationAttrTypes, diags)
	} else {
		m.NIOS = types.ObjectNull(NIOSNsgroupDelegationAttrTypes)
	}

}

// Flatten merges API response onto existing NIOS model.
func (m *NIOSNsgroupDelegationModel) Flatten(ctx context.Context, from *coremodel.NIOSNsgroupDelegationExt, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	planExtAttrs := m.ExtAttrs
	if planExtAttrs.IsUnknown() {
		planExtAttrs = types.MapNull(types.StringType)
	}
	m.Comment = flex.FlattenStringPointerEmptyAsNull(from.Comment)
	m.DelegateTo = flex.FlattenFrameworkListNestedBlock(ctx, from.DelegateTo, NsgroupDelegationDelegateToAttrTypes, diags, FlattenNsgroupDelegationDelegateTo)
	m.ExtAttrs, m.ExtAttrsAll = flex.FlattenEAs(planExtAttrs, from.ExtAttrs)
	m.Name = flex.FlattenStringPointerEmptyAsNull(from.Name)
}
