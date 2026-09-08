package dtc

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
)

// ValidateDtcMonitorHttp validates the DtcMonitorHttp configuration.
func ValidateDtcMonitorHttp(ctx context.Context, data DtcMonitorHttpModel, resp *resource.ValidateConfigResponse) {
	if nios := flex.ExpandNestedObject[NIOSDtcMonitorHttpModel](ctx, data.NIOS, &resp.Diagnostics); nios != nil {
		validateDtcMonitorHttpNIOSConfig(ctx, nios, resp)
	}
	if uddi := flex.ExpandNestedObject[UDDIDtcMonitorHttpModel](ctx, data.UDDI, &resp.Diagnostics); uddi != nil {
		validateDtcMonitorHttpUDDIConfig(ctx, uddi, resp)
	}
}

func validateDtcMonitorHttpNIOSConfig(ctx context.Context, m *NIOSDtcMonitorHttpModel, resp *resource.ValidateConfigResponse) {
}

func validateDtcMonitorHttpUDDIConfig(ctx context.Context, m *UDDIDtcMonitorHttpModel, resp *resource.ValidateConfigResponse) {
}
