package dns

import (
	niosdns "github.com/infobloxopen/infoblox-nios-go-client/dns"
	uddidnsdata "github.com/infobloxopen/universal-ddi-go-client/dnsdata"
)

// Infoblox RecordCname model
type RecordCname struct {
	Id   *string
	NIOS *NIOSRecordCnameExt
	UDDI *UDDIRecordCnameExt
}

// NIOSRecordCnameExt - NIOS specific fields for RecordCname
type NIOSRecordCnameExt struct {
	Canonical         *string
	CloudInfo         *niosdns.RecordCnameCloudInfo
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

// UDDIRecordCnameExt - UDDI specific fields for RecordCname
type UDDIRecordCnameExt struct {
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
