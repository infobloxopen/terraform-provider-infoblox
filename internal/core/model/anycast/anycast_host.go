package anycast

import (
	"time"

	uddianycast "github.com/infobloxopen/universal-ddi-go-client/anycast"
)

// Infoblox AnycastHost model
type AnycastHost struct {
	Id   *int64
	UDDI *UDDIAnycastHostExt
}

// UDDIAnycastHostExt - UDDI specific fields for AnycastHost
type UDDIAnycastHostExt struct {
	AnycastConfigRefs []uddianycast.AnycastConfigRef
	ConfigBgp         *uddianycast.BgpConfig
	ConfigOspf        *uddianycast.OspfConfig
	ConfigOspfv3      *uddianycast.Ospfv3Config
	CreatedAt         *time.Time
	IpAddress         *string
	Ipv6Address       *string
	Name              *string
	UpdatedAt         *time.Time
}
