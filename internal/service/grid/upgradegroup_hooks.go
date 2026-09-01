package grid

import (
	"context"
	"fmt"

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
	if !m.Members.IsNull() && !m.Members.IsUnknown() {
		var members []UpgradegroupMembersModel
		diags := m.Members.ElementsAs(ctx, &members, false)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}
		for i, member := range members {
			if member.Member.IsNull() || member.Member.IsUnknown() {
				resp.Diagnostics.AddError(
					"Validation Error",
					fmt.Sprintf("members.%d.member must be provided", i),
				)
			}
		}
	}
}
