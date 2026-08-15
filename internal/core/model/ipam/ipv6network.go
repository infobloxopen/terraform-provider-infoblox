package ipam

import (
	niosipam "github.com/infobloxopen/infoblox-nios-go-client/ipam"
	uddiipam "github.com/infobloxopen/universal-ddi-go-client/ipam"
)

// Infoblox Ipv6network model
type Ipv6network struct {
	Id   *string
	NIOS *NIOSIpv6networkExt
	UDDI *UDDIIpv6networkExt
}

// NIOSIpv6networkExt - NIOS specific fields for Ipv6network
type NIOSIpv6networkExt struct {
	AutoCreateReversezone            *bool
	CloudInfo                        *niosipam.Ipv6networkCloudInfo
	Comment                          *string
	DdnsDomainname                   *string
	DdnsEnableOptionFqdn             *bool
	DdnsGenerateHostname             *bool
	DdnsServerAlwaysUpdates          *bool
	DdnsTtl                          *int64
	DeleteReason                     *string
	Disable                          *bool
	DiscoveredBridgeDomain           *string
	DiscoveredTenant                 *string
	DiscoveryBasicPollSettings       *niosipam.Ipv6networkDiscoveryBasicPollSettings
	DiscoveryBlackoutSetting         *niosipam.Ipv6networkDiscoveryBlackoutSetting
	DiscoveryMember                  *string
	DomainName                       *string
	DomainNameServers                []string
	EnableDdns                       *bool
	EnableDiscovery                  *bool
	EnableIfmapPublishing            *bool
	EnableImmediateDiscovery         *bool
	ExtAttrs                         map[string]any
	FederatedRealms                  []niosipam.Ipv6networkFederatedRealms
	LogicFilterRules                 []niosipam.Ipv6networkLogicFilterRules
	Members                          []niosipam.Ipv6networkMembers
	MgmPrivate                       *bool
	Network                          *string
	NetworkView                      *string
	Options                          []niosipam.Ipv6networkOptions
	PortControlBlackoutSetting       *niosipam.Ipv6networkPortControlBlackoutSetting
	PreferredLifetime                *int64
	RecycleLeases                    *bool
	RestartIfNeeded                  *bool
	RirOrganization                  *string
	RirRegistrationAction            *string
	RirRegistrationStatus            *string
	SamePortControlDiscoveryBlackout *bool
	SendRirRequest                   *bool
	SubscribeSettings                *niosipam.Ipv6networkSubscribeSettings
	Template                         *string
	Unmanaged                        *bool
	UpdateDnsOnLeaseRenewal          *bool
	UseBlackoutSetting               *bool
	UseDdnsDomainname                *bool
	UseDdnsEnableOptionFqdn          *bool
	UseDdnsGenerateHostname          *bool
	UseDdnsTtl                       *bool
	UseDiscoveryBasicPollingSettings *bool
	UseDomainName                    *bool
	UseDomainNameServers             *bool
	UseEnableDdns                    *bool
	UseEnableDiscovery               *bool
	UseEnableIfmapPublishing         *bool
	UseLogicFilterRules              *bool
	UseMgmPrivate                    *bool
	UseOptions                       *bool
	UsePreferredLifetime             *bool
	UseRecycleLeases                 *bool
	UseSubscribeSettings             *bool
	UseUpdateDnsOnLeaseRenewal       *bool
	UseValidLifetime                 *bool
	UseZoneAssociations              *bool
	ValidLifetime                    *int64
	Vlans                            []niosipam.Ipv6networkVlans
	ZoneAssociations                 []niosipam.Ipv6networkZoneAssociations
	FuncCall                         *niosipam.FuncCall
}

// UDDIIpv6networkExt - UDDI specific fields for Ipv6network
type UDDIIpv6networkExt struct {
	Address                    *string
	AsmConfig                  *uddiipam.ASMConfig
	Cidr                       *int64
	Comment                    *string
	ConfigProfiles             []string
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
	DhcpHost                   *string
	DhcpOptions                []uddiipam.OptionItem
	DisableDhcp                *bool
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
	RebindTime                 *int64
	RenewTime                  *int64
	Space                      *string
	Tags                       map[string]any
	Threshold                  *uddiipam.UtilizationThreshold
}
