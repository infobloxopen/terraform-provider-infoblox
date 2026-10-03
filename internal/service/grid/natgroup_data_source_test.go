package grid_test

import (
	"testing"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/acctest"
)

func TestAccNatgroupDataSource(t *testing.T) {
	dsType := "infoblox_nat_group"
	resourceType := "infoblox_nat_group"

	checksByBackend := map[string]acctest.CheckFuncs{
		"nios": {
			Exists:  testAccCheckNatgroupExistsNIOS,
			Destroy: testAccCheckNatgroupDestroyNIOS,
		},
	}

	for _, backend := range []string{"nios"} {
		t.Run(backend, func(t *testing.T) {
			acctest.RunDataSourceCases(t, dsType, resourceType, "grid/nat_group/"+backend+"_datasources.hcl", checksByBackend)
		})
	}
}
