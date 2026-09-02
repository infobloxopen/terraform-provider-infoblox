package grid_test

import (
	"testing"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/acctest"
)

func TestAccUpgradegroupDataSource(t *testing.T) {
	dsType := "infoblox_upgradegroup"
	resourceType := "infoblox_upgradegroup"

	checksByBackend := map[string]acctest.CheckFuncs{
		"nios": {
			Exists:  testAccCheckUpgradegroupExistsNIOS,
			Destroy: testAccCheckUpgradegroupDestroyNIOS,
		},
	}

	for _, backend := range []string{"nios"} {
		t.Run(backend, func(t *testing.T) {
			acctest.RunDataSourceCases(t, dsType, resourceType, "grid/upgradegroup/"+backend+"_datasources.hcl", checksByBackend)
		})
	}
}
