package dns

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
)

// ValidateZoneForward validates the ZoneForward configuration.
func ValidateZoneForward(ctx context.Context, data ZoneForwardModel, resp *resource.ValidateConfigResponse) {
	if nios := flex.ExpandNestedObject[NIOSZoneForwardModel](ctx, data.NIOS, &resp.Diagnostics); nios != nil {
		validateZoneForwardNIOSConfig(ctx, nios, resp)
	}
	if uddi := flex.ExpandNestedObject[UDDIZoneForwardModel](ctx, data.UDDI, &resp.Diagnostics); uddi != nil {
		validateZoneForwardUDDIConfig(ctx, uddi, resp)
	}
}

func validateZoneForwardNIOSConfig(ctx context.Context, m *NIOSZoneForwardModel, resp *resource.ValidateConfigResponse) {
}

func validateZoneForwardUDDIConfig(ctx context.Context, m *UDDIZoneForwardModel, resp *resource.ValidateConfigResponse) {
}

func PostFlattenZoneForwardNIOS(ctx context.Context, planned, flattened *NIOSZoneForwardModel, diags *diag.Diagnostics) {
}
