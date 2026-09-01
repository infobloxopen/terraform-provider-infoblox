package ipam

// Infoblox Superhost model
type Superhost struct {
	Id   *string
	NIOS *NIOSSuperhostExt
}

// NIOSSuperhostExt - NIOS specific fields for Superhost
type NIOSSuperhostExt struct {
	Comment                 *string
	DeleteAssociatedObjects *bool
	DhcpAssociatedObjects   []string
	Disabled                *bool
	DnsAssociatedObjects    []string
	ExtAttrs                map[string]any
	Name                    *string
}
