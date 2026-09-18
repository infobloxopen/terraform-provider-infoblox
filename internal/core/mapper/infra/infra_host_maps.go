package infra

import "github.com/infobloxopen/terraform-provider-infoblox/internal/core"

// InfraHostUDDIFieldMap maps infoblox model fields to UDDI struct fields
var InfraHostUDDIFieldMap = map[string]string{
	"UDDI.Description":     "Description",
	"UDDI.DisplayName":     "DisplayName",
	"UDDI.IpSpace":         "IpSpace",
	"UDDI.LocationId":      "LocationId",
	"UDDI.MaintenanceMode": "MaintenanceMode",
	"UDDI.PoolId":          "PoolId",
	"UDDI.SerialNumber":    "SerialNumber",
	"UDDI.Tags":            "Tags",
}

// TODO: only searchable fields should be included here
// InfraHostFilterFieldMap maps infoblox filter keys to backend-specific API filter field names
var InfraHostFilterFieldMap = map[core.BackendType]map[string]string{
	core.BackendUDDI: {
		"uddi.description":      "description",
		"uddi.display_name":     "display_name",
		"uddi.ip_space":         "ip_space",
		"uddi.location_id":      "location_id",
		"uddi.maintenance_mode": "maintenance_mode",
		"uddi.pool_id":          "pool_id",
		"uddi.serial_number":    "serial_number",
		"uddi.tags":             "tags",
	},
}
