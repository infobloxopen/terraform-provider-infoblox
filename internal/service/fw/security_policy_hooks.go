package fw

import (
	"context"
	"sort"

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

// PostFlattenSecurityPolicyUDDI converts null int list fields to empty lists and
// preserves access_codes from the plan when the API omits them on read.
// The API always returns [] for unset dfps/network_lists/roaming_device_groups; the
// regular flatten returns null for empty slices, which conflicts with the Default=[].
// access_codes order in the API response may differ from the plan/state order; we
// reorder the returned list to match the plan/state reference to keep state stable.
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
	if planned != nil && !planned.AccessCodes.IsNull() && !planned.AccessCodes.IsUnknown() {
		if flattened.AccessCodes.IsNull() {
			// API omitted access_codes; preserve the plan/state values.
			flattened.AccessCodes = planned.AccessCodes
		} else {
			// API returned access_codes but may be in a different order; reorder to
			// match the plan/state so the state is stable and plan-consistent.
			flattened.AccessCodes = reorderStringListToMatch(flattened.AccessCodes, planned.AccessCodes)
		}
	}
}

// reorderStringListToMatch reorders src elements to follow the order in ref.
// Elements present in ref appear first in ref order; extra elements in src
// (not in ref) are appended in their original order at the end.
func reorderStringListToMatch(src, ref types.List) types.List {
	if src.IsNull() || src.IsUnknown() || ref.IsNull() || ref.IsUnknown() {
		return src
	}
	refElems := ref.Elements()
	refOrder := make(map[string]int, len(refElems))
	for i, e := range refElems {
		refOrder[e.(types.String).ValueString()] = i
	}
	srcElems := src.Elements()
	sort.SliceStable(srcElems, func(i, j int) bool {
		vi := srcElems[i].(types.String).ValueString()
		vj := srcElems[j].(types.String).ValueString()
		idxI, okI := refOrder[vi]
		idxJ, okJ := refOrder[vj]
		if okI && okJ {
			return idxI < idxJ
		}
		if okI {
			return true
		}
		if okJ {
			return false
		}
		return vi < vj
	})
	return types.ListValueMust(types.StringType, srcElems)
}
