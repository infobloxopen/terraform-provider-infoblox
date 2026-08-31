package ipam_test

import (
	"testing"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/acctest"
)

func TestAccNetworkList(t *testing.T) {
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
			acctest.RunListCases(t, resourceType, "ipam/network/"+backend+"_lists.hcl", checksByBackend)
		})
	}
}
