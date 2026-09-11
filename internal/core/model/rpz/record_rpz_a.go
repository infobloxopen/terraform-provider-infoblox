package rpz

// Infoblox RecordRpzA model
type RecordRpzA struct {
	Id   *string
	NIOS *NIOSRecordRpzAExt
}

// NIOSRecordRpzAExt - NIOS specific fields for RecordRpzA
type NIOSRecordRpzAExt struct {
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
