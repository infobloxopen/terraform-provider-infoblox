package dhcp

import (
	niosdhcp "github.com/infobloxopen/infoblox-nios-go-client/dhcp"
)

// Infoblox Fixedaddresstemplate model
type Fixedaddresstemplate struct {
	Id   *string
	NIOS *NIOSFixedaddresstemplateExt
}

// NIOSFixedaddresstemplateExt - NIOS specific fields for Fixedaddresstemplate
type NIOSFixedaddresstemplateExt struct {
	Bootfile                       *string
	Bootserver                     *string
	Comment                        *string
	DdnsDomainname                 *string
	DdnsHostname                   *string
	DenyBootp                      *bool
	EnableDdns                     *bool
	EnablePxeLeaseTime             *bool
	ExtAttrs                       map[string]any
	IgnoreDhcpOptionListRequest    *bool
	LogicFilterRules               []niosdhcp.FixedaddresstemplateLogicFilterRules
	Name                           *string
	Nextserver                     *string
	NumberOfAddresses              *int64
	Offset                         *int64
	Options                        []niosdhcp.FixedaddresstemplateOptions
	PxeLeaseTime                   *int64
	UseBootfile                    *bool
	UseBootserver                  *bool
	UseDdnsDomainname              *bool
	UseDenyBootp                   *bool
	UseEnableDdns                  *bool
	UseIgnoreDhcpOptionListRequest *bool
	UseLogicFilterRules            *bool
	UseNextserver                  *bool
	UseOptions                     *bool
	UsePxeLeaseTime                *bool
}
