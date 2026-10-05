package fw

// Infoblox CategoryFilter model
type CategoryFilter struct {
	Id   *int32
	UDDI *UDDICategoryFilterExt
}

// UDDICategoryFilterExt - UDDI specific fields for CategoryFilter
type UDDICategoryFilterExt struct {
	Categories  []string
	Description *string
	Name        *string
	Policies    []string
	Tags        map[string]any
}
