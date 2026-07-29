package dns

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
)

// ValidateRecordCname validates the RecordCname configuration.
func ValidateRecordCname(ctx context.Context, data RecordCnameModel, resp *resource.ValidateConfigResponse) {
	if nios := flex.ExpandNestedObject[NIOSRecordCnameModel](ctx, data.NIOS, &resp.Diagnostics); nios != nil {
		validateRecordCnameNIOSConfig(ctx, nios, resp)
	}
	if uddi := flex.ExpandNestedObject[UDDIRecordCnameModel](ctx, data.UDDI, &resp.Diagnostics); uddi != nil {
		validateRecordCnameUDDIConfig(ctx, uddi, resp)
	}
}

func validateRecordCnameNIOSConfig(ctx context.Context, m *NIOSRecordCnameModel, resp *resource.ValidateConfigResponse) {
}

func validateRecordCnameUDDIConfig(ctx context.Context, m *UDDIRecordCnameModel, resp *resource.ValidateConfigResponse) {
}
