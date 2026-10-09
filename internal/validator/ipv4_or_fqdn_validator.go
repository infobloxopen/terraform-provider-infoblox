package validator

import (
	"context"
	"fmt"
	"net"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework-validators/helpers/validatordiag"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ validator.String = iPv4OrFQDNValidator{}

// iPv4OrFQDNValidator validates if the provided value is a valid IPv4 address or a string without leading/trailing whitespace, trailing dot, or uppercase characters.
type iPv4OrFQDNValidator struct{}

func (v iPv4OrFQDNValidator) Description(ctx context.Context) string {
	return "value must be a valid IPv4 address or a string without leading/trailing whitespace, trailing dot, or uppercase characters"
}

func (v iPv4OrFQDNValidator) MarkdownDescription(ctx context.Context) string {
	return v.Description(ctx)
}

func (v iPv4OrFQDNValidator) ValidateString(ctx context.Context, req validator.StringRequest, resp *validator.StringResponse) {
	if req.ConfigValue.IsNull() || req.ConfigValue.IsUnknown() {
		return
	}

	value := req.ConfigValue.ValueString()

	if ip := net.ParseIP(value); ip != nil {
		return // valid IP address
	}

	// If it looks like an IP address but failed to parse, surface an error.
	if looksLikeIPv4(value) {
		resp.Diagnostics.Append(validatordiag.InvalidAttributeValueDiagnostic(
			req.Path,
			"Invalid IPv4 address",
			value,
		))
		return
	}

	validFQDN := domainNameValidator{
		isMultiLabel: true,
	}
	fqdnResp := &validator.StringResponse{}
	validFQDN.ValidateString(ctx, req, fqdnResp)

	if fqdnResp.Diagnostics.HasError() {
		resp.Diagnostics.Append(validatordiag.InvalidAttributeValueDiagnostic(
			req.Path,
			fmt.Sprintf("Invalid FQDN. Got Error : %s", fqdnResp.Diagnostics.Errors()[0].Detail()),
			value,
		))
	}
}

// looksLikeIPv4 returns true if the string resembles an IPv4 address (N.N.N.N with digits only)
func looksLikeIPv4(s string) bool {
	// Check if string contains at least one dot
	if !strings.Contains(s, ".") {
		return false
	}
	for i := 0; i < len(s); i++ {
		c := s[i]
		if (c < '0' || c > '9') && c != '.' {
			return false
		}
	}
	return true
}

func IsValidIPv4OrFQDN() validator.String {
	return iPv4OrFQDNValidator{}
}
