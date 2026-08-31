package dtc

import "github.com/infobloxopen/terraform-provider-infoblox/internal/core"

// DtcTopologyNIOSFieldMap maps infoblox model fields to NIOS struct fields
var DtcTopologyNIOSFieldMap = map[string]string{
	"Id":           "Ref",
	"NIOS.Comment": "Comment",
	"NIOS.Name":    "Name",
	"NIOS.Rules":   "Rules",
}

// TODO: only searchable fields should be included here
// DtcTopologyFilterFieldMap maps infoblox filter keys to backend-specific API filter field names
var DtcTopologyFilterFieldMap = map[core.BackendType]map[string]string{
	core.BackendNIOS: {
		"id":             "_ref",
		"nios.comment":   "comment",
		"nios.ext_attrs": "extattrs",
		"nios.name":      "name",
		"nios.rules":     "rules",
	},
}
