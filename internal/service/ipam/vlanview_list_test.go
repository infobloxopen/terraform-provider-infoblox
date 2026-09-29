package ipam_test

import (
	"testing"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/acctest"
)

func TestAccVlanviewList(t *testing.T) {
	resourceType := "infoblox_vlan_view"

	checksByBackend := map[string]acctest.CheckFuncs{
		"nios": {
			Exists:  testAccCheckVlanviewExistsNIOS,
			Destroy: testAccCheckVlanviewDestroyNIOS,
		},
	}

	for _, backend := range []string{"nios"} {
		t.Run(backend, func(t *testing.T) {
			acctest.RunListCases(t, resourceType, "ipam/vlan_view/"+backend+"_lists.hcl", checksByBackend)
		})
	}
}
