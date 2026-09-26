package ipam

import (
	"context"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	niosipam "github.com/infobloxopen/infoblox-nios-go-client/ipam"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/core"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/dynamicallocation"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/utils"
)

// ValidateIpv6network validates the Ipv6network configuration.
func ValidateIpv6network(ctx context.Context, data Ipv6networkModel, resp *resource.ValidateConfigResponse) {
	if nios := flex.ExpandNestedObject[NIOSIpv6networkModel](ctx, data.NIOS, &resp.Diagnostics); nios != nil {
		validateIpv6networkNIOSConfig(ctx, nios, resp)
	}
	if uddi := flex.ExpandNestedObject[UDDIIpv6networkModel](ctx, data.UDDI, &resp.Diagnostics); uddi != nil {
		validateIpv6networkUDDIConfig(ctx, uddi, resp)
	}
}

func validateIpv6networkNIOSConfig(ctx context.Context, m *NIOSIpv6networkModel, resp *resource.ValidateConfigResponse) {
	niosPath := path.Root("nios")

	// DHCP options validation
	utils.ValidateDHCPOptionsConfig(ctx, m.Options, niosPath.AtName("options"), &resp.Diagnostics)

	var hasDhcpLeaseTime bool
	var dhcpLeaseTimeValue string

	if !m.Options.IsNull() && !m.Options.IsUnknown() {
		var options []Ipv6networkOptionsModel
		resp.Diagnostics.Append(m.Options.ElementsAs(ctx, &options, false)...)
		if resp.Diagnostics.HasError() {
			return
		}

		for _, option := range options {
			if option.Name.IsNull() || option.Name.IsUnknown() ||
				option.Value.IsNull() || option.Value.IsUnknown() {
				continue
			}

			switch option.Name.ValueString() {
			case "dhcp-lease-time":
				hasDhcpLeaseTime = true
				dhcpLeaseTimeValue = option.Value.ValueString()

			case "domain-name":
				// domain_name attribute must match the value of option 'domain-name'
				if !m.DomainName.IsNull() && !m.DomainName.IsUnknown() &&
					option.Value.ValueString() != m.DomainName.ValueString() {
					resp.Diagnostics.AddAttributeError(
						niosPath.AtName("domain_name"),
						"Invalid configuration for Domain Name",
						"domain_name attribute must match the 'value' attribute for DHCP Option 'domain-name'.",
					)
				}
			}
		}

		// When dhcp-lease-time option is set, valid_lifetime attribute must have the same value as option value
		if hasDhcpLeaseTime && !m.ValidLifetime.IsNull() && !m.ValidLifetime.IsUnknown() {
			if dhcpLeaseTimeValue != strconv.FormatInt(m.ValidLifetime.ValueInt64(), 10) {
				resp.Diagnostics.AddAttributeError(
					niosPath.AtName("valid_lifetime"),
					"Invalid configuration for Valid Lifetime",
					"valid_lifetime attribute must match the 'value' attribute for DHCP Option 'dhcp-lease-time'.",
				)
			}
		}
	}

	// Preferred lifetime must be less than or equal to valid lifetime
	if !m.PreferredLifetime.IsNull() && !m.PreferredLifetime.IsUnknown() {
		switch {
		case !m.ValidLifetime.IsNull() && !m.ValidLifetime.IsUnknown():
			if m.PreferredLifetime.ValueInt64() > m.ValidLifetime.ValueInt64() {
				resp.Diagnostics.AddAttributeError(
					niosPath.AtName("preferred_lifetime"),
					"Invalid configuration",
					"The 'preferred_lifetime' must be less than or equal to 'valid_lifetime'.",
				)
			}

		case hasDhcpLeaseTime:
			// if valid_lifetime is not set, compare with DHCP lease time
			if leaseTime, err := strconv.ParseInt(dhcpLeaseTimeValue, 10, 64); err == nil {
				if m.PreferredLifetime.ValueInt64() > leaseTime {
					resp.Diagnostics.AddAttributeError(
						niosPath.AtName("preferred_lifetime"),
						"Invalid configuration",
						"The 'preferred_lifetime' must be less than or equal to 'dhcp-lease-time' (valid_lifetime) option value.",
					)
				}
			}

		case !m.ValidLifetime.IsUnknown() && !m.Options.IsUnknown():
			resp.Diagnostics.AddAttributeError(
				niosPath.AtName("preferred_lifetime"),
				"Invalid configuration",
				"Either 'valid_lifetime' attribute or 'dhcp-lease-time' option must be set when 'preferred_lifetime' is specified.",
			)
		}
	}

	// Check for valid lifetime or dhcp-lease-time when preferred_lifetime is NOT set
	if m.PreferredLifetime.IsNull() {
		// validate that valid_lifetime is >= 27000
		if !m.ValidLifetime.IsNull() && !m.ValidLifetime.IsUnknown() && m.ValidLifetime.ValueInt64() < 27000 {
			resp.Diagnostics.AddAttributeError(
				niosPath.AtName("valid_lifetime"),
				"Invalid configuration",
				"When 'preferred_lifetime' is not set ,"+
					"'valid_lifetime' must be greater than or equal to 27000.",
			)
		}

		// validate that dhcp-lease-time is >= 27000
		if hasDhcpLeaseTime {
			if leaseTime, err := strconv.ParseInt(dhcpLeaseTimeValue, 10, 64); err == nil && leaseTime < 27000 {
				resp.Diagnostics.AddAttributeError(
					niosPath.AtName("options"),
					"Invalid configuration",
					"When 'preferred_lifetime' is not set, the DHCP option "+
						"'dhcp-lease-time' must be greater than or equal to 27000.",
				)
			}
		}
	}

	// rir_organization is required when the network is registered with an RIR
	if !m.RirRegistrationStatus.IsNull() && !m.RirRegistrationStatus.IsUnknown() &&
		m.RirRegistrationStatus.ValueString() == "REGISTERED" && m.RirOrganization.IsNull() {
		resp.Diagnostics.AddAttributeError(
			niosPath.AtName("rir_organization"),
			"Missing RIR Organization",
			"The 'rir_organization' attribute must be set when 'rir_registration_status' is set to REGISTERED.",
		)
	}

	// ddns_server_always_updates cannot be false while ddns_enable_option_fqdn is false
	if !m.DdnsEnableOptionFqdn.IsNull() && !m.DdnsEnableOptionFqdn.IsUnknown() && !m.DdnsEnableOptionFqdn.ValueBool() &&
		!m.DdnsServerAlwaysUpdates.IsNull() && !m.DdnsServerAlwaysUpdates.IsUnknown() && !m.DdnsServerAlwaysUpdates.ValueBool() {
		resp.Diagnostics.AddAttributeError(
			niosPath.AtName("ddns_server_always_updates"),
			"Invalid DDNS Configuration",
			"You cannot set 'ddns_server_always_updates' to false when 'ddns_enable_option_fqdn' is false.",
		)
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

	// enabled_attributes is required when subscribe_settings is configured
	if !m.SubscribeSettings.IsNull() && !m.SubscribeSettings.IsUnknown() {
		attrs := m.SubscribeSettings.Attributes()
		enabledAttrs, exists := attrs["enabled_attributes"]
		if !exists || enabledAttrs.IsNull() {
			resp.Diagnostics.AddAttributeError(
				niosPath.AtName("subscribe_settings").AtName("enabled_attributes"),
				"Missing Required Attribute",
				"The 'enabled_attributes' attribute is required when 'subscribe_settings' is configured.",
			)
		}
	}
}

func validateIpv6networkUDDIConfig(ctx context.Context, m *UDDIIpv6networkModel, resp *resource.ValidateConfigResponse) {
}

func BuildIpv6networkFuncCall(ctx context.Context, data types.Object, diags *diag.Diagnostics) *niosipam.FuncCall {
	if data.IsNull() || data.IsUnknown() {
		return nil
	}

	var m dynamicallocation.NextAvailableNetworkModel
	diags.Append(data.As(ctx, &m, basetypes.ObjectAsOptions{})...)
	if diags.HasError() {
		return nil
	}

	return m.FuncCall(ctx, "Network", "ipv6network", diags)
}

func PostFlattenIpv6networkNIOS(ctx context.Context, planned, flattened *NIOSIpv6networkModel, diags *diag.Diagnostics) {
	if planned != nil && !planned.Options.IsUnknown() {
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

func (r *Ipv6networkResource) isIpv6networkContainerConversionError(err error) bool {
	errVal := err.Error()
	return (strings.Contains(errVal, "The search parameters") &&
		strings.Contains(errVal, "for object ipv6network did not return any result")) ||
		strings.Contains(errVal, "will overlap an existing network")
}

func (r *Ipv6networkResource) isIpv6networkConvertedToContainer(ctx context.Context, data *Ipv6networkModel) bool {
	if r.backend != core.BackendNIOS || r.containerService == nil {
		return false
	}

	var diags diag.Diagnostics
	nios := flex.ExpandNestedObject[NIOSIpv6networkModel](ctx, data.NIOS, &diags)
	if nios == nil || diags.HasError() {
		return false
	}

	// Try to fetch as Network container
	records, _, _, err := r.containerService.List(ctx, &core.ListOptions{
		Filters: map[string]string{
			"nios.network": nios.Network.ValueString(),
		},
	})
	return err == nil && len(records) > 0
}
