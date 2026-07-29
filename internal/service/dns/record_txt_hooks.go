package dns

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
)

// ValidateRecordTxt validates the RecordTxt configuration.
func ValidateRecordTxt(ctx context.Context, data RecordTxtModel, resp *resource.ValidateConfigResponse) {
	if nios := flex.ExpandNestedObject[NIOSRecordTxtModel](ctx, data.NIOS, &resp.Diagnostics); nios != nil {
		validateRecordTxtNIOSConfig(ctx, nios, resp)
	}
	if uddi := flex.ExpandNestedObject[UDDIRecordTxtModel](ctx, data.UDDI, &resp.Diagnostics); uddi != nil {
		validateRecordTxtUDDIConfig(ctx, uddi, resp)
	}
}

func validateRecordTxtNIOSConfig(ctx context.Context, m *NIOSRecordTxtModel, resp *resource.ValidateConfigResponse) {
}

func validateRecordTxtUDDIConfig(ctx context.Context, m *UDDIRecordTxtModel, resp *resource.ValidateConfigResponse) {
}
