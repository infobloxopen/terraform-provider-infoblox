package ipamfederation

import "github.com/infobloxopen/terraform-provider-infoblox/internal/core"

// ReservedBlockUDDIFieldMap maps infoblox model fields to UDDI struct fields
var ReservedBlockUDDIFieldMap = map[string]string{
	"UDDI.Address":          "Address",
	"UDDI.Cidr":             "Cidr",
	"UDDI.Comment":          "Comment",
	"UDDI.CreatedAt":        "CreatedAt",
	"UDDI.FederatedPoolId":  "FederatedPoolId",
	"UDDI.FederatedRealm":   "FederatedRealm",
	"UDDI.Metadata":         "Metadata",
	"UDDI.Name":             "Name",
	"UDDI.NetworkCompliant": "NetworkCompliant",
	"UDDI.Parent":           "Parent",
	"UDDI.Protocol":         "Protocol",
	"UDDI.Region":           "Region",
	"UDDI.Tags":             "Tags",
	"UDDI.UpdatedAt":        "UpdatedAt",
}

// TODO: only searchable fields should be included here
// ReservedBlockFilterFieldMap maps infoblox filter keys to backend-specific API filter field names
var ReservedBlockFilterFieldMap = map[core.BackendType]map[string]string{
	core.BackendUDDI: {
		"uddi.address":           "address",
		"uddi.cidr":              "cidr",
		"uddi.comment":           "comment",
		"uddi.created_at":        "created_at",
		"uddi.federated_pool_id": "federated_pool_id",
		"uddi.federated_realm":   "federated_realm",
		"uddi.metadata":          "metadata",
		"uddi.name":              "name",
		"uddi.network_compliant": "network_compliant",
		"uddi.parent":            "parent",
		"uddi.protocol":          "protocol",
		"uddi.region":            "region",
		"uddi.tags":              "tags",
		"uddi.updated_at":        "updated_at",
	},
}
