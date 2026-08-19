package fw

import "github.com/infobloxopen/terraform-provider-infoblox/internal/core"

// NamedListUDDIFieldMap maps infoblox model fields to UDDI struct fields
var NamedListUDDIFieldMap = map[string]string{
	"UDDI.ConfidenceLevel": "ConfidenceLevel",
	"UDDI.Description":     "Description",
	"UDDI.Items":           "Items",
	"UDDI.ItemsDescribed":  "ItemsDescribed",
	"UDDI.Name":            "Name",
	"UDDI.Policies":        "Policies",
	"UDDI.Tags":            "Tags",
	"UDDI.ThreatLevel":     "ThreatLevel",
	"UDDI.Type":            "Type",
}

// TODO: only searchable fields should be included here
// NamedListFilterFieldMap maps infoblox filter keys to backend-specific API filter field names
var NamedListFilterFieldMap = map[core.BackendType]map[string]string{
	core.BackendUDDI: {
		"uddi.confidence_level": "confidence_level",
		"uddi.description":      "description",
		"uddi.items":            "items",
		"uddi.items_described":  "items_described",
		"uddi.name":             "name",
		"uddi.policies":         "policies",
		"uddi.tags":             "tags",
		"uddi.threat_level":     "threat_level",
		"uddi.type":             "type",
	},
}
