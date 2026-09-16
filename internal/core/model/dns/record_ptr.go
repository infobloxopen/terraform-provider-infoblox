package dns

import (
	niosdns "github.com/infobloxopen/infoblox-nios-go-client/dns"
	uddidnsdata "github.com/infobloxopen/universal-ddi-go-client/dnsdata"
)

// Infoblox RecordPtr model
type RecordPtr struct {
	Id   *string
	NIOS *NIOSRecordPtrExt
	UDDI *UDDIRecordPtrExt
}

// NIOSRecordPtrExt - NIOS specific fields for RecordPtr
type NIOSRecordPtrExt struct {
	CloudInfo         *niosdns.RecordPtrCloudInfo
	Comment           *string
	Creator           *string
	DdnsPrincipal     *string
	DdnsProtected     *bool
	Disable           *bool
	ExtAttrs          map[string]any
	ForbidReclamation *bool
	Ipv4addr          *string
	Ipv6addr          *string
	Name              *string
	Ptrdname          *string
	Ttl               *int64
	UseTtl            *bool
	View              *string
	FuncCall          *niosdns.FuncCall
}

// UDDIRecordPtrExt - UDDI specific fields for RecordPtr
type UDDIRecordPtrExt struct {
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
