package dhcp

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
)

// ValidateHaGroup validates the HaGroup configuration.
func ValidateHaGroup(ctx context.Context, data HaGroupModel, resp *resource.ValidateConfigResponse) {
	if uddi := flex.ExpandNestedObject[UDDIHaGroupModel](ctx, data.UDDI, &resp.Diagnostics); uddi != nil {
		validateHaGroupUDDIConfig(ctx, uddi, resp)
	}
}

func validateHaGroupUDDIConfig(ctx context.Context, m *UDDIHaGroupModel, resp *resource.ValidateConfigResponse) {
}
