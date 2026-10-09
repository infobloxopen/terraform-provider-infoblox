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

// NetworkIpamThresholdSettingsModel is the Terraform model for NetworkIpamThresholdSettings
type NetworkIpamThresholdSettingsModel struct {
	TriggerValue types.Int64 `tfsdk:"trigger_value"`
	ResetValue   types.Int64 `tfsdk:"reset_value"`
}

// NetworkIpamThresholdSettingsAttrTypes contains the attribute types for NetworkIpamThresholdSettingsModel
var NetworkIpamThresholdSettingsAttrTypes = map[string]attr.Type{
	"trigger_value": types.Int64Type,
	"reset_value":   types.Int64Type,
}

// NetworkIpamThresholdSettingsResourceSchemaAttributes contains the schema attributes for NetworkIpamThresholdSettingsModel
var NetworkIpamThresholdSettingsResourceSchemaAttributes = map[string]schema.Attribute{
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

// ExpandNetworkIpamThresholdSettings converts a Terraform Object to SDK type
func ExpandNetworkIpamThresholdSettings(ctx context.Context, o types.Object, diags *diag.Diagnostics) *niosipam.NetworkIpamThresholdSettings {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m NetworkIpamThresholdSettingsModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

// Expand converts the Terraform model to SDK type
func (m *NetworkIpamThresholdSettingsModel) Expand(ctx context.Context, diags *diag.Diagnostics) *niosipam.NetworkIpamThresholdSettings {
	if m == nil {
		return nil
	}
	to := &niosipam.NetworkIpamThresholdSettings{
		TriggerValue: flex.ExpandInt64Pointer(m.TriggerValue),
		ResetValue:   flex.ExpandInt64Pointer(m.ResetValue),
	}
	return to
}

// FlattenNetworkIpamThresholdSettings converts an SDK type to Terraform Object
func FlattenNetworkIpamThresholdSettings(ctx context.Context, from *niosipam.NetworkIpamThresholdSettings, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(NetworkIpamThresholdSettingsAttrTypes)
	}
	m := &NetworkIpamThresholdSettingsModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, NetworkIpamThresholdSettingsAttrTypes, m)
	diags.Append(d...)
	return t
}

// Flatten populates the Terraform model from SDK type
func (m *NetworkIpamThresholdSettingsModel) Flatten(ctx context.Context, from *niosipam.NetworkIpamThresholdSettings, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.TriggerValue = flex.FlattenInt64Pointer(from.TriggerValue)
	m.ResetValue = flex.FlattenInt64Pointer(from.ResetValue)
}
