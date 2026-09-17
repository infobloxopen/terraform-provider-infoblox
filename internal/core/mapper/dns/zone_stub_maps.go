package dns

import "github.com/infobloxopen/terraform-provider-infoblox/internal/core"

// ZoneStubNIOSFieldMap maps infoblox model fields to NIOS struct fields
var ZoneStubNIOSFieldMap = map[string]string{
	"Id":                     "Ref",
	"NIOS.Comment":           "Comment",
	"NIOS.Disable":           "Disable",
	"NIOS.DisableForwarding": "DisableForwarding",
	"NIOS.ExternalNsGroup":   "ExternalNsGroup",
	"NIOS.Fqdn":              "Fqdn",
	"NIOS.Locked":            "Locked",
	"NIOS.MsAdIntegrated":    "MsAdIntegrated",
	"NIOS.MsDdnsMode":        "MsDdnsMode",
	"NIOS.NsGroup":           "NsGroup",
	"NIOS.Prefix":            "Prefix",
	"NIOS.StubFrom":          "StubFrom",
	"NIOS.StubMembers":       "StubMembers",
	"NIOS.StubMsservers":     "StubMsservers",
	"NIOS.View":              "View",
	"NIOS.ZoneFormat":        "ZoneFormat",
}

// TODO: only searchable fields should be included here
// ZoneStubFilterFieldMap maps infoblox filter keys to backend-specific API filter field names
var ZoneStubFilterFieldMap = map[core.BackendType]map[string]string{
	core.BackendNIOS: {
		"id":                      "_ref",
		"nios.comment":            "comment",
		"nios.disable":            "disable",
		"nios.disable_forwarding": "disable_forwarding",
		"nios.ext_attrs":          "extattrs",
		"nios.external_ns_group":  "external_ns_group",
		"nios.fqdn":               "fqdn",
		"nios.locked":             "locked",
		"nios.ms_ad_integrated":   "ms_ad_integrated",
		"nios.ms_ddns_mode":       "ms_ddns_mode",
		"nios.ns_group":           "ns_group",
		"nios.prefix":             "prefix",
		"nios.stub_from":          "stub_from",
		"nios.stub_members":       "stub_members",
		"nios.stub_msservers":     "stub_msservers",
		"nios.view":               "view",
		"nios.zone_format":        "zone_format",
	},
}
