package keys

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
)

// ValidateKerberosKey validates the KerberosKey configuration.
func ValidateKerberosKey(ctx context.Context, data KerberosKeyModel, resp *resource.ValidateConfigResponse) {
	if uddi := flex.ExpandNestedObject[UDDIKerberosKeyModel](ctx, data.UDDI, &resp.Diagnostics); uddi != nil {
		validateKerberosKeyUDDIConfig(ctx, uddi, resp)
	}
}

func validateKerberosKeyUDDIConfig(ctx context.Context, m *UDDIKerberosKeyModel, resp *resource.ValidateConfigResponse) {
}
