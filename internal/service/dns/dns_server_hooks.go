package dns

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
)

// ValidateDnsServer validates the DnsServer configuration.
func ValidateDnsServer(ctx context.Context, data DnsServerModel, resp *resource.ValidateConfigResponse) {
	if uddi := flex.ExpandNestedObject[UDDIDnsServerModel](ctx, data.UDDI, &resp.Diagnostics); uddi != nil {
		validateDnsServerUDDIConfig(ctx, uddi, resp)
	}
}

func validateDnsServerUDDIConfig(ctx context.Context, m *UDDIDnsServerModel, resp *resource.ValidateConfigResponse) {
}
