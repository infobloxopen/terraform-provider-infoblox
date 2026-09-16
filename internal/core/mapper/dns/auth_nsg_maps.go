package dns

import "github.com/infobloxopen/terraform-provider-infoblox/internal/core"

// AuthNsgUDDIFieldMap maps infoblox model fields to UDDI struct fields
var AuthNsgUDDIFieldMap = map[string]string{
	"UDDI.Comment":             "Comment",
	"UDDI.ExternalPrimaries":   "ExternalPrimaries",
	"UDDI.ExternalSecondaries": "ExternalSecondaries",
	"UDDI.InternalSecondaries": "InternalSecondaries",
	"UDDI.Name":                "Name",
	"UDDI.Nsgs":                "Nsgs",
	"UDDI.Tags":                "Tags",
}

// TODO: only searchable fields should be included here
// AuthNsgFilterFieldMap maps infoblox filter keys to backend-specific API filter field names
var AuthNsgFilterFieldMap = map[core.BackendType]map[string]string{
	core.BackendUDDI: {
		"uddi.comment":              "comment",
		"uddi.external_primaries":   "external_primaries",
		"uddi.external_secondaries": "external_secondaries",
		"uddi.internal_secondaries": "internal_secondaries",
		"uddi.name":                 "name",
		"uddi.nsgs":                 "nsgs",
		"uddi.tags":                 "tags",
	},
}
