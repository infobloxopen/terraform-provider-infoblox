package dhcp

import "github.com/infobloxopen/terraform-provider-infoblox/internal/core"

// HardwareFilterUDDIFieldMap maps infoblox model fields to UDDI struct fields
var HardwareFilterUDDIFieldMap = map[string]string{
	"UDDI.Addresses":                       "Addresses",
	"UDDI.Comment":                         "Comment",
	"UDDI.CreatedAt":                       "CreatedAt",
	"UDDI.DhcpOptions":                     "DhcpOptions",
	"UDDI.HeaderOptionFilename":            "HeaderOptionFilename",
	"UDDI.HeaderOptionServerAddress":       "HeaderOptionServerAddress",
	"UDDI.HeaderOptionServerName":          "HeaderOptionServerName",
	"UDDI.LeaseTime":                       "LeaseTime",
	"UDDI.Name":                            "Name",
	"UDDI.Role":                            "Role",
	"UDDI.Tags":                            "Tags",
	"UDDI.UpdatedAt":                       "UpdatedAt",
	"UDDI.VendorSpecificOptionOptionSpace": "VendorSpecificOptionOptionSpace",
}

// TODO: only searchable fields should be included here
// HardwareFilterFilterFieldMap maps infoblox filter keys to backend-specific API filter field names
var HardwareFilterFilterFieldMap = map[core.BackendType]map[string]string{
	core.BackendUDDI: {
		"uddi.addresses":                           "addresses",
		"uddi.comment":                             "comment",
		"uddi.created_at":                          "created_at",
		"uddi.dhcp_options":                        "dhcp_options",
		"uddi.header_option_filename":              "header_option_filename",
		"uddi.header_option_server_address":        "header_option_server_address",
		"uddi.header_option_server_name":           "header_option_server_name",
		"uddi.lease_time":                          "lease_time",
		"uddi.name":                                "name",
		"uddi.role":                                "role",
		"uddi.tags":                                "tags",
		"uddi.updated_at":                          "updated_at",
		"uddi.vendor_specific_option_option_space": "vendor_specific_option_option_space",
	},
}
