package dns

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
)

// ValidateAuthNsg validates the AuthNsg configuration.
func ValidateAuthNsg(ctx context.Context, data AuthNsgModel, resp *resource.ValidateConfigResponse) {
	if uddi := flex.ExpandNestedObject[UDDIAuthNsgModel](ctx, data.UDDI, &resp.Diagnostics); uddi != nil {
		validateAuthNsgUDDIConfig(ctx, uddi, resp)
	}
}

func validateAuthNsgUDDIConfig(ctx context.Context, m *UDDIAuthNsgModel, resp *resource.ValidateConfigResponse) {
}
