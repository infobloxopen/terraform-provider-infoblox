package ipamfederation

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	schema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
	uddiipamfederation "github.com/infobloxopen/universal-ddi-go-client/ipamfederation"
)

// NetworkComplianceModel is the Terraform model for NetworkCompliance
type NetworkComplianceModel struct {
	DefaultNetmaskLength types.Int64 `tfsdk:"default_netmask_length"`
	MaximumNetmaskLength types.Int64 `tfsdk:"maximum_netmask_length"`
	MinimumNetmaskLength types.Int64 `tfsdk:"minimum_netmask_length"`
}

// NetworkComplianceAttrTypes contains the attribute types for NetworkComplianceModel
var NetworkComplianceAttrTypes = map[string]attr.Type{
	"default_netmask_length": types.Int64Type,
	"maximum_netmask_length": types.Int64Type,
	"minimum_netmask_length": types.Int64Type,
}

// NetworkComplianceResourceSchemaAttributes contains the schema attributes for NetworkComplianceModel
var NetworkComplianceResourceSchemaAttributes = map[string]schema.Attribute{
	"default_netmask_length": schema.Int64Attribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The default netmask length used when allocating child blocks or pools.",
	},
	"maximum_netmask_length": schema.Int64Attribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The maximum netmask length for allocating child blocks or pools.",
	},
	"minimum_netmask_length": schema.Int64Attribute{
		Optional:            true,
		Computed:            true,
		MarkdownDescription: "The minimum netmask length for allocating child blocks or pools.",
	},
}

// ExpandNetworkCompliance converts a Terraform Object to SDK type
func ExpandNetworkCompliance(ctx context.Context, o types.Object, diags *diag.Diagnostics) *uddiipamfederation.NetworkCompliance {
	if o.IsNull() || o.IsUnknown() {
		return nil
	}
	var m NetworkComplianceModel
	diags.Append(o.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}
	return m.Expand(ctx, diags)
}

// Expand converts the Terraform model to SDK type
func (m *NetworkComplianceModel) Expand(ctx context.Context, diags *diag.Diagnostics) *uddiipamfederation.NetworkCompliance {
	if m == nil {
		return nil
	}
	to := &uddiipamfederation.NetworkCompliance{
		DefaultNetmaskLength: flex.ExpandInt64Pointer(m.DefaultNetmaskLength),
		MaximumNetmaskLength: flex.ExpandInt64Pointer(m.MaximumNetmaskLength),
		MinimumNetmaskLength: flex.ExpandInt64Pointer(m.MinimumNetmaskLength),
	}
	return to
}

// FlattenNetworkCompliance converts an SDK type to Terraform Object
func FlattenNetworkCompliance(ctx context.Context, from *uddiipamfederation.NetworkCompliance, diags *diag.Diagnostics) types.Object {
	if from == nil {
		return types.ObjectNull(NetworkComplianceAttrTypes)
	}
	m := &NetworkComplianceModel{}
	m.Flatten(ctx, from, diags)
	t, d := types.ObjectValueFrom(ctx, NetworkComplianceAttrTypes, m)
	diags.Append(d...)
	return t
}

// Flatten populates the Terraform model from SDK type
func (m *NetworkComplianceModel) Flatten(ctx context.Context, from *uddiipamfederation.NetworkCompliance, diags *diag.Diagnostics) {
	if from == nil || m == nil {
		return
	}
	m.DefaultNetmaskLength = flex.FlattenInt64Pointer(from.DefaultNetmaskLength)
	m.MaximumNetmaskLength = flex.FlattenInt64Pointer(from.MaximumNetmaskLength)
	m.MinimumNetmaskLength = flex.FlattenInt64Pointer(from.MinimumNetmaskLength)
}
