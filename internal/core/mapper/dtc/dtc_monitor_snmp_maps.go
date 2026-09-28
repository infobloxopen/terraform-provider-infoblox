package dtc

import "github.com/infobloxopen/terraform-provider-infoblox/internal/core"

// DtcMonitorSnmpNIOSFieldMap maps infoblox model fields to NIOS struct fields
var DtcMonitorSnmpNIOSFieldMap = map[string]string{
	"Id":             "Ref",
	"NIOS.Comment":   "Comment",
	"NIOS.Community": "Community",
	"NIOS.Context":   "Context",
	"NIOS.EngineId":  "EngineId",
	"NIOS.Interval":  "Interval",
	"NIOS.Name":      "Name",
	"NIOS.Oids":      "Oids",
	"NIOS.Port":      "Port",
	"NIOS.RetryDown": "RetryDown",
	"NIOS.RetryUp":   "RetryUp",
	"NIOS.Timeout":   "Timeout",
	"NIOS.User":      "User",
	"NIOS.Version":   "Version",
}

// DtcMonitorSnmpUDDIFieldMap maps infoblox model fields to UDDI struct fields
var DtcMonitorSnmpUDDIFieldMap = map[string]string{
	"UDDI.CheckList":         "CheckList",
	"UDDI.Comment":           "Comment",
	"UDDI.Community":         "Community",
	"UDDI.ContextEngineId":   "ContextEngineId",
	"UDDI.ContextName":       "ContextName",
	"UDDI.Disabled":          "Disabled",
	"UDDI.Interval":          "Interval",
	"UDDI.Name":              "Name",
	"UDDI.Port":              "Port",
	"UDDI.RetryDown":         "RetryDown",
	"UDDI.RetryUp":           "RetryUp",
	"UDDI.Tags":              "Tags",
	"UDDI.Timeout":           "Timeout",
	"UDDI.UserSecurityModel": "UserSecurityModel",
	"UDDI.Version":           "Version",
}

// TODO: only searchable fields should be included here
// DtcMonitorSnmpFilterFieldMap maps infoblox filter keys to backend-specific API filter field names
var DtcMonitorSnmpFilterFieldMap = map[core.BackendType]map[string]string{
	core.BackendNIOS: {
		"id":              "_ref",
		"nios.comment":    "comment",
		"nios.community":  "community",
		"nios.context":    "context",
		"nios.engine_id":  "engine_id",
		"nios.ext_attrs":  "extattrs",
		"nios.interval":   "interval",
		"nios.name":       "name",
		"nios.oids":       "oids",
		"nios.port":       "port",
		"nios.retry_down": "retry_down",
		"nios.retry_up":   "retry_up",
		"nios.timeout":    "timeout",
		"nios.user":       "user",
		"nios.version":    "version",
	},
	core.BackendUDDI: {
		"uddi.check_list":          "check_list",
		"uddi.comment":             "comment",
		"uddi.community":           "community",
		"uddi.context_engine_id":   "context_engine_id",
		"uddi.context_name":        "context_name",
		"uddi.disabled":            "disabled",
		"uddi.interval":            "interval",
		"uddi.name":                "name",
		"uddi.port":                "port",
		"uddi.retry_down":          "retry_down",
		"uddi.retry_up":            "retry_up",
		"uddi.tags":                "tags",
		"uddi.timeout":             "timeout",
		"uddi.user_security_model": "user_security_model",
		"uddi.version":             "version",
	},
}
