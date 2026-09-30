package fw

import "github.com/infobloxopen/terraform-provider-infoblox/internal/core"

// CategoryFilterUDDIFieldMap maps infoblox model fields to UDDI struct fields
var CategoryFilterUDDIFieldMap = map[string]string{
	"UDDI.Categories":  "Categories",
	"UDDI.Description": "Description",
	"UDDI.Name":        "Name",
	"UDDI.Tags":        "Tags",
}

// TODO: only searchable fields should be included here
// CategoryFilterFilterFieldMap maps infoblox filter keys to backend-specific API filter field names
var CategoryFilterFilterFieldMap = map[core.BackendType]map[string]string{
	core.BackendUDDI: {
		"uddi.categories":  "categories",
		"uddi.description": "description",
		"uddi.name":        "name",
		"uddi.tags":        "tags",
	},
}
