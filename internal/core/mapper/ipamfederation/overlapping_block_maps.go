package ipamfederation

import "github.com/infobloxopen/terraform-provider-infoblox/internal/core"

// OverlappingBlockUDDIFieldMap maps infoblox model fields to UDDI struct fields
var OverlappingBlockUDDIFieldMap = map[string]string{
	"UDDI.Address":          "Address",
	"UDDI.Cidr":             "Cidr",
	"UDDI.Comment":          "Comment",
	"UDDI.FederatedPoolId":  "FederatedPoolId",
	"UDDI.FederatedRealm":   "FederatedRealm",
	"UDDI.Name":             "Name",
	"UDDI.NetworkCompliant": "NetworkCompliant",
	"UDDI.Parent":           "Parent",
	"UDDI.Protocol":         "Protocol",
	"UDDI.Tags":             "Tags",
}

// TODO: only searchable fields should be included here
// OverlappingBlockFilterFieldMap maps infoblox filter keys to backend-specific API filter field names
var OverlappingBlockFilterFieldMap = map[core.BackendType]map[string]string{
	core.BackendUDDI: {
		"uddi.address":           "address",
		"uddi.cidr":              "cidr",
		"uddi.comment":           "comment",
		"uddi.federated_pool_id": "federated_pool_id",
		"uddi.federated_realm":   "federated_realm",
		"uddi.name":              "name",
		"uddi.network_compliant": "network_compliant",
		"uddi.parent":            "parent",
		"uddi.protocol":          "protocol",
		"uddi.tags":              "tags",
	},
}
