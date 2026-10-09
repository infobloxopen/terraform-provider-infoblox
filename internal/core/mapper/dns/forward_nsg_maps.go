package dns

import "github.com/infobloxopen/terraform-provider-infoblox/internal/core"

// ForwardNsgUDDIFieldMap maps infoblox model fields to UDDI struct fields
var ForwardNsgUDDIFieldMap = map[string]string{
	"UDDI.Comment":            "Comment",
	"UDDI.ExternalForwarders": "ExternalForwarders",
	"UDDI.ForwardersOnly":     "ForwardersOnly",
	"UDDI.Hosts":              "Hosts",
	"UDDI.InternalForwarders": "InternalForwarders",
	"UDDI.Name":               "Name",
	"UDDI.Nsgs":               "Nsgs",
	"UDDI.Tags":               "Tags",
}

// TODO: only searchable fields should be included here
// ForwardNsgFilterFieldMap maps infoblox filter keys to backend-specific API filter field names
var ForwardNsgFilterFieldMap = map[core.BackendType]map[string]string{
	core.BackendUDDI: {
		"uddi.comment":             "comment",
		"uddi.external_forwarders": "external_forwarders",
		"uddi.forwarders_only":     "forwarders_only",
		"uddi.hosts":               "hosts",
		"uddi.internal_forwarders": "internal_forwarders",
		"uddi.name":                "name",
		"uddi.nsgs":                "nsgs",
		"uddi.tags":                "tags",
	},
}
