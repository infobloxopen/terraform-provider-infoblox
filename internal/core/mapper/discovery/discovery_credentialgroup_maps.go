package discovery

import "github.com/infobloxopen/terraform-provider-infoblox/internal/core"

// DiscoveryCredentialgroupNIOSFieldMap maps infoblox model fields to NIOS struct fields
var DiscoveryCredentialgroupNIOSFieldMap = map[string]string{
	"Id":        "Ref",
	"NIOS.Name": "Name",
}

// TODO: only searchable fields should be included here
// DiscoveryCredentialgroupFilterFieldMap maps infoblox filter keys to backend-specific API filter field names
var DiscoveryCredentialgroupFilterFieldMap = map[core.BackendType]map[string]string{
	core.BackendNIOS: {
		"id":        "_ref",
		"nios.name": "name",
	},
}
