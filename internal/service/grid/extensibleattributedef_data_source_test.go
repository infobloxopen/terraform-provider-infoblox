package grid_test

import (
	"testing"

	"github.com/infobloxopen/terraform-provider-infoblox/internal/acctest"
)

func TestAccExtensibleattributedefDataSource(t *testing.T) {
	dsType := "infoblox_extensibleattributedef"
	resourceType := "infoblox_extensibleattributedef"

	checksByBackend := map[string]acctest.CheckFuncs{
		"nios": {
			Exists:  testAccCheckExtensibleattributedefExistsNIOS,
			Destroy: testAccCheckExtensibleattributedefDestroyNIOS,
		},
	}

	for _, backend := range []string{"nios"} {
		t.Run(backend, func(t *testing.T) {
			acctest.RunDataSourceCases(t, dsType, resourceType, "grid/extensibleattributedef/"+backend+"_datasources.hcl", checksByBackend)
		})
	}
}
