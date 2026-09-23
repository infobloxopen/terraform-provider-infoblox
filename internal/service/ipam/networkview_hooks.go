package ipam

import (
	"context"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	coremodel "github.com/infobloxopen/terraform-provider-infoblox/internal/core/model/ipam"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
)

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

func PostExpandNetworkviewNIOS(ctx context.Context, ext *coremodel.NIOSNetworkviewExt, diags *diag.Diagnostics) *coremodel.NIOSNetworkviewExt {
	if ext == nil || ext.DdnsDnsView == nil || ext.Name == nil {
		return ext
	}
	suffix := "." + *ext.Name
	if !strings.HasSuffix(*ext.DdnsDnsView, suffix) {
		qualified := *ext.DdnsDnsView + suffix
		ext.DdnsDnsView = &qualified
	}
	return ext
}

func PostFlattenNetworkviewNIOS(ctx context.Context, planned, flattened *NIOSNetworkviewModel, diags *diag.Diagnostics) {
	if flattened.DdnsDnsView.IsNull() || flattened.Name.IsNull() {
		return
	}
	name := flattened.Name.ValueString()
	val := flattened.DdnsDnsView.ValueString()
	suffix := "." + name
	if !strings.HasSuffix(val, suffix) {
		return
	}
	stripped := strings.TrimSuffix(val, suffix)
	if planned == nil || planned.DdnsDnsView.IsNull() || planned.DdnsDnsView.IsUnknown() || planned.DdnsDnsView.ValueString() == stripped {
		flattened.DdnsDnsView = types.StringValue(stripped)
	}
}
