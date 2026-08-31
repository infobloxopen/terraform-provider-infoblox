package dns

import "github.com/infobloxopen/terraform-provider-infoblox/internal/core"

// ZoneForwardNIOSFieldMap maps infoblox model fields to NIOS struct fields
var ZoneForwardNIOSFieldMap = map[string]string{
	"Id":                       "Ref",
	"NIOS.Comment":             "Comment",
	"NIOS.Disable":             "Disable",
	"NIOS.DisableNsGeneration": "DisableNsGeneration",
	"NIOS.ExternalNsGroup":     "ExternalNsGroup",
	"NIOS.ForwardTo":           "ForwardTo",
	"NIOS.ForwardersOnly":      "ForwardersOnly",
	"NIOS.ForwardingServers":   "ForwardingServers",
	"NIOS.Fqdn":                "Fqdn",
	"NIOS.Locked":              "Locked",
	"NIOS.MsAdIntegrated":      "MsAdIntegrated",
	"NIOS.MsDdnsMode":          "MsDdnsMode",
	"NIOS.NsGroup":             "NsGroup",
	"NIOS.Prefix":              "Prefix",
	"NIOS.View":                "View",
	"NIOS.ZoneFormat":          "ZoneFormat",
}

// ZoneForwardUDDIFieldMap maps infoblox model fields to UDDI struct fields
var ZoneForwardUDDIFieldMap = map[string]string{
	"UDDI.Comment":            "Comment",
	"UDDI.CompartmentId":      "CompartmentId",
	"UDDI.Disabled":           "Disabled",
	"UDDI.ExternalForwarders": "ExternalForwarders",
	"UDDI.ForwardOnly":        "ForwardOnly",
	"UDDI.Fqdn":               "Fqdn",
	"UDDI.Hosts":              "Hosts",
	"UDDI.InternalForwarders": "InternalForwarders",
	"UDDI.Nsgs":               "Nsgs",
	"UDDI.Parent":             "Parent",
	"UDDI.Tags":               "Tags",
	"UDDI.View":               "View",
}

// TODO: only searchable fields should be included here
// ZoneForwardFilterFieldMap maps infoblox filter keys to backend-specific API filter field names
var ZoneForwardFilterFieldMap = map[core.BackendType]map[string]string{
	core.BackendNIOS: {
		"id":                         "_ref",
		"nios.comment":               "comment",
		"nios.disable":               "disable",
		"nios.disable_ns_generation": "disable_ns_generation",
		"nios.ext_attrs":             "extattrs",
		"nios.external_ns_group":     "external_ns_group",
		"nios.forward_to":            "forward_to",
		"nios.forwarders_only":       "forwarders_only",
		"nios.forwarding_servers":    "forwarding_servers",
		"nios.fqdn":                  "fqdn",
		"nios.locked":                "locked",
		"nios.ms_ad_integrated":      "ms_ad_integrated",
		"nios.ms_ddns_mode":          "ms_ddns_mode",
		"nios.ns_group":              "ns_group",
		"nios.prefix":                "prefix",
		"nios.view":                  "view",
		"nios.zone_format":           "zone_format",
	},
	core.BackendUDDI: {
		"uddi.comment":             "comment",
		"uddi.compartment_id":      "compartment_id",
		"uddi.disabled":            "disabled",
		"uddi.external_forwarders": "external_forwarders",
		"uddi.forward_only":        "forward_only",
		"uddi.fqdn":                "fqdn",
		"uddi.hosts":               "hosts",
		"uddi.internal_forwarders": "internal_forwarders",
		"uddi.nsgs":                "nsgs",
		"uddi.parent":              "parent",
		"uddi.tags":                "tags",
		"uddi.view":                "view",
	},
}
