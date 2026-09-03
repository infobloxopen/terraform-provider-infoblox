package dns

import "github.com/infobloxopen/terraform-provider-infoblox/internal/core"

// NsgroupStubmemberNIOSFieldMap maps infoblox model fields to NIOS struct fields
var NsgroupStubmemberNIOSFieldMap = map[string]string{
	"Id":               "Ref",
	"NIOS.Comment":     "Comment",
	"NIOS.Name":        "Name",
	"NIOS.StubMembers": "StubMembers",
}

// TODO: only searchable fields should be included here
// NsgroupStubmemberFilterFieldMap maps infoblox filter keys to backend-specific API filter field names
var NsgroupStubmemberFilterFieldMap = map[core.BackendType]map[string]string{
	core.BackendNIOS: {
		"id":                "_ref",
		"nios.comment":      "comment",
		"nios.ext_attrs":    "extattrs",
		"nios.name":         "name",
		"nios.stub_members": "stub_members",
	},
}
