package dhcp

// Infoblox DhcpHost model
type DhcpHost struct {
	Id   *string
	UDDI *UDDIDhcpHostExt
}

// UDDIDhcpHostExt - UDDI specific fields for DhcpHost
type UDDIDhcpHostExt struct {
	IpSpace *string
	Server  *string
	Tags    map[string]any
}
