package dns

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
)

// ValidateZoneAuth validates the ZoneAuth configuration.
func ValidateZoneAuth(ctx context.Context, data ZoneAuthModel, resp *resource.ValidateConfigResponse) {
	if nios := flex.ExpandNestedObject[NIOSZoneAuthModel](ctx, data.NIOS, &resp.Diagnostics); nios != nil {
		validateZoneAuthNIOSConfig(ctx, nios, resp)
	}
	if uddi := flex.ExpandNestedObject[UDDIZoneAuthModel](ctx, data.UDDI, &resp.Diagnostics); uddi != nil {
		validateZoneAuthUDDIConfig(ctx, uddi, resp)
	}
}

func validateZoneAuthNIOSConfig(ctx context.Context, m *NIOSZoneAuthModel, resp *resource.ValidateConfigResponse) {
}

func validateZoneAuthUDDIConfig(ctx context.Context, m *UDDIZoneAuthModel, resp *resource.ValidateConfigResponse) {
}
