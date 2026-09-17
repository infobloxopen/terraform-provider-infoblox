package dns

import "github.com/infobloxopen/terraform-provider-infoblox/internal/core"

// NsgroupForwardstubserverNIOSFieldMap maps infoblox model fields to NIOS struct fields
var NsgroupForwardstubserverNIOSFieldMap = map[string]string{
	"Id":                   "Ref",
	"NIOS.Comment":         "Comment",
	"NIOS.ExternalServers": "ExternalServers",
	"NIOS.Name":            "Name",
}

// TODO: only searchable fields should be included here
// NsgroupForwardstubserverFilterFieldMap maps infoblox filter keys to backend-specific API filter field names
var NsgroupForwardstubserverFilterFieldMap = map[core.BackendType]map[string]string{
	core.BackendNIOS: {
		"id":                    "_ref",
		"nios.comment":          "comment",
		"nios.ext_attrs":        "extattrs",
		"nios.external_servers": "external_servers",
		"nios.name":             "name",
	},
}
