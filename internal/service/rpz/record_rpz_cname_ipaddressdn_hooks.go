package rpz

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
)

// ValidateRecordRpzCnameIpaddressdn validates the RecordRpzCnameIpaddressdn configuration.
func ValidateRecordRpzCnameIpaddressdn(ctx context.Context, data RecordRpzCnameIpaddressdnModel, resp *resource.ValidateConfigResponse) {
	if nios := flex.ExpandNestedObject[NIOSRecordRpzCnameIpaddressdnModel](ctx, data.NIOS, &resp.Diagnostics); nios != nil {
		validateRecordRpzCnameIpaddressdnNIOSConfig(ctx, nios, resp)
	}
}

func validateRecordRpzCnameIpaddressdnNIOSConfig(ctx context.Context, m *NIOSRecordRpzCnameIpaddressdnModel, resp *resource.ValidateConfigResponse) {
}
