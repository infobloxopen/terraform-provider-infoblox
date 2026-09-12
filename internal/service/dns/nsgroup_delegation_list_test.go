package dns_test

import (
	"testing"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/acctest"
)

func TestAccNsgroupDelegationList(t *testing.T) {
	resourceType := "infoblox_nsgroup_delegation"

	checksByBackend := map[string]acctest.CheckFuncs{
		"nios": {
			Exists:  testAccCheckNsgroupDelegationExistsNIOS,
			Destroy: testAccCheckNsgroupDelegationDestroyNIOS,
		},
	}

	for _, backend := range []string{"nios"} {
		t.Run(backend, func(t *testing.T) {
			acctest.RunListCases(t, resourceType, "dns/nsgroup_delegation/"+backend+"_lists.hcl", checksByBackend)
		})
	}
}
