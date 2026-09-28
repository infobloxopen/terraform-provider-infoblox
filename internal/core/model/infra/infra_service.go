package infra

import (
	"time"

	uddiinframgmt "github.com/infobloxopen/universal-ddi-go-client/inframgmt"
)

// Infoblox InfraService model
type InfraService struct {
	Id   *string
	UDDI *UDDIInfraServiceExt
}

// UDDIInfraServiceExt - UDDI specific fields for InfraService
type UDDIInfraServiceExt struct {
	Configs         []uddiinframgmt.ServiceHostConfig
	CreatedAt       *time.Time
	Description     *string
	DesiredState    *string
	DesiredVersion  *string
	InterfaceLabels []string
	Name            string
	PoolId          string
	ServiceType     string
	Tags            map[string]any
	UpdatedAt       *time.Time
}
