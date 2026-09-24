package dhcp

import "github.com/infobloxopen/terraform-provider-infoblox/internal/core"

// Ipv6filteroptionNIOSFieldMap maps infoblox model fields to NIOS struct fields
var Ipv6filteroptionNIOSFieldMap = map[string]string{
	"Id":                "Ref",
	"NIOS.ApplyAsClass": "ApplyAsClass",
	"NIOS.Comment":      "Comment",
	"NIOS.Expression":   "Expression",
	"NIOS.LeaseTime":    "LeaseTime",
	"NIOS.Name":         "Name",
	"NIOS.OptionList":   "OptionList",
	"NIOS.OptionSpace":  "OptionSpace",
}

// Ipv6filteroptionUDDIFieldMap maps infoblox model fields to UDDI struct fields
var Ipv6filteroptionUDDIFieldMap = map[string]string{
	"UDDI.Comment":                         "Comment",
	"UDDI.DhcpOptions":                     "DhcpOptions",
	"UDDI.LeaseTime":                       "LeaseTime",
	"UDDI.Name":                            "Name",
	"UDDI.Protocol":                        "Protocol",
	"UDDI.Role":                            "Role",
	"UDDI.Rules":                           "Rules",
	"UDDI.Tags":                            "Tags",
	"UDDI.VendorSpecificOptionOptionSpace": "VendorSpecificOptionOptionSpace",
}

// TODO: only searchable fields should be included here
// Ipv6filteroptionFilterFieldMap maps infoblox filter keys to backend-specific API filter field names
var Ipv6filteroptionFilterFieldMap = map[core.BackendType]map[string]string{
	core.BackendNIOS: {
		"id":                  "_ref",
		"nios.apply_as_class": "apply_as_class",
		"nios.comment":        "comment",
		"nios.expression":     "expression",
		"nios.ext_attrs":      "extattrs",
		"nios.lease_time":     "lease_time",
		"nios.name":           "name",
		"nios.option_list":    "option_list",
		"nios.option_space":   "option_space",
	},
	core.BackendUDDI: {
		"uddi.comment":      "comment",
		"uddi.dhcp_options": "dhcp_options",
		"uddi.lease_time":   "lease_time",
		"uddi.name":         "name",
		"uddi.protocol":     "protocol",
		"uddi.role":         "role",
		"uddi.rules":        "rules",
		"uddi.tags":         "tags",
		"uddi.vendor_specific_option_option_space": "vendor_specific_option_option_space",
	},
}
