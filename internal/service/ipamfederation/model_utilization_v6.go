package ipamfederation

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	uddiipamfederation "github.com/infobloxopen/universal-ddi-go-client/ipamfederation"
)

// UtilizationV6Model is the Terraform model for UtilizationV6
type UtilizationV6Model struct {
	Total types.String `tfsdk:"total"`
	Used  types.String `tfsdk:"used"`
}

// UtilizationV6AttrTypes contains the attribute types for UtilizationV6Model
var UtilizationV6AttrTypes = map[string]attr.Type{
	"total": types.StringType,
	"used":  types.StringType,
}

// UtilizationV6ResourceSchemaAttributes contains the schema attributes for UtilizationV6Model
var UtilizationV6ResourceSchemaAttributes = map[string]schema.Attribute{
	"total": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "Total IPv6 addresses.",
	},
	"used": schema.StringAttribute{
		Computed:            true,
		MarkdownDescription: "Used IPv6 addresses.",
	},
}

// ExpandUtilizationV6 converts a Terraform Object to SDK type
func ExpandUtilizationV6(ctx context.Context, o types.Object, diags *diag.Diagnostics) *uddiipamfederation.UtilizationV6 {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m UtilizationV6Model
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

// Expand converts the Terraform model to SDK type
func (m *UtilizationV6Model) Expand(ctx context.Context, diags *diag.Diagnostics) *uddiipamfederation.UtilizationV6 {
	if m == nil {
		return nil
	}
	to := &uddiipamfederation.UtilizationV6{
		Total: ExpandInteger128Pointer(m.Total),
		Used:  ExpandInteger128Pointer(m.Used),
	}
	return to
}

// FlattenUtilizationV6 converts an SDK type to Terraform Object
func FlattenUtilizationV6(ctx context.Context, from *uddiipamfederation.UtilizationV6, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(UtilizationV6AttrTypes)
	}
	m := &UtilizationV6Model{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, UtilizationV6AttrTypes, m)
	diags.Append(d...)
	return t
}

// Flatten populates the Terraform model from SDK type
func (m *UtilizationV6Model) Flatten(ctx context.Context, from *uddiipamfederation.UtilizationV6, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.Total = FlattenInteger128Pointer(from.Total)
	m.Used = FlattenInteger128Pointer(from.Used)
}
