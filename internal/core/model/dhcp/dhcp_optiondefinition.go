package dhcp

// Infoblox DhcpOptiondefinition model
type DhcpOptiondefinition struct {
	Id   *string
	NIOS *NIOSDhcpOptiondefinitionExt
	UDDI *UDDIDhcpOptiondefinitionExt
}

// NIOSDhcpOptiondefinitionExt - NIOS specific fields for DhcpOptiondefinition
type NIOSDhcpOptiondefinitionExt struct {
	Code  *int64
	Name  *string
	Space *string
	Type  *string
}

// UDDIDhcpOptiondefinitionExt - UDDI specific fields for DhcpOptiondefinition
type UDDIDhcpOptiondefinitionExt struct {
	Array       *bool
	Code        int64
	Comment     *string
	Name        string
	OptionSpace string
	Type        string
}
