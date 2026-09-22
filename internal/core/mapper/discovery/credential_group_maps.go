package discovery

import "github.com/infobloxopen/terraform-provider-infoblox/internal/core"

// CredentialGroupNIOSFieldMap maps infoblox model fields to NIOS struct fields
var CredentialGroupNIOSFieldMap = map[string]string{
	"Id":        "Ref",
	"NIOS.Name": "Name",
}

// TODO: only searchable fields should be included here
// CredentialGroupFilterFieldMap maps infoblox filter keys to backend-specific API filter field names
var CredentialGroupFilterFieldMap = map[core.BackendType]map[string]string{
	core.BackendNIOS: {
		"id":        "_ref",
		"nios.name": "name",
	},
}
