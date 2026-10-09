package rpz

// Infoblox RecordRpzCnameIpaddress model
type RecordRpzCnameIpaddress struct {
	Id   *string
	NIOS *NIOSRecordRpzCnameIpaddressExt
}

// NIOSRecordRpzCnameIpaddressExt - NIOS specific fields for RecordRpzCnameIpaddress
type NIOSRecordRpzCnameIpaddressExt struct {
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
