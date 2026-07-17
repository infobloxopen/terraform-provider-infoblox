package option

import (
	"net/http"

	"github.com/infobloxopen/universal-ddi-go-client/internal"
)

// ClientOption is a function that applies configuration options to the API Client.
type ClientOption func(configuration *internal.Configuration)

// WithCSPUrl returns a ClientOption that sets the URL for Universal DDI Cloud Services Portal.
// Can also be configured using the `INFOBLOX_PORTAL_URL` environment variable.
// Optional. Default is https://csp.infoblox.com
func WithCSPUrl(cspURL string) ClientOption {
	return func(configuration *internal.Configuration) {
		if cspURL != "" {
			configuration.CSPURL = cspURL
		}
	}
}

// WithAPIKey returns a ClientOption that sets the API Key for accessing the Universal DDI API.
// Can also be configured by using the `INFOBLOX_PORTAL_KEY` environment variable.
// You can configure an API key for your user account in the Universal DDI Cloud Services Portal.
// Please refer to the following link for more information: https://docs.infoblox.com/space/BloxOneCloud/35430405/Configuring+User+API+Keys
// Required
func WithAPIKey(apiKey string) ClientOption {
	return func(configuration *internal.Configuration) {
		if apiKey != "" {
			configuration.APIKey = apiKey
		}
	}
}

// WithHTTPClient returns a ClientOption that sets the HTTPClient to use for the SDK.
// Optional. The default HTTPClient will be used if not provided.
func WithHTTPClient(httpClient *http.Client) ClientOption {
	return func(configuration *internal.Configuration) {
		if httpClient != nil {
			configuration.HTTPClient = httpClient
		}
	}
}

// WithDefaultTags returns a ClientOption that sets the tags the client can set by default for objects that has tags support.
// Optional.
func WithDefaultTags(defaultTags map[string]string) ClientOption {
	return func(configuration *internal.Configuration) {
		configuration.DefaultTags = defaultTags
	}
}

// WithClientName returns a ClientOption that sets the name of the client using the SDK.
// This can be used to identify the client in the audit logs.
// Optional. If not provided, the client name will be set to "universal-ddi-go-client".
func WithClientName(clientName string) ClientOption {
	return func(configuration *internal.Configuration) {
		if clientName != "" {
			configuration.ClientName = clientName
		}
	}
}

// WithDebug returns a ClientOption that sets the debug mode.
// Enabling the debug flag will write the request and response to the log.
func WithDebug(debug bool) ClientOption {
	return func(configuration *internal.Configuration) {
		configuration.Debug = debug
	}
}
