package ipamfederation

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
)

// ValidateOverlappingBlock validates the OverlappingBlock configuration.
func ValidateOverlappingBlock(ctx context.Context, data OverlappingBlockModel, resp *resource.ValidateConfigResponse) {
	if uddi := flex.ExpandNestedObject[UDDIOverlappingBlockModel](ctx, data.UDDI, &resp.Diagnostics); uddi != nil {
		validateOverlappingBlockUDDIConfig(ctx, uddi, resp)
	}
}

func validateOverlappingBlockUDDIConfig(ctx context.Context, m *UDDIOverlappingBlockModel, resp *resource.ValidateConfigResponse) {
}
