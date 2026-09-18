package fw

import (
	uddifw "github.com/infobloxopen/universal-ddi-go-client/fw"
)

// Infoblox SecurityPolicy model
type SecurityPolicy struct {
	Id   *int32
	UDDI *UDDISecurityPolicyExt
}

// UDDISecurityPolicyExt - UDDI specific fields for SecurityPolicy
type UDDISecurityPolicyExt struct {
	AccessCodes         []string
	DefaultAction       *string
	DefaultRedirectName *string
	Description         *string
	DfpServices         []string
	Dfps                []int32
	Ecs                 *bool
	Name                *string
	NetAddressDfps      []uddifw.NetAddrDfpAssignment
	NetworkLists        []int64
	OnpremResolve       *bool
	Precedence          *int32
	RoamingDeviceGroups []int32
	Rules               []uddifw.SecurityPolicyRule
	SafeSearch          *bool
	Tags                map[string]any
	UserGroups          []string
}
