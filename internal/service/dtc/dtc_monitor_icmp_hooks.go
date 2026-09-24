package dtc

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
)

// ValidateDtcMonitorIcmp validates the DtcMonitorIcmp configuration.
func ValidateDtcMonitorIcmp(ctx context.Context, data DtcMonitorIcmpModel, resp *resource.ValidateConfigResponse) {
	if nios := flex.ExpandNestedObject[NIOSDtcMonitorIcmpModel](ctx, data.NIOS, &resp.Diagnostics); nios != nil {
		validateDtcMonitorIcmpNIOSConfig(ctx, nios, resp)
	}
	if uddi := flex.ExpandNestedObject[UDDIDtcMonitorIcmpModel](ctx, data.UDDI, &resp.Diagnostics); uddi != nil {
		validateDtcMonitorIcmpUDDIConfig(ctx, uddi, resp)
	}
}

func validateDtcMonitorIcmpNIOSConfig(ctx context.Context, m *NIOSDtcMonitorIcmpModel, resp *resource.ValidateConfigResponse) {
}

func validateDtcMonitorIcmpUDDIConfig(ctx context.Context, m *UDDIDtcMonitorIcmpModel, resp *resource.ValidateConfigResponse) {
}
