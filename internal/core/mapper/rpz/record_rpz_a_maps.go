package rpz

import "github.com/infobloxopen/terraform-provider-infoblox/internal/core"

// RecordRpzANIOSFieldMap maps infoblox model fields to NIOS struct fields
var RecordRpzANIOSFieldMap = map[string]string{
	"Id":            "Ref",
	"NIOS.Comment":  "Comment",
	"NIOS.Disable":  "Disable",
	"NIOS.Ipv4addr": "Ipv4addr",
	"NIOS.Name":     "Name",
	"NIOS.RpZone":   "RpZone",
	"NIOS.Ttl":      "Ttl",
	"NIOS.UseTtl":   "UseTtl",
	"NIOS.View":     "View",
}

// TODO: only searchable fields should be included here
// RecordRpzAFilterFieldMap maps infoblox filter keys to backend-specific API filter field names
var RecordRpzAFilterFieldMap = map[core.BackendType]map[string]string{
	core.BackendNIOS: {
		"id":             "_ref",
		"nios.comment":   "comment",
		"nios.disable":   "disable",
		"nios.ext_attrs": "extattrs",
		"nios.ipv4addr":  "ipv4addr",
		"nios.name":      "name",
		"nios.rp_zone":   "rp_zone",
		"nios.ttl":       "ttl",
		"nios.use_ttl":   "use_ttl",
		"nios.view":      "view",
	},
}
