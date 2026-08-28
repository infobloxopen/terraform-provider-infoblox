package dns_test

import (
	"testing"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/acctest"
)

func TestAccNsgroupList(t *testing.T) {
	resourceType := "infoblox_nsgroup"

	checksByBackend := map[string]acctest.CheckFuncs{
		"nios": {
			Exists:  testAccCheckNsgroupExistsNIOS,
			Destroy: testAccCheckNsgroupDestroyNIOS,
		},
	}

	for _, backend := range []string{"nios"} {
		t.Run(backend, func(t *testing.T) {
			acctest.RunListCases(t, resourceType, "dns/nsgroup/"+backend+"_lists.hcl", checksByBackend)
		})
	}
}
