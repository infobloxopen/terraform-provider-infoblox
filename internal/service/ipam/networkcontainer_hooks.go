package ipam

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
)

// ValidateNetworkcontainer validates the Networkcontainer configuration.
func ValidateNetworkcontainer(ctx context.Context, data NetworkcontainerModel, resp *resource.ValidateConfigResponse) {
	if nios := flex.ExpandNestedObject[NIOSNetworkcontainerModel](ctx, data.NIOS, &resp.Diagnostics); nios != nil {
		validateNetworkcontainerNIOSConfig(ctx, nios, resp)
	}
	if uddi := flex.ExpandNestedObject[UDDINetworkcontainerModel](ctx, data.UDDI, &resp.Diagnostics); uddi != nil {
		validateNetworkcontainerUDDIConfig(ctx, uddi, resp)
	}
}

func validateNetworkcontainerNIOSConfig(ctx context.Context, m *NIOSNetworkcontainerModel, resp *resource.ValidateConfigResponse) {
}

func validateNetworkcontainerUDDIConfig(ctx context.Context, m *UDDINetworkcontainerModel, resp *resource.ValidateConfigResponse) {
}
