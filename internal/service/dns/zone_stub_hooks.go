package dns

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/utils"
)

// ValidateZoneStub validates the ZoneStub configuration.
func ValidateZoneStub(ctx context.Context, data ZoneStubModel, resp *resource.ValidateConfigResponse) {
	if nios := flex.ExpandNestedObject[NIOSZoneStubModel](ctx, data.NIOS, &resp.Diagnostics); nios != nil {
		validateZoneStubNIOSConfig(ctx, nios, resp)
	}
}

func validateZoneStubNIOSConfig(ctx context.Context, m *NIOSZoneStubModel, resp *resource.ValidateConfigResponse) {
}

func PostFlattenZoneStubNIOS(ctx context.Context, planned, flattened *NIOSZoneStubModel, diags *diag.Diagnostics) {
	if planned == nil || flattened == nil {
		return
	}

	// NIOS does not return 'stealth' for stub_msservers, so carry the planned value over to state
	if result, d := utils.CopyFieldFromPlanToRespList(ctx, planned.StubMsservers, flattened.StubMsservers, "stealth"); !d.HasError() {
		if resultList, ok := result.(basetypes.ListValue); ok {
			flattened.StubMsservers = resultList
		}
	}
}
