package rpz

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
)

// ValidateRecordRpzAaaaIpaddress validates the RecordRpzAaaaIpaddress configuration.
func ValidateRecordRpzAaaaIpaddress(ctx context.Context, data RecordRpzAaaaIpaddressModel, resp *resource.ValidateConfigResponse) {
	if nios := flex.ExpandNestedObject[NIOSRecordRpzAaaaIpaddressModel](ctx, data.NIOS, &resp.Diagnostics); nios != nil {
		validateRecordRpzAaaaIpaddressNIOSConfig(ctx, nios, resp)
	}
}

func validateRecordRpzAaaaIpaddressNIOSConfig(ctx context.Context, m *NIOSRecordRpzAaaaIpaddressModel, resp *resource.ValidateConfigResponse) {
}
