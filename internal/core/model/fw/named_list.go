package fw

import (
	uddifw "github.com/infobloxopen/universal-ddi-go-client/fw"
)

// Infoblox NamedList model
type NamedList struct {
	Id   *int32
	UDDI *UDDINamedListExt
}

// UDDINamedListExt - UDDI specific fields for NamedList
type UDDINamedListExt struct {
	ConfidenceLevel *string
	Description     *string
	Items           []string
	ItemsDescribed  []uddifw.ItemStructs
	Name            *string
	Policies        []string
	Tags            map[string]any
	ThreatLevel     *string
	Type            *string
}
