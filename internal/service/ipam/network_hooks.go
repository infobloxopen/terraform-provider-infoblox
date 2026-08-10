package ipam

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	niosipam "github.com/infobloxopen/infoblox-nios-go-client/ipam"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/core"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/dynamicallocation"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/utils"
)

// ValidateNetwork validates the Network configuration.
func ValidateNetwork(ctx context.Context, data NetworkModel, resp *resource.ValidateConfigResponse) {
	if nios := flex.ExpandNestedObject[NIOSNetworkModel](ctx, data.NIOS, &resp.Diagnostics); nios != nil {
		validateNetworkNIOSConfig(ctx, nios, resp)
	}
	if uddi := flex.ExpandNestedObject[UDDINetworkModel](ctx, data.UDDI, &resp.Diagnostics); uddi != nil {
		validateNetworkUDDIConfig(ctx, uddi, resp)
	}
}

func validateNetworkNIOSConfig(ctx context.Context, m *NIOSNetworkModel, resp *resource.ValidateConfigResponse) {
}

func validateNetworkUDDIConfig(ctx context.Context, m *UDDINetworkModel, resp *resource.ValidateConfigResponse) {
}

func BuildNetworkFuncCall(ctx context.Context, data types.Object, diags *diag.Diagnostics) *niosipam.FuncCall {
	if data.IsNull() || data.IsUnknown() {
		return nil
	}

	var m dynamicallocation.NextAvailableNetworkModel
	diags.Append(data.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}

	return m.FuncCall(ctx, "Network", "network", diags)
}

func PostFlattenNetworkNIOS(ctx context.Context, planned, flattened *NIOSNetworkModel, diags *diag.Diagnostics) {
	if planned != nil && !planned.Options.IsUnknown() {
		if reordered, d := utils.ReorderAndFilterDHCPOptions(ctx, planned.Options, flattened.Options); !d.HasError() {
			flattened.Options = reordered.(basetypes.ListValue)
		}
	}
}

func BuildNetworkAllocation(ctx context.Context, allocObj types.Object, diags *diag.Diagnostics) *string {
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

// LockNetworkAllocation serializes concurrent next-available allocations
// that target the same parent scope by acquiring a per-scope mutex keyed on the next_available_id
func LockNetworkAllocation(ctx context.Context, uddiBlock types.Object, diags *diag.Diagnostics) func() {
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

func (r *NetworkResource) isNetworkConvertedToContainer(ctx context.Context, data *NetworkModel) bool {
	if r.backend != core.BackendNIOS || r.containerService == nil {
		return false
	}

	var diags diag.Diagnostics
	nios := flex.ExpandNestedObject[NIOSNetworkModel](ctx, data.NIOS, &diags)
	if nios == nil || diags.HasError() {
		return false
	}

	// Try to fetch as Network container
	_, _, _, err := r.containerService.List(ctx, &core.ListOptions{
		Filters: map[string]string{
			"nios.network": nios.Network.ValueString(),
		},
	})
	return err == nil
}
