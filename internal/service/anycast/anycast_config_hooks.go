package anycast

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
)

// ValidateAnycastConfig validates the AnycastConfig configuration.
func ValidateAnycastConfig(ctx context.Context, data AnycastConfigModel, resp *resource.ValidateConfigResponse) {
	if uddi := flex.ExpandNestedObject[UDDIAnycastConfigModel](ctx, data.UDDI, &resp.Diagnostics); uddi != nil {
		validateAnycastConfigUDDIConfig(ctx, uddi, resp)
	}
}

func validateAnycastConfigUDDIConfig(ctx context.Context, m *UDDIAnycastConfigModel, resp *resource.ValidateConfigResponse) {
}
