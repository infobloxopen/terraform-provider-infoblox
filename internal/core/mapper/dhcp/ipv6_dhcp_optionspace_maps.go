package dhcp

import "github.com/infobloxopen/terraform-provider-infoblox/internal/core"

// Ipv6DhcpOptionspaceNIOSFieldMap maps infoblox model fields to NIOS struct fields
var Ipv6DhcpOptionspaceNIOSFieldMap = map[string]string{
	"Id":                    "Ref",
	"NIOS.Comment":          "Comment",
	"NIOS.EnterpriseNumber": "EnterpriseNumber",
	"NIOS.Name":             "Name",
}

// Ipv6DhcpOptionspaceUDDIFieldMap maps infoblox model fields to UDDI struct fields
var Ipv6DhcpOptionspaceUDDIFieldMap = map[string]string{
	"UDDI.Comment":  "Comment",
	"UDDI.Name":     "Name",
	"UDDI.Protocol": "Protocol",
	"UDDI.Tags":     "Tags",
}

// TODO: only searchable fields should be included here
// Ipv6DhcpOptionspaceFilterFieldMap maps infoblox filter keys to backend-specific API filter field names
var Ipv6DhcpOptionspaceFilterFieldMap = map[core.BackendType]map[string]string{
	core.BackendNIOS: {
		"id":                     "_ref",
		"nios.comment":           "comment",
		"nios.enterprise_number": "enterprise_number",
		"nios.name":              "name",
	},
	core.BackendUDDI: {
		"uddi.comment":  "comment",
		"uddi.name":     "name",
		"uddi.protocol": "protocol",
		"uddi.tags":     "tags",
	},
}
