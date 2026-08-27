package grid

import (
	niosgrid "github.com/infobloxopen/infoblox-nios-go-client/grid"
)

// Infoblox Extensibleattributedef model
type Extensibleattributedef struct {
	Id   *string
	NIOS *NIOSExtensibleattributedefExt
}

// NIOSExtensibleattributedefExt - NIOS specific fields for Extensibleattributedef
type NIOSExtensibleattributedefExt struct {
	AllowedObjectTypes []string
	Comment            *string
	DefaultValue       *niosgrid.ExtensibleattributedefDefaultValue
	DescendantsAction  *niosgrid.ExtensibleattributedefDescendantsAction
	Flags              *string
	ListValues         []niosgrid.ExtensibleattributedefListValues
	Max                *int64
	Min                *int64
	Name               *string
	Type               *string
}
