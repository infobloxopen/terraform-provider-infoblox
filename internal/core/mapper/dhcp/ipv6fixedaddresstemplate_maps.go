package dhcp

import "github.com/infobloxopen/terraform-provider-infoblox/internal/core"

// Ipv6fixedaddresstemplateNIOSFieldMap maps infoblox model fields to NIOS struct fields
var Ipv6fixedaddresstemplateNIOSFieldMap = map[string]string{
	"Id":                        "Ref",
	"NIOS.Comment":              "Comment",
	"NIOS.DomainName":           "DomainName",
	"NIOS.DomainNameServers":    "DomainNameServers",
	"NIOS.LogicFilterRules":     "LogicFilterRules",
	"NIOS.Name":                 "Name",
	"NIOS.NumberOfAddresses":    "NumberOfAddresses",
	"NIOS.Offset":               "Offset",
	"NIOS.Options":              "Options",
	"NIOS.PreferredLifetime":    "PreferredLifetime",
	"NIOS.UseDomainName":        "UseDomainName",
	"NIOS.UseDomainNameServers": "UseDomainNameServers",
	"NIOS.UseLogicFilterRules":  "UseLogicFilterRules",
	"NIOS.UseOptions":           "UseOptions",
	"NIOS.UsePreferredLifetime": "UsePreferredLifetime",
	"NIOS.UseValidLifetime":     "UseValidLifetime",
	"NIOS.ValidLifetime":        "ValidLifetime",
}

// TODO: only searchable fields should be included here
// Ipv6fixedaddresstemplateFilterFieldMap maps infoblox filter keys to backend-specific API filter field names
var Ipv6fixedaddresstemplateFilterFieldMap = map[core.BackendType]map[string]string{
	core.BackendNIOS: {
		"id":                           "_ref",
		"nios.comment":                 "comment",
		"nios.domain_name":             "domain_name",
		"nios.domain_name_servers":     "domain_name_servers",
		"nios.ext_attrs":               "extattrs",
		"nios.logic_filter_rules":      "logic_filter_rules",
		"nios.name":                    "name",
		"nios.number_of_addresses":     "number_of_addresses",
		"nios.offset":                  "offset",
		"nios.options":                 "options",
		"nios.preferred_lifetime":      "preferred_lifetime",
		"nios.use_domain_name":         "use_domain_name",
		"nios.use_domain_name_servers": "use_domain_name_servers",
		"nios.use_logic_filter_rules":  "use_logic_filter_rules",
		"nios.use_options":             "use_options",
		"nios.use_preferred_lifetime":  "use_preferred_lifetime",
		"nios.use_valid_lifetime":      "use_valid_lifetime",
		"nios.valid_lifetime":          "valid_lifetime",
	},
}
