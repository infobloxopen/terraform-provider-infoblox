package dns

import "github.com/infobloxopen/terraform-provider-infoblox/internal/core"

// SharedrecordCnameNIOSFieldMap maps infoblox model fields to NIOS struct fields
var SharedrecordCnameNIOSFieldMap = map[string]string{
	"Id":                     "Ref",
	"NIOS.Canonical":         "Canonical",
	"NIOS.Comment":           "Comment",
	"NIOS.Disable":           "Disable",
	"NIOS.Name":              "Name",
	"NIOS.SharedRecordGroup": "SharedRecordGroup",
	"NIOS.Ttl":               "Ttl",
	"NIOS.UseTtl":            "UseTtl",
}

// TODO: only searchable fields should be included here
// SharedrecordCnameFilterFieldMap maps infoblox filter keys to backend-specific API filter field names
var SharedrecordCnameFilterFieldMap = map[core.BackendType]map[string]string{
	core.BackendNIOS: {
		"id":                       "_ref",
		"nios.canonical":           "canonical",
		"nios.comment":             "comment",
		"nios.disable":             "disable",
		"nios.ext_attrs":           "extattrs",
		"nios.name":                "name",
		"nios.shared_record_group": "shared_record_group",
		"nios.ttl":                 "ttl",
		"nios.use_ttl":             "use_ttl",
	},
}
