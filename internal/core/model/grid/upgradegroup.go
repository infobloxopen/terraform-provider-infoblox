package grid

import (
	niosgrid "github.com/infobloxopen/infoblox-nios-go-client/grid"
)

// Infoblox Upgradegroup model
type Upgradegroup struct {
	Id   *string
	NIOS *NIOSUpgradegroupExt
}

// NIOSUpgradegroupExt - NIOS specific fields for Upgradegroup
type NIOSUpgradegroupExt struct {
	Comment                    *string
	DistributionDependentGroup *string
	DistributionPolicy         *string
	DistributionTime           *int64
	Members                    []niosgrid.UpgradegroupMembers
	Name                       *string
	UpgradeDependentGroup      *string
	UpgradePolicy              *string
	UpgradeTime                *int64
}
