package grid

import "github.com/infobloxopen/terraform-provider-infoblox/internal/core"

// UpgradegroupNIOSFieldMap maps infoblox model fields to NIOS struct fields
var UpgradegroupNIOSFieldMap = map[string]string{
	"Id":                              "Ref",
	"NIOS.Comment":                    "Comment",
	"NIOS.DistributionDependentGroup": "DistributionDependentGroup",
	"NIOS.DistributionPolicy":         "DistributionPolicy",
	"NIOS.DistributionTime":           "DistributionTime",
	"NIOS.Members":                    "Members",
	"NIOS.Name":                       "Name",
	"NIOS.UpgradeDependentGroup":      "UpgradeDependentGroup",
	"NIOS.UpgradePolicy":              "UpgradePolicy",
	"NIOS.UpgradeTime":                "UpgradeTime",
}

// TODO: only searchable fields should be included here
// UpgradegroupFilterFieldMap maps infoblox filter keys to backend-specific API filter field names
var UpgradegroupFilterFieldMap = map[core.BackendType]map[string]string{
	core.BackendNIOS: {
		"id":                                "_ref",
		"nios.comment":                      "comment",
		"nios.distribution_dependent_group": "distribution_dependent_group",
		"nios.distribution_policy":          "distribution_policy",
		"nios.distribution_time":            "distribution_time",
		"nios.members":                      "members",
		"nios.name":                         "name",
		"nios.upgrade_dependent_group":      "upgrade_dependent_group",
		"nios.upgrade_policy":               "upgrade_policy",
		"nios.upgrade_time":                 "upgrade_time",
	},
}
