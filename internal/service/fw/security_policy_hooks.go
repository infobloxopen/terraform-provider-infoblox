package fw

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
)

// ValidateSecurityPolicy validates the SecurityPolicy configuration.
func ValidateSecurityPolicy(ctx context.Context, data SecurityPolicyModel, resp *resource.ValidateConfigResponse) {
	if uddi := flex.ExpandNestedObject[UDDISecurityPolicyModel](ctx, data.UDDI, &resp.Diagnostics); uddi != nil {
		validateSecurityPolicyUDDIConfig(ctx, uddi, resp)
	}
}

func validateSecurityPolicyUDDIConfig(ctx context.Context, m *UDDISecurityPolicyModel, resp *resource.ValidateConfigResponse) {
}

// PostFlattenSecurityPolicyUDDI converts null int list fields to empty lists.
// The API always returns [] for unset dfps/network_lists/roaming_device_groups; the
// regular flatten returns null for empty slices, which conflicts with the Default=[].
func PostFlattenSecurityPolicyUDDI(ctx context.Context, planned, flattened *UDDISecurityPolicyModel, diags *diag.Diagnostics) {
	if flattened == nil {
		return
	}
	if flattened.Dfps.IsNull() {
		flattened.Dfps = types.ListValueMust(types.Int32Type, []attr.Value{})
	}
	if flattened.NetworkLists.IsNull() {
		flattened.NetworkLists = types.ListValueMust(types.Int64Type, []attr.Value{})
	}
	if flattened.RoamingDeviceGroups.IsNull() {
		flattened.RoamingDeviceGroups = types.ListValueMust(types.Int32Type, []attr.Value{})
	}
}
