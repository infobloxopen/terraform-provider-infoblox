package ipam

import (
	uddiipam "github.com/infobloxopen/universal-ddi-go-client/ipam"
)

// Infoblox IpamHost model
type IpamHost struct {
	Id   *string
	UDDI *UDDIIpamHostExt
}

// UDDIIpamHostExt - UDDI specific fields for IpamHost
type UDDIIpamHostExt struct {
	Addresses           []uddiipam.HostAddress
	AutoGenerateRecords *bool
	Comment             *string
	HostNames           []uddiipam.HostName
	Name                string
	Tags                map[string]any
}
