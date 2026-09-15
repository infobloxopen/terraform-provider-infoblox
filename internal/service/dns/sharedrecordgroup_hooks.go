package dns

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
)

// ValidateSharedrecordgroup validates the Sharedrecordgroup configuration.
func ValidateSharedrecordgroup(ctx context.Context, data SharedrecordgroupModel, resp *resource.ValidateConfigResponse) {
	if nios := flex.ExpandNestedObject[NIOSSharedrecordgroupModel](ctx, data.NIOS, &resp.Diagnostics); nios != nil {
		validateSharedrecordgroupNIOSConfig(ctx, nios, resp)
	}
}

func validateSharedrecordgroupNIOSConfig(ctx context.Context, m *NIOSSharedrecordgroupModel, resp *resource.ValidateConfigResponse) {
	// fqdn is required for every zone_associations entry.
	if m.ZoneAssociations.IsNull() || m.ZoneAssociations.IsUnknown() {
		return
	}

	var zoneAssociations []SharedrecordgroupZoneAssociationsModel
	resp.Diagnostics.Append(m.ZoneAssociations.ElementsAs(ctx, &zoneAssociations, false)...)
	if resp.Diagnostics.HasError() {
		return
	}

	for i, zoneAssociation := range zoneAssociations {
		if zoneAssociation.Fqdn.IsUnknown() {
			continue
		}
		if zoneAssociation.Fqdn.IsNull() || zoneAssociation.Fqdn.ValueString() == "" {
			resp.Diagnostics.AddAttributeError(
				path.Root("nios").AtName("zone_associations").AtListIndex(i).AtName("fqdn"),
				"Invalid Configuration",
				fmt.Sprintf("The 'fqdn' attribute is required for each item in 'zone_associations'. Please provide a valid FQDN for item index %d.", i),
			)
		}
	}
}
