package dns

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
)

// ValidateSharedrecordgroup validates the Sharedrecordgroup configuration.
func ValidateSharedrecordgroup(ctx context.Context, data SharedrecordgroupModel, resp *resource.ValidateConfigResponse) {
	if nios := flex.ExpandNestedObject[NIOSSharedrecordgroupModel](ctx, data.NIOS, &resp.Diagnostics); nios != nil {
		validateSharedrecordgroupNIOSConfig(ctx, nios, resp)
	}
}

func validateSharedrecordgroupNIOSConfig(ctx context.Context, m *NIOSSharedrecordgroupModel, resp *resource.ValidateConfigResponse) {
}
