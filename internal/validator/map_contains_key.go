package validator

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ validator.Map = mapContainsKeyValidator{}

type mapContainsKeyValidator struct {
	requiredKey string
}

func (v mapContainsKeyValidator) Description(_ context.Context) string {
	return fmt.Sprintf("map must contain the key '%s'", v.requiredKey)
}

func (v mapContainsKeyValidator) MarkdownDescription(_ context.Context) string {
	return fmt.Sprintf("This map must contain the key '%s'", v.requiredKey)
}

func (v mapContainsKeyValidator) ValidateMap(ctx context.Context, req validator.MapRequest, resp *validator.MapResponse) {
	if req.ConfigValue.IsNull() || req.ConfigValue.IsUnknown() {
		return
	}

	if _, ok := req.ConfigValue.Elements()[v.requiredKey]; !ok {
		resp.Diagnostics.AddAttributeError(
			req.Path,
			"Missing Required Key",
			fmt.Sprintf("The map is missing the required key '%s'.", v.requiredKey),
		)
	}
}

// MapContainsKey returns a validator.Map that ensures a map contains the specified key.
func MapContainsKey(requiredKey string) validator.Map {
	return mapContainsKeyValidator{
		requiredKey: requiredKey,
	}
}
