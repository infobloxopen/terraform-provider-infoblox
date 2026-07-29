package dns

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
)

// ValidateRecordCaa validates the RecordCaa configuration.
func ValidateRecordCaa(ctx context.Context, data RecordCaaModel, resp *resource.ValidateConfigResponse) {
	if nios := flex.ExpandNestedObject[NIOSRecordCaaModel](ctx, data.NIOS, &resp.Diagnostics); nios != nil {
		validateRecordCaaNIOSConfig(ctx, nios, resp)
	}
	if uddi := flex.ExpandNestedObject[UDDIRecordCaaModel](ctx, data.UDDI, &resp.Diagnostics); uddi != nil {
		validateRecordCaaUDDIConfig(ctx, uddi, resp)
	}
}

func validateRecordCaaNIOSConfig(ctx context.Context, m *NIOSRecordCaaModel, resp *resource.ValidateConfigResponse) {
}

func validateRecordCaaUDDIConfig(ctx context.Context, m *UDDIRecordCaaModel, resp *resource.ValidateConfigResponse) {
}
