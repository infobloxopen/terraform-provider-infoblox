package dns

import "github.com/infobloxopen/terraform-provider-infoblox/internal/core"

// RecordPtrNIOSFieldMap maps infoblox model fields to NIOS struct fields
var RecordPtrNIOSFieldMap = map[string]string{
	"Id":                     "Ref",
	"NIOS.Comment":           "Comment",
	"NIOS.Creator":           "Creator",
	"NIOS.DdnsPrincipal":     "DdnsPrincipal",
	"NIOS.DdnsProtected":     "DdnsProtected",
	"NIOS.Disable":           "Disable",
	"NIOS.ForbidReclamation": "ForbidReclamation",
	"NIOS.Ipv4addr":          "Ipv4addr.String",
	"NIOS.Ipv6addr":          "Ipv6addr.String",
	"NIOS.Name":              "Name",
	"NIOS.Ptrdname":          "Ptrdname",
	"NIOS.Ttl":               "Ttl",
	"NIOS.UseTtl":            "UseTtl",
	"NIOS.View":              "View",
	"NIOS.FuncCall":          "FuncCall",
}

// RecordPtrUDDIFieldMap maps infoblox model fields to UDDI struct fields
var RecordPtrUDDIFieldMap = map[string]string{
	"UDDI.AbsoluteNameSpec":   "AbsoluteNameSpec",
	"UDDI.Comment":            "Comment",
	"UDDI.Disabled":           "Disabled",
	"UDDI.InheritanceSources": "InheritanceSources",
	"UDDI.NameInZone":         "NameInZone",
	"UDDI.Options":            "Options",
	"UDDI.Rdata":              "Rdata",
	"UDDI.Tags":               "Tags",
	"UDDI.Ttl":                "Ttl",
	"UDDI.Type":               "Type",
	"UDDI.View":               "View",
	"UDDI.Zone":               "Zone",
}

// TODO: only searchable fields should be included here
// RecordPtrFilterFieldMap maps infoblox filter keys to backend-specific API filter field names
var RecordPtrFilterFieldMap = map[core.BackendType]map[string]string{
	core.BackendNIOS: {
		"id":                      "_ref",
		"nios.comment":            "comment",
		"nios.creator":            "creator",
		"nios.ddns_principal":     "ddns_principal",
		"nios.ddns_protected":     "ddns_protected",
		"nios.disable":            "disable",
		"nios.ext_attrs":          "extattrs",
		"nios.forbid_reclamation": "forbid_reclamation",
		"nios.ipv4addr":           "ipv4addr",
		"nios.ipv6addr":           "ipv6addr",
		"nios.name":               "name",
		"nios.ptrdname":           "ptrdname",
		"nios.ttl":                "ttl",
		"nios.use_ttl":            "use_ttl",
		"nios.view":               "view",
	},
	core.BackendUDDI: {
		"uddi.absolute_name_spec":  "absolute_name_spec",
		"uddi.comment":             "comment",
		"uddi.disabled":            "disabled",
		"uddi.inheritance_sources": "inheritance_sources",
		"uddi.name_in_zone":        "name_in_zone",
		"uddi.options":             "options",
		"uddi.rdata":               "rdata",
		"uddi.tags":                "tags",
		"uddi.ttl":                 "ttl",
		"uddi.type":                "type",
		"uddi.view":                "view",
		"uddi.zone":                "zone",
	},
}
