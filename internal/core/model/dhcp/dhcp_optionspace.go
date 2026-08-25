package dhcp

// Infoblox DhcpOptionspace model
type DhcpOptionspace struct {
	Id   *string
	NIOS *NIOSDhcpOptionspaceExt
	UDDI *UDDIDhcpOptionspaceExt
}

// NIOSDhcpOptionspaceExt - NIOS specific fields for DhcpOptionspace
type NIOSDhcpOptionspaceExt struct {
	Comment           *string
	Name              *string
	OptionDefinitions []string
}

// UDDIDhcpOptionspaceExt - UDDI specific fields for DhcpOptionspace
type UDDIDhcpOptionspaceExt struct {
	Comment  *string
	Name     string
	Protocol *string
	Tags     map[string]any
}
