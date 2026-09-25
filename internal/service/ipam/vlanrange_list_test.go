package ipam_test

import (
	"testing"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/acctest"
)

func TestAccVlanrangeList(t *testing.T) {
	resourceType := "infoblox_vlan_range"

	checksByBackend := map[string]acctest.CheckFuncs{
		"nios": {
			Exists:  testAccCheckVlanrangeExistsNIOS,
			Destroy: testAccCheckVlanrangeDestroyNIOS,
		},
	}

	for _, backend := range []string{"nios"} {
		t.Run(backend, func(t *testing.T) {
			acctest.RunListCases(t, resourceType, "ipam/vlan_range/"+backend+"_lists.hcl", checksByBackend)
		})
	}
}
