package dhcp

import (
	"context"
	"strings"

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

const (
	niosDefaultIpv6OptionSpace = "DHCPv6"
	niosIpv6OptionNamePrefix   = "dhcp6."
)

func validateIpv6DhcpOptiondefinitionNIOSConfig(ctx context.Context, m *NIOSIpv6DhcpOptiondefinitionModel, resp *resource.ValidateConfigResponse) {
	if m.Space.IsUnknown() || m.Name.IsUnknown() || m.Name.IsNull() {
		return
	}

	space := niosDefaultIpv6OptionSpace
	if !m.Space.IsNull() {
		space = m.Space.ValueString()
	}
	hasPrefix := strings.HasPrefix(m.Name.ValueString(), niosIpv6OptionNamePrefix)

	switch {
	case space == niosDefaultIpv6OptionSpace && !hasPrefix:
		resp.Diagnostics.AddError(
			"Invalid Name for DHCPv6 Option Definition",
			"The name of a DHCP IPv6 option definition object in the default space (DHCPv6) must start with 'dhcp6.'.",
		)
	case space != niosDefaultIpv6OptionSpace && hasPrefix:
		resp.Diagnostics.AddError(
			"Invalid Name for Custom DHCPv6 Option Definition",
			"The name of a DHCP IPv6 option definition object in a custom space must not start with 'dhcp6.'.",
		)
	}
}

func validateIpv6DhcpOptiondefinitionUDDIConfig(ctx context.Context, m *UDDIIpv6DhcpOptiondefinitionModel, resp *resource.ValidateConfigResponse) {
}
