package grid_test

import (
	"testing"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/acctest"
)

func TestAccUpgradegroupList(t *testing.T) {
	resourceType := "infoblox_upgradegroup"

	checksByBackend := map[string]acctest.CheckFuncs{
		"nios": {
			Exists:  testAccCheckUpgradegroupExistsNIOS,
			Destroy: testAccCheckUpgradegroupDestroyNIOS,
		},
	}

	for _, backend := range []string{"nios"} {
		t.Run(backend, func(t *testing.T) {
			acctest.RunListCases(t, resourceType, "grid/upgradegroup/"+backend+"_lists.hcl", checksByBackend)
		})
	}
}
