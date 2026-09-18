package infra

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
)

// ValidateInfraHost validates the InfraHost configuration.
func ValidateInfraHost(ctx context.Context, data InfraHostModel, resp *resource.ValidateConfigResponse) {
	if uddi := flex.ExpandNestedObject[UDDIInfraHostModel](ctx, data.UDDI, &resp.Diagnostics); uddi != nil {
		validateInfraHostUDDIConfig(ctx, uddi, resp)
	}
}

func validateInfraHostUDDIConfig(ctx context.Context, m *UDDIInfraHostModel, resp *resource.ValidateConfigResponse) {
}
