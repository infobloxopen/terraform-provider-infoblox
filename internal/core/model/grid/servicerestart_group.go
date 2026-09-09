package grid

import (
	niosgrid "github.com/infobloxopen/infoblox-nios-go-client/grid"
)

// Infoblox ServicerestartGroup model
type ServicerestartGroup struct {
	Id   *string
	NIOS *NIOSServicerestartGroupExt
}

// NIOSServicerestartGroupExt - NIOS specific fields for ServicerestartGroup
type NIOSServicerestartGroupExt struct {
	Comment           *string
	ExtAttrs          map[string]any
	Members           []string
	Mode              *string
	Name              *string
	RecurringSchedule *niosgrid.GridServicerestartGroupRecurringSchedule
	Service           *string
}
