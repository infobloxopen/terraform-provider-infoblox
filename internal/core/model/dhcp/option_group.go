package dhcp

import (
	"time"

	uddiipam "github.com/infobloxopen/universal-ddi-go-client/ipam"
)

// Infoblox OptionGroup model
type OptionGroup struct {
	Id   *string
	UDDI *UDDIOptionGroupExt
}

// UDDIOptionGroupExt - UDDI specific fields for OptionGroup
type UDDIOptionGroupExt struct {
	Comment     *string
	CreatedAt   *time.Time
	DhcpOptions []uddiipam.OptionItem
	Name        string
	Protocol    *string
	Tags        map[string]any
	UpdatedAt   *time.Time
}
