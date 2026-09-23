package infra

import "github.com/infobloxopen/terraform-provider-infoblox/internal/core"

// InfraServiceUDDIFieldMap maps infoblox model fields to UDDI struct fields
var InfraServiceUDDIFieldMap = map[string]string{
	"UDDI.Configs":         "Configs",
	"UDDI.CreatedAt":       "CreatedAt",
	"UDDI.Description":     "Description",
	"UDDI.DesiredState":    "DesiredState",
	"UDDI.DesiredVersion":  "DesiredVersion",
	"UDDI.InterfaceLabels": "InterfaceLabels",
	"UDDI.Name":            "Name",
	"UDDI.PoolId":          "PoolId",
	"UDDI.ServiceType":     "ServiceType",
	"UDDI.Tags":            "Tags",
	"UDDI.UpdatedAt":       "UpdatedAt",
}

// TODO: only searchable fields should be included here
// InfraServiceFilterFieldMap maps infoblox filter keys to backend-specific API filter field names
var InfraServiceFilterFieldMap = map[core.BackendType]map[string]string{
	core.BackendUDDI: {
		"uddi.configs":          "configs",
		"uddi.created_at":       "created_at",
		"uddi.description":      "description",
		"uddi.desired_state":    "desired_state",
		"uddi.desired_version":  "desired_version",
		"uddi.interface_labels": "interface_labels",
		"uddi.name":             "name",
		"uddi.pool_id":          "pool_id",
		"uddi.service_type":     "service_type",
		"uddi.tags":             "tags",
		"uddi.updated_at":       "updated_at",
	},
}
