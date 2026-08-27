package dhcp

import "github.com/infobloxopen/terraform-provider-infoblox/internal/core"

// DhcpHostUDDIFieldMap maps infoblox model fields to UDDI struct fields.
// Only Server is writable; all other fields are system-managed read-only.
var DhcpHostUDDIFieldMap = map[string]string{
	"UDDI.Server": "Server",
}

// DhcpHostFilterFieldMap maps infoblox filter keys to backend-specific API filter field names
var DhcpHostFilterFieldMap = map[core.BackendType]map[string]string{
	core.BackendUDDI: {
		"uddi.ip_space": "ip_space",
		"uddi.name":     "name",
		"uddi.server":   "server",
	},
}
