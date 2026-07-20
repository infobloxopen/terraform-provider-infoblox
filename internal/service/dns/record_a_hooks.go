package dns

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	niosdns "github.com/infobloxopen/infoblox-nios-go-client/dns"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/dynamicallocation"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
)

// ValidateRecordA validates the RecordA configuration. Hooks live in the service
// package so they can work with the typed TF model directly (no import cycle).
func ValidateRecordA(ctx context.Context, data RecordAModel, resp *resource.ValidateConfigResponse) {
	if nios := flex.ExpandNestedObject[NIOSRecordAModel](ctx, data.NIOS, &resp.Diagnostics); nios != nil {
		validateRecordANIOSConfig(ctx, nios, resp)
	}
	if uddi := flex.ExpandNestedObject[UDDIRecordAModel](ctx, data.UDDI, &resp.Diagnostics); uddi != nil {
		validateRecordAUDDIConfig(ctx, uddi, resp)
	}
}

func validateRecordANIOSConfig(ctx context.Context, m *NIOSRecordAModel, resp *resource.ValidateConfigResponse) {
}

func validateRecordAUDDIConfig(ctx context.Context, m *UDDIRecordAModel, resp *resource.ValidateConfigResponse) {
}

func BuildRecordAFuncCall(ctx context.Context, data types.Object, diags *diag.Diagnostics) *niosdns.FuncCall {
	if data.IsNull() || data.IsUnknown() {
		return nil
	}

	var m dynamicallocation.NextAvailableIpModel
	diags.Append(data.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}

	return m.FuncCall(ctx, "Ipv4addr", "network", diags)
}
