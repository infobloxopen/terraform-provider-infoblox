package dns_test

import (
	"testing"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/acctest"
)

func TestAccZoneRpList(t *testing.T) {
	resourceType := "infoblox_zone_rp"

	checksByBackend := map[string]acctest.CheckFuncs{
		"nios": {
			Exists:  testAccCheckZoneRpExistsNIOS,
			Destroy: testAccCheckZoneRpDestroyNIOS,
		},
	}

	for _, backend := range []string{"nios"} {
		t.Run(backend, func(t *testing.T) {
			acctest.RunListCases(t, resourceType, "dns/zone_rp/"+backend+"_lists.hcl", checksByBackend)
		})
	}
}
