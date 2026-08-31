package dtc

import "github.com/infobloxopen/terraform-provider-infoblox/internal/core"

// DtcPoolNIOSFieldMap maps infoblox model fields to NIOS struct fields
var DtcPoolNIOSFieldMap = map[string]string{
	"Id":                            "Ref",
	"NIOS.AutoConsolidatedMonitors": "AutoConsolidatedMonitors",
	"NIOS.Availability":             "Availability",
	"NIOS.Comment":                  "Comment",
	"NIOS.ConsolidatedMonitors":     "ConsolidatedMonitors",
	"NIOS.Disable":                  "Disable",
	"NIOS.LbAlternateMethod":        "LbAlternateMethod",
	"NIOS.LbAlternateTopology":      "LbAlternateTopology",
	"NIOS.LbDynamicRatioAlternate":  "LbDynamicRatioAlternate",
	"NIOS.LbDynamicRatioPreferred":  "LbDynamicRatioPreferred",
	"NIOS.LbPreferredMethod":        "LbPreferredMethod",
	"NIOS.LbPreferredTopology":      "LbPreferredTopology",
	"NIOS.Monitors":                 "Monitors",
	"NIOS.Name":                     "Name",
	"NIOS.Quorum":                   "Quorum",
	"NIOS.Servers":                  "Servers",
	"NIOS.Ttl":                      "Ttl",
	"NIOS.UseTtl":                   "UseTtl",
}

// DtcPoolUDDIFieldMap maps infoblox model fields to UDDI struct fields
var DtcPoolUDDIFieldMap = map[string]string{
	"UDDI.Comment":                   "Comment",
	"UDDI.ConsolidatedHealthEnabled": "ConsolidatedHealthEnabled",
	"UDDI.Disabled":                  "Disabled",
	"UDDI.HealthChecks":              "HealthChecks",
	"UDDI.InheritanceSources":        "InheritanceSources",
	"UDDI.Method":                    "Method",
	"UDDI.Name":                      "Name",
	"UDDI.PoolAvailability":          "PoolAvailability",
	"UDDI.PoolServersQuorum":         "PoolServersQuorum",
	"UDDI.ServerAvailability":        "ServerAvailability",
	"UDDI.ServerHealthChecksQuorum":  "ServerHealthChecksQuorum",
	"UDDI.Servers":                   "Servers",
	"UDDI.Tags":                      "Tags",
	"UDDI.Ttl":                       "Ttl",
}

// TODO: only searchable fields should be included here
// DtcPoolFilterFieldMap maps infoblox filter keys to backend-specific API filter field names
var DtcPoolFilterFieldMap = map[core.BackendType]map[string]string{
	core.BackendNIOS: {
		"id":                              "_ref",
		"nios.auto_consolidated_monitors": "auto_consolidated_monitors",
		"nios.availability":               "availability",
		"nios.comment":                    "comment",
		"nios.consolidated_monitors":      "consolidated_monitors",
		"nios.disable":                    "disable",
		"nios.ext_attrs":                  "extattrs",
		"nios.lb_alternate_method":        "lb_alternate_method",
		"nios.lb_alternate_topology":      "lb_alternate_topology",
		"nios.lb_dynamic_ratio_alternate": "lb_dynamic_ratio_alternate",
		"nios.lb_dynamic_ratio_preferred": "lb_dynamic_ratio_preferred",
		"nios.lb_preferred_method":        "lb_preferred_method",
		"nios.lb_preferred_topology":      "lb_preferred_topology",
		"nios.monitors":                   "monitors",
		"nios.name":                       "name",
		"nios.quorum":                     "quorum",
		"nios.servers":                    "servers",
		"nios.ttl":                        "ttl",
		"nios.use_ttl":                    "use_ttl",
	},
	core.BackendUDDI: {
		"uddi.comment":                     "comment",
		"uddi.consolidated_health_enabled": "consolidated_health_enabled",
		"uddi.disabled":                    "disabled",
		"uddi.health_checks":               "health_checks",
		"uddi.inheritance_sources":         "inheritance_sources",
		"uddi.method":                      "method",
		"uddi.name":                        "name",
		"uddi.pool_availability":           "pool_availability",
		"uddi.pool_servers_quorum":         "pool_servers_quorum",
		"uddi.server_availability":         "server_availability",
		"uddi.server_health_checks_quorum": "server_health_checks_quorum",
		"uddi.servers":                     "servers",
		"uddi.tags":                        "tags",
		"uddi.ttl":                         "ttl",
	},
}
