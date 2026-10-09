package dns

import "github.com/infobloxopen/terraform-provider-infoblox/internal/core"

// SharedrecordANIOSFieldMap maps infoblox model fields to NIOS struct fields
var SharedrecordANIOSFieldMap = map[string]string{
	"Id":                     "Ref",
	"NIOS.Comment":           "Comment",
	"NIOS.Disable":           "Disable",
	"NIOS.Ipv4addr":          "Ipv4addr",
	"NIOS.Name":              "Name",
	"NIOS.SharedRecordGroup": "SharedRecordGroup",
	"NIOS.Ttl":               "Ttl",
	"NIOS.UseTtl":            "UseTtl",
}

// TODO: only searchable fields should be included here
// SharedrecordAFilterFieldMap maps infoblox filter keys to backend-specific API filter field names
var SharedrecordAFilterFieldMap = map[core.BackendType]map[string]string{
	core.BackendNIOS: {
		"id":                       "_ref",
		"nios.comment":             "comment",
		"nios.disable":             "disable",
		"nios.ext_attrs":           "extattrs",
		"nios.ipv4addr":            "ipv4addr",
		"nios.name":                "name",
		"nios.shared_record_group": "shared_record_group",
		"nios.ttl":                 "ttl",
		"nios.use_ttl":             "use_ttl",
	},
}
