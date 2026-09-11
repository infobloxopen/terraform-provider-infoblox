package rpz

import (
	"context"
	"fmt"
	"net/netip"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
)

// ValidateRecordRpzCnameClientipaddress validates the RecordRpzCnameClientipaddress configuration.
func ValidateRecordRpzCnameClientipaddress(ctx context.Context, data RecordRpzCnameClientipaddressModel, resp *resource.ValidateConfigResponse) {
	if nios := flex.ExpandNestedObject[NIOSRecordRpzCnameClientipaddressModel](ctx, data.NIOS, &resp.Diagnostics); nios != nil {
		validateRecordRpzCnameClientipaddressNIOSConfig(ctx, nios, resp)
	}
}

func validateRecordRpzCnameClientipaddressNIOSConfig(ctx context.Context, m *NIOSRecordRpzCnameClientipaddressModel, resp *resource.ValidateConfigResponse) {
	if m.Canonical.IsNull() || m.Canonical.IsUnknown() {
		return
	}

	canonical := m.Canonical.ValueString()
	if canonical == "" || canonical == "*" || canonical == "rpz-passthru" {
		return
	}

	if _, err := netip.ParseAddr(canonical); err != nil {
		resp.Diagnostics.AddAttributeError(
			path.Root("nios").AtName("canonical"),
			"Invalid Canonical Value",
			fmt.Sprintf("The canonical value must be empty, '*', 'rpz-passthru', or a valid IP address. Got: %q", canonical),
		)
	}
}
