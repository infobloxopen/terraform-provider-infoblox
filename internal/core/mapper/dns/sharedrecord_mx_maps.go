package dns

import "github.com/infobloxopen/terraform-provider-infoblox/internal/core"

// SharedrecordMxNIOSFieldMap maps infoblox model fields to NIOS struct fields
var SharedrecordMxNIOSFieldMap = map[string]string{
	"Id":                     "Ref",
	"NIOS.Comment":           "Comment",
	"NIOS.Disable":           "Disable",
	"NIOS.MailExchanger":     "MailExchanger",
	"NIOS.Name":              "Name",
	"NIOS.Preference":        "Preference",
	"NIOS.SharedRecordGroup": "SharedRecordGroup",
	"NIOS.Ttl":               "Ttl",
	"NIOS.UseTtl":            "UseTtl",
}

// TODO: only searchable fields should be included here
// SharedrecordMxFilterFieldMap maps infoblox filter keys to backend-specific API filter field names
var SharedrecordMxFilterFieldMap = map[core.BackendType]map[string]string{
	core.BackendNIOS: {
		"id":                       "_ref",
		"nios.comment":             "comment",
		"nios.disable":             "disable",
		"nios.ext_attrs":           "extattrs",
		"nios.mail_exchanger":      "mail_exchanger",
		"nios.name":                "name",
		"nios.preference":          "preference",
		"nios.shared_record_group": "shared_record_group",
		"nios.ttl":                 "ttl",
		"nios.use_ttl":             "use_ttl",
	},
}
