package dns

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
)

// ValidateSharedrecordSrv validates the SharedrecordSrv configuration.
func ValidateSharedrecordSrv(ctx context.Context, data SharedrecordSrvModel, resp *resource.ValidateConfigResponse) {
	if nios := flex.ExpandNestedObject[NIOSSharedrecordSrvModel](ctx, data.NIOS, &resp.Diagnostics); nios != nil {
		validateSharedrecordSrvNIOSConfig(ctx, nios, resp)
	}
}

func validateSharedrecordSrvNIOSConfig(ctx context.Context, m *NIOSSharedrecordSrvModel, resp *resource.ValidateConfigResponse) {
}
