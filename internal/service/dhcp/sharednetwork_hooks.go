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

	if !m.DdnsServerAlwaysUpdates.IsNull() && !m.DdnsServerAlwaysUpdates.IsUnknown() {
		// Check if ddns_use_option81 is set to false
		if !m.DdnsUseOption81.IsUnknown() && (m.DdnsUseOption81.IsNull() || !m.DdnsUseOption81.ValueBool()) {
			resp.Diagnostics.AddAttributeError(
				niosPath.AtName("ddns_server_always_updates"),
				"Invalid Configuration",
				"ddns_use_option81 must be set to true if ddns_server_always_updates is configured.",
			)
		}
	}

	validateSharednetworkIgnoreClientIdentifier(m, niosPath, &resp.Diagnostics)
}

// ignore_id and ignore_client_identifier validation
func validateSharednetworkIgnoreClientIdentifier(m *NIOSSharednetworkModel, niosPath path.Path, diags *diag.Diagnostics) {
	if m.IgnoreClientIdentifier.IsUnknown() || m.IgnoreId.IsUnknown() {
		return
	}

	ignoreClientIdentifier := !m.IgnoreClientIdentifier.IsNull() && m.IgnoreClientIdentifier.ValueBool()
	ignoreIdClient := !m.IgnoreId.IsNull() && m.IgnoreId.ValueString() == "CLIENT"

	switch {
	case ignoreClientIdentifier && !ignoreIdClient:
		diags.AddAttributeError(
			niosPath.AtName("ignore_id"),
			"Invalid Configuration",
			"ignore_id must be set to \"CLIENT\" when ignore_client_identifier is set to true.",
		)
	case ignoreIdClient && !ignoreClientIdentifier:
		diags.AddAttributeError(
			niosPath.AtName("ignore_client_identifier"),
			"Invalid Configuration",
			"ignore_client_identifier must be set to true when ignore_id is set to \"CLIENT\".",
		)
	}
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
