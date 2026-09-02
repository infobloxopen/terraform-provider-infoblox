package rpz

import "github.com/infobloxopen/terraform-provider-infoblox/internal/core"

// RecordRpzNaptrNIOSFieldMap maps infoblox model fields to NIOS struct fields
var RecordRpzNaptrNIOSFieldMap = map[string]string{
	"Id":               "Ref",
	"NIOS.Comment":     "Comment",
	"NIOS.Disable":     "Disable",
	"NIOS.Flags":       "Flags",
	"NIOS.Name":        "Name",
	"NIOS.Order":       "Order",
	"NIOS.Preference":  "Preference",
	"NIOS.Regexp":      "Regexp",
	"NIOS.Replacement": "Replacement",
	"NIOS.RpZone":      "RpZone",
	"NIOS.Services":    "Services",
	"NIOS.Ttl":         "Ttl",
	"NIOS.UseTtl":      "UseTtl",
	"NIOS.View":        "View",
}

// TODO: only searchable fields should be included here
// RecordRpzNaptrFilterFieldMap maps infoblox filter keys to backend-specific API filter field names
var RecordRpzNaptrFilterFieldMap = map[core.BackendType]map[string]string{
	core.BackendNIOS: {
		"id":               "_ref",
		"nios.comment":     "comment",
		"nios.disable":     "disable",
		"nios.ext_attrs":   "extattrs",
		"nios.flags":       "flags",
		"nios.name":        "name",
		"nios.order":       "order",
		"nios.preference":  "preference",
		"nios.regexp":      "regexp",
		"nios.replacement": "replacement",
		"nios.rp_zone":     "rp_zone",
		"nios.services":    "services",
		"nios.ttl":         "ttl",
		"nios.use_ttl":     "use_ttl",
		"nios.view":        "view",
	},
}
