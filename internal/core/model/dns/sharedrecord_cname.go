package dns

// Infoblox SharedrecordCname model
type SharedrecordCname struct {
	Id   *string
	NIOS *NIOSSharedrecordCnameExt
}

// NIOSSharedrecordCnameExt - NIOS specific fields for SharedrecordCname
type NIOSSharedrecordCnameExt struct {
	Canonical         *string
	Comment           *string
	Disable           *bool
	ExtAttrs          map[string]any
	Name              *string
	SharedRecordGroup *string
	Ttl               *int64
	UseTtl            *bool
}
