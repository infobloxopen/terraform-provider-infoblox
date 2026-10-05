package fw

import "github.com/infobloxopen/terraform-provider-infoblox/internal/core"

// InternalDomainListUDDIFieldMap maps infoblox model fields to UDDI struct fields
var InternalDomainListUDDIFieldMap = map[string]string{
	"UDDI.Description":     "Description",
	"UDDI.InternalDomains": "InternalDomains",
	"UDDI.IsDefault":       "IsDefault",
	"UDDI.Name":            "Name",
	"UDDI.Tags":            "Tags",
}

// TODO: only searchable fields should be included here
// InternalDomainListFilterFieldMap maps infoblox filter keys to backend-specific API filter field names
var InternalDomainListFilterFieldMap = map[core.BackendType]map[string]string{
	core.BackendUDDI: {
		"uddi.description":      "description",
		"uddi.internal_domains": "internal_domains",
		"uddi.is_default":       "is_default",
		"uddi.name":             "name",
		"uddi.tags":             "tags",
	},
}
