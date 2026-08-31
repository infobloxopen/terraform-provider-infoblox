package rpz

import "github.com/infobloxopen/terraform-provider-infoblox/internal/core"

// RecordRpzTxtNIOSFieldMap maps infoblox model fields to NIOS struct fields
var RecordRpzTxtNIOSFieldMap = map[string]string{
	"Id":           "Ref",
	"NIOS.Comment": "Comment",
	"NIOS.Disable": "Disable",
	"NIOS.Name":    "Name",
	"NIOS.RpZone":  "RpZone",
	"NIOS.Text":    "Text",
	"NIOS.Ttl":     "Ttl",
	"NIOS.UseTtl":  "UseTtl",
	"NIOS.View":    "View",
}

// TODO: only searchable fields should be included here
// RecordRpzTxtFilterFieldMap maps infoblox filter keys to backend-specific API filter field names
var RecordRpzTxtFilterFieldMap = map[core.BackendType]map[string]string{
	core.BackendNIOS: {
		"id":             "_ref",
		"nios.comment":   "comment",
		"nios.disable":   "disable",
		"nios.ext_attrs": "extattrs",
		"nios.name":      "name",
		"nios.rp_zone":   "rp_zone",
		"nios.text":      "text",
		"nios.ttl":       "ttl",
		"nios.use_ttl":   "use_ttl",
		"nios.view":      "view",
	},
}
