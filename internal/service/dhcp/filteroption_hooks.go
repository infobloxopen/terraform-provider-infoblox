package dhcp

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/utils"
)

// ValidateFilteroption validates the Filteroption configuration.
func ValidateFilteroption(ctx context.Context, data FilteroptionModel, resp *resource.ValidateConfigResponse) {
	if nios := flex.ExpandNestedObject[NIOSFilteroptionModel](ctx, data.NIOS, &resp.Diagnostics); nios != nil {
		validateFilteroptionNIOSConfig(ctx, nios, resp)
	}
	if uddi := flex.ExpandNestedObject[UDDIFilteroptionModel](ctx, data.UDDI, &resp.Diagnostics); uddi != nil {
		validateFilteroptionUDDIConfig(ctx, uddi, resp)
	}
}

func validateFilteroptionNIOSConfig(ctx context.Context, m *NIOSFilteroptionModel, resp *resource.ValidateConfigResponse) {
}

func validateFilteroptionUDDIConfig(ctx context.Context, m *UDDIFilteroptionModel, resp *resource.ValidateConfigResponse) {
}

func PostFlattenFilteroptionNIOS(ctx context.Context, planned, flattened *NIOSFilteroptionModel, diags *diag.Diagnostics) {
	if planned == nil || flattened == nil {
		return
	}

	if !planned.OptionList.IsUnknown() {
		if reordered, d := utils.ReorderAndFilterDHCPOptions(ctx, planned.OptionList, flattened.OptionList); !d.HasError() {
			flattened.OptionList = reordered.(basetypes.ListValue)
		}
	}
}
