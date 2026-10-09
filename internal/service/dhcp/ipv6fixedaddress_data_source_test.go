package dhcp_test

import (
	"testing"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/acctest"
)

func TestAccIpv6fixedaddressDataSource(t *testing.T) {
	dsType := "infoblox_ipv6_fixed_address"
	resourceType := "infoblox_ipv6_fixed_address"

	checksByBackend := map[string]acctest.CheckFuncs{
		"nios": {
			Exists:  testAccCheckIpv6fixedaddressExistsNIOS,
			Destroy: testAccCheckIpv6fixedaddressDestroyNIOS,
		},
		"uddi": {
			Exists:  testAccCheckIpv6fixedaddressExistsUDDI,
			Destroy: testAccCheckIpv6fixedaddressDestroyUDDI,
		},
	}

	for _, backend := range []string{"nios", "uddi"} {
		t.Run(backend, func(t *testing.T) {
			acctest.RunDataSourceCases(t, dsType, resourceType, "dhcp/ipv6_fixed_address/"+backend+"_datasources.hcl", checksByBackend)
		})
	}
}
