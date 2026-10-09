package cloud

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
)

// ValidateAwsuser validates the Awsuser configuration.
func ValidateAwsuser(ctx context.Context, data AwsuserModel, resp *resource.ValidateConfigResponse) {
	if nios := flex.ExpandNestedObject[NIOSAwsuserModel](ctx, data.NIOS, &resp.Diagnostics); nios != nil {
		validateAwsuserNIOSConfig(ctx, nios, resp)
	}
}

func validateAwsuserNIOSConfig(ctx context.Context, m *NIOSAwsuserModel, resp *resource.ValidateConfigResponse) {
}
