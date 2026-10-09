package rpz

import "github.com/infobloxopen/terraform-provider-infoblox/internal/core"

// RecordRpzCnameIpaddressNIOSFieldMap maps infoblox model fields to NIOS struct fields
var RecordRpzCnameIpaddressNIOSFieldMap = map[string]string{
	"Id":             "Ref",
	"NIOS.Canonical": "Canonical",
	"NIOS.Comment":   "Comment",
	"NIOS.Disable":   "Disable",
	"NIOS.Name":      "Name",
	"NIOS.RpZone":    "RpZone",
	"NIOS.Ttl":       "Ttl",
	"NIOS.UseTtl":    "UseTtl",
	"NIOS.View":      "View",
}

// TODO: only searchable fields should be included here
// RecordRpzCnameIpaddressFilterFieldMap maps infoblox filter keys to backend-specific API filter field names
var RecordRpzCnameIpaddressFilterFieldMap = map[core.BackendType]map[string]string{
	core.BackendNIOS: {
		"id":             "_ref",
		"nios.canonical": "canonical",
		"nios.comment":   "comment",
		"nios.disable":   "disable",
		"nios.ext_attrs": "extattrs",
		"nios.name":      "name",
		"nios.rp_zone":   "rp_zone",
		"nios.ttl":       "ttl",
		"nios.use_ttl":   "use_ttl",
		"nios.view":      "view",
	},
}
