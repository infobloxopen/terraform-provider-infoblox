package ipamfederation

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
	Provider    *string
	Region      *string
	Tags        map[string]any
	Utilization *int64
}
