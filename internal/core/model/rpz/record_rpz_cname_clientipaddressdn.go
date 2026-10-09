package rpz

// Infoblox RecordRpzCnameClientipaddressdn model
type RecordRpzCnameClientipaddressdn struct {
	Id   *string
	NIOS *NIOSRecordRpzCnameClientipaddressdnExt
}

// NIOSRecordRpzCnameClientipaddressdnExt - NIOS specific fields for RecordRpzCnameClientipaddressdn
type NIOSRecordRpzCnameClientipaddressdnExt struct {
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
