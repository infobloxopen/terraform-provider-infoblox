package dns

import (
	niosdns "github.com/infobloxopen/infoblox-nios-go-client/dns"
)

// Infoblox NsgroupForwardstubserver model
type NsgroupForwardstubserver struct {
	Id   *string
	NIOS *NIOSNsgroupForwardstubserverExt
}

// NIOSNsgroupForwardstubserverExt - NIOS specific fields for NsgroupForwardstubserver
type NIOSNsgroupForwardstubserverExt struct {
	Comment         *string
	ExtAttrs        map[string]any
	ExternalServers []niosdns.NsgroupForwardstubserverExternalServers
	Name            *string
}
