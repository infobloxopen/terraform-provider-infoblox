package dhcp

import "github.com/infobloxopen/terraform-provider-infoblox/internal/core"

// DhcpOptionspaceNIOSFieldMap maps infoblox model fields to NIOS struct fields
var DhcpOptionspaceNIOSFieldMap = map[string]string{
	"Id":           "Ref",
	"NIOS.Comment": "Comment",
	"NIOS.Name":    "Name",
}

// DhcpOptionspaceUDDIFieldMap maps infoblox model fields to UDDI struct fields
var DhcpOptionspaceUDDIFieldMap = map[string]string{
	"UDDI.Comment":  "Comment",
	"UDDI.Name":     "Name",
	"UDDI.Protocol": "Protocol",
	"UDDI.Tags":     "Tags",
}

// TODO: only searchable fields should be included here
// DhcpOptionspaceFilterFieldMap maps infoblox filter keys to backend-specific API filter field names
var DhcpOptionspaceFilterFieldMap = map[core.BackendType]map[string]string{
	core.BackendNIOS: {
		"id":           "_ref",
		"nios.comment": "comment",
		"nios.name":    "name",
	},
	core.BackendUDDI: {
		"uddi.comment":  "comment",
		"uddi.name":     "name",
		"uddi.protocol": "protocol",
		"uddi.tags":     "tags",
	},
}
