package ipam

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	niosipam "github.com/infobloxopen/infoblox-nios-go-client/ipam"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
)

// NetworkcontainerIpamThresholdSettingsModel is the Terraform model for NetworkcontainerIpamThresholdSettings
type NetworkcontainerIpamThresholdSettingsModel struct {
	TriggerValue types.Int64 `tfsdk:"trigger_value"`
	ResetValue   types.Int64 `tfsdk:"reset_value"`
}

// NetworkcontainerIpamThresholdSettingsAttrTypes contains the attribute types for NetworkcontainerIpamThresholdSettingsModel
var NetworkcontainerIpamThresholdSettingsAttrTypes = map[string]attr.Type{
	"trigger_value": types.Int64Type,
	"reset_value":   types.Int64Type,
}

// NetworkcontainerIpamThresholdSettingsResourceSchemaAttributes contains the schema attributes for NetworkcontainerIpamThresholdSettingsModel
var NetworkcontainerIpamThresholdSettingsResourceSchemaAttributes = map[string]schema.Attribute{
	"trigger_value": schema.Int64Attribute{
		Optional:            true,
		Computed:            true,
		Default:             int64default.StaticInt64(95),
		MarkdownDescription: "Indicates the percentage point which triggers the email/SNMP trap sending.",
	},
	"reset_value": schema.Int64Attribute{
		Optional:            true,
		Computed:            true,
		Default:             int64default.StaticInt64(85),
		MarkdownDescription: "Indicates the percentage point which resets the email/SNMP trap sending.",
	},
}

// ExpandNetworkcontainerIpamThresholdSettings converts a Terraform Object to SDK type
func ExpandNetworkcontainerIpamThresholdSettings(ctx context.Context, o types.Object, diags *diag.Diagnostics) *niosipam.NetworkcontainerIpamThresholdSettings {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m NetworkcontainerIpamThresholdSettingsModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

// Expand converts the Terraform model to SDK type
func (m *NetworkcontainerIpamThresholdSettingsModel) Expand(ctx context.Context, diags *diag.Diagnostics) *niosipam.NetworkcontainerIpamThresholdSettings {
	if m == nil {
		return nil
	}
	to := &niosipam.NetworkcontainerIpamThresholdSettings{
		TriggerValue: flex.ExpandInt64Pointer(m.TriggerValue),
		ResetValue:   flex.ExpandInt64Pointer(m.ResetValue),
	}
	return to
}

// FlattenNetworkcontainerIpamThresholdSettings converts an SDK type to Terraform Object
func FlattenNetworkcontainerIpamThresholdSettings(ctx context.Context, from *niosipam.NetworkcontainerIpamThresholdSettings, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(NetworkcontainerIpamThresholdSettingsAttrTypes)
	}
	m := &NetworkcontainerIpamThresholdSettingsModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, NetworkcontainerIpamThresholdSettingsAttrTypes, m)
	diags.Append(d...)
	return t
}

// Flatten populates the Terraform model from SDK type
func (m *NetworkcontainerIpamThresholdSettingsModel) Flatten(ctx context.Context, from *niosipam.NetworkcontainerIpamThresholdSettings, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.TriggerValue = flex.FlattenInt64Pointer(from.TriggerValue)
	m.ResetValue = flex.FlattenInt64Pointer(from.ResetValue)
}
