package dhcp

import (
	uddiipam "github.com/infobloxopen/universal-ddi-go-client/ipam"
)

// Infoblox HaGroup model
type HaGroup struct {
	Id   *string
	UDDI *UDDIHaGroupExt
}

// UDDIHaGroupExt - UDDI specific fields for HaGroup
type UDDIHaGroupExt struct {
	AnycastConfigId *string
	Comment         *string
	Hosts           []uddiipam.HAGroupHost
	IpSpace         *string
	Mode            *string
	Name            string
	Status          *string
	StatusV6        *string
	Tags            map[string]any
}
