package infra

import (
	uddiinframgmt "github.com/infobloxopen/universal-ddi-go-client/inframgmt"
)

// Infoblox InfraHost model
type InfraHost struct {
	Id   *string
	UDDI *UDDIInfraHostExt
}

// UDDIInfraHostExt - UDDI specific fields for InfraHost
type UDDIInfraHostExt struct {
	Configs         []uddiinframgmt.ServiceHostConfig
	Description     *string
	DisplayName     string
	HostType        *string
	IpAddress       *string
	IpSpace         *string
	LegacyId        *string
	LocationId      *string
	MacAddress      *string
	MaintenanceMode *string
	Ophid           *string
	PoolId          *string
	SerialNumber    *string
	Tags            map[string]any
	Timezone        *string
}
