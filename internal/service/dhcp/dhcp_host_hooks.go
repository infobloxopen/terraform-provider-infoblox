package dhcp

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
)

// ValidateDhcpHost validates the DhcpHost configuration.
func ValidateDhcpHost(ctx context.Context, data DhcpHostModel, resp *resource.ValidateConfigResponse) {
	if uddi := flex.ExpandNestedObject[UDDIDhcpHostModel](ctx, data.UDDI, &resp.Diagnostics); uddi != nil {
		validateDhcpHostUDDIConfig(ctx, uddi, resp)
	}
}

func validateDhcpHostUDDIConfig(ctx context.Context, m *UDDIDhcpHostModel, resp *resource.ValidateConfigResponse) {
}
