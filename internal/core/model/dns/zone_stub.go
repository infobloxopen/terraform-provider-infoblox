package dns

import (
	niosdns "github.com/infobloxopen/infoblox-nios-go-client/dns"
)

// Infoblox ZoneStub model
type ZoneStub struct {
	Id   *string
	NIOS *NIOSZoneStubExt
}

// NIOSZoneStubExt - NIOS specific fields for ZoneStub
type NIOSZoneStubExt struct {
	Comment           *string
	Disable           *bool
	DisableForwarding *bool
	ExtAttrs          map[string]any
	ExternalNsGroup   *string
	Fqdn              *string
	Locked            *bool
	MsAdIntegrated    *bool
	MsDdnsMode        *string
	NsGroup           *string
	Prefix            *string
	StubFrom          []niosdns.ZoneStubStubFrom
	StubMembers       []niosdns.ZoneStubStubMembers
	StubMsservers     []niosdns.ZoneStubStubMsservers
	View              *string
	ZoneFormat        *string
}
