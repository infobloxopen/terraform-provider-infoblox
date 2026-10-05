package fw

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
)

// ValidateInternalDomainList validates the InternalDomainList configuration.
func ValidateInternalDomainList(ctx context.Context, data InternalDomainListModel, resp *resource.ValidateConfigResponse) {
	if uddi := flex.ExpandNestedObject[UDDIInternalDomainListModel](ctx, data.UDDI, &resp.Diagnostics); uddi != nil {
		validateInternalDomainListUDDIConfig(ctx, uddi, resp)
	}
}

func validateInternalDomainListUDDIConfig(ctx context.Context, m *UDDIInternalDomainListModel, resp *resource.ValidateConfigResponse) {
}
