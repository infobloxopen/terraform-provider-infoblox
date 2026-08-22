package dhcp

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
)

// ValidateIpv6DhcpOptiondefinition validates the Ipv6DhcpOptiondefinition configuration.
func ValidateIpv6DhcpOptiondefinition(ctx context.Context, data Ipv6DhcpOptiondefinitionModel, resp *resource.ValidateConfigResponse) {
	if nios := flex.ExpandNestedObject[NIOSIpv6DhcpOptiondefinitionModel](ctx, data.NIOS, &resp.Diagnostics); nios != nil {
		validateIpv6DhcpOptiondefinitionNIOSConfig(ctx, nios, resp)
	}
	if uddi := flex.ExpandNestedObject[UDDIIpv6DhcpOptiondefinitionModel](ctx, data.UDDI, &resp.Diagnostics); uddi != nil {
		validateIpv6DhcpOptiondefinitionUDDIConfig(ctx, uddi, resp)
	}
}

func validateIpv6DhcpOptiondefinitionNIOSConfig(ctx context.Context, m *NIOSIpv6DhcpOptiondefinitionModel, resp *resource.ValidateConfigResponse) {
}

func validateIpv6DhcpOptiondefinitionUDDIConfig(ctx context.Context, m *UDDIIpv6DhcpOptiondefinitionModel, resp *resource.ValidateConfigResponse) {
}
