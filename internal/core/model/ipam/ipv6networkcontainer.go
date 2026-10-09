package ipam

import (
	niosipam "github.com/infobloxopen/infoblox-nios-go-client/ipam"
	uddiipam "github.com/infobloxopen/universal-ddi-go-client/ipam"
)

// Infoblox Ipv6networkcontainer model
type Ipv6networkcontainer struct {
	Id   *string
	NIOS *NIOSIpv6networkcontainerExt
	UDDI *UDDIIpv6networkcontainerExt
}

// NIOSIpv6networkcontainerExt - NIOS specific fields for Ipv6networkcontainer
type NIOSIpv6networkcontainerExt struct {
	AutoCreateReversezone            *bool
	CloudInfo                        *niosipam.Ipv6networkcontainerCloudInfo
	Comment                          *string
	DdnsDomainname                   *string
	DdnsEnableOptionFqdn             *bool
	DdnsGenerateHostname             *bool
	DdnsServerAlwaysUpdates          *bool
	DdnsTtl                          *int64
	DeleteReason                     *string
	DiscoveryBasicPollSettings       *niosipam.Ipv6networkcontainerDiscoveryBasicPollSettings
	DiscoveryBlackoutSetting         *niosipam.Ipv6networkcontainerDiscoveryBlackoutSetting
	DiscoveryMember                  *string
	DomainNameServers                []string
	EnableDdns                       *bool
	EnableDiscovery                  *bool
	EnableImmediateDiscovery         *bool
	ExtAttrs                         map[string]any
	FederatedRealms                  []niosipam.Ipv6networkcontainerFederatedRealms
	LogicFilterRules                 []niosipam.Ipv6networkcontainerLogicFilterRules
	MgmPrivate                       *bool
	Network                          *string
	NetworkView                      *string
	Options                          []niosipam.Ipv6networkcontainerOptions
	PortControlBlackoutSetting       *niosipam.Ipv6networkcontainerPortControlBlackoutSetting
	PreferredLifetime                *int64
	RestartIfNeeded                  *bool
	RirOrganization                  *string
	RirRegistrationAction            *string
	RirRegistrationStatus            *string
	SamePortControlDiscoveryBlackout *bool
	SendRirRequest                   *bool
	SubscribeSettings                *niosipam.Ipv6networkcontainerSubscribeSettings
	Unmanaged                        *bool
	UpdateDnsOnLeaseRenewal          *bool
	UseBlackoutSetting               *bool
	UseDdnsDomainname                *bool
	UseDdnsEnableOptionFqdn          *bool
	UseDdnsGenerateHostname          *bool
	UseDdnsTtl                       *bool
	UseDiscoveryBasicPollingSettings *bool
	UseDomainNameServers             *bool
	UseEnableDdns                    *bool
	UseEnableDiscovery               *bool
	UseLogicFilterRules              *bool
	UseMgmPrivate                    *bool
	UseOptions                       *bool
	UsePreferredLifetime             *bool
	UseSubscribeSettings             *bool
	UseUpdateDnsOnLeaseRenewal       *bool
	UseValidLifetime                 *bool
	UseZoneAssociations              *bool
	ValidLifetime                    *int64
	ZoneAssociations                 []niosipam.Ipv6networkcontainerZoneAssociations
	FuncCall                         *niosipam.FuncCall
}

// UDDIIpv6networkcontainerExt - UDDI specific fields for Ipv6networkcontainer
type UDDIIpv6networkcontainerExt struct {
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
