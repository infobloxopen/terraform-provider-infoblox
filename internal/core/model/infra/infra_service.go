package infra

// Infoblox InfraService model
type InfraService struct {
	Id   *string
	UDDI *UDDIInfraServiceExt
}

// UDDIInfraServiceExt - UDDI specific fields for InfraService
type UDDIInfraServiceExt struct {
	Description     *string
	DesiredState    *string
	DesiredVersion  *string
	InterfaceLabels []string
	Name            string
	PoolId          string
	ServiceType     string
	Tags            map[string]any
}
