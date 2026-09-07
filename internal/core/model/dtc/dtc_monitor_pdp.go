package dtc

// Infoblox DtcMonitorPdp model
type DtcMonitorPdp struct {
	Id   *string
	NIOS *NIOSDtcMonitorPdpExt
}

// NIOSDtcMonitorPdpExt - NIOS specific fields for DtcMonitorPdp
type NIOSDtcMonitorPdpExt struct {
	Comment   *string
	ExtAttrs  map[string]any
	Interval  *int64
	Name      *string
	Port      *int64
	RetryDown *int64
	RetryUp   *int64
	Timeout   *int64
}
