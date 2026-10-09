package dns

// Infoblox SharedrecordA model
type SharedrecordA struct {
	Id   *string
	NIOS *NIOSSharedrecordAExt
}

// NIOSSharedrecordAExt - NIOS specific fields for SharedrecordA
type NIOSSharedrecordAExt struct {
	Comment           *string
	Disable           *bool
	ExtAttrs          map[string]any
	Ipv4addr          *string
	Name              *string
	SharedRecordGroup *string
	Ttl               *int64
	UseTtl            *bool
}
