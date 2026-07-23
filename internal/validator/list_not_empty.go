package validator

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ validator.List = listNotEmptyValidator{}

type listNotEmptyValidator struct{}

func (v listNotEmptyValidator) Description(ctx context.Context) string {
	return "Validates that a list is not empty"
}

func (v listNotEmptyValidator) MarkdownDescription(ctx context.Context) string {
	return "Validates that a list is not empty"
}

func (v listNotEmptyValidator) ValidateList(ctx context.Context, request validator.ListRequest, response *validator.ListResponse) {
	if !request.ConfigValue.IsNull() && !request.ConfigValue.IsUnknown() && len(request.ConfigValue.Elements()) == 0 {
		response.Diagnostics.AddAttributeError(
			request.Path,
			"Empty Value Not Allowed",
			"This attribute must not be set to an empty list ([]). "+
				"If you want to unset this value, remove the attribute from your configuration instead.",
		)
	}
}

func ListNotEmpty() validator.List {
	return listNotEmptyValidator{}
}
