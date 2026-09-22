package dtc

import (
	uddidtc "github.com/infobloxopen/universal-ddi-go-client/dtc"
)

// Infoblox DtcMonitorIcmp model
type DtcMonitorIcmp struct {
	Id   *string
	NIOS *NIOSDtcMonitorIcmpExt
	UDDI *UDDIDtcMonitorIcmpExt
}

// NIOSDtcMonitorIcmpExt - NIOS specific fields for DtcMonitorIcmp
type NIOSDtcMonitorIcmpExt struct {
	Comment   *string
	ExtAttrs  map[string]any
	Interval  *int64
	Name      *string
	RetryDown *int64
	RetryUp   *int64
	Timeout   *int64
}

// UDDIDtcMonitorIcmpExt - UDDI specific fields for DtcMonitorIcmp
type UDDIDtcMonitorIcmpExt struct {
	Comment   *string
	Disabled  *bool
	Interval  *int64
	Metadata  *uddidtc.Metadata
	Name      string
	RetryDown *int64
	RetryUp   *int64
	Tags      map[string]any
	Timeout   *int64
}
