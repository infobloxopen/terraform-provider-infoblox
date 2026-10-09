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

// Ipv6networkcontainerPortControlBlackoutSettingModel is the Terraform model for Ipv6networkcontainerPortControlBlackoutSetting
type Ipv6networkcontainerPortControlBlackoutSettingModel struct {
	EnableBlackout   types.Bool   `tfsdk:"enable_blackout"`
	BlackoutDuration types.Int64  `tfsdk:"blackout_duration"`
	BlackoutSchedule types.Object `tfsdk:"blackout_schedule"`
}

// Ipv6networkcontainerPortControlBlackoutSettingAttrTypes contains the attribute types for Ipv6networkcontainerPortControlBlackoutSettingModel
var Ipv6networkcontainerPortControlBlackoutSettingAttrTypes = map[string]attr.Type{
	"enable_blackout":   types.BoolType,
	"blackout_duration": types.Int64Type,
	"blackout_schedule": types.ObjectType{AttrTypes: Ipv6networkcontainerportcontrolblackoutsettingBlackoutScheduleAttrTypes},
}

// Ipv6networkcontainerPortControlBlackoutSettingResourceSchemaAttributes contains the schema attributes for Ipv6networkcontainerPortControlBlackoutSettingModel
var Ipv6networkcontainerPortControlBlackoutSettingResourceSchemaAttributes = map[string]schema.Attribute{
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
		Attributes:          Ipv6networkcontainerportcontrolblackoutsettingBlackoutScheduleResourceSchemaAttributes,
		Optional:            true,
		MarkdownDescription: "A Schedule Setting struct that determines blackout schedule.",
	},
}

// ExpandIpv6networkcontainerPortControlBlackoutSetting converts a Terraform Object to SDK type
func ExpandIpv6networkcontainerPortControlBlackoutSetting(ctx context.Context, o types.Object, diags *diag.Diagnostics) *niosipam.Ipv6networkcontainerPortControlBlackoutSetting {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m Ipv6networkcontainerPortControlBlackoutSettingModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

// Expand converts the Terraform model to SDK type
func (m *Ipv6networkcontainerPortControlBlackoutSettingModel) Expand(ctx context.Context, diags *diag.Diagnostics) *niosipam.Ipv6networkcontainerPortControlBlackoutSetting {
	if m == nil {
		return nil
	}
	to := &niosipam.Ipv6networkcontainerPortControlBlackoutSetting{
		EnableBlackout:   flex.ExpandBoolPointer(m.EnableBlackout),
		BlackoutDuration: flex.ExpandInt64Pointer(m.BlackoutDuration),
		BlackoutSchedule: ExpandIpv6networkcontainerportcontrolblackoutsettingBlackoutSchedule(ctx, m.BlackoutSchedule, diags),
	}
	return to
}

// FlattenIpv6networkcontainerPortControlBlackoutSetting converts an SDK type to Terraform Object
func FlattenIpv6networkcontainerPortControlBlackoutSetting(ctx context.Context, from *niosipam.Ipv6networkcontainerPortControlBlackoutSetting, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(Ipv6networkcontainerPortControlBlackoutSettingAttrTypes)
	}
	m := &Ipv6networkcontainerPortControlBlackoutSettingModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, Ipv6networkcontainerPortControlBlackoutSettingAttrTypes, m)
	diags.Append(d...)
	return t
}

// Flatten populates the Terraform model from SDK type
func (m *Ipv6networkcontainerPortControlBlackoutSettingModel) Flatten(ctx context.Context, from *niosipam.Ipv6networkcontainerPortControlBlackoutSetting, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.EnableBlackout = flex.FlattenBoolPointer(from.EnableBlackout)
	m.BlackoutDuration = flex.FlattenInt64Pointer(from.BlackoutDuration)
	m.BlackoutSchedule = FlattenIpv6networkcontainerportcontrolblackoutsettingBlackoutSchedule(ctx, from.BlackoutSchedule, diags)
}
