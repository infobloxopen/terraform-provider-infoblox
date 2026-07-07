package dns

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	niosdns "github.com/infobloxopen/infoblox-nios-go-client/dns"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/dynamicallocation"
)

// ValidateRecordA validates the RecordA configuration.
func ValidateRecordA(ctx context.Context, niosBlock, uddiBlock types.Object, resp *resource.ValidateConfigResponse) {
	if !niosBlock.IsNull() && !niosBlock.IsUnknown() {
		validateRecordANIOSConfig(ctx, niosBlock, resp)
	}
	if !uddiBlock.IsNull() && !uddiBlock.IsUnknown() {
		validateRecordAUDDIConfig(ctx, uddiBlock, resp)
	}
}

func validateRecordANIOSConfig(ctx context.Context, data types.Object, resp *resource.ValidateConfigResponse) {
}

func validateRecordAUDDIConfig(ctx context.Context, data types.Object, resp *resource.ValidateConfigResponse) {
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
