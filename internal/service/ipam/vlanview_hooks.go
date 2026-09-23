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
	if !m.VlanNamePrefix.IsUnknown() && !m.VlanNamePrefix.IsNull() {
		if !m.PreCreateVlan.IsNull() && !m.PreCreateVlan.IsUnknown() && !m.PreCreateVlan.ValueBool() {
			resp.Diagnostics.AddError(
				"Configuration Error",
				"`vlan_name_prefix` can only be set when `pre_create_vlan` is set to true.",
			)
		}
	}
	if !m.StartVlanId.IsUnknown() && !m.StartVlanId.IsNull() &&
		!m.EndVlanId.IsUnknown() && !m.EndVlanId.IsNull() {
		if m.StartVlanId.ValueInt64() > m.EndVlanId.ValueInt64() {
			resp.Diagnostics.AddError(
				"Configuration Error",
				"`start_vlan_id` must be less than or equal to `end_vlan_id`.",
			)
		}
	}
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
