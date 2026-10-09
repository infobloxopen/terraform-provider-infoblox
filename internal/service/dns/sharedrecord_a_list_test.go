package dns_test

import (
	"testing"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/acctest"
)

func TestAccSharedrecordAList(t *testing.T) {
	resourceType := "infoblox_sharedrecord_a"

	checksByBackend := map[string]acctest.CheckFuncs{
		"nios": {
			Exists:  testAccCheckSharedrecordAExistsNIOS,
			Destroy: testAccCheckSharedrecordADestroyNIOS,
		},
	}

	for _, backend := range []string{"nios"} {
		t.Run(backend, func(t *testing.T) {
			acctest.RunListCases(t, resourceType, "dns/sharedrecord_a/"+backend+"_lists.hcl", checksByBackend)
		})
	}
}
