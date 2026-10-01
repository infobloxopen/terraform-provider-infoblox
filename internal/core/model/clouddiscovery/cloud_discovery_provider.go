package clouddiscovery

import (
	uddiclouddiscovery "github.com/infobloxopen/universal-ddi-go-client/clouddiscovery"
)

// Infoblox CloudDiscoveryProvider model
type CloudDiscoveryProvider struct {
	Id   *string
	UDDI *UDDICloudDiscoveryProviderExt
}

// UDDICloudDiscoveryProviderExt - UDDI specific fields for CloudDiscoveryProvider
type UDDICloudDiscoveryProviderExt struct {
	AccountPreference       string
	AdditionalConfig        *uddiclouddiscovery.AdditionalConfig
	CredentialPreference    *uddiclouddiscovery.CredentialPreference
	Description             *string
	DesiredState            *string
	DestinationTypesEnabled []string
	Destinations            []uddiclouddiscovery.Destination
	IsDisabled              *bool
	LabsProvider            *bool
	Name                    string
	ProviderType            string
	SourceConfigs           []uddiclouddiscovery.SourceConfig
	SyncInterval            *string
	Tags                    map[string]any
}
