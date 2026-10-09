package dns

import (
	uddidnsconfig "github.com/infobloxopen/universal-ddi-go-client/dnsconfig"
)

// Infoblox ForwardNsg model
type ForwardNsg struct {
	Id   *string
	UDDI *UDDIForwardNsgExt
}

// UDDIForwardNsgExt - UDDI specific fields for ForwardNsg
type UDDIForwardNsgExt struct {
	Comment            *string
	ExternalForwarders []uddidnsconfig.Forwarder
	ForwardersOnly     *bool
	Hosts              []string
	InternalForwarders []string
	Name               string
	Nsgs               []string
	Tags               map[string]any
}
