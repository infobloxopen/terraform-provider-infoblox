package dns

import (
	niosdns "github.com/infobloxopen/infoblox-nios-go-client/dns"
	uddidnsconfig "github.com/infobloxopen/universal-ddi-go-client/dnsconfig"
)

// Infoblox ZoneForward model
type ZoneForward struct {
	Id   *string
	NIOS *NIOSZoneForwardExt
	UDDI *UDDIZoneForwardExt
}

// NIOSZoneForwardExt - NIOS specific fields for ZoneForward
type NIOSZoneForwardExt struct {
	Comment             *string
	Disable             *bool
	DisableNsGeneration *bool
	ExtAttrs            map[string]any
	ExternalNsGroup     *string
	ForwardTo           []niosdns.ZoneForwardForwardTo
	ForwardersOnly      *bool
	ForwardingServers   []niosdns.ZoneForwardForwardingServers
	Fqdn                *string
	Locked              *bool
	MsAdIntegrated      *bool
	MsDdnsMode          *string
	NsGroup             *string
	Prefix              *string
	View                *string
	ZoneFormat          *string
}

// UDDIZoneForwardExt - UDDI specific fields for ZoneForward
type UDDIZoneForwardExt struct {
	Comment            *string
	CompartmentId      *string
	Disabled           *bool
	ExternalForwarders []uddidnsconfig.Forwarder
	ForwardOnly        *bool
	Fqdn               *string
	Hosts              []string
	InternalForwarders []string
	Nsgs               []string
	Parent             *string
	Tags               map[string]any
	View               *string
}
