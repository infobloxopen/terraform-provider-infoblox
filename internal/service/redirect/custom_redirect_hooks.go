package redirect

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
)

// ValidateCustomRedirect validates the CustomRedirect configuration.
func ValidateCustomRedirect(ctx context.Context, data CustomRedirectModel, resp *resource.ValidateConfigResponse) {
	if uddi := flex.ExpandNestedObject[UDDICustomRedirectModel](ctx, data.UDDI, &resp.Diagnostics); uddi != nil {
		validateCustomRedirectUDDIConfig(ctx, uddi, resp)
	}
}

func validateCustomRedirectUDDIConfig(ctx context.Context, m *UDDICustomRedirectModel, resp *resource.ValidateConfigResponse) {
}
