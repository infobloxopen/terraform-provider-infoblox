package client

import (
	"github.com/infobloxopen/infoblox-nios-go-client/acl"
	"github.com/infobloxopen/infoblox-nios-go-client/cloud"
	"github.com/infobloxopen/infoblox-nios-go-client/dhcp"
	"github.com/infobloxopen/infoblox-nios-go-client/discovery"
	"github.com/infobloxopen/infoblox-nios-go-client/dns"
	"github.com/infobloxopen/infoblox-nios-go-client/dtc"
	"github.com/infobloxopen/infoblox-nios-go-client/federatedrealms"
	"github.com/infobloxopen/infoblox-nios-go-client/grid"
	"github.com/infobloxopen/infoblox-nios-go-client/ipam"
	"github.com/infobloxopen/infoblox-nios-go-client/microsoftserver"
	"github.com/infobloxopen/infoblox-nios-go-client/misc"
	"github.com/infobloxopen/infoblox-nios-go-client/notification"
	"github.com/infobloxopen/infoblox-nios-go-client/option"
	"github.com/infobloxopen/infoblox-nios-go-client/parentalcontrol"
	"github.com/infobloxopen/infoblox-nios-go-client/rir"
	"github.com/infobloxopen/infoblox-nios-go-client/rpz"
	"github.com/infobloxopen/infoblox-nios-go-client/security"
	"github.com/infobloxopen/infoblox-nios-go-client/smartfolder"
	"github.com/infobloxopen/infoblox-nios-go-client/threatinsight"
	"github.com/infobloxopen/infoblox-nios-go-client/threatprotection"
)

// APIClient is an aggregation of different NIOS WAPI clients.
type APIClient struct {
	ACLAPI              *acl.APIClient
	CloudAPI            *cloud.APIClient
	DHCPAPI             *dhcp.APIClient
	DiscoveryAPI        *discovery.APIClient
	DNSAPI              *dns.APIClient
	DTCAPI              *dtc.APIClient
	FederatedRealmsAPI  *federatedrealms.APIClient
	GridAPI             *grid.APIClient
	IPAMAPI             *ipam.APIClient
	MicrosoftServerAPI  *microsoftserver.APIClient
	MiscAPI             *misc.APIClient
	NotificationAPI     *notification.APIClient
	ParentalControlAPI  *parentalcontrol.APIClient
	RIRAPI              *rir.APIClient
	RPZAPI              *rpz.APIClient
	SecurityAPI         *security.APIClient
	SmartFolderAPI      *smartfolder.APIClient
	ThreatInsightAPI    *threatinsight.APIClient
	ThreatProtectionAPI *threatprotection.APIClient
}

// NewAPIClient creates a new NIOS WAPI Client.
// This is an aggregation of different NIOS WAPI clients.
// The following clients are available:
// - ACLAPI
// - CloudAPI
// - DHCPAPI
// - DiscoveryAPI
// - DNSAPI
// - DTCAPI
// - FederatedRealmsAPI
// - GridAPI
// - IPAMAPI
// - MicrosoftServerAPI
// - MiscAPI
// - NotificationAPI
// - ParentalControlAPI
// - RIRAPI
// - RPZAPI
// - SecurityAPI
// - SmartFolderAPI
// - ThreatInsightAPI
// - ThreatProtectionAPI
// The client can be configured with a variadic option. The following options are available:
// - WithClientName(string) sets the name of the client using the SDK.
// - WithNIOSHostUrl(string) sets the URL for NIOS Portal.
// - WithNIOSUsername(string) sets the Username for the NIOS Portal.
// - WithNIOSPassword(string) sets the Password for the NIOS Portal.
// - WithHTTPClient(*http.Client) sets the HTTPClient to use for the SDK.
// - WithDefaultExtAttrs(map[string]struct{ Value String }) sets the Extensible Attributes the client can set by default for objects that has Extensible Attributes support.
// - WithDebug() sets the debug mode.
func NewAPIClient(options ...option.ClientOption) *APIClient {
	return &APIClient{
		ACLAPI:              acl.NewAPIClient(options...),
		CloudAPI:            cloud.NewAPIClient(options...),
		DHCPAPI:             dhcp.NewAPIClient(options...),
		DiscoveryAPI:        discovery.NewAPIClient(options...),
		DNSAPI:              dns.NewAPIClient(options...),
		DTCAPI:              dtc.NewAPIClient(options...),
		FederatedRealmsAPI:  federatedrealms.NewAPIClient(options...),
		GridAPI:             grid.NewAPIClient(options...),
		IPAMAPI:             ipam.NewAPIClient(options...),
		MicrosoftServerAPI:  microsoftserver.NewAPIClient(options...),
		MiscAPI:             misc.NewAPIClient(options...),
		NotificationAPI:     notification.NewAPIClient(options...),
		ParentalControlAPI:  parentalcontrol.NewAPIClient(options...),
		RIRAPI:              rir.NewAPIClient(options...),
		RPZAPI:              rpz.NewAPIClient(options...),
		SecurityAPI:         security.NewAPIClient(options...),
		SmartFolderAPI:      smartfolder.NewAPIClient(options...),
		ThreatInsightAPI:    threatinsight.NewAPIClient(options...),
		ThreatProtectionAPI: threatprotection.NewAPIClient(options...),
	}
}
