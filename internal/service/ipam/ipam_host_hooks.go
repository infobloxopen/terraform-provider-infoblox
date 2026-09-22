package ipam

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
)

// ValidateIpamHost validates the IpamHost configuration.
func ValidateIpamHost(ctx context.Context, data IpamHostModel, resp *resource.ValidateConfigResponse) {
	if uddi := flex.ExpandNestedObject[UDDIIpamHostModel](ctx, data.UDDI, &resp.Diagnostics); uddi != nil {
		validateIpamHostUDDIConfig(ctx, uddi, resp)
	}
}

func validateIpamHostUDDIConfig(ctx context.Context, m *UDDIIpamHostModel, resp *resource.ValidateConfigResponse) {
}
