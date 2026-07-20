package validator

import (
	"context"
	"fmt"
	"regexp"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

type ArpaIPv4RPZValidator struct{}

func (v ArpaIPv4RPZValidator) Description(ctx context.Context) string {
	return "Validator to check if string is in the RPZ IPv4 ARPA format 'x.x.x.x.in-addr.arpa.<rp-zone>'"
}

func (v ArpaIPv4RPZValidator) MarkdownDescription(ctx context.Context) string {
	return "Validator to check if string is in the RPZ IPv4 ARPA format 'x.x.x.x.in-addr.arpa.<rp-zone>'"
}

func (v ArpaIPv4RPZValidator) ValidateString(ctx context.Context, request validator.StringRequest, response *validator.StringResponse) {
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

	re := regexp.MustCompile(`^((25[0-5]|2[0-4][0-9]|1[0-9]{2}|[1-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|1[0-9]{2}|[1-9]?[0-9])\.in-addr\.arpa\.[A-Za-z0-9.-]+$`)

	if !re.MatchString(value) {
		response.Diagnostics.AddAttributeError(
			request.Path,
			"Invalid RPZ ARPA IPv4 Format",
			fmt.Sprintf("The value '%s' is not a valid RPZ ARPA IPv4 name. Expected format: 'x.x.x.x.in-addr.arpa.<rp-zone>' (example: '1.0.10.10.in-addr.arpa.rpz.example.com').", value),
		)
	}
}

// IsValidArpaIPv4RPZ returns a validator that ensures the input is a valid ARPA IPv4 RPZ address.
func IsValidArpaIPv4RPZ() validator.String {
	return ArpaIPv4RPZValidator{}
}
