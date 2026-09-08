package dns

import "github.com/infobloxopen/terraform-provider-infoblox/internal/core"

// ZoneDelegatedNIOSFieldMap maps infoblox model fields to NIOS struct fields
var ZoneDelegatedNIOSFieldMap = map[string]string{
	"Id":                          "Ref",
	"NIOS.Comment":                "Comment",
	"NIOS.DelegateTo":             "DelegateTo",
	"NIOS.DelegatedTtl":           "DelegatedTtl",
	"NIOS.Disable":                "Disable",
	"NIOS.EnableRfc2317Exclusion": "EnableRfc2317Exclusion",
	"NIOS.Fqdn":                   "Fqdn",
	"NIOS.Locked":                 "Locked",
	"NIOS.MsAdIntegrated":         "MsAdIntegrated",
	"NIOS.MsDdnsMode":             "MsDdnsMode",
	"NIOS.NsGroup":                "NsGroup",
	"NIOS.Prefix":                 "Prefix",
	"NIOS.UseDelegatedTtl":        "UseDelegatedTtl",
	"NIOS.View":                   "View",
	"NIOS.ZoneFormat":             "ZoneFormat",
}

// ZoneDelegatedUDDIFieldMap maps infoblox model fields to UDDI struct fields
var ZoneDelegatedUDDIFieldMap = map[string]string{
	"UDDI.Comment":           "Comment",
	"UDDI.CompartmentId":     "CompartmentId",
	"UDDI.DelegationServers": "DelegationServers",
	"UDDI.Disabled":          "Disabled",
	"UDDI.Fqdn":              "Fqdn",
	"UDDI.Parent":            "Parent",
	"UDDI.Tags":              "Tags",
	"UDDI.View":              "View",
}

// TODO: only searchable fields should be included here
// ZoneDelegatedFilterFieldMap maps infoblox filter keys to backend-specific API filter field names
var ZoneDelegatedFilterFieldMap = map[core.BackendType]map[string]string{
	core.BackendNIOS: {
		"id":                            "_ref",
		"nios.comment":                  "comment",
		"nios.delegate_to":              "delegate_to",
		"nios.delegated_ttl":            "delegated_ttl",
		"nios.disable":                  "disable",
		"nios.enable_rfc2317_exclusion": "enable_rfc2317_exclusion",
		"nios.ext_attrs":                "extattrs",
		"nios.fqdn":                     "fqdn",
		"nios.locked":                   "locked",
		"nios.ms_ad_integrated":         "ms_ad_integrated",
		"nios.ms_ddns_mode":             "ms_ddns_mode",
		"nios.ns_group":                 "ns_group",
		"nios.prefix":                   "prefix",
		"nios.use_delegated_ttl":        "use_delegated_ttl",
		"nios.view":                     "view",
		"nios.zone_format":              "zone_format",
	},
	core.BackendUDDI: {
		"uddi.comment":            "comment",
		"uddi.compartment_id":     "compartment_id",
		"uddi.delegation_servers": "delegation_servers",
		"uddi.disabled":           "disabled",
		"uddi.fqdn":               "fqdn",
		"uddi.parent":             "parent",
		"uddi.tags":               "tags",
		"uddi.view":               "view",
	},
}
