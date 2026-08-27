package dtc

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/path"
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
	if m.Availability.IsNull() || m.Availability.IsUnknown() {
		return
	}

	if m.Availability.ValueString() == "QUORUM" {
		if m.Quorum.IsNull() {
			resp.Diagnostics.AddAttributeError(
				path.Root("nios").AtName("quorum"),
				"Missing Required Attribute",
				"When nios.availability is set to 'QUORUM', the 'nios.quorum' attribute must be specified.",
			)
		}
	} else {
		if !m.Quorum.IsNull() && !m.Quorum.IsUnknown() {
			resp.Diagnostics.AddAttributeError(
				path.Root("nios").AtName("quorum"),
				"Invalid Attribute Combination",
				fmt.Sprintf("The 'nios.quorum' attribute can only be set when 'nios.availability' is 'QUORUM', but got '%s'.", m.Availability.ValueString()),
			)
		}
	}
}

func validateDtcPoolUDDIConfig(ctx context.Context, m *UDDIDtcPoolModel, resp *resource.ValidateConfigResponse) {
}
