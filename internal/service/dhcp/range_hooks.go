package dhcp

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
)

// ValidateRange validates the Range configuration.
func ValidateRange(ctx context.Context, data RangeModel, resp *resource.ValidateConfigResponse) {
	if nios := flex.ExpandNestedObject[NIOSRangeModel](ctx, data.NIOS, &resp.Diagnostics); nios != nil {
		validateRangeNIOSConfig(ctx, nios, resp)
	}
	if uddi := flex.ExpandNestedObject[UDDIRangeModel](ctx, data.UDDI, &resp.Diagnostics); uddi != nil {
		validateRangeUDDIConfig(ctx, uddi, resp)
	}
}

func validateRangeNIOSConfig(ctx context.Context, m *NIOSRangeModel, resp *resource.ValidateConfigResponse) {
}

func validateRangeUDDIConfig(ctx context.Context, m *UDDIRangeModel, resp *resource.ValidateConfigResponse) {
}

func PostFlattenRangeNIOS(ctx context.Context, planned, flattened *NIOSRangeModel, diags *diag.Diagnostics) {
}
