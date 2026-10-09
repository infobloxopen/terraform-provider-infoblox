package ipam

import (
	niosipam "github.com/infobloxopen/infoblox-nios-go-client/ipam"
	uddiipam "github.com/infobloxopen/universal-ddi-go-client/ipam"
)

// Infoblox Networkcontainer model
type Networkcontainer struct {
	Id   *string
	NIOS *NIOSNetworkcontainerExt
	UDDI *UDDINetworkcontainerExt
}

// NIOSNetworkcontainerExt - NIOS specific fields for Networkcontainer
type NIOSNetworkcontainerExt struct {
	Authority                        *bool
	AutoCreateReversezone            *bool
	Bootfile                         *string
	Bootserver                       *string
	CloudInfo                        *niosipam.NetworkcontainerCloudInfo
	Comment                          *string
	DdnsDomainname                   *string
	DdnsGenerateHostname             *bool
	DdnsServerAlwaysUpdates          *bool
	DdnsTtl                          *int64
	DdnsUpdateFixedAddresses         *bool
	DdnsUseOption81                  *bool
	DeleteReason                     *string
	DenyBootp                        *bool
	DiscoveryBasicPollSettings       *niosipam.NetworkcontainerDiscoveryBasicPollSettings
	DiscoveryBlackoutSetting         *niosipam.NetworkcontainerDiscoveryBlackoutSetting
	DiscoveryMember                  *string
	EmailList                        []string
	EnableDdns                       *bool
	EnableDhcpThresholds             *bool
	EnableDiscovery                  *bool
	EnableEmailWarnings              *bool
	EnableImmediateDiscovery         *bool
	EnablePxeLeaseTime               *bool
	EnableSnmpWarnings               *bool
	ExtAttrs                         map[string]any
	FederatedRealms                  []niosipam.NetworkcontainerFederatedRealms
	HighWaterMark                    *int64
	HighWaterMarkReset               *int64
	IgnoreDhcpOptionListRequest      *bool
	IgnoreId                         *string
	IgnoreMacAddresses               []string
	IpamEmailAddresses               []string
	IpamThresholdSettings            *niosipam.NetworkcontainerIpamThresholdSettings
	IpamTrapSettings                 *niosipam.NetworkcontainerIpamTrapSettings
	LeaseScavengeTime                *int64
	LogicFilterRules                 []niosipam.NetworkcontainerLogicFilterRules
	LowWaterMark                     *int64
	LowWaterMarkReset                *int64
	MgmPrivate                       *bool
	Network                          *string
	NetworkView                      *string
	Nextserver                       *string
	Options                          []niosipam.NetworkcontainerOptions
	PortControlBlackoutSetting       *niosipam.NetworkcontainerPortControlBlackoutSetting
	PxeLeaseTime                     *int64
	RecycleLeases                    *bool
	RestartIfNeeded                  *bool
	RirOrganization                  *string
	RirRegistrationAction            *string
	RirRegistrationStatus            *string
	SamePortControlDiscoveryBlackout *bool
	SendRirRequest                   *bool
	SubscribeSettings                *niosipam.NetworkcontainerSubscribeSettings
	Unmanaged                        *bool
	UpdateDnsOnLeaseRenewal          *bool
	UseAuthority                     *bool
	UseBlackoutSetting               *bool
	UseBootfile                      *bool
	UseBootserver                    *bool
	UseDdnsDomainname                *bool
	UseDdnsGenerateHostname          *bool
	UseDdnsTtl                       *bool
	UseDdnsUpdateFixedAddresses      *bool
	UseDdnsUseOption81               *bool
	UseDenyBootp                     *bool
	UseDiscoveryBasicPollingSettings *bool
	UseEmailList                     *bool
	UseEnableDdns                    *bool
	UseEnableDhcpThresholds          *bool
	UseEnableDiscovery               *bool
	UseIgnoreDhcpOptionListRequest   *bool
	UseIgnoreId                      *bool
	UseIpamEmailAddresses            *bool
	UseIpamThresholdSettings         *bool
	UseIpamTrapSettings              *bool
	UseLeaseScavengeTime             *bool
	UseLogicFilterRules              *bool
	UseMgmPrivate                    *bool
	UseNextserver                    *bool
	UseOptions                       *bool
	UsePxeLeaseTime                  *bool
	UseRecycleLeases                 *bool
	UseSubscribeSettings             *bool
	UseUpdateDnsOnLeaseRenewal       *bool
	UseZoneAssociations              *bool
	ZoneAssociations                 []niosipam.NetworkcontainerZoneAssociations
	FuncCall                         *niosipam.FuncCall
}

// UDDINetworkcontainerExt - UDDI specific fields for Networkcontainer
type UDDINetworkcontainerExt struct {
	Address                    *string
	AsmConfig                  *uddiipam.ASMConfig
	Cidr                       *int64
	Comment                    *string
	CompartmentId              *string
	DdnsClientUpdate           *string
	DdnsConflictResolutionMode *string
	DdnsDomain                 *string
	DdnsGenerateName           *bool
	DdnsGeneratedPrefix        *string
	DdnsSendUpdates            *bool
	DdnsTtlPercent             *float32
	DdnsUpdateOnRenew          *bool
	DdnsUseConflictResolution  *bool
	DhcpConfig                 *uddiipam.DHCPConfig
	DhcpOptions                []uddiipam.OptionItem
	ExternalKeys               map[string]any
	FederatedRealms            []string
	HeaderOptionFilename       *string
	HeaderOptionServerAddress  *string
	HeaderOptionServerName     *string
	HostnameRewriteChar        *string
	HostnameRewriteEnabled     *bool
	HostnameRewriteRegex       *string
	InheritanceParent          *string
	InheritanceSources         *uddiipam.DHCPInheritance
	Name                       *string
	Parent                     *string
	Space                      *string
	Tags                       map[string]any
	Threshold                  *uddiipam.UtilizationThreshold
}
