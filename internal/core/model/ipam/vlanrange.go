package ipam

import (
	niosipam "github.com/infobloxopen/infoblox-nios-go-client/ipam"
)

// Infoblox Vlanrange model
type Vlanrange struct {
	Id   *string
	NIOS *NIOSVlanrangeExt
}

// NIOSVlanrangeExt - NIOS specific fields for Vlanrange
type NIOSVlanrangeExt struct {
	Comment        *string
	DeleteVlans    *bool
	EndVlanId      *int64
	ExtAttrs       map[string]any
	Name           *string
	PreCreateVlan  *bool
	StartVlanId    *int64
	VlanNamePrefix *string
	VlanView       *niosipam.VlanrangeVlanView
}
