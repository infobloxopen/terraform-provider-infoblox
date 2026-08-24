package ipam_test

import (
	"testing"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/acctest"
)

func TestAccNetworkcontainerDataSource(t *testing.T) {
	dsType := "infoblox_networkcontainer"
	resourceType := "infoblox_networkcontainer"

	checksByBackend := map[string]acctest.CheckFuncs{
		"nios": {
			Exists:  testAccCheckNetworkcontainerExistsNIOS,
			Destroy: testAccCheckNetworkcontainerDestroyNIOS,
		},
		"uddi": {
			Exists:  testAccCheckNetworkcontainerExistsUDDI,
			Destroy: testAccCheckNetworkcontainerDestroyUDDI,
		},
	}

	for _, backend := range []string{"nios", "uddi"} {
		t.Run(backend, func(t *testing.T) {
			acctest.RunDataSourceCases(t, dsType, resourceType, "ipam/networkcontainer/"+backend+"_datasources.tfvars", checksByBackend)
		})
	}
}
