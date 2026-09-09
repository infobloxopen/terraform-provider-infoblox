package dhcp

import (
	niosdhcp "github.com/infobloxopen/infoblox-nios-go-client/dhcp"
)

// Infoblox Ipv6rangetemplate model
type Ipv6rangetemplate struct {
	Id   *string
	NIOS *NIOSIpv6rangetemplateExt
}

// NIOSIpv6rangetemplateExt - NIOS specific fields for Ipv6rangetemplate
type NIOSIpv6rangetemplateExt struct {
	CloudApiCompatible    *bool
	Comment               *string
	DelegatedMember       *niosdhcp.Ipv6rangetemplateDelegatedMember
	Exclude               []niosdhcp.Ipv6rangetemplateExclude
	LogicFilterRules      []niosdhcp.Ipv6rangetemplateLogicFilterRules
	Member                *niosdhcp.Ipv6rangetemplateMember
	Name                  *string
	NumberOfAddresses     *int64
	Offset                *int64
	OptionFilterRules     []niosdhcp.Ipv6rangetemplateOptionFilterRules
	RecycleLeases         *bool
	ServerAssociationType *string
	UseLogicFilterRules   *bool
	UseRecycleLeases      *bool
}
