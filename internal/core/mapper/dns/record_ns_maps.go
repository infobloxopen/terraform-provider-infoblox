package dns

import "github.com/infobloxopen/terraform-provider-infoblox/internal/core"

// RecordNsNIOSFieldMap maps infoblox model fields to NIOS struct fields
var RecordNsNIOSFieldMap = map[string]string{
	"Id":                    "Ref",
	"NIOS.Addresses":        "Addresses",
	"NIOS.MsDelegationName": "MsDelegationName",
	"NIOS.Name":             "Name",
	"NIOS.Nameserver":       "Nameserver",
	"NIOS.View":             "View",
}

// RecordNsUDDIFieldMap maps infoblox model fields to UDDI struct fields
var RecordNsUDDIFieldMap = map[string]string{
	"UDDI.AbsoluteNameSpec":   "AbsoluteNameSpec",
	"UDDI.Comment":            "Comment",
	"UDDI.Disabled":           "Disabled",
	"UDDI.InheritanceSources": "InheritanceSources",
	"UDDI.NameInZone":         "NameInZone",
	"UDDI.Rdata":              "Rdata",
	"UDDI.Tags":               "Tags",
	"UDDI.Ttl":                "Ttl",
	"UDDI.Type":               "Type",
	"UDDI.View":               "View",
	"UDDI.Zone":               "Zone",
}

// TODO: only searchable fields should be included here
// RecordNsFilterFieldMap maps infoblox filter keys to backend-specific API filter field names
var RecordNsFilterFieldMap = map[core.BackendType]map[string]string{
	core.BackendNIOS: {
		"id":                      "_ref",
		"nios.addresses":          "addresses",
		"nios.ms_delegation_name": "ms_delegation_name",
		"nios.name":               "name",
		"nios.nameserver":         "nameserver",
		"nios.view":               "view",
	},
	core.BackendUDDI: {
		"uddi.absolute_name_spec":  "absolute_name_spec",
		"uddi.comment":             "comment",
		"uddi.disabled":            "disabled",
		"uddi.inheritance_sources": "inheritance_sources",
		"uddi.name_in_zone":        "name_in_zone",
		"uddi.rdata":               "rdata",
		"uddi.tags":                "tags",
		"uddi.ttl":                 "ttl",
		"uddi.type":                "type",
		"uddi.view":                "view",
		"uddi.zone":                "zone",
	},
}
