package validator

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework-validators/helpers/validatordiag"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ validator.String = uddiDomainNameValidator{}

// uddiDomainNameValidator validates that a value is in the FQDN form expected by UDDI.
type uddiDomainNameValidator struct{}

func (v uddiDomainNameValidator) Description(ctx context.Context) string {
	return "value must be a fully qualified domain name ending with a trailing dot"
}

func (v uddiDomainNameValidator) MarkdownDescription(ctx context.Context) string {
	return v.Description(ctx)
}

func (v uddiDomainNameValidator) ValidateString(ctx context.Context, req validator.StringRequest, resp *validator.StringResponse) {
	if req.ConfigValue.IsNull() || req.ConfigValue.IsUnknown() {
		return
	}
	value := req.ConfigValue.ValueString()
	if !strings.HasSuffix(value, ".") {
		resp.Diagnostics.Append(validatordiag.InvalidAttributeValueDiagnostic(
			req.Path,
			fmt.Sprintf("must be a fully qualified domain name ending with a trailing dot, e.g. %q", value+"."),
			value,
		))
	}
}

// IsValidUDDIDomainName returns an AttributeValidator which ensures a configured FQDN is terminated by a dot.
func IsValidUDDIDomainName() validator.String {
	return uddiDomainNameValidator{}
}
