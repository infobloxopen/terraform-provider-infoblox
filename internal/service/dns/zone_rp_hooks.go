package dns

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
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

// validateZoneRpNIOSConfig ports the primary/secondary name server rules from the
// legacy NIOS provider's ZoneRpResource.ValidateConfig.
//
// The legacy provider additionally rejects SOA timer fields when use_grid_zone_timer
// is false. That rule is not ported because the unified schema does not expose the
// use_* inheritance flags -- codegen drops them for every object (zone_auth included).
func validateZoneRpNIOSConfig(ctx context.Context, m *NIOSZoneRpModel, resp *resource.ValidateConfigResponse) {
	// Only one of the primary server attributes may be specified. The grid_primary
	// field validator (list_conflicts_with) already reports the conflict itself, so
	// bail out here instead of raising a second, duplicate diagnostic.
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

	// A secondary server requires exactly one primary server. This rule spans four
	// attributes, so it cannot be expressed as a single-field validator.
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
}

func primariesOrNone(p []string) string {
	if len(p) == 0 {
		return "none"
	}
	return strings.Join(p, ", ")
}

// PostFlattenZoneRpNIOS reconciles the create/update response with the plan.
//
// NIOS never echoes tsig_key_name back on external_primaries / external_secondaries,
// so the flattened list would drop a value the user configured. Copy it back from
// the plan. Same treatment as zone_auth, which shares these extserver structs.
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

	// NIOS returns the grid name servers in its own order, not the order they were
	// configured in, which surfaces as a per-index diff on name/stealth. Re-key the
	// response against the plan by server name.
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

	// On a FIREEYE zone NIOS always materialises the complete fireeye_alert_mapping
	// set -- one entry per alert type -- no matter what the user declared. Computed
	// cannot absorb that, because the user configured the parent object explicitly,
	// so the nested value in config wins over state and every later plan proposes
	// deleting the entries NIOS added. Narrow the response to what the user actually
	// declared. The undeclared entries still exist on the grid; they are simply not
	// Terraform-managed until they appear in the configuration.
	reconcileFireeyeAlertMapping(ctx, planned, flattened, diags)
}

func reconcileFireeyeAlertMapping(ctx context.Context, planned, flattened *NIOSZoneRpModel, diags *diag.Diagnostics) {
	if planned.FireeyeRuleMapping.IsNull() || planned.FireeyeRuleMapping.IsUnknown() {
		return
	}
	if flattened.FireeyeRuleMapping.IsNull() || flattened.FireeyeRuleMapping.IsUnknown() {
		return
	}

	plannedAlert, ok := planned.FireeyeRuleMapping.Attributes()["fireeye_alert_mapping"]
	if !ok {
		return
	}

	var reconciled attr.Value
	if plannedAlert.IsNull() || plannedAlert.IsUnknown() {
		// Null on the read path (prior state), unknown on create/update because the
		// parent object is Optional+Computed. Both mean "the user declared none".
		reconciled = types.ListNull(
			types.ObjectType{AttrTypes: ZonerpfireeyerulemappingFireeyeAlertMappingAttrTypes},
		)
	} else {
		// The user declared a subset: keep those entries, in their configured order.
		filtered, d := utils.ReorderAndFilterNestedListResponse(
			ctx, plannedAlert, flattened.FireeyeRuleMapping.Attributes()["fireeye_alert_mapping"], "alert_type")
		if d.HasError() {
			diags.Append(*d...)
			return
		}
		reconciled = filtered
	}

	attrs := make(map[string]attr.Value, len(flattened.FireeyeRuleMapping.Attributes()))
	for k, v := range flattened.FireeyeRuleMapping.Attributes() {
		attrs[k] = v
	}
	attrs["fireeye_alert_mapping"] = reconciled

	obj, d := types.ObjectValue(ZoneRpFireeyeRuleMappingAttrTypes, attrs)
	if d.HasError() {
		diags.Append(d...)
		return
	}
	flattened.FireeyeRuleMapping = obj
}
