package dns

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
)

// ValidateForwardNsg validates the ForwardNsg configuration.
func ValidateForwardNsg(ctx context.Context, data ForwardNsgModel, resp *resource.ValidateConfigResponse) {
	if uddi := flex.ExpandNestedObject[UDDIForwardNsgModel](ctx, data.UDDI, &resp.Diagnostics); uddi != nil {
		validateForwardNsgUDDIConfig(ctx, uddi, resp)
	}
}

func validateForwardNsgUDDIConfig(ctx context.Context, m *UDDIForwardNsgModel, resp *resource.ValidateConfigResponse) {
}
