package dns_test

import (
	"testing"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/acctest"
)

func TestAccZoneDelegatedDataSource(t *testing.T) {
	dsType := "infoblox_zone_delegated"
	resourceType := "infoblox_zone_delegated"

	checksByBackend := map[string]acctest.CheckFuncs{
		"nios": {
			Exists:  testAccCheckZoneDelegatedExistsNIOS,
			Destroy: testAccCheckZoneDelegatedDestroyNIOS,
		},
		"uddi": {
			Exists:  testAccCheckZoneDelegatedExistsUDDI,
			Destroy: testAccCheckZoneDelegatedDestroyUDDI,
		},
	}

	for _, backend := range []string{"nios", "uddi"} {
		t.Run(backend, func(t *testing.T) {
			acctest.RunDataSourceCases(t, dsType, resourceType, "dns/zone_delegated/"+backend+"_datasources.hcl", checksByBackend)
		})
	}
}
