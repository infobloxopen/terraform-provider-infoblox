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

// ValidateRecordAaaa validates the RecordAaaa configuration.
func ValidateRecordAaaa(ctx context.Context, data RecordAaaaModel, resp *resource.ValidateConfigResponse) {
	if nios := flex.ExpandNestedObject[NIOSRecordAaaaModel](ctx, data.NIOS, &resp.Diagnostics); nios != nil {
		validateRecordAaaaNIOSConfig(ctx, nios, resp)
	}
	if uddi := flex.ExpandNestedObject[UDDIRecordAaaaModel](ctx, data.UDDI, &resp.Diagnostics); uddi != nil {
		validateRecordAaaaUDDIConfig(ctx, uddi, resp)
	}
}

func validateRecordAaaaNIOSConfig(ctx context.Context, m *NIOSRecordAaaaModel, resp *resource.ValidateConfigResponse) {
}

func validateRecordAaaaUDDIConfig(ctx context.Context, m *UDDIRecordAaaaModel, resp *resource.ValidateConfigResponse) {
}

func BuildRecordAaaaFuncCall(ctx context.Context, data types.Object, diags *diag.Diagnostics) *niosdns.FuncCall {
	if data.IsNull() || data.IsUnknown() {
		return nil
	}

	var m dynamicallocation.NextAvailableIpModel
	diags.Append(data.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}

	return m.FuncCall(ctx, "Ipv6addr", "ipv6network", diags)
}
