package dhcp

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
)

// ValidateDhcpOptiondefinition validates the DhcpOptiondefinition configuration.
func ValidateDhcpOptiondefinition(ctx context.Context, data DhcpOptiondefinitionModel, resp *resource.ValidateConfigResponse) {
	if nios := flex.ExpandNestedObject[NIOSDhcpOptiondefinitionModel](ctx, data.NIOS, &resp.Diagnostics); nios != nil {
		validateDhcpOptiondefinitionNIOSConfig(ctx, nios, resp)
	}
	if uddi := flex.ExpandNestedObject[UDDIDhcpOptiondefinitionModel](ctx, data.UDDI, &resp.Diagnostics); uddi != nil {
		validateDhcpOptiondefinitionUDDIConfig(ctx, uddi, resp)
	}
}

func validateDhcpOptiondefinitionNIOSConfig(ctx context.Context, m *NIOSDhcpOptiondefinitionModel, resp *resource.ValidateConfigResponse) {
}

func validateDhcpOptiondefinitionUDDIConfig(ctx context.Context, m *UDDIDhcpOptiondefinitionModel, resp *resource.ValidateConfigResponse) {
}
