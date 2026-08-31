package dns_test

import (
	"testing"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/acctest"
)

func TestAccZoneForwardDataSource(t *testing.T) {
	dsType := "infoblox_zone_forward"
	resourceType := "infoblox_zone_forward"

	checksByBackend := map[string]acctest.CheckFuncs{
		"nios": {
			Exists:  testAccCheckZoneForwardExistsNIOS,
			Destroy: testAccCheckZoneForwardDestroyNIOS,
		},
		"uddi": {
			Exists:  testAccCheckZoneForwardExistsUDDI,
			Destroy: testAccCheckZoneForwardDestroyUDDI,
		},
	}

	for _, backend := range []string{"nios", "uddi"} {
		t.Run(backend, func(t *testing.T) {
			acctest.RunDataSourceCases(t, dsType, resourceType, "dns/zone_forward/"+backend+"_datasources.hcl", checksByBackend)
		})
	}
}
