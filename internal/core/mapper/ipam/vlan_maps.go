package ipam

import "github.com/infobloxopen/terraform-provider-infoblox/internal/core"

// VlanNIOSFieldMap maps infoblox model fields to NIOS struct fields
var VlanNIOSFieldMap = map[string]string{
	"Id":               "Ref",
	"NIOS.Comment":     "Comment",
	"NIOS.Contact":     "Contact",
	"NIOS.Department":  "Department",
	"NIOS.Description": "Description",
	"NIOS.Id":          "Id",
	"NIOS.Name":        "Name",
	"NIOS.Parent":      "Parent",
	"NIOS.Reserved":    "Reserved",
}

// TODO: only searchable fields should be included here
// VlanFilterFieldMap maps infoblox filter keys to backend-specific API filter field names
var VlanFilterFieldMap = map[core.BackendType]map[string]string{
	core.BackendNIOS: {
		"id":               "_ref",
		"nios.comment":     "comment",
		"nios.contact":     "contact",
		"nios.department":  "department",
		"nios.description": "description",
		"nios.ext_attrs":   "extattrs",
		"nios.id":          "id",
		"nios.name":        "name",
		"nios.parent":      "parent",
		"nios.reserved":    "reserved",
	},
}
