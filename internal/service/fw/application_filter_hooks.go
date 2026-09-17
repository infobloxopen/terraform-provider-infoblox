package fw

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
)

// ValidateApplicationFilter validates the ApplicationFilter configuration.
func ValidateApplicationFilter(ctx context.Context, data ApplicationFilterModel, resp *resource.ValidateConfigResponse) {
	if uddi := flex.ExpandNestedObject[UDDIApplicationFilterModel](ctx, data.UDDI, &resp.Diagnostics); uddi != nil {
		validateApplicationFilterUDDIConfig(ctx, uddi, resp)
	}
}

func validateApplicationFilterUDDIConfig(ctx context.Context, m *UDDIApplicationFilterModel, resp *resource.ValidateConfigResponse) {
}
