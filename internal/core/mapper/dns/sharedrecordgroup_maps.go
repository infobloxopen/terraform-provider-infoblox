package dns

import "github.com/infobloxopen/terraform-provider-infoblox/internal/core"

// SharedrecordgroupNIOSFieldMap maps infoblox model fields to NIOS struct fields
var SharedrecordgroupNIOSFieldMap = map[string]string{
	"Id":                       "Ref",
	"NIOS.Comment":             "Comment",
	"NIOS.Name":                "Name",
	"NIOS.RecordNamePolicy":    "RecordNamePolicy",
	"NIOS.UseRecordNamePolicy": "UseRecordNamePolicy",
	"NIOS.ZoneAssociations":    "ZoneAssociations",
}

// TODO: only searchable fields should be included here
// SharedrecordgroupFilterFieldMap maps infoblox filter keys to backend-specific API filter field names
var SharedrecordgroupFilterFieldMap = map[core.BackendType]map[string]string{
	core.BackendNIOS: {
		"id":                          "_ref",
		"nios.comment":                "comment",
		"nios.ext_attrs":              "extattrs",
		"nios.name":                   "name",
		"nios.record_name_policy":     "record_name_policy",
		"nios.use_record_name_policy": "use_record_name_policy",
		"nios.zone_associations":      "zone_associations",
	},
}
