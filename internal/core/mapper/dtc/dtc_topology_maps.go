package dtc

import "github.com/infobloxopen/terraform-provider-infoblox/internal/core"

// DtcTopologyNIOSFieldMap maps infoblox model fields to NIOS struct fields
var DtcTopologyNIOSFieldMap = map[string]string{
	"Id":           "Ref",
	"NIOS.Comment": "Comment",
	"NIOS.Name":    "Name",
	"NIOS.Rules":   "Rules",
}

// DtcTopologyUDDIFieldMap maps infoblox model fields to UDDI struct fields
var DtcTopologyUDDIFieldMap = map[string]string{
	"UDDI.Comment":  "Comment",
	"UDDI.Disabled": "Disabled",
	"UDDI.Name":     "Name",
	"UDDI.Sources":  "Sources",
	"UDDI.Tags":     "Tags",
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
	core.BackendUDDI: {
		"uddi.comment":  "comment",
		"uddi.disabled": "disabled",
		"uddi.name":     "name",
		"uddi.sources":  "sources",
		"uddi.tags":     "tags",
	},
}
