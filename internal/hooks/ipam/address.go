package ipam

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/infobloxopen/terraform-provider-unified/internal/dynamicallocation"
	"github.com/infobloxopen/terraform-provider-unified/internal/utils"
)

// ValidateAddress validates the Address configuration.
func ValidateAddress(ctx context.Context, niosBlock, uddiBlock types.Object, resp *resource.ValidateConfigResponse) {
	if !uddiBlock.IsNull() && !uddiBlock.IsUnknown() {
		validateAddressUDDIConfig(ctx, uddiBlock, resp)
	}
}

func validateAddressUDDIConfig(ctx context.Context, data types.Object, resp *resource.ValidateConfigResponse) {
}

func BuildAddressAllocation(ctx context.Context, allocObj types.Object, diags *diag.Diagnostics) *string {
	if allocObj.IsNull() || allocObj.IsUnknown() {
		return nil
	}

	var m dynamicallocation.NextAvailableAddressModel
	diags.Append(allocObj.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}

	v := m.Suffixed("/nextavailableip")
	return &v
}

// LockAddressAllocation serializes concurrent next-available allocations that
// target the same parent scope by acquiring a per-scope mutex keyed on the next_available_id
func LockAddressAllocation(ctx context.Context, uddiBlock types.Object, diags *diag.Diagnostics) func() {
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

	var m dynamicallocation.NextAvailableAddressModel
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
