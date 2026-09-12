package dns

import "github.com/infobloxopen/terraform-provider-infoblox/internal/core"

// NsgroupDelegationNIOSFieldMap maps infoblox model fields to NIOS struct fields
var NsgroupDelegationNIOSFieldMap = map[string]string{
	"Id":              "Ref",
	"NIOS.Comment":    "Comment",
	"NIOS.DelegateTo": "DelegateTo",
	"NIOS.Name":       "Name",
}

// TODO: only searchable fields should be included here
// NsgroupDelegationFilterFieldMap maps infoblox filter keys to backend-specific API filter field names
var NsgroupDelegationFilterFieldMap = map[core.BackendType]map[string]string{
	core.BackendNIOS: {
		"id":               "_ref",
		"nios.comment":     "comment",
		"nios.delegate_to": "delegate_to",
		"nios.ext_attrs":   "extattrs",
		"nios.name":        "name",
	},
}
