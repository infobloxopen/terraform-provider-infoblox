package dtc

import (
	niosdtc "github.com/infobloxopen/infoblox-nios-go-client/dtc"
)

// Infoblox DtcTopology model
type DtcTopology struct {
	Id   *string
	NIOS *NIOSDtcTopologyExt
}

// NIOSDtcTopologyExt - NIOS specific fields for DtcTopology
type NIOSDtcTopologyExt struct {
	Comment  *string
	ExtAttrs map[string]any
	Name     *string
	Rules    []niosdtc.DtcTopologyRulesInner
}
