package anycast

import "github.com/infobloxopen/terraform-provider-infoblox/internal/core"

// AnycastConfigUDDIFieldMap maps infoblox model fields to UDDI struct fields
var AnycastConfigUDDIFieldMap = map[string]string{
	"UDDI.AnycastIpAddress":   "AnycastIpAddress",
	"UDDI.AnycastIpv6Address": "AnycastIpv6Address",
	"UDDI.CreatedAt":          "CreatedAt",
	"UDDI.Description":        "Description",
	"UDDI.Fields":             "Fields",
	"UDDI.IsConfigured":       "IsConfigured",
	"UDDI.Name":               "Name",
	"UDDI.OnpremHosts":        "OnpremHosts",
	"UDDI.RuntimeStatus":      "RuntimeStatus",
	"UDDI.Service":            "Service",
	"UDDI.Tags":               "Tags",
	"UDDI.UpdatedAt":          "UpdatedAt",
}

// TODO: only searchable fields should be included here
// AnycastConfigFilterFieldMap maps infoblox filter keys to backend-specific API filter field names
var AnycastConfigFilterFieldMap = map[core.BackendType]map[string]string{
	core.BackendUDDI: {
		"uddi.anycast_ip_address":   "anycast_ip_address",
		"uddi.anycast_ipv6_address": "anycast_ipv6_address",
		"uddi.created_at":           "created_at",
		"uddi.description":          "description",
		"uddi.fields":               "fields",
		"uddi.is_configured":        "is_configured",
		"uddi.name":                 "name",
		"uddi.onprem_hosts":         "onprem_hosts",
		"uddi.runtime_status":       "runtime_status",
		"uddi.service":              "service",
		"uddi.tags":                 "tags",
		"uddi.updated_at":           "updated_at",
	},
}
