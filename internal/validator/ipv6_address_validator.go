package validator

import (
	"context"
	"fmt"
	"net/netip"

	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

// IPv6AddressValidator validates that the provided string is a valid IPv6 address.
type IPv6AddressValidator struct{}

func (v IPv6AddressValidator) Description(ctx context.Context) string {
	return "String must be a valid IPv6 address"
}

func (v IPv6AddressValidator) MarkdownDescription(ctx context.Context) string {
	return "String must be a valid IPv6 address"
}

func (v IPv6AddressValidator) ValidateString(ctx context.Context, request validator.StringRequest, response *validator.StringResponse) {
	if request.ConfigValue.IsNull() || request.ConfigValue.IsUnknown() {
		return
	}

	value := request.ConfigValue.ValueString()

	// Check if the value is a valid IPv6 address
	if ip, err := netip.ParseAddr(value); err == nil && ip.Is6() {
		return
	}

	// Add error if it is not a valid IPv6 address
	response.Diagnostics.AddAttributeError(
		request.Path,
		"Invalid IPv6 Address",
		fmt.Sprintf("Invalid value for %s. The value '%s' is not a valid IPv6 address.", request.Path, value),
	)
}

// IsValidIPv6Address returns a validator that ensures the input is a valid IPv6 address.
func IsValidIPv6Address() validator.String {
	return IPv6AddressValidator{}
}
