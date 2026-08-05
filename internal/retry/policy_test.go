package retry

import (
	"errors"
	"testing"
	"time"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/core"
)

// Core models are named after the object, which is what supplies the policy key.
type (
	RecordA struct{}
	View    struct{}
)

func TestFor(t *testing.T) {
	tests := []struct {
		name          string
		backend       core.BackendType
		op            Operation
		err           error
		wantRetryable bool
		wantTimeout   time.Duration
	}{
		{"uddi create retries missing zone", core.BackendUDDI, OpCreate, errors.New("404 Not Found, 'zone not found'"), true, 2 * time.Minute},
		{"uddi update retries missing record", core.BackendUDDI, OpUpdate, errors.New("404 Not Found, 'record not found'"), true, 2 * time.Minute},
		{"uddi update ignores other not found", core.BackendUDDI, OpUpdate, errors.New("404 Not Found, 'zone not found'"), false, 2 * time.Minute},
		{"uddi read falls back to transient", core.BackendUDDI, OpRead, errors.New("404 Not Found, 'record not found'"), false, 0},
		{"uddi delete falls back to transient", core.BackendUDDI, OpDelete, errors.New("404 Not Found, 'record not found'"), false, 0},
		{"nios create falls back to transient", core.BackendNIOS, OpCreate, errors.New("zone not found"), false, 0},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			p := For[RecordA](tc.backend, tc.op)

			if p.Retryable == nil {
				t.Fatal("expected a non-nil Retryable, got nil")
			}
			if got := p.Retryable(tc.err); got != tc.wantRetryable {
				t.Errorf("Retryable(%v) = %v, want %v", tc.err, got, tc.wantRetryable)
			}
			if p.Timeout != tc.wantTimeout {
				t.Errorf("Timeout = %v, want %v", p.Timeout, tc.wantTimeout)
			}
		})
	}
}

// TestFor_UnlistedObject keeps objects without an override on Transient, which
// also keeps them wired to TransientErrors for when it classifies something.
func TestFor_UnlistedObject(t *testing.T) {
	err := errors.New("404 Not Found, 'zone not found'")

	p := For[View](core.BackendUDDI, OpCreate)
	if p.Retryable == nil {
		t.Fatal("expected fallback policy to carry a predicate")
	}
	if p.Retryable(err) != TransientErrors(err) {
		t.Error("fallback predicate does not match TransientErrors")
	}
	if p.Timeout != 0 {
		t.Errorf("Timeout = %v, want 0", p.Timeout)
	}
}

func TestIsNotFound(t *testing.T) {
	tests := []struct {
		name string
		err  error
		want bool
	}{
		{"nil", nil, false},
		{"zone", errors.New("404 Not Found, 'zone not found'"), true},
		{"record", errors.New("404 Not Found, 'record not found'"), true},
		{"unrelated", errors.New("403 Forbidden"), false},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := IsNotFound(tc.err); got != tc.want {
				t.Errorf("IsNotFound(%v) = %v, want %v", tc.err, got, tc.want)
			}
		})
	}
}

func TestIsRecordNotFound(t *testing.T) {
	tests := []struct {
		name string
		err  error
		want bool
	}{
		{"nil", nil, false},
		{"record", errors.New("404 Not Found, 'record not found'"), true},
		{"zone", errors.New("404 Not Found, 'zone not found'"), false},
		{"unrelated", errors.New("403 Forbidden"), false},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := IsRecordNotFound(tc.err); got != tc.want {
				t.Errorf("IsRecordNotFound(%v) = %v, want %v", tc.err, got, tc.want)
			}
		})
	}
}
