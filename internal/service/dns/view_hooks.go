package dns

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
)

// ValidateView validates the View configuration.
func ValidateView(ctx context.Context, data ViewModel, resp *resource.ValidateConfigResponse) {
	if nios := flex.ExpandNestedObject[NIOSViewModel](ctx, data.NIOS, &resp.Diagnostics); nios != nil {
		validateViewNIOSConfig(ctx, nios, resp)
	}
	if uddi := flex.ExpandNestedObject[UDDIViewModel](ctx, data.UDDI, &resp.Diagnostics); uddi != nil {
		validateViewUDDIConfig(ctx, uddi, resp)
	}
}

func validateViewNIOSConfig(ctx context.Context, m *NIOSViewModel, resp *resource.ValidateConfigResponse) {
}

func validateViewUDDIConfig(ctx context.Context, m *UDDIViewModel, resp *resource.ValidateConfigResponse) {
}
