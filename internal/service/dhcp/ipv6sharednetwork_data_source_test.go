package dhcp_test

import (
	"testing"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/acctest"
)

func TestAccIpv6sharednetworkDataSource(t *testing.T) {
	dsType := "infoblox_ipv6_sharednetwork"
	resourceType := "infoblox_ipv6_sharednetwork"

	checksByBackend := map[string]acctest.CheckFuncs{
		"nios": {
			Exists:  testAccCheckIpv6sharednetworkExistsNIOS,
			Destroy: testAccCheckIpv6sharednetworkDestroyNIOS,
		},
	}

	for _, backend := range []string{"nios"} {
		t.Run(backend, func(t *testing.T) {
			acctest.RunDataSourceCases(t, dsType, resourceType, "dhcp/ipv6_sharednetwork/"+backend+"_datasources.hcl", checksByBackend)
		})
	}
}
