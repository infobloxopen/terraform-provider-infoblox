package dns

import (
	niosdns "github.com/infobloxopen/infoblox-nios-go-client/dns"
	uddidnsdata "github.com/infobloxopen/universal-ddi-go-client/dnsdata"
)

// Infoblox RecordCaa model
type RecordCaa struct {
	Id   *string
	NIOS *NIOSRecordCaaExt
	UDDI *UDDIRecordCaaExt
}

// NIOSRecordCaaExt - NIOS specific fields for RecordCaa
type NIOSRecordCaaExt struct {
	CaFlag            *int64
	CaTag             *string
	CaValue           *string
	CloudInfo         *niosdns.RecordCaaCloudInfo
	Comment           *string
	Creator           *string
	DdnsPrincipal     *string
	DdnsProtected     *bool
	Disable           *bool
	ExtAttrs          map[string]any
	ForbidReclamation *bool
	Name              *string
	Ttl               *int64
	UseTtl            *bool
	View              *string
}

// UDDIRecordCaaExt - UDDI specific fields for RecordCaa
type UDDIRecordCaaExt struct {
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
