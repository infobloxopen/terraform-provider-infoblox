package grid

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
	internaltypes "github.com/infobloxopen/terraform-provider-infoblox/internal/types"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/utils"
)

// ValidateServicerestartGroup validates the ServicerestartGroup configuration.
func ValidateServicerestartGroup(ctx context.Context, data ServicerestartGroupModel, resp *resource.ValidateConfigResponse) {
	if nios := flex.ExpandNestedObject[NIOSServicerestartGroupModel](ctx, data.NIOS, &resp.Diagnostics); nios != nil {
		validateServicerestartGroupNIOSConfig(ctx, nios, resp)
	}
}

func validateServicerestartGroupNIOSConfig(ctx context.Context, m *NIOSServicerestartGroupModel, resp *resource.ValidateConfigResponse) {
	if m.RecurringSchedule.IsNull() || m.RecurringSchedule.IsUnknown() {
		return
	}

	recurringScheduleAttrs := m.RecurringSchedule.Attributes()

	// If both DHCP and DNS are selected in services without every other service,
	// the caller must use "ALL" instead of enumerating them individually.
	if servicesAttr, ok := recurringScheduleAttrs["services"]; ok && !servicesAttr.IsNull() && !servicesAttr.IsUnknown() {
		servicesList, ok := servicesAttr.(internaltypes.UnorderedListValue)
		if !ok {
			resp.Diagnostics.AddAttributeError(
				path.Root("nios").AtName("recurring_schedule").AtName("services"),
				"Invalid Services Attribute",
				"Expected services to be a list but got a different type",
			)
			return
		}

		hasDHCP, hasDNS := false, false
		for _, v := range servicesList.Elements() {
			service, ok := v.(types.String)
			if !ok {
				resp.Diagnostics.AddAttributeError(
					path.Root("nios").AtName("recurring_schedule").AtName("services"),
					"Invalid Service Value",
					"Expected service value to be a string but got a different type",
				)
				return
			}
			switch service.ValueString() {
			case "DHCP":
				hasDHCP = true
			case "DNS":
				hasDNS = true
			}
			if hasDHCP && hasDNS {
				resp.Diagnostics.AddAttributeError(
					path.Root("nios").AtName("recurring_schedule").AtName("services"),
					"Invalid Services Configuration",
					"If both DHCP and DNS are selected in services, then services must be set to ALL",
				)
				return
			}
		}
	}

	utils.ValidateScheduleConfig(
		m.RecurringSchedule,
		"schedule",
		path.Root("nios").AtName("recurring_schedule"),
		&resp.Diagnostics,
	)
}
