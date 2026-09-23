package dhcp_test

import (
	"testing"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/acctest"
)

func TestAccIpv6filteroptionDataSource(t *testing.T) {
	dsType := "infoblox_ipv6_filteroption"
	resourceType := "infoblox_ipv6_filteroption"

	checksByBackend := map[string]acctest.CheckFuncs{
		"nios": {
			Exists:  testAccCheckIpv6filteroptionExistsNIOS,
			Destroy: testAccCheckIpv6filteroptionDestroyNIOS,
		},
		"uddi": {
			Exists:  testAccCheckIpv6filteroptionExistsUDDI,
			Destroy: testAccCheckIpv6filteroptionDestroyUDDI,
		},
	}

	for _, backend := range []string{"nios", "uddi"} {
		t.Run(backend, func(t *testing.T) {
			acctest.RunDataSourceCases(t, dsType, resourceType, "dhcp/ipv6_filteroption/"+backend+"_datasources.hcl", checksByBackend)
		})
	}
}
