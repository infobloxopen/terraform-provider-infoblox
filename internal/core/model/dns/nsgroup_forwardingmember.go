package dns

import (
	niosdns "github.com/infobloxopen/infoblox-nios-go-client/dns"
)

// Infoblox NsgroupForwardingmember model
type NsgroupForwardingmember struct {
	Id   *string
	NIOS *NIOSNsgroupForwardingmemberExt
}

// NIOSNsgroupForwardingmemberExt - NIOS specific fields for NsgroupForwardingmember
type NIOSNsgroupForwardingmemberExt struct {
	Comment           *string
	ExtAttrs          map[string]any
	ForwardingServers []niosdns.NsgroupForwardingmemberForwardingServers
	Name              *string
}
