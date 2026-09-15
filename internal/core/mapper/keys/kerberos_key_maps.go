package keys

import "github.com/infobloxopen/terraform-provider-infoblox/internal/core"

// KerberosKeyUDDIFieldMap maps infoblox model fields to UDDI struct fields
var KerberosKeyUDDIFieldMap = map[string]string{
	"UDDI.Algorithm":  "Algorithm",
	"UDDI.Comment":    "Comment",
	"UDDI.Domain":     "Domain",
	"UDDI.Principal":  "Principal",
	"UDDI.Tags":       "Tags",
	"UDDI.UploadedAt": "UploadedAt",
	"UDDI.Version":    "Version",
}

// TODO: only searchable fields should be included here
// KerberosKeyFilterFieldMap maps infoblox filter keys to backend-specific API filter field names
var KerberosKeyFilterFieldMap = map[core.BackendType]map[string]string{
	core.BackendUDDI: {
		"uddi.algorithm":   "algorithm",
		"uddi.comment":     "comment",
		"uddi.domain":      "domain",
		"uddi.principal":   "principal",
		"uddi.tags":        "tags",
		"uddi.uploaded_at": "uploaded_at",
		"uddi.version":     "version",
	},
}
