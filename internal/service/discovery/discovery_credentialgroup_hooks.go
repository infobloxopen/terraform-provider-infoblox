package discovery

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
)

// ValidateDiscoveryCredentialgroup validates the DiscoveryCredentialgroup configuration.
func ValidateDiscoveryCredentialgroup(ctx context.Context, data DiscoveryCredentialgroupModel, resp *resource.ValidateConfigResponse) {
	if nios := flex.ExpandNestedObject[NIOSDiscoveryCredentialgroupModel](ctx, data.NIOS, &resp.Diagnostics); nios != nil {
		validateDiscoveryCredentialgroupNIOSConfig(ctx, nios, resp)
	}
}

func validateDiscoveryCredentialgroupNIOSConfig(ctx context.Context, m *NIOSDiscoveryCredentialgroupModel, resp *resource.ValidateConfigResponse) {
}
