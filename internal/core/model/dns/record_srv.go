package dns

import (
	niosdns "github.com/infobloxopen/infoblox-nios-go-client/dns"
	uddidnsdata "github.com/infobloxopen/universal-ddi-go-client/dnsdata"
)

// Infoblox RecordSrv model
type RecordSrv struct {
	Id   *string
	NIOS *NIOSRecordSrvExt
	UDDI *UDDIRecordSrvExt
}

// NIOSRecordSrvExt - NIOS specific fields for RecordSrv
type NIOSRecordSrvExt struct {
	CloudInfo         *niosdns.RecordSrvCloudInfo
	Comment           *string
	Creator           *string
	DdnsPrincipal     *string
	DdnsProtected     *bool
	Disable           *bool
	ExtAttrs          map[string]any
	ForbidReclamation *bool
	Name              *string
	Port              *int64
	Priority          *int64
	Target            *string
	Ttl               *int64
	UseTtl            *bool
	View              *string
	Weight            *int64
}

// UDDIRecordSrvExt - UDDI specific fields for RecordSrv
type UDDIRecordSrvExt struct {
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
