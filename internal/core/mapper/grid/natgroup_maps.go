package grid

import "github.com/infobloxopen/terraform-provider-infoblox/internal/core"

// NatgroupNIOSFieldMap maps infoblox model fields to NIOS struct fields
var NatgroupNIOSFieldMap = map[string]string{
	"Id":           "Ref",
	"NIOS.Comment": "Comment",
	"NIOS.Name":    "Name",
}

// TODO: only searchable fields should be included here
// NatgroupFilterFieldMap maps infoblox filter keys to backend-specific API filter field names
var NatgroupFilterFieldMap = map[core.BackendType]map[string]string{
	core.BackendNIOS: {
		"id":           "_ref",
		"nios.comment": "comment",
		"nios.name":    "name",
	},
}
