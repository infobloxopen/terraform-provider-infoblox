package dns

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
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
	niosPath := path.Root("nios")
	externalPrimariesSet := !m.ExternalPrimaries.IsNull() && !m.ExternalPrimaries.IsUnknown()

	// If external_primaries is set, at least one grid secondary is required.
	if externalPrimariesSet && m.GridSecondaries.IsNull() {
		resp.Diagnostics.AddAttributeError(
			niosPath.AtName("grid_secondaries"),
			"Missing Grid Secondary Server",
			"An NS group must contain at least one grid secondary server if an external primary server is specified.",
		)
	}

	// Multimaster validations only apply when is_multimaster is true.
	if m.IsMultimaster.IsNull() || m.IsMultimaster.IsUnknown() || !m.IsMultimaster.ValueBool() {
		return
	}

	if !m.GridPrimary.IsNull() && !m.GridPrimary.IsUnknown() && len(m.GridPrimary.Elements()) <= 1 {
		resp.Diagnostics.AddAttributeError(
			niosPath.AtName("grid_primary"),
			"Invalid Multimaster Configuration",
			"When 'is_multimaster' is set to true, 'grid_primary' must contain more than one entry.",
		)
	}

	if externalPrimariesSet && len(m.ExternalPrimaries.Elements()) <= 1 {
		resp.Diagnostics.AddAttributeError(
			niosPath.AtName("external_primaries"),
			"Invalid Multimaster Configuration",
			"When 'is_multimaster' is set to true, 'external_primaries' must contain more than one entry.",
		)
	}
}

func PostFlattenNsgroupNIOS(ctx context.Context, planned, flattened *NIOSNsgroupModel, diags *diag.Diagnostics) {
	if flattened == nil {
		return
	}

	if flattened.Comment.IsNull() {
		flattened.Comment = types.StringValue("")
	}

	if planned == nil {
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
