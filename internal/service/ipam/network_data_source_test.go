package ipam_test

import (
	"testing"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/acctest"
)

func TestAccNetworkDataSource(t *testing.T) {
	dsType := "infoblox_network"
	resourceType := "infoblox_network"

	checksByBackend := map[string]acctest.CheckFuncs{
		"nios": {
			Exists:  testAccCheckNetworkExistsNIOS,
			Destroy: testAccCheckNetworkDestroyNIOS,
		},
		"uddi": {
			Exists:  testAccCheckNetworkExistsUDDI,
			Destroy: testAccCheckNetworkDestroyUDDI,
		},
	}

	for _, backend := range []string{"nios", "uddi"} {
		t.Run(backend, func(t *testing.T) {
			acctest.RunDataSourceCases(t, dsType, resourceType, "ipam/network/"+backend+"_datasources.hcl", checksByBackend)
		})
	}
}
