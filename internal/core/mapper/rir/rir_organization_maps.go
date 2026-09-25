package rir

import "github.com/infobloxopen/terraform-provider-infoblox/internal/core"

// RirOrganizationNIOSFieldMap maps infoblox model fields to NIOS struct fields
var RirOrganizationNIOSFieldMap = map[string]string{
	"Id":               "Ref",
	"NIOS.Id":          "Id",
	"NIOS.Maintainer":  "Maintainer",
	"NIOS.Name":        "Name",
	"NIOS.Password":    "Password",
	"NIOS.Rir":         "Rir",
	"NIOS.SenderEmail": "SenderEmail",
}

// TODO: only searchable fields should be included here
// RirOrganizationFilterFieldMap maps infoblox filter keys to backend-specific API filter field names
var RirOrganizationFilterFieldMap = map[core.BackendType]map[string]string{
	core.BackendNIOS: {
		"id":                "_ref",
		"nios.ext_attrs":    "extattrs",
		"nios.id":           "id",
		"nios.maintainer":   "maintainer",
		"nios.name":         "name",
		"nios.password":     "password",
		"nios.rir":          "rir",
		"nios.sender_email": "sender_email",
	},
}
