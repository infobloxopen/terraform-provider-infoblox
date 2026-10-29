package dns

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/utils"
)

// ValidateZoneRp validates the ZoneRp configuration.
func ValidateZoneRp(ctx context.Context, data ZoneRpModel, resp *resource.ValidateConfigResponse) {
	if nios := flex.ExpandNestedObject[NIOSZoneRpModel](ctx, data.NIOS, &resp.Diagnostics); nios != nil {
		validateZoneRpNIOSConfig(ctx, nios, resp)
	}
}

func validateZoneRpNIOSConfig(ctx context.Context, m *NIOSZoneRpModel, resp *resource.ValidateConfigResponse) {
	var specifiedPrimaries []string
	if !m.GridPrimary.IsNull() && !m.GridPrimary.IsUnknown() {
		specifiedPrimaries = append(specifiedPrimaries, "grid_primary")
	}
	if !m.ExternalPrimaries.IsNull() && !m.ExternalPrimaries.IsUnknown() {
		specifiedPrimaries = append(specifiedPrimaries, "external_primaries")
	}
	if len(specifiedPrimaries) > 1 {
		return
	}

	primaryUnknown := m.GridPrimary.IsUnknown() || m.ExternalPrimaries.IsUnknown()
	secondarySpecified := (!m.GridSecondaries.IsNull() && !m.GridSecondaries.IsUnknown()) ||
		(!m.ExternalSecondaries.IsNull() && !m.ExternalSecondaries.IsUnknown())

	if secondarySpecified && !primaryUnknown && len(specifiedPrimaries) != 1 {
		resp.Diagnostics.AddError(
			"Secondary Server Requires Exactly One Primary Server",
			fmt.Sprintf(
				"When secondary servers (grid_secondaries or external_secondaries) are specified, "+
					"exactly one primary server (grid_primary or external_primaries) is required. Found: %s.",
				primariesOrNone(specifiedPrimaries),
			),
		)
	}

	// NIOS discards SOA values for zones with an external primary or without a primary.
	hasSoaEmail := !m.SoaEmail.IsNull() && !m.SoaEmail.IsUnknown()
	hasSoaSerialNumber := !m.SoaSerialNumber.IsNull() && !m.SoaSerialNumber.IsUnknown()
	hasMemberSoaMnames := !m.MemberSoaMnames.IsNull() && !m.MemberSoaMnames.IsUnknown()
	hasSoaTimers := (!m.SoaDefaultTtl.IsNull() && !m.SoaDefaultTtl.IsUnknown()) ||
		(!m.SoaExpire.IsNull() && !m.SoaExpire.IsUnknown()) ||
		(!m.SoaNegativeTtl.IsNull() && !m.SoaNegativeTtl.IsUnknown()) ||
		(!m.SoaRefresh.IsNull() && !m.SoaRefresh.IsUnknown()) ||
		(!m.SoaRetry.IsNull() && !m.SoaRetry.IsUnknown())

	if hasSoaEmail || hasSoaSerialNumber || hasMemberSoaMnames || hasSoaTimers {
		hasGridPrimary := !m.GridPrimary.IsNull() && !m.GridPrimary.IsUnknown()
		hasNsGroup := !m.NsGroup.IsNull() && !m.NsGroup.IsUnknown()
		isPrimaryOrNsGroupUnknown := m.GridPrimary.IsUnknown() || m.NsGroup.IsUnknown()
		usesExternalPrimary := !m.ExternalPrimaries.IsNull() && !m.ExternalPrimaries.IsUnknown()

		switch {
		case usesExternalPrimary:
			resp.Diagnostics.AddError(
				"Invalid Configuration",
				"soa_email, soa_serial_number, member_soa_mnames, soa_default_ttl, soa_expire, soa_negative_ttl, soa_refresh, and soa_retry cannot be specified when external_primaries is specified.",
			)
		case !hasGridPrimary && !hasNsGroup && !isPrimaryOrNsGroupUnknown:
			resp.Diagnostics.AddError(
				"Invalid Configuration",
				"When soa_email, soa_serial_number, member_soa_mnames, soa_default_ttl, soa_expire, soa_negative_ttl, soa_refresh, or soa_retry is specified, either grid_primary or ns_group must be provided.",
			)
		}
	}

	// External secondaries don't count, NIOS needs a grid secondary to serve the RPZ.
	if !m.ExternalPrimaries.IsNull() && !m.ExternalPrimaries.IsUnknown() &&
		!m.GridSecondaries.IsUnknown() && m.GridSecondaries.IsNull() {
		resp.Diagnostics.AddAttributeError(
			path.Root("nios").AtName("grid_secondaries"),
			"Invalid Configuration",
			"At least one 'grid_secondaries' entry must be specified when 'external_primaries' is specified.",
		)
	}
}

func primariesOrNone(p []string) string {
	if len(p) == 0 {
		return "none"
	}
	return strings.Join(p, ", ")
}

func PostFlattenZoneRpNIOS(ctx context.Context, planned, flattened *NIOSZoneRpModel, diags *diag.Diagnostics) {
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
