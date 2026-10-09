package dns

import (
	niosdns "github.com/infobloxopen/infoblox-nios-go-client/dns"
	uddidnsdata "github.com/infobloxopen/universal-ddi-go-client/dnsdata"
)

// Infoblox RecordTxt model
type RecordTxt struct {
	Id   *string
	NIOS *NIOSRecordTxtExt
	UDDI *UDDIRecordTxtExt
}

// NIOSRecordTxtExt - NIOS specific fields for RecordTxt
type NIOSRecordTxtExt struct {
	CloudInfo         *niosdns.RecordTxtCloudInfo
	Comment           *string
	Creator           *string
	DdnsPrincipal     *string
	DdnsProtected     *bool
	Disable           *bool
	ExtAttrs          map[string]any
	ForbidReclamation *bool
	Name              *string
	Text              *string
	Ttl               *int64
	UseTtl            *bool
	View              *string
}

// UDDIRecordTxtExt - UDDI specific fields for RecordTxt
type UDDIRecordTxtExt struct {
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
