package ipamfederation

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
)

// ValidateForwardLookingDelegation validates the ForwardLookingDelegation configuration.
func ValidateForwardLookingDelegation(ctx context.Context, data ForwardLookingDelegationModel, resp *resource.ValidateConfigResponse) {
	if uddi := flex.ExpandNestedObject[UDDIForwardLookingDelegationModel](ctx, data.UDDI, &resp.Diagnostics); uddi != nil {
		validateForwardLookingDelegationUDDIConfig(ctx, uddi, resp)
	}
}

func validateForwardLookingDelegationUDDIConfig(ctx context.Context, m *UDDIForwardLookingDelegationModel, resp *resource.ValidateConfigResponse) {
}
