package infra

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
)

// ValidateInfraService validates the InfraService configuration.
func ValidateInfraService(ctx context.Context, data InfraServiceModel, resp *resource.ValidateConfigResponse) {
	if uddi := flex.ExpandNestedObject[UDDIInfraServiceModel](ctx, data.UDDI, &resp.Diagnostics); uddi != nil {
		validateInfraServiceUDDIConfig(ctx, uddi, resp)
	}
}

func validateInfraServiceUDDIConfig(ctx context.Context, m *UDDIInfraServiceModel, resp *resource.ValidateConfigResponse) {
}
