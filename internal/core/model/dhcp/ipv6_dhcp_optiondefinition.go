package dhcp

// Infoblox Ipv6DhcpOptiondefinition model
type Ipv6DhcpOptiondefinition struct {
	Id   *string
	NIOS *NIOSIpv6DhcpOptiondefinitionExt
	UDDI *UDDIIpv6DhcpOptiondefinitionExt
}

// NIOSIpv6DhcpOptiondefinitionExt - NIOS specific fields for Ipv6DhcpOptiondefinition
type NIOSIpv6DhcpOptiondefinitionExt struct {
	Code  *int64
	Name  *string
	Space *string
	Type  *string
}

// UDDIIpv6DhcpOptiondefinitionExt - UDDI specific fields for Ipv6DhcpOptiondefinition
type UDDIIpv6DhcpOptiondefinitionExt struct {
	Array       *bool
	Code        int64
	Comment     *string
	Name        string
	OptionSpace string
	Type        string
}
