package ipam

import "github.com/infobloxopen/terraform-provider-infoblox/internal/core"

// IpamHostUDDIFieldMap maps infoblox model fields to UDDI struct fields
var IpamHostUDDIFieldMap = map[string]string{
	"UDDI.Addresses":           "Addresses",
	"UDDI.AutoGenerateRecords": "AutoGenerateRecords",
	"UDDI.Comment":             "Comment",
	"UDDI.HostNames":           "HostNames",
	"UDDI.Name":                "Name",
	"UDDI.Tags":                "Tags",
}

// TODO: only searchable fields should be included here
// IpamHostFilterFieldMap maps infoblox filter keys to backend-specific API filter field names
var IpamHostFilterFieldMap = map[core.BackendType]map[string]string{
	core.BackendUDDI: {
		"uddi.addresses":             "addresses",
		"uddi.auto_generate_records": "auto_generate_records",
		"uddi.comment":               "comment",
		"uddi.host_names":            "host_names",
		"uddi.name":                  "name",
		"uddi.tags":                  "tags",
	},
}
