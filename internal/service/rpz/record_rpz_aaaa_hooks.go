package rpz

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
)

// ValidateRecordRpzAaaa validates the RecordRpzAaaa configuration.
func ValidateRecordRpzAaaa(ctx context.Context, data RecordRpzAaaaModel, resp *resource.ValidateConfigResponse) {
	if nios := flex.ExpandNestedObject[NIOSRecordRpzAaaaModel](ctx, data.NIOS, &resp.Diagnostics); nios != nil {
		validateRecordRpzAaaaNIOSConfig(ctx, nios, resp)
	}
}

func validateRecordRpzAaaaNIOSConfig(ctx context.Context, m *NIOSRecordRpzAaaaModel, resp *resource.ValidateConfigResponse) {
}
