package rpz

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
)

// ValidateRecordRpzA validates the RecordRpzA configuration.
func ValidateRecordRpzA(ctx context.Context, data RecordRpzAModel, resp *resource.ValidateConfigResponse) {
	if nios := flex.ExpandNestedObject[NIOSRecordRpzAModel](ctx, data.NIOS, &resp.Diagnostics); nios != nil {
		validateRecordRpzANIOSConfig(ctx, nios, resp)
	}
}

func validateRecordRpzANIOSConfig(ctx context.Context, m *NIOSRecordRpzAModel, resp *resource.ValidateConfigResponse) {
}
