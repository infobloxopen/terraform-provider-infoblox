package grid_test

import (
	"testing"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/acctest"
)

func TestAccDistributionscheduleDataSource(t *testing.T) {
	dsType := "infoblox_distribution_schedule"
	resourceType := "infoblox_distribution_schedule"

	checksByBackend := map[string]acctest.CheckFuncs{
		"nios": {
			Exists: testAccCheckDistributionscheduleExistsNIOS,
		},
	}

	for _, backend := range []string{"nios"} {
		t.Run(backend, func(t *testing.T) {
			acctest.RunDataSourceCases(t, dsType, resourceType, "grid/distribution_schedule/"+backend+"_datasources.hcl", checksByBackend)
		})
	}
}
