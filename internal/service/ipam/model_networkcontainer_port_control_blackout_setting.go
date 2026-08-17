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

// NetworkcontainerPortControlBlackoutSettingModel is the Terraform model for NetworkcontainerPortControlBlackoutSetting
type NetworkcontainerPortControlBlackoutSettingModel struct {
	EnableBlackout   types.Bool   `tfsdk:"enable_blackout"`
	BlackoutDuration types.Int64  `tfsdk:"blackout_duration"`
	BlackoutSchedule types.Object `tfsdk:"blackout_schedule"`
}

// NetworkcontainerPortControlBlackoutSettingAttrTypes contains the attribute types for NetworkcontainerPortControlBlackoutSettingModel
var NetworkcontainerPortControlBlackoutSettingAttrTypes = map[string]attr.Type{
	"enable_blackout":   types.BoolType,
	"blackout_duration": types.Int64Type,
	"blackout_schedule": types.ObjectType{AttrTypes: NetworkcontainerportcontrolblackoutsettingBlackoutScheduleAttrTypes},
}

// NetworkcontainerPortControlBlackoutSettingResourceSchemaAttributes contains the schema attributes for NetworkcontainerPortControlBlackoutSettingModel
var NetworkcontainerPortControlBlackoutSettingResourceSchemaAttributes = map[string]schema.Attribute{
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
		Attributes:          NetworkcontainerportcontrolblackoutsettingBlackoutScheduleResourceSchemaAttributes,
		Optional:            true,
		MarkdownDescription: "A Schedule Setting struct that determines blackout schedule.",
	},
}

// ExpandNetworkcontainerPortControlBlackoutSetting converts a Terraform Object to SDK type
func ExpandNetworkcontainerPortControlBlackoutSetting(ctx context.Context, o types.Object, diags *diag.Diagnostics) *niosipam.NetworkcontainerPortControlBlackoutSetting {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m NetworkcontainerPortControlBlackoutSettingModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

// Expand converts the Terraform model to SDK type
func (m *NetworkcontainerPortControlBlackoutSettingModel) Expand(ctx context.Context, diags *diag.Diagnostics) *niosipam.NetworkcontainerPortControlBlackoutSetting {
	if m == nil {
		return nil
	}
	to := &niosipam.NetworkcontainerPortControlBlackoutSetting{
		EnableBlackout:   flex.ExpandBoolPointer(m.EnableBlackout),
		BlackoutDuration: flex.ExpandInt64Pointer(m.BlackoutDuration),
		BlackoutSchedule: ExpandNetworkcontainerportcontrolblackoutsettingBlackoutSchedule(ctx, m.BlackoutSchedule, diags),
	}
	return to
}

// FlattenNetworkcontainerPortControlBlackoutSetting converts an SDK type to Terraform Object
func FlattenNetworkcontainerPortControlBlackoutSetting(ctx context.Context, from *niosipam.NetworkcontainerPortControlBlackoutSetting, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(NetworkcontainerPortControlBlackoutSettingAttrTypes)
	}
	m := &NetworkcontainerPortControlBlackoutSettingModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, NetworkcontainerPortControlBlackoutSettingAttrTypes, m)
	diags.Append(d...)
	return t
}

// Flatten populates the Terraform model from SDK type
func (m *NetworkcontainerPortControlBlackoutSettingModel) Flatten(ctx context.Context, from *niosipam.NetworkcontainerPortControlBlackoutSetting, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.EnableBlackout = flex.FlattenBoolPointer(from.EnableBlackout)
	m.BlackoutDuration = flex.FlattenInt64Pointer(from.BlackoutDuration)
	m.BlackoutSchedule = FlattenNetworkcontainerportcontrolblackoutsettingBlackoutSchedule(ctx, from.BlackoutSchedule, diags)
}
