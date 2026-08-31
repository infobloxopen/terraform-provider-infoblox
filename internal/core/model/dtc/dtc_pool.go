package dtc

import (
	niosdtc "github.com/infobloxopen/infoblox-nios-go-client/dtc"
	uddidtc "github.com/infobloxopen/universal-ddi-go-client/dtc"
)

// Infoblox DtcPool model
type DtcPool struct {
	Id   *string
	NIOS *NIOSDtcPoolExt
	UDDI *UDDIDtcPoolExt
}

// NIOSDtcPoolExt - NIOS specific fields for DtcPool
type NIOSDtcPoolExt struct {
	AutoConsolidatedMonitors *bool
	Availability             *string
	Comment                  *string
	ConsolidatedMonitors     []niosdtc.DtcPoolConsolidatedMonitors
	Disable                  *bool
	ExtAttrs                 map[string]any
	Health                   *niosdtc.DtcPoolHealth
	LbAlternateMethod        *string
	LbAlternateTopology      *string
	LbDynamicRatioAlternate  *niosdtc.DtcPoolLbDynamicRatioAlternate
	LbDynamicRatioPreferred  *niosdtc.DtcPoolLbDynamicRatioPreferred
	LbPreferredMethod        *string
	LbPreferredTopology      *string
	Monitors                 []string
	Name                     *string
	Quorum                   *int64
	Servers                  []niosdtc.DtcPoolServers
	Ttl                      *int64
	UseTtl                   *bool
}

// UDDIDtcPoolExt - UDDI specific fields for DtcPool
type UDDIDtcPoolExt struct {
	Comment                   *string
	ConsolidatedHealthEnabled *bool
	Disabled                  *bool
	HealthChecks              []uddidtc.PoolHealthCheck
	InheritanceSources        *uddidtc.TTLInheritance
	Metadata                  *uddidtc.Metadata
	Method                    string
	Name                      string
	PoolAvailability          *string
	PoolServersQuorum         *int64
	ServerAvailability        *string
	ServerHealthChecksQuorum  *int64
	Servers                   []uddidtc.PoolServer
	Tags                      map[string]any
	Ttl                       *int64
}
