package dns

import (
	niosdns "github.com/infobloxopen/infoblox-nios-go-client/dns"
)

// Infoblox NsgroupDelegation model
type NsgroupDelegation struct {
	Id   *string
	NIOS *NIOSNsgroupDelegationExt
}

// NIOSNsgroupDelegationExt - NIOS specific fields for NsgroupDelegation
type NIOSNsgroupDelegationExt struct {
	Comment    *string
	DelegateTo []niosdns.NsgroupDelegationDelegateTo
	ExtAttrs   map[string]any
	Name       *string
}
