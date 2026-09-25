package retry

import (
	"reflect"
	"strings"
	"time"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/core"
)

type Operation string

const (
	OpCreate Operation = "create"
	OpRead   Operation = "read"
	OpUpdate Operation = "update"
	OpDelete Operation = "delete"
)

// Policy is the retry behaviour of a single API call. A nil Retryable attempts
// the call exactly once, a 0 Timeout falls back to OperationTimeout.
// IgnoreError, when set, suppresses a final-attempt error: if the predicate
// returns true, Do returns nil (success). Use for delete operations where the
// API returns a non-404 status to signal "already deleted" (e.g. the redirect
// API returns 400 "Non-existent" instead of 404).
type Policy struct {
	Retryable   RetryableFunc
	Timeout     time.Duration
	IgnoreError func(error) bool
}

// Transient is the default policy for every call without an override.
func Transient() Policy {
	return Policy{Retryable: TransientErrors}
}

type override struct {
	object  string
	backend core.BackendType
	op      Operation
}

// overrides contains the calls needing something other than Transient. The same
// operation on the same object can differ per backend.
var overrides = map[override]Policy{
	// UDDI calls can fail with "not found" if the zone of a record is not yet created, hence we retry for little longer.
	{"RecordA", core.BackendUDDI, OpCreate}: {Retryable: IsNotFound, Timeout: 2 * time.Minute},
	{"RecordA", core.BackendUDDI, OpUpdate}: {Retryable: IsRecordNotFound, Timeout: 2 * time.Minute},

	{"RecordAaaa", core.BackendUDDI, OpCreate}: {Retryable: IsNotFound, Timeout: 2 * time.Minute},
	{"RecordAaaa", core.BackendUDDI, OpUpdate}: {Retryable: IsRecordNotFound, Timeout: 2 * time.Minute},

	{"RecordCaa", core.BackendUDDI, OpCreate}: {Retryable: IsNotFound, Timeout: 2 * time.Minute},
	{"RecordCaa", core.BackendUDDI, OpUpdate}: {Retryable: IsRecordNotFound, Timeout: 2 * time.Minute},

	{"RecordCname", core.BackendUDDI, OpCreate}: {Retryable: IsNotFound, Timeout: 2 * time.Minute},
	{"RecordCname", core.BackendUDDI, OpUpdate}: {Retryable: IsRecordNotFound, Timeout: 2 * time.Minute},

	{"RecordDname", core.BackendUDDI, OpCreate}: {Retryable: IsNotFound, Timeout: 2 * time.Minute},
	{"RecordDname", core.BackendUDDI, OpUpdate}: {Retryable: IsRecordNotFound, Timeout: 2 * time.Minute},

	{"RecordNaptr", core.BackendUDDI, OpCreate}: {Retryable: IsNotFound, Timeout: 2 * time.Minute},
	{"RecordNaptr", core.BackendUDDI, OpUpdate}: {Retryable: IsRecordNotFound, Timeout: 2 * time.Minute},

	{"RecordSrv", core.BackendUDDI, OpCreate}: {Retryable: IsNotFound, Timeout: 2 * time.Minute},
	{"RecordSrv", core.BackendUDDI, OpUpdate}: {Retryable: IsRecordNotFound, Timeout: 2 * time.Minute},

	{"RecordTxt", core.BackendUDDI, OpCreate}: {Retryable: IsNotFound, Timeout: 2 * time.Minute},
	{"RecordTxt", core.BackendUDDI, OpUpdate}: {Retryable: IsRecordNotFound, Timeout: 2 * time.Minute},

	{"RecordNs", core.BackendUDDI, OpCreate}: {Retryable: IsNotFound, Timeout: 2 * time.Minute},
	{"RecordNs", core.BackendUDDI, OpUpdate}: {Retryable: IsRecordNotFound, Timeout: 2 * time.Minute},

	{"ZoneAuth", core.BackendUDDI, OpDelete}: {Retryable: IsZoneReferenced, Timeout: 2 * time.Minute},
	{"ZoneAuth", core.BackendUDDI, OpDelete}: {Retryable: IsZoneReferenced, Timeout: 2 * time.Minute},

	{"RecordPtr", core.BackendUDDI, OpCreate}: {Retryable: IsNotFound, Timeout: 2 * time.Minute},
	{"RecordPtr", core.BackendUDDI, OpUpdate}: {Retryable: IsRecordNotFound, Timeout: 2 * time.Minute},

	{"Networkview", core.BackendUDDI, OpDelete}: {Retryable: IsNetworkViewReferenced, Timeout: 2 * time.Minute},

	{"TsigKey", core.BackendUDDI, OpDelete}: {Retryable: IsTsigReferenced, Timeout: 2 * time.Minute},

	// The redirect API returns 400 "Non-existent: <id>" (not 404) when deleting an already-deleted
	// resource. IgnoreError makes Do return nil for that case, which the generated delete function
	// interprets as success (no error → resource is gone, which is the desired outcome).
	{"CustomRedirect", core.BackendUDDI, OpDelete}: {IgnoreError: IsNonExistent},
}

// For resolves the policy for op on backend. T is the core model of the object,
// whose type name supplies the object key: coremodel.RecordA -> RecordA.
func For[T any](backend core.BackendType, op Operation) Policy {
	if p, ok := overrides[override{reflect.TypeFor[T]().Name(), backend, op}]; ok {
		return p
	}
	return Transient()
}

func IsNotFound(err error) bool {
	return err != nil && strings.Contains(err.Error(), "not found")
}

func IsRecordNotFound(err error) bool {
	return err != nil && strings.Contains(err.Error(), "record not found")
}

func IsZoneReferenced(err error) bool {
	return err != nil && strings.Contains(err.Error(), "object is referenced by a 'Zone' object")
}

func IsNetworkViewReferenced(err error) bool {
	return err != nil && strings.Contains(err.Error(), "Cannot delete IP Space") &&
		strings.Contains(err.Error(), "it is being used")
}

func IsTsigReferenced(err error) bool {
	return err != nil && strings.Contains(err.Error(), "object is referenced by a 'TSIG Key' object")
}

func IsNonExistent(err error) bool {
	return err != nil && strings.Contains(err.Error(), "Non-existent")
}
