package client

import (
	"github.com/infobloxopen/bloxone-go-client/anycast"
	"github.com/infobloxopen/bloxone-go-client/clouddiscovery"
	"github.com/infobloxopen/bloxone-go-client/dfp"
	"github.com/infobloxopen/bloxone-go-client/dnsconfig"
	"github.com/infobloxopen/bloxone-go-client/dnsdata"
	"github.com/infobloxopen/bloxone-go-client/fw"
	"github.com/infobloxopen/bloxone-go-client/inframgmt"
	"github.com/infobloxopen/bloxone-go-client/infraprovision"
	"github.com/infobloxopen/bloxone-go-client/ipam"
	"github.com/infobloxopen/bloxone-go-client/ipamfederation"
	"github.com/infobloxopen/bloxone-go-client/keys"
	"github.com/infobloxopen/bloxone-go-client/option"
	"github.com/infobloxopen/bloxone-go-client/redirect"
	"github.com/infobloxopen/bloxone-go-client/upgradepolicy"
)

// APIClient is an aggregation of different BloxOne API clients.
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

// NewAPIClient creates a new BloxOne API Client.
// This is an aggregation of different BloxOne API clients.
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
// - WithCSPUrl(string) sets the URL for BloxOne Cloud Services Portal.
// - WithAPIKey(string) sets the APIKey for accessing the BloxOne API.
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
