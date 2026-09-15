package dtc

import "github.com/infobloxopen/terraform-provider-infoblox/internal/core"

// DtcLbdnNIOSFieldMap maps infoblox model fields to NIOS struct fields
var DtcLbdnNIOSFieldMap = map[string]string{
	"Id":                            "Ref",
	"NIOS.AuthZones":                "AuthZones",
	"NIOS.AutoConsolidatedMonitors": "AutoConsolidatedMonitors",
	"NIOS.Comment":                  "Comment",
	"NIOS.Disable":                  "Disable",
	"NIOS.LbMethod":                 "LbMethod",
	"NIOS.Name":                     "Name",
	"NIOS.Patterns":                 "Patterns",
	"NIOS.Persistence":              "Persistence",
	"NIOS.Pools":                    "Pools",
	"NIOS.Priority":                 "Priority",
	"NIOS.Topology":                 "Topology",
	"NIOS.Ttl":                      "Ttl",
	"NIOS.Types":                    "Types",
	"NIOS.UseTtl":                   "UseTtl",
}

// DtcLbdnUDDIFieldMap maps infoblox model fields to UDDI struct fields
var DtcLbdnUDDIFieldMap = map[string]string{
	"UDDI.Comment":            "Comment",
	"UDDI.Disabled":           "Disabled",
	"UDDI.DtcPolicy":          "DtcPolicy",
	"UDDI.InheritanceSources": "InheritanceSources",
	"UDDI.Name":               "Name",
	"UDDI.Precedence":         "Precedence",
	"UDDI.Tags":               "Tags",
	"UDDI.Ttl":                "Ttl",
	"UDDI.View":               "View",
}

// TODO: only searchable fields should be included here
// DtcLbdnFilterFieldMap maps infoblox filter keys to backend-specific API filter field names
var DtcLbdnFilterFieldMap = map[core.BackendType]map[string]string{
	core.BackendNIOS: {
		"id":                              "_ref",
		"nios.auth_zones":                 "auth_zones",
		"nios.auto_consolidated_monitors": "auto_consolidated_monitors",
		"nios.comment":                    "comment",
		"nios.disable":                    "disable",
		"nios.ext_attrs":                  "extattrs",
		"nios.lb_method":                  "lb_method",
		"nios.name":                       "name",
		"nios.patterns":                   "patterns",
		"nios.persistence":                "persistence",
		"nios.pools":                      "pools",
		"nios.priority":                   "priority",
		"nios.topology":                   "topology",
		"nios.ttl":                        "ttl",
		"nios.types":                      "types",
		"nios.use_ttl":                    "use_ttl",
	},
	core.BackendUDDI: {
		"uddi.comment":             "comment",
		"uddi.disabled":            "disabled",
		"uddi.dtc_policy":          "dtc_policy",
		"uddi.inheritance_sources": "inheritance_sources",
		"uddi.name":                "name",
		"uddi.precedence":          "precedence",
		"uddi.tags":                "tags",
		"uddi.ttl":                 "ttl",
		"uddi.view":                "view",
	},
}
