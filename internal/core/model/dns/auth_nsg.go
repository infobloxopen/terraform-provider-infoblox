package dns

import (
	uddidnsconfig "github.com/infobloxopen/universal-ddi-go-client/dnsconfig"
)

// Infoblox AuthNsg model
type AuthNsg struct {
	Id   *string
	UDDI *UDDIAuthNsgExt
}

// UDDIAuthNsgExt - UDDI specific fields for AuthNsg
type UDDIAuthNsgExt struct {
	Comment             *string
	ExternalPrimaries   []uddidnsconfig.ExternalPrimary
	ExternalSecondaries []uddidnsconfig.ExternalSecondary
	InternalSecondaries []uddidnsconfig.InternalSecondary
	Name                string
	Nsgs                []string
	Tags                map[string]any
}
