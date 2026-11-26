package dns

import (
	"testing"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/core"
)

func TestDefaultStaticCreatorFilter(t *testing.T) {
	if got := defaultStaticCreatorFilter(core.BackendNIOS, nil); got["creator"] != "STATIC" {
		t.Fatalf("expected creator=STATIC by default, got %v", got)
	}
	if got := defaultStaticCreatorFilter(core.BackendNIOS, map[string]string{"creator": "SYSTEM"}); got["creator"] != "SYSTEM" {
		t.Fatalf("user creator filter must win, got %v", got)
	}
	if got := defaultStaticCreatorFilter(core.BackendNIOS, map[string]string{"nios.creator": "DYNAMIC"}); len(got) != 1 {
		t.Fatalf("nios.creator must suppress the default, got %v", got)
	}
	if got := defaultStaticCreatorFilter(core.BackendUDDI, nil); got != nil {
		t.Fatalf("UDDI filters must be unchanged, got %v", got)
	}
}
