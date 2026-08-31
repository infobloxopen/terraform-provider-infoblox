package rpz

// Infoblox RecordRpzNaptr model
type RecordRpzNaptr struct {
	Id   *string
	NIOS *NIOSRecordRpzNaptrExt
}

// NIOSRecordRpzNaptrExt - NIOS specific fields for RecordRpzNaptr
type NIOSRecordRpzNaptrExt struct {
	Comment     *string
	Disable     *bool
	ExtAttrs    map[string]any
	Flags       *string
	Name        *string
	Order       *int64
	Preference  *int64
	Regexp      *string
	Replacement *string
	RpZone      *string
	Services    *string
	Ttl         *int64
	UseTtl      *bool
	View        *string
}
