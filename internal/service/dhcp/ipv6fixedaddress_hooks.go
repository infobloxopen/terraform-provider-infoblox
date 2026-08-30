package dhcp

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
)

// ValidateIpv6fixedaddress validates the Ipv6fixedaddress configuration.
func ValidateIpv6fixedaddress(ctx context.Context, data Ipv6fixedaddressModel, resp *resource.ValidateConfigResponse) {
	if nios := flex.ExpandNestedObject[NIOSIpv6fixedaddressModel](ctx, data.NIOS, &resp.Diagnostics); nios != nil {
		validateIpv6fixedaddressNIOSConfig(ctx, nios, resp)
	}
	if uddi := flex.ExpandNestedObject[UDDIIpv6fixedaddressModel](ctx, data.UDDI, &resp.Diagnostics); uddi != nil {
		validateIpv6fixedaddressUDDIConfig(ctx, uddi, resp)
	}
}

func validateIpv6fixedaddressNIOSConfig(ctx context.Context, m *NIOSIpv6fixedaddressModel, resp *resource.ValidateConfigResponse) {
}

func validateIpv6fixedaddressUDDIConfig(ctx context.Context, m *UDDIIpv6fixedaddressModel, resp *resource.ValidateConfigResponse) {
}
