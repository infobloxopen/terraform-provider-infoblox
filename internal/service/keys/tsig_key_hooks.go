package keys

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
)

// ValidateTsigKey validates the TsigKey configuration.
func ValidateTsigKey(ctx context.Context, data TsigKeyModel, resp *resource.ValidateConfigResponse) {
	if uddi := flex.ExpandNestedObject[UDDITsigKeyModel](ctx, data.UDDI, &resp.Diagnostics); uddi != nil {
		validateTsigKeyUDDIConfig(ctx, uddi, resp)
	}
}

func validateTsigKeyUDDIConfig(ctx context.Context, m *UDDITsigKeyModel, resp *resource.ValidateConfigResponse) {
}
