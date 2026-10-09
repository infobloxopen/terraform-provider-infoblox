package dtc

import "github.com/infobloxopen/terraform-provider-infoblox/internal/core"

// DtcMonitorIcmpNIOSFieldMap maps infoblox model fields to NIOS struct fields
var DtcMonitorIcmpNIOSFieldMap = map[string]string{
	"Id":             "Ref",
	"NIOS.Comment":   "Comment",
	"NIOS.Interval":  "Interval",
	"NIOS.Name":      "Name",
	"NIOS.RetryDown": "RetryDown",
	"NIOS.RetryUp":   "RetryUp",
	"NIOS.Timeout":   "Timeout",
}

// DtcMonitorIcmpUDDIFieldMap maps infoblox model fields to UDDI struct fields
var DtcMonitorIcmpUDDIFieldMap = map[string]string{
	"UDDI.Comment":   "Comment",
	"UDDI.Disabled":  "Disabled",
	"UDDI.Interval":  "Interval",
	"UDDI.Name":      "Name",
	"UDDI.RetryDown": "RetryDown",
	"UDDI.RetryUp":   "RetryUp",
	"UDDI.Tags":      "Tags",
	"UDDI.Timeout":   "Timeout",
}

// TODO: only searchable fields should be included here
// DtcMonitorIcmpFilterFieldMap maps infoblox filter keys to backend-specific API filter field names
var DtcMonitorIcmpFilterFieldMap = map[core.BackendType]map[string]string{
	core.BackendNIOS: {
		"id":              "_ref",
		"nios.comment":    "comment",
		"nios.ext_attrs":  "extattrs",
		"nios.interval":   "interval",
		"nios.name":       "name",
		"nios.retry_down": "retry_down",
		"nios.retry_up":   "retry_up",
		"nios.timeout":    "timeout",
	},
	core.BackendUDDI: {
		"uddi.comment":    "comment",
		"uddi.disabled":   "disabled",
		"uddi.interval":   "interval",
		"uddi.name":       "name",
		"uddi.retry_down": "retry_down",
		"uddi.retry_up":   "retry_up",
		"uddi.tags":       "tags",
		"uddi.timeout":    "timeout",
	},
}
