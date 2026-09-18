package fw

import "github.com/infobloxopen/terraform-provider-infoblox/internal/core"

// ApplicationFilterUDDIFieldMap maps infoblox model fields to UDDI struct fields
var ApplicationFilterUDDIFieldMap = map[string]string{
	"UDDI.Criteria":    "Criteria",
	"UDDI.Description": "Description",
	"UDDI.Name":        "Name",
	"UDDI.Readonly":    "Readonly",
	"UDDI.Tags":        "Tags",
}

// TODO: only searchable fields should be included here
// ApplicationFilterFilterFieldMap maps infoblox filter keys to backend-specific API filter field names
var ApplicationFilterFilterFieldMap = map[core.BackendType]map[string]string{
	core.BackendUDDI: {
		"uddi.criteria":    "criteria",
		"uddi.description": "description",
		"uddi.name":        "name",
		"uddi.readonly":    "readonly",
		"uddi.tags":        "tags",
	},
}
