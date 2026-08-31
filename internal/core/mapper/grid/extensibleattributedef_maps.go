package grid

import "github.com/infobloxopen/terraform-provider-infoblox/internal/core"

// ExtensibleattributedefNIOSFieldMap maps infoblox model fields to NIOS struct fields
var ExtensibleattributedefNIOSFieldMap = map[string]string{
	"Id":                      "Ref",
	"NIOS.AllowedObjectTypes": "AllowedObjectTypes",
	"NIOS.Comment":            "Comment",
	"NIOS.DefaultValue":       "DefaultValue",
	"NIOS.Flags":              "Flags",
	"NIOS.ListValues":         "ListValues",
	"NIOS.Max":                "Max",
	"NIOS.Min":                "Min",
	"NIOS.Name":               "Name",
	"NIOS.Type":               "Type",
}

// TODO: only searchable fields should be included here
// ExtensibleattributedefFilterFieldMap maps infoblox filter keys to backend-specific API filter field names
var ExtensibleattributedefFilterFieldMap = map[core.BackendType]map[string]string{
	core.BackendNIOS: {
		"id":                        "_ref",
		"nios.allowed_object_types": "allowed_object_types",
		"nios.comment":              "comment",
		"nios.default_value":        "default_value",
		"nios.flags":                "flags",
		"nios.list_values":          "list_values",
		"nios.max":                  "max",
		"nios.min":                  "min",
		"nios.name":                 "name",
		"nios.type":                 "type",
	},
}
