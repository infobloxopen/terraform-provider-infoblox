package ipam_test

import (
	"testing"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/acctest"
)

func TestAccNetworkviewList(t *testing.T) {
	resourceType := "infoblox_networkview"

	checksByBackend := map[string]acctest.CheckFuncs{
		"nios": {
			Exists:  testAccCheckNetworkviewExistsNIOS,
			Destroy: testAccCheckNetworkviewDestroyNIOS,
		},
		"uddi": {
			Exists:  testAccCheckNetworkviewExistsUDDI,
			Destroy: testAccCheckNetworkviewDestroyUDDI,
		},
	}

	for _, backend := range []string{"nios", "uddi"} {
		t.Run(backend, func(t *testing.T) {
			acctest.RunListCases(t, resourceType, "ipam/networkview/"+backend+"_lists.tfvars", checksByBackend)
		})
	}
}
