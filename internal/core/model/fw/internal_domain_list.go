package fw

// Infoblox InternalDomainList model
type InternalDomainList struct {
	Id   *int32
	UDDI *UDDIInternalDomainListExt
}

// UDDIInternalDomainListExt - UDDI specific fields for InternalDomainList
type UDDIInternalDomainListExt struct {
	Description     *string
	InternalDomains []string
	IsDefault       *bool
	Name            *string
	Tags            map[string]any
}
