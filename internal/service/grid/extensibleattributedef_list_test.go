package grid_test

import (
	"testing"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/acctest"
)

func TestAccExtensibleattributedefList(t *testing.T) {
	resourceType := "infoblox_extensibleattributedef"

	checksByBackend := map[string]acctest.CheckFuncs{
		"nios": {
			Exists:  testAccCheckExtensibleattributedefExistsNIOS,
			Destroy: testAccCheckExtensibleattributedefDestroyNIOS,
		},
	}

	for _, backend := range []string{"nios"} {
		t.Run(backend, func(t *testing.T) {
			acctest.RunListCases(t, resourceType, "grid/extensibleattributedef/"+backend+"_lists.hcl", checksByBackend)
		})
	}
}
