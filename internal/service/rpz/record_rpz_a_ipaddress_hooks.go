package rpz

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
)

// ValidateRecordRpzAIpaddress validates the RecordRpzAIpaddress configuration.
func ValidateRecordRpzAIpaddress(ctx context.Context, data RecordRpzAIpaddressModel, resp *resource.ValidateConfigResponse) {
	if nios := flex.ExpandNestedObject[NIOSRecordRpzAIpaddressModel](ctx, data.NIOS, &resp.Diagnostics); nios != nil {
		validateRecordRpzAIpaddressNIOSConfig(ctx, nios, resp)
	}
}

func validateRecordRpzAIpaddressNIOSConfig(ctx context.Context, m *NIOSRecordRpzAIpaddressModel, resp *resource.ValidateConfigResponse) {
}
