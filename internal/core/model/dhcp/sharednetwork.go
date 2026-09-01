package dhcp

import (
	niosdhcp "github.com/infobloxopen/infoblox-nios-go-client/dhcp"
)

// Infoblox Sharednetwork model
type Sharednetwork struct {
	Id   *string
	NIOS *NIOSSharednetworkExt
}

// NIOSSharednetworkExt - NIOS specific fields for Sharednetwork
type NIOSSharednetworkExt struct {
	Authority                      *bool
	Bootfile                       *string
	Bootserver                     *string
	Comment                        *string
	DdnsGenerateHostname           *bool
	DdnsServerAlwaysUpdates        *bool
	DdnsTtl                        *int64
	DdnsUpdateFixedAddresses       *bool
	DdnsUseOption81                *bool
	DenyBootp                      *bool
	Disable                        *bool
	EnableDdns                     *bool
	EnablePxeLeaseTime             *bool
	ExtAttrs                       map[string]any
	IgnoreClientIdentifier         *bool
	IgnoreDhcpOptionListRequest    *bool
	IgnoreId                       *string
	IgnoreMacAddresses             []string
	LeaseScavengeTime              *int64
	LogicFilterRules               []niosdhcp.SharednetworkLogicFilterRules
	Name                           *string
	NetworkView                    *string
	Networks                       []niosdhcp.SharednetworkNetworks
	Nextserver                     *string
	Options                        []niosdhcp.SharednetworkOptions
	PxeLeaseTime                   *int64
	UpdateDnsOnLeaseRenewal        *bool
	UseAuthority                   *bool
	UseBootfile                    *bool
	UseBootserver                  *bool
	UseDdnsGenerateHostname        *bool
	UseDdnsTtl                     *bool
	UseDdnsUpdateFixedAddresses    *bool
	UseDdnsUseOption81             *bool
	UseDenyBootp                   *bool
	UseEnableDdns                  *bool
	UseIgnoreClientIdentifier      *bool
	UseIgnoreDhcpOptionListRequest *bool
	UseIgnoreId                    *bool
	UseLeaseScavengeTime           *bool
	UseLogicFilterRules            *bool
	UseNextserver                  *bool
	UseOptions                     *bool
	UsePxeLeaseTime                *bool
	UseUpdateDnsOnLeaseRenewal     *bool
}
