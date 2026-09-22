package ipam

import (
	"context"

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
