package validator

import (
	"context"
	"fmt"
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

// ArpaIPv4Validator validates that the string matches a valid ARPA IPv4 format
type ArpaIPv4Validator struct{}

func (v ArpaIPv4Validator) Description(ctx context.Context) string {
	return "Validator to check if string is in the correct IPv4 ARPA format 'x.x.x.x.in-addr.arpa'"
}

func (v ArpaIPv4Validator) MarkdownDescription(ctx context.Context) string {
	return "Validator to check if string is in the correct IPv4 ARPA format 'x.x.x.x.in-addr.arpa'"
}

func (v ArpaIPv4Validator) ValidateString(ctx context.Context, request validator.StringRequest, response *validator.StringResponse) {
	if request.ConfigValue.IsNull() || request.ConfigValue.IsUnknown() {
		return
	}

	value := request.ConfigValue.ValueString()

	// Define the ARPA IPv4 regex
	re := regexp.MustCompile(`^((25[0-5]|2[0-4][0-9]|1[0-9]{2}|[1-9][0-9]?|0)(\.)){3}(25[0-5]|2[0-4][0-9]|1[0-9]{2}|[1-9][0-9]?|0)\.in-addr\.arpa$`)

	if !re.MatchString(value) {
		response.Diagnostics.AddAttributeError(
			request.Path,
			"Invalid ARPA IPv4 Format",
			fmt.Sprintf("The value '%s' is not a valid ARPA IPv4 address. Expected format: 'x.x.x.x.in-addr.arpa'", value),
		)
	}
}

// IsValidArpaIPv4 returns a validator that ensures the input is a valid ARPA IPv4 address.
func IsValidArpaIPv4() validator.String {
	return ArpaIPv4Validator{}
}
