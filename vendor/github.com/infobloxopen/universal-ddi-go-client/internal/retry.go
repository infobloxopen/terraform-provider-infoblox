package internal

import (
	"crypto/x509"
	"errors"
	"math"
	"math/rand"
	"net"
	"net/http"
	"strconv"
	"time"
)

// RetryConfig controls automatic retry behavior for transient HTTP errors.
type RetryConfig struct {
	MaxRetries int
	MinWait    time.Duration
	MaxWait    time.Duration
}

func isRetryableStatusCode(statusCode int) bool {
	switch statusCode {
	case http.StatusTooManyRequests,
		http.StatusInternalServerError,
		http.StatusBadGateway,
		http.StatusServiceUnavailable,
		http.StatusGatewayTimeout:
		return true
	}
	return false
}

func parseRetryAfter(resp *http.Response) time.Duration {
	header := resp.Header.Get("Retry-After")
	if header == "" {
		return 0
	}

	if seconds, err := strconv.Atoi(header); err == nil && seconds > 0 {
		return time.Duration(seconds) * time.Second
	}

	if t, err := http.ParseTime(header); err == nil {
		delay := time.Until(t)
		if delay > 0 {
			return delay
		}
	}

	return 0
}

func retryBackoff(attempt int, minWait, maxWait time.Duration) time.Duration {
	cap := float64(minWait) * math.Pow(2, float64(attempt))
	if cap > float64(maxWait) {
		cap = float64(maxWait)
	}
	return time.Duration(rand.Float64() * cap)
}

// isRetryableError reports whether err represents a transient condition
// worth retrying. Permanent transport errors (DNS not found, invalid or
// unknown-authority TLS certificates) are excluded from retry.
func isRetryableError(err error) bool {
	if err == nil {
		return false
	}
	var dnsErr *net.DNSError
	if errors.As(err, &dnsErr) && dnsErr.IsNotFound {
		return false
	}
	var certInvalidErr *x509.CertificateInvalidError
	if errors.As(err, &certInvalidErr) {
		return false
	}
	var unknownAuthErr *x509.UnknownAuthorityError
	if errors.As(err, &unknownAuthErr) {
		return false
	}
	return true
}
