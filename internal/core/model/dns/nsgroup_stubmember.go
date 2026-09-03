package dns

import (
	niosdns "github.com/infobloxopen/infoblox-nios-go-client/dns"
)

// Infoblox NsgroupStubmember model
type NsgroupStubmember struct {
	Id   *string
	NIOS *NIOSNsgroupStubmemberExt
}

// NIOSNsgroupStubmemberExt - NIOS specific fields for NsgroupStubmember
type NIOSNsgroupStubmemberExt struct {
	Comment     *string
	ExtAttrs    map[string]any
	Name        *string
	StubMembers []niosdns.NsgroupStubmemberStubMembers
}
