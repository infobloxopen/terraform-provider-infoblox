package client

import (
	"github.com/infobloxopen/universal-ddi-go-client/anycast"
	"github.com/infobloxopen/universal-ddi-go-client/clouddiscovery"
	"github.com/infobloxopen/universal-ddi-go-client/dfp"
	"github.com/infobloxopen/universal-ddi-go-client/dnsconfig"
	"github.com/infobloxopen/universal-ddi-go-client/dnsdata"
	"github.com/infobloxopen/universal-ddi-go-client/fw"
	"github.com/infobloxopen/universal-ddi-go-client/inframgmt"
	"github.com/infobloxopen/universal-ddi-go-client/infraprovision"
	"github.com/infobloxopen/universal-ddi-go-client/ipam"
	"github.com/infobloxopen/universal-ddi-go-client/ipamfederation"
	"github.com/infobloxopen/universal-ddi-go-client/keys"
	"github.com/infobloxopen/universal-ddi-go-client/option"
	"github.com/infobloxopen/universal-ddi-go-client/redirect"
	"github.com/infobloxopen/universal-ddi-go-client/upgradepolicy"
)

// APIClient is an aggregation of different Universal DDI API clients.
type APIClient struct {
	IPAddressManagementAPI      *ipam.APIClient
	IPAMFederationAPI           *ipamfederation.APIClient
	DiscoveryConfigurationAPIV2 *clouddiscovery.APIClient
	DNSConfigurationAPI         *dnsconfig.APIClient
	DNSDataAPI                  *dnsdata.APIClient
	HostActivationAPI           *infraprovision.APIClient
	InfraManagementAPI          *inframgmt.APIClient
	KeysAPI                     *keys.APIClient
	DNSForwardingProxyAPI       *dfp.APIClient
	FWAPI                       *fw.APIClient
	AnycastAPI                  *anycast.APIClient
	RedirectAPI                 *redirect.APIClient
	UpgradePolicyClientAPI      *upgradepolicy.APIClient
}

// NewAPIClient creates a new Universal DDI API Client.
// This is an aggregation of different Universal DDI API clients.
// The following clients are available:
// - IPAddressManagementAPI
// - IPAMFederationAPI
// - DiscoveryConfigurationAPIV2
// - DNSConfigurationAPI
// - DNSDataAPI
// - HostActivationAPI
// - InfraManagementAPI
// - KeysAPI
// - DNSForwardingProxyAPI
// - FWAPI
// - AnycastAPI
// - UpgradePolicyClientAPI
//
// The client can be configured with a variadic option. The following options are available:
// - WithClientName(string) sets the name of the client using the SDK.
// - WithCSPUrl(string) sets the URL for Universal DDI Cloud Services Portal.
// - WithAPIKey(string) sets the APIKey for accessing the Universal DDI API.
// - WithHTTPClient(*http.Client) sets the HTTPClient to use for the SDK.
// - WithDefaultTags(map[string]string) sets the tags the client can set by default for objects that has tags support.
// - WithDebug() sets the debug mode.
func NewAPIClient(options ...option.ClientOption) *APIClient {
	return &APIClient{
		IPAddressManagementAPI:      ipam.NewAPIClient(options...),
		IPAMFederationAPI:           ipamfederation.NewAPIClient(options...),
		DiscoveryConfigurationAPIV2: clouddiscovery.NewAPIClient(options...),
		DNSConfigurationAPI:         dnsconfig.NewAPIClient(options...),
		DNSDataAPI:                  dnsdata.NewAPIClient(options...),
		HostActivationAPI:           infraprovision.NewAPIClient(options...),
		InfraManagementAPI:          inframgmt.NewAPIClient(options...),
		KeysAPI:                     keys.NewAPIClient(options...),
		DNSForwardingProxyAPI:       dfp.NewAPIClient(options...),
		FWAPI:                       fw.NewAPIClient(options...),
		AnycastAPI:                  anycast.NewAPIClient(options...),
		RedirectAPI:                 redirect.NewAPIClient(options...),
		UpgradePolicyClientAPI:      upgradepolicy.NewAPIClient(options...),
	}
}
