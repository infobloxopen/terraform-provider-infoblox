package grid_test

import (
	"testing"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/acctest"
)

func TestAccServicerestartGroupList(t *testing.T) {
	resourceType := "infoblox_servicerestart_group"

	checksByBackend := map[string]acctest.CheckFuncs{
		"nios": {
			Exists:  testAccCheckServicerestartGroupExistsNIOS,
			Destroy: testAccCheckServicerestartGroupDestroyNIOS,
		},
	}

	for _, backend := range []string{"nios"} {
		t.Run(backend, func(t *testing.T) {
			acctest.RunListCases(t, resourceType, "grid/servicerestart_group/"+backend+"_lists.hcl", checksByBackend)
		})
	}
}
