package validator

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

// NotArpaValidator validates that the provided string is NOT an ARPA format address
type NotArpaValidator struct{}

func (v NotArpaValidator) Description(ctx context.Context) string {
	return "String must not be an ARPA format address (*.in-addr.arpa or *.ip6.arpa)"
}

func (v NotArpaValidator) MarkdownDescription(ctx context.Context) string {
	return "String must not be an ARPA format address (*.in-addr.arpa or *.ip6.arpa)"
}

func (v NotArpaValidator) ValidateString(ctx context.Context, request validator.StringRequest, response *validator.StringResponse) {
	if request.ConfigValue.IsNull() || request.ConfigValue.IsUnknown() {
		return
	}

	value := strings.ToLower(request.ConfigValue.ValueString())

	// Check if it's ARPA format
	if strings.HasSuffix(value, ".in-addr.arpa") || strings.HasSuffix(value, ".ip6.arpa") {
		response.Diagnostics.AddAttributeError(
			request.Path,
			"ARPA Format Not Allowed",
			fmt.Sprintf("ARPA format addresses are not allowed: '%s'. Please use IP CIDR Notation.", request.ConfigValue.ValueString()),
		)
	}
}

// IsNotArpa returns a validator that ensures the input is NOT an ARPA format address.
func IsNotArpa() validator.String {
	return NotArpaValidator{}
}
