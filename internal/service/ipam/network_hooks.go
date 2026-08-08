package ipam

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	niosipam "github.com/infobloxopen/infoblox-nios-go-client/ipam"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/dynamicallocation"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/utils"
)

// ValidateNetwork validates the Network configuration.
func ValidateNetwork(ctx context.Context, data NetworkModel, resp *resource.ValidateConfigResponse) {
	if nios := flex.ExpandNestedObject[NIOSNetworkModel](ctx, data.NIOS, &resp.Diagnostics); nios != nil {
		validateNetworkNIOSConfig(ctx, nios, resp)
	}
	if uddi := flex.ExpandNestedObject[UDDINetworkModel](ctx, data.UDDI, &resp.Diagnostics); uddi != nil {
		validateNetworkUDDIConfig(ctx, uddi, resp)
	}
}

func validateNetworkNIOSConfig(ctx context.Context, m *NIOSNetworkModel, resp *resource.ValidateConfigResponse) {
}

func validateNetworkUDDIConfig(ctx context.Context, m *UDDINetworkModel, resp *resource.ValidateConfigResponse) {
}

func BuildNetworkFuncCall(ctx context.Context, data types.Object, diags *diag.Diagnostics) *niosipam.FuncCall {
	if data.IsNull() || data.IsUnknown() {
		return nil
	}

	var m dynamicallocation.NextAvailableNetworkModel
	diags.Append(data.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}

	return m.FuncCall(ctx, "Network", "network", diags)
}

func PostFlattenNetworkNIOS(ctx context.Context, planned, flattened *NIOSNetworkModel, diags *diag.Diagnostics) {
	if planned != nil && !planned.Options.IsUnknown() {
		if reordered, d := utils.ReorderAndFilterDHCPOptions(ctx, planned.Options, flattened.Options); !d.HasError() {
			flattened.Options = reordered.(basetypes.ListValue)
		}
	}
}
