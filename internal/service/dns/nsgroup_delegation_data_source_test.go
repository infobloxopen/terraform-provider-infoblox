package dns_test

import (
	"testing"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/acctest"
)

func TestAccNsgroupDelegationDataSource(t *testing.T) {
	dsType := "infoblox_nsgroup_delegation"
	resourceType := "infoblox_nsgroup_delegation"

	checksByBackend := map[string]acctest.CheckFuncs{
		"nios": {
			Exists:  testAccCheckNsgroupDelegationExistsNIOS,
			Destroy: testAccCheckNsgroupDelegationDestroyNIOS,
		},
	}

	for _, backend := range []string{"nios"} {
		t.Run(backend, func(t *testing.T) {
			acctest.RunDataSourceCases(t, dsType, resourceType, "dns/nsgroup_delegation/"+backend+"_datasources.hcl", checksByBackend)
		})
	}
}
