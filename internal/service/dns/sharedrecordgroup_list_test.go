package dns_test

import (
	"testing"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/acctest"
)

func TestAccSharedrecordgroupList(t *testing.T) {
	resourceType := "infoblox_sharedrecordgroup"

	checksByBackend := map[string]acctest.CheckFuncs{
		"nios": {
			Exists:  testAccCheckSharedrecordgroupExistsNIOS,
			Destroy: testAccCheckSharedrecordgroupDestroyNIOS,
		},
	}

	for _, backend := range []string{"nios"} {
		t.Run(backend, func(t *testing.T) {
			acctest.RunListCases(t, resourceType, "dns/sharedrecordgroup/"+backend+"_lists.hcl", checksByBackend)
		})
	}
}
