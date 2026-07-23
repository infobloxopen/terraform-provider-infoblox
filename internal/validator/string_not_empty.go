package validator

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ validator.String = stringNotEmptyValidator{}

type stringNotEmptyValidator struct{}

func (s stringNotEmptyValidator) Description(ctx context.Context) string {
	return "Validates that a string is not empty"
}

func (s stringNotEmptyValidator) MarkdownDescription(ctx context.Context) string {
	return "Validates that a string is not empty"
}

func (s stringNotEmptyValidator) ValidateString(ctx context.Context, request validator.StringRequest, response *validator.StringResponse) {
	if !request.ConfigValue.IsNull() && !request.ConfigValue.IsUnknown() && request.ConfigValue.ValueString() == "" {
		response.Diagnostics.AddAttributeError(
			request.Path,
			"Empty Value Not Allowed",
			"This attribute must not be set to an empty string (\"\"). "+
				"If you want to unset this value, remove the attribute from your configuration instead.",
		)
	}
}

func StringNotEmpty() validator.String {
	return stringNotEmptyValidator{}
}
