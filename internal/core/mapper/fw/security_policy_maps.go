package fw

import "github.com/infobloxopen/terraform-provider-infoblox/internal/core"

// SecurityPolicyUDDIFieldMap maps infoblox model fields to UDDI struct fields
var SecurityPolicyUDDIFieldMap = map[string]string{
	"UDDI.AccessCodes":         "AccessCodes",
	"UDDI.DefaultAction":       "DefaultAction",
	"UDDI.DefaultRedirectName": "DefaultRedirectName",
	"UDDI.Description":         "Description",
	"UDDI.DfpServices":         "DfpServices",
	"UDDI.Dfps":                "Dfps",
	"UDDI.Ecs":                 "Ecs",
	"UDDI.Name":                "Name",
	"UDDI.NetAddressDfps":      "NetAddressDfps",
	"UDDI.NetworkLists":        "NetworkLists",
	"UDDI.OnpremResolve":       "OnpremResolve",
	"UDDI.Precedence":          "Precedence",
	"UDDI.RoamingDeviceGroups": "RoamingDeviceGroups",
	"UDDI.Rules":               "Rules",
	"UDDI.SafeSearch":          "SafeSearch",
	"UDDI.Tags":                "Tags",
	"UDDI.UserGroups":          "UserGroups",
}

// TODO: only searchable fields should be included here
// SecurityPolicyFilterFieldMap maps infoblox filter keys to backend-specific API filter field names
var SecurityPolicyFilterFieldMap = map[core.BackendType]map[string]string{
	core.BackendUDDI: {
		"uddi.access_codes":          "access_codes",
		"uddi.default_action":        "default_action",
		"uddi.default_redirect_name": "default_redirect_name",
		"uddi.description":           "description",
		"uddi.dfp_services":          "dfp_services",
		"uddi.dfps":                  "dfps",
		"uddi.ecs":                   "ecs",
		"uddi.name":                  "name",
		"uddi.net_address_dfps":      "net_address_dfps",
		"uddi.network_lists":         "network_lists",
		"uddi.onprem_resolve":        "onprem_resolve",
		"uddi.precedence":            "precedence",
		"uddi.roaming_device_groups": "roaming_device_groups",
		"uddi.rules":                 "rules",
		"uddi.safe_search":           "safe_search",
		"uddi.tags":                  "tags",
		"uddi.user_groups":           "user_groups",
	},
}
