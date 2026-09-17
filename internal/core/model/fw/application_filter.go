package fw

import (
	uddifw "github.com/infobloxopen/universal-ddi-go-client/fw"
)

// Infoblox ApplicationFilter model
type ApplicationFilter struct {
	Id   *int32
	UDDI *UDDIApplicationFilterExt
}

// UDDIApplicationFilterExt - UDDI specific fields for ApplicationFilter
type UDDIApplicationFilterExt struct {
	Criteria    []uddifw.ApplicationCriterion
	Description *string
	Name        *string
	Readonly    *bool
	Tags        map[string]any
}
