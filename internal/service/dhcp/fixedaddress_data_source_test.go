package dhcp_test

import (
	"testing"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/acctest"
)

func TestAccFixedaddressDataSource(t *testing.T) {
	dsType := "infoblox_fixed_address"
	resourceType := "infoblox_fixed_address"

	checksByBackend := map[string]acctest.CheckFuncs{
		"nios": {
			Exists:  testAccCheckFixedaddressExistsNIOS,
			Destroy: testAccCheckFixedaddressDestroyNIOS,
		},
		"uddi": {
			Exists:  testAccCheckFixedaddressExistsUDDI,
			Destroy: testAccCheckFixedaddressDestroyUDDI,
		},
	}

	for _, backend := range []string{"nios", "uddi"} {
		t.Run(backend, func(t *testing.T) {
			acctest.RunDataSourceCases(t, dsType, resourceType, "dhcp/fixed_address/"+backend+"_datasources.hcl", checksByBackend)
		})
	}
}
