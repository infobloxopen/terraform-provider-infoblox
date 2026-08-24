package dhcp

// Infoblox Ipv6DhcpOptionspace model
type Ipv6DhcpOptionspace struct {
	Id   *string
	NIOS *NIOSIpv6DhcpOptionspaceExt
	UDDI *UDDIIpv6DhcpOptionspaceExt
}

// NIOSIpv6DhcpOptionspaceExt - NIOS specific fields for Ipv6DhcpOptionspace
type NIOSIpv6DhcpOptionspaceExt struct {
	Comment           *string
	EnterpriseNumber  *int64
	Name              *string
	OptionDefinitions []string
}

// UDDIIpv6DhcpOptionspaceExt - UDDI specific fields for Ipv6DhcpOptionspace
type UDDIIpv6DhcpOptionspaceExt struct {
	Comment  *string
	Name     string
	Protocol *string
	Tags     map[string]any
}
