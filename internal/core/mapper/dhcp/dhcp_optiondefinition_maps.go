package dhcp

import "github.com/infobloxopen/terraform-provider-infoblox/internal/core"

// DhcpOptiondefinitionNIOSFieldMap maps infoblox model fields to NIOS struct fields
var DhcpOptiondefinitionNIOSFieldMap = map[string]string{
	"Id":         "Ref",
	"NIOS.Code":  "Code",
	"NIOS.Name":  "Name",
	"NIOS.Space": "Space",
	"NIOS.Type":  "Type",
}

// DhcpOptiondefinitionUDDIFieldMap maps infoblox model fields to UDDI struct fields
var DhcpOptiondefinitionUDDIFieldMap = map[string]string{
	"UDDI.Array":       "Array",
	"UDDI.Code":        "Code",
	"UDDI.Comment":     "Comment",
	"UDDI.Name":        "Name",
	"UDDI.OptionSpace": "OptionSpace",
	"UDDI.Type":        "Type",
}

// TODO: only searchable fields should be included here
// DhcpOptiondefinitionFilterFieldMap maps infoblox filter keys to backend-specific API filter field names
var DhcpOptiondefinitionFilterFieldMap = map[core.BackendType]map[string]string{
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
