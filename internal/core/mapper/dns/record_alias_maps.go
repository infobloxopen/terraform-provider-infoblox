package dns

import "github.com/infobloxopen/terraform-provider-infoblox/internal/core"

// RecordAliasNIOSFieldMap maps infoblox model fields to NIOS struct fields
var RecordAliasNIOSFieldMap = map[string]string{
	"Id":              "Ref",
	"NIOS.Comment":    "Comment",
	"NIOS.Creator":    "Creator",
	"NIOS.Disable":    "Disable",
	"NIOS.Name":       "Name",
	"NIOS.TargetName": "TargetName",
	"NIOS.TargetType": "TargetType",
	"NIOS.Ttl":        "Ttl",
	"NIOS.UseTtl":     "UseTtl",
	"NIOS.View":       "View",
}

// TODO: only searchable fields should be included here
// RecordAliasFilterFieldMap maps infoblox filter keys to backend-specific API filter field names
var RecordAliasFilterFieldMap = map[core.BackendType]map[string]string{
	core.BackendNIOS: {
		"id":               "_ref",
		"nios.comment":     "comment",
		"nios.creator":     "creator",
		"nios.disable":     "disable",
		"nios.ext_attrs":   "extattrs",
		"nios.name":        "name",
		"nios.target_name": "target_name",
		"nios.target_type": "target_type",
		"nios.ttl":         "ttl",
		"nios.use_ttl":     "use_ttl",
		"nios.view":        "view",
	},
}
