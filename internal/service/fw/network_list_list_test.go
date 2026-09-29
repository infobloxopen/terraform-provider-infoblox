package fw_test

import (
	"testing"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/acctest"
)

func TestAccNetworkListList(t *testing.T) {
	resourceType := "infoblox_network_list"

	checksByBackend := map[string]acctest.CheckFuncs{
		"uddi": {
			Exists:  testAccCheckNetworkListExistsUDDI,
			Destroy: testAccCheckNetworkListDestroyUDDI,
		},
	}

	for _, backend := range []string{"uddi"} {
		t.Run(backend, func(t *testing.T) {
			acctest.RunListCases(t, resourceType, "fw/network_list/"+backend+"_lists.hcl", checksByBackend)
		})
	}
}
