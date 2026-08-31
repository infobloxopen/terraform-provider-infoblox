package dhcp

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
)

// ValidateIpv6DhcpOptionspace validates the Ipv6DhcpOptionspace configuration.
func ValidateIpv6DhcpOptionspace(ctx context.Context, data Ipv6DhcpOptionspaceModel, resp *resource.ValidateConfigResponse) {
	if nios := flex.ExpandNestedObject[NIOSIpv6DhcpOptionspaceModel](ctx, data.NIOS, &resp.Diagnostics); nios != nil {
		validateIpv6DhcpOptionspaceNIOSConfig(ctx, nios, resp)
	}
	if uddi := flex.ExpandNestedObject[UDDIIpv6DhcpOptionspaceModel](ctx, data.UDDI, &resp.Diagnostics); uddi != nil {
		validateIpv6DhcpOptionspaceUDDIConfig(ctx, uddi, resp)
	}
}

func validateIpv6DhcpOptionspaceNIOSConfig(ctx context.Context, m *NIOSIpv6DhcpOptionspaceModel, resp *resource.ValidateConfigResponse) {
}

func validateIpv6DhcpOptionspaceUDDIConfig(ctx context.Context, m *UDDIIpv6DhcpOptionspaceModel, resp *resource.ValidateConfigResponse) {
}
