package ipam

import (
	"context"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
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

// PostFlattenNetworkviewNIOS normalizes fields that NIOS modifies on read-back.
//
// When ddns_dns_view is left unconfigured (null), NIOS auto-assigns the default
// DNS view for the network view and names it "default.<networkview_name>". Strip
// the ".<networkview_name>" suffix so Terraform sees "default", matching what
// NIOS does when the field is left unconfigured.
//
// When the user explicitly sets ddns_dns_view (non-null), NIOS stores and returns
// the exact value — no suffix is appended, so no stripping is needed.
func PostFlattenNetworkviewNIOS(ctx context.Context, planned, flattened *NIOSNetworkviewModel, diags *diag.Diagnostics) {
	if flattened == nil {
		return
	}
	if flattened.DdnsDnsView.IsNull() || flattened.DdnsDnsView.IsUnknown() {
		return
	}
	if flattened.Name.IsNull() || flattened.Name.IsUnknown() {
		return
	}
	// planned == nil means datasource context (no prior config). For resources,
	// only strip when the field was not explicitly set in config.
	if planned != nil && !planned.DdnsDnsView.IsNull() && !planned.DdnsDnsView.IsUnknown() {
		return
	}
	flattenedVal := flattened.DdnsDnsView.ValueString()
	name := flattened.Name.ValueString()
	if strings.HasSuffix(flattenedVal, "."+name) {
		flattened.DdnsDnsView = types.StringValue(strings.TrimSuffix(flattenedVal, "."+name))
	}
}
