package rpz

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
)

// ValidateRecordRpzPtr validates the RecordRpzPtr configuration.
func ValidateRecordRpzPtr(ctx context.Context, data RecordRpzPtrModel, resp *resource.ValidateConfigResponse) {
	if nios := flex.ExpandNestedObject[NIOSRecordRpzPtrModel](ctx, data.NIOS, &resp.Diagnostics); nios != nil {
		validateRecordRpzPtrNIOSConfig(ctx, nios, resp)
	}
}

func validateRecordRpzPtrNIOSConfig(ctx context.Context, m *NIOSRecordRpzPtrModel, resp *resource.ValidateConfigResponse) {
}
