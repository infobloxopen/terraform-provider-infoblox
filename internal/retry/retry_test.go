package retry

import (
	"context"
	"errors"
	"testing"
	"time"
)

// alwaysRetry is a policy that retries every error.
var alwaysRetry = Policy{Retryable: func(error) bool { return true }}

// TestDo_Success tests that Do returns nil when the function succeeds
func TestDo_Success(t *testing.T) {
	callCount := 0
	fn := func(ctx context.Context) (int, error) {
		callCount++
		return 200, nil
	}

	err := Do(context.Background(), Policy{}, fn)
	if err != nil {
		t.Errorf("Expected no error, got: %v", err)
	}
	if callCount != 1 {
		t.Errorf("Expected 1 call, got: %d", callCount)
	}
}

// TestDo_NonRetryableError tests that Do returns immediately on non-retryable errors
func TestDo_NonRetryableError(t *testing.T) {
	callCount := 0
	expectedErr := errors.New("non-retryable error")

	fn := func(ctx context.Context) (int, error) {
		callCount++
		return 500, expectedErr
	}

	p := Policy{Retryable: func(err error) bool { return false }}

	err := Do(context.Background(), p, fn)
	if err == nil {
		t.Error("Expected error, got nil")
	}
	if !errors.Is(err, expectedErr) {
		t.Errorf("Expected error %v, got: %v", expectedErr, err)
	}
	if callCount != 1 {
		t.Errorf("Expected 1 call for non-retryable error, got: %d", callCount)
	}
}

// TestDo_RetryableErrorWithEventualSuccess tests retrying until success
func TestDo_RetryableErrorWithEventualSuccess(t *testing.T) {
	// Reduce timeout for faster test execution
	originalTimeout := OperationTimeout
	SetOperationTimeout(5)
	defer func() { OperationTimeout = originalTimeout }()

	callCount := 0
	fn := func(ctx context.Context) (int, error) {
		callCount++
		if callCount < 3 {
			return 500, errors.New("temporary error")
		}
		return 200, nil
	}

	err := Do(context.Background(), alwaysRetry, fn)
	if err != nil {
		t.Errorf("Expected no error after retries, got: %v", err)
	}
	if callCount != 3 {
		t.Errorf("Expected 3 calls, got: %d", callCount)
	}
}

// TestDo_ContextTimeout tests that Do respects the global timeout
func TestDo_ContextTimeout(t *testing.T) {
	// Set a very short timeout
	originalTimeout := OperationTimeout
	SetOperationTimeout(1)
	defer func() { OperationTimeout = originalTimeout }()

	callCount := 0
	fn := func(ctx context.Context) (int, error) {
		callCount++
		// Block until the retry context times out or is cancelled
		<-ctx.Done()
		return 500, ctx.Err()
	}

	err := Do(context.Background(), alwaysRetry, fn)
	if err == nil {
		t.Error("Expected timeout error, got nil")
	}
	if err.Error() != operationTimeoutMsg {
		t.Errorf("Expected timeout message, got: %v", err)
	}
	if callCount != 1 {
		t.Errorf("Expected 1 call, got: %d", callCount)
	}
}

// TestDo_ContextCancellation tests that Do respects context cancellation
func TestDo_ContextCancellation(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())

	callCount := 0
	fn := func(ctx context.Context) (int, error) {
		callCount++
		if callCount == 2 {
			cancel()
		}
		return 500, errors.New("retryable error")
	}

	err := Do(ctx, alwaysRetry, fn)
	if err == nil {
		t.Errorf("expected cancellation error, got nil")
	}
	if err.Error() != operationTimeoutMsg {
		t.Errorf("expected retry timeout message, got: %v", err)
	}
	if callCount < 2 {
		t.Errorf("expected at least 2 calls, got: %d", callCount)
	}
}

// TestDo_NilRetryableFunc tests that a policy without a predicate runs fn once
func TestDo_NilRetryableFunc(t *testing.T) {
	callCount := 0
	expectedErr := errors.New("some error")

	fn := func(ctx context.Context) (int, error) {
		callCount++
		return 500, expectedErr
	}

	err := Do(context.Background(), Policy{}, fn)
	if err == nil {
		t.Error("Expected error, got nil")
	}
	if !errors.Is(err, expectedErr) {
		t.Errorf("Expected error %v, got: %v", expectedErr, err)
	}
	if callCount != 1 {
		t.Errorf("Expected 1 call when Retryable is nil, got: %d", callCount)
	}
}

// TestDo_PolicyTimeoutOverridesGlobal tests that Policy.Timeout wins over OperationTimeout
func TestDo_PolicyTimeoutOverridesGlobal(t *testing.T) {
	originalTimeout := OperationTimeout
	SetOperationTimeout(3600)
	defer func() { OperationTimeout = originalTimeout }()

	fn := func(ctx context.Context) (int, error) {
		<-ctx.Done()
		return 500, ctx.Err()
	}

	start := time.Now()
	err := Do(context.Background(), Policy{Retryable: func(error) bool { return true }, Timeout: 500 * time.Millisecond}, fn)
	elapsed := time.Since(start)

	if err == nil {
		t.Fatal("Expected timeout error, got nil")
	}
	if err.Error() != operationTimeoutMsg {
		t.Errorf("Expected timeout message, got: %v", err)
	}
	if elapsed > 5*time.Second {
		t.Errorf("Policy timeout was not applied, call took %s", elapsed)
	}
}

// TestSetOperationTimeout tests the SetOperationTimeout function
func TestSetOperationTimeout(t *testing.T) {
	originalTimeout := OperationTimeout
	defer func() { OperationTimeout = originalTimeout }()

	SetOperationTimeout(10)
	if OperationTimeout != 10*time.Second {
		t.Fatalf("expected 10s, got %v", OperationTimeout)
	}

	OperationTimeout = 5 * time.Second
	SetOperationTimeout(0)
	if OperationTimeout != 5*time.Second {
		t.Fatalf("zero timeout should not change value, got %v", OperationTimeout)
	}

	SetOperationTimeout(-1)
	if OperationTimeout != 5*time.Second {
		t.Fatalf("negative timeout should not change value")
	}
}

// TestTransientErrors documents that every error is non-retryable today
func TestTransientErrors(t *testing.T) {
	if TransientErrors(nil) {
		t.Error("expected nil error to be non-retryable")
	}
	if TransientErrors(errors.New("dial tcp: connection refused")) {
		t.Error("expected all errors to be non-retryable for now")
	}
}
