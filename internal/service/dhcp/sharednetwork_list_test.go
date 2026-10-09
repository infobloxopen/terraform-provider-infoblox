package dhcp_test

import (
	"testing"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/acctest"
)

func TestAccSharednetworkList(t *testing.T) {
	resourceType := "infoblox_sharednetwork"

	checksByBackend := map[string]acctest.CheckFuncs{
		"nios": {
			Exists:  testAccCheckSharednetworkExistsNIOS,
			Destroy: testAccCheckSharednetworkDestroyNIOS,
		},
	}

	for _, backend := range []string{"nios"} {
		t.Run(backend, func(t *testing.T) {
			acctest.RunListCases(t, resourceType, "dhcp/sharednetwork/"+backend+"_lists.hcl", checksByBackend)
		})
	}
}
