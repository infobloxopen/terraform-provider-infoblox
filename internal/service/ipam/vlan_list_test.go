package ipam_test

import (
	"testing"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/acctest"
)

func TestAccVlanList(t *testing.T) {
	resourceType := "infoblox_vlan"

	checksByBackend := map[string]acctest.CheckFuncs{
		"nios": {
			Exists:  testAccCheckVlanExistsNIOS,
			Destroy: testAccCheckVlanDestroyNIOS,
		},
	}

	for _, backend := range []string{"nios"} {
		t.Run(backend, func(t *testing.T) {
			acctest.RunListCases(t, resourceType, "ipam/vlan/"+backend+"_lists.hcl", checksByBackend)
		})
	}
}
