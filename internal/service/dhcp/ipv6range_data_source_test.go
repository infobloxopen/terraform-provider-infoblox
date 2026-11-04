package dhcp_test

import (
	"testing"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/acctest"
)

func TestAccIpv6rangeDataSource(t *testing.T) {
	dsType := "infoblox_ipv6_range"
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
			acctest.RunDataSourceCases(t, dsType, resourceType, "dhcp/ipv6_range/"+backend+"_datasources.hcl", checksByBackend)
		})
	}
}
