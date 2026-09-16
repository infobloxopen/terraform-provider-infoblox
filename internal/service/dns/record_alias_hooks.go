package dns

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
)

// ValidateRecordAlias validates the RecordAlias configuration.
func ValidateRecordAlias(ctx context.Context, data RecordAliasModel, resp *resource.ValidateConfigResponse) {
	if nios := flex.ExpandNestedObject[NIOSRecordAliasModel](ctx, data.NIOS, &resp.Diagnostics); nios != nil {
		validateRecordAliasNIOSConfig(ctx, nios, resp)
	}
}

func validateRecordAliasNIOSConfig(ctx context.Context, m *NIOSRecordAliasModel, resp *resource.ValidateConfigResponse) {
}
