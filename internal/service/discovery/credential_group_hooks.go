package discovery

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
)

// ValidateCredentialGroup validates the CredentialGroup configuration.
func ValidateCredentialGroup(ctx context.Context, data CredentialGroupModel, resp *resource.ValidateConfigResponse) {
	if nios := flex.ExpandNestedObject[NIOSCredentialGroupModel](ctx, data.NIOS, &resp.Diagnostics); nios != nil {
		validateCredentialGroupNIOSConfig(ctx, nios, resp)
	}
}

func validateCredentialGroupNIOSConfig(ctx context.Context, m *NIOSCredentialGroupModel, resp *resource.ValidateConfigResponse) {
}
