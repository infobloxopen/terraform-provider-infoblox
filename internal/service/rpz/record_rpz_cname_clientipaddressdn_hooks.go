package rpz

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
)

// ValidateRecordRpzCnameClientipaddressdn validates the RecordRpzCnameClientipaddressdn configuration.
func ValidateRecordRpzCnameClientipaddressdn(ctx context.Context, data RecordRpzCnameClientipaddressdnModel, resp *resource.ValidateConfigResponse) {
	if nios := flex.ExpandNestedObject[NIOSRecordRpzCnameClientipaddressdnModel](ctx, data.NIOS, &resp.Diagnostics); nios != nil {
		validateRecordRpzCnameClientipaddressdnNIOSConfig(ctx, nios, resp)
	}
}

func validateRecordRpzCnameClientipaddressdnNIOSConfig(ctx context.Context, m *NIOSRecordRpzCnameClientipaddressdnModel, resp *resource.ValidateConfigResponse) {
}
