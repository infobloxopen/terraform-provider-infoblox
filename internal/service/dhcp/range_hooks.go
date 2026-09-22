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

// ValidateRange validates the Range configuration.
func ValidateRange(ctx context.Context, data RangeModel, resp *resource.ValidateConfigResponse) {
	if nios := flex.ExpandNestedObject[NIOSRangeModel](ctx, data.NIOS, &resp.Diagnostics); nios != nil {
		validateRangeNIOSConfig(ctx, nios, resp)
	}
	if uddi := flex.ExpandNestedObject[UDDIRangeModel](ctx, data.UDDI, &resp.Diagnostics); uddi != nil {
		validateRangeUDDIConfig(ctx, uddi, resp)
	}
}

func validateRangeNIOSConfig(ctx context.Context, m *NIOSRangeModel, resp *resource.ValidateConfigResponse) {
	niosPath := path.Root("nios")
	// DHCP options validation
	utils.ValidateDHCPOptionsConfig(ctx, m.Options, niosPath.AtName("options"), &resp.Diagnostics)

	if !m.ServerAssociationType.IsUnknown() {
		serverAssociationType := "NONE"
		if !m.ServerAssociationType.IsNull() {
			serverAssociationType = m.ServerAssociationType.ValueString()
		}

		// If server_association_type is MEMBER, member field must be set
		if serverAssociationType == "MEMBER" {
			if m.Member.IsNull() {
				resp.Diagnostics.AddAttributeError(
					path.Root("member"),
					"Invalid Configuration",
					"The 'member' field must be set when 'server_association_type' is set to 'MEMBER'.",
				)
			}
		}

		// If server_association_type is NONE, member field cannot be set
		if serverAssociationType == "NONE" {
			if !m.Member.IsNull() && !m.Member.IsUnknown() {
				resp.Diagnostics.AddAttributeError(
					path.Root("member"),
					"Invalid Configuration",
					"The 'member' field cannot be set when 'server_association_type' is set to 'NONE' (default).",
				)
			}
			if !m.MsServer.IsNull() && !m.MsServer.IsUnknown() {
				resp.Diagnostics.AddAttributeError(
					path.Root("ms_server"),
					"Invalid Configuration",
					"The 'ms_server' field cannot be set when 'server_association_type' is set to 'NONE' (default). "+
						"Modify the 'server_association_type' field to 'MS_SERVER' to allow setting 'ms_server'.",
				)
			}
		}
	}
	// Validate discovery_blackout_setting blackout_schedule
	if !m.DiscoveryBlackoutSetting.IsNull() && !m.DiscoveryBlackoutSetting.IsUnknown() {
		utils.ValidateScheduleConfig(
			m.DiscoveryBlackoutSetting,
			"blackout_schedule",
			path.Root("discovery_blackout_setting"),
			&resp.Diagnostics,
		)
	}

	// Validate port_control_blackout_setting blackout_schedule
	if !m.PortControlBlackoutSetting.IsNull() && !m.PortControlBlackoutSetting.IsUnknown() {
		utils.ValidateScheduleConfig(
			m.PortControlBlackoutSetting,
			"blackout_schedule",
			path.Root("port_control_blackout_setting"),
			&resp.Diagnostics,
		)
	}
}

func validateRangeUDDIConfig(ctx context.Context, m *UDDIRangeModel, resp *resource.ValidateConfigResponse) {
}

func PostFlattenRangeNIOS(ctx context.Context, planned, flattened *NIOSRangeModel, diags *diag.Diagnostics) {
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
}
