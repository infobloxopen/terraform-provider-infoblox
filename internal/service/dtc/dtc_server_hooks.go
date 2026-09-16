package dtc

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
)

// ValidateDtcServer validates the DtcServer configuration.
func ValidateDtcServer(ctx context.Context, data DtcServerModel, resp *resource.ValidateConfigResponse) {
	if nios := flex.ExpandNestedObject[NIOSDtcServerModel](ctx, data.NIOS, &resp.Diagnostics); nios != nil {
		validateDtcServerNIOSConfig(ctx, nios, resp)
	}
	if uddi := flex.ExpandNestedObject[UDDIDtcServerModel](ctx, data.UDDI, &resp.Diagnostics); uddi != nil {
		validateDtcServerUDDIConfig(ctx, uddi, resp)
	}
}

func validateDtcServerNIOSConfig(ctx context.Context, m *NIOSDtcServerModel, resp *resource.ValidateConfigResponse) {
}

func validateDtcServerUDDIConfig(ctx context.Context, m *UDDIDtcServerModel, resp *resource.ValidateConfigResponse) {
}
