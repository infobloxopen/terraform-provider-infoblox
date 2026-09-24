package dhcp

import (
	"context"
	"strconv"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	niosdhcp "github.com/infobloxopen/infoblox-nios-go-client/dhcp"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
	internaltypes "github.com/infobloxopen/terraform-provider-infoblox/internal/types"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/utils"
)

// ValidateIpv6sharednetwork validates the Ipv6sharednetwork configuration.
func ValidateIpv6sharednetwork(ctx context.Context, data Ipv6sharednetworkModel, resp *resource.ValidateConfigResponse) {
	if nios := flex.ExpandNestedObject[NIOSIpv6sharednetworkModel](ctx, data.NIOS, &resp.Diagnostics); nios != nil {
		validateIpv6sharednetworkNIOSConfig(ctx, nios, resp)
	}
}

func validateIpv6sharednetworkNIOSConfig(ctx context.Context, m *NIOSIpv6sharednetworkModel, resp *resource.ValidateConfigResponse) {
	niosPath := path.Root("nios")
	// DHCP options validation
	utils.ValidateDHCPOptionsConfig(ctx, m.Options, niosPath.AtName("options"), &resp.Diagnostics)

	if !m.DdnsServerAlwaysUpdates.IsNull() && !m.DdnsServerAlwaysUpdates.IsUnknown() {
		// Check if ddns_use_option81 is not set to true.
		if !m.DdnsUseOption81.IsUnknown() && (m.DdnsUseOption81.IsNull() || !m.DdnsUseOption81.ValueBool()) {
			resp.Diagnostics.AddAttributeError(
				niosPath.AtName("ddns_server_always_updates"),
				"Invalid Configuration",
				"ddns_use_option81 must be set to true if ddns_server_always_updates is configured.",
			)
		}
	}

	var dhcpLeaseTimeValue string
	var hasDhcpLeaseTime bool

	// Check if options are defined
	if !m.Options.IsNull() && !m.Options.IsUnknown() {
		var options []Ipv6sharednetworkOptionsModel
		resp.Diagnostics.Append(m.Options.ElementsAs(ctx, &options, false)...)
		if resp.Diagnostics.HasError() {
			return
		}

		for _, option := range options {
			if option.Name.ValueString() == "dhcp-lease-time" && !option.Value.IsNull() && !option.Value.IsUnknown() {
				hasDhcpLeaseTime = true
				dhcpLeaseTimeValue = option.Value.ValueString()
			}

			// domain_name attribute must match the value of option 'domain-name'
			if option.Name.ValueString() == "domain-name" {
				if !m.DomainName.IsNull() && !m.DomainName.IsUnknown() &&
					!option.Value.IsNull() && !option.Value.IsUnknown() &&
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
		if m.ValidLifetime.IsNull() && !hasDhcpLeaseTime && !m.Options.IsUnknown() {
			resp.Diagnostics.AddAttributeError(
				niosPath.AtName("preferred_lifetime"),
				"Invalid configuration",
				"Either 'valid_lifetime' attribute or 'dhcp-lease-time' option must be set when 'preferred_lifetime' is specified.",
			)
		} else if !m.ValidLifetime.IsNull() && !m.ValidLifetime.IsUnknown() {
			if m.PreferredLifetime.ValueInt64() > m.ValidLifetime.ValueInt64() {
				resp.Diagnostics.AddAttributeError(
					niosPath.AtName("preferred_lifetime"),
					"Invalid configuration",
					"The 'preferred_lifetime' must be less than or equal to 'valid_lifetime'.",
				)
			}
		} else if hasDhcpLeaseTime {
			// if valid_lifetime is not set, compare with DHCP lease time
			if dhcpLeaseTimeInt, err := strconv.ParseInt(dhcpLeaseTimeValue, 10, 64); err == nil {
				if m.PreferredLifetime.ValueInt64() > dhcpLeaseTimeInt {
					resp.Diagnostics.AddAttributeError(
						niosPath.AtName("preferred_lifetime"),
						"Invalid configuration",
						"The 'preferred_lifetime' must be less than or equal to 'dhcp-lease-time' (valid_lifetime) option value.",
					)
				}
			}
		}
	}

	// Check for valid lifetime or dhcp-lease-time when preferred_lifetime is NOT set
	if m.PreferredLifetime.IsNull() {
		// validate that valid_lifetime is >= 27000
		if !m.ValidLifetime.IsNull() && !m.ValidLifetime.IsUnknown() {
			if m.ValidLifetime.ValueInt64() < 27000 {
				resp.Diagnostics.AddAttributeError(
					niosPath.AtName("valid_lifetime"),
					"Invalid configuration",
					"When 'preferred_lifetime' is not set ,"+
						"'valid_lifetime' must be greater than or equal to 27000.",
				)
			}
		}

		// validate that dhcp-lease-time  is >= 27000
		if hasDhcpLeaseTime {
			if dhcpLeaseTimeInt, err := strconv.ParseInt(dhcpLeaseTimeValue, 10, 64); err == nil {
				if dhcpLeaseTimeInt < 27000 {
					resp.Diagnostics.AddAttributeError(
						niosPath.AtName("options"),
						"Invalid configuration",
						"When 'preferred_lifetime' is not set, the DHCP option "+
							"'dhcp-lease-time' must be greater than or equal to 27000.",
					)
				}
			}
		}
	}
}

func PostFlattenIpv6sharednetworkNIOS(ctx context.Context, planned, flattened *NIOSIpv6sharednetworkModel, diags *diag.Diagnostics) {
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

func ExpandNetworks(ctx context.Context, networks internaltypes.UnorderedListValue, diags *diag.Diagnostics) []niosdhcp.Ipv6sharednetworkNetworks {
	if networks.IsNull() || networks.IsUnknown() {
		return nil
	}

	var networkRefs []string
	diags.Append(networks.ElementsAs(ctx, &networkRefs, false)...)
	if diags.HasError() {
		return nil
	}

	result := make([]niosdhcp.Ipv6sharednetworkNetworks, len(networkRefs))
	for i, ref := range networkRefs {
		result[i] = niosdhcp.Ipv6sharednetworkNetworks{
			Ref: &ref,
		}
	}
	return result
}

func FlattenNetworks(ctx context.Context, networks []niosdhcp.Ipv6sharednetworkNetworks, diags *diag.Diagnostics) internaltypes.UnorderedListValue {
	if networks == nil {
		return internaltypes.NewUnorderedListValueNull(types.StringType)
	}

	networkRefs := make([]string, len(networks))
	for i, network := range networks {
		if network.Ref != nil {
			networkRefs[i] = *network.Ref
		}
	}

	listValue, d := internaltypes.NewUnorderedListValueFrom(ctx, types.StringType, networkRefs)
	diags.Append(d...)
	return listValue
}
