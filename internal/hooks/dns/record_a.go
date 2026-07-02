package dns

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"

	niosdns "github.com/infobloxopen/infoblox-nios-go-client/dns"
	"github.com/infobloxopen/terraform-provider-unified/internal/dynamicallocation"
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

func validateRecordANIOSConfig(_ctx context.Context, _data types.Object, _resp *resource.ValidateConfigResponse) {
}

func validateRecordAUDDIConfig(_ctx context.Context, _data types.Object, _resp *resource.ValidateConfigResponse) {
}

func BuildRecordAFuncCall(ctx context.Context, allocObj types.Object, diags *diag.Diagnostics) *niosdns.FuncCall {
	if allocObj.IsNull() || allocObj.IsUnknown() {
		return nil
	}

	var m dynamicallocation.NextAvailableIpModel
	diags.Append(allocObj.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}

	return m.FuncCall(ctx, "Ipv4addr", "network", diags)
}
