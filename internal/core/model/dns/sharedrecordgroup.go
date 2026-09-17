package dns

import (
	niosdns "github.com/infobloxopen/infoblox-nios-go-client/dns"
)

// Infoblox Sharedrecordgroup model
type Sharedrecordgroup struct {
	Id   *string
	NIOS *NIOSSharedrecordgroupExt
}

// NIOSSharedrecordgroupExt - NIOS specific fields for Sharedrecordgroup
type NIOSSharedrecordgroupExt struct {
	Comment             *string
	ExtAttrs            map[string]any
	Name                *string
	RecordNamePolicy    *string
	UseRecordNamePolicy *bool
	ZoneAssociations    []niosdns.SharedrecordgroupZoneAssociations
}
