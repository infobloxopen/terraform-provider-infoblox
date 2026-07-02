package ipam

import "github.com/infobloxopen/terraform-provider-unified/internal/core"

// AddressUDDIFieldMap maps unified model fields to UDDI struct fields
var AddressUDDIFieldMap = map[string]string{
	"UDDI.Address":      "Address",
	"UDDI.Comment":      "Comment",
	"UDDI.ExternalKeys": "ExternalKeys",
	"UDDI.Host":         "Host",
	"UDDI.Hwaddr":       "Hwaddr",
	"UDDI.Interface":    "Interface",
	"UDDI.Names":        "Names",
	"UDDI.Range":        "Range",
	"UDDI.Space":        "Space",
	"UDDI.Tags":         "Tags",
}

// TODO: only searchable fields should be included here
// AddressFilterFieldMap maps unified filter keys to backend-specific API filter field names
var AddressFilterFieldMap = map[core.BackendType]map[string]string{
	core.BackendUDDI: {
		"uddi.address":       "address",
		"uddi.comment":       "comment",
		"uddi.external_keys": "external_keys",
		"uddi.host":          "host",
		"uddi.hwaddr":        "hwaddr",
		"uddi.interface":     "interface",
		"uddi.names":         "names",
		"uddi.range":         "range",
		"uddi.space":         "space",
		"uddi.tags":          "tags",
	},
}
