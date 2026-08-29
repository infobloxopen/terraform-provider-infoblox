package ipamfederation

// Infoblox FederatedRealm model
type FederatedRealm struct {
	Id   *string
	UDDI *UDDIFederatedRealmExt
}

// UDDIFederatedRealmExt - UDDI specific fields for FederatedRealm
type UDDIFederatedRealmExt struct {
	Comment *string
	Name    string
	Tags    map[string]any
}
