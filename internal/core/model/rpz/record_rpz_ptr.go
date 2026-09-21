package rpz

// Infoblox RecordRpzPtr model
type RecordRpzPtr struct {
	Id   *string
	NIOS *NIOSRecordRpzPtrExt
}

// NIOSRecordRpzPtrExt - NIOS specific fields for RecordRpzPtr
type NIOSRecordRpzPtrExt struct {
	Comment  *string
	Disable  *bool
	ExtAttrs map[string]any
	Ipv4addr *string
	Ipv6addr *string
	Name     *string
	Ptrdname *string
	RpZone   *string
	Ttl      *int64
	UseTtl   *bool
	View     *string
}
