package ipam

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/infobloxopen/infoblox-nios-go-client/ipam"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
)

// ValidateVlan validates the Vlan configuration.
func ValidateVlan(ctx context.Context, data VlanModel, resp *resource.ValidateConfigResponse) {
	if nios := flex.ExpandNestedObject[NIOSVlanModel](ctx, data.NIOS, &resp.Diagnostics); nios != nil {
		validateVlanNIOSConfig(ctx, nios, resp)
	}
}

func validateVlanNIOSConfig(ctx context.Context, m *NIOSVlanModel, resp *resource.ValidateConfigResponse) {
}

func ExpandVlanParent(str types.String) *ipam.VlanParent {
	if str.IsNull() {
		return &ipam.VlanParent{}
	}
	var m ipam.VlanParent
	m.String = flex.ExpandStringPointer(str)

	return &m
}

func FlattenVlanParent(from *ipam.VlanParent) types.String {
	if from.VlanParentOneOf == nil {
		return types.StringNull()
	}
	m := flex.FlattenStringPointer(from.VlanParentOneOf.Ref)
	return m
}

func ExpandVlanId(val types.Int64) *ipam.VlanId {
	if val.IsNull() {
		return &ipam.VlanId{}
	}
	var m ipam.VlanId
	m.Int64 = flex.ExpandInt64Pointer(val)

	return &m
}

func FlattenVlanId(from *ipam.VlanId) types.Int64 {
	if from.Int64 == nil {
		return types.Int64Null()
	}
	m := flex.FlattenInt64Pointer(from.Int64)
	return m
}
