package ipam

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
)

// ValidateNetworkview validates the Networkview configuration.
func ValidateNetworkview(ctx context.Context, data NetworkviewModel, resp *resource.ValidateConfigResponse) {
	if nios := flex.ExpandNestedObject[NIOSNetworkviewModel](ctx, data.NIOS, &resp.Diagnostics); nios != nil {
		validateNetworkviewNIOSConfig(ctx, nios, resp)
	}
	if uddi := flex.ExpandNestedObject[UDDINetworkviewModel](ctx, data.UDDI, &resp.Diagnostics); uddi != nil {
		validateNetworkviewUDDIConfig(ctx, uddi, resp)
	}
}

func validateNetworkviewNIOSConfig(ctx context.Context, m *NIOSNetworkviewModel, resp *resource.ValidateConfigResponse) {
}

func validateNetworkviewUDDIConfig(ctx context.Context, m *UDDINetworkviewModel, resp *resource.ValidateConfigResponse) {
}
