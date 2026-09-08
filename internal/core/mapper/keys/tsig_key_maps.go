package keys

import "github.com/infobloxopen/terraform-provider-infoblox/internal/core"

// TsigKeyUDDIFieldMap maps infoblox model fields to UDDI struct fields
var TsigKeyUDDIFieldMap = map[string]string{
	"UDDI.Algorithm": "Algorithm",
	"UDDI.Comment":   "Comment",
	"UDDI.Name":      "Name",
	"UDDI.Secret":    "Secret",
	"UDDI.Tags":      "Tags",
}

// TODO: only searchable fields should be included here
// TsigKeyFilterFieldMap maps infoblox filter keys to backend-specific API filter field names
var TsigKeyFilterFieldMap = map[core.BackendType]map[string]string{
	core.BackendUDDI: {
		"uddi.algorithm": "algorithm",
		"uddi.comment":   "comment",
		"uddi.name":      "name",
		"uddi.secret":    "secret",
		"uddi.tags":      "tags",
	},
}
