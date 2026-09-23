package rpz

// Infoblox RecordRpzCnameIpaddressdn model
type RecordRpzCnameIpaddressdn struct {
	Id   *string
	NIOS *NIOSRecordRpzCnameIpaddressdnExt
}

// NIOSRecordRpzCnameIpaddressdnExt - NIOS specific fields for RecordRpzCnameIpaddressdn
type NIOSRecordRpzCnameIpaddressdnExt struct {
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
