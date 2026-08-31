package grid

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	niosgrid "github.com/infobloxopen/infoblox-nios-go-client/grid"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
	internaltypes "github.com/infobloxopen/terraform-provider-infoblox/internal/types"
	customvalidator "github.com/infobloxopen/terraform-provider-infoblox/internal/validator"
)

// GridServicerestartGroupRecurringScheduleModel is the Terraform model for GridServicerestartGroupRecurringSchedule
type GridServicerestartGroupRecurringScheduleModel struct {
	Services internaltypes.UnorderedListValue `tfsdk:"services"`
	Mode     types.String                     `tfsdk:"mode"`
	Schedule types.Object                     `tfsdk:"schedule"`
	Force    types.Bool                       `tfsdk:"force"`
}

// GridServicerestartGroupRecurringScheduleAttrTypes contains the attribute types for GridServicerestartGroupRecurringScheduleModel
var GridServicerestartGroupRecurringScheduleAttrTypes = map[string]attr.Type{
	"services": internaltypes.UnorderedListOfStringType,
	"mode":     types.StringType,
	"schedule": types.ObjectType{AttrTypes: GridservicerestartgrouprecurringscheduleScheduleAttrTypes},
	"force":    types.BoolType,
}

// GridServicerestartGroupRecurringScheduleResourceSchemaAttributes contains the schema attributes for GridServicerestartGroupRecurringScheduleModel
var GridServicerestartGroupRecurringScheduleResourceSchemaAttributes = map[string]schema.Attribute{
	"services": schema.ListAttribute{
		ElementType: types.StringType,
		Optional:    true,
		Computed:    true,
		CustomType:  internaltypes.UnorderedListOfStringType,
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
			customvalidator.StringsInSlice([]string{"ALL", "DHCP", "DHCPV4", "DHCPV6", "DNS"}),
		},
		MarkdownDescription: "The list of applicable services for the restart.",
	},
	"mode": schema.StringAttribute{
		Validators: []validator.String{
			stringvalidator.OneOf("GROUPED", "SEQUENTIAL", "SIMULTANEOUS"),
		},
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The restart method for a Grid restart.",
	},
	"schedule": schema.SingleNestedAttribute{
		Attributes:          GridservicerestartgrouprecurringscheduleScheduleResourceSchemaAttributes,
		Optional:            true,
		MarkdownDescription: "The Schedule Setting struct that determines the schedule for the restart.",
	},
	"force": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Determines if the Restart Group should have a force restart.",
	},
}

// ExpandGridServicerestartGroupRecurringSchedule converts a Terraform Object to SDK type
func ExpandGridServicerestartGroupRecurringSchedule(ctx context.Context, o types.Object, diags *diag.Diagnostics) *niosgrid.GridServicerestartGroupRecurringSchedule {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m GridServicerestartGroupRecurringScheduleModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

// Expand converts the Terraform model to SDK type
func (m *GridServicerestartGroupRecurringScheduleModel) Expand(ctx context.Context, diags *diag.Diagnostics) *niosgrid.GridServicerestartGroupRecurringSchedule {
	if m == nil {
		return nil
	}
	to := &niosgrid.GridServicerestartGroupRecurringSchedule{
		Services: flex.ExpandFrameworkListString(ctx, m.Services, diags),
		Mode:     flex.ExpandStringPointer(m.Mode),
		Schedule: ExpandGridservicerestartgrouprecurringscheduleSchedule(ctx, m.Schedule, diags),
		Force:    flex.ExpandBoolPointer(m.Force),
	}
	return to
}

// FlattenGridServicerestartGroupRecurringSchedule converts an SDK type to Terraform Object
func FlattenGridServicerestartGroupRecurringSchedule(ctx context.Context, from *niosgrid.GridServicerestartGroupRecurringSchedule, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(GridServicerestartGroupRecurringScheduleAttrTypes)
	}
	m := &GridServicerestartGroupRecurringScheduleModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, GridServicerestartGroupRecurringScheduleAttrTypes, m)
	diags.Append(d...)
	return t
}

// Flatten populates the Terraform model from SDK type
func (m *GridServicerestartGroupRecurringScheduleModel) Flatten(ctx context.Context, from *niosgrid.GridServicerestartGroupRecurringSchedule, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.Services = flex.FlattenFrameworkUnorderedListString(ctx, from.Services, diags)
	m.Mode = flex.FlattenStringPointerEmptyAsNull(from.Mode)
	m.Schedule = FlattenGridservicerestartgrouprecurringscheduleSchedule(ctx, from.Schedule, diags)
	m.Force = flex.FlattenBoolPointer(from.Force)
}
