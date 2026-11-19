package ipamfederation

import (
	"time"

	uddiipamfederation "github.com/infobloxopen/universal-ddi-go-client/ipamfederation"
)

// Infoblox FederatedBlock model
type FederatedBlock struct {
	Id   *string
	UDDI *UDDIFederatedBlockExt
}

// UDDIFederatedBlockExt - UDDI specific fields for FederatedBlock
type UDDIFederatedBlockExt struct {
	Address           *string
	AllocationV4      *uddiipamfederation.Allocation
	Cidr              *int64
	Comment           *string
	CreatedAt         *time.Time
	FederatedPoolId   *string
	FederatedRealm    string
	Metadata          map[string]any
	Name              *string
	NetworkCompliance *uddiipamfederation.NetworkCompliance
	NetworkCompliant  *bool
	Parent            *string
	Protocol          *string
	Region            *string
	State             *string
	Tags              map[string]any
	UpdatedAt         *time.Time
	Utilization       *int64
	UtilizationV6     *uddiipamfederation.UtilizationV6
}
