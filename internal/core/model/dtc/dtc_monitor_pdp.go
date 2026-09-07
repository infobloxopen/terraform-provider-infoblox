package dtc

import (
	uddidtc "github.com/infobloxopen/universal-ddi-go-client/dtc"
)

// Infoblox DtcMonitorPdp model
type DtcMonitorPdp struct {
	Id   *string
	NIOS *NIOSDtcMonitorPdpExt
	UDDI *UDDIDtcMonitorPdpExt
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

// UDDIDtcMonitorPdpExt - UDDI specific fields for DtcMonitorPdp
type UDDIDtcMonitorPdpExt struct {
	Comment   *string
	Disabled  *bool
	Interval  *int64
	Metadata  *uddidtc.Metadata
	Name      string
	Port      *int64
	RetryDown *int64
	RetryUp   *int64
	Tags      map[string]any
	Timeout   *int64
}
