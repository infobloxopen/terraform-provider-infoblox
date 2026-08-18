package ipam_test

import (
	"testing"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/acctest"
)

func TestAccNetworkcontainerList(t *testing.T) {
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
			acctest.RunListCases(t, resourceType, "ipam/networkcontainer/"+backend+"_lists.tfvars", checksByBackend)
		})
	}
}
