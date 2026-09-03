package rpz

// Infoblox RecordRpzAaaaIpaddress model
type RecordRpzAaaaIpaddress struct {
	Id   *string
	NIOS *NIOSRecordRpzAaaaIpaddressExt
}

// NIOSRecordRpzAaaaIpaddressExt - NIOS specific fields for RecordRpzAaaaIpaddress
type NIOSRecordRpzAaaaIpaddressExt struct {
	Comment  *string
	Disable  *bool
	ExtAttrs map[string]any
	Ipv6addr *string
	Name     *string
	RpZone   *string
	Ttl      *int64
	UseTtl   *bool
	View     *string
}
