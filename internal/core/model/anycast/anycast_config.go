package anycast

import (
	"time"

	uddianycast "github.com/infobloxopen/universal-ddi-go-client/anycast"
)

// Infoblox AnycastConfig model
type AnycastConfig struct {
	Id   *int64
	UDDI *UDDIAnycastConfigExt
}

// UDDIAnycastConfigExt - UDDI specific fields for AnycastConfig
type UDDIAnycastConfigExt struct {
	AccountId          *int64
	AnycastIpAddress   *string
	AnycastIpv6Address *string
	CreatedAt          *time.Time
	Description        *string
	Fields             *uddianycast.ProtobufFieldMask
	IsConfigured       *bool
	Name               *string
	OnpremHosts        []uddianycast.OnpremHostRef
	RuntimeStatus      *string
	Service            *string
	Tags               map[string]any
	UpdatedAt          *time.Time
}
