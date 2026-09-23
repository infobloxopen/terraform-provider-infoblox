package dns

// Infoblox SharedrecordMx model
type SharedrecordMx struct {
	Id   *string
	NIOS *NIOSSharedrecordMxExt
}

// NIOSSharedrecordMxExt - NIOS specific fields for SharedrecordMx
type NIOSSharedrecordMxExt struct {
	Comment           *string
	Disable           *bool
	ExtAttrs          map[string]any
	MailExchanger     *string
	Name              *string
	Preference        *int64
	SharedRecordGroup *string
	Ttl               *int64
	UseTtl            *bool
}
