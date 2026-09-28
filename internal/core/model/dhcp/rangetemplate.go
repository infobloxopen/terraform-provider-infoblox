package dhcp

import (
	niosdhcp "github.com/infobloxopen/infoblox-nios-go-client/dhcp"
)

// Infoblox Rangetemplate model
type Rangetemplate struct {
	Id   *string
	NIOS *NIOSRangetemplateExt
}

// NIOSRangetemplateExt - NIOS specific fields for Rangetemplate
type NIOSRangetemplateExt struct {
	Bootfile                       *string
	Bootserver                     *string
	CloudApiCompatible             *bool
	Comment                        *string
	DdnsDomainname                 *string
	DdnsGenerateHostname           *bool
	DelegatedMember                *niosdhcp.RangetemplateDelegatedMember
	DenyAllClients                 *bool
	DenyBootp                      *bool
	EmailList                      []string
	EnableDdns                     *bool
	EnableDhcpThresholds           *bool
	EnableEmailWarnings            *bool
	EnablePxeLeaseTime             *bool
	EnableSnmpWarnings             *bool
	Exclude                        []niosdhcp.RangetemplateExclude
	ExtAttrs                       map[string]any
	FailoverAssociation            *string
	FingerprintFilterRules         []niosdhcp.RangetemplateFingerprintFilterRules
	HighWaterMark                  *int64
	HighWaterMarkReset             *int64
	IgnoreDhcpOptionListRequest    *bool
	KnownClients                   *string
	LeaseScavengeTime              *int64
	LogicFilterRules               []niosdhcp.RangetemplateLogicFilterRules
	LowWaterMark                   *int64
	LowWaterMarkReset              *int64
	MacFilterRules                 []niosdhcp.RangetemplateMacFilterRules
	Member                         *niosdhcp.RangetemplateMember
	MsOptions                      []niosdhcp.RangetemplateMsOptions
	MsServer                       *niosdhcp.RangetemplateMsServer
	NacFilterRules                 []niosdhcp.RangetemplateNacFilterRules
	Name                           *string
	Nextserver                     *string
	NumberOfAddresses              *int64
	Offset                         *int64
	OptionFilterRules              []niosdhcp.RangetemplateOptionFilterRules
	Options                        []niosdhcp.RangetemplateOptions
	PxeLeaseTime                   *int64
	RecycleLeases                  *bool
	RelayAgentFilterRules          []niosdhcp.RangetemplateRelayAgentFilterRules
	ServerAssociationType          *string
	UnknownClients                 *string
	UpdateDnsOnLeaseRenewal        *bool
	UseBootfile                    *bool
	UseBootserver                  *bool
	UseDdnsDomainname              *bool
	UseDdnsGenerateHostname        *bool
	UseDenyBootp                   *bool
	UseEmailList                   *bool
	UseEnableDdns                  *bool
	UseEnableDhcpThresholds        *bool
	UseIgnoreDhcpOptionListRequest *bool
	UseKnownClients                *bool
	UseLeaseScavengeTime           *bool
	UseLogicFilterRules            *bool
	UseMsOptions                   *bool
	UseNextserver                  *bool
	UseOptions                     *bool
	UsePxeLeaseTime                *bool
	UseRecycleLeases               *bool
	UseUnknownClients              *bool
	UseUpdateDnsOnLeaseRenewal     *bool
}
