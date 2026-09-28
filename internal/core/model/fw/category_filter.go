package fw

import (
	"time"
)

// Infoblox CategoryFilter model
type CategoryFilter struct {
	Id   *int32
	UDDI *UDDICategoryFilterExt
}

// UDDICategoryFilterExt - UDDI specific fields for CategoryFilter
type UDDICategoryFilterExt struct {
	Categories  []string
	CreatedTime *time.Time
	Description *string
	Name        *string
	Policies    []string
	Tags        map[string]any
	UpdatedTime *time.Time
}
