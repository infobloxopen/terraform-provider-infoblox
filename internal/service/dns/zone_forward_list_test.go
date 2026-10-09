package dns_test

import (
	"testing"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/acctest"
)

func TestAccZoneForwardList(t *testing.T) {
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
			acctest.RunListCases(t, resourceType, "dns/zone_forward/"+backend+"_lists.hcl", checksByBackend)
		})
	}
}
