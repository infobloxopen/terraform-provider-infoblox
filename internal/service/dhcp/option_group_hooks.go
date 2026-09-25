package dhcp

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
)

// ValidateOptionGroup validates the OptionGroup configuration.
func ValidateOptionGroup(ctx context.Context, data OptionGroupModel, resp *resource.ValidateConfigResponse) {
	if uddi := flex.ExpandNestedObject[UDDIOptionGroupModel](ctx, data.UDDI, &resp.Diagnostics); uddi != nil {
		validateOptionGroupUDDIConfig(ctx, uddi, resp)
	}
}

func validateOptionGroupUDDIConfig(ctx context.Context, m *UDDIOptionGroupModel, resp *resource.ValidateConfigResponse) {
}
