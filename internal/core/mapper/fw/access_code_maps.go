package fw

import "github.com/infobloxopen/terraform-provider-infoblox/internal/core"

// AccessCodeUDDIFieldMap maps infoblox model fields to UDDI struct fields
var AccessCodeUDDIFieldMap = map[string]string{
	"Id":               "AccessKey",
	"UDDI.AccessKey":   "AccessKey",
	"UDDI.Activation":  "Activation",
	"UDDI.CreatedTime": "CreatedTime",
	"UDDI.Description": "Description",
	"UDDI.Expiration":  "Expiration",
	"UDDI.Name":        "Name",
	"UDDI.PolicyIds":   "PolicyIds",
	"UDDI.Rules":       "Rules",
	"UDDI.UpdatedTime": "UpdatedTime",
}

// TODO: only searchable fields should be included here
// AccessCodeFilterFieldMap maps infoblox filter keys to backend-specific API filter field names
var AccessCodeFilterFieldMap = map[core.BackendType]map[string]string{
	core.BackendUDDI: {
		"id":                "access_key",
		"uddi.access_key":   "access_key",
		"uddi.activation":   "activation",
		"uddi.created_time": "created_time",
		"uddi.description":  "description",
		"uddi.expiration":   "expiration",
		"uddi.name":         "name",
		"uddi.policy_ids":   "policy_ids",
		"uddi.rules":        "rules",
		"uddi.updated_time": "updated_time",
	},
}
