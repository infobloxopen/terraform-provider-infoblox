package notification

import (
	niosnotification "github.com/infobloxopen/infoblox-nios-go-client/notification"
)

// Infoblox NotificationRule model
type NotificationRule struct {
	Id   *string
	NIOS *NIOSNotificationRuleExt
}

// NIOSNotificationRuleExt - NIOS specific fields for NotificationRule
type NIOSNotificationRuleExt struct {
	AllMembers                       *bool
	Comment                          *string
	Disable                          *bool
	EnableEventDeduplication         *bool
	EnableEventDeduplicationLog      *bool
	EventDeduplicationFields         []string
	EventDeduplicationLookbackPeriod *int64
	EventPriority                    *string
	EventType                        *string
	ExpressionList                   []niosnotification.NotificationRuleExpressionList
	Name                             *string
	NotificationAction               *string
	NotificationTarget               *string
	PublishSettings                  *niosnotification.NotificationRulePublishSettings
	ScheduledEvent                   *niosnotification.NotificationRuleScheduledEvent
	SelectedMembers                  []string
	TemplateInstance                 *niosnotification.NotificationRuleTemplateInstance
	UsePublishSettings               *bool
}
