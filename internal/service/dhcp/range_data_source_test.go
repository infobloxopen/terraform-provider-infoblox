package dhcp_test

import (
	"testing"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/acctest"
)

func TestAccRangeDataSource(t *testing.T) {
	dsType := "infoblox_range"
	resourceType := "infoblox_range"

	checksByBackend := map[string]acctest.CheckFuncs{
		"nios": {
			Exists:  testAccCheckRangeExistsNIOS,
			Destroy: testAccCheckRangeDestroyNIOS,
		},
		"uddi": {
			Exists:  testAccCheckRangeExistsUDDI,
			Destroy: testAccCheckRangeDestroyUDDI,
		},
	}

	for _, backend := range []string{"nios", "uddi"} {
		t.Run(backend, func(t *testing.T) {
			acctest.RunDataSourceCases(t, dsType, resourceType, "dhcp/range/"+backend+"_datasources.hcl", checksByBackend)
		})
	}
}
