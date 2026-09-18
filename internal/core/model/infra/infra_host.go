package infra

// Infoblox InfraHost model
type InfraHost struct {
	Id   *string
	UDDI *UDDIInfraHostExt
}

// UDDIInfraHostExt - UDDI specific fields for InfraHost
type UDDIInfraHostExt struct {
	Description     *string
	DisplayName     string
	IpSpace         *string
	LocationId      *string
	MaintenanceMode *string
	PoolId          *string
	SerialNumber    *string
	Tags            map[string]any
}
