package dhcp

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/utils"
)

// ValidateFixedaddresstemplate validates the Fixedaddresstemplate configuration.
func ValidateFixedaddresstemplate(ctx context.Context, data FixedaddresstemplateModel, resp *resource.ValidateConfigResponse) {
	if nios := flex.ExpandNestedObject[NIOSFixedaddresstemplateModel](ctx, data.NIOS, &resp.Diagnostics); nios != nil {
		validateFixedaddresstemplateNIOSConfig(ctx, nios, resp)
	}
}

func validateFixedaddresstemplateNIOSConfig(ctx context.Context, m *NIOSFixedaddresstemplateModel, resp *resource.ValidateConfigResponse) {
	// The legacy NIOS provider's only ValidateConfig check is the DHCP options rules:
	// every option needs a value and a name or num, an empty value is only allowed for
	// special options, and use_option is only valid on special options. The shared helper
	// holds those rules and skips unknown values, so a value from a variable never fails here.
	utils.ValidateDHCPOptionsConfig(ctx, m.Options, path.Root("nios").AtName("options"), &resp.Diagnostics)
}

// PostFlattenFixedaddresstemplateNIOS keeps the options list in the planned order and drops
// the options NIOS adds on its own. A template with no options still comes back with a default
// dhcp-lease-time option (use_option false), which would otherwise appear as a new element.
func PostFlattenFixedaddresstemplateNIOS(ctx context.Context, planned, flattened *NIOSFixedaddresstemplateModel, diags *diag.Diagnostics) {
	if planned == nil || flattened == nil {
		return
	}

	if !planned.Options.IsUnknown() {
		if reordered, d := utils.ReorderAndFilterDHCPOptions(ctx, planned.Options, flattened.Options); !d.HasError() {
			if reorderedList, ok := reordered.(basetypes.ListValue); ok {
				flattened.Options = reorderedList
			}
		}
	}
}
