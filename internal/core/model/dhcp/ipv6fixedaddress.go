package dhcp

import (
	niosdhcp "github.com/infobloxopen/infoblox-nios-go-client/dhcp"
	uddiipam "github.com/infobloxopen/universal-ddi-go-client/ipam"
)

// Infoblox Ipv6fixedaddress model
type Ipv6fixedaddress struct {
	Id   *string
	NIOS *NIOSIpv6fixedaddressExt
	UDDI *UDDIIpv6fixedaddressExt
}

// NIOSIpv6fixedaddressExt - NIOS specific fields for Ipv6fixedaddress
type NIOSIpv6fixedaddressExt struct {
	AddressType              *string
	AllowTelnet              *bool
	CliCredentials           []niosdhcp.Ipv6fixedaddressCliCredentials
	CloudInfo                *niosdhcp.Ipv6fixedaddressCloudInfo
	Comment                  *string
	DeviceDescription        *string
	DeviceLocation           *string
	DeviceType               *string
	DeviceVendor             *string
	Disable                  *bool
	DisableDiscovery         *bool
	DomainName               *string
	DomainNameServers        []string
	Duid                     *string
	EnableImmediateDiscovery *bool
	ExtAttrs                 map[string]any
	Ipv6addr                 *string
	Ipv6prefix               *string
	Ipv6prefixBits           *int64
	LogicFilterRules         []niosdhcp.Ipv6fixedaddressLogicFilterRules
	MacAddress               *string
	MatchClient              *string
	Name                     *string
	Network                  *string
	NetworkView              *string
	Options                  []niosdhcp.Ipv6fixedaddressOptions
	PreferredLifetime        *int64
	ReservedInterface        *string
	RestartIfNeeded          *bool
	Snmp3Credential          *niosdhcp.Ipv6fixedaddressSnmp3Credential
	SnmpCredential           *niosdhcp.Ipv6fixedaddressSnmpCredential
	Template                 *string
	UseCliCredentials        *bool
	UseDomainName            *bool
	UseDomainNameServers     *bool
	UseLogicFilterRules      *bool
	UseOptions               *bool
	UsePreferredLifetime     *bool
	UseSnmp3Credential       *bool
	UseSnmpCredential        *bool
	UseValidLifetime         *bool
	ValidLifetime            *int64
	FuncCall                 *niosdhcp.FuncCall
}

// UDDIIpv6fixedaddressExt - UDDI specific fields for Ipv6fixedaddress
type UDDIIpv6fixedaddressExt struct {
	Address                   string
	Comment                   *string
	DhcpOptions               []uddiipam.OptionItem
	DisableDhcp               *bool
	HeaderOptionFilename      *string
	HeaderOptionServerAddress *string
	HeaderOptionServerName    *string
	Hostname                  *string
	InheritanceParent         *string
	InheritanceSources        *uddiipam.FixedAddressInheritance
	IpSpace                   *string
	MatchType                 string
	MatchValue                string
	Name                      *string
	Parent                    *string
	Tags                      map[string]any
}
