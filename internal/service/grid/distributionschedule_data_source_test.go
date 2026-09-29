package grid_test

import (
	"testing"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/acctest"
)

func TestAccDistributionscheduleDataSource(t *testing.T) {
	dsType := "infoblox_distributionschedule"
	resourceType := "infoblox_distributionschedule"

	checksByBackend := map[string]acctest.CheckFuncs{
		"nios": {
			Exists:  testAccCheckDistributionscheduleExistsNIOS,
			Destroy: testAccCheckDistributionscheduleDestroyNIOS,
		},
	}

	for _, backend := range []string{"nios"} {
		t.Run(backend, func(t *testing.T) {
			acctest.RunDataSourceCases(t, dsType, resourceType, "grid/distributionschedule/"+backend+"_datasources.hcl", checksByBackend)
		})
	}
}
