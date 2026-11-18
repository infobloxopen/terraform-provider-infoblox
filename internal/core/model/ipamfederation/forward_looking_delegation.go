package ipamfederation

// Infoblox ForwardLookingDelegation model
type ForwardLookingDelegation struct {
	Id   *string
	UDDI *UDDIForwardLookingDelegationExt
}

// UDDIForwardLookingDelegationExt - UDDI specific fields for ForwardLookingDelegation
type UDDIForwardLookingDelegationExt struct {
	Address          *string
	Cidr             *int64
	Comment          *string
	FederatedPoolId  *string
	FederatedRealms  []string
	Name             *string
	NetworkCompliant *bool
	Protocol         *string
	Tags             map[string]any
}
