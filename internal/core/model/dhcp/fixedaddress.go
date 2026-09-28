package dhcp

import (
	niosdhcp "github.com/infobloxopen/infoblox-nios-go-client/dhcp"
	uddiipam "github.com/infobloxopen/universal-ddi-go-client/ipam"
)

// Infoblox Fixedaddress model
type Fixedaddress struct {
	Id   *string
	NIOS *NIOSFixedaddressExt
	UDDI *UDDIFixedaddressExt
}

// NIOSFixedaddressExt - NIOS specific fields for Fixedaddress
type NIOSFixedaddressExt struct {
	AgentCircuitId                 *string
	AgentRemoteId                  *string
	AllowTelnet                    *bool
	AlwaysUpdateDns                *bool
	Bootfile                       *string
	Bootserver                     *string
	CliCredentials                 []niosdhcp.FixedaddressCliCredentials
	ClientIdentifierPrependZero    *bool
	CloudInfo                      *niosdhcp.FixedaddressCloudInfo
	Comment                        *string
	DdnsDomainname                 *string
	DdnsHostname                   *string
	DenyBootp                      *bool
	DeviceDescription              *string
	DeviceLocation                 *string
	DeviceType                     *string
	DeviceVendor                   *string
	DhcpClientIdentifier           *string
	Disable                        *bool
	DisableDiscovery               *bool
	EnableDdns                     *bool
	EnableImmediateDiscovery       *bool
	EnablePxeLeaseTime             *bool
	ExtAttrs                       map[string]any
	IgnoreDhcpOptionListRequest    *bool
	Ipv4addr                       *string
	LogicFilterRules               []niosdhcp.FixedaddressLogicFilterRules
	Mac                            *string
	MatchClient                    *string
	MsOptions                      []niosdhcp.FixedaddressMsOptions
	MsServer                       *niosdhcp.FixedaddressMsServer
	Name                           *string
	Network                        *string
	NetworkView                    *string
	Nextserver                     *string
	Options                        []niosdhcp.FixedaddressOptions
	PxeLeaseTime                   *int64
	ReservedInterface              *string
	RestartIfNeeded                *bool
	Snmp3Credential                *niosdhcp.FixedaddressSnmp3Credential
	SnmpCredential                 *niosdhcp.FixedaddressSnmpCredential
	Template                       *string
	UseBootfile                    *bool
	UseBootserver                  *bool
	UseCliCredentials              *bool
	UseDdnsDomainname              *bool
	UseDenyBootp                   *bool
	UseEnableDdns                  *bool
	UseIgnoreDhcpOptionListRequest *bool
	UseLogicFilterRules            *bool
	UseMsOptions                   *bool
	UseNextserver                  *bool
	UseOptions                     *bool
	UsePxeLeaseTime                *bool
	UseSnmp3Credential             *bool
	UseSnmpCredential              *bool
	FuncCall                       *niosdhcp.FuncCall
}

// UDDIFixedaddressExt - UDDI specific fields for Fixedaddress
type UDDIFixedaddressExt struct {
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
