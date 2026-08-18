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

// NetworkcontainerDiscoveryBlackoutSettingModel is the Terraform model for NetworkcontainerDiscoveryBlackoutSetting
type NetworkcontainerDiscoveryBlackoutSettingModel struct {
	EnableBlackout   types.Bool   `tfsdk:"enable_blackout"`
	BlackoutDuration types.Int64  `tfsdk:"blackout_duration"`
	BlackoutSchedule types.Object `tfsdk:"blackout_schedule"`
}

// NetworkcontainerDiscoveryBlackoutSettingAttrTypes contains the attribute types for NetworkcontainerDiscoveryBlackoutSettingModel
var NetworkcontainerDiscoveryBlackoutSettingAttrTypes = map[string]attr.Type{
	"enable_blackout":   types.BoolType,
	"blackout_duration": types.Int64Type,
	"blackout_schedule": types.ObjectType{AttrTypes: NetworkcontainerdiscoveryblackoutsettingBlackoutScheduleAttrTypes},
}

// NetworkcontainerDiscoveryBlackoutSettingResourceSchemaAttributes contains the schema attributes for NetworkcontainerDiscoveryBlackoutSettingModel
var NetworkcontainerDiscoveryBlackoutSettingResourceSchemaAttributes = map[string]schema.Attribute{
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
		Attributes:          NetworkcontainerdiscoveryblackoutsettingBlackoutScheduleResourceSchemaAttributes,
		Optional:            true,
		MarkdownDescription: "A Schedule Setting struct that determines blackout schedule.",
	},
}

// ExpandNetworkcontainerDiscoveryBlackoutSetting converts a Terraform Object to SDK type
func ExpandNetworkcontainerDiscoveryBlackoutSetting(ctx context.Context, o types.Object, diags *diag.Diagnostics) *niosipam.NetworkcontainerDiscoveryBlackoutSetting {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m NetworkcontainerDiscoveryBlackoutSettingModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

// Expand converts the Terraform model to SDK type
func (m *NetworkcontainerDiscoveryBlackoutSettingModel) Expand(ctx context.Context, diags *diag.Diagnostics) *niosipam.NetworkcontainerDiscoveryBlackoutSetting {
	if m == nil {
		return nil
	}
	to := &niosipam.NetworkcontainerDiscoveryBlackoutSetting{
		EnableBlackout:   flex.ExpandBoolPointer(m.EnableBlackout),
		BlackoutDuration: flex.ExpandInt64Pointer(m.BlackoutDuration),
		BlackoutSchedule: ExpandNetworkcontainerdiscoveryblackoutsettingBlackoutSchedule(ctx, m.BlackoutSchedule, diags),
	}
	return to
}

// FlattenNetworkcontainerDiscoveryBlackoutSetting converts an SDK type to Terraform Object
func FlattenNetworkcontainerDiscoveryBlackoutSetting(ctx context.Context, from *niosipam.NetworkcontainerDiscoveryBlackoutSetting, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(NetworkcontainerDiscoveryBlackoutSettingAttrTypes)
	}
	m := &NetworkcontainerDiscoveryBlackoutSettingModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, NetworkcontainerDiscoveryBlackoutSettingAttrTypes, m)
	diags.Append(d...)
	return t
}

// Flatten populates the Terraform model from SDK type
func (m *NetworkcontainerDiscoveryBlackoutSettingModel) Flatten(ctx context.Context, from *niosipam.NetworkcontainerDiscoveryBlackoutSetting, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.EnableBlackout = flex.FlattenBoolPointer(from.EnableBlackout)
	m.BlackoutDuration = flex.FlattenInt64Pointer(from.BlackoutDuration)
	m.BlackoutSchedule = FlattenNetworkcontainerdiscoveryblackoutsettingBlackoutSchedule(ctx, from.BlackoutSchedule, diags)
}
