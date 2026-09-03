package ipam

import "github.com/infobloxopen/terraform-provider-infoblox/internal/core"

// BulkhostnametemplateNIOSFieldMap maps infoblox model fields to NIOS struct fields
var BulkhostnametemplateNIOSFieldMap = map[string]string{
	"Id":                  "Ref",
	"NIOS.TemplateFormat": "TemplateFormat",
	"NIOS.TemplateName":   "TemplateName",
}

// TODO: only searchable fields should be included here
// BulkhostnametemplateFilterFieldMap maps infoblox filter keys to backend-specific API filter field names
var BulkhostnametemplateFilterFieldMap = map[core.BackendType]map[string]string{
	core.BackendNIOS: {
		"id":                   "_ref",
		"nios.template_format": "template_format",
		"nios.template_name":   "template_name",
	},
}
