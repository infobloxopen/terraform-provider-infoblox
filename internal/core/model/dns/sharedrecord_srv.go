package dns

// Infoblox SharedrecordSrv model
type SharedrecordSrv struct {
	Id   *string
	NIOS *NIOSSharedrecordSrvExt
}

// NIOSSharedrecordSrvExt - NIOS specific fields for SharedrecordSrv
type NIOSSharedrecordSrvExt struct {
	Comment           *string
	Disable           *bool
	ExtAttrs          map[string]any
	Name              *string
	Port              *int64
	Priority          *int64
	SharedRecordGroup *string
	Target            *string
	Ttl               *int64
	UseTtl            *bool
	Weight            *int64
}
