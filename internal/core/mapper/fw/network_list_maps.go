package fw

import "github.com/infobloxopen/terraform-provider-infoblox/internal/core"

// NetworkListUDDIFieldMap maps infoblox model fields to UDDI struct fields
var NetworkListUDDIFieldMap = map[string]string{
	"UDDI.AddrBlock":   "AddrBlock",
	"UDDI.Description": "Description",
	"UDDI.Name":        "Name",
}

// TODO: only searchable fields should be included here
// NetworkListFilterFieldMap maps infoblox filter keys to backend-specific API filter field names
var NetworkListFilterFieldMap = map[core.BackendType]map[string]string{
	core.BackendUDDI: {
		"uddi.addr_block":  "addr_block",
		"uddi.description": "description",
		"uddi.name":        "name",
	},
}
