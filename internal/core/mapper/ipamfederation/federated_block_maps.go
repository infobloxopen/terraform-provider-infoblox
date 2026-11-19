package ipamfederation

import "github.com/infobloxopen/terraform-provider-infoblox/internal/core"

// FederatedBlockUDDIFieldMap maps infoblox model fields to UDDI struct fields
var FederatedBlockUDDIFieldMap = map[string]string{
	"UDDI.Address":           "Address",
	"UDDI.AllocationV4":      "AllocationV4",
	"UDDI.Cidr":              "Cidr",
	"UDDI.Comment":           "Comment",
	"UDDI.CreatedAt":         "CreatedAt",
	"UDDI.FederatedPoolId":   "FederatedPoolId",
	"UDDI.FederatedRealm":    "FederatedRealm",
	"UDDI.Metadata":          "Metadata",
	"UDDI.Name":              "Name",
	"UDDI.NetworkCompliance": "NetworkCompliance",
	"UDDI.NetworkCompliant":  "NetworkCompliant",
	"UDDI.Parent":            "Parent",
	"UDDI.Protocol":          "Protocol",
	"UDDI.Region":            "Region",
	"UDDI.State":             "State",
	"UDDI.Tags":              "Tags",
	"UDDI.UpdatedAt":         "UpdatedAt",
	"UDDI.Utilization":       "Utilization",
	"UDDI.UtilizationV6":     "UtilizationV6",
}

// TODO: only searchable fields should be included here
// FederatedBlockFilterFieldMap maps infoblox filter keys to backend-specific API filter field names
var FederatedBlockFilterFieldMap = map[core.BackendType]map[string]string{
	core.BackendUDDI: {
		"uddi.address":            "address",
		"uddi.allocation_v4":      "allocation_v4",
		"uddi.cidr":               "cidr",
		"uddi.comment":            "comment",
		"uddi.created_at":         "created_at",
		"uddi.federated_pool_id":  "federated_pool_id",
		"uddi.federated_realm":    "federated_realm",
		"uddi.metadata":           "metadata",
		"uddi.name":               "name",
		"uddi.network_compliance": "network_compliance",
		"uddi.network_compliant":  "network_compliant",
		"uddi.parent":             "parent",
		"uddi.protocol":           "protocol",
		"uddi.region":             "region",
		"uddi.state":              "state",
		"uddi.tags":               "tags",
		"uddi.updated_at":         "updated_at",
		"uddi.utilization":        "utilization",
		"uddi.utilization_v6":     "utilization_v6",
	},
}
