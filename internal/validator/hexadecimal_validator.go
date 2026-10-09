package validator

import (
	"context"
	"fmt"
	"regexp"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

// hexadecimalValidator validates if the provided value is a valid hexadecimal string.
type hexadecimalValidator struct{}

var hexPattern = regexp.MustCompile(`^[0-9a-fA-F]+$`)

func IsValidHexadecimal() validator.String {
	return hexadecimalValidator{}
}

func (v hexadecimalValidator) Description(ctx context.Context) string {
	return "Validator to check if a value is a valid hexadecimal string"
}

func (v hexadecimalValidator) MarkdownDescription(ctx context.Context) string {
	return v.Description(ctx)
}

func (v hexadecimalValidator) ValidateString(ctx context.Context, req validator.StringRequest, resp *validator.StringResponse) {
	if req.ConfigValue.IsNull() || req.ConfigValue.IsUnknown() {
		return
	}

	value := req.ConfigValue.ValueString()
	if value == "" {
		return
	}

	if !isValidHexadecimal(value) {
		resp.Diagnostics.AddAttributeError(
			req.Path,
			"Invalid Hexadecimal Format",
			fmt.Sprintf("Only hexadecimal digits are allowed. Value %q contains invalid characters. Supported format: 0-9, a-f, A-F (optional 0x or 0X prefix).", value),
		)
	}
}

// isValidHexadecimal checks if the string contains only hexadecimal characters
// Accepts optional 0x or 0X prefix
func isValidHexadecimal(value string) bool {
	value = strings.TrimSpace(value)

	// Remove optional 0x or 0X prefix
	if strings.HasPrefix(strings.ToLower(value), "0x") {
		value = value[2:]
	}

	// Must have at least one hex digit after removing prefix
	if value == "" {
		return false
	}

	return hexPattern.MatchString(value)
}
