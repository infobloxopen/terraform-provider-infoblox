package validator

import (
	"context"
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework-validators/helpers/validatordiag"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ validator.String = syslogEndpointNameValidator{}

// syslogEndpointNameValidator validates that the provided string does not contain spaces or special characters.
type syslogEndpointNameValidator struct{}

func (v syslogEndpointNameValidator) Description(ctx context.Context) string {
	return "value must not contain spaces or special characters"
}

func (v syslogEndpointNameValidator) MarkdownDescription(ctx context.Context) string {
	return v.Description(ctx)
}

func (v syslogEndpointNameValidator) ValidateString(ctx context.Context, req validator.StringRequest, resp *validator.StringResponse) {
	if req.ConfigValue.IsNull() || req.ConfigValue.IsUnknown() {
		return
	}

	value := req.ConfigValue.ValueString()

	// Regex pattern: ^[a-zA-Z0-9]*$
	// Only allows alphanumeric characters (letters and numbers)
	pattern := `^[a-zA-Z0-9]*$`
	matched, err := regexp.MatchString(pattern, value)
	if err != nil {
		resp.Diagnostics.Append(validatordiag.InvalidAttributeTypeDiagnostic(
			req.Path,
			"failed to validate syslog endpoint name pattern",
			value,
		))
		return
	}

	if !matched {
		resp.Diagnostics.Append(validatordiag.InvalidAttributeTypeDiagnostic(
			req.Path,
			"Space and special characters are not allowed in syslog endpoint name.",
			value,
		))
	}
}

// ValidateSyslogEndpointName returns a validator that ensures the string does not contain spaces or special characters.
// Only alphanumeric characters (letters and numbers) are allowed.
func ValidateSyslogEndpointName() validator.String {
	return syslogEndpointNameValidator{}
}
