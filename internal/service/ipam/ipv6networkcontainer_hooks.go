package ipam

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
)

// ValidateIpv6networkcontainer validates the Ipv6networkcontainer configuration.
func ValidateIpv6networkcontainer(ctx context.Context, data Ipv6networkcontainerModel, resp *resource.ValidateConfigResponse) {
	if nios := flex.ExpandNestedObject[NIOSIpv6networkcontainerModel](ctx, data.NIOS, &resp.Diagnostics); nios != nil {
		validateIpv6networkcontainerNIOSConfig(ctx, nios, resp)
	}
}

func validateIpv6networkcontainerNIOSConfig(ctx context.Context, m *NIOSIpv6networkcontainerModel, resp *resource.ValidateConfigResponse) {
}
