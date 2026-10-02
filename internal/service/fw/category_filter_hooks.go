package fw

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
)

// ValidateCategoryFilter validates the CategoryFilter configuration.
func ValidateCategoryFilter(ctx context.Context, data CategoryFilterModel, resp *resource.ValidateConfigResponse) {
	if uddi := flex.ExpandNestedObject[UDDICategoryFilterModel](ctx, data.UDDI, &resp.Diagnostics); uddi != nil {
		validateCategoryFilterUDDIConfig(ctx, uddi, resp)
	}
}

func validateCategoryFilterUDDIConfig(ctx context.Context, m *UDDICategoryFilterModel, resp *resource.ValidateConfigResponse) {
}
