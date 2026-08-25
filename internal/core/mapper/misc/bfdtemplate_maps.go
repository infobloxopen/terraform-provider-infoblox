package misc

import "github.com/infobloxopen/terraform-provider-infoblox/internal/core"

// BfdtemplateNIOSFieldMap maps infoblox model fields to NIOS struct fields
var BfdtemplateNIOSFieldMap = map[string]string{
	"Id":                       "Ref",
	"NIOS.AuthenticationKey":   "AuthenticationKey",
	"NIOS.AuthenticationKeyId": "AuthenticationKeyId",
	"NIOS.AuthenticationType":  "AuthenticationType",
	"NIOS.DetectionMultiplier": "DetectionMultiplier",
	"NIOS.MinRxInterval":       "MinRxInterval",
	"NIOS.MinTxInterval":       "MinTxInterval",
	"NIOS.Name":                "Name",
}

// TODO: only searchable fields should be included here
// BfdtemplateFilterFieldMap maps infoblox filter keys to backend-specific API filter field names
var BfdtemplateFilterFieldMap = map[core.BackendType]map[string]string{
	core.BackendNIOS: {
		"id":                         "_ref",
		"nios.authentication_key":    "authentication_key",
		"nios.authentication_key_id": "authentication_key_id",
		"nios.authentication_type":   "authentication_type",
		"nios.detection_multiplier":  "detection_multiplier",
		"nios.min_rx_interval":       "min_rx_interval",
		"nios.min_tx_interval":       "min_tx_interval",
		"nios.name":                  "name",
	},
}
