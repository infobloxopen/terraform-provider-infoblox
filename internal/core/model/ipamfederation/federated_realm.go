package ipamfederation

import (
	uddiipamfederation "github.com/infobloxopen/universal-ddi-go-client/ipamfederation"
)

// Infoblox FederatedRealm model
type FederatedRealm struct {
	Id   *string
	UDDI *UDDIFederatedRealmExt
}

// UDDIFederatedRealmExt - UDDI specific fields for FederatedRealm
type UDDIFederatedRealmExt struct {
	Comment     *string
	Metadata    map[string]any
	Name        string
	Provider    *uddiipamfederation.ProviderType
	Region      *string
	Tags        map[string]any
	Utilization *int64
}
