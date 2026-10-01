package anycast

import "github.com/infobloxopen/terraform-provider-infoblox/internal/core"

// AnycastConfigUDDIFieldMap maps infoblox model fields to UDDI struct fields
var AnycastConfigUDDIFieldMap = map[string]string{
	"UDDI.AnycastIpAddress":   "AnycastIpAddress",
	"UDDI.AnycastIpv6Address": "AnycastIpv6Address",
	"UDDI.Description":        "Description",
	"UDDI.Name":               "Name",
	"UDDI.OnpremHosts":        "OnpremHosts",
	"UDDI.Service":            "Service",
	"UDDI.Tags":               "Tags",
}

// TODO: only searchable fields should be included here
// AnycastConfigFilterFieldMap maps infoblox filter keys to backend-specific API filter field names
var AnycastConfigFilterFieldMap = map[core.BackendType]map[string]string{
	core.BackendUDDI: {
		"uddi.anycast_ip_address":   "anycast_ip_address",
		"uddi.anycast_ipv6_address": "anycast_ipv6_address",
		"uddi.description":          "description",
		"uddi.name":                 "name",
		"uddi.onprem_hosts":         "onprem_hosts",
		"uddi.service":              "service",
		"uddi.tags":                 "tags",
	},
}
