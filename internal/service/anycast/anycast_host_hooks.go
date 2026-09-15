package anycast

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
)

// ValidateAnycastHost validates the AnycastHost configuration.
func ValidateAnycastHost(ctx context.Context, data AnycastHostModel, resp *resource.ValidateConfigResponse) {
	if uddi := flex.ExpandNestedObject[UDDIAnycastHostModel](ctx, data.UDDI, &resp.Diagnostics); uddi != nil {
		validateAnycastHostUDDIConfig(ctx, uddi, resp)
	}
}

func validateAnycastHostUDDIConfig(ctx context.Context, m *UDDIAnycastHostModel, resp *resource.ValidateConfigResponse) {
}
