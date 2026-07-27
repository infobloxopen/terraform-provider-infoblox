package dns

import (
	niosdns "github.com/infobloxopen/infoblox-nios-go-client/dns"
	uddidnsdata "github.com/infobloxopen/universal-ddi-go-client/dnsdata"
)

// Infoblox RecordAaaa model
type RecordAaaa struct {
	Id   *string
	NIOS *NIOSRecordAaaaExt
	UDDI *UDDIRecordAaaaExt
}

// NIOSRecordAaaaExt - NIOS specific fields for RecordAaaa
type NIOSRecordAaaaExt struct {
	CloudInfo         *niosdns.RecordAaaaCloudInfo
	Comment           *string
	Creator           *string
	DdnsPrincipal     *string
	DdnsProtected     *bool
	Disable           *bool
	ExtAttrs          map[string]any
	ForbidReclamation *bool
	Ipv6addr          *string
	Name              *string
	Ttl               *int64
	UseTtl            *bool
	View              *string
}

// UDDIRecordAaaaExt - UDDI specific fields for RecordAaaa
type UDDIRecordAaaaExt struct {
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
