package grid_test

import (
	"testing"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/acctest"
)

func TestAccServicerestartGroupDataSource(t *testing.T) {
	dsType := "infoblox_servicerestart_group"
	resourceType := "infoblox_servicerestart_group"

	checksByBackend := map[string]acctest.CheckFuncs{
		"nios": {
			Exists:  testAccCheckServicerestartGroupExistsNIOS,
			Destroy: testAccCheckServicerestartGroupDestroyNIOS,
		},
	}

	for _, backend := range []string{"nios"} {
		t.Run(backend, func(t *testing.T) {
			acctest.RunDataSourceCases(t, dsType, resourceType, "grid/servicerestart_group/"+backend+"_datasources.hcl", checksByBackend)
		})
	}
}
