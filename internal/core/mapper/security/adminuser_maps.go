package security

import "github.com/infobloxopen/terraform-provider-infoblox/internal/core"

// AdminuserNIOSFieldMap maps infoblox model fields to NIOS struct fields
var AdminuserNIOSFieldMap = map[string]string{
	"Id":                                   "Ref",
	"NIOS.AdminGroups":                     "AdminGroups",
	"NIOS.AuthMethod":                      "AuthMethod",
	"NIOS.AuthType":                        "AuthType",
	"NIOS.CaCertificateIssuer":             "CaCertificateIssuer",
	"NIOS.ClientCertificateSerialNumber":   "ClientCertificateSerialNumber",
	"NIOS.Comment":                         "Comment",
	"NIOS.Disable":                         "Disable",
	"NIOS.Email":                           "Email",
	"NIOS.EnableCertificateAuthentication": "EnableCertificateAuthentication",
	"NIOS.Name":                            "Name",
	"NIOS.Password":                        "Password",
	"NIOS.SshKeys":                         "SshKeys",
	"NIOS.TimeZone":                        "TimeZone",
	"NIOS.UseSshKeys":                      "UseSshKeys",
	"NIOS.UseTimeZone":                     "UseTimeZone",
}

// TODO: only searchable fields should be included here
// AdminuserFilterFieldMap maps infoblox filter keys to backend-specific API filter field names
var AdminuserFilterFieldMap = map[core.BackendType]map[string]string{
	core.BackendNIOS: {
		"id":                                     "_ref",
		"nios.admin_groups":                      "admin_groups",
		"nios.auth_method":                       "auth_method",
		"nios.auth_type":                         "auth_type",
		"nios.ca_certificate_issuer":             "ca_certificate_issuer",
		"nios.client_certificate_serial_number":  "client_certificate_serial_number",
		"nios.comment":                           "comment",
		"nios.disable":                           "disable",
		"nios.email":                             "email",
		"nios.enable_certificate_authentication": "enable_certificate_authentication",
		"nios.ext_attrs":                         "extattrs",
		"nios.name":                              "name",
		"nios.password":                          "password",
		"nios.ssh_keys":                          "ssh_keys",
		"nios.time_zone":                         "time_zone",
		"nios.use_ssh_keys":                      "use_ssh_keys",
		"nios.use_time_zone":                     "use_time_zone",
	},
}
