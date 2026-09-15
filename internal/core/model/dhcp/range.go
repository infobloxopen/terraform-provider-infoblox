package dhcp

import (
	niosdhcp "github.com/infobloxopen/infoblox-nios-go-client/dhcp"
	uddiipam "github.com/infobloxopen/universal-ddi-go-client/ipam"
)

// Infoblox Range model
type Range struct {
	Id   *string
	NIOS *NIOSRangeExt
	UDDI *UDDIRangeExt
}

// NIOSRangeExt - NIOS specific fields for Range
type NIOSRangeExt struct {
	AlwaysUpdateDns                  *bool
	Bootfile                         *string
	Bootserver                       *string
	CloudInfo                        *niosdhcp.RangeCloudInfo
	Comment                          *string
	DdnsDomainname                   *string
	DdnsGenerateHostname             *bool
	DenyAllClients                   *bool
	DenyBootp                        *bool
	Disable                          *bool
	DiscoveryBasicPollSettings       *niosdhcp.RangeDiscoveryBasicPollSettings
	DiscoveryBlackoutSetting         *niosdhcp.RangeDiscoveryBlackoutSetting
	DiscoveryMember                  *string
	EmailList                        []string
	EnableDdns                       *bool
	EnableDhcpThresholds             *bool
	EnableDiscovery                  *bool
	EnableEmailWarnings              *bool
	EnableIfmapPublishing            *bool
	EnableImmediateDiscovery         *bool
	EnablePxeLeaseTime               *bool
	EnableSnmpWarnings               *bool
	EndAddr                          *string
	Exclude                          []niosdhcp.RangeExclude
	ExtAttrs                         map[string]any
	FailoverAssociation              *string
	FingerprintFilterRules           []niosdhcp.RangeFingerprintFilterRules
	HighWaterMark                    *int64
	HighWaterMarkReset               *int64
	IgnoreDhcpOptionListRequest      *bool
	IgnoreId                         *string
	IgnoreMacAddresses               []string
	KnownClients                     *string
	LeaseScavengeTime                *int64
	LogicFilterRules                 []niosdhcp.RangeLogicFilterRules
	LowWaterMark                     *int64
	LowWaterMarkReset                *int64
	MacFilterRules                   []niosdhcp.RangeMacFilterRules
	Member                           *niosdhcp.RangeMember
	MsOptions                        []niosdhcp.RangeMsOptions
	MsServer                         *niosdhcp.RangeMsServer
	NacFilterRules                   []niosdhcp.RangeNacFilterRules
	Name                             *string
	Network                          *string
	NetworkView                      *string
	Nextserver                       *string
	OptionFilterRules                []niosdhcp.RangeOptionFilterRules
	Options                          []niosdhcp.RangeOptions
	PortControlBlackoutSetting       *niosdhcp.RangePortControlBlackoutSetting
	PxeLeaseTime                     *int64
	RecycleLeases                    *bool
	RelayAgentFilterRules            []niosdhcp.RangeRelayAgentFilterRules
	RestartIfNeeded                  *bool
	SamePortControlDiscoveryBlackout *bool
	ServerAssociationType            *string
	SplitMember                      *niosdhcp.RangeSplitMember
	SplitScopeExclusionPercent       *int64
	StartAddr                        *string
	SubscribeSettings                *niosdhcp.RangeSubscribeSettings
	Template                         *string
	UnknownClients                   *string
	UpdateDnsOnLeaseRenewal          *bool
	UseBlackoutSetting               *bool
	UseBootfile                      *bool
	UseBootserver                    *bool
	UseDdnsDomainname                *bool
	UseDdnsGenerateHostname          *bool
	UseDenyBootp                     *bool
	UseDiscoveryBasicPollingSettings *bool
	UseEmailList                     *bool
	UseEnableDdns                    *bool
	UseEnableDhcpThresholds          *bool
	UseEnableDiscovery               *bool
	UseEnableIfmapPublishing         *bool
	UseIgnoreDhcpOptionListRequest   *bool
	UseIgnoreId                      *bool
	UseKnownClients                  *bool
	UseLeaseScavengeTime             *bool
	UseLogicFilterRules              *bool
	UseMsOptions                     *bool
	UseNextserver                    *bool
	UseOptions                       *bool
	UsePxeLeaseTime                  *bool
	UseRecycleLeases                 *bool
	UseSubscribeSettings             *bool
	UseUnknownClients                *bool
	UseUpdateDnsOnLeaseRenewal       *bool
}

// UDDIRangeExt - UDDI specific fields for Range
type UDDIRangeExt struct {
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
