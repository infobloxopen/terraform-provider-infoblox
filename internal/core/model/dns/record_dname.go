package dns

import (
	niosdns "github.com/infobloxopen/infoblox-nios-go-client/dns"
	uddidnsdata "github.com/infobloxopen/universal-ddi-go-client/dnsdata"
)

// Infoblox RecordDname model
type RecordDname struct {
	Id   *string
	NIOS *NIOSRecordDnameExt
	UDDI *UDDIRecordDnameExt
}

// NIOSRecordDnameExt - NIOS specific fields for RecordDname
type NIOSRecordDnameExt struct {
	CloudInfo         *niosdns.RecordDnameCloudInfo
	Comment           *string
	Creator           *string
	DdnsPrincipal     *string
	DdnsProtected     *bool
	Disable           *bool
	ExtAttrs          map[string]any
	ForbidReclamation *bool
	Name              *string
	Target            *string
	Ttl               *int64
	UseTtl            *bool
	View              *string
}

// UDDIRecordDnameExt - UDDI specific fields for RecordDname
type UDDIRecordDnameExt struct {
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
