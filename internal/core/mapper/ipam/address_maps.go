package ipam

import "github.com/infobloxopen/terraform-provider-infoblox/internal/core"

// AddressUDDIFieldMap maps infoblox model fields to UDDI struct fields
var AddressUDDIFieldMap = map[string]string{
	"UDDI.Address":      "Address",
	"UDDI.Comment":      "Comment",
	"UDDI.ExternalKeys": "ExternalKeys",
	"UDDI.Hwaddr":       "Hwaddr",
	"UDDI.Interface":    "Interface",
	"UDDI.Names":        "Names",
	"UDDI.Space":        "Space",
	"UDDI.Tags":         "Tags",
}

// TODO: only searchable fields should be included here
// AddressFilterFieldMap maps infoblox filter keys to backend-specific API filter field names
var AddressFilterFieldMap = map[core.BackendType]map[string]string{
	core.BackendUDDI: {
		"uddi.address":       "address",
		"uddi.comment":       "comment",
		"uddi.external_keys": "external_keys",
		"uddi.hwaddr":        "hwaddr",
		"uddi.interface":     "interface",
		"uddi.names":         "names",
		"uddi.space":         "space",
		"uddi.tags":          "tags",
	},
}
