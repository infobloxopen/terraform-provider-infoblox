package grid_test

import (
	"testing"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/acctest"
)

func TestAccNatgroupDataSource(t *testing.T) {
	dsType := "infoblox_natgroup"
	resourceType := "infoblox_natgroup"

	checksByBackend := map[string]acctest.CheckFuncs{
		"nios": {
			Exists:  testAccCheckNatgroupExistsNIOS,
			Destroy: testAccCheckNatgroupDestroyNIOS,
		},
	}

	for _, backend := range []string{"nios"} {
		t.Run(backend, func(t *testing.T) {
			acctest.RunDataSourceCases(t, dsType, resourceType, "grid/natgroup/"+backend+"_datasources.hcl", checksByBackend)
		})
	}
}
