package rpz

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
)

// ValidateRecordRpzNaptr validates the RecordRpzNaptr configuration.
func ValidateRecordRpzNaptr(ctx context.Context, data RecordRpzNaptrModel, resp *resource.ValidateConfigResponse) {
	if nios := flex.ExpandNestedObject[NIOSRecordRpzNaptrModel](ctx, data.NIOS, &resp.Diagnostics); nios != nil {
		validateRecordRpzNaptrNIOSConfig(ctx, nios, resp)
	}
}

func validateRecordRpzNaptrNIOSConfig(ctx context.Context, m *NIOSRecordRpzNaptrModel, resp *resource.ValidateConfigResponse) {
}
