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

// NetworkIpamTrapSettingsModel is the Terraform model for NetworkIpamTrapSettings
type NetworkIpamTrapSettingsModel struct {
	EnableEmailWarnings types.Bool `tfsdk:"enable_email_warnings"`
	EnableSnmpWarnings  types.Bool `tfsdk:"enable_snmp_warnings"`
}

// NetworkIpamTrapSettingsAttrTypes contains the attribute types for NetworkIpamTrapSettingsModel
var NetworkIpamTrapSettingsAttrTypes = map[string]attr.Type{
	"enable_email_warnings": types.BoolType,
	"enable_snmp_warnings":  types.BoolType,
}

// NetworkIpamTrapSettingsResourceSchemaAttributes contains the schema attributes for NetworkIpamTrapSettingsModel
var NetworkIpamTrapSettingsResourceSchemaAttributes = map[string]schema.Attribute{
	"enable_email_warnings": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(false),
		MarkdownDescription: "Determines whether sending warnings by email is enabled or not.",
	},
	"enable_snmp_warnings": schema.BoolAttribute{
		Optional:            true,
		Computed:            true,
		Default:             booldefault.StaticBool(true),
		MarkdownDescription: "Determines whether sending warnings by SNMP is enabled or not.",
	},
}

// ExpandNetworkIpamTrapSettings converts a Terraform Object to SDK type
func ExpandNetworkIpamTrapSettings(ctx context.Context, o types.Object, diags *diag.Diagnostics) *niosipam.NetworkIpamTrapSettings {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m NetworkIpamTrapSettingsModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

// Expand converts the Terraform model to SDK type
func (m *NetworkIpamTrapSettingsModel) Expand(ctx context.Context, diags *diag.Diagnostics) *niosipam.NetworkIpamTrapSettings {
	if m == nil {
		return nil
	}
	to := &niosipam.NetworkIpamTrapSettings{
		EnableEmailWarnings: flex.ExpandBoolPointer(m.EnableEmailWarnings),
		EnableSnmpWarnings:  flex.ExpandBoolPointer(m.EnableSnmpWarnings),
	}
	return to
}

// FlattenNetworkIpamTrapSettings converts an SDK type to Terraform Object
func FlattenNetworkIpamTrapSettings(ctx context.Context, from *niosipam.NetworkIpamTrapSettings, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(NetworkIpamTrapSettingsAttrTypes)
	}
	m := &NetworkIpamTrapSettingsModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, NetworkIpamTrapSettingsAttrTypes, m)
	diags.Append(d...)
	return t
}

// Flatten populates the Terraform model from SDK type
func (m *NetworkIpamTrapSettingsModel) Flatten(ctx context.Context, from *niosipam.NetworkIpamTrapSettings, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.EnableEmailWarnings = flex.FlattenBoolPointer(from.EnableEmailWarnings)
	m.EnableSnmpWarnings = flex.FlattenBoolPointer(from.EnableSnmpWarnings)
}
