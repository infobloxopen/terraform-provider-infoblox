package ipam

import (
	niosipam "github.com/infobloxopen/infoblox-nios-go-client/ipam"
)

// Infoblox Vlan model
type Vlan struct {
	Id   *string
	NIOS *NIOSVlanExt
}

// NIOSVlanExt - NIOS specific fields for Vlan
type NIOSVlanExt struct {
	Comment     *string
	Contact     *string
	Department  *string
	Description *string
	ExtAttrs    map[string]any
	Id          *niosipam.VlanId
	Name        *string
	Parent      *niosipam.VlanParent
	Reserved    *bool
}
