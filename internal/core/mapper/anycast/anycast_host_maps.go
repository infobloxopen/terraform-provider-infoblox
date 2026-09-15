package anycast

import "github.com/infobloxopen/terraform-provider-infoblox/internal/core"

// AnycastHostUDDIFieldMap maps infoblox model fields to UDDI struct fields
var AnycastHostUDDIFieldMap = map[string]string{
	"UDDI.AnycastConfigRefs": "AnycastConfigRefs",
	"UDDI.ConfigBgp":         "ConfigBgp",
	"UDDI.ConfigOspf":        "ConfigOspf",
	"UDDI.ConfigOspfv3":      "ConfigOspfv3",
	"UDDI.CreatedAt":         "CreatedAt",
	"UDDI.IpAddress":         "IpAddress",
	"UDDI.Ipv6Address":       "Ipv6Address",
	"UDDI.Name":              "Name",
	"UDDI.UpdatedAt":         "UpdatedAt",
}

// TODO: only searchable fields should be included here
// AnycastHostFilterFieldMap maps infoblox filter keys to backend-specific API filter field names
var AnycastHostFilterFieldMap = map[core.BackendType]map[string]string{
	core.BackendUDDI: {
		"uddi.anycast_config_refs": "anycast_config_refs",
		"uddi.config_bgp":          "config_bgp",
		"uddi.config_ospf":         "config_ospf",
		"uddi.config_ospfv3":       "config_ospfv3",
		"uddi.created_at":          "created_at",
		"uddi.ip_address":          "ip_address",
		"uddi.ipv6_address":        "ipv6_address",
		"uddi.name":                "name",
		"uddi.updated_at":          "updated_at",
	},
}
