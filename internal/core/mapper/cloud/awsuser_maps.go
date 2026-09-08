package cloud

import "github.com/infobloxopen/terraform-provider-infoblox/internal/core"

// AwsuserNIOSFieldMap maps infoblox model fields to NIOS struct fields
var AwsuserNIOSFieldMap = map[string]string{
	"Id":                   "Ref",
	"NIOS.AccessKeyId":     "AccessKeyId",
	"NIOS.AccountId":       "AccountId",
	"NIOS.GovcloudEnabled": "GovcloudEnabled",
	"NIOS.Name":            "Name",
	"NIOS.NiosUserName":    "NiosUserName",
	"NIOS.SecretAccessKey": "SecretAccessKey",
}

// TODO: only searchable fields should be included here
// AwsuserFilterFieldMap maps infoblox filter keys to backend-specific API filter field names
var AwsuserFilterFieldMap = map[core.BackendType]map[string]string{
	core.BackendNIOS: {
		"id":                     "_ref",
		"nios.access_key_id":     "access_key_id",
		"nios.account_id":        "account_id",
		"nios.govcloud_enabled":  "govcloud_enabled",
		"nios.name":              "name",
		"nios.nios_user_name":    "nios_user_name",
		"nios.secret_access_key": "secret_access_key",
	},
}
