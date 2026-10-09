package dhcp

import "github.com/infobloxopen/terraform-provider-infoblox/internal/core"

// Ipv6DhcpOptiondefinitionNIOSFieldMap maps infoblox model fields to NIOS struct fields
var Ipv6DhcpOptiondefinitionNIOSFieldMap = map[string]string{
	"Id":         "Ref",
	"NIOS.Code":  "Code",
	"NIOS.Name":  "Name",
	"NIOS.Space": "Space",
	"NIOS.Type":  "Type",
}

// Ipv6DhcpOptiondefinitionUDDIFieldMap maps infoblox model fields to UDDI struct fields
var Ipv6DhcpOptiondefinitionUDDIFieldMap = map[string]string{
	"UDDI.Array":       "Array",
	"UDDI.Code":        "Code",
	"UDDI.Comment":     "Comment",
	"UDDI.Name":        "Name",
	"UDDI.OptionSpace": "OptionSpace",
	"UDDI.Type":        "Type",
}

// TODO: only searchable fields should be included here
// Ipv6DhcpOptiondefinitionFilterFieldMap maps infoblox filter keys to backend-specific API filter field names
var Ipv6DhcpOptiondefinitionFilterFieldMap = map[core.BackendType]map[string]string{
	core.BackendNIOS: {
		"id":         "_ref",
		"nios.code":  "code",
		"nios.name":  "name",
		"nios.space": "space",
		"nios.type":  "type",
	},
	core.BackendUDDI: {
		"uddi.array":        "array",
		"uddi.code":         "code",
		"uddi.comment":      "comment",
		"uddi.name":         "name",
		"uddi.option_space": "option_space",
		"uddi.type":         "type",
	},
}
