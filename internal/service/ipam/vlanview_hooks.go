package ipam

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
)

// ValidateVlanview validates the Vlanview configuration.
func ValidateVlanview(ctx context.Context, data VlanviewModel, resp *resource.ValidateConfigResponse) {
	if nios := flex.ExpandNestedObject[NIOSVlanviewModel](ctx, data.NIOS, &resp.Diagnostics); nios != nil {
		validateVlanviewNIOSConfig(ctx, nios, resp)
	}
}

func validateVlanviewNIOSConfig(ctx context.Context, m *NIOSVlanviewModel, resp *resource.ValidateConfigResponse) {
}

func PostFlattenVlanviewNIOS(ctx context.Context, planned, flattened *NIOSVlanviewModel, diags *diag.Diagnostics) {
	if planned == nil || flattened == nil {
		return
	}

	if !planned.PreCreateVlan.IsUnknown() && !planned.PreCreateVlan.IsNull() {
		flattened.PreCreateVlan = planned.PreCreateVlan
	}

	if !planned.VlanNamePrefix.IsUnknown() && !planned.VlanNamePrefix.IsNull() {
		flattened.VlanNamePrefix = planned.VlanNamePrefix
	}
}
