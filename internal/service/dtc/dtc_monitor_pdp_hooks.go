package dtc

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
)

// ValidateDtcMonitorPdp validates the DtcMonitorPdp configuration.
func ValidateDtcMonitorPdp(ctx context.Context, data DtcMonitorPdpModel, resp *resource.ValidateConfigResponse) {
	if nios := flex.ExpandNestedObject[NIOSDtcMonitorPdpModel](ctx, data.NIOS, &resp.Diagnostics); nios != nil {
		validateDtcMonitorPdpNIOSConfig(ctx, nios, resp)
	}
	if uddi := flex.ExpandNestedObject[UDDIDtcMonitorPdpModel](ctx, data.UDDI, &resp.Diagnostics); uddi != nil {
		validateDtcMonitorPdpUDDIConfig(ctx, uddi, resp)
	}
}

func validateDtcMonitorPdpNIOSConfig(ctx context.Context, m *NIOSDtcMonitorPdpModel, resp *resource.ValidateConfigResponse) {
}

func validateDtcMonitorPdpUDDIConfig(ctx context.Context, m *UDDIDtcMonitorPdpModel, resp *resource.ValidateConfigResponse) {
}
