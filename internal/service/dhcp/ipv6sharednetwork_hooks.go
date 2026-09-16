package dhcp

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
)

// ValidateIpv6sharednetwork validates the Ipv6sharednetwork configuration.
func ValidateIpv6sharednetwork(ctx context.Context, data Ipv6sharednetworkModel, resp *resource.ValidateConfigResponse) {
	if nios := flex.ExpandNestedObject[NIOSIpv6sharednetworkModel](ctx, data.NIOS, &resp.Diagnostics); nios != nil {
		validateIpv6sharednetworkNIOSConfig(ctx, nios, resp)
	}
}

func validateIpv6sharednetworkNIOSConfig(ctx context.Context, m *NIOSIpv6sharednetworkModel, resp *resource.ValidateConfigResponse) {
}
