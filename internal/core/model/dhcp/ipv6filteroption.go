package dhcp

import (
	niosdhcp "github.com/infobloxopen/infoblox-nios-go-client/dhcp"
	uddiipam "github.com/infobloxopen/universal-ddi-go-client/ipam"
)

// Infoblox Ipv6filteroption model
type Ipv6filteroption struct {
	Id   *string
	NIOS *NIOSIpv6filteroptionExt
	UDDI *UDDIIpv6filteroptionExt
}

// NIOSIpv6filteroptionExt - NIOS specific fields for Ipv6filteroption
type NIOSIpv6filteroptionExt struct {
	ApplyAsClass *bool
	Comment      *string
	Expression   *string
	ExtAttrs     map[string]any
	LeaseTime    *int64
	Name         *string
	OptionList   []niosdhcp.Ipv6filteroptionOptionList
	OptionSpace  *string
}

// UDDIIpv6filteroptionExt - UDDI specific fields for Ipv6filteroption
type UDDIIpv6filteroptionExt struct {
	Comment                         *string
	DhcpOptions                     []uddiipam.OptionItem
	HeaderOptionFilename            *string
	HeaderOptionServerAddress       *string
	HeaderOptionServerName          *string
	LeaseTime                       *int64
	Name                            string
	Protocol                        *string
	Role                            *string
	Rules                           *uddiipam.OptionFilterRuleList
	Tags                            map[string]any
	VendorSpecificOptionOptionSpace *string
}
