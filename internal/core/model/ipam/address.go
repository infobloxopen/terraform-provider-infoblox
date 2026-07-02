package ipam

import (
	uddiipam "github.com/infobloxopen/bloxone-go-client/ipam"
)

// Unified Address model
type Address struct {
	Id   *string
	UDDI *UDDIAddressExt
}

// UDDIAddressExt - UDDI specific fields for Address
type UDDIAddressExt struct {
	Address      string
	Comment      *string
	ExternalKeys map[string]any
	Host         *string
	Hwaddr       *string
	Interface    *string
	Names        []uddiipam.Name
	Parent       *string
	Range        *string
	Space        *string
	Tags         map[string]any
}
