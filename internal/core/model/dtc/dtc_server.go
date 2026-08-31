package dtc

import (
	niosdtc "github.com/infobloxopen/infoblox-nios-go-client/dtc"
	uddidtc "github.com/infobloxopen/universal-ddi-go-client/dtc"
)

// Infoblox DtcServer model
type DtcServer struct {
	Id   *string
	NIOS *NIOSDtcServerExt
	UDDI *UDDIDtcServerExt
}

// NIOSDtcServerExt - NIOS specific fields for DtcServer
type NIOSDtcServerExt struct {
	AutoCreateHostRecord *bool
	Comment              *string
	Disable              *bool
	ExtAttrs             map[string]any
	Health               *niosdtc.DtcServerHealth
	Host                 *string
	Monitors             []niosdtc.DtcServerMonitors
	Name                 *string
	SniHostname          *string
	UseSniHostname       *bool
}

// UDDIDtcServerExt - UDDI specific fields for DtcServer
type UDDIDtcServerExt struct {
	Address                   *string
	AutoCreateResponseRecords *bool
	Comment                   *string
	Disabled                  *bool
	EndpointType              *string
	Fqdn                      *string
	Metadata                  *uddidtc.Metadata
	Name                      string
	Records                   []uddidtc.Record
	Tags                      map[string]any
}
