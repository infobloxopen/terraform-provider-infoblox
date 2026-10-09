package validator

import (
	"context"
	"fmt"
	"regexp"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

// ArpaIPv6Validator validates that the string matches a valid ARPA IPv6 format
type ArpaIPv6Validator struct{}

func (v ArpaIPv6Validator) Description(ctx context.Context) string {
	return "Validator to check if string is in the correct IPv6 ARPA format 'x.x.x.x...x.ip6.arpa' with exactly 32 hex nibbles."
}

func (v ArpaIPv6Validator) MarkdownDescription(ctx context.Context) string {
	return "Validator to check if string is in the correct IPv6 ARPA format 'x.x.x.x...x.ip6.arpa' with exactly 32 hex nibbles."
}

func (v ArpaIPv6Validator) ValidateString(ctx context.Context, request validator.StringRequest, response *validator.StringResponse) {
	if request.ConfigValue.IsNull() || request.ConfigValue.IsUnknown() {
		return
	}

	value := request.ConfigValue.ValueString()

	// Check for whitespace
	if strings.TrimSpace(value) != value {
		response.Diagnostics.AddAttributeError(
			request.Path,
			"Invalid Whitespace",
			"The value must not contain leading or trailing whitespace.",
		)
		return
	}

	// Check for trailing dot
	if strings.HasSuffix(value, ".") {
		response.Diagnostics.AddAttributeError(
			request.Path,
			"Trailing Dot Not Allowed",
			"The value must not end with a dot.",
		)
		return
	}

	// Define the ARPA IPv6 regex
	re := regexp.MustCompile(`^([0-9a-f]\.){31}[0-9a-f]\.ip6\.arpa$`)

	if !re.MatchString(value) {
		response.Diagnostics.AddAttributeError(
			request.Path,
			"Invalid ARPA IPv6 Format",
			fmt.Sprintf("The value '%s' is not a valid ARPA IPv6 address. Expected format: 'x.x.x.x...x.ip6.arpa' with exactly 32 hex nibbles.", value),
		)
	}
}

// IsValidArpaIPv6 returns a validator that ensures the input is a valid ARPA IPv6 address.
func IsValidArpaIPv6() validator.String {
	return ArpaIPv6Validator{}
}
