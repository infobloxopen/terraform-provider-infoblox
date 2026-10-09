package grid

import "github.com/infobloxopen/terraform-provider-infoblox/internal/core"

// ServicerestartGroupNIOSFieldMap maps infoblox model fields to NIOS struct fields
var ServicerestartGroupNIOSFieldMap = map[string]string{
	"Id":                     "Ref",
	"NIOS.Comment":           "Comment",
	"NIOS.Members":           "Members",
	"NIOS.Mode":              "Mode",
	"NIOS.Name":              "Name",
	"NIOS.RecurringSchedule": "RecurringSchedule",
	"NIOS.Service":           "Service",
}

// TODO: only searchable fields should be included here
// ServicerestartGroupFilterFieldMap maps infoblox filter keys to backend-specific API filter field names
var ServicerestartGroupFilterFieldMap = map[core.BackendType]map[string]string{
	core.BackendNIOS: {
		"id":                      "_ref",
		"nios.comment":            "comment",
		"nios.ext_attrs":          "extattrs",
		"nios.members":            "members",
		"nios.mode":               "mode",
		"nios.name":               "name",
		"nios.recurring_schedule": "recurring_schedule",
		"nios.service":            "service",
	},
}
