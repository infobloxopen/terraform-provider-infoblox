package ipam

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
)

// ValidateSuperhost validates the Superhost configuration.
func ValidateSuperhost(ctx context.Context, data SuperhostModel, resp *resource.ValidateConfigResponse) {
	if nios := flex.ExpandNestedObject[NIOSSuperhostModel](ctx, data.NIOS, &resp.Diagnostics); nios != nil {
		validateSuperhostNIOSConfig(ctx, nios, resp)
	}
}

func validateSuperhostNIOSConfig(ctx context.Context, m *NIOSSuperhostModel, resp *resource.ValidateConfigResponse) {
}
