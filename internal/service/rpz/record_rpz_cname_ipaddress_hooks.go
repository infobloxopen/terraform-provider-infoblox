package rpz

import (
	"context"
	"fmt"
	"net/netip"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
)

// ValidateRecordRpzCnameIpaddress validates the RecordRpzCnameIpaddress configuration.
func ValidateRecordRpzCnameIpaddress(ctx context.Context, data RecordRpzCnameIpaddressModel, resp *resource.ValidateConfigResponse) {
	if nios := flex.ExpandNestedObject[NIOSRecordRpzCnameIpaddressModel](ctx, data.NIOS, &resp.Diagnostics); nios != nil {
		validateRecordRpzCnameIpaddressNIOSConfig(ctx, nios, resp)
	}
}

func validateRecordRpzCnameIpaddressNIOSConfig(ctx context.Context, m *NIOSRecordRpzCnameIpaddressModel, resp *resource.ValidateConfigResponse) {
	if m.Canonical.IsUnknown() {
		return
	}

	canonical := m.Canonical.ValueString()
	// Empty string is a valid passthru rule in NIOS RPZ
	if canonical != "*" && canonical != "" {
		if _, err := netip.ParseAddr(canonical); err != nil {
			resp.Diagnostics.AddError(
				"Invalid Canonical Value",
				fmt.Sprintf("The canonical value must be '*', '' (passthru), or a valid IP address. Got: %s", canonical),
			)
		}
	}
}
