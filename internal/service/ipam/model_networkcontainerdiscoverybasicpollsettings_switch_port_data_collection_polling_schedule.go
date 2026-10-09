package ipam

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	niosipam "github.com/infobloxopen/infoblox-nios-go-client/ipam"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
	customvalidator "github.com/infobloxopen/terraform-provider-infoblox/internal/validator"
)

// NetworkcontainerdiscoverybasicpollsettingsSwitchPortDataCollectionPollingScheduleModel is the Terraform model for NetworkcontainerdiscoverybasicpollsettingsSwitchPortDataCollectionPollingSchedule
type NetworkcontainerdiscoverybasicpollsettingsSwitchPortDataCollectionPollingScheduleModel struct {
	Weekdays        types.List   `tfsdk:"weekdays"`
	TimeZone        types.String `tfsdk:"time_zone"`
	RecurringTime   types.Int64  `tfsdk:"recurring_time"`
	Frequency       types.String `tfsdk:"frequency"`
	Every           types.Int64  `tfsdk:"every"`
	MinutesPastHour types.Int64  `tfsdk:"minutes_past_hour"`
	HourOfDay       types.Int64  `tfsdk:"hour_of_day"`
	Year            types.Int64  `tfsdk:"year"`
	Month           types.Int64  `tfsdk:"month"`
	DayOfMonth      types.Int64  `tfsdk:"day_of_month"`
	Repeat          types.String `tfsdk:"repeat"`
	Disable         types.Bool   `tfsdk:"disable"`
}

// NetworkcontainerdiscoverybasicpollsettingsSwitchPortDataCollectionPollingScheduleAttrTypes contains the attribute types for NetworkcontainerdiscoverybasicpollsettingsSwitchPortDataCollectionPollingScheduleModel
var NetworkcontainerdiscoverybasicpollsettingsSwitchPortDataCollectionPollingScheduleAttrTypes = map[string]attr.Type{
	"weekdays":          types.ListType{ElemType: types.StringType},
	"time_zone":         types.StringType,
	"recurring_time":    types.Int64Type,
	"frequency":         types.StringType,
	"every":             types.Int64Type,
	"minutes_past_hour": types.Int64Type,
	"hour_of_day":       types.Int64Type,
	"year":              types.Int64Type,
	"month":             types.Int64Type,
	"day_of_month":      types.Int64Type,
	"repeat":            types.StringType,
	"disable":           types.BoolType,
}

// NetworkcontainerdiscoverybasicpollsettingsSwitchPortDataCollectionPollingScheduleResourceSchemaAttributes contains the schema attributes for NetworkcontainerdiscoverybasicpollsettingsSwitchPortDataCollectionPollingScheduleModel
var NetworkcontainerdiscoverybasicpollsettingsSwitchPortDataCollectionPollingScheduleResourceSchemaAttributes = map[string]schema.Attribute{
	"weekdays": schema.ListAttribute{
		ElementType: types.StringType,
		Optional:    true,
		Validators: []validator.List{
			customvalidator.ListNotEmpty(),
		},
		MarkdownDescription: "Days of the week when scheduling is triggered.",
	},
	"time_zone": schema.StringAttribute{
		Optional: true,
		Validators: []validator.String{
			customvalidator.StringNotEmpty(),
		},
		MarkdownDescription: "The time zone for the schedule.",
	},
	"recurring_time": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "The recurring time for the schedule in Epoch seconds format. This field is obsolete and is preserved only for backward compatibility purposes. Please use other applicable fields to define the recurring schedule. DO NOT use recurring_time together with these fields. If you use recurring_time with other fields to define the recurring schedule, recurring_time has priority over year, hour_of_day, and minutes_past_hour and will override the values of these fields, although it does not override month and day_of_month. In this case, the recurring time value might be different than the intended value that you define.",
	},
	"frequency": schema.StringAttribute{
		Validators: []validator.String{
			stringvalidator.OneOf("DAILY", "HOURLY", "MONTHLY", "WEEKLY"),
		},
		Optional:            true,
		MarkdownDescription: "The frequency for the scheduled task.",
	},
	"every": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "The number of frequency to wait before repeating the scheduled task.",
	},
	"minutes_past_hour": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "The minutes past the hour for the scheduled task.",
	},
	"hour_of_day": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "The hour of day for the scheduled task.",
	},
	"year": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "The year for the scheduled task.",
	},
	"month": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "The month for the scheduled task.",
	},
	"day_of_month": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "The day of the month for the scheduled task.",
	},
	"repeat": schema.StringAttribute{
		Validators: []validator.String{
			stringvalidator.OneOf("ONCE", "RECUR"),
		},
		Optional:            true,
		MarkdownDescription: "Indicates if the scheduled task will be repeated or run only once.",
	},
	"disable": schema.BoolAttribute{
		Optional:            true,
		MarkdownDescription: "If set to True, the scheduled task is disabled.",
	},
}

// ExpandNetworkcontainerdiscoverybasicpollsettingsSwitchPortDataCollectionPollingSchedule converts a Terraform Object to SDK type
func ExpandNetworkcontainerdiscoverybasicpollsettingsSwitchPortDataCollectionPollingSchedule(ctx context.Context, o types.Object, diags *diag.Diagnostics) *niosipam.NetworkcontainerdiscoverybasicpollsettingsSwitchPortDataCollectionPollingSchedule {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m NetworkcontainerdiscoverybasicpollsettingsSwitchPortDataCollectionPollingScheduleModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

// Expand converts the Terraform model to SDK type
func (m *NetworkcontainerdiscoverybasicpollsettingsSwitchPortDataCollectionPollingScheduleModel) Expand(ctx context.Context, diags *diag.Diagnostics) *niosipam.NetworkcontainerdiscoverybasicpollsettingsSwitchPortDataCollectionPollingSchedule {
	if m == nil {
		return nil
	}
	to := &niosipam.NetworkcontainerdiscoverybasicpollsettingsSwitchPortDataCollectionPollingSchedule{
		Weekdays:        flex.ExpandFrameworkListString(ctx, m.Weekdays, diags),
		TimeZone:        flex.ExpandStringPointerNullAsEmpty(m.TimeZone),
		RecurringTime:   flex.ExpandInt64Pointer(m.RecurringTime),
		Frequency:       flex.ExpandStringPointerNullAsEmpty(m.Frequency),
		Every:           flex.ExpandInt64Pointer(m.Every),
		MinutesPastHour: flex.ExpandInt64Pointer(m.MinutesPastHour),
		HourOfDay:       flex.ExpandInt64Pointer(m.HourOfDay),
		Year:            flex.ExpandInt64Pointer(m.Year),
		Month:           flex.ExpandInt64Pointer(m.Month),
		DayOfMonth:      flex.ExpandInt64Pointer(m.DayOfMonth),
		Repeat:          flex.ExpandStringPointerNullAsEmpty(m.Repeat),
		Disable:         flex.ExpandBoolPointer(m.Disable),
	}
	return to
}

// FlattenNetworkcontainerdiscoverybasicpollsettingsSwitchPortDataCollectionPollingSchedule converts an SDK type to Terraform Object
func FlattenNetworkcontainerdiscoverybasicpollsettingsSwitchPortDataCollectionPollingSchedule(ctx context.Context, from *niosipam.NetworkcontainerdiscoverybasicpollsettingsSwitchPortDataCollectionPollingSchedule, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(NetworkcontainerdiscoverybasicpollsettingsSwitchPortDataCollectionPollingScheduleAttrTypes)
	}
	m := &NetworkcontainerdiscoverybasicpollsettingsSwitchPortDataCollectionPollingScheduleModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, NetworkcontainerdiscoverybasicpollsettingsSwitchPortDataCollectionPollingScheduleAttrTypes, m)
	diags.Append(d...)
	return t
}

// Flatten populates the Terraform model from SDK type
func (m *NetworkcontainerdiscoverybasicpollsettingsSwitchPortDataCollectionPollingScheduleModel) Flatten(ctx context.Context, from *niosipam.NetworkcontainerdiscoverybasicpollsettingsSwitchPortDataCollectionPollingSchedule, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.Weekdays = flex.FlattenFrameworkListString(ctx, from.Weekdays, diags)
	m.TimeZone = flex.FlattenStringPointerEmptyAsNull(from.TimeZone)
	m.RecurringTime = flex.FlattenInt64Pointer(from.RecurringTime)
	m.Frequency = flex.FlattenStringPointerEmptyAsNull(from.Frequency)
	m.Every = flex.FlattenInt64Pointer(from.Every)
	m.MinutesPastHour = flex.FlattenInt64Pointer(from.MinutesPastHour)
	m.HourOfDay = flex.FlattenInt64Pointer(from.HourOfDay)
	m.Year = flex.FlattenInt64Pointer(from.Year)
	m.Month = flex.FlattenInt64Pointer(from.Month)
	m.DayOfMonth = flex.FlattenInt64Pointer(from.DayOfMonth)
	m.Repeat = flex.FlattenStringPointerEmptyAsNull(from.Repeat)
	m.Disable = flex.FlattenBoolPointer(from.Disable)
}
