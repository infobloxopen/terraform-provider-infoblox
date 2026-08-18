package ipam

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
	uddiipam "github.com/infobloxopen/universal-ddi-go-client/ipam"
)

// UtilizationThresholdModel is the Terraform model for UtilizationThreshold
type UtilizationThresholdModel struct {
	Enabled types.Bool  `tfsdk:"enabled"`
	High    types.Int64 `tfsdk:"high"`
	Low     types.Int64 `tfsdk:"low"`
}

// UtilizationThresholdAttrTypes contains the attribute types for UtilizationThresholdModel
var UtilizationThresholdAttrTypes = map[string]attr.Type{
	"enabled": types.BoolType,
	"high":    types.Int64Type,
	"low":     types.Int64Type,
}

// UtilizationThresholdResourceSchemaAttributes contains the schema attributes for UtilizationThresholdModel
var UtilizationThresholdResourceSchemaAttributes = map[string]schema.Attribute{
	"enabled": schema.BoolAttribute{
		Required:            true,
		MarkdownDescription: "Indicates whether the utilization threshold for IP addresses is enabled or not.",
	},
	"high": schema.Int64Attribute{
		Required:            true,
		MarkdownDescription: "The high threshold value for the percentage of used IP addresses relative to the total IP addresses available in the scope of the object. Thresholds are inclusive in the comparison test.",
	},
	"low": schema.Int64Attribute{
		Required:            true,
		MarkdownDescription: "The low threshold value for the percentage of used IP addresses relative to the total IP addresses available in the scope of the object. Thresholds are inclusive in the comparison test.",
	},
}

// ExpandUtilizationThreshold converts a Terraform Object to SDK type
func ExpandUtilizationThreshold(ctx context.Context, o types.Object, diags *diag.Diagnostics) *uddiipam.UtilizationThreshold {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m UtilizationThresholdModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

// Expand converts the Terraform model to SDK type
func (m *UtilizationThresholdModel) Expand(ctx context.Context, diags *diag.Diagnostics) *uddiipam.UtilizationThreshold {
	if m == nil {
		return nil
	}
	to := &uddiipam.UtilizationThreshold{
		Enabled: flex.ExpandBool(m.Enabled),
		High:    flex.ExpandInt64(m.High),
		Low:     flex.ExpandInt64(m.Low),
	}
	return to
}

// FlattenUtilizationThreshold converts an SDK type to Terraform Object
func FlattenUtilizationThreshold(ctx context.Context, from *uddiipam.UtilizationThreshold, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(UtilizationThresholdAttrTypes)
	}
	m := &UtilizationThresholdModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, UtilizationThresholdAttrTypes, m)
	diags.Append(d...)
	return t
}

// Flatten populates the Terraform model from SDK type
func (m *UtilizationThresholdModel) Flatten(ctx context.Context, from *uddiipam.UtilizationThreshold, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.Enabled = flex.FlattenBool(from.Enabled)
	m.High = flex.FlattenInt64(from.High)
	m.Low = flex.FlattenInt64(from.Low)
}
