package infra

import "github.com/infobloxopen/terraform-provider-infoblox/internal/core"

// JoinTokenUDDIFieldMap maps infoblox model fields to UDDI struct fields
var JoinTokenUDDIFieldMap = map[string]string{
	"UDDI.DeletedAt":   "DeletedAt",
	"UDDI.Description": "Description",
	"UDDI.ExpiresAt":   "ExpiresAt",
	"UDDI.LastUsedAt":  "LastUsedAt",
	"UDDI.Name":        "Name",
	"UDDI.Tags":        "Tags",
	"UDDI.TokenId":     "TokenId",
	"UDDI.UseCounter":  "UseCounter",
}

// TODO: only searchable fields should be included here
// JoinTokenFilterFieldMap maps infoblox filter keys to backend-specific API filter field names
var JoinTokenFilterFieldMap = map[core.BackendType]map[string]string{
	core.BackendUDDI: {
		"uddi.deleted_at":   "deleted_at",
		"uddi.description":  "description",
		"uddi.expires_at":   "expires_at",
		"uddi.join_token":   "join_token",
		"uddi.last_used_at": "last_used_at",
		"uddi.name":         "name",
		"uddi.tags":         "tags",
		"uddi.token_id":     "token_id",
		"uddi.use_counter":  "use_counter",
	},
}
