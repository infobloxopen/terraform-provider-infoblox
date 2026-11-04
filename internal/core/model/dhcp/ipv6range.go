package dhcp

import (
	niosdhcp "github.com/infobloxopen/infoblox-nios-go-client/dhcp"
	uddiipam "github.com/infobloxopen/universal-ddi-go-client/ipam"
)

// Infoblox Ipv6range model
type Ipv6range struct {
	Id   *string
	NIOS *NIOSIpv6rangeExt
	UDDI *UDDIIpv6rangeExt
}

// NIOSIpv6rangeExt - NIOS specific fields for Ipv6range
type NIOSIpv6rangeExt struct {
	AddressType                      *string
	CloudInfo                        *niosdhcp.Ipv6rangeCloudInfo
	Comment                          *string
	Disable                          *bool
	DiscoveryBasicPollSettings       *niosdhcp.Ipv6rangeDiscoveryBasicPollSettings
	DiscoveryBlackoutSetting         *niosdhcp.Ipv6rangeDiscoveryBlackoutSetting
	DiscoveryMember                  *string
	EnableDiscovery                  *bool
	EnableImmediateDiscovery         *bool
	EndAddr                          *string
	Exclude                          []niosdhcp.Ipv6rangeExclude
	ExtAttrs                         map[string]any
	Ipv6EndPrefix                    *string
	Ipv6PrefixBits                   *int64
	Ipv6StartPrefix                  *string
	LogicFilterRules                 []niosdhcp.Ipv6rangeLogicFilterRules
	Member                           *niosdhcp.Ipv6rangeMember
	Name                             *string
	Network                          *string
	NetworkView                      *string
	OptionFilterRules                []niosdhcp.Ipv6rangeOptionFilterRules
	PortControlBlackoutSetting       *niosdhcp.Ipv6rangePortControlBlackoutSetting
	RecycleLeases                    *bool
	RestartIfNeeded                  *bool
	SamePortControlDiscoveryBlackout *bool
	ServerAssociationType            *string
	StartAddr                        *string
	SubscribeSettings                *niosdhcp.Ipv6rangeSubscribeSettings
	Template                         *string
	UseBlackoutSetting               *bool
	UseDiscoveryBasicPollingSettings *bool
	UseEnableDiscovery               *bool
	UseLogicFilterRules              *bool
	UseRecycleLeases                 *bool
	UseSubscribeSettings             *bool
}

// UDDIIpv6rangeExt - UDDI specific fields for Ipv6range
type UDDIIpv6rangeExt struct {
	Comment            *string
	DhcpHost           *string
	DhcpOptions        []uddiipam.OptionItem
	DisableDhcp        *bool
	End                string
	ExclusionRanges    []uddiipam.ExclusionRange
	Filters            []uddiipam.AccessFilter
	InheritanceParent  *string
	InheritanceSources *uddiipam.DHCPOptionsInheritance
	Name               *string
	Parent             *string
	Space              *string
	Start              string
	Tags               map[string]any
	Threshold          *uddiipam.UtilizationThreshold
}
