package infra

import "github.com/infobloxopen/terraform-provider-infoblox/internal/core"

// InfraServiceUDDIFieldMap maps infoblox model fields to UDDI struct fields
var InfraServiceUDDIFieldMap = map[string]string{
	"UDDI.Description":     "Description",
	"UDDI.DesiredState":    "DesiredState",
	"UDDI.DesiredVersion":  "DesiredVersion",
	"UDDI.InterfaceLabels": "InterfaceLabels",
	"UDDI.Name":            "Name",
	"UDDI.PoolId":          "PoolId",
	"UDDI.ServiceType":     "ServiceType",
	"UDDI.Tags":            "Tags",
}

// TODO: only searchable fields should be included here
// InfraServiceFilterFieldMap maps infoblox filter keys to backend-specific API filter field names
var InfraServiceFilterFieldMap = map[core.BackendType]map[string]string{
	core.BackendUDDI: {
		"uddi.description":      "description",
		"uddi.desired_state":    "desired_state",
		"uddi.desired_version":  "desired_version",
		"uddi.interface_labels": "interface_labels",
		"uddi.name":             "name",
		"uddi.pool_id":          "pool_id",
		"uddi.service_type":     "service_type",
		"uddi.tags":             "tags",
	},
}
