package redirect

import "github.com/infobloxopen/terraform-provider-infoblox/internal/core"

// CustomRedirectUDDIFieldMap maps infoblox model fields to UDDI struct fields
var CustomRedirectUDDIFieldMap = map[string]string{
	"UDDI.Data":      "Data",
	"UDDI.Name":      "Name",
	"UDDI.PolicyIds": "PolicyIds",
}

// TODO: only searchable fields should be included here
// CustomRedirectFilterFieldMap maps infoblox filter keys to backend-specific API filter field names
var CustomRedirectFilterFieldMap = map[core.BackendType]map[string]string{
	core.BackendUDDI: {
		"uddi.data":       "data",
		"uddi.name":       "name",
		"uddi.policy_ids": "policy_ids",
	},
}
