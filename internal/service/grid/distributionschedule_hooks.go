package grid

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	niosgrid "github.com/infobloxopen/infoblox-nios-go-client/grid"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/core"
	coremodel "github.com/infobloxopen/terraform-provider-infoblox/internal/core/model/grid"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/utils"
)

// ValidateDistributionschedule validates the Distributionschedule configuration.
func ValidateDistributionschedule(ctx context.Context, data DistributionscheduleModel, resp *resource.ValidateConfigResponse) {
	if nios := flex.ExpandNestedObject[NIOSDistributionscheduleModel](ctx, data.NIOS, &resp.Diagnostics); nios != nil {
		validateDistributionscheduleNIOSConfig(ctx, nios, resp)
	}
}

func validateDistributionscheduleNIOSConfig(ctx context.Context, m *NIOSDistributionscheduleModel, resp *resource.ValidateConfigResponse) {
	if m.UpgradeGroups.IsNull() || m.UpgradeGroups.IsUnknown() {
		return
	}
	var groups []DistributionscheduleUpgradeGroupsModel
	resp.Diagnostics.Append(m.UpgradeGroups.ElementsAs(ctx, &groups, false)...)
	if resp.Diagnostics.HasError() {
		return
	}

	groupsPath := path.Root("nios").AtName("upgrade_groups")
	for idx, group := range groups {
		if group.Name.IsUnknown() {
			continue
		}
		if group.Name.IsNull() {
			resp.Diagnostics.AddAttributeError(groupsPath.AtListIndex(idx).AtName("name"),
				"Invalid upgrade_groups.name",
				fmt.Sprintf("upgrade_groups[%d].name must be set", idx))
			continue
		}
		if group.Name.ValueString() == "Grid Master" {
			resp.Diagnostics.AddAttributeError(groupsPath.AtListIndex(idx).AtName("name"),
				"Invalid upgrade group",
				"\"Grid Master\" is not a valid upgrade group. It is upgraded as part of the \"Default\" group and cannot be scheduled explicitly.")
		}
	}
}

// lookupDistributionschedule finds the grid's singleton schedule and targets it for the Create-time Update.
func (r *DistributionscheduleResource) lookupDistributionschedule(ctx context.Context, data *DistributionscheduleModel, obj *coremodel.Distributionschedule, diags *diag.Diagnostics) {
	results, _, _, err := r.lookupService.List(ctx, &core.ListOptions{ReturnFields: DistributionscheduleReturnFields})
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to look up Distributionschedule: %s", err))
		return
	}
	if len(results) == 0 || results[0] == nil || results[0].Id == nil {
		diags.AddError("Client Error", "No Distributionschedule found on the grid")
		return
	}
	data.Id = types.StringPointerValue(results[0].Id)

	if obj != nil && obj.NIOS != nil && results[0].NIOS != nil {
		obj.NIOS.UpgradeGroups = mergeDistributionscheduleUpgradeGroups(results[0].NIOS.UpgradeGroups, obj.NIOS.UpgradeGroups)
	}
}

// preDistributionscheduleUpdate merges the configured upgrade groups into the grid's current
// list before the Update call, as the lookup hook does on Create.
func (r *DistributionscheduleResource) preDistributionscheduleUpdate(ctx context.Context, data *DistributionscheduleModel, obj *coremodel.Distributionschedule, diags *diag.Diagnostics) {
	if obj == nil || obj.NIOS == nil || len(obj.NIOS.UpgradeGroups) == 0 {
		return
	}
	current, _, err := r.service.Read(ctx, data.Id.ValueString(), &core.Options{ReturnFields: "upgrade_groups"})
	if err != nil {
		diags.AddError("Client Error", fmt.Sprintf("Unable to read current Distributionschedule for merge: %s", err))
		return
	}
	if current != nil && current.NIOS != nil {
		obj.NIOS.UpgradeGroups = mergeDistributionscheduleUpgradeGroups(current.NIOS.UpgradeGroups, obj.NIOS.UpgradeGroups)
	}
}

// mergeDistributionscheduleUpgradeGroups merges the grid's current groups with the configured ones:
// NIOS rejects a schedule that omits any existing group ("Missing upgrade groups"), so configured
// groups replace their grid counterpart by name and grid-only groups are kept as-is.
// Configuring no groups sends none, leaving the grid's schedule untouched.
func mergeDistributionscheduleUpgradeGroups(current, configured []niosgrid.DistributionscheduleUpgradeGroups) []niosgrid.DistributionscheduleUpgradeGroups {
	if len(configured) == 0 {
		return nil
	}
	byName := make(map[string]niosgrid.DistributionscheduleUpgradeGroups, len(configured))
	for _, g := range configured {
		if g.Name != nil {
			byName[*g.Name] = g
		}
	}

	// time_zone is read-only on the wire; drop it from every group sent back.
	merged := make([]niosgrid.DistributionscheduleUpgradeGroups, 0, len(current)+len(configured))
	seen := make(map[string]bool, len(current))
	for _, g := range current {
		name := ""
		if g.Name != nil {
			name = *g.Name
		}
		if cg, ok := byName[name]; ok {
			g = cg
		}
		g.TimeZone = nil
		merged = append(merged, g)
		seen[name] = true
	}
	for _, g := range configured {
		if g.Name == nil || !seen[*g.Name] {
			g.TimeZone = nil
			merged = append(merged, g)
		}
	}
	return merged
}

func PostExpandDistributionscheduleNIOS(ctx context.Context, ext *coremodel.NIOSDistributionscheduleExt, diags *diag.Diagnostics) *coremodel.NIOSDistributionscheduleExt {
	if ext == nil {
		return ext
	}
	// NIOS rejects a PUT whose upgrade_groups omits any existing group ("Missing upgrade groups"),
	// so when none are configured leave the field out entirely and keep the grid's current groups.
	if len(ext.UpgradeGroups) == 0 {
		ext.UpgradeGroups = nil
		return ext
	}
	// upgrade_time is computed-only but NIOS rejects an upgrade group without it; send 0 when unset.
	for i := range ext.UpgradeGroups {
		if ext.UpgradeGroups[i].UpgradeTime == nil {
			ext.UpgradeGroups[i].UpgradeTime = new(int64)
		}
	}
	return ext
}

func PostFlattenDistributionscheduleNIOS(ctx context.Context, planned, flattened *NIOSDistributionscheduleModel, diags *diag.Diagnostics) {
	if planned == nil || flattened == nil {
		return
	}

	if !planned.UpgradeGroups.IsUnknown() {
		if reordered, d := utils.ReorderAndFilterNestedListResponse(ctx, planned.UpgradeGroups, flattened.UpgradeGroups, "name"); !d.HasError() {
			if reorderedList, ok := reordered.(basetypes.ListValue); ok {
				flattened.UpgradeGroups = reorderedList
			}
		}
	}
}
