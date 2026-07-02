package validator

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

// ConditionalArpaOrFQDNValidator validates input based on the presence of "arpa".
type ConditionalArpaOrFQDNValidator struct {
	IPv4Validator validator.String
	IPv6Validator validator.String
	FQDNValidator validator.String
}

func (v ConditionalArpaOrFQDNValidator) Description(ctx context.Context) string {
	return "Validates input as either a valid IPv4/IPv6 ARPA address or a valid FQDN based on the presence of 'arpa'."
}

func (v ConditionalArpaOrFQDNValidator) MarkdownDescription(ctx context.Context) string {
	return v.Description(ctx)
}

func (v ConditionalArpaOrFQDNValidator) ValidateString(ctx context.Context, request validator.StringRequest, response *validator.StringResponse) {
	if request.ConfigValue.IsNull() || request.ConfigValue.IsUnknown() {
		return
	}

	value := request.ConfigValue.ValueString()
	lowerValue := strings.ToLower(value)

	if value == "" {
		response.Diagnostics.AddAttributeError(request.Path, "Invalid Value", "An empty value is not allowed for the zone name.")
	}

	if strings.HasSuffix(lowerValue, ".ip6.arpa") || strings.HasSuffix(lowerValue, ".in-addr.arpa") {
		// Run both ARPA validators and collect errors
		ipv4Resp := &validator.StringResponse{}
		ipv6Resp := &validator.StringResponse{}

		v.IPv4Validator.ValidateString(ctx, request, ipv4Resp)
		v.IPv6Validator.ValidateString(ctx, request, ipv6Resp)

		if ipv4Resp.Diagnostics.HasError() && ipv6Resp.Diagnostics.HasError() {
			// Both failed → add combined error
			response.Diagnostics.AddAttributeError(
				request.Path,
				"Invalid ARPA Format",
				fmt.Sprintf("The value '%s' must be a valid IPv4 ARPA (format: x.x.x.x.in-addr.arpa) or IPv6 ARPA (format: 32 hexadecimal digits separated by dots, ending with .ip6.arpa).", value),
			)
		}
	} else {
		// Validate as FQDN
		v.FQDNValidator.ValidateString(ctx, request, response)
	}
}

// IsValidArpaOrFQDN creates a new instance of the validator.
func IsValidArpaOrFQDN(ipv4, ipv6, fqdn validator.String) validator.String {
	return ConditionalArpaOrFQDNValidator{
		IPv4Validator: ipv4,
		IPv6Validator: ipv6,
		FQDNValidator: fqdn,
	}
}
