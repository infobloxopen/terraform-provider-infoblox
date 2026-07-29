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
	// ms_primaries is required when ms_sync_disabled is set.
	if !m.MsSyncDisabled.IsNull() && !m.MsSyncDisabled.IsUnknown() && m.MsSyncDisabled.ValueBool() && m.MsPrimaries.IsNull() {
		resp.Diagnostics.AddAttributeError(
			path.Root("nios").AtName("ms_primaries"),
			"Invalid Configuration",
			"'ms_primaries' must be provided when 'ms_sync_disabled' is set.",
		)
	}

	// SOA authority fields require a primary name server source.
	hasSoaEmail := !m.SoaEmail.IsNull() && !m.SoaEmail.IsUnknown()
	hasSoaSerialNumber := !m.SoaSerialNumber.IsNull() && !m.SoaSerialNumber.IsUnknown()
	hasMemberSoaMnames := !m.MemberSoaMnames.IsNull() && !m.MemberSoaMnames.IsUnknown()

	if hasSoaEmail || hasSoaSerialNumber || hasMemberSoaMnames {
		hasGridPrimary := !m.GridPrimary.IsNull() && !m.GridPrimary.IsUnknown()
		hasNsGroup := !m.NsGroup.IsNull() && !m.NsGroup.IsUnknown()
		isPrimaryOrNsGroupUnknown := m.GridPrimary.IsUnknown() || m.NsGroup.IsUnknown()

		if !hasGridPrimary && !hasNsGroup && !isPrimaryOrNsGroupUnknown {
			resp.Diagnostics.AddError(
				"Invalid Configuration",
				"When soa_email, soa_serial_number, or member_soa_mnames is specified, either grid_primary or ns_group must be provided.",
			)
		}
	}

	// Only one of the primary server attributes may be specified.
	var specifiedPrimaries []string
	if !m.GridPrimary.IsNull() && !m.GridPrimary.IsUnknown() {
		specifiedPrimaries = append(specifiedPrimaries, "grid_primary")
	}
	if !m.ExternalPrimaries.IsNull() && !m.ExternalPrimaries.IsUnknown() {
		specifiedPrimaries = append(specifiedPrimaries, "external_primaries")
	}
	if !m.MsPrimaries.IsNull() && !m.MsPrimaries.IsUnknown() {
		specifiedPrimaries = append(specifiedPrimaries, "ms_primaries")
	}

	if len(specifiedPrimaries) > 1 {
		resp.Diagnostics.AddError(
			"Conflicting Primary Servers",
			fmt.Sprintf(
				"Only one of grid_primary, external_primaries, or ms_primaries can be specified. Found: %s.",
				strings.Join(specifiedPrimaries, ", "),
			),
		)
		return
	}

	primaryUnknown := m.GridPrimary.IsUnknown() || m.ExternalPrimaries.IsUnknown() || m.MsPrimaries.IsUnknown()
	secondarySpecified := (!m.GridSecondaries.IsNull() && !m.GridSecondaries.IsUnknown()) ||
		(!m.ExternalSecondaries.IsNull() && !m.ExternalSecondaries.IsUnknown()) ||
		(!m.MsSecondaries.IsNull() && !m.MsSecondaries.IsUnknown())

	// A secondary server requires exactly one primary server.
	if secondarySpecified && !primaryUnknown && len(specifiedPrimaries) != 1 {
		resp.Diagnostics.AddError(
			"Secondary Server Requires Exactly One Primary Server",
			"When secondary servers (grid_secondaries, external_secondaries, or ms_secondaries) are specified, exactly one primary server (grid_primary, external_primaries, or ms_primaries) is required.",
		)
	}

	// ns_group can't be specified when any primary or secondary server is specified.
	if !m.NsGroup.IsNull() && !m.NsGroup.IsUnknown() && (len(specifiedPrimaries) > 0 || secondarySpecified) {
		resp.Diagnostics.AddAttributeError(
			path.Root("nios").AtName("ns_group"),
			"NS Group Not Allowed",
			"The ns_group attribute cannot be specified when any of primary (grid_primary, external_primaries, or ms_primaries) or secondary server (grid_secondaries, external_secondaries, or ms_secondaries) is specified. Please remove the ns_group attribute or the primary/secondary server attributes.",
		)
	}
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
