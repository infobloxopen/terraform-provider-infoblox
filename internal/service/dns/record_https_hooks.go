package dns

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
)

// ValidateRecordHttps validates the RecordHttps configuration.
func ValidateRecordHttps(ctx context.Context, data RecordHttpsModel, resp *resource.ValidateConfigResponse) {
	if uddi := flex.ExpandNestedObject[UDDIRecordHttpsModel](ctx, data.UDDI, &resp.Diagnostics); uddi != nil {
		validateRecordHttpsUDDIConfig(ctx, uddi, resp)
	}
}

func validateRecordHttpsUDDIConfig(ctx context.Context, m *UDDIRecordHttpsModel, resp *resource.ValidateConfigResponse) {
}
