package dns_test

import (
	"testing"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/acctest"
)

func TestAccNsgroupDataSource(t *testing.T) {
	dsType := "infoblox_nsgroup"
	resourceType := "infoblox_nsgroup"

	checksByBackend := map[string]acctest.CheckFuncs{
		"nios": {
			Exists:  testAccCheckNsgroupExistsNIOS,
			Destroy: testAccCheckNsgroupDestroyNIOS,
		},
	}

	for _, backend := range []string{"nios"} {
		t.Run(backend, func(t *testing.T) {
			acctest.RunDataSourceCases(t, dsType, resourceType, "dns/nsgroup/"+backend+"_datasources.hcl", checksByBackend)
		})
	}
}
