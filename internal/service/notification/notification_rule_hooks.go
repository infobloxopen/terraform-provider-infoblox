package notification

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/flex"
	"github.com/infobloxopen/terraform-provider-infoblox/internal/utils"
)

// ValidateNotificationRule validates the NotificationRule configuration.
func ValidateNotificationRule(ctx context.Context, data NotificationRuleModel, resp *resource.ValidateConfigResponse) {
	if nios := flex.ExpandNestedObject[NIOSNotificationRuleModel](ctx, data.NIOS, &resp.Diagnostics); nios != nil {
		validateNotificationRuleNIOSConfig(ctx, nios, resp)
	}
}

func validateNotificationRuleNIOSConfig(ctx context.Context, m *NIOSNotificationRuleModel, resp *resource.ValidateConfigResponse) {
	utils.ValidateScheduleConfig(
		m.ScheduledEvent,
		"",
		path.Root("nios").AtName("scheduled_event"),
		&resp.Diagnostics,
	)
}
