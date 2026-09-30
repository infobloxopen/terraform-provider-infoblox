package option

import (
	"context"
	"net/http"
	"time"

	"github.com/infobloxopen/universal-ddi-go-client/internal"
	"golang.org/x/time/rate"
)

// ClientOption is a function that applies configuration options to the API Client.
type ClientOption func(configuration *internal.Configuration)

// RateLimiter is the interface for custom rate limiter implementations.
// Implementations must be safe for concurrent use.
type RateLimiter interface {
	Wait(ctx context.Context) error
}

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

// WithRateLimit returns a ClientOption that configures a token-bucket rate limiter.
// requestsPerSecond controls the sustained request rate. burst controls the maximum
// number of requests that can be made instantly before rate limiting kicks in.
// The limiter instance is created once and shared across all services when used
// with the aggregated client.NewAPIClient.
// Can also be configured using the INFOBLOX_RATE_LIMIT and
// INFOBLOX_RATE_LIMIT_BURST environment variables.
func WithRateLimit(requestsPerSecond float64, burst int) ClientOption {
	if requestsPerSecond <= 0 {
		return WithRateLimitDisabled()
	}
	if burst < 1 {
		burst = 1
	}
	limiter := rate.NewLimiter(rate.Limit(requestsPerSecond), burst)
	return func(configuration *internal.Configuration) {
		configuration.RateLimiter = limiter
	}
}

// WithRateLimitDisabled returns a ClientOption that disables the built-in rate limiter.
func WithRateLimitDisabled() ClientOption {
	return func(configuration *internal.Configuration) {
		configuration.RateLimiter = nil
	}
}

// WithRateLimiter returns a ClientOption that sets a custom rate limiter implementation.
// Use this for advanced use cases like distributed rate limiting or
// per-resource-type limiting at the provider level.
func WithRateLimiter(limiter RateLimiter) ClientOption {
	return func(configuration *internal.Configuration) {
		if limiter != nil {
			configuration.RateLimiter = limiter
		}
	}
}

// WithRetry returns a ClientOption that configures automatic retry with exponential
// backoff for transient HTTP errors (429, 500, 502, 503, 504).
// maxRetries is the number of retries after the initial request.
// minWait and maxWait control the exponential backoff range.
// Can also be configured using the INFOBLOX_MAX_RETRIES,
// INFOBLOX_RETRY_MIN_WAIT, and INFOBLOX_RETRY_MAX_WAIT environment variables.
func WithRetry(maxRetries int, minWait, maxWait time.Duration) ClientOption {
	return func(configuration *internal.Configuration) {
		if maxRetries > 0 {
			configuration.RetryConfig = &internal.RetryConfig{
				MaxRetries: maxRetries,
				MinWait:    minWait,
				MaxWait:    maxWait,
			}
		} else {
			configuration.RetryConfig = nil
		}
	}
}

// WithRetryDisabled returns a ClientOption that disables automatic retry.
func WithRetryDisabled() ClientOption {
	return func(configuration *internal.Configuration) {
		configuration.RetryConfig = nil
	}
}
