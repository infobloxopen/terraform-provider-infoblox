package grid

import (
	niosgrid "github.com/infobloxopen/infoblox-nios-go-client/grid"
)

// Infoblox Distributionschedule model
type Distributionschedule struct {
	Id   *string
	NIOS *NIOSDistributionscheduleExt
}

// NIOSDistributionscheduleExt - NIOS specific fields for Distributionschedule
type NIOSDistributionscheduleExt struct {
	Active        *bool
	StartTime     *int64
	TimeZone      *string
	UpgradeGroups []niosgrid.DistributionscheduleUpgradeGroups
}
