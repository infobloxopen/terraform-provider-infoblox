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
type Policy struct {
	Retryable RetryableFunc
	Timeout   time.Duration
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
