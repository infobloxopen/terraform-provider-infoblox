package grid_test

import (
	"testing"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/acctest"
)

func TestAccDistributionscheduleList(t *testing.T) {
	resourceType := "infoblox_distributionschedule"

	checksByBackend := map[string]acctest.CheckFuncs{
		"nios": {
			Exists:  testAccCheckDistributionscheduleExistsNIOS,
			Destroy: testAccCheckDistributionscheduleDestroyNIOS,
		},
	}

	for _, backend := range []string{"nios"} {
		t.Run(backend, func(t *testing.T) {
			acctest.RunListCases(t, resourceType, "grid/distributionschedule/"+backend+"_lists.hcl", checksByBackend)
		})
	}
}
