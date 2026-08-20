package dhcp

import (
	"context"
	"strconv"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
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
	niosPath := path.Root("nios")

	utils.ValidateDHCPOptionsConfig(ctx, m.OptionList, niosPath.AtName("option_list"), &resp.Diagnostics)

	if m.LeaseTime.IsNull() || m.LeaseTime.IsUnknown() || m.OptionList.IsNull() || m.OptionList.IsUnknown() {
		return
	}
	var options []FilteroptionOptionListModel
	resp.Diagnostics.Append(m.OptionList.ElementsAs(ctx, &options, false)...)
	if resp.Diagnostics.HasError() {
		return
	}
	for i, option := range options {
		if !option.Name.IsNull() && !option.Name.IsUnknown() && option.Name.ValueString() == "dhcp-lease-time" {
			if !option.Value.IsNull() && !option.Value.IsUnknown() &&
				option.Value.ValueString() != strconv.FormatInt(m.LeaseTime.ValueInt64(), 10) {
				resp.Diagnostics.AddAttributeError(
					niosPath.AtName("option_list").AtListIndex(i).AtName("value"),
					"Invalid configuration for Lease Time",
					"lease_time attribute must match the 'value' attribute for DHCP Option 'dhcp-lease-time'.",
				)
			}
		}
	}
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
