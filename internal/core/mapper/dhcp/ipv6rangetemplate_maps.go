package dhcp

import "github.com/infobloxopen/terraform-provider-infoblox/internal/core"

// Ipv6rangetemplateNIOSFieldMap maps infoblox model fields to NIOS struct fields
var Ipv6rangetemplateNIOSFieldMap = map[string]string{
	"Id":                         "Ref",
	"NIOS.CloudApiCompatible":    "CloudApiCompatible",
	"NIOS.Comment":               "Comment",
	"NIOS.DelegatedMember":       "DelegatedMember",
	"NIOS.Exclude":               "Exclude",
	"NIOS.LogicFilterRules":      "LogicFilterRules",
	"NIOS.Member":                "Member",
	"NIOS.Name":                  "Name",
	"NIOS.NumberOfAddresses":     "NumberOfAddresses",
	"NIOS.Offset":                "Offset",
	"NIOS.OptionFilterRules":     "OptionFilterRules",
	"NIOS.RecycleLeases":         "RecycleLeases",
	"NIOS.ServerAssociationType": "ServerAssociationType",
	"NIOS.UseLogicFilterRules":   "UseLogicFilterRules",
	"NIOS.UseRecycleLeases":      "UseRecycleLeases",
}

// TODO: only searchable fields should be included here
// Ipv6rangetemplateFilterFieldMap maps infoblox filter keys to backend-specific API filter field names
var Ipv6rangetemplateFilterFieldMap = map[core.BackendType]map[string]string{
	core.BackendNIOS: {
		"id":                           "_ref",
		"nios.cloud_api_compatible":    "cloud_api_compatible",
		"nios.comment":                 "comment",
		"nios.delegated_member":        "delegated_member",
		"nios.exclude":                 "exclude",
		"nios.logic_filter_rules":      "logic_filter_rules",
		"nios.member":                  "member",
		"nios.name":                    "name",
		"nios.number_of_addresses":     "number_of_addresses",
		"nios.offset":                  "offset",
		"nios.option_filter_rules":     "option_filter_rules",
		"nios.recycle_leases":          "recycle_leases",
		"nios.server_association_type": "server_association_type",
		"nios.use_logic_filter_rules":  "use_logic_filter_rules",
		"nios.use_recycle_leases":      "use_recycle_leases",
	},
}
