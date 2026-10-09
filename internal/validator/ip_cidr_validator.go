package validator

import (
	"context"
	"fmt"
	"net"

	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

// IPCIDRValidator validates that the provided string is either:
// - A valid FQDN
// - A valid IPv4 address or IPV4 CIDR
// - A valid IPv6 address or IPV6 CIDR
type IPCIDRValidator struct{}

// Description returns a plain text description of the validator's behavior.
func (v IPCIDRValidator) Description(ctx context.Context) string {
	return "String must be a valid FQDN, IPv4, IPv4 CIDR, IPv6, or IPv6 CIDR"
}

// MarkdownDescription returns a markdown formatted description of the validator's behavior.
func (v IPCIDRValidator) MarkdownDescription(ctx context.Context) string {
	return "String must be a valid IPv4, IPv4 CIDR, IPv6, or IPv6 CIDR"
}

// Validate performs the validation.
func (v IPCIDRValidator) ValidateString(ctx context.Context, request validator.StringRequest, response *validator.StringResponse) {
	if request.ConfigValue.IsNull() || request.ConfigValue.IsUnknown() {
		return
	}

	value := request.ConfigValue.ValueString()

	// Check if the value is a valid IPv4 or IPv6 address
	if ip := net.ParseIP(value); ip != nil {
		// Valid IP address (either IPv4 or IPv6)
		return
	}

	// Check if the value is a valid CIDR notation (IPv4 or IPv6)
	if _, _, err := net.ParseCIDR(value); err == nil {
		// Valid CIDR notation (either IPv4 or IPv6)
		return
	}
	// If we get here, it's not a valid format
	response.Diagnostics.AddAttributeError(
		request.Path,
		"Invalid Value Format",
		fmt.Sprintf("The value '%s' is not a valid IPv4, IPv4 CIDR, IPv6, or IPv6 CIDR.", value),
	)
}

// IsValidIPCIDR returns a validator that ensures the input is a valid IPv4, IPv6 address, or CIDR notation.
func IsValidIPCIDR() validator.String {
	return IPCIDRValidator{}
}
