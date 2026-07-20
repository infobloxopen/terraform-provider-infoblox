package validator

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/helpers/validatordiag"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ validator.String = stringNotNullValidator{}

// stringNotNullValidator validates that the value is not null.
// This is required for some fields that are "required" but cannot be marked required in the schema.
// See - https://github.com/hashicorp/terraform-plugin-framework/issues/898#issuecomment-1871470240
type stringNotNullValidator struct{}

func (s stringNotNullValidator) Description(ctx context.Context) string {
	return "string must not be null"
}

func (s stringNotNullValidator) MarkdownDescription(ctx context.Context) string {
	return "string must not be null"
}

func (s stringNotNullValidator) ValidateString(ctx context.Context, request validator.StringRequest, response *validator.StringResponse) {
	if request.ConfigValue.IsNull() {
		response.Diagnostics.Append(validatordiag.InvalidAttributeValueLengthDiagnostic(
			request.Path,
			s.Description(ctx),
			"null",
		))
	}
}

func StringNotNull() validator.String {
	return stringNotNullValidator{}
}
