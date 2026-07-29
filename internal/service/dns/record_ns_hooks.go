package dns

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
)

// ValidateRecordNs validates the RecordNs configuration.
func ValidateRecordNs(ctx context.Context, data RecordNsModel, resp *resource.ValidateConfigResponse) {
	if nios := flex.ExpandNestedObject[NIOSRecordNsModel](ctx, data.NIOS, &resp.Diagnostics); nios != nil {
		validateRecordNsNIOSConfig(ctx, nios, resp)
	}
	if uddi := flex.ExpandNestedObject[UDDIRecordNsModel](ctx, data.UDDI, &resp.Diagnostics); uddi != nil {
		validateRecordNsUDDIConfig(ctx, uddi, resp)
	}
}

func validateRecordNsNIOSConfig(ctx context.Context, m *NIOSRecordNsModel, resp *resource.ValidateConfigResponse) {
}

func validateRecordNsUDDIConfig(ctx context.Context, m *UDDIRecordNsModel, resp *resource.ValidateConfigResponse) {
}
