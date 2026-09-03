package dhcp

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
)

// ValidateIpv6fixedaddresstemplate validates the Ipv6fixedaddresstemplate configuration.
func ValidateIpv6fixedaddresstemplate(ctx context.Context, data Ipv6fixedaddresstemplateModel, resp *resource.ValidateConfigResponse) {
	if nios := flex.ExpandNestedObject[NIOSIpv6fixedaddresstemplateModel](ctx, data.NIOS, &resp.Diagnostics); nios != nil {
		validateIpv6fixedaddresstemplateNIOSConfig(ctx, nios, resp)
	}
}

func validateIpv6fixedaddresstemplateNIOSConfig(ctx context.Context, m *NIOSIpv6fixedaddresstemplateModel, resp *resource.ValidateConfigResponse) {
}
