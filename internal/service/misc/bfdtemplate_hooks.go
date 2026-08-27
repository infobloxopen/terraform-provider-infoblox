package misc

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
)

// ValidateBfdtemplate validates the Bfdtemplate configuration.
func ValidateBfdtemplate(ctx context.Context, data BfdtemplateModel, resp *resource.ValidateConfigResponse) {
	if nios := flex.ExpandNestedObject[NIOSBfdtemplateModel](ctx, data.NIOS, &resp.Diagnostics); nios != nil {
		validateBfdtemplateNIOSConfig(ctx, nios, resp)
	}
}

func validateBfdtemplateNIOSConfig(ctx context.Context, m *NIOSBfdtemplateModel, resp *resource.ValidateConfigResponse) {
}
