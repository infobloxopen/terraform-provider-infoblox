package ipam

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	niosipam "github.com/infobloxopen/infoblox-nios-go-client/ipam"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/dynamicallocation"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/utils"
)

// ValidateIpv6network validates the Ipv6network configuration.
func ValidateIpv6network(ctx context.Context, data Ipv6networkModel, resp *resource.ValidateConfigResponse) {
	if nios := flex.ExpandNestedObject[NIOSIpv6networkModel](ctx, data.NIOS, &resp.Diagnostics); nios != nil {
		validateIpv6networkNIOSConfig(ctx, nios, resp)
	}
	if uddi := flex.ExpandNestedObject[UDDIIpv6networkModel](ctx, data.UDDI, &resp.Diagnostics); uddi != nil {
		validateIpv6networkUDDIConfig(ctx, uddi, resp)
	}
}

func validateIpv6networkNIOSConfig(ctx context.Context, m *NIOSIpv6networkModel, resp *resource.ValidateConfigResponse) {
}

func validateIpv6networkUDDIConfig(ctx context.Context, m *UDDIIpv6networkModel, resp *resource.ValidateConfigResponse) {
}

func BuildIpv6networkFuncCall(ctx context.Context, data types.Object, diags *diag.Diagnostics) *niosipam.FuncCall {
	if data.IsNull() || data.IsUnknown() {
		return nil
	}

	var m dynamicallocation.NextAvailableNetworkModel
	diags.Append(data.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}

	return m.FuncCall(ctx, "Network", "ipv6network", diags)
}

func PostFlattenIpv6networkNIOS(ctx context.Context, planned, flattened *NIOSIpv6networkModel, diags *diag.Diagnostics) {
	if planned != nil && !planned.Options.IsUnknown() {
		reordered, d := utils.ReorderAndFilterDHCPOptions(ctx, planned.Options, flattened.Options)
		diags.Append(*d...)
		if d.HasError() {
			return
		}
		if reorderedList, ok := reordered.(basetypes.ListValue); ok {
			flattened.Options = reorderedList
		}
	}
}

func BuildIpv6networkAllocation(ctx context.Context, allocObj types.Object, diags *diag.Diagnostics) *string {
	if allocObj.IsNull() || allocObj.IsUnknown() {
		return nil
	}

	var m dynamicallocation.NextAvailableSubnetModel
	diags.Append(allocObj.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}

	if m.NextAvailableId.IsNull() || m.NextAvailableId.IsUnknown() {
		return nil
	}

	allocated := m.Suffixed("/nextavailablesubnet")
	return &allocated
}

// LockIpv6networkAllocation serializes concurrent next-available allocations
// that target the same parent scope by acquiring a per-scope mutex keyed on the next_available_id
func LockIpv6networkAllocation(ctx context.Context, uddiBlock types.Object, diags *diag.Diagnostics) func() {
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

	var m dynamicallocation.NextAvailableSubnetModel
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
