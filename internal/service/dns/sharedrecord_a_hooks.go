package dns

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
)

// ValidateSharedrecordA validates the SharedrecordA configuration.
func ValidateSharedrecordA(ctx context.Context, data SharedrecordAModel, resp *resource.ValidateConfigResponse) {
	if nios := flex.ExpandNestedObject[NIOSSharedrecordAModel](ctx, data.NIOS, &resp.Diagnostics); nios != nil {
		validateSharedrecordANIOSConfig(ctx, nios, resp)
	}
}

func validateSharedrecordANIOSConfig(ctx context.Context, m *NIOSSharedrecordAModel, resp *resource.ValidateConfigResponse) {
}
