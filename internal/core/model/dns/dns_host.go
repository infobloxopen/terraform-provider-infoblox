package dns

import (
	uddidnsconfig "github.com/infobloxopen/universal-ddi-go-client/dnsconfig"
)

// Infoblox DnsHost model
type DnsHost struct {
	Id   *string
	UDDI *UDDIDnsHostExt
}

// UDDIDnsHostExt - UDDI specific fields for DnsHost
type UDDIDnsHostExt struct {
	AbsoluteName       *string
	Address            *string
	AnycastAddresses   []string
	AssociatedServer   *uddidnsconfig.HostAssociatedServer
	DfpService         *string
	InheritanceSources *uddidnsconfig.HostInheritance
	KerberosKeys       []uddidnsconfig.KerberosKey
	Name               *string
	Ophid              *string
	ProviderId         *string
	Server             *string
	Tags               map[string]any
	Type               *string
}
