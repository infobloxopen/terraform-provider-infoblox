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

// Ipv6networkPortControlBlackoutSettingModel is the Terraform model for Ipv6networkPortControlBlackoutSetting
type Ipv6networkPortControlBlackoutSettingModel struct {
	EnableBlackout   types.Bool   `tfsdk:"enable_blackout"`
	BlackoutDuration types.Int64  `tfsdk:"blackout_duration"`
	BlackoutSchedule types.Object `tfsdk:"blackout_schedule"`
}

// Ipv6networkPortControlBlackoutSettingAttrTypes contains the attribute types for Ipv6networkPortControlBlackoutSettingModel
var Ipv6networkPortControlBlackoutSettingAttrTypes = map[string]attr.Type{
	"enable_blackout":   types.BoolType,
	"blackout_duration": types.Int64Type,
	"blackout_schedule": types.ObjectType{AttrTypes: Ipv6networkportcontrolblackoutsettingBlackoutScheduleAttrTypes},
}

// Ipv6networkPortControlBlackoutSettingResourceSchemaAttributes contains the schema attributes for Ipv6networkPortControlBlackoutSettingModel
var Ipv6networkPortControlBlackoutSettingResourceSchemaAttributes = map[string]schema.Attribute{
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
		Attributes:          Ipv6networkportcontrolblackoutsettingBlackoutScheduleResourceSchemaAttributes,
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "A Schedule Setting struct that determines blackout schedule.",
	},
}

// ExpandIpv6networkPortControlBlackoutSetting converts a Terraform Object to SDK type
func ExpandIpv6networkPortControlBlackoutSetting(ctx context.Context, o types.Object, diags *diag.Diagnostics) *niosipam.Ipv6networkPortControlBlackoutSetting {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m Ipv6networkPortControlBlackoutSettingModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

// Expand converts the Terraform model to SDK type
func (m *Ipv6networkPortControlBlackoutSettingModel) Expand(ctx context.Context, diags *diag.Diagnostics) *niosipam.Ipv6networkPortControlBlackoutSetting {
	if m == nil {
		return nil
	}
	to := &niosipam.Ipv6networkPortControlBlackoutSetting{
		EnableBlackout:   flex.ExpandBoolPointer(m.EnableBlackout),
		BlackoutDuration: flex.ExpandInt64Pointer(m.BlackoutDuration),
		BlackoutSchedule: ExpandIpv6networkportcontrolblackoutsettingBlackoutSchedule(ctx, m.BlackoutSchedule, diags),
	}
	return to
}

// FlattenIpv6networkPortControlBlackoutSetting converts an SDK type to Terraform Object
func FlattenIpv6networkPortControlBlackoutSetting(ctx context.Context, from *niosipam.Ipv6networkPortControlBlackoutSetting, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(Ipv6networkPortControlBlackoutSettingAttrTypes)
	}
	m := &Ipv6networkPortControlBlackoutSettingModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, Ipv6networkPortControlBlackoutSettingAttrTypes, m)
	diags.Append(d...)
	return t
}

// Flatten populates the Terraform model from SDK type
func (m *Ipv6networkPortControlBlackoutSettingModel) Flatten(ctx context.Context, from *niosipam.Ipv6networkPortControlBlackoutSetting, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.EnableBlackout = flex.FlattenBoolPointer(from.EnableBlackout)
	m.BlackoutDuration = flex.FlattenInt64Pointer(from.BlackoutDuration)
	m.BlackoutSchedule = FlattenIpv6networkportcontrolblackoutsettingBlackoutSchedule(ctx, from.BlackoutSchedule, diags)
}
