package option

import (
	"net/http"
	"net/url"
	"strings"

	"github.com/infobloxopen/infoblox-nios-go-client/internal"
)

// ClientOption is a function that applies configuration options to the API Client.
type ClientOption func(configuration *internal.Configuration)

// WithNIOSHostUrl returns a ClientOption that sets the URL for Infoblox NIOS Portal
// Can also be configured using the `NIOS_HOST_URL` environment variable.
// Required
func WithNIOSHostUrl(NIOSHostURL string) ClientOption {
	return func(configuration *internal.Configuration) {
		if NIOSHostURL != "" {
			configuration.NIOSHostURL = NIOSHostURL
		}
	}
}

// WithNIOSUsername returns a ClientOption that sets the NIOSUsername for Infoblox NIOS Portal
// Can also be configured using the `NIOS_USERNAME` environment variable.
func WithNIOSUsername(NIOSUsername string) ClientOption {
	return func(configuration *internal.Configuration) {
		if NIOSUsername != "" {
			configuration.NIOSUsername = NIOSUsername
		}
	}
}

// WithNIOSPassword returns a ClientOption that sets the NIOSPassword for Infoblox NIOS Portal.
// Can also be configured using the `NIOS_PASSWORD` environment variable.
func WithNIOSPassword(NIOSPassword string) ClientOption {
	return func(configuration *internal.Configuration) {
		if NIOSPassword != "" {
			configuration.NIOSPassword = NIOSPassword
		}
	}
}

// WithNIOSPassthrough returns a ClientOption that reaches NIOS through the Infoblox Portal.
func WithNIOSPassthrough(enabled bool) ClientOption {
	return func(configuration *internal.Configuration) {
		configuration.NIOSPassthrough = enabled
	}
}

// WithNIOSLicenseUID returns a ClientOption that sets the license UID of the NIOS Grid to reach
// through the Infoblox Portal. Can also be configured using the `NIOS_LICENSE_UID` environment variable.
func WithNIOSLicenseUID(NIOSLicenseUID string) ClientOption {
	return func(configuration *internal.Configuration) {
		if NIOSLicenseUID != "" {
			configuration.NIOSLicenseUID = NIOSLicenseUID
		}
	}
}

// WithPortalUrl returns a ClientOption that sets the Infoblox Portal WAPI endpoint used in
// passthrough mode. Can also be configured using the `INFOBLOX_PORTAL_URL` environment variable.
func WithPortalUrl(portalURL string) ClientOption {
	return func(configuration *internal.Configuration) {
		if portalURL != "" {
			configuration.PortalURL = portalURL
		}
	}
}

// WithPortalAPIKey returns a ClientOption that sets the Infoblox Portal API key authenticating
// passthrough requests. Can also be configured using the `INFOBLOX_PORTAL_KEY` environment variable.
func WithPortalAPIKey(portalAPIKey string) ClientOption {
	return func(configuration *internal.Configuration) {
		if portalAPIKey != "" {
			configuration.PortalAPIKey = portalAPIKey
		}
	}
}

// ValidatePassthrough reports whether the options describe a usable passthrough setup.
func ValidatePassthrough(options ...ClientOption) error {
	configuration := internal.NewConfiguration()
	for _, opt := range options {
		opt(configuration)
	}
	return configuration.CheckPortalConfig()
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

// WithDefaultExtAttrs returns a ClientOption that sets the Extensible Attributes the client can set by default for objects that has Extensible Attributes support.
// Optional.
func WithDefaultExtAttrs(defaultExtAttrs map[string]struct{ Value string }) ClientOption {
	return func(configuration *internal.Configuration) {
		configuration.DefaultExtAttrs = defaultExtAttrs
	}
}

// WithClientName returns a ClientOption that sets the name of the client using the SDK.
// This can be used to identify the client in the audit logs.
// Optional. If not provided, the client name will be set to "nios-go-client".
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

// WithProxyURL returns a ClientOption that sets the URL for Proxy Server
func WithProxyURL(proxyURL string) ClientOption {
	return func(configuration *internal.Configuration) {
		if strings.TrimSpace(proxyURL) != "" {
			parsedURL, err := url.Parse(strings.TrimSpace(proxyURL))
			if err != nil {
				configuration.ProxyURL = nil
			} else {
				configuration.ProxyURL = parsedURL
			}
		}
	}
}
