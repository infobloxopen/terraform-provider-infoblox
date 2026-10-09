package dhcp

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
)

// ValidateDhcpOptionspace validates the DhcpOptionspace configuration.
func ValidateDhcpOptionspace(ctx context.Context, data DhcpOptionspaceModel, resp *resource.ValidateConfigResponse) {
	if nios := flex.ExpandNestedObject[NIOSDhcpOptionspaceModel](ctx, data.NIOS, &resp.Diagnostics); nios != nil {
		validateDhcpOptionspaceNIOSConfig(ctx, nios, resp)
	}
	if uddi := flex.ExpandNestedObject[UDDIDhcpOptionspaceModel](ctx, data.UDDI, &resp.Diagnostics); uddi != nil {
		validateDhcpOptionspaceUDDIConfig(ctx, uddi, resp)
	}
}

func validateDhcpOptionspaceNIOSConfig(ctx context.Context, m *NIOSDhcpOptionspaceModel, resp *resource.ValidateConfigResponse) {
}

func validateDhcpOptionspaceUDDIConfig(ctx context.Context, m *UDDIDhcpOptionspaceModel, resp *resource.ValidateConfigResponse) {
}
