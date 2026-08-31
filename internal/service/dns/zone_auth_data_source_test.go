package dns_test

import (
	"testing"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/acctest"
)

func TestAccZoneAuthDataSource(t *testing.T) {
	dsType := "infoblox_zone_auth"
	resourceType := "infoblox_zone_auth"

	checksByBackend := map[string]acctest.CheckFuncs{
		"nios": {
			Exists:  testAccCheckZoneAuthExistsNIOS,
			Destroy: testAccCheckZoneAuthDestroyNIOS,
		},
		"uddi": {
			Exists:  testAccCheckZoneAuthExistsUDDI,
			Destroy: testAccCheckZoneAuthDestroyUDDI,
		},
	}

	for _, backend := range []string{"nios", "uddi"} {
		t.Run(backend, func(t *testing.T) {
			acctest.RunDataSourceCases(t, dsType, resourceType, "dns/zone_auth/"+backend+"_datasources.hcl", checksByBackend)
		})
	}
}
