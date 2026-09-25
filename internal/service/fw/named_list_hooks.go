package fw

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/utils"
)

// ValidateNamedList validates the NamedList configuration.
func ValidateNamedList(ctx context.Context, data NamedListModel, resp *resource.ValidateConfigResponse) {
	if uddi := flex.ExpandNestedObject[UDDINamedListModel](ctx, data.UDDI, &resp.Diagnostics); uddi != nil {
		validateNamedListUDDIConfig(ctx, uddi, resp)
	}
}

func validateNamedListUDDIConfig(ctx context.Context, m *UDDINamedListModel, resp *resource.ValidateConfigResponse) {
}

// PostFlattenNamedListUDDI reorders items_described to match the plan order,
// since the UDDI backend does not preserve submission order in its response.
func PostFlattenNamedListUDDI(ctx context.Context, planned, flattened *UDDINamedListModel, diags *diag.Diagnostics) {
	if flattened == nil || planned == nil {
		return
	}
	if !planned.ItemsDescribed.IsUnknown() {
		if reordered, d := utils.ReorderAndFilterNestedListResponse(ctx, planned.ItemsDescribed, flattened.ItemsDescribed, "item"); !d.HasError() {
			flattened.ItemsDescribed = reordered.(basetypes.ListValue)
		}
	}
}
