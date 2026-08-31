package misc

import "github.com/infobloxopen/terraform-provider-infoblox/internal/core"

// RulesetNIOSFieldMap maps infoblox model fields to NIOS struct fields
var RulesetNIOSFieldMap = map[string]string{
	"Id":                 "Ref",
	"NIOS.Comment":       "Comment",
	"NIOS.Disabled":      "Disabled",
	"NIOS.Name":          "Name",
	"NIOS.NxdomainRules": "NxdomainRules",
	"NIOS.Type":          "Type",
}

// TODO: only searchable fields should be included here
// RulesetFilterFieldMap maps infoblox filter keys to backend-specific API filter field names
var RulesetFilterFieldMap = map[core.BackendType]map[string]string{
	core.BackendNIOS: {
		"id":                  "_ref",
		"nios.comment":        "comment",
		"nios.disabled":       "disabled",
		"nios.name":           "name",
		"nios.nxdomain_rules": "nxdomain_rules",
		"nios.type":           "type",
	},
}
