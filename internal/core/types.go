package core

import (
	uddiclient "github.com/infobloxopen/bloxone-go-client/client"
	niosclient "github.com/infobloxopen/infoblox-nios-go-client/client"
)

type BackendType string

const (
	BackendNIOS BackendType = "nios"
	BackendUDDI BackendType = "uddi"
)

// Options contains common options for create/update operations
type Options struct {
	ReturnFields string
	Inherit      string
}

// ListOptions contains options for listing records (datasource)
type ListOptions struct {
	Filters         map[string]string
	InternalFilters map[string]string
	ExtAttrFilter   map[string]string
	TagFilter       map[string]string
	// NIOS pagination params
	MaxResults   int32
	Paging       int32
	ReturnFields string
	PageID       string
	// UDDI pagination params
	Limit  int32
	Offset int32
}

type InfobloxClient struct {
	NIOS *niosclient.APIClient
	UDDI *uddiclient.APIClient
}
