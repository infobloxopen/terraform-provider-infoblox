package fw

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
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
	if !m.DefaultRedirectName.IsNull() && !m.DefaultRedirectName.IsUnknown() && m.DefaultRedirectName.ValueString() != "" {
		if !m.DefaultAction.IsNull() && !m.DefaultAction.IsUnknown() && m.DefaultAction.ValueString() != "action_redirect" {
			resp.Diagnostics.AddAttributeError(
				path.Root("uddi").AtName("default_redirect_name"),
				"Invalid Configuration",
				`default_redirect_name can only be set when default_action is "action_redirect".`,
			)
		}
	}
}

// PostFlattenSecurityPolicyUDDI converts null list fields to their correct values after flatten.
// - dfps/network_lists/roaming_device_groups: API omits empty slices; null conflicts with Default=[].
// - access_codes: API never echoes access_codes back in create/update/read responses; preserve planned value.
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
		}
	}
}
