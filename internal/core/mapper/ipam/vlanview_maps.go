package ipam

import "github.com/infobloxopen/terraform-provider-infoblox/internal/core"

// VlanviewNIOSFieldMap maps infoblox model fields to NIOS struct fields
var VlanviewNIOSFieldMap = map[string]string{
	"Id":                         "Ref",
	"NIOS.AllowRangeOverlapping": "AllowRangeOverlapping",
	"NIOS.Comment":               "Comment",
	"NIOS.EndVlanId":             "EndVlanId",
	"NIOS.Name":                  "Name",
	"NIOS.PreCreateVlan":         "PreCreateVlan",
	"NIOS.StartVlanId":           "StartVlanId",
	"NIOS.VlanNamePrefix":        "VlanNamePrefix",
}

// TODO: only searchable fields should be included here
// VlanviewFilterFieldMap maps infoblox filter keys to backend-specific API filter field names
var VlanviewFilterFieldMap = map[core.BackendType]map[string]string{
	core.BackendNIOS: {
		"id":                           "_ref",
		"nios.allow_range_overlapping": "allow_range_overlapping",
		"nios.comment":                 "comment",
		"nios.end_vlan_id":             "end_vlan_id",
		"nios.ext_attrs":               "extattrs",
		"nios.name":                    "name",
		"nios.pre_create_vlan":         "pre_create_vlan",
		"nios.start_vlan_id":           "start_vlan_id",
		"nios.vlan_name_prefix":        "vlan_name_prefix",
	},
}
