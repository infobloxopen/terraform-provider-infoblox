package ipamfederation

import "github.com/infobloxopen/terraform-provider-infoblox/internal/core"

// FederatedRealmUDDIFieldMap maps infoblox model fields to UDDI struct fields
var FederatedRealmUDDIFieldMap = map[string]string{
	"UDDI.Comment": "Comment",
	"UDDI.Name":    "Name",
	"UDDI.Tags":    "Tags",
}

// TODO: only searchable fields should be included here
// FederatedRealmFilterFieldMap maps infoblox filter keys to backend-specific API filter field names
var FederatedRealmFilterFieldMap = map[core.BackendType]map[string]string{
	core.BackendUDDI: {
		"uddi.comment": "comment",
		"uddi.name":    "name",
		"uddi.tags":    "tags",
	},
}
