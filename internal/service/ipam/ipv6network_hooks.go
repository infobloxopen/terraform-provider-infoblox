package ipam

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
)

// ValidateIpv6network validates the Ipv6network configuration.
func ValidateIpv6network(ctx context.Context, data Ipv6networkModel, resp *resource.ValidateConfigResponse) {
	if nios := flex.ExpandNestedObject[NIOSIpv6networkModel](ctx, data.NIOS, &resp.Diagnostics); nios != nil {
		validateIpv6networkNIOSConfig(ctx, nios, resp)
	}
	if uddi := flex.ExpandNestedObject[UDDIIpv6networkModel](ctx, data.UDDI, &resp.Diagnostics); uddi != nil {
		validateIpv6networkUDDIConfig(ctx, uddi, resp)
	}
}

func validateIpv6networkNIOSConfig(ctx context.Context, m *NIOSIpv6networkModel, resp *resource.ValidateConfigResponse) {
}

func validateIpv6networkUDDIConfig(ctx context.Context, m *UDDIIpv6networkModel, resp *resource.ValidateConfigResponse) {
}
