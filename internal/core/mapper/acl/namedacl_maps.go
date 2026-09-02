package acl

import "github.com/infobloxopen/terraform-provider-infoblox/internal/core"

// NamedaclNIOSFieldMap maps infoblox model fields to NIOS struct fields
var NamedaclNIOSFieldMap = map[string]string{
	"Id":              "Ref",
	"NIOS.AccessList": "AccessList",
	"NIOS.Comment":    "Comment",
	"NIOS.Name":       "Name",
}

// NamedaclUDDIFieldMap maps infoblox model fields to UDDI struct fields
var NamedaclUDDIFieldMap = map[string]string{
	"UDDI.Comment":       "Comment",
	"UDDI.CompartmentId": "CompartmentId",
	"UDDI.List":          "List",
	"UDDI.Name":          "Name",
	"UDDI.Tags":          "Tags",
}

// TODO: only searchable fields should be included here
// NamedaclFilterFieldMap maps infoblox filter keys to backend-specific API filter field names
var NamedaclFilterFieldMap = map[core.BackendType]map[string]string{
	core.BackendNIOS: {
		"id":               "_ref",
		"nios.access_list": "access_list",
		"nios.comment":     "comment",
		"nios.ext_attrs":   "extattrs",
		"nios.name":        "name",
	},
	core.BackendUDDI: {
		"uddi.comment":        "comment",
		"uddi.compartment_id": "compartment_id",
		"uddi.list":           "list",
		"uddi.name":           "name",
		"uddi.tags":           "tags",
	},
}
