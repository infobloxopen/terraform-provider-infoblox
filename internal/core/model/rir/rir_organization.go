package rir

// Infoblox RirOrganization model
type RirOrganization struct {
	Id   *string
	NIOS *NIOSRirOrganizationExt
}

// NIOSRirOrganizationExt - NIOS specific fields for RirOrganization
type NIOSRirOrganizationExt struct {
	ExtAttrs    map[string]any
	Id          *string
	Maintainer  *string
	Name        *string
	Password    *string
	Rir         *string
	SenderEmail *string
}
