package dns

import (
	niosdns "github.com/infobloxopen/infoblox-nios-go-client/dns"
	uddidnsdata "github.com/infobloxopen/universal-ddi-go-client/dnsdata"
)

// Infoblox RecordA model
type RecordA struct {
	Id   *string
	NIOS *NIOSRecordAExt
	UDDI *UDDIRecordAExt
}

// NIOSRecordAExt - NIOS specific fields for RecordA
type NIOSRecordAExt struct {
	CloudInfo         *niosdns.RecordACloudInfo
	Comment           *string
	Creator           *string
	DdnsPrincipal     *string
	DdnsProtected     *bool
	Disable           *bool
	ExtAttrs          map[string]any
	ForbidReclamation *bool
	Ipv4addr          *string
	Name              *string
	Ttl               *int64
	UseTtl            *bool
	View              *string
	FuncCall          *niosdns.FuncCall
}

// UDDIRecordAExt - UDDI specific fields for RecordA
type UDDIRecordAExt struct {
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
