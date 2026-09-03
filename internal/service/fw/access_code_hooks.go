package fw

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
)

// ValidateAccessCode validates the AccessCode configuration.
func ValidateAccessCode(ctx context.Context, data AccessCodeModel, resp *resource.ValidateConfigResponse) {
	if uddi := flex.ExpandNestedObject[UDDIAccessCodeModel](ctx, data.UDDI, &resp.Diagnostics); uddi != nil {
		validateAccessCodeUDDIConfig(ctx, uddi, resp)
	}
}

func validateAccessCodeUDDIConfig(ctx context.Context, m *UDDIAccessCodeModel, resp *resource.ValidateConfigResponse) {
}
