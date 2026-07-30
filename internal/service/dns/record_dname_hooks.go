package dns

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
)

// ValidateRecordDname validates the RecordDname configuration.
func ValidateRecordDname(ctx context.Context, data RecordDnameModel, resp *resource.ValidateConfigResponse) {
	if nios := flex.ExpandNestedObject[NIOSRecordDnameModel](ctx, data.NIOS, &resp.Diagnostics); nios != nil {
		validateRecordDnameNIOSConfig(ctx, nios, resp)
	}
	if uddi := flex.ExpandNestedObject[UDDIRecordDnameModel](ctx, data.UDDI, &resp.Diagnostics); uddi != nil {
		validateRecordDnameUDDIConfig(ctx, uddi, resp)
	}
}

func validateRecordDnameNIOSConfig(ctx context.Context, m *NIOSRecordDnameModel, resp *resource.ValidateConfigResponse) {
}

func validateRecordDnameUDDIConfig(ctx context.Context, m *UDDIRecordDnameModel, resp *resource.ValidateConfigResponse) {
}
