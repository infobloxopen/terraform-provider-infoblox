package validator

import (
	"context"
	"fmt"
	"regexp"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

type ArpaIPv6RPZValidator struct{}

func (v ArpaIPv6RPZValidator) Description(ctx context.Context) string {
	return "Validator to check if string is in the RPZ IPv6 ARPA format: 'x.x.x...x.ip6.arpa.<rp-zone>'."
}

func (v ArpaIPv6RPZValidator) MarkdownDescription(ctx context.Context) string {
	return "Validator for IPv6 ARPA + RPZ domain: 'x.x.x...x.ip6.arpa.<rp-zone>'."
}

func (v ArpaIPv6RPZValidator) ValidateString(ctx context.Context, request validator.StringRequest, response *validator.StringResponse) {
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

	re := regexp.MustCompile(`(?i)^([0-9a-f]\.){31}[0-9a-f]\.ip6\.arpa\.[A-Za-z0-9.-]+$`)

	if !re.MatchString(value) {
		response.Diagnostics.AddAttributeError(
			request.Path,
			"Invalid RPZ ARPA IPv6 Format",
			fmt.Sprintf("The value '%s' must be in RPZ ARPA IPv6 form: 'x.x.x...x.ip6.arpa.<rp-zone>'", value),
		)
	}
}

// IsValidArpaIPv6RPZ returns a validator that ensures the input is a valid ARPA IPv6 RPZ address.
func IsValidArpaIPv6RPZ() validator.String {
	return ArpaIPv6RPZValidator{}
}
