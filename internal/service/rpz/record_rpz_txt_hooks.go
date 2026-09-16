package rpz

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
)

// ValidateRecordRpzTxt validates the RecordRpzTxt configuration.
func ValidateRecordRpzTxt(ctx context.Context, data RecordRpzTxtModel, resp *resource.ValidateConfigResponse) {
	if nios := flex.ExpandNestedObject[NIOSRecordRpzTxtModel](ctx, data.NIOS, &resp.Diagnostics); nios != nil {
		validateRecordRpzTxtNIOSConfig(ctx, nios, resp)
	}
}

func validateRecordRpzTxtNIOSConfig(ctx context.Context, m *NIOSRecordRpzTxtModel, resp *resource.ValidateConfigResponse) {
}
