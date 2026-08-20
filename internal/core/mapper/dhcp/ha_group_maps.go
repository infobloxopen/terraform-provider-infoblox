package dhcp

import "github.com/infobloxopen/terraform-provider-infoblox/internal/core"

// HaGroupUDDIFieldMap maps infoblox model fields to UDDI struct fields
var HaGroupUDDIFieldMap = map[string]string{
	"UDDI.AnycastConfigId": "AnycastConfigId",
	"UDDI.Comment":         "Comment",
	"UDDI.Hosts":           "Hosts",
	"UDDI.IpSpace":         "IpSpace",
	"UDDI.Mode":            "Mode",
	"UDDI.Name":            "Name",
	"UDDI.Tags":            "Tags",
}

// TODO: only searchable fields should be included here
// HaGroupFilterFieldMap maps infoblox filter keys to backend-specific API filter field names
var HaGroupFilterFieldMap = map[core.BackendType]map[string]string{
	core.BackendUDDI: {
		"uddi.anycast_config_id": "anycast_config_id",
		"uddi.comment":           "comment",
		"uddi.hosts":             "hosts",
		"uddi.ip_space":          "ip_space",
		"uddi.mode":              "mode",
		"uddi.name":              "name",
		"uddi.tags":              "tags",
	},
}
