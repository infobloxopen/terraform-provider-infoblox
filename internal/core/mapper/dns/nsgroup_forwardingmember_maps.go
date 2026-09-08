package dns

import "github.com/infobloxopen/terraform-provider-infoblox/internal/core"

// NsgroupForwardingmemberNIOSFieldMap maps infoblox model fields to NIOS struct fields
var NsgroupForwardingmemberNIOSFieldMap = map[string]string{
	"Id":                     "Ref",
	"NIOS.Comment":           "Comment",
	"NIOS.ForwardingServers": "ForwardingServers",
	"NIOS.Name":              "Name",
}

// TODO: only searchable fields should be included here
// NsgroupForwardingmemberFilterFieldMap maps infoblox filter keys to backend-specific API filter field names
var NsgroupForwardingmemberFilterFieldMap = map[core.BackendType]map[string]string{
	core.BackendNIOS: {
		"id":                      "_ref",
		"nios.comment":            "comment",
		"nios.ext_attrs":          "extattrs",
		"nios.forwarding_servers": "forwarding_servers",
		"nios.name":               "name",
	},
}
