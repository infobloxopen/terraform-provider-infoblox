package validator

import (
	"context"
	"fmt"
	"regexp"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

// macAddressValidator validates if the provided value is a valid MAC Address.
type macAddressValidator struct{}

func IsValidMacAddress() validator.String {
	return macAddressValidator{}
}

func (v macAddressValidator) Description(ctx context.Context) string {
	return "value must be a valid MAC address format"
}

func (v macAddressValidator) MarkdownDescription(ctx context.Context) string {
	return v.Description(ctx)
}

func (v macAddressValidator) ValidateString(ctx context.Context, req validator.StringRequest, resp *validator.StringResponse) {
	if req.ConfigValue.IsNull() || req.ConfigValue.IsUnknown() {
		return
	}

	value := req.ConfigValue.ValueString()
	if value == "" {
		return
	}

	if !isValidMACFormat(value) {
		resp.Diagnostics.AddAttributeError(
			req.Path,
			"Invalid MAC Address Format",
			fmt.Sprintf("Value %q is not a valid MAC address format. Supported formats: AA:BB:CC:DD:EE:FF, AA-BB-CC-DD-EE-FF, AABB.CCDD.EEFF, AABBCC-DDEEFF, AABBCCDDEEFF", value),
		)
	}
}

// isValidMACFormat checks if the MAC address format is valid
func isValidMACFormat(address string) bool {
	if address == "" {
		return true
	}

	address = strings.TrimSpace(address)
	patterns := []*regexp.Regexp{
		// AA:BB:CC:DD:EE:FF
		regexp.MustCompile(`^[0-9a-fA-F]{1,2}:[0-9a-fA-F]{1,2}:[0-9a-fA-F]{1,2}:[0-9a-fA-F]{1,2}:[0-9a-fA-F]{1,2}:[0-9a-fA-F]{1,2}$`),
		// AA-BB-CC-DD-EE-FF
		regexp.MustCompile(`^[0-9a-fA-F]{1,2}-[0-9a-fA-F]{1,2}-[0-9a-fA-F]{1,2}-[0-9a-fA-F]{1,2}-[0-9a-fA-F]{1,2}-[0-9a-fA-F]{1,2}$`),
		// AABB.CCDD.EEFF
		regexp.MustCompile(`^[0-9a-fA-F]{1,4}\.[0-9a-fA-F]{1,4}\.[0-9a-fA-F]{1,4}$`),
		// AABBCC-DDEEFF
		regexp.MustCompile(`^[0-9a-fA-F]{1,6}-[0-9a-fA-F]{1,6}$`),
		// AABBCCDDEEFF
		regexp.MustCompile(`^[0-9a-fA-F]{12}$`),
	}

	for _, pattern := range patterns {
		if pattern.MatchString(address) {
			return true
		}
	}
	return false
}
