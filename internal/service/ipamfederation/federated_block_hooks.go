package ipamfederation

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/dynamicallocation"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/utils"
	uddiipamfederation "github.com/infobloxopen/universal-ddi-go-client/ipamfederation"
)

// ValidateFederatedBlock validates the FederatedBlock configuration.
func ValidateFederatedBlock(ctx context.Context, data FederatedBlockModel, resp *resource.ValidateConfigResponse) {
	if uddi := flex.ExpandNestedObject[UDDIFederatedBlockModel](ctx, data.UDDI, &resp.Diagnostics); uddi != nil {
		validateFederatedBlockUDDIConfig(ctx, uddi, resp)
	}
}

func validateFederatedBlockUDDIConfig(ctx context.Context, m *UDDIFederatedBlockModel, resp *resource.ValidateConfigResponse) {
	if m.NetworkCompliance.IsNull() || m.NetworkCompliance.IsUnknown() {
		return
	}

	var nc NetworkComplianceModel
	resp.Diagnostics.Append(m.NetworkCompliance.As(ctx, &nc, basetypes.ObjectAsOptions{})...)
	if resp.Diagnostics.HasError() {
		return
	}

	ncPath := path.Root("uddi").AtName("network_compliance")

	if !nc.MinimumNetmaskLength.IsNull() && !nc.MinimumNetmaskLength.IsUnknown() &&
		!m.Cidr.IsNull() && !m.Cidr.IsUnknown() {
		min := nc.MinimumNetmaskLength.ValueInt64()
		cidr := m.Cidr.ValueInt64()

		if min <= cidr {
			resp.Diagnostics.AddAttributeError(
				ncPath.AtName("minimum_netmask_length"),
				"Invalid Network Compliance",
				fmt.Sprintf("\"minimum_netmask_length\" must be greater than the block's own \"cidr\" (/%d).", cidr),
			)
		}
	}

	if nc.MinimumNetmaskLength.IsNull() || nc.MinimumNetmaskLength.IsUnknown() ||
		nc.MaximumNetmaskLength.IsNull() || nc.MaximumNetmaskLength.IsUnknown() {
		return
	}

	min := nc.MinimumNetmaskLength.ValueInt64()
	max := nc.MaximumNetmaskLength.ValueInt64()

	if min >= max {
		resp.Diagnostics.AddAttributeError(
			ncPath.AtName("minimum_netmask_length"),
			"Invalid Network Compliance",
			"\"minimum_netmask_length\" must be less than \"maximum_netmask_length\".",
		)
		resp.Diagnostics.AddAttributeError(
			ncPath.AtName("maximum_netmask_length"),
			"Invalid Network Compliance",
			"\"maximum_netmask_length\" must be greater than \"minimum_netmask_length\".",
		)
	}

	if nc.DefaultNetmaskLength.IsNull() || nc.DefaultNetmaskLength.IsUnknown() {
		return
	}
	def := nc.DefaultNetmaskLength.ValueInt64()

	if def > max {
		resp.Diagnostics.AddAttributeError(
			ncPath.AtName("default_netmask_length"),
			"Invalid Network Compliance",
			"\"default_netmask_length\" must be less than or equal to \"maximum_netmask_length\".",
		)
	}
	if def < min {
		resp.Diagnostics.AddAttributeError(
			ncPath.AtName("default_netmask_length"),
			"Invalid Network Compliance",
			"\"default_netmask_length\" must be greater than or equal to \"minimum_netmask_length\".",
		)
	}
}

// BuildFederatedBlockAllocation intentionally always returns nil. Unlike the usual
// suffix-on-Address allocation pattern, the Federation API requires a completely
// different SDK call for next-available allocation (NextAvailablePoolBlockAPI or
// NextAvailableFederatedBlockAPI, chosen by next_available_id's prefix), so the actual
// routing is hand-written in FederatedBlockResource.Create/createNextAvailable instead of
// here. This hook stays a no-op so the standard Expand() path is unaffected; it is not an
// unfinished stub.
func BuildFederatedBlockAllocation(ctx context.Context, allocObj types.Object, diags *diag.Diagnostics) *string {
	return nil
}

// LockFederatedBlockAllocation serializes concurrent next-available allocations
// that target the same parent scope by acquiring a per-scope mutex keyed on the next_available_id
func LockFederatedBlockAllocation(ctx context.Context, uddiBlock types.Object, diags *diag.Diagnostics) func() {
	noop := func() {}
	if uddiBlock.IsNull() || uddiBlock.IsUnknown() {
		return noop
	}

	allocVal, ok := uddiBlock.Attributes()["dynamic_allocation"]
	if !ok {
		return noop
	}
	allocObj, ok := allocVal.(types.Object)
	if !ok || allocObj.IsNull() || allocObj.IsUnknown() {
		return noop
	}

	var m dynamicallocation.NextAvailableFederatedBlockModel
	diags.Append(allocObj.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return noop
	}

	key := m.NextAvailableId.ValueString()
	if key == "" {
		return noop
	}

	utils.GlobalMutexStore.Lock(key)
	return func() { utils.GlobalMutexStore.Unlock(key) }
}

// ExpandInteger128Pointer converts a TF string to the SDK's Integer128 wrapper.
//
// Integer128 is a struct (not a plain string alias) because the swagger
// definition (federationInteger128: type=string, format=byte) still generates
// a wrapper type with a RawValue field, unlike other named-scalar string refs.
func ExpandInteger128Pointer(val types.String) *uddiipamfederation.Integer128 {
	if val.IsNull() || val.IsUnknown() {
		return nil
	}
	return &uddiipamfederation.Integer128{RawValue: val.ValueStringPointer()}
}

// FlattenInteger128Pointer converts the SDK's Integer128 wrapper to a TF string.
func FlattenInteger128Pointer(from *uddiipamfederation.Integer128) types.String {
	if from == nil || from.RawValue == nil {
		return types.StringNull()
	}
	return types.StringValue(*from.RawValue)
}
