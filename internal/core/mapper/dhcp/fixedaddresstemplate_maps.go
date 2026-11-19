package dhcp

import "github.com/infobloxopen/terraform-provider-infoblox/internal/core"

// FixedaddresstemplateNIOSFieldMap maps infoblox model fields to NIOS struct fields
var FixedaddresstemplateNIOSFieldMap = map[string]string{
	"Id":                                  "Ref",
	"NIOS.Bootfile":                       "Bootfile",
	"NIOS.Bootserver":                     "Bootserver",
	"NIOS.Comment":                        "Comment",
	"NIOS.DdnsDomainname":                 "DdnsDomainname",
	"NIOS.DdnsHostname":                   "DdnsHostname",
	"NIOS.DenyBootp":                      "DenyBootp",
	"NIOS.EnableDdns":                     "EnableDdns",
	"NIOS.EnablePxeLeaseTime":             "EnablePxeLeaseTime",
	"NIOS.IgnoreDhcpOptionListRequest":    "IgnoreDhcpOptionListRequest",
	"NIOS.LogicFilterRules":               "LogicFilterRules",
	"NIOS.Name":                           "Name",
	"NIOS.Nextserver":                     "Nextserver",
	"NIOS.NumberOfAddresses":              "NumberOfAddresses",
	"NIOS.Offset":                         "Offset",
	"NIOS.Options":                        "Options",
	"NIOS.PxeLeaseTime":                   "PxeLeaseTime",
	"NIOS.UseBootfile":                    "UseBootfile",
	"NIOS.UseBootserver":                  "UseBootserver",
	"NIOS.UseDdnsDomainname":              "UseDdnsDomainname",
	"NIOS.UseDenyBootp":                   "UseDenyBootp",
	"NIOS.UseEnableDdns":                  "UseEnableDdns",
	"NIOS.UseIgnoreDhcpOptionListRequest": "UseIgnoreDhcpOptionListRequest",
	"NIOS.UseLogicFilterRules":            "UseLogicFilterRules",
	"NIOS.UseNextserver":                  "UseNextserver",
	"NIOS.UseOptions":                     "UseOptions",
	"NIOS.UsePxeLeaseTime":                "UsePxeLeaseTime",
}

// TODO: only searchable fields should be included here
// FixedaddresstemplateFilterFieldMap maps infoblox filter keys to backend-specific API filter field names
var FixedaddresstemplateFilterFieldMap = map[core.BackendType]map[string]string{
	core.BackendNIOS: {
		"id":                                       "_ref",
		"nios.bootfile":                            "bootfile",
		"nios.bootserver":                          "bootserver",
		"nios.comment":                             "comment",
		"nios.ddns_domainname":                     "ddns_domainname",
		"nios.ddns_hostname":                       "ddns_hostname",
		"nios.deny_bootp":                          "deny_bootp",
		"nios.enable_ddns":                         "enable_ddns",
		"nios.enable_pxe_lease_time":               "enable_pxe_lease_time",
		"nios.ext_attrs":                           "extattrs",
		"nios.ignore_dhcp_option_list_request":     "ignore_dhcp_option_list_request",
		"nios.logic_filter_rules":                  "logic_filter_rules",
		"nios.name":                                "name",
		"nios.nextserver":                          "nextserver",
		"nios.number_of_addresses":                 "number_of_addresses",
		"nios.offset":                              "offset",
		"nios.options":                             "options",
		"nios.pxe_lease_time":                      "pxe_lease_time",
		"nios.use_bootfile":                        "use_bootfile",
		"nios.use_bootserver":                      "use_bootserver",
		"nios.use_ddns_domainname":                 "use_ddns_domainname",
		"nios.use_deny_bootp":                      "use_deny_bootp",
		"nios.use_enable_ddns":                     "use_enable_ddns",
		"nios.use_ignore_dhcp_option_list_request": "use_ignore_dhcp_option_list_request",
		"nios.use_logic_filter_rules":              "use_logic_filter_rules",
		"nios.use_nextserver":                      "use_nextserver",
		"nios.use_options":                         "use_options",
		"nios.use_pxe_lease_time":                  "use_pxe_lease_time",
	},
}
