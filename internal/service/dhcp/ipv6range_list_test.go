package dhcp_test

import (
	"testing"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/acctest"
)

func TestAccIpv6rangeList(t *testing.T) {
	resourceType := "infoblox_ipv6_range"

	checksByBackend := map[string]acctest.CheckFuncs{
		"nios": {
			Exists:  testAccCheckIpv6rangeExistsNIOS,
			Destroy: testAccCheckIpv6rangeDestroyNIOS,
		},
		"uddi": {
			Exists:  testAccCheckIpv6rangeExistsUDDI,
			Destroy: testAccCheckIpv6rangeDestroyUDDI,
		},
	}

	for _, backend := range []string{"nios", "uddi"} {
		t.Run(backend, func(t *testing.T) {
			acctest.RunListCases(t, resourceType, "dhcp/ipv6_range/"+backend+"_lists.hcl", checksByBackend)
		})
	}
}
