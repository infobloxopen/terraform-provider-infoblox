package dtc

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
)

// ValidateDtcLbdn validates the DtcLbdn configuration.
func ValidateDtcLbdn(ctx context.Context, data DtcLbdnModel, resp *resource.ValidateConfigResponse) {
	if nios := flex.ExpandNestedObject[NIOSDtcLbdnModel](ctx, data.NIOS, &resp.Diagnostics); nios != nil {
		validateDtcLbdnNIOSConfig(ctx, nios, resp)
	}
	if uddi := flex.ExpandNestedObject[UDDIDtcLbdnModel](ctx, data.UDDI, &resp.Diagnostics); uddi != nil {
		validateDtcLbdnUDDIConfig(ctx, uddi, resp)
	}
}

func validateDtcLbdnNIOSConfig(ctx context.Context, m *NIOSDtcLbdnModel, resp *resource.ValidateConfigResponse) {
}

func validateDtcLbdnUDDIConfig(ctx context.Context, m *UDDIDtcLbdnModel, resp *resource.ValidateConfigResponse) {
}
