package dhcp

import (
	niosdhcp "github.com/infobloxopen/infoblox-nios-go-client/dhcp"
	uddiipam "github.com/infobloxopen/universal-ddi-go-client/ipam"
)

// Infoblox Filteroption model
type Filteroption struct {
	Id   *string
	NIOS *NIOSFilteroptionExt
	UDDI *UDDIFilteroptionExt
}

// NIOSFilteroptionExt - NIOS specific fields for Filteroption
type NIOSFilteroptionExt struct {
	ApplyAsClass *bool
	Bootfile     *string
	Bootserver   *string
	Comment      *string
	Expression   *string
	ExtAttrs     map[string]any
	LeaseTime    *int64
	Name         *string
	NextServer   *string
	OptionList   []niosdhcp.FilteroptionOptionList
	OptionSpace  *string
	PxeLeaseTime *int64
}

// UDDIFilteroptionExt - UDDI specific fields for Filteroption
type UDDIFilteroptionExt struct {
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
