package ipam

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
)

// ValidateBulkhostnametemplate validates the Bulkhostnametemplate configuration.
func ValidateBulkhostnametemplate(ctx context.Context, data BulkhostnametemplateModel, resp *resource.ValidateConfigResponse) {
	if nios := flex.ExpandNestedObject[NIOSBulkhostnametemplateModel](ctx, data.NIOS, &resp.Diagnostics); nios != nil {
		validateBulkhostnametemplateNIOSConfig(ctx, nios, resp)
	}
}

func validateBulkhostnametemplateNIOSConfig(ctx context.Context, m *NIOSBulkhostnametemplateModel, resp *resource.ValidateConfigResponse) {
}
