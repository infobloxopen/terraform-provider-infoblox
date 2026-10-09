package acl

import (
	niosacl "github.com/infobloxopen/infoblox-nios-go-client/acl"
	uddidnsconfig "github.com/infobloxopen/universal-ddi-go-client/dnsconfig"
)

// Infoblox Namedacl model
type Namedacl struct {
	Id   *string
	NIOS *NIOSNamedaclExt
	UDDI *UDDINamedaclExt
}

// NIOSNamedaclExt - NIOS specific fields for Namedacl
type NIOSNamedaclExt struct {
	AccessList []niosacl.NamedaclAccessList
	Comment    *string
	ExtAttrs   map[string]any
	Name       *string
}

// UDDINamedaclExt - UDDI specific fields for Namedacl
type UDDINamedaclExt struct {
	Comment       *string
	CompartmentId *string
	List          []uddidnsconfig.ACLItem
	Name          string
	Tags          map[string]any
}
