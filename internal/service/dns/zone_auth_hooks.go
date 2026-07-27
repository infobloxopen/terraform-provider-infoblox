package dns

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/utils"
)

// ValidateZoneAuth validates the ZoneAuth configuration.
func ValidateZoneAuth(ctx context.Context, data ZoneAuthModel, resp *resource.ValidateConfigResponse) {
	if nios := flex.ExpandNestedObject[NIOSZoneAuthModel](ctx, data.NIOS, &resp.Diagnostics); nios != nil {
		validateZoneAuthNIOSConfig(ctx, nios, resp)
	}
	if uddi := flex.ExpandNestedObject[UDDIZoneAuthModel](ctx, data.UDDI, &resp.Diagnostics); uddi != nil {
		validateZoneAuthUDDIConfig(ctx, uddi, resp)
	}
}

func validateZoneAuthNIOSConfig(ctx context.Context, m *NIOSZoneAuthModel, resp *resource.ValidateConfigResponse) {
}

func validateZoneAuthUDDIConfig(ctx context.Context, m *UDDIZoneAuthModel, resp *resource.ValidateConfigResponse) {
}

func PostFlattenZoneAuthNIOS(ctx context.Context, planned, flattened *NIOSZoneAuthModel, diags *diag.Diagnostics) {
	if planned == nil || flattened == nil {
		return
	}

	if !planned.AllowQuery.IsNull() {
		if result, d := utils.CopyFieldFromPlanToRespList(ctx, planned.AllowQuery, flattened.AllowQuery, "use_tsig_key_name"); !d.HasError() {
			flattened.AllowQuery = result.(basetypes.ListValue)
		}
	}
	if !planned.AllowTransfer.IsNull() {
		if result, d := utils.CopyFieldFromPlanToRespList(ctx, planned.AllowTransfer, flattened.AllowTransfer, "use_tsig_key_name"); !d.HasError() {
			flattened.AllowTransfer = result.(basetypes.ListValue)
		}
	}
	if !planned.AllowUpdate.IsNull() {
		if result, d := utils.CopyFieldFromPlanToRespList(ctx, planned.AllowUpdate, flattened.AllowUpdate, "use_tsig_key_name"); !d.HasError() {
			flattened.AllowUpdate = result.(basetypes.ListValue)
		}
	}
	if !planned.ExternalPrimaries.IsNull() {
		if result, d := utils.CopyFieldFromPlanToRespList(ctx, planned.ExternalPrimaries, flattened.ExternalPrimaries, "tsig_key_name"); !d.HasError() {
			flattened.ExternalPrimaries = result.(basetypes.ListValue)
		}
	}
	if !planned.ExternalSecondaries.IsNull() {
		if result, d := utils.CopyFieldFromPlanToRespList(ctx, planned.ExternalSecondaries, flattened.ExternalSecondaries, "tsig_key_name"); !d.HasError() {
			flattened.ExternalSecondaries = result.(basetypes.ListValue)
		}
	}

	if !planned.GridPrimary.IsUnknown() {
		if reordered, d := utils.ReorderAndFilterNestedListResponse(ctx, planned.GridPrimary, flattened.GridPrimary, "name"); !d.HasError() {
			flattened.GridPrimary = reordered.(basetypes.ListValue)
		}
	}
	if !planned.GridSecondaries.IsUnknown() {
		if reordered, d := utils.ReorderAndFilterNestedListResponse(ctx, planned.GridSecondaries, flattened.GridSecondaries, "name"); !d.HasError() {
			flattened.GridSecondaries = reordered.(basetypes.ListValue)
		}
	}
}
