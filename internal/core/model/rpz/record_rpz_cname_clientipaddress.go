package rpz

// Infoblox RecordRpzCnameClientipaddress model
type RecordRpzCnameClientipaddress struct {
	Id   *string
	NIOS *NIOSRecordRpzCnameClientipaddressExt
}

// NIOSRecordRpzCnameClientipaddressExt - NIOS specific fields for RecordRpzCnameClientipaddress
type NIOSRecordRpzCnameClientipaddressExt struct {
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
