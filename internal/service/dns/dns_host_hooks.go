package dns

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
)

// ValidateDnsHost validates the DnsHost configuration.
func ValidateDnsHost(ctx context.Context, data DnsHostModel, resp *resource.ValidateConfigResponse) {
	if uddi := flex.ExpandNestedObject[UDDIDnsHostModel](ctx, data.UDDI, &resp.Diagnostics); uddi != nil {
		validateDnsHostUDDIConfig(ctx, uddi, resp)
	}
}

func validateDnsHostUDDIConfig(ctx context.Context, m *UDDIDnsHostModel, resp *resource.ValidateConfigResponse) {
}
