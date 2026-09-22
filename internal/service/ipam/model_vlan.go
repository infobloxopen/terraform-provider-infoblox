package ipam

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
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"

	coremodel "github.com/infobloxopen/terraform-provider-infoblox/internal/core/model/ipam"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
	importmod "github.com/infobloxopen/terraform-provider-infoblox/internal/planmodifiers/import"
	customvalidator "github.com/infobloxopen/terraform-provider-infoblox/internal/validator"
)

type VlanModel struct {
	Id   types.String `tfsdk:"id"`
	NIOS types.Object `tfsdk:"nios"`
}

var VlanAttrTypes = map[string]attr.Type{
	"id":   types.StringType,
	"nios": types.ObjectType{AttrTypes: NIOSVlanAttrTypes},
}

type NIOSVlanModel struct {
	Comment     types.String `tfsdk:"comment"`
	Contact     types.String `tfsdk:"contact"`
	Department  types.String `tfsdk:"department"`
	Description types.String `tfsdk:"description"`
	ExtAttrs    types.Map    `tfsdk:"ext_attrs"`
	ExtAttrsAll types.Map    `tfsdk:"ext_attrs_all"`
	Id          types.Int64  `tfsdk:"id"`
	Name        types.String `tfsdk:"name"`
	Parent      types.String `tfsdk:"parent"`
	Reserved    types.Bool   `tfsdk:"reserved"`
}

var NIOSVlanAttrTypes = map[string]attr.Type{
	"comment":       types.StringType,
	"contact":       types.StringType,
	"department":    types.StringType,
	"description":   types.StringType,
	"ext_attrs":     types.MapType{ElemType: types.StringType},
	"ext_attrs_all": types.MapType{ElemType: types.StringType},
	"id":            types.Int64Type,
	"name":          types.StringType,
	"parent":        types.StringType,
	"reserved":      types.BoolType,
}

const (
	VlanReturnFields = "assigned_to,comment,contact,department,description,extattrs,id,name,parent,reserved,status"
)

var VlanResourceSchemaAttributes = map[string]schema.Attribute{
	"id": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The reference to the object.",
	},
	"nios": schema.SingleNestedAttribute{
		Optional:            true,
		MarkdownDescription: "NIOS backend-specific fields.",
		Attributes:          VlanResourceNiosSchemaAttributes,
	},
}

var VlanResourceNiosSchemaAttributes = map[string]schema.Attribute{
	"comment": schema.StringAttribute{
		Optional: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
			customvalidator.ValidateTrimmedString(),
			stringvalidator.LengthBetween(0, 256),
		},
		MarkdownDescription: "A descriptive comment for this VLAN.",
	},
	"contact": schema.StringAttribute{
		Optional: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "Contact information for person/team managing or using VLAN.",
	},
	"department": schema.StringAttribute{
		Optional: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "Department where VLAN is used.",
	},
	"description": schema.StringAttribute{
		Optional: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "Description for the VLAN object, may be potentially used for longer VLAN names.",
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
	"id": schema.Int64Attribute{
		Required:            true,
		MarkdownDescription: "VLAN ID value.",
	},
	"name": schema.StringAttribute{
		Required: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "Name of the VLAN.",
	},
	"parent": schema.StringAttribute{
		Required: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "The VLAN View or VLAN Range to which this VLAN belongs.",
	},
	"reserved": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "When set VLAN can only be assigned to IPAM object manually.",
	},
}

// Expand converts the TF model to the infoblox core model
func (m *VlanModel) Expand(ctx context.Context, diags *diag.Diagnostics, isCreate bool) *coremodel.Vlan {
	if m == nil {
		return nil
	}

	obj := &coremodel.Vlan{}

	// Expand NIOS nested attribute (returns nil if not present)
	niosModel := flex.ExpandNestedObject[NIOSVlanModel](ctx, m.NIOS, diags)
	if niosModel != nil {
		obj.NIOS = niosModel.Expand(ctx, diags)
	}

	return obj
}

// Expand converts the NIOS TF model to the core model.
func (m *NIOSVlanModel) Expand(ctx context.Context, diags *diag.Diagnostics) *coremodel.NIOSVlanExt {
	return &coremodel.NIOSVlanExt{
		Comment:     flex.ExpandStringPointerNullAsEmpty(m.Comment),
		Contact:     flex.ExpandStringPointerNullAsEmpty(m.Contact),
		Department:  flex.ExpandStringPointerNullAsEmpty(m.Department),
		Description: flex.ExpandStringPointerNullAsEmpty(m.Description),
		ExtAttrs:    flex.ExpandMapStringAny(ctx, m.ExtAttrs, diags),
		Id:          ExpandVlanId(m.Id),
		Name:        flex.ExpandStringPointerNullAsEmpty(m.Name),
		Parent:      ExpandVlanParent(m.Parent),
		Reserved:    flex.ExpandBoolPointer(m.Reserved),
	}
}

// Flatten populates the TF model from a core response.
func (m *VlanModel) Flatten(ctx context.Context, resp *coremodel.Vlan, diags *diag.Diagnostics) {
	if resp == nil {
		return
	}

	m.Id = flex.FlattenStringPointer(resp.Id)

	// Extract existing NIOS model, flatten API response onto it, convert back
	niosModel := flex.ExpandNestedObject[NIOSVlanModel](ctx, m.NIOS, diags)
	if niosModel == nil {
		niosModel = &NIOSVlanModel{}
	}
	niosModel.Flatten(ctx, resp.NIOS, diags)
	if resp.NIOS != nil {
		m.NIOS = flex.FlattenNestedObject(ctx, niosModel, NIOSVlanAttrTypes, diags)
	} else {
		m.NIOS = types.ObjectNull(NIOSVlanAttrTypes)
	}

}

// Flatten merges API response onto existing NIOS model.
func (m *NIOSVlanModel) Flatten(ctx context.Context, from *coremodel.NIOSVlanExt, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	planExtAttrs := m.ExtAttrs
	if planExtAttrs.IsUnknown() {
		planExtAttrs = types.MapNull(types.StringType)
	}
	m.Comment = flex.FlattenStringPointerEmptyAsNull(from.Comment)
	m.Contact = flex.FlattenStringPointerEmptyAsNull(from.Contact)
	m.Department = flex.FlattenStringPointerEmptyAsNull(from.Department)
	m.Description = flex.FlattenStringPointerEmptyAsNull(from.Description)
	m.ExtAttrs, m.ExtAttrsAll = flex.FlattenEAs(planExtAttrs, from.ExtAttrs)
	m.Id = FlattenVlanId(from.Id)
	m.Name = flex.FlattenStringPointerEmptyAsNull(from.Name)
	m.Parent = FlattenVlanParent(from.Parent)
	m.Reserved = flex.FlattenBoolPointer(from.Reserved)
}
