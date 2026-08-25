package dns_test

import (
	"testing"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/acctest"
)

func TestAccForwardNsgDataSource(t *testing.T) {
	dsType := "infoblox_forward_nsg"
	resourceType := "infoblox_forward_nsg"

	checksByBackend := map[string]acctest.CheckFuncs{
		"uddi": {
			Exists:  testAccCheckForwardNsgExistsUDDI,
			Destroy: testAccCheckForwardNsgDestroyUDDI,
		},
	}

	for _, backend := range []string{"uddi"} {
		t.Run(backend, func(t *testing.T) {
			acctest.RunDataSourceCases(t, dsType, resourceType, "dns/forward_nsg/"+backend+"_datasources.hcl", checksByBackend)
		})
	}
}
