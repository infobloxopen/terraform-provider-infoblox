package dhcp

// Infoblox DhcpHost model
type DhcpHost struct {
	Id   *string
	UDDI *UDDIDhcpHostExt
}

// UDDIDhcpHostExt - UDDI specific fields for DhcpHost
type UDDIDhcpHostExt struct {
	Address          *string
	AnycastAddresses []string
	Comment          *string
	CurrentVersion   *string
	IpSpace          *string
	Name             *string
	Ophid            *string
	ProviderId       *string
	Server           *string
	Tags             map[string]any
	Type             *string
}
