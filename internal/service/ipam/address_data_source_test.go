package ipam_test

import (
	"testing"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/acctest"
)

func TestAccAddressDataSource(t *testing.T) {
	dsType := "infoblox_address"
	resourceType := "infoblox_address"

	checksByBackend := map[string]acctest.CheckFuncs{
		"uddi": {
			Exists:  testAccCheckAddressExistsUDDI,
			Destroy: testAccCheckAddressDestroyUDDI,
		},
	}

	for _, backend := range []string{"uddi"} {
		t.Run(backend, func(t *testing.T) {
			acctest.RunDataSourceCases(t, dsType, resourceType, "ipam/address/"+backend+"_datasources.hcl", checksByBackend)
		})
	}
}
