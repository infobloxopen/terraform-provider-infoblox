package clouddiscovery

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
)

// ValidateCloudDiscoveryProvider validates the CloudDiscoveryProvider configuration.
func ValidateCloudDiscoveryProvider(ctx context.Context, data CloudDiscoveryProviderModel, resp *resource.ValidateConfigResponse) {
	if uddi := flex.ExpandNestedObject[UDDICloudDiscoveryProviderModel](ctx, data.UDDI, &resp.Diagnostics); uddi != nil {
		validateCloudDiscoveryProviderUDDIConfig(ctx, uddi, resp)
	}
}

func validateCloudDiscoveryProviderUDDIConfig(ctx context.Context, m *UDDICloudDiscoveryProviderModel, resp *resource.ValidateConfigResponse) {
}
