package clouddiscovery

import "github.com/infobloxopen/terraform-provider-infoblox/internal/core"

// CloudDiscoveryProviderUDDIFieldMap maps infoblox model fields to UDDI struct fields
var CloudDiscoveryProviderUDDIFieldMap = map[string]string{
	"UDDI.AccountPreference":       "AccountPreference",
	"UDDI.AdditionalConfig":        "AdditionalConfig",
	"UDDI.CredentialPreference":    "CredentialPreference",
	"UDDI.Description":             "Description",
	"UDDI.DesiredState":            "DesiredState",
	"UDDI.DestinationTypesEnabled": "DestinationTypesEnabled",
	"UDDI.Destinations":            "Destinations",
	"UDDI.IsDisabled":              "IsDisabled",
	"UDDI.Name":                    "Name",
	"UDDI.ProviderType":            "ProviderType",
	"UDDI.SourceConfigs":           "SourceConfigs",
	"UDDI.SyncInterval":            "SyncInterval",
	"UDDI.Tags":                    "Tags",
}

// TODO: only searchable fields should be included here
// CloudDiscoveryProviderFilterFieldMap maps infoblox filter keys to backend-specific API filter field names
var CloudDiscoveryProviderFilterFieldMap = map[core.BackendType]map[string]string{
	core.BackendUDDI: {
		"uddi.account_preference":        "account_preference",
		"uddi.additional_config":         "additional_config",
		"uddi.credential_preference":     "credential_preference",
		"uddi.description":               "description",
		"uddi.desired_state":             "desired_state",
		"uddi.destination_types_enabled": "destination_types_enabled",
		"uddi.destinations":              "destinations",
		"uddi.is_disabled":               "is_disabled",
		"uddi.name":                      "name",
		"uddi.provider_type":             "provider_type",
		"uddi.source_configs":            "source_configs",
		"uddi.sync_interval":             "sync_interval",
		"uddi.tags":                      "tags",
	},
}
