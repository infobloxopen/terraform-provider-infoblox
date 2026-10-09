package ipam

import "github.com/infobloxopen/terraform-provider-infoblox/internal/core"

// SuperhostNIOSFieldMap maps infoblox model fields to NIOS struct fields
var SuperhostNIOSFieldMap = map[string]string{
	"Id":                           "Ref",
	"NIOS.Comment":                 "Comment",
	"NIOS.DeleteAssociatedObjects": "DeleteAssociatedObjects",
	"NIOS.DhcpAssociatedObjects":   "DhcpAssociatedObjects",
	"NIOS.Disabled":                "Disabled",
	"NIOS.DnsAssociatedObjects":    "DnsAssociatedObjects",
	"NIOS.Name":                    "Name",
}

// TODO: only searchable fields should be included here
// SuperhostFilterFieldMap maps infoblox filter keys to backend-specific API filter field names
var SuperhostFilterFieldMap = map[core.BackendType]map[string]string{
	core.BackendNIOS: {
		"id":                             "_ref",
		"nios.comment":                   "comment",
		"nios.delete_associated_objects": "delete_associated_objects",
		"nios.dhcp_associated_objects":   "dhcp_associated_objects",
		"nios.disabled":                  "disabled",
		"nios.dns_associated_objects":    "dns_associated_objects",
		"nios.ext_attrs":                 "extattrs",
		"nios.name":                      "name",
	},
}
