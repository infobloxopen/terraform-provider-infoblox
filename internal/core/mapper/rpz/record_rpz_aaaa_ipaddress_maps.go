package rpz

import "github.com/infobloxopen/terraform-provider-infoblox/internal/core"

// RecordRpzAaaaIpaddressNIOSFieldMap maps infoblox model fields to NIOS struct fields
var RecordRpzAaaaIpaddressNIOSFieldMap = map[string]string{
	"Id":            "Ref",
	"NIOS.Comment":  "Comment",
	"NIOS.Disable":  "Disable",
	"NIOS.Ipv6addr": "Ipv6addr",
	"NIOS.Name":     "Name",
	"NIOS.RpZone":   "RpZone",
	"NIOS.Ttl":      "Ttl",
	"NIOS.UseTtl":   "UseTtl",
	"NIOS.View":     "View",
}

// TODO: only searchable fields should be included here
// RecordRpzAaaaIpaddressFilterFieldMap maps infoblox filter keys to backend-specific API filter field names
var RecordRpzAaaaIpaddressFilterFieldMap = map[core.BackendType]map[string]string{
	core.BackendNIOS: {
		"id":             "_ref",
		"nios.comment":   "comment",
		"nios.disable":   "disable",
		"nios.ext_attrs": "extattrs",
		"nios.ipv6addr":  "ipv6addr",
		"nios.name":      "name",
		"nios.rp_zone":   "rp_zone",
		"nios.ttl":       "ttl",
		"nios.use_ttl":   "use_ttl",
		"nios.view":      "view",
	},
}
