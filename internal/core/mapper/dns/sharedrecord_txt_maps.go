package dns

import "github.com/infobloxopen/terraform-provider-infoblox/internal/core"

// SharedrecordTxtNIOSFieldMap maps infoblox model fields to NIOS struct fields
var SharedrecordTxtNIOSFieldMap = map[string]string{
	"Id":                     "Ref",
	"NIOS.Comment":           "Comment",
	"NIOS.Disable":           "Disable",
	"NIOS.Name":              "Name",
	"NIOS.SharedRecordGroup": "SharedRecordGroup",
	"NIOS.Text":              "Text",
	"NIOS.Ttl":               "Ttl",
	"NIOS.UseTtl":            "UseTtl",
}

// TODO: only searchable fields should be included here
// SharedrecordTxtFilterFieldMap maps infoblox filter keys to backend-specific API filter field names
var SharedrecordTxtFilterFieldMap = map[core.BackendType]map[string]string{
	core.BackendNIOS: {
		"id":                       "_ref",
		"nios.comment":             "comment",
		"nios.disable":             "disable",
		"nios.ext_attrs":           "extattrs",
		"nios.name":                "name",
		"nios.shared_record_group": "shared_record_group",
		"nios.text":                "text",
		"nios.ttl":                 "ttl",
		"nios.use_ttl":             "use_ttl",
	},
}
