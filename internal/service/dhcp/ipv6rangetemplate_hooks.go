package dhcp

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
)

// ValidateIpv6rangetemplate validates the Ipv6rangetemplate configuration.
func ValidateIpv6rangetemplate(ctx context.Context, data Ipv6rangetemplateModel, resp *resource.ValidateConfigResponse) {
	if nios := flex.ExpandNestedObject[NIOSIpv6rangetemplateModel](ctx, data.NIOS, &resp.Diagnostics); nios != nil {
		validateIpv6rangetemplateNIOSConfig(ctx, nios, resp)
	}
}

func validateIpv6rangetemplateNIOSConfig(ctx context.Context, m *NIOSIpv6rangetemplateModel, resp *resource.ValidateConfigResponse) {
}
