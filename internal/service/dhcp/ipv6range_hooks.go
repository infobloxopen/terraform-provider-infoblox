package dhcp

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/utils"
)

// ValidateIpv6range validates the Ipv6range configuration.
func ValidateIpv6range(ctx context.Context, data Ipv6rangeModel, resp *resource.ValidateConfigResponse) {
	if nios := flex.ExpandNestedObject[NIOSIpv6rangeModel](ctx, data.NIOS, &resp.Diagnostics); nios != nil {
		validateIpv6rangeNIOSConfig(ctx, nios, resp)
	}
	if uddi := flex.ExpandNestedObject[UDDIIpv6rangeModel](ctx, data.UDDI, &resp.Diagnostics); uddi != nil {
		validateIpv6rangeUDDIConfig(ctx, uddi, resp)
	}
}

func validateIpv6rangeNIOSConfig(ctx context.Context, m *NIOSIpv6rangeModel, resp *resource.ValidateConfigResponse) {
	niosPath := path.Root("nios")

	// A value from a variable is unknown at validation time, so each check below
	// is skipped (never failed) while any value it reads is still unknown.
	//
	// The legacy provider's checks that depend on the use_* flags (use_enable_discovery,
	// use_discovery_basic_polling_settings, use_blackout_setting) are not ported: here
	// those flags are derived from whether the field is set, so they cannot disagree.
	if !m.AddressType.IsUnknown() {
		addressType := "ADDRESS"
		if !m.AddressType.IsNull() {
			addressType = m.AddressType.ValueString()
		}

		switch addressType {
		case "ADDRESS":
			if !m.StartAddr.IsUnknown() && !m.EndAddr.IsUnknown() {
				if m.StartAddr.IsNull() || m.EndAddr.IsNull() {
					resp.Diagnostics.AddError(
						"Configuration Error",
						"When address_type is set to 'ADDRESS' (default), both start_addr and end_addr must be specified.",
					)
				}
			}
			if !m.Ipv6StartPrefix.IsUnknown() && !m.Ipv6EndPrefix.IsUnknown() && !m.Ipv6PrefixBits.IsUnknown() {
				if !m.Ipv6StartPrefix.IsNull() || !m.Ipv6EndPrefix.IsNull() || !m.Ipv6PrefixBits.IsNull() {
					resp.Diagnostics.AddError(
						"Configuration Error",
						"When address_type is 'ADDRESS' (default), ipv6_start_prefix, ipv6_end_prefix, and ipv6_prefix_bits cannot be specified.",
					)
				}
			}
		case "PREFIX":
			if !m.Ipv6StartPrefix.IsUnknown() && !m.Ipv6EndPrefix.IsUnknown() && !m.Ipv6PrefixBits.IsUnknown() {
				if m.Ipv6StartPrefix.IsNull() || m.Ipv6EndPrefix.IsNull() || m.Ipv6PrefixBits.IsNull() {
					resp.Diagnostics.AddError(
						"Configuration Error",
						"When address_type is set to 'PREFIX', ipv6_start_prefix, ipv6_end_prefix, and ipv6_prefix_bits must be specified.",
					)
				}
			}
			if !m.StartAddr.IsUnknown() && !m.EndAddr.IsUnknown() {
				if !m.StartAddr.IsNull() || !m.EndAddr.IsNull() {
					resp.Diagnostics.AddError(
						"Configuration Error",
						"When address_type is 'PREFIX', start_addr and end_addr cannot be specified.",
					)
				}
			}
		case "BOTH":
			if !m.StartAddr.IsUnknown() && !m.EndAddr.IsUnknown() && !m.Ipv6StartPrefix.IsUnknown() && !m.Ipv6EndPrefix.IsUnknown() && !m.Ipv6PrefixBits.IsUnknown() {
				if m.StartAddr.IsNull() || m.EndAddr.IsNull() || m.Ipv6StartPrefix.IsNull() || m.Ipv6EndPrefix.IsNull() || m.Ipv6PrefixBits.IsNull() {
					resp.Diagnostics.AddError(
						"Configuration Error",
						"When address_type is set to 'BOTH', start_addr, end_addr, ipv6_start_prefix, ipv6_end_prefix, and ipv6_prefix_bits must be specified.",
					)
				}
			}
		}
	}

	// Validate discovery_blackout_setting blackout_schedule
	if !m.DiscoveryBlackoutSetting.IsNull() && !m.DiscoveryBlackoutSetting.IsUnknown() {
		utils.ValidateScheduleConfig(
			m.DiscoveryBlackoutSetting,
			"blackout_schedule",
			niosPath.AtName("discovery_blackout_setting"),
			&resp.Diagnostics,
		)
	}

	// Validate port_control_blackout_setting blackout_schedule
	if !m.PortControlBlackoutSetting.IsNull() && !m.PortControlBlackoutSetting.IsUnknown() {
		utils.ValidateScheduleConfig(
			m.PortControlBlackoutSetting,
			"blackout_schedule",
			niosPath.AtName("port_control_blackout_setting"),
			&resp.Diagnostics,
		)
	}

	if !m.ServerAssociationType.IsUnknown() {
		serverAssociationType := "NONE"
		if !m.ServerAssociationType.IsNull() {
			serverAssociationType = m.ServerAssociationType.ValueString()
		}

		// If server_association_type is MEMBER, member field must be set
		if serverAssociationType == "MEMBER" {
			if !m.Member.IsUnknown() && m.Member.IsNull() {
				resp.Diagnostics.AddAttributeError(
					niosPath.AtName("member"),
					"Invalid Configuration",
					"The 'member' field must be set when 'server_association_type' is set to 'MEMBER'.",
				)
			}
		}

		// If server_association_type is NONE, member field cannot be set
		if serverAssociationType == "NONE" {
			if !m.Member.IsNull() && !m.Member.IsUnknown() {
				resp.Diagnostics.AddAttributeError(
					niosPath.AtName("member"),
					"Invalid Configuration",
					"The 'member' field cannot be set when 'server_association_type' is set to 'NONE' (default).",
				)
			}
		}
	}
}

func validateIpv6rangeUDDIConfig(ctx context.Context, m *UDDIIpv6rangeModel, resp *resource.ValidateConfigResponse) {
}
