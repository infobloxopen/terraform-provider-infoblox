package grid_test

import (
	"testing"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/acctest"
)

func TestAccNatgroupList(t *testing.T) {
	resourceType := "infoblox_natgroup"

	checksByBackend := map[string]acctest.CheckFuncs{
		"nios": {
			Exists:  testAccCheckNatgroupExistsNIOS,
			Destroy: testAccCheckNatgroupDestroyNIOS,
		},
	}

	for _, backend := range []string{"nios"} {
		t.Run(backend, func(t *testing.T) {
			acctest.RunListCases(t, resourceType, "grid/natgroup/"+backend+"_lists.hcl", checksByBackend)
		})
	}
}
