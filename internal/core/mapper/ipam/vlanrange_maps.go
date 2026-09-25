package ipam

import "github.com/infobloxopen/terraform-provider-infoblox/internal/core"

// VlanrangeNIOSFieldMap maps infoblox model fields to NIOS struct fields
var VlanrangeNIOSFieldMap = map[string]string{
	"Id":                  "Ref",
	"NIOS.Comment":        "Comment",
	"NIOS.DeleteVlans":    "DeleteVlans",
	"NIOS.EndVlanId":      "EndVlanId",
	"NIOS.Name":           "Name",
	"NIOS.PreCreateVlan":  "PreCreateVlan",
	"NIOS.StartVlanId":    "StartVlanId",
	"NIOS.VlanNamePrefix": "VlanNamePrefix",
	"NIOS.VlanView":       "VlanView",
}

// TODO: only searchable fields should be included here
// VlanrangeFilterFieldMap maps infoblox filter keys to backend-specific API filter field names
var VlanrangeFilterFieldMap = map[core.BackendType]map[string]string{
	core.BackendNIOS: {
		"id":                    "_ref",
		"nios.comment":          "comment",
		"nios.delete_vlans":     "delete_vlans",
		"nios.end_vlan_id":      "end_vlan_id",
		"nios.ext_attrs":        "extattrs",
		"nios.name":             "name",
		"nios.pre_create_vlan":  "pre_create_vlan",
		"nios.start_vlan_id":    "start_vlan_id",
		"nios.vlan_name_prefix": "vlan_name_prefix",
		"nios.vlan_view":        "vlan_view",
	},
}
