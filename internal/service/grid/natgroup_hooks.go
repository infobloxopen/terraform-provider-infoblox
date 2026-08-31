package grid

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
)

// ValidateNatgroup validates the Natgroup configuration.
func ValidateNatgroup(ctx context.Context, data NatgroupModel, resp *resource.ValidateConfigResponse) {
	if nios := flex.ExpandNestedObject[NIOSNatgroupModel](ctx, data.NIOS, &resp.Diagnostics); nios != nil {
		validateNatgroupNIOSConfig(ctx, nios, resp)
	}
}

func validateNatgroupNIOSConfig(ctx context.Context, m *NIOSNatgroupModel, resp *resource.ValidateConfigResponse) {
}
