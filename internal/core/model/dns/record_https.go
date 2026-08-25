package dns

import (
	uddidnsdata "github.com/infobloxopen/universal-ddi-go-client/dnsdata"
)

// Infoblox RecordHttps model
type RecordHttps struct {
	Id   *string
	UDDI *UDDIRecordHttpsExt
}

// UDDIRecordHttpsExt - UDDI specific fields for RecordHttps
type UDDIRecordHttpsExt struct {
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
