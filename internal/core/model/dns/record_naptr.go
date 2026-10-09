package dns

import (
	niosdns "github.com/infobloxopen/infoblox-nios-go-client/dns"
	uddidnsdata "github.com/infobloxopen/universal-ddi-go-client/dnsdata"
)

// Infoblox RecordNaptr model
type RecordNaptr struct {
	Id   *string
	NIOS *NIOSRecordNaptrExt
	UDDI *UDDIRecordNaptrExt
}

// NIOSRecordNaptrExt - NIOS specific fields for RecordNaptr
type NIOSRecordNaptrExt struct {
	CloudInfo         *niosdns.RecordNaptrCloudInfo
	Comment           *string
	Creator           *string
	DdnsPrincipal     *string
	DdnsProtected     *bool
	Disable           *bool
	ExtAttrs          map[string]any
	Flags             *string
	ForbidReclamation *bool
	Name              *string
	Order             *int64
	Preference        *int64
	Regexp            *string
	Replacement       *string
	Services          *string
	Ttl               *int64
	UseTtl            *bool
	View              *string
}

// UDDIRecordNaptrExt - UDDI specific fields for RecordNaptr
type UDDIRecordNaptrExt struct {
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
