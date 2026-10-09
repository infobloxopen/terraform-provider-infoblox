package dns_test

import (
	"testing"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/acctest"
)

func TestAccSharedrecordAaaaList(t *testing.T) {
	resourceType := "infoblox_sharedrecord_aaaa"

	checksByBackend := map[string]acctest.CheckFuncs{
		"nios": {
			Exists:  testAccCheckSharedrecordAaaaExistsNIOS,
			Destroy: testAccCheckSharedrecordAaaaDestroyNIOS,
		},
	}

	for _, backend := range []string{"nios"} {
		t.Run(backend, func(t *testing.T) {
			acctest.RunListCases(t, resourceType, "dns/sharedrecord_aaaa/"+backend+"_lists.hcl", checksByBackend)
		})
	}
}
