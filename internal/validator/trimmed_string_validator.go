package validator

import (
	"context"
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework-validators/helpers/validatordiag"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ validator.String = trimmedStringValidator{}

// trimmedStringValidator validates that the provided string does not have leading or trailing whitespace.
type trimmedStringValidator struct{}

func (v trimmedStringValidator) Description(ctx context.Context) string {
	return "value must not have leading or trailing whitespace"
}

func (v trimmedStringValidator) MarkdownDescription(ctx context.Context) string {
	return v.Description(ctx)
}

func (v trimmedStringValidator) ValidateString(ctx context.Context, req validator.StringRequest, resp *validator.StringResponse) {
	if req.ConfigValue.IsNull() || req.ConfigValue.IsUnknown() {
		return
	}

	value := req.ConfigValue.ValueString()

	// Regex pattern: ^$|^\S(?:.*\S)?$
	// ^$ - matches empty string
	// ^\S(?:.*\S)?$ - matches non-whitespace at start, optionally any characters in middle, non-whitespace at end
	pattern := `^$|^\S(?:.*\S)?$`
	matched, err := regexp.MatchString(pattern, value)
	if err != nil {
		resp.Diagnostics.Append(validatordiag.InvalidAttributeTypeDiagnostic(
			req.Path,
			"failed to validate whitespace pattern",
			value,
		))
		return
	}

	if !matched {
		resp.Diagnostics.Append(validatordiag.InvalidAttributeTypeDiagnostic(
			req.Path,
			"must not have leading or trailing whitespace",
			value,
		))
	}
}

// ValidateTrimmedString returns a validator that ensures the string does not have leading or trailing whitespace.
// It allows empty strings and strings that start and end with non-whitespace characters.
func ValidateTrimmedString() validator.String {
	return trimmedStringValidator{}
}
