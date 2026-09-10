package dtc

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
)

// ValidateDtcMonitorSnmp validates the DtcMonitorSnmp configuration.
func ValidateDtcMonitorSnmp(ctx context.Context, data DtcMonitorSnmpModel, resp *resource.ValidateConfigResponse) {
	if nios := flex.ExpandNestedObject[NIOSDtcMonitorSnmpModel](ctx, data.NIOS, &resp.Diagnostics); nios != nil {
		validateDtcMonitorSnmpNIOSConfig(ctx, nios, resp)
	}
	if uddi := flex.ExpandNestedObject[UDDIDtcMonitorSnmpModel](ctx, data.UDDI, &resp.Diagnostics); uddi != nil {
		validateDtcMonitorSnmpUDDIConfig(ctx, uddi, resp)
	}
}

func validateDtcMonitorSnmpNIOSConfig(ctx context.Context, m *NIOSDtcMonitorSnmpModel, resp *resource.ValidateConfigResponse) {
}

func validateDtcMonitorSnmpUDDIConfig(ctx context.Context, m *UDDIDtcMonitorSnmpModel, resp *resource.ValidateConfigResponse) {
}
