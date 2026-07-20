package validator

import (
	"context"
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework-validators/helpers/validatordiag"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ validator.String = passwordStrengthValidator{}

type passwordStrengthValidator struct{}

var (
	lowercaseRegex = regexp.MustCompile(`[a-z]`)
	uppercaseRegex = regexp.MustCompile(`[A-Z]`)
	digitRegex     = regexp.MustCompile(`[0-9]`)
	symbolRegex    = regexp.MustCompile(`[^a-zA-Z0-9]`)
)

func (v passwordStrengthValidator) Description(ctx context.Context) string {
	return "Password must be at least 4 characters long and include at least one lowercase letter, one uppercase letter, one digit, and one symbol."
}

func (v passwordStrengthValidator) MarkdownDescription(ctx context.Context) string {
	return v.Description(ctx)
}

func (v passwordStrengthValidator) ValidateString(ctx context.Context, req validator.StringRequest, resp *validator.StringResponse) {
	if req.ConfigValue.IsNull() || req.ConfigValue.IsUnknown() {
		return
	}

	value := req.ConfigValue.ValueString()

	if len(value) < 4 ||
		!lowercaseRegex.MatchString(value) ||
		!uppercaseRegex.MatchString(value) ||
		!digitRegex.MatchString(value) ||
		!symbolRegex.MatchString(value) {
		resp.Diagnostics.Append(validatordiag.InvalidAttributeValueDiagnostic(
			req.Path,
			"Password does not meet requirements. Password must contain at least 4 characters, including at least 1 lowercase character, 1 uppercase character, 1 numeric character, 1 symbol character.",
			value,
		))
	}
}

// IsStrongPassword returns a validator.String that enforces password strength rules.
func IsStrongPassword() validator.String {
	return passwordStrengthValidator{}
}
