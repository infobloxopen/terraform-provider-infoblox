package dns

import "github.com/infobloxopen/terraform-provider-infoblox/internal/core"

// SharedrecordSrvNIOSFieldMap maps infoblox model fields to NIOS struct fields
var SharedrecordSrvNIOSFieldMap = map[string]string{
	"Id":                     "Ref",
	"NIOS.Comment":           "Comment",
	"NIOS.Disable":           "Disable",
	"NIOS.Name":              "Name",
	"NIOS.Port":              "Port",
	"NIOS.Priority":          "Priority",
	"NIOS.SharedRecordGroup": "SharedRecordGroup",
	"NIOS.Target":            "Target",
	"NIOS.Ttl":               "Ttl",
	"NIOS.UseTtl":            "UseTtl",
	"NIOS.Weight":            "Weight",
}

// TODO: only searchable fields should be included here
// SharedrecordSrvFilterFieldMap maps infoblox filter keys to backend-specific API filter field names
var SharedrecordSrvFilterFieldMap = map[core.BackendType]map[string]string{
	core.BackendNIOS: {
		"id":                       "_ref",
		"nios.comment":             "comment",
		"nios.disable":             "disable",
		"nios.ext_attrs":           "extattrs",
		"nios.name":                "name",
		"nios.port":                "port",
		"nios.priority":            "priority",
		"nios.shared_record_group": "shared_record_group",
		"nios.target":              "target",
		"nios.ttl":                 "ttl",
		"nios.use_ttl":             "use_ttl",
		"nios.weight":              "weight",
	},
}
