package validator

import (
	"context"
	"fmt"
	"net"

	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ validator.String = ipOrFQDNValidator{}

// ipOrFQDNValidator validates that the value is either a valid IP (IPv4/IPv6) or FQDN.
type ipOrFQDNValidator struct{}

func (v ipOrFQDNValidator) Description(ctx context.Context) string {
	return "value must be a valid IPv4 address, IPv6 address, or FQDN"
}

func (v ipOrFQDNValidator) MarkdownDescription(ctx context.Context) string {
	return v.Description(ctx)
}

func (v ipOrFQDNValidator) ValidateString(ctx context.Context, req validator.StringRequest, resp *validator.StringResponse) {

	if req.ConfigValue.IsNull() || req.ConfigValue.IsUnknown() {
		return
	}

	value := req.ConfigValue.ValueString()

	// Check if valid IP
	if net.ParseIP(value) != nil {
		// valid IP address
		return
	}

	// Check if valid FQDN
	fqdnValidator := domainNameValidator{
		isMultiLabel: true,
	}

	fqdnResp := &validator.StringResponse{}
	fqdnValidator.ValidateString(ctx, req, fqdnResp)

	if !fqdnResp.Diagnostics.HasError() {
		// valid FQDN
		return
	}

	// If it is neither a valid IP nor a valid FQDN, return error
	resp.Diagnostics.AddAttributeError(
		req.Path,
		"Invalid Address Value",
		fmt.Sprintf(
			"Value %q is not a valid IP address or FQDN.",
			value,
		),
	)
}

func IsValidIPOrFQDN() validator.String {
	return ipOrFQDNValidator{}
}
