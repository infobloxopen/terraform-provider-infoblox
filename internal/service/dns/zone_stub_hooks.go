package dns

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
)

// ValidateZoneStub validates the ZoneStub configuration.
func ValidateZoneStub(ctx context.Context, data ZoneStubModel, resp *resource.ValidateConfigResponse) {
	if nios := flex.ExpandNestedObject[NIOSZoneStubModel](ctx, data.NIOS, &resp.Diagnostics); nios != nil {
		validateZoneStubNIOSConfig(ctx, nios, resp)
	}
}

func validateZoneStubNIOSConfig(ctx context.Context, m *NIOSZoneStubModel, resp *resource.ValidateConfigResponse) {
}
