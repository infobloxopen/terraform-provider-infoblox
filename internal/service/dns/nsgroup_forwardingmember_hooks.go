package dns

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
)

// ValidateNsgroupForwardingmember validates the NsgroupForwardingmember configuration.
func ValidateNsgroupForwardingmember(ctx context.Context, data NsgroupForwardingmemberModel, resp *resource.ValidateConfigResponse) {
	if nios := flex.ExpandNestedObject[NIOSNsgroupForwardingmemberModel](ctx, data.NIOS, &resp.Diagnostics); nios != nil {
		validateNsgroupForwardingmemberNIOSConfig(ctx, nios, resp)
	}
}

func validateNsgroupForwardingmemberNIOSConfig(ctx context.Context, m *NIOSNsgroupForwardingmemberModel, resp *resource.ValidateConfigResponse) {
}
