package dns

import "github.com/infobloxopen/terraform-provider-infoblox/internal/core"

// NsgroupNIOSFieldMap maps infoblox model fields to NIOS struct fields
var NsgroupNIOSFieldMap = map[string]string{
	"Id":                       "Ref",
	"NIOS.Comment":             "Comment",
	"NIOS.ExternalPrimaries":   "ExternalPrimaries",
	"NIOS.ExternalSecondaries": "ExternalSecondaries",
	"NIOS.GridPrimary":         "GridPrimary",
	"NIOS.GridSecondaries":     "GridSecondaries",
	"NIOS.IsGridDefault":       "IsGridDefault",
	"NIOS.IsMultimaster":       "IsMultimaster",
	"NIOS.Name":                "Name",
	"NIOS.UseExternalPrimary":  "UseExternalPrimary",
}

// TODO: only searchable fields should be included here
// NsgroupFilterFieldMap maps infoblox filter keys to backend-specific API filter field names
var NsgroupFilterFieldMap = map[core.BackendType]map[string]string{
	core.BackendNIOS: {
		"id":                        "_ref",
		"nios.comment":              "comment",
		"nios.ext_attrs":            "extattrs",
		"nios.external_primaries":   "external_primaries",
		"nios.external_secondaries": "external_secondaries",
		"nios.grid_primary":         "grid_primary",
		"nios.grid_secondaries":     "grid_secondaries",
		"nios.is_grid_default":      "is_grid_default",
		"nios.is_multimaster":       "is_multimaster",
		"nios.name":                 "name",
		"nios.use_external_primary": "use_external_primary",
	},
}
