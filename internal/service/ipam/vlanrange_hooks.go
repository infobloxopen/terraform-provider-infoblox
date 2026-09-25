package ipam

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/infobloxopen/infoblox-nios-go-client/ipam"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
)

// ValidateVlanrange validates the Vlanrange configuration.
func ValidateVlanrange(ctx context.Context, data VlanrangeModel, resp *resource.ValidateConfigResponse) {
	if nios := flex.ExpandNestedObject[NIOSVlanrangeModel](ctx, data.NIOS, &resp.Diagnostics); nios != nil {
		validateVlanrangeNIOSConfig(ctx, nios, resp)
	}
}

func validateVlanrangeNIOSConfig(ctx context.Context, m *NIOSVlanrangeModel, resp *resource.ValidateConfigResponse) {
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

func ExpandVlanrangeVlanView(str types.String) *ipam.VlanrangeVlanView {
	if str.IsNull() {
		return &ipam.VlanrangeVlanView{}
	}
	var m ipam.VlanrangeVlanView
	m.String = flex.ExpandStringPointer(str)

	return &m
}

func FlattenVlanrangeVlanView(from *ipam.VlanrangeVlanView) types.String {
	if from == nil || from.VlanrangeVlanViewOneOf == nil {
		return types.StringNull()
	}
	m := flex.FlattenStringPointer(from.VlanrangeVlanViewOneOf.Ref)
	return m
}

func PostFlattenVlanrangeNIOS(ctx context.Context, planned, flattened *NIOSVlanrangeModel, diags *diag.Diagnostics) {
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
