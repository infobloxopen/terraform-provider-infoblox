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

// RangeDiscoveryBlackoutSettingModel is the Terraform model for RangeDiscoveryBlackoutSetting
type RangeDiscoveryBlackoutSettingModel struct {
	EnableBlackout   types.Bool   `tfsdk:"enable_blackout"`
	BlackoutDuration types.Int64  `tfsdk:"blackout_duration"`
	BlackoutSchedule types.Object `tfsdk:"blackout_schedule"`
}

// RangeDiscoveryBlackoutSettingAttrTypes contains the attribute types for RangeDiscoveryBlackoutSettingModel
var RangeDiscoveryBlackoutSettingAttrTypes = map[string]attr.Type{
	"enable_blackout":   types.BoolType,
	"blackout_duration": types.Int64Type,
	"blackout_schedule": types.ObjectType{AttrTypes: RangediscoveryblackoutsettingBlackoutScheduleAttrTypes},
}

// RangeDiscoveryBlackoutSettingResourceSchemaAttributes contains the schema attributes for RangeDiscoveryBlackoutSettingModel
var RangeDiscoveryBlackoutSettingResourceSchemaAttributes = map[string]schema.Attribute{
	"enable_blackout": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Determines whether a blackout is enabled or not.",
	},
	"blackout_duration": schema.Int64Attribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The blackout duration in seconds; minimum value is 1 minute.",
	},
	"blackout_schedule": schema.SingleNestedAttribute{
		Attributes:          RangediscoveryblackoutsettingBlackoutScheduleResourceSchemaAttributes,
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "A Schedule Setting struct that determines blackout schedule.",
	},
}

// ExpandRangeDiscoveryBlackoutSetting converts a Terraform Object to SDK type
func ExpandRangeDiscoveryBlackoutSetting(ctx context.Context, o types.Object, diags *diag.Diagnostics) *niosdhcp.RangeDiscoveryBlackoutSetting {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m RangeDiscoveryBlackoutSettingModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

// Expand converts the Terraform model to SDK type
func (m *RangeDiscoveryBlackoutSettingModel) Expand(ctx context.Context, diags *diag.Diagnostics) *niosdhcp.RangeDiscoveryBlackoutSetting {
	if m == nil {
		return nil
	}
	to := &niosdhcp.RangeDiscoveryBlackoutSetting{
		EnableBlackout:   flex.ExpandBoolPointer(m.EnableBlackout),
		BlackoutDuration: flex.ExpandInt64Pointer(m.BlackoutDuration),
		BlackoutSchedule: ExpandRangediscoveryblackoutsettingBlackoutSchedule(ctx, m.BlackoutSchedule, diags),
	}
	return to
}

// FlattenRangeDiscoveryBlackoutSetting converts an SDK type to Terraform Object
func FlattenRangeDiscoveryBlackoutSetting(ctx context.Context, from *niosdhcp.RangeDiscoveryBlackoutSetting, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(RangeDiscoveryBlackoutSettingAttrTypes)
	}
	m := &RangeDiscoveryBlackoutSettingModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, RangeDiscoveryBlackoutSettingAttrTypes, m)
	diags.Append(d...)
	return t
}

// Flatten populates the Terraform model from SDK type
func (m *RangeDiscoveryBlackoutSettingModel) Flatten(ctx context.Context, from *niosdhcp.RangeDiscoveryBlackoutSetting, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.EnableBlackout = flex.FlattenBoolPointer(from.EnableBlackout)
	m.BlackoutDuration = flex.FlattenInt64Pointer(from.BlackoutDuration)
	m.BlackoutSchedule = FlattenRangediscoveryblackoutsettingBlackoutSchedule(ctx, from.BlackoutSchedule, diags)
}
