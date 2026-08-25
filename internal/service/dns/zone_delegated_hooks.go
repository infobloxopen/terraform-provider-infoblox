package dns

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
)

// ValidateZoneDelegated validates the ZoneDelegated configuration.
func ValidateZoneDelegated(ctx context.Context, data ZoneDelegatedModel, resp *resource.ValidateConfigResponse) {
	if nios := flex.ExpandNestedObject[NIOSZoneDelegatedModel](ctx, data.NIOS, &resp.Diagnostics); nios != nil {
		validateZoneDelegatedNIOSConfig(ctx, nios, resp)
	}
	if uddi := flex.ExpandNestedObject[UDDIZoneDelegatedModel](ctx, data.UDDI, &resp.Diagnostics); uddi != nil {
		validateZoneDelegatedUDDIConfig(ctx, uddi, resp)
	}
}

func validateZoneDelegatedNIOSConfig(ctx context.Context, m *NIOSZoneDelegatedModel, resp *resource.ValidateConfigResponse) {
}

func validateZoneDelegatedUDDIConfig(ctx context.Context, m *UDDIZoneDelegatedModel, resp *resource.ValidateConfigResponse) {
}
