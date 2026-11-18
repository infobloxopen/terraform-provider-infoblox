package ipamfederation

import "github.com/infobloxopen/terraform-provider-infoblox/internal/core"

// ForwardLookingDelegationUDDIFieldMap maps infoblox model fields to UDDI struct fields
var ForwardLookingDelegationUDDIFieldMap = map[string]string{
	"UDDI.Address":          "Address",
	"UDDI.Cidr":             "Cidr",
	"UDDI.Comment":          "Comment",
	"UDDI.FederatedPoolId":  "FederatedPoolId",
	"UDDI.FederatedRealms":  "FederatedRealms",
	"UDDI.Name":             "Name",
	"UDDI.NetworkCompliant": "NetworkCompliant",
	"UDDI.Protocol":         "Protocol",
	"UDDI.Tags":             "Tags",
}

// TODO: only searchable fields should be included here
// ForwardLookingDelegationFilterFieldMap maps infoblox filter keys to backend-specific API filter field names
var ForwardLookingDelegationFilterFieldMap = map[core.BackendType]map[string]string{
	core.BackendUDDI: {
		"uddi.address":           "address",
		"uddi.cidr":              "cidr",
		"uddi.comment":           "comment",
		"uddi.federated_pool_id": "federated_pool_id",
		"uddi.federated_realms":  "federated_realms",
		"uddi.name":              "name",
		"uddi.network_compliant": "network_compliant",
		"uddi.protocol":          "protocol",
		"uddi.tags":              "tags",
	},
}
