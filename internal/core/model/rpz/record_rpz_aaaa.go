package rpz

// Infoblox RecordRpzAaaa model
type RecordRpzAaaa struct {
	Id   *string
	NIOS *NIOSRecordRpzAaaaExt
}

// NIOSRecordRpzAaaaExt - NIOS specific fields for RecordRpzAaaa
type NIOSRecordRpzAaaaExt struct {
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
