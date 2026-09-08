package notification

import (
	niosnotification "github.com/infobloxopen/infoblox-nios-go-client/notification"
)

// Infoblox NotificationRestEndpoint model
type NotificationRestEndpoint struct {
	Id   *string
	NIOS *NIOSNotificationRestEndpointExt
}

// NIOSNotificationRestEndpointExt - NIOS specific fields for NotificationRestEndpoint
type NIOSNotificationRestEndpointExt struct {
	ClientCertificateToken *string
	Comment                *string
	ExtAttrs               map[string]any
	LogLevel               *string
	Name                   *string
	OutboundMemberType     *string
	OutboundMembers        []string
	Password               *string
	ServerCertValidation   *string
	SyncDisabled           *bool
	TemplateInstance       *niosnotification.NotificationRestEndpointTemplateInstance
	Timeout                *int64
	Uri                    *string
	Username               *string
	VendorIdentifier       *string
	WapiUserName           *string
	WapiUserPassword       *string
}
