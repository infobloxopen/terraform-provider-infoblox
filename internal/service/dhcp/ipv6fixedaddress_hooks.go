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

// ValidateIpv6fixedaddress validates the Ipv6fixedaddress configuration.
func ValidateIpv6fixedaddress(ctx context.Context, data Ipv6fixedaddressModel, resp *resource.ValidateConfigResponse) {
	if nios := flex.ExpandNestedObject[NIOSIpv6fixedaddressModel](ctx, data.NIOS, &resp.Diagnostics); nios != nil {
		validateIpv6fixedaddressNIOSConfig(ctx, nios, resp)
	}
	if uddi := flex.ExpandNestedObject[UDDIIpv6fixedaddressModel](ctx, data.UDDI, &resp.Diagnostics); uddi != nil {
		validateIpv6fixedaddressUDDIConfig(ctx, uddi, resp)
	}
}

func validateIpv6fixedaddressNIOSConfig(ctx context.Context, m *NIOSIpv6fixedaddressModel, resp *resource.ValidateConfigResponse) {
	niosPath := path.Root("nios")

	if !m.AddressType.IsUnknown() {
		addressType := "ADDRESS"
		if !m.AddressType.IsNull() {
			addressType = m.AddressType.ValueString()
		}

		requireAddr := func() {
			if !m.Ipv6addr.IsUnknown() && m.Ipv6addr.IsNull() {
				resp.Diagnostics.AddAttributeError(
					niosPath.AtName("ipv6addr"),
					"Missing Required Attribute",
					"When address_type is set to '"+addressType+"', the 'ipv6addr' attribute must be specified.",
				)
			}
		}
		requirePrefix := func() {
			if !m.Ipv6prefix.IsUnknown() && m.Ipv6prefix.IsNull() {
				resp.Diagnostics.AddAttributeError(
					niosPath.AtName("ipv6prefix"),
					"Missing Required Attribute",
					"When address_type is set to '"+addressType+"', the 'ipv6prefix' attribute must be specified.",
				)
			}
			if !m.Ipv6prefixBits.IsUnknown() && m.Ipv6prefixBits.IsNull() {
				resp.Diagnostics.AddAttributeError(
					niosPath.AtName("ipv6prefix_bits"),
					"Missing Required Attribute",
					"When address_type is set to '"+addressType+"', the 'ipv6prefix_bits' attribute must be specified.",
				)
			}
		}

		switch addressType {
		case "ADDRESS":
			requireAddr()
		case "PREFIX":
			requirePrefix()
		case "BOTH":
			requireAddr()
			requirePrefix()
		}
	}

	// DHCP options validation
	utils.ValidateDHCPOptionsConfig(ctx, m.Options, niosPath.AtName("options"), &resp.Diagnostics)

	var hasDhcpLeaseTime bool
	var dhcpLeaseTimeValue string

	if !m.Options.IsNull() && !m.Options.IsUnknown() {
		var options []Ipv6fixedaddressOptionsModel
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

	// allow_telnet requires cli_credentials holding both an SSH and a TELNET credential
	if !m.AllowTelnet.IsNull() && !m.AllowTelnet.IsUnknown() && m.AllowTelnet.ValueBool() {
		var hasSSH, hasTelnet bool

		if !m.CliCredentials.IsNull() && !m.CliCredentials.IsUnknown() {
			var cliCredentials []Ipv6fixedaddressCliCredentialsModel
			resp.Diagnostics.Append(m.CliCredentials.ElementsAs(ctx, &cliCredentials, false)...)
			if resp.Diagnostics.HasError() {
				return
			}

			for _, credential := range cliCredentials {
				if credential.CredentialType.IsNull() || credential.CredentialType.IsUnknown() {
					continue
				}
				switch credential.CredentialType.ValueString() {
				case "SSH":
					hasSSH = true
				case "TELNET":
					hasTelnet = true
				}
			}
		}

		if !hasSSH {
			resp.Diagnostics.AddAttributeError(
				niosPath.AtName("allow_telnet"),
				"Invalid configuration",
				"The 'cli_credentials' must contain credentials with 'credential_type' set to 'SSH'.",
			)
		}
		if !hasTelnet {
			resp.Diagnostics.AddAttributeError(
				niosPath.AtName("allow_telnet"),
				"Invalid configuration",
				"The 'allow_telnet' attribute must be set to false when 'cli_credentials' does not contain any credentials with 'credential_type' set to 'TELNET'.",
			)
		}
	}
}

func validateIpv6fixedaddressUDDIConfig(ctx context.Context, m *UDDIIpv6fixedaddressModel, resp *resource.ValidateConfigResponse) {
}

func PostFlattenIpv6fixedaddressNIOS(ctx context.Context, planned, flattened *NIOSIpv6fixedaddressModel, diags *diag.Diagnostics) {
	if planned == nil || flattened == nil {
		return
	}

	// NIOS returns cli_credentials and options in its own order; restore the
	// configured order so Terraform does not report ordering-only diffs.
	if !planned.CliCredentials.IsUnknown() {
		if reordered, d := utils.ReorderAndFilterNestedListResponse(ctx, planned.CliCredentials, flattened.CliCredentials, "credential_type"); !d.HasError() {
			if reorderedList, ok := reordered.(basetypes.ListValue); ok {
				flattened.CliCredentials = reorderedList
			}
		}
	}

	if !planned.Options.IsUnknown() {
		if reordered, d := utils.ReorderAndFilterDHCPOptions(ctx, planned.Options, flattened.Options); !d.HasError() {
			if reorderedList, ok := reordered.(basetypes.ListValue); ok {
				flattened.Options = reorderedList
			}
		}
	}

	// template is a write-only NIOS attribute; NIOS never echoes it back, so keep
	// the configured value instead of nulling it out on read.
	if flattened.Template.IsNull() && !planned.Template.IsNull() && !planned.Template.IsUnknown() {
		flattened.Template = planned.Template
	}
}
