package ipam

import (
	niosipam "github.com/infobloxopen/infoblox-nios-go-client/ipam"
	uddiipam "github.com/infobloxopen/universal-ddi-go-client/ipam"
)

// Infoblox Networkview model
type Networkview struct {
	Id   *string
	NIOS *NIOSNetworkviewExt
	UDDI *UDDINetworkviewExt
}

// NIOSNetworkviewExt - NIOS specific fields for Networkview
type NIOSNetworkviewExt struct {
	CloudInfo            *niosipam.NetworkviewCloudInfo
	Comment              *string
	DdnsDnsView          *string
	DdnsZonePrimaries    []niosipam.NetworkviewDdnsZonePrimaries
	ExtAttrs             map[string]any
	FederatedRealms      []niosipam.NetworkviewFederatedRealms
	InternalForwardZones []string
	MgmPrivate           *bool
	Name                 *string
	RemoteForwardZones   []niosipam.NetworkviewRemoteForwardZones
	RemoteReverseZones   []niosipam.NetworkviewRemoteReverseZones
}

// UDDINetworkviewExt - UDDI specific fields for Networkview
type UDDINetworkviewExt struct {
	AsmConfig                       *uddiipam.ASMConfig
	Comment                         *string
	CompartmentId                   *string
	DdnsClientUpdate                *string
	DdnsConflictResolutionMode      *string
	DdnsDomain                      *string
	DdnsGenerateName                *bool
	DdnsGeneratedPrefix             *string
	DdnsSendUpdates                 *bool
	DdnsTtlPercent                  *float32
	DdnsUpdateOnRenew               *bool
	DdnsUseConflictResolution       *bool
	DefaultRealms                   []string
	DhcpConfig                      *uddiipam.DHCPConfig
	DhcpOptions                     []uddiipam.OptionItem
	DhcpOptionsV6                   []uddiipam.OptionItem
	HeaderOptionFilename            *string
	HeaderOptionServerAddress       *string
	HeaderOptionServerName          *string
	HostnameRewriteChar             *string
	HostnameRewriteEnabled          *bool
	HostnameRewriteRegex            *string
	InheritanceSources              *uddiipam.IPSpaceInheritance
	Name                            string
	Tags                            map[string]any
	VendorSpecificOptionOptionSpace *string
}
