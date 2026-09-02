package dns

// Infoblox SharedrecordAaaa model
type SharedrecordAaaa struct {
	Id   *string
	NIOS *NIOSSharedrecordAaaaExt
}

// NIOSSharedrecordAaaaExt - NIOS specific fields for SharedrecordAaaa
type NIOSSharedrecordAaaaExt struct {
	Comment           *string
	Disable           *bool
	ExtAttrs          map[string]any
	Ipv6addr          *string
	Name              *string
	SharedRecordGroup *string
	Ttl               *int64
	UseTtl            *bool
}
