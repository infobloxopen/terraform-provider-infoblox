package rpz

// Infoblox RecordRpzTxt model
type RecordRpzTxt struct {
	Id   *string
	NIOS *NIOSRecordRpzTxtExt
}

// NIOSRecordRpzTxtExt - NIOS specific fields for RecordRpzTxt
type NIOSRecordRpzTxtExt struct {
	Comment  *string
	Disable  *bool
	ExtAttrs map[string]any
	Name     *string
	RpZone   *string
	Text     *string
	Ttl      *int64
	UseTtl   *bool
	View     *string
}
