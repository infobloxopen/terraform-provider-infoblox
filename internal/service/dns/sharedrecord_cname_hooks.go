package dns

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
)

// ValidateSharedrecordCname validates the SharedrecordCname configuration.
func ValidateSharedrecordCname(ctx context.Context, data SharedrecordCnameModel, resp *resource.ValidateConfigResponse) {
	if nios := flex.ExpandNestedObject[NIOSSharedrecordCnameModel](ctx, data.NIOS, &resp.Diagnostics); nios != nil {
		validateSharedrecordCnameNIOSConfig(ctx, nios, resp)
	}
}

func validateSharedrecordCnameNIOSConfig(ctx context.Context, m *NIOSSharedrecordCnameModel, resp *resource.ValidateConfigResponse) {
}
