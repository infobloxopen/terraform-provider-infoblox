package fw

import (
	uddifw "github.com/infobloxopen/universal-ddi-go-client/fw"
)

// Infoblox NetworkList model
type NetworkList struct {
	Id   *int32
	UDDI *UDDINetworkListExt
}

// UDDINetworkListExt - UDDI specific fields for NetworkList
type UDDINetworkListExt struct {
	AddrBlock   []uddifw.AddrBlock
	Description *string
	Items       []string
	Name        *string
}
