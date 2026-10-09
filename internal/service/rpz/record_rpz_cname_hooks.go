package rpz

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
)

// ValidateRecordRpzCname validates the RecordRpzCname configuration.
func ValidateRecordRpzCname(ctx context.Context, data RecordRpzCnameModel, resp *resource.ValidateConfigResponse) {
	if nios := flex.ExpandNestedObject[NIOSRecordRpzCnameModel](ctx, data.NIOS, &resp.Diagnostics); nios != nil {
		validateRecordRpzCnameNIOSConfig(ctx, nios, resp)
	}
}

func validateRecordRpzCnameNIOSConfig(ctx context.Context, m *NIOSRecordRpzCnameModel, resp *resource.ValidateConfigResponse) {
}
