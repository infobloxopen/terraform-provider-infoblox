package dns

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
)

// ValidateSharedrecordTxt validates the SharedrecordTxt configuration.
func ValidateSharedrecordTxt(ctx context.Context, data SharedrecordTxtModel, resp *resource.ValidateConfigResponse) {
	if nios := flex.ExpandNestedObject[NIOSSharedrecordTxtModel](ctx, data.NIOS, &resp.Diagnostics); nios != nil {
		validateSharedrecordTxtNIOSConfig(ctx, nios, resp)
	}
}

func validateSharedrecordTxtNIOSConfig(ctx context.Context, m *NIOSSharedrecordTxtModel, resp *resource.ValidateConfigResponse) {
}
