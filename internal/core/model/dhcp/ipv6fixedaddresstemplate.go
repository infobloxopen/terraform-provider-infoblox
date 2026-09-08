package dhcp

import (
	niosdhcp "github.com/infobloxopen/infoblox-nios-go-client/dhcp"
)

// Infoblox Ipv6fixedaddresstemplate model
type Ipv6fixedaddresstemplate struct {
	Id   *string
	NIOS *NIOSIpv6fixedaddresstemplateExt
}

// NIOSIpv6fixedaddresstemplateExt - NIOS specific fields for Ipv6fixedaddresstemplate
type NIOSIpv6fixedaddresstemplateExt struct {
	Comment              *string
	DomainName           *string
	DomainNameServers    []string
	ExtAttrs             map[string]any
	LogicFilterRules     []niosdhcp.Ipv6fixedaddresstemplateLogicFilterRules
	Name                 *string
	NumberOfAddresses    *int64
	Offset               *int64
	Options              []niosdhcp.Ipv6fixedaddresstemplateOptions
	PreferredLifetime    *int64
	UseDomainName        *bool
	UseDomainNameServers *bool
	UseLogicFilterRules  *bool
	UseOptions           *bool
	UsePreferredLifetime *bool
	UseValidLifetime     *bool
	ValidLifetime        *int64
}
