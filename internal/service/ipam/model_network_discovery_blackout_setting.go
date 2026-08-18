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

// NetworkDiscoveryBlackoutSettingModel is the Terraform model for NetworkDiscoveryBlackoutSetting
type NetworkDiscoveryBlackoutSettingModel struct {
	EnableBlackout   types.Bool   `tfsdk:"enable_blackout"`
	BlackoutDuration types.Int64  `tfsdk:"blackout_duration"`
	BlackoutSchedule types.Object `tfsdk:"blackout_schedule"`
}

// NetworkDiscoveryBlackoutSettingAttrTypes contains the attribute types for NetworkDiscoveryBlackoutSettingModel
var NetworkDiscoveryBlackoutSettingAttrTypes = map[string]attr.Type{
	"enable_blackout":   types.BoolType,
	"blackout_duration": types.Int64Type,
	"blackout_schedule": types.ObjectType{AttrTypes: NetworkdiscoveryblackoutsettingBlackoutScheduleAttrTypes},
}

// NetworkDiscoveryBlackoutSettingResourceSchemaAttributes contains the schema attributes for NetworkDiscoveryBlackoutSettingModel
var NetworkDiscoveryBlackoutSettingResourceSchemaAttributes = map[string]schema.Attribute{
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
		Attributes:          NetworkdiscoveryblackoutsettingBlackoutScheduleResourceSchemaAttributes,
		Optional:            true,
		MarkdownDescription: "A Schedule Setting struct that determines blackout schedule.",
	},
}

// ExpandNetworkDiscoveryBlackoutSetting converts a Terraform Object to SDK type
func ExpandNetworkDiscoveryBlackoutSetting(ctx context.Context, o types.Object, diags *diag.Diagnostics) *niosipam.NetworkDiscoveryBlackoutSetting {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m NetworkDiscoveryBlackoutSettingModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

// Expand converts the Terraform model to SDK type
func (m *NetworkDiscoveryBlackoutSettingModel) Expand(ctx context.Context, diags *diag.Diagnostics) *niosipam.NetworkDiscoveryBlackoutSetting {
	if m == nil {
		return nil
	}
	to := &niosipam.NetworkDiscoveryBlackoutSetting{
		EnableBlackout:   flex.ExpandBoolPointer(m.EnableBlackout),
		BlackoutDuration: flex.ExpandInt64Pointer(m.BlackoutDuration),
		BlackoutSchedule: ExpandNetworkdiscoveryblackoutsettingBlackoutSchedule(ctx, m.BlackoutSchedule, diags),
	}
	return to
}

// FlattenNetworkDiscoveryBlackoutSetting converts an SDK type to Terraform Object
func FlattenNetworkDiscoveryBlackoutSetting(ctx context.Context, from *niosipam.NetworkDiscoveryBlackoutSetting, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(NetworkDiscoveryBlackoutSettingAttrTypes)
	}
	m := &NetworkDiscoveryBlackoutSettingModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, NetworkDiscoveryBlackoutSettingAttrTypes, m)
	diags.Append(d...)
	return t
}

// Flatten populates the Terraform model from SDK type
func (m *NetworkDiscoveryBlackoutSettingModel) Flatten(ctx context.Context, from *niosipam.NetworkDiscoveryBlackoutSetting, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.EnableBlackout = flex.FlattenBoolPointer(from.EnableBlackout)
	m.BlackoutDuration = flex.FlattenInt64Pointer(from.BlackoutDuration)
	m.BlackoutSchedule = FlattenNetworkdiscoveryblackoutsettingBlackoutSchedule(ctx, from.BlackoutSchedule, diags)
}
