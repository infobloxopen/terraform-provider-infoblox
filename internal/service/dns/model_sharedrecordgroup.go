package dns

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/mapvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/mapdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"

	coremodel "github.com/infobloxopen/terraform-provider-infoblox/internal/core/model/dns"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
	importmod "github.com/infobloxopen/terraform-provider-infoblox/internal/planmodifiers/import"
	customvalidator "github.com/infobloxopen/terraform-provider-infoblox/internal/validator"
)

type SharedrecordgroupModel struct {
	Id   types.String `tfsdk:"id"`
	NIOS types.Object `tfsdk:"nios"`
}

var SharedrecordgroupAttrTypes = map[string]attr.Type{
	"id":   types.StringType,
	"nios": types.ObjectType{AttrTypes: NIOSSharedrecordgroupAttrTypes},
}

type NIOSSharedrecordgroupModel struct {
	Comment          types.String `tfsdk:"comment"`
	ExtAttrs         types.Map    `tfsdk:"ext_attrs"`
	ExtAttrsAll      types.Map    `tfsdk:"ext_attrs_all"`
	Name             types.String `tfsdk:"name"`
	RecordNamePolicy types.String `tfsdk:"record_name_policy"`
	ZoneAssociations types.List   `tfsdk:"zone_associations"`
}

var NIOSSharedrecordgroupAttrTypes = map[string]attr.Type{
	"comment":            types.StringType,
	"ext_attrs":          types.MapType{ElemType: types.StringType},
	"ext_attrs_all":      types.MapType{ElemType: types.StringType},
	"name":               types.StringType,
	"record_name_policy": types.StringType,
	"zone_associations":  types.ListType{ElemType: types.ObjectType{AttrTypes: SharedrecordgroupZoneAssociationsAttrTypes}},
}

const (
	SharedrecordgroupReturnFields = "comment,extattrs,name,record_name_policy,use_record_name_policy,zone_associations"
)

var SharedrecordgroupResourceSchemaAttributes = map[string]schema.Attribute{
	"id": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The reference to the object.",
	},
	"nios": schema.SingleNestedAttribute{
		Optional:            true,
		MarkdownDescription: "NIOS backend-specific fields.",
		Attributes:          SharedrecordgroupResourceNiosSchemaAttributes,
	},
}

var SharedrecordgroupResourceNiosSchemaAttributes = map[string]schema.Attribute{
	"comment": schema.StringAttribute{
		Optional: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
			stringvalidator.LengthBetween(0, 256),
		},
		MarkdownDescription: "The descriptive comment of this shared record group.",
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
		MarkdownDescription: "The name of this shared record group.",
	},
	"record_name_policy": schema.StringAttribute{
		Optional: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "The record name policy of this shared record group.",
	},
	"zone_associations": schema.ListNestedAttribute{
		NestedObject: schema.NestedAttributeObject{
			Attributes: SharedrecordgroupZoneAssociationsResourceSchemaAttributes,
		},
		Optional: true,
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
		},
		MarkdownDescription: "The list of zones associated with this shared record group. Starting from NIOS-9.0.6, this field has been updated to a structure that includes FQDN and DNS view details.",
	},
}

// Expand converts the TF model to the infoblox core model
func (m *SharedrecordgroupModel) Expand(ctx context.Context, diags *diag.Diagnostics, isCreate bool) *coremodel.Sharedrecordgroup {
	if m == nil {
		return nil
	}

	obj := &coremodel.Sharedrecordgroup{}

	// Expand NIOS nested attribute (returns nil if not present)
	niosModel := flex.ExpandNestedObject[NIOSSharedrecordgroupModel](ctx, m.NIOS, diags)
	if niosModel != nil {
		obj.NIOS = niosModel.Expand(ctx, diags)
	}

	return obj
}

// Expand converts the NIOS TF model to the core model.
func (m *NIOSSharedrecordgroupModel) Expand(ctx context.Context, diags *diag.Diagnostics) *coremodel.NIOSSharedrecordgroupExt {
	return &coremodel.NIOSSharedrecordgroupExt{
		Comment:          flex.ExpandStringPointerNullAsEmpty(m.Comment),
		ExtAttrs:         flex.ExpandMapStringAny(ctx, m.ExtAttrs, diags),
		Name:             flex.ExpandStringPointerNullAsEmpty(m.Name),
		RecordNamePolicy: flex.ExpandStringPointer(m.RecordNamePolicy),
		ZoneAssociations: flex.ExpandFrameworkListNestedBlock(ctx, m.ZoneAssociations, diags, ExpandSharedrecordgroupZoneAssociations),
	}
}

// ApplySharedrecordgroupNIOSUseFlags derives NIOS use flags from the raw config
// value(s) and writes them onto the core model. A flag is true when the user
// set any of its governed value fields in config.
func ApplySharedrecordgroupNIOSUseFlags(ctx context.Context, config tfsdk.Config, obj *coremodel.Sharedrecordgroup, diags *diag.Diagnostics) {
	if obj == nil || obj.NIOS == nil {
		return
	}
	obj.NIOS.UseRecordNamePolicy = flex.DeriveUseFlag(ctx, config, diags, path.Root("nios").AtName("record_name_policy"))
}

// Flatten populates the TF model from a core response.
func (m *SharedrecordgroupModel) Flatten(ctx context.Context, resp *coremodel.Sharedrecordgroup, diags *diag.Diagnostics) {
	if resp == nil {
		return
	}

	m.Id = flex.FlattenStringPointer(resp.Id)

	// Extract existing NIOS model, flatten API response onto it, convert back
	niosModel := flex.ExpandNestedObject[NIOSSharedrecordgroupModel](ctx, m.NIOS, diags)
	if niosModel == nil {
		niosModel = &NIOSSharedrecordgroupModel{}
	}
	niosModel.Flatten(ctx, resp.NIOS, diags)
	if resp.NIOS != nil {
		m.NIOS = flex.FlattenNestedObject(ctx, niosModel, NIOSSharedrecordgroupAttrTypes, diags)
	} else {
		m.NIOS = types.ObjectNull(NIOSSharedrecordgroupAttrTypes)
	}

}

// Flatten merges API response onto existing NIOS model.
func (m *NIOSSharedrecordgroupModel) Flatten(ctx context.Context, from *coremodel.NIOSSharedrecordgroupExt, diags *diag.Diagnostics) {
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
	m.RecordNamePolicy = flex.FlattenStringPointerEmptyAsNull(from.RecordNamePolicy)
	m.ZoneAssociations = flex.FlattenFrameworkListNestedBlock(ctx, from.ZoneAssociations, SharedrecordgroupZoneAssociationsAttrTypes, diags, FlattenSharedrecordgroupZoneAssociations)
}
