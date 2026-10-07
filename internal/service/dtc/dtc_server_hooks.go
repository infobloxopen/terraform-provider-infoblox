package dtc

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/utils"
)

// ValidateDtcServer validates the DtcServer configuration.
func ValidateDtcServer(ctx context.Context, data DtcServerModel, resp *resource.ValidateConfigResponse) {
	if nios := flex.ExpandNestedObject[NIOSDtcServerModel](ctx, data.NIOS, &resp.Diagnostics); nios != nil {
		validateDtcServerNIOSConfig(ctx, nios, resp)
	}
	if uddi := flex.ExpandNestedObject[UDDIDtcServerModel](ctx, data.UDDI, &resp.Diagnostics); uddi != nil {
		validateDtcServerUDDIConfig(ctx, uddi, resp)
	}
}

func validateDtcServerNIOSConfig(ctx context.Context, m *NIOSDtcServerModel, resp *resource.ValidateConfigResponse) {
}

func validateDtcServerUDDIConfig(ctx context.Context, m *UDDIDtcServerModel, resp *resource.ValidateConfigResponse) {
}

// PostFlattenDtcServerUDDI reorders the records list to match the plan order so
// that the API returning records in a different sequence does not produce a spurious diff.
func PostFlattenDtcServerUDDI(ctx context.Context, planned, flattened *UDDIDtcServerModel, diags *diag.Diagnostics) {
	if planned == nil || flattened == nil {
		return
	}

	if !planned.Records.IsUnknown() {
		if reordered, d := utils.ReorderAndFilterNestedListResponse(ctx, planned.Records, flattened.Records, "type"); !d.HasError() {
			if reorderedList, ok := reordered.(basetypes.ListValue); ok {
				flattened.Records = reorderedList
			}
		}
	}
}
