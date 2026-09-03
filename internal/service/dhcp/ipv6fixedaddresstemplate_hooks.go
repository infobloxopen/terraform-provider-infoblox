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

// ValidateIpv6fixedaddresstemplate validates the Ipv6fixedaddresstemplate configuration.
func ValidateIpv6fixedaddresstemplate(ctx context.Context, data Ipv6fixedaddresstemplateModel, resp *resource.ValidateConfigResponse) {
	if nios := flex.ExpandNestedObject[NIOSIpv6fixedaddresstemplateModel](ctx, data.NIOS, &resp.Diagnostics); nios != nil {
		validateIpv6fixedaddresstemplateNIOSConfig(ctx, nios, resp)
	}
}

func validateIpv6fixedaddresstemplateNIOSConfig(ctx context.Context, m *NIOSIpv6fixedaddresstemplateModel, resp *resource.ValidateConfigResponse) {
	niosPath := path.Root("nios")

	// DHCP options validation
	utils.ValidateDHCPOptionsConfig(ctx, m.Options, niosPath.AtName("options"), &resp.Diagnostics)

	var hasDhcpLeaseTime bool
	var dhcpLeaseTimeValue string

	if !m.Options.IsNull() && !m.Options.IsUnknown() {
		var options []Ipv6fixedaddresstemplateOptionsModel
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
}

func PostFlattenIpv6fixedaddresstemplateNIOS(ctx context.Context, planned, flattened *NIOSIpv6fixedaddresstemplateModel, diags *diag.Diagnostics) {
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
