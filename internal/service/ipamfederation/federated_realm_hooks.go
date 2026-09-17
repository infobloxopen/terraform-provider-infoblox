package ipamfederation

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
)

// ValidateFederatedRealm validates the FederatedRealm configuration.
func ValidateFederatedRealm(ctx context.Context, data FederatedRealmModel, resp *resource.ValidateConfigResponse) {
	if uddi := flex.ExpandNestedObject[UDDIFederatedRealmModel](ctx, data.UDDI, &resp.Diagnostics); uddi != nil {
		validateFederatedRealmUDDIConfig(ctx, uddi, resp)
	}
}

func validateFederatedRealmUDDIConfig(ctx context.Context, m *UDDIFederatedRealmModel, resp *resource.ValidateConfigResponse) {
}
