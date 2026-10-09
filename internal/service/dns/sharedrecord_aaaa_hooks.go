package dns

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
)

// ValidateSharedrecordAaaa validates the SharedrecordAaaa configuration.
func ValidateSharedrecordAaaa(ctx context.Context, data SharedrecordAaaaModel, resp *resource.ValidateConfigResponse) {
	if nios := flex.ExpandNestedObject[NIOSSharedrecordAaaaModel](ctx, data.NIOS, &resp.Diagnostics); nios != nil {
		validateSharedrecordAaaaNIOSConfig(ctx, nios, resp)
	}
}

func validateSharedrecordAaaaNIOSConfig(ctx context.Context, m *NIOSSharedrecordAaaaModel, resp *resource.ValidateConfigResponse) {
}
