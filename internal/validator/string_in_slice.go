package validator

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure the validator implements the List interface
var _ validator.List = &stringsInSliceValidator{}

// stringsInSliceValidator checks that all elements in a list are in a predefined set
type stringsInSliceValidator struct {
	allowedValues map[string]struct{}
	description   string
}

// Constructor function
func StringsInSlice(allowed []string) validator.List {
	allowedMap := make(map[string]struct{}, len(allowed))
	for _, v := range allowed {
		allowedMap[v] = struct{}{}
	}

	return &stringsInSliceValidator{
		allowedValues: allowedMap,
		description:   fmt.Sprintf("Each value must be one of: %v", allowed),
	}
}

func (v *stringsInSliceValidator) Description(_ context.Context) string {
	return v.description
}

func (v *stringsInSliceValidator) MarkdownDescription(_ context.Context) string {
	return v.description
}

func (v *stringsInSliceValidator) ValidateList(ctx context.Context, req validator.ListRequest, resp *validator.ListResponse) {
	if req.ConfigValue.IsNull() || req.ConfigValue.IsUnknown() {
		return
	}

	for i, elem := range req.ConfigValue.Elements() {
		strElem, ok := elem.(types.String)
		if !ok || strElem.IsNull() || strElem.IsUnknown() {
			continue
		}

		if _, valid := v.allowedValues[strElem.ValueString()]; !valid {
			resp.Diagnostics.AddAttributeError(
				req.Path,
				"Invalid Value",
				fmt.Sprintf("Element %d has invalid value %q. Allowed values are: %v.", i, strElem.ValueString(), v.description),
			)
		}
	}
}
