package dns

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/path"
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
	// If any of the items in filter_aaaa_list has a non-null ref, then filter_aaaa must be set to YES or BREAK_DNSSEC.
	if m.FilterAaaaList.IsNull() || m.FilterAaaaList.IsUnknown() {
		return
	}

	var items []ViewFilterAaaaListModel
	resp.Diagnostics.Append(m.FilterAaaaList.ElementsAs(ctx, &items, false)...)
	if resp.Diagnostics.HasError() {
		return
	}

	hasRef := false
	for _, item := range items {
		if !item.Ref.IsNull() && !item.Ref.IsUnknown() && item.Ref.ValueString() != "" {
			hasRef = true
			break
		}
	}
	if !hasRef {
		return
	}

	// filter_aaaa may be set from a variable(tfvars), so during plan-time ValidateConfig
	// it can still be unknown. We only want to validate if it is known.
	if m.FilterAaaa.IsUnknown() {
		return
	}

	if m.FilterAaaa.IsNull() {
		resp.Diagnostics.AddAttributeError(
			path.Root("nios").AtName("filter_aaaa"),
			"Missing Filter AAAA Configuration",
			"When 'ref' field is provided in filter_aaaa_list, filter_aaaa must be set to 'YES' or 'BREAK_DNSSEC', it cannot be null or empty.",
		)
		return
	}

	if m.FilterAaaa.ValueString() == "NO" {
		resp.Diagnostics.AddAttributeError(
			path.Root("nios").AtName("filter_aaaa"),
			"Invalid Filter AAAA Configuration",
			"When 'ref' field is provided in filter_aaaa_list, filter_aaaa must be set to 'YES' or 'BREAK_DNSSEC', not 'NO'.",
		)
	}
}

func validateViewUDDIConfig(ctx context.Context, m *UDDIViewModel, resp *resource.ValidateConfigResponse) {
}
