package dhcp

import (
	"time"

	uddiipam "github.com/infobloxopen/universal-ddi-go-client/ipam"
)

// Infoblox HardwareFilter model
type HardwareFilter struct {
	Id   *string
	UDDI *UDDIHardwareFilterExt
}

// UDDIHardwareFilterExt - UDDI specific fields for HardwareFilter
type UDDIHardwareFilterExt struct {
	Addresses                       []string
	Comment                         *string
	CreatedAt                       *time.Time
	DhcpOptions                     []uddiipam.OptionItem
	HeaderOptionFilename            *string
	HeaderOptionServerAddress       *string
	HeaderOptionServerName          *string
	LeaseTime                       *int64
	Name                            string
	Role                            *string
	Tags                            map[string]any
	UpdatedAt                       *time.Time
	VendorSpecificOptionOptionSpace *string
}
