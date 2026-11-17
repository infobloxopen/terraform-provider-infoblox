package ipamfederation

import (
	"time"
)

// Infoblox ReservedBlock model
type ReservedBlock struct {
	Id   *string
	UDDI *UDDIReservedBlockExt
}

// UDDIReservedBlockExt - UDDI specific fields for ReservedBlock
type UDDIReservedBlockExt struct {
	Address          *string
	Cidr             *int64
	Comment          *string
	CreatedAt        *time.Time
	FederatedPoolId  *string
	FederatedRealm   string
	Metadata         map[string]any
	Name             *string
	NetworkCompliant *bool
	Parent           *string
	Protocol         *string
	Region           *string
	Tags             map[string]any
	UpdatedAt        *time.Time
}
