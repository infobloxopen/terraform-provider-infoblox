package dns

import (
	niosdns "github.com/infobloxopen/infoblox-nios-go-client/dns"
)

// Infoblox RecordHost model
type RecordHost struct {
	Id   *string
	NIOS *NIOSRecordHostExt
}

// NIOSRecordHostExt - NIOS specific fields for RecordHost
type NIOSRecordHostExt struct {
	Aliases                  []string
	AllowTelnet              *bool
	CliCredentials           []niosdns.RecordHostCliCredentials
	CloudInfo                *niosdns.RecordHostCloudInfo
	Comment                  *string
	ConfigureForDns          *bool
	DdnsProtected            *bool
	DeviceDescription        *string
	DeviceLocation           *string
	DeviceType               *string
	DeviceVendor             *string
	Disable                  *bool
	DisableDiscovery         *bool
	DnsAliases               []string
	EnableImmediateDiscovery *bool
	ExtAttrs                 map[string]any
	Ipv4addrs                []niosdns.RecordHostIpv4addr
	Ipv6addrs                []niosdns.RecordHostIpv6addr
	Name                     *string
	NetworkView              *string
	RestartIfNeeded          *bool
	RrsetOrder               *string
	Snmp3Credential          *niosdns.RecordHostSnmp3Credential
	SnmpCredential           *niosdns.RecordHostSnmpCredential
	Ttl                      *int64
	UseCliCredentials        *bool
	UseDnsEaInheritance      *bool
	UseSnmp3Credential       *bool
	UseSnmpCredential        *bool
	UseTtl                   *bool
	View                     *string
}
