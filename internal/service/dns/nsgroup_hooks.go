package dns

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/utils"
)

// ValidateNsgroup validates the Nsgroup configuration.
func ValidateNsgroup(ctx context.Context, data NsgroupModel, resp *resource.ValidateConfigResponse) {
	if nios := flex.ExpandNestedObject[NIOSNsgroupModel](ctx, data.NIOS, &resp.Diagnostics); nios != nil {
		validateNsgroupNIOSConfig(ctx, nios, resp)
	}
}

func validateNsgroupNIOSConfig(ctx context.Context, m *NIOSNsgroupModel, resp *resource.ValidateConfigResponse) {
}

// PostFlattenNsgroupNIOS restores tsig_key_name on the external server lists.
// NIOS does not echo tsig_key_name in its response, so the flattened value is
// null even when the user configured it. Copy it back from the plan to keep the
// applied state consistent with the configuration.
func PostFlattenNsgroupNIOS(ctx context.Context, planned, flattened *NIOSNsgroupModel, diags *diag.Diagnostics) {
	if planned == nil || flattened == nil {
		return
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
}
