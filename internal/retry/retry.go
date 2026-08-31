package retry

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-log/tflog"
)

const (
	initialBackoff      = 1 * time.Second
	maxBackoff          = 30 * time.Second
	operationTimeoutMsg = "operation timeout exceeded while waiting for the request to complete, the failure may be due to a transient issue or request cancellation"
)

// OperationTimeout is the default time for an operation, covering its
// first attempt and every retry. It can be overridden at the provider config
// level via SetOperationTimeout, or per call through Policy.Timeout.
var OperationTimeout = 60 * time.Second

type (
	// RetryableFunc reports whether an error is worth retrying.
	RetryableFunc func(error) bool

	// RetryFunc performs a retryable operation. It returns the HTTP status code of
	// the attempt (0 when no response was received) and the resulting error.
	RetryFunc func(ctx context.Context) (int, error)
)

// SetOperationTimeout sets the global operation timeout, in seconds.
func SetOperationTimeout(timeout int64) {
	if timeout <= 0 {
		return
	}
	OperationTimeout = time.Duration(timeout) * time.Second
}

// Do runs fn until:
// - fn succeeds
// - error is non-retryable
// - context is canceled or times out
func Do(parentCtx context.Context, p Policy, fn RetryFunc) error {
	timeout := p.Timeout
	if timeout <= 0 {
		timeout = OperationTimeout
	}

	ctx, cancel := context.WithTimeout(parentCtx, timeout)
	defer cancel()

	backoff := initialBackoff
	attempt := 0

	for {
		attempt++
		_, err := fn(ctx)
		if err == nil {
			return nil
		}

		// Stop retrying on context deadline exceeded, cancellation
		if errors.Is(err, context.DeadlineExceeded) || errors.Is(err, context.Canceled) ||
			errors.Is(ctx.Err(), context.DeadlineExceeded) || errors.Is(ctx.Err(), context.Canceled) {
			// Overriding the ctx deadline/cancellation error message for better user understanding
			return errors.New(operationTimeoutMsg)
		}

		// Stop retrying if error is not retryable
		if p.Retryable == nil || !p.Retryable(err) {
			return err
		}

		tflog.Warn(ctx, fmt.Sprintf(
			"Transient error detected, retrying request (attempt=%d, backoff=%s, err=%v)",
			attempt, backoff, err,
		))

		// Wait before retrying with exponential backoff
		select {
		case <-ctx.Done():
			return fmt.Errorf("%s: %w", operationTimeoutMsg, err)
		case <-time.After(backoff):
		}

		backoff *= 2
		if backoff > maxBackoff {
			backoff = maxBackoff
		}
	}
}

// TransientErrors determines if an error is retryable based on transient conditions.
// TODO: Currently returns false, treating all errors as non-retryable. Extend
// this with predicates for network errors, temporary service unavailability (5xx errors), etc.
func TransientErrors(_ error) bool {
	return false
}
