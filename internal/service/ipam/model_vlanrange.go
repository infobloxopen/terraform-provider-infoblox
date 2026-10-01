package ipam

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/mapvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
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

type VlanrangeModel struct {
	Id            types.String `tfsdk:"id"`
	UpdateTrigger types.String `tfsdk:"update_trigger"`
	NIOS          types.Object `tfsdk:"nios"`
}

var VlanrangeAttrTypes = map[string]attr.Type{
	"id":             types.StringType,
	"update_trigger": types.StringType,
	"nios":           types.ObjectType{AttrTypes: NIOSVlanrangeAttrTypes},
}

type NIOSVlanrangeModel struct {
	Comment        types.String `tfsdk:"comment"`
	DeleteVlans    types.Bool   `tfsdk:"delete_vlans"`
	EndVlanId      types.Int64  `tfsdk:"end_vlan_id"`
	ExtAttrs       types.Map    `tfsdk:"ext_attrs"`
	ExtAttrsAll    types.Map    `tfsdk:"ext_attrs_all"`
	Name           types.String `tfsdk:"name"`
	PreCreateVlan  types.Bool   `tfsdk:"pre_create_vlan"`
	StartVlanId    types.Int64  `tfsdk:"start_vlan_id"`
	VlanNamePrefix types.String `tfsdk:"vlan_name_prefix"`
	VlanView       types.String `tfsdk:"vlan_view"`
}

var NIOSVlanrangeAttrTypes = map[string]attr.Type{
	"comment":          types.StringType,
	"delete_vlans":     types.BoolType,
	"end_vlan_id":      types.Int64Type,
	"ext_attrs":        types.MapType{ElemType: types.StringType},
	"ext_attrs_all":    types.MapType{ElemType: types.StringType},
	"name":             types.StringType,
	"pre_create_vlan":  types.BoolType,
	"start_vlan_id":    types.Int64Type,
	"vlan_name_prefix": types.StringType,
	"vlan_view":        types.StringType,
}

const (
	VlanrangeReturnFields = "comment,end_vlan_id,extattrs,name,pre_create_vlan,start_vlan_id,vlan_name_prefix,vlan_view"
)

var VlanrangeResourceSchemaAttributes = map[string]schema.Attribute{
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
		Attributes:          VlanrangeResourceNiosSchemaAttributes,
	},
}

var VlanrangeResourceNiosSchemaAttributes = map[string]schema.Attribute{
	"comment": schema.StringAttribute{
		Optional: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
			customvalidator.ValidateTrimmedString(),
			stringvalidator.LengthBetween(0, 256),
		},
		MarkdownDescription: "A descriptive comment for this VLAN Range.",
	},
	"delete_vlans": schema.BoolAttribute{
		Computed:            true,
		MarkdownDescription: "Vlans delete option. Determines whether all child objects should be removed alongside with the VLAN Range or child objects should be assigned to another parental VLAN Range/View. By default child objects are re-parented.",
	},
	"end_vlan_id": schema.Int64Attribute{
		Required: true,
		Validators: []validator.Int64{
			int64validator.Between(1, 4094),
		},
		MarkdownDescription: "End ID for VLAN Range.",
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
		MarkdownDescription: "Name of the VLAN Range.",
	},
	"pre_create_vlan": schema.BoolAttribute{
		Optional: true,
		Computed: true,
		PlanModifiers: []planmodifier.Bool{
			immutable.ImmutableBool(),
		},
		MarkdownDescription: "If set on creation VLAN objects will be created once VLAN Range created.",
	},
	"start_vlan_id": schema.Int64Attribute{
		Required: true,
		Validators: []validator.Int64{
			int64validator.Between(1, 4094),
		},
		MarkdownDescription: "Start ID for VLAN Range.",
	},
	"vlan_name_prefix": schema.StringAttribute{
		Optional: true,
		Computed: true,
		PlanModifiers: []planmodifier.String{
			immutable.ImmutableString(),
		},
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "If set on creation prefix string will be used for VLAN name.",
	},
	"vlan_view": schema.StringAttribute{
		Required: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "The VLAN View to which this VLAN Range belongs.",
	},
}

// Expand converts the TF model to the infoblox core model
func (m *VlanrangeModel) Expand(ctx context.Context, diags *diag.Diagnostics, isCreate bool) *coremodel.Vlanrange {
	if m == nil {
		return nil
	}

	obj := &coremodel.Vlanrange{}

	// Expand NIOS nested attribute (returns nil if not present)
	niosModel := flex.ExpandNestedObject[NIOSVlanrangeModel](ctx, m.NIOS, diags)
	if niosModel != nil {
		obj.NIOS = niosModel.Expand(ctx, diags, isCreate)
	}

	return obj
}

// Expand converts the NIOS TF model to the core model.
func (m *NIOSVlanrangeModel) Expand(ctx context.Context, diags *diag.Diagnostics, isCreate bool) *coremodel.NIOSVlanrangeExt {
	ext := &coremodel.NIOSVlanrangeExt{
		Comment:     flex.ExpandStringPointerNullAsEmpty(m.Comment),
		DeleteVlans: flex.ExpandBoolPointer(m.DeleteVlans),
		EndVlanId:   flex.ExpandInt64Pointer(m.EndVlanId),
		ExtAttrs:    flex.ExpandMapStringAny(ctx, m.ExtAttrs, diags),
		Name:        flex.ExpandStringPointerNullAsEmpty(m.Name),
		StartVlanId: flex.ExpandInt64Pointer(m.StartVlanId),
		VlanView:    ExpandVlanrangeVlanView(m.VlanView),
	}
	if isCreate {
		ext.PreCreateVlan = flex.ExpandBoolPointer(m.PreCreateVlan)
		ext.VlanNamePrefix = flex.ExpandStringPointerNullAsEmpty(m.VlanNamePrefix)
	}
	return ext
}

// Flatten populates the TF model from a core response.
func (m *VlanrangeModel) Flatten(ctx context.Context, resp *coremodel.Vlanrange, diags *diag.Diagnostics) {
	if resp == nil {
		return
	}

	m.Id = flex.FlattenStringPointer(resp.Id)

	// Extract existing NIOS model, flatten API response onto it, convert back
	niosModel := flex.ExpandNestedObject[NIOSVlanrangeModel](ctx, m.NIOS, diags)
	if niosModel == nil {
		niosModel = &NIOSVlanrangeModel{}
	}
	plannedNIOS := flex.ExpandNestedObject[NIOSVlanrangeModel](ctx, m.NIOS, diags)
	niosModel.Flatten(ctx, resp.NIOS, diags)
	if resp.NIOS != nil {
		PostFlattenVlanrangeNIOS(ctx, plannedNIOS, niosModel, diags)
		m.NIOS = flex.FlattenNestedObject(ctx, niosModel, NIOSVlanrangeAttrTypes, diags)
	} else {
		m.NIOS = types.ObjectNull(NIOSVlanrangeAttrTypes)
	}

}

// Flatten merges API response onto existing NIOS model.
func (m *NIOSVlanrangeModel) Flatten(ctx context.Context, from *coremodel.NIOSVlanrangeExt, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	planExtAttrs := m.ExtAttrs
	if planExtAttrs.IsUnknown() {
		planExtAttrs = types.MapNull(types.StringType)
	}
	m.Comment = flex.FlattenStringPointerEmptyAsNull(from.Comment)
	m.DeleteVlans = flex.FlattenBoolPointer(from.DeleteVlans)
	m.EndVlanId = flex.FlattenInt64Pointer(from.EndVlanId)
	m.ExtAttrs, m.ExtAttrsAll = flex.FlattenEAs(planExtAttrs, from.ExtAttrs)
	m.Name = flex.FlattenStringPointerEmptyAsNull(from.Name)
	m.PreCreateVlan = flex.FlattenBoolPointer(from.PreCreateVlan)
	m.StartVlanId = flex.FlattenInt64Pointer(from.StartVlanId)
	m.VlanNamePrefix = flex.FlattenStringPointerEmptyAsNull(from.VlanNamePrefix)
	m.VlanView = FlattenVlanrangeVlanView(from.VlanView)
}
