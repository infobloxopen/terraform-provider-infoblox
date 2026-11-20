package dhcp

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	niosdhcp "github.com/infobloxopen/infoblox-nios-go-client/dhcp"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
)

// Ipv6rangePortControlBlackoutSettingModel is the Terraform model for Ipv6rangePortControlBlackoutSetting
type Ipv6rangePortControlBlackoutSettingModel struct {
	EnableBlackout   types.Bool   `tfsdk:"enable_blackout"`
	BlackoutDuration types.Int64  `tfsdk:"blackout_duration"`
	BlackoutSchedule types.Object `tfsdk:"blackout_schedule"`
}

// Ipv6rangePortControlBlackoutSettingAttrTypes contains the attribute types for Ipv6rangePortControlBlackoutSettingModel
var Ipv6rangePortControlBlackoutSettingAttrTypes = map[string]attr.Type{
	"enable_blackout":   types.BoolType,
	"blackout_duration": types.Int64Type,
	"blackout_schedule": types.ObjectType{AttrTypes: Ipv6rangeportcontrolblackoutsettingBlackoutScheduleAttrTypes},
}

// Ipv6rangePortControlBlackoutSettingResourceSchemaAttributes contains the schema attributes for Ipv6rangePortControlBlackoutSettingModel
var Ipv6rangePortControlBlackoutSettingResourceSchemaAttributes = map[string]schema.Attribute{
	"enable_blackout": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Determines whether a blackout is enabled or not.",
	},
	"blackout_duration": schema.Int64Attribute{
		Optional:            true,
		MarkdownDescription: "The blackout duration in seconds; minimum value is 1 minute.",
	},
	"blackout_schedule": schema.SingleNestedAttribute{
		Attributes:          Ipv6rangeportcontrolblackoutsettingBlackoutScheduleResourceSchemaAttributes,
		Optional:            true,
		MarkdownDescription: "A Schedule Setting struct that determines blackout schedule.",
	},
}

// ExpandIpv6rangePortControlBlackoutSetting converts a Terraform Object to SDK type
func ExpandIpv6rangePortControlBlackoutSetting(ctx context.Context, o types.Object, diags *diag.Diagnostics) *niosdhcp.Ipv6rangePortControlBlackoutSetting {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m Ipv6rangePortControlBlackoutSettingModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

// Expand converts the Terraform model to SDK type
func (m *Ipv6rangePortControlBlackoutSettingModel) Expand(ctx context.Context, diags *diag.Diagnostics) *niosdhcp.Ipv6rangePortControlBlackoutSetting {
	if m == nil {
		return nil
	}
	to := &niosdhcp.Ipv6rangePortControlBlackoutSetting{
		EnableBlackout:   flex.ExpandBoolPointer(m.EnableBlackout),
		BlackoutDuration: flex.ExpandInt64Pointer(m.BlackoutDuration),
		BlackoutSchedule: ExpandIpv6rangeportcontrolblackoutsettingBlackoutSchedule(ctx, m.BlackoutSchedule, diags),
	}
	return to
}

// FlattenIpv6rangePortControlBlackoutSetting converts an SDK type to Terraform Object
func FlattenIpv6rangePortControlBlackoutSetting(ctx context.Context, from *niosdhcp.Ipv6rangePortControlBlackoutSetting, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(Ipv6rangePortControlBlackoutSettingAttrTypes)
	}
	m := &Ipv6rangePortControlBlackoutSettingModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, Ipv6rangePortControlBlackoutSettingAttrTypes, m)
	diags.Append(d...)
	return t
}

// Flatten populates the Terraform model from SDK type
func (m *Ipv6rangePortControlBlackoutSettingModel) Flatten(ctx context.Context, from *niosdhcp.Ipv6rangePortControlBlackoutSetting, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.EnableBlackout = flex.FlattenBoolPointer(from.EnableBlackout)
	m.BlackoutDuration = flex.FlattenInt64Pointer(from.BlackoutDuration)
	m.BlackoutSchedule = FlattenIpv6rangeportcontrolblackoutsettingBlackoutSchedule(ctx, from.BlackoutSchedule, diags)
}
