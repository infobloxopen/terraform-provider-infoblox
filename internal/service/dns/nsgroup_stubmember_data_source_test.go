package dns_test

import (
	"testing"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/acctest"
)

func TestAccNsgroupStubmemberDataSource(t *testing.T) {
	dsType := "infoblox_nsgroup_stubmember"
	resourceType := "infoblox_nsgroup_stubmember"

	checksByBackend := map[string]acctest.CheckFuncs{
		"nios": {
			Exists:  testAccCheckNsgroupStubmemberExistsNIOS,
			Destroy: testAccCheckNsgroupStubmemberDestroyNIOS,
		},
	}

	for _, backend := range []string{"nios"} {
		t.Run(backend, func(t *testing.T) {
			acctest.RunDataSourceCases(t, dsType, resourceType, "dns/nsgroup_stubmember/"+backend+"_datasources.hcl", checksByBackend)
		})
	}
}
