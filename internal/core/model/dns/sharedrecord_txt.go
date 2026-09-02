package dns

// Infoblox SharedrecordTxt model
type SharedrecordTxt struct {
	Id   *string
	NIOS *NIOSSharedrecordTxtExt
}

// NIOSSharedrecordTxtExt - NIOS specific fields for SharedrecordTxt
type NIOSSharedrecordTxtExt struct {
	Comment           *string
	Disable           *bool
	ExtAttrs          map[string]any
	Name              *string
	SharedRecordGroup *string
	Text              *string
	Ttl               *int64
	UseTtl            *bool
}
