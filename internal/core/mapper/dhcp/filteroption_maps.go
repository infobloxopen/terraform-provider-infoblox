package dhcp

import "github.com/infobloxopen/terraform-provider-infoblox/internal/core"

// FilteroptionNIOSFieldMap maps infoblox model fields to NIOS struct fields
var FilteroptionNIOSFieldMap = map[string]string{
	"Id":                "Ref",
	"NIOS.ApplyAsClass": "ApplyAsClass",
	"NIOS.Bootfile":     "Bootfile",
	"NIOS.Bootserver":   "Bootserver",
	"NIOS.Comment":      "Comment",
	"NIOS.Expression":   "Expression",
	"NIOS.LeaseTime":    "LeaseTime",
	"NIOS.Name":         "Name",
	"NIOS.NextServer":   "NextServer",
	"NIOS.OptionList":   "OptionList",
	"NIOS.OptionSpace":  "OptionSpace",
	"NIOS.PxeLeaseTime": "PxeLeaseTime",
}

// FilteroptionUDDIFieldMap maps infoblox model fields to UDDI struct fields
var FilteroptionUDDIFieldMap = map[string]string{
	"UDDI.Comment":                         "Comment",
	"UDDI.DhcpOptions":                     "DhcpOptions",
	"UDDI.HeaderOptionFilename":            "HeaderOptionFilename",
	"UDDI.HeaderOptionServerAddress":       "HeaderOptionServerAddress",
	"UDDI.HeaderOptionServerName":          "HeaderOptionServerName",
	"UDDI.LeaseTime":                       "LeaseTime",
	"UDDI.Name":                            "Name",
	"UDDI.Protocol":                        "Protocol",
	"UDDI.Role":                            "Role",
	"UDDI.Rules":                           "Rules",
	"UDDI.Tags":                            "Tags",
	"UDDI.VendorSpecificOptionOptionSpace": "VendorSpecificOptionOptionSpace",
}

// TODO: only searchable fields should be included here
// FilteroptionFilterFieldMap maps infoblox filter keys to backend-specific API filter field names
var FilteroptionFilterFieldMap = map[core.BackendType]map[string]string{
	core.BackendNIOS: {
		"id":                  "_ref",
		"nios.apply_as_class": "apply_as_class",
		"nios.bootfile":       "bootfile",
		"nios.bootserver":     "bootserver",
		"nios.comment":        "comment",
		"nios.expression":     "expression",
		"nios.ext_attrs":      "extattrs",
		"nios.lease_time":     "lease_time",
		"nios.name":           "name",
		"nios.next_server":    "next_server",
		"nios.option_list":    "option_list",
		"nios.option_space":   "option_space",
		"nios.pxe_lease_time": "pxe_lease_time",
	},
	core.BackendUDDI: {
		"uddi.comment":                             "comment",
		"uddi.dhcp_options":                        "dhcp_options",
		"uddi.header_option_filename":              "header_option_filename",
		"uddi.header_option_server_address":        "header_option_server_address",
		"uddi.header_option_server_name":           "header_option_server_name",
		"uddi.lease_time":                          "lease_time",
		"uddi.name":                                "name",
		"uddi.protocol":                            "protocol",
		"uddi.role":                                "role",
		"uddi.rules":                               "rules",
		"uddi.tags":                                "tags",
		"uddi.vendor_specific_option_option_space": "vendor_specific_option_option_space",
	},
}
