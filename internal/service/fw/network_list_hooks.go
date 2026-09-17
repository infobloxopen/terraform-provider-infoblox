package fw

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
)

// ValidateNetworkList validates the NetworkList configuration.
func ValidateNetworkList(ctx context.Context, data NetworkListModel, resp *resource.ValidateConfigResponse) {
	if uddi := flex.ExpandNestedObject[UDDINetworkListModel](ctx, data.UDDI, &resp.Diagnostics); uddi != nil {
		validateNetworkListUDDIConfig(ctx, uddi, resp)
	}
}

func validateNetworkListUDDIConfig(ctx context.Context, m *UDDINetworkListModel, resp *resource.ValidateConfigResponse) {
}
