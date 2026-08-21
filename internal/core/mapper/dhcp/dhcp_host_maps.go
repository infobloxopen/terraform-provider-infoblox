package dhcp

import "github.com/infobloxopen/terraform-provider-infoblox/internal/core"

// DhcpHostUDDIFieldMap maps infoblox model fields to UDDI struct fields.
// Tags are omitted because the DhcpHost Update API rejects them as read-only.
var DhcpHostUDDIFieldMap = map[string]string{
	"UDDI.IpSpace": "IpSpace",
	"UDDI.Server":  "Server",
}

// DhcpHostFilterFieldMap maps infoblox filter keys to backend-specific API filter field names
var DhcpHostFilterFieldMap = map[core.BackendType]map[string]string{
	core.BackendUDDI: {
		"uddi.ip_space": "ip_space",
		"uddi.server":   "server",
	},
}
