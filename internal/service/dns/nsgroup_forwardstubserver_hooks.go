package dns

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
)

// ValidateNsgroupForwardstubserver validates the NsgroupForwardstubserver configuration.
func ValidateNsgroupForwardstubserver(ctx context.Context, data NsgroupForwardstubserverModel, resp *resource.ValidateConfigResponse) {
	if nios := flex.ExpandNestedObject[NIOSNsgroupForwardstubserverModel](ctx, data.NIOS, &resp.Diagnostics); nios != nil {
		validateNsgroupForwardstubserverNIOSConfig(ctx, nios, resp)
	}
}

func validateNsgroupForwardstubserverNIOSConfig(ctx context.Context, m *NIOSNsgroupForwardstubserverModel, resp *resource.ValidateConfigResponse) {
}
