package dns_test

import (
	"testing"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/acctest"
)

func TestAccZoneStubDataSource(t *testing.T) {
	dsType := "infoblox_zone_stub"
	resourceType := "infoblox_zone_stub"

	checksByBackend := map[string]acctest.CheckFuncs{
		"nios": {
			Exists:  testAccCheckZoneStubExistsNIOS,
			Destroy: testAccCheckZoneStubDestroyNIOS,
		},
	}

	for _, backend := range []string{"nios"} {
		t.Run(backend, func(t *testing.T) {
			acctest.RunDataSourceCases(t, dsType, resourceType, "dns/zone_stub/"+backend+"_datasources.hcl", checksByBackend)
		})
	}
}
