package internal

import "context"

// RateLimiter controls the rate of API requests.
// Implementations must be safe for concurrent use.
type RateLimiter interface {
	// Wait blocks until the rate limiter allows the request to proceed,
	// or returns an error if the context is cancelled.
	Wait(ctx context.Context) error
}
