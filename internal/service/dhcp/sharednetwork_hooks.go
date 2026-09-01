package dhcp

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
)

// ValidateSharednetwork validates the Sharednetwork configuration.
func ValidateSharednetwork(ctx context.Context, data SharednetworkModel, resp *resource.ValidateConfigResponse) {
	if nios := flex.ExpandNestedObject[NIOSSharednetworkModel](ctx, data.NIOS, &resp.Diagnostics); nios != nil {
		validateSharednetworkNIOSConfig(ctx, nios, resp)
	}
}

func validateSharednetworkNIOSConfig(ctx context.Context, m *NIOSSharednetworkModel, resp *resource.ValidateConfigResponse) {
}
