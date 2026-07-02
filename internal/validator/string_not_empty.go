package validator

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/helpers/validatordiag"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ validator.String = stringNotEmptyValidator{}

// stringNotEmptyValidator validates that the value is not null.
type stringNotEmptyValidator struct{}

func (s stringNotEmptyValidator) Description(ctx context.Context) string {
	return "Validates that a string is not empty"
}

func (s stringNotEmptyValidator) MarkdownDescription(ctx context.Context) string {
	return "Validates that a string is not empty"
}

func (s stringNotEmptyValidator) ValidateString(ctx context.Context, request validator.StringRequest, response *validator.StringResponse) {
	if !request.ConfigValue.IsNull() && !request.ConfigValue.IsUnknown() && request.ConfigValue.ValueString() == "" {
		response.Diagnostics.Append(validatordiag.InvalidAttributeValueDiagnostic(
			request.Path,
			"must not be empty",
			"\"\"",
		))
	}
}

func StringNotEmpty() validator.String {
	return stringNotEmptyValidator{}
}
