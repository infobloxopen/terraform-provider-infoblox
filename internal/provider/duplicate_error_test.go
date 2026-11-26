package provider

import (
	"errors"
	"testing"
)

func TestIsDuplicateObjectError(t *testing.T) {
	dup := errors.New(`400 Bad Request, '{ "Error": "Duplicate object 'Terraform Internal ID' of type 'extensible_attributes_def' already exists in the database." }'`)
	if !isDuplicateObjectError(dup) {
		t.Fatal("expected duplicate object error to match")
	}
	if isDuplicateObjectError(errors.New("401 Unauthorized")) || isDuplicateObjectError(nil) {
		t.Fatal("unexpected match")
	}
}
