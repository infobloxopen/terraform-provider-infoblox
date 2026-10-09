package grid_test

import (
	"testing"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/acctest"
)

func TestAccUpgradegroupDataSource(t *testing.T) {
	dsType := "infoblox_upgrade_group"
	resourceType := "infoblox_upgrade_group"

	checksByBackend := map[string]acctest.CheckFuncs{
		"nios": {
			Exists:  testAccCheckUpgradegroupExistsNIOS,
			Destroy: testAccCheckUpgradegroupDestroyNIOS,
		},
	}

	for _, backend := range []string{"nios"} {
		t.Run(backend, func(t *testing.T) {
			acctest.RunDataSourceCases(t, dsType, resourceType, "grid/upgrade_group/"+backend+"_datasources.hcl", checksByBackend)
		})
	}
}
