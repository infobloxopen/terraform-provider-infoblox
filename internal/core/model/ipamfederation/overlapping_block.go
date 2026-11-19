package ipamfederation

import (
	"time"
)

// Infoblox OverlappingBlock model
type OverlappingBlock struct {
	Id   *string
	UDDI *UDDIOverlappingBlockExt
}

// UDDIOverlappingBlockExt - UDDI specific fields for OverlappingBlock
type UDDIOverlappingBlockExt struct {
	Address          *string
	Cidr             *int64
	Comment          *string
	CreatedAt        *time.Time
	FederatedPoolId  *string
	FederatedRealm   string
	Name             *string
	NetworkCompliant *bool
	Parent           *string
	Protocol         *string
	Tags             map[string]any
	UpdatedAt        *time.Time
}
