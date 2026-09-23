package dns_test

import (
	"testing"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/acctest"
)

func TestAccSharedrecordMxList(t *testing.T) {
	resourceType := "infoblox_sharedrecord_mx"

	checksByBackend := map[string]acctest.CheckFuncs{
		"nios": {
			Exists:  testAccCheckSharedrecordMxExistsNIOS,
			Destroy: testAccCheckSharedrecordMxDestroyNIOS,
		},
	}

	for _, backend := range []string{"nios"} {
		t.Run(backend, func(t *testing.T) {
			acctest.RunListCases(t, resourceType, "dns/sharedrecord_mx/"+backend+"_lists.hcl", checksByBackend)
		})
	}
}
