package security

import (
	niossecurity "github.com/infobloxopen/infoblox-nios-go-client/security"
)

// Infoblox Adminuser model
type Adminuser struct {
	Id   *string
	NIOS *NIOSAdminuserExt
}

// NIOSAdminuserExt - NIOS specific fields for Adminuser
type NIOSAdminuserExt struct {
	AdminGroups                     []string
	AuthMethod                      *string
	AuthType                        *string
	CaCertificateIssuer             *string
	ClientCertificateSerialNumber   *string
	Comment                         *string
	Disable                         *bool
	Email                           *string
	EnableCertificateAuthentication *bool
	ExtAttrs                        map[string]any
	Name                            *string
	Password                        *string
	SshKeys                         []niossecurity.AdminuserSshKeys
	TimeZone                        *string
	UseSshKeys                      *bool
	UseTimeZone                     *bool
}
