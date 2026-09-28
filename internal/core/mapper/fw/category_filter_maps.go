package fw

import "github.com/infobloxopen/terraform-provider-infoblox/internal/core"

// CategoryFilterUDDIFieldMap maps infoblox model fields to UDDI struct fields
var CategoryFilterUDDIFieldMap = map[string]string{
	"UDDI.Categories":  "Categories",
	"UDDI.CreatedTime": "CreatedTime",
	"UDDI.Description": "Description",
	"UDDI.Name":        "Name",
	"UDDI.Policies":    "Policies",
	"UDDI.Tags":        "Tags",
	"UDDI.UpdatedTime": "UpdatedTime",
}

// TODO: only searchable fields should be included here
// CategoryFilterFilterFieldMap maps infoblox filter keys to backend-specific API filter field names
var CategoryFilterFilterFieldMap = map[core.BackendType]map[string]string{
	core.BackendUDDI: {
		"uddi.categories":   "categories",
		"uddi.created_time": "created_time",
		"uddi.description":  "description",
		"uddi.name":         "name",
		"uddi.policies":     "policies",
		"uddi.tags":         "tags",
		"uddi.updated_time": "updated_time",
	},
}
