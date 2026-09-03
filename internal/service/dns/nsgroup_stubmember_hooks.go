package dns

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
)

// ValidateNsgroupStubmember validates the NsgroupStubmember configuration.
func ValidateNsgroupStubmember(ctx context.Context, data NsgroupStubmemberModel, resp *resource.ValidateConfigResponse) {
	if nios := flex.ExpandNestedObject[NIOSNsgroupStubmemberModel](ctx, data.NIOS, &resp.Diagnostics); nios != nil {
		validateNsgroupStubmemberNIOSConfig(ctx, nios, resp)
	}
}

func validateNsgroupStubmemberNIOSConfig(ctx context.Context, m *NIOSNsgroupStubmemberModel, resp *resource.ValidateConfigResponse) {
}
