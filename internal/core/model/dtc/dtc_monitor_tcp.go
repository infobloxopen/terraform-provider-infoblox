package dtc

import (
	uddidtc "github.com/infobloxopen/universal-ddi-go-client/dtc"
)

// Infoblox DtcMonitorTcp model
type DtcMonitorTcp struct {
	Id   *string
	NIOS *NIOSDtcMonitorTcpExt
	UDDI *UDDIDtcMonitorTcpExt
}

// NIOSDtcMonitorTcpExt - NIOS specific fields for DtcMonitorTcp
type NIOSDtcMonitorTcpExt struct {
	Comment   *string
	ExtAttrs  map[string]any
	Interval  *int64
	Name      *string
	Port      *int64
	RetryDown *int64
	RetryUp   *int64
	Timeout   *int64
}

// UDDIDtcMonitorTcpExt - UDDI specific fields for DtcMonitorTcp
type UDDIDtcMonitorTcpExt struct {
	Comment   *string
	Disabled  *bool
	Interval  *int64
	Metadata  *uddidtc.Metadata
	Name      string
	Port      int64
	RetryDown *int64
	RetryUp   *int64
	Tags      map[string]any
	Timeout   *int64
}
