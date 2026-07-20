package dns

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
	uddidns "github.com/infobloxopen/universal-ddi-go-client/dnsconfig"
)

// ECSZoneModel is the Terraform model for ECSZone
type ECSZoneModel struct {
	Access       types.String `tfsdk:"access"`
	Fqdn         types.String `tfsdk:"fqdn"`
	ProtocolFqdn types.String `tfsdk:"protocol_fqdn"`
}

// ECSZoneAttrTypes contains the attribute types for ECSZoneModel
var ECSZoneAttrTypes = map[string]attr.Type{
	"access":        types.StringType,
	"fqdn":          types.StringType,
	"protocol_fqdn": types.StringType,
}

// ECSZoneResourceSchemaAttributes contains the schema attributes for ECSZoneModel
var ECSZoneResourceSchemaAttributes = map[string]schema.Attribute{
	"access": schema.StringAttribute{
		Required:            true,
		MarkdownDescription: "Access control for zone.  Allowed values: * _allow_, * _deny_.",
	},
	"fqdn": schema.StringAttribute{
		Required:            true,
		MarkdownDescription: "Zone FQDN.",
	},
	"protocol_fqdn": schema.StringAttribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "Zone FQDN in punycode.",
	},
}

// ExpandECSZone converts a Terraform Object to SDK type
func ExpandECSZone(ctx context.Context, o types.Object, diags *diag.Diagnostics) *uddidns.ECSZone {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m ECSZoneModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

// Expand converts the Terraform model to SDK type
func (m *ECSZoneModel) Expand(ctx context.Context, diags *diag.Diagnostics) *uddidns.ECSZone {
	if m == nil {
		return nil
	}
	to := &uddidns.ECSZone{
		Access:       flex.ExpandString(m.Access),
		Fqdn:         flex.ExpandString(m.Fqdn),
		ProtocolFqdn: flex.ExpandStringPointer(m.ProtocolFqdn),
	}
	return to
}

// FlattenECSZone converts an SDK type to Terraform Object
func FlattenECSZone(ctx context.Context, from *uddidns.ECSZone, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(ECSZoneAttrTypes)
	}
	m := &ECSZoneModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, ECSZoneAttrTypes, m)
	diags.Append(d...)
	return t
}

// Flatten populates the Terraform model from SDK type
func (m *ECSZoneModel) Flatten(ctx context.Context, from *uddidns.ECSZone, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.Access = flex.FlattenString(from.Access)
	m.Fqdn = flex.FlattenString(from.Fqdn)
	m.ProtocolFqdn = flex.FlattenStringPointer(from.ProtocolFqdn)
}
