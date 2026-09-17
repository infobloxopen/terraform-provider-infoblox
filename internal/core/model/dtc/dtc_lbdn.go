package dtc

import (
	niosdtc "github.com/infobloxopen/infoblox-nios-go-client/dtc"
	uddidnsconfig "github.com/infobloxopen/universal-ddi-go-client/dnsconfig"
)

// Infoblox DtcLbdn model
type DtcLbdn struct {
	Id   *string
	NIOS *NIOSDtcLbdnExt
	UDDI *UDDIDtcLbdnExt
}

// NIOSDtcLbdnExt - NIOS specific fields for DtcLbdn
type NIOSDtcLbdnExt struct {
	AuthZones                []string
	AutoConsolidatedMonitors *bool
	Comment                  *string
	Disable                  *bool
	ExtAttrs                 map[string]any
	Health                   *niosdtc.DtcLbdnHealth
	LbMethod                 *string
	Name                     *string
	Patterns                 []string
	Persistence              *int64
	Pools                    []niosdtc.DtcLbdnPools
	Priority                 *int64
	Topology                 *string
	Ttl                      *int64
	Types                    []string
	UseTtl                   *bool
}

// UDDIDtcLbdnExt - UDDI specific fields for DtcLbdn
type UDDIDtcLbdnExt struct {
	Comment            *string
	Disabled           *bool
	DtcPolicy          *uddidnsconfig.DTCPolicy
	InheritanceSources *uddidnsconfig.TTLInheritance
	Name               string
	Precedence         *int64
	Tags               map[string]any
	Ttl                *int64
	View               string
}
