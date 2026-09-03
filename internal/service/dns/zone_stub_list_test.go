package dns_test

import (
	"testing"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/acctest"
)

func TestAccZoneStubList(t *testing.T) {
	resourceType := "infoblox_zone_stub"

	checksByBackend := map[string]acctest.CheckFuncs{
		"nios": {
			Exists:  testAccCheckZoneStubExistsNIOS,
			Destroy: testAccCheckZoneStubDestroyNIOS,
		},
	}

	for _, backend := range []string{"nios"} {
		t.Run(backend, func(t *testing.T) {
			acctest.RunListCases(t, resourceType, "dns/zone_stub/"+backend+"_lists.hcl", checksByBackend)
		})
	}
}
