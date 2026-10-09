package grid

import "github.com/infobloxopen/terraform-provider-infoblox/internal/core"

// DistributionscheduleNIOSFieldMap maps infoblox model fields to NIOS struct fields
var DistributionscheduleNIOSFieldMap = map[string]string{
	"Id":                 "Ref",
	"NIOS.Active":        "Active",
	"NIOS.StartTime":     "StartTime",
	"NIOS.TimeZone":      "TimeZone",
	"NIOS.UpgradeGroups": "UpgradeGroups",
}

// TODO: only searchable fields should be included here
// DistributionscheduleFilterFieldMap maps infoblox filter keys to backend-specific API filter field names
var DistributionscheduleFilterFieldMap = map[core.BackendType]map[string]string{
	core.BackendNIOS: {
		"id":                  "_ref",
		"nios.active":         "active",
		"nios.start_time":     "start_time",
		"nios.time_zone":      "time_zone",
		"nios.upgrade_groups": "upgrade_groups",
	},
}
