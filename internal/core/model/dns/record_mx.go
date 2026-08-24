package dns

import (
	niosdns "github.com/infobloxopen/infoblox-nios-go-client/dns"
	uddidnsdata "github.com/infobloxopen/universal-ddi-go-client/dnsdata"
)

// Infoblox RecordMx model
type RecordMx struct {
	Id   *string
	NIOS *NIOSRecordMxExt
	UDDI *UDDIRecordMxExt
}

// NIOSRecordMxExt - NIOS specific fields for RecordMx
type NIOSRecordMxExt struct {
	CloudInfo         *niosdns.RecordMxCloudInfo
	Comment           *string
	Creator           *string
	DdnsPrincipal     *string
	DdnsProtected     *bool
	Disable           *bool
	ExtAttrs          map[string]any
	ForbidReclamation *bool
	MailExchanger     *string
	Name              *string
	Preference        *int64
	Ttl               *int64
	UseTtl            *bool
	View              *string
}

// UDDIRecordMxExt - UDDI specific fields for RecordMx
type UDDIRecordMxExt struct {
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
