package dtc

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
)

// ValidateDtcPool validates the DtcPool configuration.
func ValidateDtcPool(ctx context.Context, data DtcPoolModel, resp *resource.ValidateConfigResponse) {
	if nios := flex.ExpandNestedObject[NIOSDtcPoolModel](ctx, data.NIOS, &resp.Diagnostics); nios != nil {
		validateDtcPoolNIOSConfig(ctx, nios, resp)
	}
	if uddi := flex.ExpandNestedObject[UDDIDtcPoolModel](ctx, data.UDDI, &resp.Diagnostics); uddi != nil {
		validateDtcPoolUDDIConfig(ctx, uddi, resp)
	}
}

func validateDtcPoolNIOSConfig(ctx context.Context, m *NIOSDtcPoolModel, resp *resource.ValidateConfigResponse) {
}

func validateDtcPoolUDDIConfig(ctx context.Context, m *UDDIDtcPoolModel, resp *resource.ValidateConfigResponse) {
}
