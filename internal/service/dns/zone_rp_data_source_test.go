package dns_test

import (
	"testing"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/acctest"
)

func TestAccZoneRpDataSource(t *testing.T) {
	dsType := "infoblox_zone_rp"
	resourceType := "infoblox_zone_rp"

	checksByBackend := map[string]acctest.CheckFuncs{
		"nios": {
			Exists:  testAccCheckZoneRpExistsNIOS,
			Destroy: testAccCheckZoneRpDestroyNIOS,
		},
	}

	for _, backend := range []string{"nios"} {
		t.Run(backend, func(t *testing.T) {
			acctest.RunDataSourceCases(t, dsType, resourceType, "dns/zone_rp/"+backend+"_datasources.hcl", checksByBackend)
		})
	}
}
