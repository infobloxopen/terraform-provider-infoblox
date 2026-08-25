package grid

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
)

// ValidateUpgradegroup validates the Upgradegroup configuration.
func ValidateUpgradegroup(ctx context.Context, data UpgradegroupModel, resp *resource.ValidateConfigResponse) {
	if nios := flex.ExpandNestedObject[NIOSUpgradegroupModel](ctx, data.NIOS, &resp.Diagnostics); nios != nil {
		validateUpgradegroupNIOSConfig(ctx, nios, resp)
	}
}

func validateUpgradegroupNIOSConfig(ctx context.Context, m *NIOSUpgradegroupModel, resp *resource.ValidateConfigResponse) {
}
