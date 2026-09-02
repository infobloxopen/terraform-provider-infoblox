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

// ValidateSharednetwork validates the Sharednetwork configuration.
func ValidateSharednetwork(ctx context.Context, data SharednetworkModel, resp *resource.ValidateConfigResponse) {
	if nios := flex.ExpandNestedObject[NIOSSharednetworkModel](ctx, data.NIOS, &resp.Diagnostics); nios != nil {
		validateSharednetworkNIOSConfig(ctx, nios, resp)
	}
}

func validateSharednetworkNIOSConfig(ctx context.Context, m *NIOSSharednetworkModel, resp *resource.ValidateConfigResponse) {
	niosPath := path.Root("nios")
	// DHCP options validation
	utils.ValidateDHCPOptionsConfig(ctx, m.Options, niosPath.AtName("options"), &resp.Diagnostics)
}

func PostFlattenSharednetworkNIOS(ctx context.Context, planned, flattened *NIOSSharednetworkModel, diags *diag.Diagnostics) {
	if planned == nil {
		return
	}
	if !planned.Options.IsUnknown() {
		reordered, d := utils.ReorderAndFilterDHCPOptions(ctx, planned.Options, flattened.Options)
		diags.Append(*d...)
		if d.HasError() {
			return
		}
		if reorderedList, ok := reordered.(basetypes.ListValue); ok {
			flattened.Options = reorderedList
		}
	}
	reOrderedNetworks, d := utils.ReorderAndFilterNestedListResponse(ctx, planned.Networks, flattened.Networks, "ref")
	diags.Append(*d...)
	if !diags.HasError() {
		flattened.Networks = reOrderedNetworks.(basetypes.ListValue)
	}
}
