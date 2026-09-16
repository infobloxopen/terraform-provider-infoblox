package notification

import "github.com/infobloxopen/terraform-provider-infoblox/internal/core"

// NotificationRestEndpointNIOSFieldMap maps infoblox model fields to NIOS struct fields
var NotificationRestEndpointNIOSFieldMap = map[string]string{
	"Id":                          "Ref",
	"NIOS.ClientCertificateToken": "ClientCertificateToken",
	"NIOS.Comment":                "Comment",
	"NIOS.LogLevel":               "LogLevel",
	"NIOS.Name":                   "Name",
	"NIOS.OutboundMemberType":     "OutboundMemberType",
	"NIOS.OutboundMembers":        "OutboundMembers",
	"NIOS.Password":               "Password",
	"NIOS.ServerCertValidation":   "ServerCertValidation",
	"NIOS.SyncDisabled":           "SyncDisabled",
	"NIOS.TemplateInstance":       "TemplateInstance",
	"NIOS.Timeout":                "Timeout",
	"NIOS.Uri":                    "Uri",
	"NIOS.Username":               "Username",
	"NIOS.VendorIdentifier":       "VendorIdentifier",
	"NIOS.WapiUserName":           "WapiUserName",
	"NIOS.WapiUserPassword":       "WapiUserPassword",
}

// TODO: only searchable fields should be included here
// NotificationRestEndpointFilterFieldMap maps infoblox filter keys to backend-specific API filter field names
var NotificationRestEndpointFilterFieldMap = map[core.BackendType]map[string]string{
	core.BackendNIOS: {
		"id":                            "_ref",
		"nios.client_certificate_token": "client_certificate_token",
		"nios.comment":                  "comment",
		"nios.ext_attrs":                "extattrs",
		"nios.log_level":                "log_level",
		"nios.name":                     "name",
		"nios.outbound_member_type":     "outbound_member_type",
		"nios.outbound_members":         "outbound_members",
		"nios.password":                 "password",
		"nios.server_cert_validation":   "server_cert_validation",
		"nios.sync_disabled":            "sync_disabled",
		"nios.template_instance":        "template_instance",
		"nios.timeout":                  "timeout",
		"nios.uri":                      "uri",
		"nios.username":                 "username",
		"nios.vendor_identifier":        "vendor_identifier",
		"nios.wapi_user_name":           "wapi_user_name",
		"nios.wapi_user_password":       "wapi_user_password",
	},
}
