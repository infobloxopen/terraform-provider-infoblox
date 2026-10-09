package ipam

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
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
	immutable "github.com/infobloxopen/terraform-provider-infoblox/internal/planmodifiers/immutable"
	importmod "github.com/infobloxopen/terraform-provider-infoblox/internal/planmodifiers/import"
	customvalidator "github.com/infobloxopen/terraform-provider-infoblox/internal/validator"
)

type VlanviewModel struct {
	Id            types.String `tfsdk:"id"`
	UpdateTrigger types.String `tfsdk:"update_trigger"`
	NIOS          types.Object `tfsdk:"nios"`
}

var VlanviewAttrTypes = map[string]attr.Type{
	"id":             types.StringType,
	"update_trigger": types.StringType,
	"nios":           types.ObjectType{AttrTypes: NIOSVlanviewAttrTypes},
}

type NIOSVlanviewModel struct {
	AllowRangeOverlapping types.Bool   `tfsdk:"allow_range_overlapping"`
	Comment               types.String `tfsdk:"comment"`
	EndVlanId             types.Int64  `tfsdk:"end_vlan_id"`
	ExtAttrs              types.Map    `tfsdk:"ext_attrs"`
	ExtAttrsAll           types.Map    `tfsdk:"ext_attrs_all"`
	Name                  types.String `tfsdk:"name"`
	PreCreateVlan         types.Bool   `tfsdk:"pre_create_vlan"`
	StartVlanId           types.Int64  `tfsdk:"start_vlan_id"`
	VlanNamePrefix        types.String `tfsdk:"vlan_name_prefix"`
}

var NIOSVlanviewAttrTypes = map[string]attr.Type{
	"allow_range_overlapping": types.BoolType,
	"comment":                 types.StringType,
	"end_vlan_id":             types.Int64Type,
	"ext_attrs":               types.MapType{ElemType: types.StringType},
	"ext_attrs_all":           types.MapType{ElemType: types.StringType},
	"name":                    types.StringType,
	"pre_create_vlan":         types.BoolType,
	"start_vlan_id":           types.Int64Type,
	"vlan_name_prefix":        types.StringType,
}

const (
	VlanviewReturnFields = "allow_range_overlapping,comment,end_vlan_id,extattrs,name,pre_create_vlan,start_vlan_id,vlan_name_prefix"
)

var VlanviewResourceSchemaAttributes = map[string]schema.Attribute{
	"id": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The reference to the object.",
	},
	"update_trigger": schema.StringAttribute{
		Optional:            true,
		MarkdownDescription: "An arbitrary value used to trigger an update. Not sent to the API. Change it when Terraform reports no infrastructure changes.",
	},
	"nios": schema.SingleNestedAttribute{
		Optional:            true,
		MarkdownDescription: "NIOS backend-specific fields.",
		Attributes:          VlanviewResourceNiosSchemaAttributes,
	},
}

var VlanviewResourceNiosSchemaAttributes = map[string]schema.Attribute{
	"allow_range_overlapping": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "When set to true VLAN Ranges under VLAN View can have overlapping ID.",
	},
	"comment": schema.StringAttribute{
		Optional: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
			customvalidator.ValidateTrimmedString(),
			stringvalidator.LengthBetween(0, 256),
		},
		MarkdownDescription: "A descriptive comment for this VLAN View.",
	},
	"end_vlan_id": schema.Int64Attribute{
		Required: true,
		Validators: []validator.Int64{
			int64validator.Between(1, 4094),
		},
		MarkdownDescription: "End ID for VLAN View.",
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
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "Name of the VLAN View.",
	},
	"pre_create_vlan": schema.BoolAttribute{
		Optional: true,
		Computed: true,
		Default:  booldefault.StaticBool(false),
		PlanModifiers: []planmodifier.Bool{
			immutable.ImmutableBool(),
		},
		MarkdownDescription: "If set on creation VLAN objects will be created once VLAN View created.",
	},
	"start_vlan_id": schema.Int64Attribute{
		Required: true,
		Validators: []validator.Int64{
			int64validator.Between(1, 4094),
		},
		MarkdownDescription: "Start ID for VLAN View.",
	},
	"vlan_name_prefix": schema.StringAttribute{
		Optional: true,
		PlanModifiers: []planmodifier.String{
			immutable.ImmutableString(),
		},
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "If set on creation prefix string will be used for VLAN name.",
	},
}

// Expand converts the TF model to the infoblox core model
func (m *VlanviewModel) Expand(ctx context.Context, diags *diag.Diagnostics, isCreate bool) *coremodel.Vlanview {
	if m == nil {
		return nil
	}

	obj := &coremodel.Vlanview{}

	// Expand NIOS nested attribute (returns nil if not present)
	niosModel := flex.ExpandNestedObject[NIOSVlanviewModel](ctx, m.NIOS, diags)
	if niosModel != nil {
		obj.NIOS = niosModel.Expand(ctx, diags, isCreate)
	}

	return obj
}

// Expand converts the NIOS TF model to the core model.
func (m *NIOSVlanviewModel) Expand(ctx context.Context, diags *diag.Diagnostics, isCreate bool) *coremodel.NIOSVlanviewExt {
	ext := &coremodel.NIOSVlanviewExt{
		AllowRangeOverlapping: flex.ExpandBoolPointer(m.AllowRangeOverlapping),
		Comment:               flex.ExpandStringPointerNullAsEmpty(m.Comment),
		EndVlanId:             flex.ExpandInt64Pointer(m.EndVlanId),
		ExtAttrs:              flex.ExpandMapStringAny(ctx, m.ExtAttrs, diags),
		Name:                  flex.ExpandStringPointerNullAsEmpty(m.Name),
		StartVlanId:           flex.ExpandInt64Pointer(m.StartVlanId),
	}
	if isCreate {
		ext.PreCreateVlan = flex.ExpandBoolPointer(m.PreCreateVlan)
		ext.VlanNamePrefix = flex.ExpandStringPointerNullAsEmpty(m.VlanNamePrefix)
	}
	return ext
}

// Flatten populates the TF model from a core response.
func (m *VlanviewModel) Flatten(ctx context.Context, resp *coremodel.Vlanview, diags *diag.Diagnostics) {
	if resp == nil {
		return
	}

	m.Id = flex.FlattenStringPointer(resp.Id)

	// Extract existing NIOS model, flatten API response onto it, convert back
	niosModel := flex.ExpandNestedObject[NIOSVlanviewModel](ctx, m.NIOS, diags)
	if niosModel == nil {
		niosModel = &NIOSVlanviewModel{}
	}
	plannedNIOS := flex.ExpandNestedObject[NIOSVlanviewModel](ctx, m.NIOS, diags)
	niosModel.Flatten(ctx, resp.NIOS, diags)
	if resp.NIOS != nil {
		PostFlattenVlanviewNIOS(ctx, plannedNIOS, niosModel, diags)
		m.NIOS = flex.FlattenNestedObject(ctx, niosModel, NIOSVlanviewAttrTypes, diags)
	} else {
		m.NIOS = types.ObjectNull(NIOSVlanviewAttrTypes)
	}

}

// Flatten merges API response onto existing NIOS model.
func (m *NIOSVlanviewModel) Flatten(ctx context.Context, from *coremodel.NIOSVlanviewExt, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	planExtAttrs := m.ExtAttrs
	if planExtAttrs.IsUnknown() {
		planExtAttrs = types.MapNull(types.StringType)
	}
	m.AllowRangeOverlapping = flex.FlattenBoolPointer(from.AllowRangeOverlapping)
	m.Comment = flex.FlattenStringPointerEmptyAsNull(from.Comment)
	m.EndVlanId = flex.FlattenInt64Pointer(from.EndVlanId)
	m.ExtAttrs, m.ExtAttrsAll = flex.FlattenEAs(planExtAttrs, from.ExtAttrs)
	m.Name = flex.FlattenStringPointerEmptyAsNull(from.Name)
	m.PreCreateVlan = flex.FlattenBoolPointer(from.PreCreateVlan)
	m.StartVlanId = flex.FlattenInt64Pointer(from.StartVlanId)
	m.VlanNamePrefix = flex.FlattenStringPointerEmptyAsNull(from.VlanNamePrefix)
}
