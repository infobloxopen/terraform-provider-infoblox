package grid_test

import (
	"testing"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/acctest"
)

func TestAccDistributionscheduleList(t *testing.T) {
	resourceType := "infoblox_distribution_schedule"

	checksByBackend := map[string]acctest.CheckFuncs{
		"nios": {
			Exists: testAccCheckDistributionscheduleExistsNIOS,
		},
	}

	for _, backend := range []string{"nios"} {
		t.Run(backend, func(t *testing.T) {
			acctest.RunListCases(t, resourceType, "grid/distribution_schedule/"+backend+"_lists.hcl", checksByBackend)
		})
	}
}
