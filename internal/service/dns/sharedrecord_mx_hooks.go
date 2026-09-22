package dns

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
)

// ValidateSharedrecordMx validates the SharedrecordMx configuration.
func ValidateSharedrecordMx(ctx context.Context, data SharedrecordMxModel, resp *resource.ValidateConfigResponse) {
	if nios := flex.ExpandNestedObject[NIOSSharedrecordMxModel](ctx, data.NIOS, &resp.Diagnostics); nios != nil {
		validateSharedrecordMxNIOSConfig(ctx, nios, resp)
	}
}

func validateSharedrecordMxNIOSConfig(ctx context.Context, m *NIOSSharedrecordMxModel, resp *resource.ValidateConfigResponse) {
}
