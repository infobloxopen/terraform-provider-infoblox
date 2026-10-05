package dns

import "github.com/infobloxopen/terraform-provider-infoblox/internal/core"

// DnsHostUDDIFieldMap maps infoblox model fields to UDDI struct fields
var DnsHostUDDIFieldMap = map[string]string{
	"UDDI.AbsoluteName":       "AbsoluteName",
	"UDDI.Address":            "Address",
	"UDDI.AnycastAddresses":   "AnycastAddresses",
	"UDDI.AssociatedServer":   "AssociatedServer",
	"UDDI.DfpService":         "DfpService",
	"UDDI.InheritanceSources": "InheritanceSources",
	"UDDI.KerberosKeys":       "KerberosKeys",
	"UDDI.Name":               "Name",
	"UDDI.Ophid":              "Ophid",
	"UDDI.ProviderId":         "ProviderId",
	"UDDI.Server":             "Server",
	"UDDI.Tags":               "Tags",
	"UDDI.Type":               "Type",
}

// TODO: only searchable fields should be included here
// DnsHostFilterFieldMap maps infoblox filter keys to backend-specific API filter field names
var DnsHostFilterFieldMap = map[core.BackendType]map[string]string{
	core.BackendUDDI: {
		"uddi.absolute_name":       "absolute_name",
		"uddi.address":             "address",
		"uddi.anycast_addresses":   "anycast_addresses",
		"uddi.associated_server":   "associated_server",
		"uddi.dfp_service":         "dfp_service",
		"uddi.inheritance_sources": "inheritance_sources",
		"uddi.kerberos_keys":       "kerberos_keys",
		"uddi.name":                "name",
		"uddi.ophid":               "ophid",
		"uddi.provider_id":         "provider_id",
		"uddi.server":              "server",
		"uddi.tags":                "tags",
		"uddi.type":                "type",
	},
}
