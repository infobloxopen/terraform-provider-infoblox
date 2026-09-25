package dhcp

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
)

// ValidateHardwareFilter validates the HardwareFilter configuration.
func ValidateHardwareFilter(ctx context.Context, data HardwareFilterModel, resp *resource.ValidateConfigResponse) {
	if uddi := flex.ExpandNestedObject[UDDIHardwareFilterModel](ctx, data.UDDI, &resp.Diagnostics); uddi != nil {
		validateHardwareFilterUDDIConfig(ctx, uddi, resp)
	}
}

func validateHardwareFilterUDDIConfig(ctx context.Context, m *UDDIHardwareFilterModel, resp *resource.ValidateConfigResponse) {
}
