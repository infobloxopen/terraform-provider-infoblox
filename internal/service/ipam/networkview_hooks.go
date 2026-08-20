package ipam

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
)

// ValidateNetworkview validates the Networkview configuration.
func ValidateNetworkview(ctx context.Context, data NetworkviewModel, resp *resource.ValidateConfigResponse) {
	if nios := flex.ExpandNestedObject[NIOSNetworkviewModel](ctx, data.NIOS, &resp.Diagnostics); nios != nil {
		validateNetworkviewNIOSConfig(ctx, nios, resp)
	}
	if uddi := flex.ExpandNestedObject[UDDINetworkviewModel](ctx, data.UDDI, &resp.Diagnostics); uddi != nil {
		validateNetworkviewUDDIConfig(ctx, uddi, resp)
	}
}

func validateNetworkviewNIOSConfig(ctx context.Context, m *NIOSNetworkviewModel, resp *resource.ValidateConfigResponse) {
}

func validateNetworkviewUDDIConfig(ctx context.Context, m *UDDINetworkviewModel, resp *resource.ValidateConfigResponse) {
}

func PostFlattenNetworkviewNIOS(ctx context.Context, planned, flattened *NIOSNetworkviewModel, diags *diag.Diagnostics) {
}

func PostFlattenNetworkviewUDDI(ctx context.Context, planned, flattened *UDDINetworkviewModel, diags *diag.Diagnostics) {
	if flattened == nil || planned == nil {
		return
	}
	// FlattenFrameworkListNestedBlock returns ListNull for empty API responses.
	// If the plan had an explicit empty list, restore it to avoid Terraform
	// "provider produced inconsistent result" errors.
	if flattened.DhcpOptionsV6.IsNull() && !planned.DhcpOptionsV6.IsNull() && !planned.DhcpOptionsV6.IsUnknown() && len(planned.DhcpOptionsV6.Elements()) == 0 {
		flattened.DhcpOptionsV6 = planned.DhcpOptionsV6
	}
	if flattened.DhcpOptions.IsNull() && !planned.DhcpOptions.IsNull() && !planned.DhcpOptions.IsUnknown() && len(planned.DhcpOptions.Elements()) == 0 {
		flattened.DhcpOptions = planned.DhcpOptions
	}
	if flattened.DefaultRealms.IsNull() && !planned.DefaultRealms.IsNull() && !planned.DefaultRealms.IsUnknown() && len(planned.DefaultRealms.Elements()) == 0 {
		flattened.DefaultRealms = planned.DefaultRealms
	}
}
