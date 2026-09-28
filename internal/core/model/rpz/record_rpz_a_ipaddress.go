package rpz

// Infoblox RecordRpzAIpaddress model
type RecordRpzAIpaddress struct {
	Id   *string
	NIOS *NIOSRecordRpzAIpaddressExt
}

// NIOSRecordRpzAIpaddressExt - NIOS specific fields for RecordRpzAIpaddress
type NIOSRecordRpzAIpaddressExt struct {
	Comment  *string
	Disable  *bool
	ExtAttrs map[string]any
	Ipv4addr *string
	Name     *string
	RpZone   *string
	Ttl      *int64
	UseTtl   *bool
	View     *string
}
