package notification

import "github.com/infobloxopen/terraform-provider-infoblox/internal/core"

// NotificationRuleNIOSFieldMap maps infoblox model fields to NIOS struct fields
var NotificationRuleNIOSFieldMap = map[string]string{
	"Id":                                    "Ref",
	"NIOS.AllMembers":                       "AllMembers",
	"NIOS.Comment":                          "Comment",
	"NIOS.Disable":                          "Disable",
	"NIOS.EnableEventDeduplication":         "EnableEventDeduplication",
	"NIOS.EnableEventDeduplicationLog":      "EnableEventDeduplicationLog",
	"NIOS.EventDeduplicationFields":         "EventDeduplicationFields",
	"NIOS.EventDeduplicationLookbackPeriod": "EventDeduplicationLookbackPeriod",
	"NIOS.EventPriority":                    "EventPriority",
	"NIOS.EventType":                        "EventType",
	"NIOS.ExpressionList":                   "ExpressionList",
	"NIOS.Name":                             "Name",
	"NIOS.NotificationAction":               "NotificationAction",
	"NIOS.NotificationTarget":               "NotificationTarget",
	"NIOS.PublishSettings":                  "PublishSettings",
	"NIOS.ScheduledEvent":                   "ScheduledEvent",
	"NIOS.SelectedMembers":                  "SelectedMembers",
	"NIOS.TemplateInstance":                 "TemplateInstance",
	"NIOS.UsePublishSettings":               "UsePublishSettings",
}

// TODO: only searchable fields should be included here
// NotificationRuleFilterFieldMap maps infoblox filter keys to backend-specific API filter field names
var NotificationRuleFilterFieldMap = map[core.BackendType]map[string]string{
	core.BackendNIOS: {
		"id":                                       "_ref",
		"nios.all_members":                         "all_members",
		"nios.comment":                             "comment",
		"nios.disable":                             "disable",
		"nios.enable_event_deduplication":          "enable_event_deduplication",
		"nios.enable_event_deduplication_log":      "enable_event_deduplication_log",
		"nios.event_deduplication_fields":          "event_deduplication_fields",
		"nios.event_deduplication_lookback_period": "event_deduplication_lookback_period",
		"nios.event_priority":                      "event_priority",
		"nios.event_type":                          "event_type",
		"nios.expression_list":                     "expression_list",
		"nios.name":                                "name",
		"nios.notification_action":                 "notification_action",
		"nios.notification_target":                 "notification_target",
		"nios.publish_settings":                    "publish_settings",
		"nios.scheduled_event":                     "scheduled_event",
		"nios.selected_members":                    "selected_members",
		"nios.template_instance":                   "template_instance",
		"nios.use_publish_settings":                "use_publish_settings",
	},
}
