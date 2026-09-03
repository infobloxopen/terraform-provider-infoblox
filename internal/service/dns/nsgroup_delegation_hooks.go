package dns

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
)

// ValidateNsgroupDelegation validates the NsgroupDelegation configuration.
func ValidateNsgroupDelegation(ctx context.Context, data NsgroupDelegationModel, resp *resource.ValidateConfigResponse) {
	if nios := flex.ExpandNestedObject[NIOSNsgroupDelegationModel](ctx, data.NIOS, &resp.Diagnostics); nios != nil {
		validateNsgroupDelegationNIOSConfig(ctx, nios, resp)
	}
}

func validateNsgroupDelegationNIOSConfig(ctx context.Context, m *NIOSNsgroupDelegationModel, resp *resource.ValidateConfigResponse) {
}

func PostFlattenNsgroupDelegationNIOS(ctx context.Context, planned, flattened *NIOSNsgroupDelegationModel, diags *diag.Diagnostics) {
}
