package dtc

import (
	niosdtc "github.com/infobloxopen/infoblox-nios-go-client/dtc"
	uddidtc "github.com/infobloxopen/universal-ddi-go-client/dtc"
)

// Infoblox DtcTopology model
type DtcTopology struct {
	Id   *string
	NIOS *NIOSDtcTopologyExt
	UDDI *UDDIDtcTopologyExt
}

// NIOSDtcTopologyExt - NIOS specific fields for DtcTopology
type NIOSDtcTopologyExt struct {
	Comment  *string
	ExtAttrs map[string]any
	Name     *string
	Rules    []niosdtc.DtcTopologyRulesInner
}

// UDDIDtcTopologyExt - UDDI specific fields for DtcTopology
type UDDIDtcTopologyExt struct {
	Comment  *string
	Disabled *bool
	Metadata *uddidtc.Metadata
	Name     string
	Rules    []uddidtc.TopologyRulePreset
	Tags     map[string]any
}
