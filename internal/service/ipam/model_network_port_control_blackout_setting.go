package ipam

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	niosipam "github.com/infobloxopen/infoblox-nios-go-client/ipam"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
)

// NetworkPortControlBlackoutSettingModel is the Terraform model for NetworkPortControlBlackoutSetting
type NetworkPortControlBlackoutSettingModel struct {
	EnableBlackout   types.Bool   `tfsdk:"enable_blackout"`
	BlackoutDuration types.Int64  `tfsdk:"blackout_duration"`
	BlackoutSchedule types.Object `tfsdk:"blackout_schedule"`
}

// NetworkPortControlBlackoutSettingAttrTypes contains the attribute types for NetworkPortControlBlackoutSettingModel
var NetworkPortControlBlackoutSettingAttrTypes = map[string]attr.Type{
	"enable_blackout":   types.BoolType,
	"blackout_duration": types.Int64Type,
	"blackout_schedule": types.ObjectType{AttrTypes: NetworkportcontrolblackoutsettingBlackoutScheduleAttrTypes},
}

// NetworkPortControlBlackoutSettingResourceSchemaAttributes contains the schema attributes for NetworkPortControlBlackoutSettingModel
var NetworkPortControlBlackoutSettingResourceSchemaAttributes = map[string]schema.Attribute{
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
		Attributes:          NetworkportcontrolblackoutsettingBlackoutScheduleResourceSchemaAttributes,
		Optional:            true,
		MarkdownDescription: "A Schedule Setting struct that determines blackout schedule.",
	},
}

// ExpandNetworkPortControlBlackoutSetting converts a Terraform Object to SDK type
func ExpandNetworkPortControlBlackoutSetting(ctx context.Context, o types.Object, diags *diag.Diagnostics) *niosipam.NetworkPortControlBlackoutSetting {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m NetworkPortControlBlackoutSettingModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

// Expand converts the Terraform model to SDK type
func (m *NetworkPortControlBlackoutSettingModel) Expand(ctx context.Context, diags *diag.Diagnostics) *niosipam.NetworkPortControlBlackoutSetting {
	if m == nil {
		return nil
	}
	to := &niosipam.NetworkPortControlBlackoutSetting{
		EnableBlackout:   flex.ExpandBoolPointer(m.EnableBlackout),
		BlackoutDuration: flex.ExpandInt64Pointer(m.BlackoutDuration),
		BlackoutSchedule: ExpandNetworkportcontrolblackoutsettingBlackoutSchedule(ctx, m.BlackoutSchedule, diags),
	}
	return to
}

// FlattenNetworkPortControlBlackoutSetting converts an SDK type to Terraform Object
func FlattenNetworkPortControlBlackoutSetting(ctx context.Context, from *niosipam.NetworkPortControlBlackoutSetting, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(NetworkPortControlBlackoutSettingAttrTypes)
	}
	m := &NetworkPortControlBlackoutSettingModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, NetworkPortControlBlackoutSettingAttrTypes, m)
	diags.Append(d...)
	return t
}

// Flatten populates the Terraform model from SDK type
func (m *NetworkPortControlBlackoutSettingModel) Flatten(ctx context.Context, from *niosipam.NetworkPortControlBlackoutSetting, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.EnableBlackout = flex.FlattenBoolPointer(from.EnableBlackout)
	m.BlackoutDuration = flex.FlattenInt64Pointer(from.BlackoutDuration)
	m.BlackoutSchedule = FlattenNetworkportcontrolblackoutsettingBlackoutSchedule(ctx, from.BlackoutSchedule, diags)
}
