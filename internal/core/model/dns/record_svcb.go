package dns

import (
	uddidnsdata "github.com/infobloxopen/universal-ddi-go-client/dnsdata"
)

// Infoblox RecordSvcb model
type RecordSvcb struct {
	Id   *string
	UDDI *UDDIRecordSvcbExt
}

// UDDIRecordSvcbExt - UDDI specific fields for RecordSvcb
type UDDIRecordSvcbExt struct {
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
