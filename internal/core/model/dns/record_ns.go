package dns

import (
	niosdns "github.com/infobloxopen/infoblox-nios-go-client/dns"
	uddidnsdata "github.com/infobloxopen/universal-ddi-go-client/dnsdata"
)

// Infoblox RecordNs model
type RecordNs struct {
	Id   *string
	NIOS *NIOSRecordNsExt
	UDDI *UDDIRecordNsExt
}

// NIOSRecordNsExt - NIOS specific fields for RecordNs
type NIOSRecordNsExt struct {
	Addresses        []niosdns.RecordNsAddresses
	CloudInfo        *niosdns.RecordNsCloudInfo
	MsDelegationName *string
	Name             *string
	Nameserver       *string
	View             *string
}

// UDDIRecordNsExt - UDDI specific fields for RecordNs
type UDDIRecordNsExt struct {
	AbsoluteNameSpec   *string
	Comment            *string
	Disabled           *bool
	InheritanceSources *uddidnsdata.RecordInheritance
	NameInZone         *string
	Options            map[string]any
	Rdata              map[string]any
	Tags               map[string]any
	Ttl                *int64
	Type               *string
	View               *string
	Zone               *string
}
