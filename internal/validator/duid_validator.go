package validator

import (
	"context"
	"fmt"
	"regexp"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

// duidValidator validates if the provided value is a valid DUID.
type duidValidator struct{}

func IsValidDUID() validator.String {
	return duidValidator{}
}

func (v duidValidator) Description(ctx context.Context) string {
	return "value must be a valid DUID format"
}

func (v duidValidator) MarkdownDescription(ctx context.Context) string {
	return v.Description(ctx)
}

func (v duidValidator) ValidateString(ctx context.Context, req validator.StringRequest, resp *validator.StringResponse) {
	if req.ConfigValue.IsNull() || req.ConfigValue.IsUnknown() {
		return
	}

	value := req.ConfigValue.ValueString()
	if value == "" {
		return
	}

	if !isValidDUIDFormat(value) {
		resp.Diagnostics.AddAttributeError(
			req.Path,
			"Invalid DUID Format",
			fmt.Sprintf("Value %q is not a valid DUID format. Must contain 2-130 octets (4-260 hex characters). Supported formats: 0123456789ab, 01:23:45:67:89:ab, 01-23-45-67-89-ab", value),
		)
	}
}

// isValidDUIDFormat checks if the DUID format is valid
func isValidDUIDFormat(address string) bool {
	if address == "" {
		return true
	}

	address = strings.TrimSpace(address)

	patterns := []*regexp.Regexp{
		// AABBCCDDEEFF
		regexp.MustCompile(`^[0-9a-fA-F]{4,260}$`),
		// AA:BB:CC:DD:EE:FF
		regexp.MustCompile(`^([0-9a-fA-F]{1,2}:){1,129}[0-9a-fA-F]{1,2}$`),
		// AA-BB-CC-DD-EE-FF
		regexp.MustCompile(`^([0-9a-fA-F]{1,2}-){1,129}[0-9a-fA-F]{1,2}$`),
	}

	for _, pattern := range patterns {
		if pattern.MatchString(address) {
			// Additional check for continuous format: must have even length
			if !strings.Contains(address, ":") && !strings.Contains(address, "-") {
				return len(address)%2 == 0
			}
			return true
		}
	}
	return false
}
