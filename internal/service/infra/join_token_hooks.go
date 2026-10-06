package infra

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
)

// ValidateJoinToken validates the JoinToken configuration.
func ValidateJoinToken(ctx context.Context, data JoinTokenModel, resp *resource.ValidateConfigResponse) {
	if uddi := flex.ExpandNestedObject[UDDIJoinTokenModel](ctx, data.UDDI, &resp.Diagnostics); uddi != nil {
		validateJoinTokenUDDIConfig(ctx, uddi, resp)
	}
}

func validateJoinTokenUDDIConfig(ctx context.Context, m *UDDIJoinTokenModel, resp *resource.ValidateConfigResponse) {
}

// PostFlattenJoinTokenUDDI restores the create-response fields onto the flattened
// model. join_token is returned once, on CreateJoinTokenResponse, and never
// again -- an unguarded flatten on read would null it out and lose the secret.
// The state value is authoritative once created.
func PostFlattenJoinTokenUDDI(ctx context.Context, planned, flattened *UDDIJoinTokenModel, diags *diag.Diagnostics) {
	if planned == nil || flattened == nil {
		return
	}
	if !planned.JoinToken.IsUnknown() && !planned.JoinToken.IsNull() {
		flattened.JoinToken = planned.JoinToken
	}
}
