package dtc

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
)

// ValidateDtcMonitorTcp validates the DtcMonitorTcp configuration.
func ValidateDtcMonitorTcp(ctx context.Context, data DtcMonitorTcpModel, resp *resource.ValidateConfigResponse) {
	if nios := flex.ExpandNestedObject[NIOSDtcMonitorTcpModel](ctx, data.NIOS, &resp.Diagnostics); nios != nil {
		validateDtcMonitorTcpNIOSConfig(ctx, nios, resp)
	}
	if uddi := flex.ExpandNestedObject[UDDIDtcMonitorTcpModel](ctx, data.UDDI, &resp.Diagnostics); uddi != nil {
		validateDtcMonitorTcpUDDIConfig(ctx, uddi, resp)
	}
}

func validateDtcMonitorTcpNIOSConfig(ctx context.Context, m *NIOSDtcMonitorTcpModel, resp *resource.ValidateConfigResponse) {
}

func validateDtcMonitorTcpUDDIConfig(ctx context.Context, m *UDDIDtcMonitorTcpModel, resp *resource.ValidateConfigResponse) {
}
