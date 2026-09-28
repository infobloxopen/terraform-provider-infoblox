package infra

import "github.com/infobloxopen/terraform-provider-infoblox/internal/core"

// InfraHostUDDIFieldMap maps infoblox model fields to UDDI struct fields
var InfraHostUDDIFieldMap = map[string]string{
	"UDDI.Configs":         "Configs",
	"UDDI.Description":     "Description",
	"UDDI.DisplayName":     "DisplayName",
	"UDDI.HostType":        "HostType",
	"UDDI.IpAddress":       "IpAddress",
	"UDDI.IpSpace":         "IpSpace",
	"UDDI.LegacyId":        "LegacyId",
	"UDDI.LocationId":      "LocationId",
	"UDDI.MacAddress":      "MacAddress",
	"UDDI.MaintenanceMode": "MaintenanceMode",
	"UDDI.Ophid":           "Ophid",
	"UDDI.PoolId":          "PoolId",
	"UDDI.SerialNumber":    "SerialNumber",
	"UDDI.Tags":            "Tags",
	"UDDI.Timezone":        "Timezone",
}

// TODO: only searchable fields should be included here
// InfraHostFilterFieldMap maps infoblox filter keys to backend-specific API filter field names
var InfraHostFilterFieldMap = map[core.BackendType]map[string]string{
	core.BackendUDDI: {
		"uddi.configs":          "configs",
		"uddi.description":      "description",
		"uddi.display_name":     "display_name",
		"uddi.host_type":        "host_type",
		"uddi.ip_address":       "ip_address",
		"uddi.ip_space":         "ip_space",
		"uddi.legacy_id":        "legacy_id",
		"uddi.location_id":      "location_id",
		"uddi.mac_address":      "mac_address",
		"uddi.maintenance_mode": "maintenance_mode",
		"uddi.ophid":            "ophid",
		"uddi.pool_id":          "pool_id",
		"uddi.serial_number":    "serial_number",
		"uddi.tags":             "tags",
		"uddi.timezone":         "timezone",
	},
}
