package dtc

import (
	niosdtc "github.com/infobloxopen/infoblox-nios-go-client/dtc"
	uddidtc "github.com/infobloxopen/universal-ddi-go-client/dtc"
)

// Infoblox DtcMonitorSnmp model
type DtcMonitorSnmp struct {
	Id   *string
	NIOS *NIOSDtcMonitorSnmpExt
	UDDI *UDDIDtcMonitorSnmpExt
}

// NIOSDtcMonitorSnmpExt - NIOS specific fields for DtcMonitorSnmp
type NIOSDtcMonitorSnmpExt struct {
	Comment   *string
	Community *string
	Context   *string
	EngineId  *string
	ExtAttrs  map[string]any
	Interval  *int64
	Name      *string
	Oids      []niosdtc.DtcMonitorSnmpOids
	Port      *int64
	RetryDown *int64
	RetryUp   *int64
	Timeout   *int64
	User      *string
	Version   *string
}

// UDDIDtcMonitorSnmpExt - UDDI specific fields for DtcMonitorSnmp
type UDDIDtcMonitorSnmpExt struct {
	CheckList         []uddidtc.SNMPHealthCheckEntryCheck
	Comment           *string
	Community         *string
	ContextEngineId   *string
	ContextName       *string
	Disabled          *bool
	Interval          *int64
	Metadata          *uddidtc.Metadata
	Name              string
	Port              *int64
	RetryDown         *int64
	RetryUp           *int64
	Tags              map[string]any
	Timeout           *int64
	UserSecurityModel *string
	Version           string
}
