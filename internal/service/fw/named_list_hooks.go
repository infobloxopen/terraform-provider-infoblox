package fw

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
)

// ValidateNamedList validates the NamedList configuration.
func ValidateNamedList(ctx context.Context, data NamedListModel, resp *resource.ValidateConfigResponse) {
	if uddi := flex.ExpandNestedObject[UDDINamedListModel](ctx, data.UDDI, &resp.Diagnostics); uddi != nil {
		validateNamedListUDDIConfig(ctx, uddi, resp)
	}
}

func validateNamedListUDDIConfig(ctx context.Context, m *UDDINamedListModel, resp *resource.ValidateConfigResponse) {
}
