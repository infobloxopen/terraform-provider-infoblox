package dhcp

import (
	niosdhcp "github.com/infobloxopen/infoblox-nios-go-client/dhcp"
)

// Infoblox Ipv6sharednetwork model
type Ipv6sharednetwork struct {
	Id   *string
	NIOS *NIOSIpv6sharednetworkExt
}

// NIOSIpv6sharednetworkExt - NIOS specific fields for Ipv6sharednetwork
type NIOSIpv6sharednetworkExt struct {
	Comment                    *string
	DdnsDomainname             *string
	DdnsGenerateHostname       *bool
	DdnsServerAlwaysUpdates    *bool
	DdnsTtl                    *int64
	DdnsUseOption81            *bool
	Disable                    *bool
	DomainName                 *string
	DomainNameServers          []string
	EnableDdns                 *bool
	ExtAttrs                   map[string]any
	LogicFilterRules           []niosdhcp.Ipv6sharednetworkLogicFilterRules
	Name                       *string
	NetworkView                *string
	Networks                   []niosdhcp.Ipv6sharednetworkNetworks
	Options                    []niosdhcp.Ipv6sharednetworkOptions
	PreferredLifetime          *int64
	UpdateDnsOnLeaseRenewal    *bool
	UseDdnsDomainname          *bool
	UseDdnsGenerateHostname    *bool
	UseDdnsTtl                 *bool
	UseDdnsUseOption81         *bool
	UseDomainName              *bool
	UseDomainNameServers       *bool
	UseEnableDdns              *bool
	UseLogicFilterRules        *bool
	UseOptions                 *bool
	UsePreferredLifetime       *bool
	UseUpdateDnsOnLeaseRenewal *bool
	UseValidLifetime           *bool
	ValidLifetime              *int64
}
