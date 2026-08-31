package grid

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

	coremodel "github.com/infobloxopen/terraform-provider-infoblox/internal/core/model/grid"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
	importmod "github.com/infobloxopen/terraform-provider-infoblox/internal/planmodifiers/import"
	customvalidator "github.com/infobloxopen/terraform-provider-infoblox/internal/validator"
)

type ServicerestartGroupModel struct {
	Id   types.String `tfsdk:"id"`
	NIOS types.Object `tfsdk:"nios"`
}

var ServicerestartGroupAttrTypes = map[string]attr.Type{
	"id":   types.StringType,
	"nios": types.ObjectType{AttrTypes: NIOSServicerestartGroupAttrTypes},
}

type NIOSServicerestartGroupModel struct {
	Comment           types.String `tfsdk:"comment"`
	ExtAttrs          types.Map    `tfsdk:"ext_attrs"`
	ExtAttrsAll       types.Map    `tfsdk:"ext_attrs_all"`
	Members           types.List   `tfsdk:"members"`
	Mode              types.String `tfsdk:"mode"`
	Name              types.String `tfsdk:"name"`
	RecurringSchedule types.Object `tfsdk:"recurring_schedule"`
	Service           types.String `tfsdk:"service"`
}

var NIOSServicerestartGroupAttrTypes = map[string]attr.Type{
	"comment":            types.StringType,
	"ext_attrs":          types.MapType{ElemType: types.StringType},
	"ext_attrs_all":      types.MapType{ElemType: types.StringType},
	"members":            types.ListType{ElemType: types.StringType},
	"mode":               types.StringType,
	"name":               types.StringType,
	"recurring_schedule": types.ObjectType{AttrTypes: GridServicerestartGroupRecurringScheduleAttrTypes},
	"service":            types.StringType,
}

const (
	ServicerestartGroupReturnFields = "comment,extattrs,is_default,last_updated_time,members,mode,name,position,recurring_schedule,requests,service,status"
)

var ServicerestartGroupResourceSchemaAttributes = map[string]schema.Attribute{
	"id": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "The reference to the object.",
	},
	"nios": schema.SingleNestedAttribute{
		Optional:            true,
		MarkdownDescription: "NIOS backend-specific fields.",
		Attributes:          ServicerestartGroupResourceNiosSchemaAttributes,
	},
}

var ServicerestartGroupResourceNiosSchemaAttributes = map[string]schema.Attribute{
	"comment": schema.StringAttribute{
		Optional: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
			customvalidator.ValidateTrimmedString(),
			stringvalidator.LengthBetween(0, 256),
		},
		MarkdownDescription: "Comment for the Restart Group; maximum 256 characters.",
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
	"members": schema.ListAttribute{
		ElementType: types.StringType,
		Optional:    true,
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
		},
		MarkdownDescription: "The list of members belonging to the group.",
	},
	"mode": schema.StringAttribute{
		Default: stringdefault.StaticString("SIMULTANEOUS"),
		Validators: []validator.String{
			stringvalidator.OneOf("SEQUENTIAL", "SIMULTANEOUS"),
		},
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The default restart method for this Restart Group.",
	},
	"name": schema.StringAttribute{
		Required: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
			customvalidator.ValidateTrimmedString(),
		},
		MarkdownDescription: "The name of this Restart Group.",
	},
	"recurring_schedule": schema.SingleNestedAttribute{
		Attributes:          GridServicerestartGroupRecurringScheduleResourceSchemaAttributes,
		Optional:            true,
		MarkdownDescription: "The recurring schedule for restart of a group.",
	},
	"service": schema.StringAttribute{
		Validators: []validator.String{
			stringvalidator.OneOf("DHCP", "DNS"),
		},
		Required:            true,
		MarkdownDescription: "The applicable service for this Restart Group.",
	},
}

// Expand converts the TF model to the infoblox core model
func (m *ServicerestartGroupModel) Expand(ctx context.Context, diags *diag.Diagnostics, isCreate bool) *coremodel.ServicerestartGroup {
	if m == nil {
		return nil
	}

	obj := &coremodel.ServicerestartGroup{}

	// Expand NIOS nested attribute (returns nil if not present)
	niosModel := flex.ExpandNestedObject[NIOSServicerestartGroupModel](ctx, m.NIOS, diags)
	if niosModel != nil {
		obj.NIOS = niosModel.Expand(ctx, diags)
	}

	return obj
}

// Expand converts the NIOS TF model to the core model.
func (m *NIOSServicerestartGroupModel) Expand(ctx context.Context, diags *diag.Diagnostics) *coremodel.NIOSServicerestartGroupExt {
	return &coremodel.NIOSServicerestartGroupExt{
		Comment:           flex.ExpandStringPointerNullAsEmpty(m.Comment),
		ExtAttrs:          flex.ExpandMapStringAny(ctx, m.ExtAttrs, diags),
		Members:           flex.ExpandFrameworkListString(ctx, m.Members, diags),
		Mode:              flex.ExpandStringPointerNullAsEmpty(m.Mode),
		Name:              flex.ExpandStringPointerNullAsEmpty(m.Name),
		RecurringSchedule: ExpandGridServicerestartGroupRecurringSchedule(ctx, m.RecurringSchedule, diags),
		Service:           flex.ExpandStringPointerNullAsEmpty(m.Service),
	}
}

// Flatten populates the TF model from a core response.
func (m *ServicerestartGroupModel) Flatten(ctx context.Context, resp *coremodel.ServicerestartGroup, diags *diag.Diagnostics) {
	if resp == nil {
		return
	}

	m.Id = flex.FlattenStringPointer(resp.Id)

	// Extract existing NIOS model, flatten API response onto it, convert back
	niosModel := flex.ExpandNestedObject[NIOSServicerestartGroupModel](ctx, m.NIOS, diags)
	if niosModel == nil {
		niosModel = &NIOSServicerestartGroupModel{}
	}
	niosModel.Flatten(ctx, resp.NIOS, diags)
	if resp.NIOS != nil {
		m.NIOS = flex.FlattenNestedObject(ctx, niosModel, NIOSServicerestartGroupAttrTypes, diags)
	} else {
		m.NIOS = types.ObjectNull(NIOSServicerestartGroupAttrTypes)
	}

}

// Flatten merges API response onto existing NIOS model.
func (m *NIOSServicerestartGroupModel) Flatten(ctx context.Context, from *coremodel.NIOSServicerestartGroupExt, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	planExtAttrs := m.ExtAttrs
	if planExtAttrs.IsUnknown() {
		planExtAttrs = types.MapNull(types.StringType)
	}
	m.Comment = flex.FlattenStringPointerEmptyAsNull(from.Comment)
	m.ExtAttrs, m.ExtAttrsAll = flex.FlattenEAs(planExtAttrs, from.ExtAttrs)
	m.Members = flex.FlattenFrameworkListString(ctx, from.Members, diags)
	m.Mode = flex.FlattenStringPointerEmptyAsNull(from.Mode)
	m.Name = flex.FlattenStringPointerEmptyAsNull(from.Name)
	m.RecurringSchedule = FlattenGridServicerestartGroupRecurringSchedule(ctx, from.RecurringSchedule, diags)
	m.Service = flex.FlattenStringPointerEmptyAsNull(from.Service)
}
