package dns

import (
	niosdns "github.com/infobloxopen/infoblox-nios-go-client/dns"
	uddidnsconfig "github.com/infobloxopen/universal-ddi-go-client/dnsconfig"
)

// Infoblox ZoneDelegated model
type ZoneDelegated struct {
	Id   *string
	NIOS *NIOSZoneDelegatedExt
	UDDI *UDDIZoneDelegatedExt
}

// NIOSZoneDelegatedExt - NIOS specific fields for ZoneDelegated
type NIOSZoneDelegatedExt struct {
	Comment                *string
	DelegateTo             []niosdns.ZoneDelegatedDelegateTo
	DelegatedTtl           *int64
	Disable                *bool
	EnableRfc2317Exclusion *bool
	ExtAttrs               map[string]any
	Fqdn                   *string
	Locked                 *bool
	MsAdIntegrated         *bool
	MsDdnsMode             *string
	NsGroup                *string
	Prefix                 *string
	UseDelegatedTtl        *bool
	View                   *string
	ZoneFormat             *string
}

// UDDIZoneDelegatedExt - UDDI specific fields for ZoneDelegated
type UDDIZoneDelegatedExt struct {
	Comment           *string
	CompartmentId     *string
	DelegationServers []uddidnsconfig.DelegationServer
	Disabled          *bool
	Fqdn              *string
	Parent            *string
	Tags              map[string]any
	View              *string
}
