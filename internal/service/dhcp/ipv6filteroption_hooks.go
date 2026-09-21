package dhcp

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/utils"
)

// ValidateIpv6filteroption validates the Ipv6filteroption configuration.
func ValidateIpv6filteroption(ctx context.Context, data Ipv6filteroptionModel, resp *resource.ValidateConfigResponse) {
	if nios := flex.ExpandNestedObject[NIOSIpv6filteroptionModel](ctx, data.NIOS, &resp.Diagnostics); nios != nil {
		validateIpv6filteroptionNIOSConfig(ctx, nios, resp)
	}
	if uddi := flex.ExpandNestedObject[UDDIIpv6filteroptionModel](ctx, data.UDDI, &resp.Diagnostics); uddi != nil {
		validateIpv6filteroptionUDDIConfig(ctx, uddi, resp)
	}
}

func validateIpv6filteroptionNIOSConfig(ctx context.Context, m *NIOSIpv6filteroptionModel, resp *resource.ValidateConfigResponse) {
}

func validateIpv6filteroptionUDDIConfig(ctx context.Context, m *UDDIIpv6filteroptionModel, resp *resource.ValidateConfigResponse) {
}

func PostFlattenIpv6filteroptionNIOS(ctx context.Context, planned, flattened *NIOSIpv6filteroptionModel, diags *diag.Diagnostics) {
	if planned == nil || flattened == nil {
		return
	}

	if !planned.OptionList.IsUnknown() {
		if reordered, d := utils.ReorderAndFilterDHCPOptions(ctx, planned.OptionList, flattened.OptionList); !d.HasError() {
			flattened.OptionList = reordered.(basetypes.ListValue)
		}
	}
}
