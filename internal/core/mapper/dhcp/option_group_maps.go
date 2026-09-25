package dhcp

import "github.com/infobloxopen/terraform-provider-infoblox/internal/core"

// OptionGroupUDDIFieldMap maps infoblox model fields to UDDI struct fields
var OptionGroupUDDIFieldMap = map[string]string{
	"UDDI.Comment":     "Comment",
	"UDDI.DhcpOptions": "DhcpOptions",
	"UDDI.Name":        "Name",
	"UDDI.Protocol":    "Protocol",
	"UDDI.Tags":        "Tags",
}

// TODO: only searchable fields should be included here
// OptionGroupFilterFieldMap maps infoblox filter keys to backend-specific API filter field names
var OptionGroupFilterFieldMap = map[core.BackendType]map[string]string{
	core.BackendUDDI: {
		"uddi.comment":      "comment",
		"uddi.dhcp_options": "dhcp_options",
		"uddi.name":         "name",
		"uddi.protocol":     "protocol",
		"uddi.tags":         "tags",
	},
}
