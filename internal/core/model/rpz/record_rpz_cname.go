package rpz

// Infoblox RecordRpzCname model
type RecordRpzCname struct {
	Id   *string
	NIOS *NIOSRecordRpzCnameExt
}

// NIOSRecordRpzCnameExt - NIOS specific fields for RecordRpzCname
type NIOSRecordRpzCnameExt struct {
	Canonical *string
	Comment   *string
	Disable   *bool
	ExtAttrs  map[string]any
	Name      *string
	RpZone    *string
	Ttl       *int64
	UseTtl    *bool
	View      *string
}
